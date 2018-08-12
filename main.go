package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	addr     = flag.String("addr", ":8080", "The address to listen on for HTTP requests.")
	path     = flag.String("file", "/sys/class/thermal/thermal_zone0/temp", "Path to the thermal zone file.")
	interval = flag.Duration("interval", 10*time.Second, "The interval at which the temperature is checked.")
	verbose  = flag.Bool("debug", false, "Enable debug output.")
)

func main() {
	log.SetFlags(0)
	ctx := Context()
	flag.Parse()

	s := Server(*addr)
	temperature := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "pitemp",
		Name:      "temperature_celsius",
		Help:      "The CPU temperature of this Raspberry Pi in degrees Celsius.",
	})
	prometheus.MustRegister(temperature)

	debug := log.New(ioutil.Discard, "", 0)
	if *verbose {
		debug.SetOutput(os.Stderr)
	} else {
		log.Printf("Starting pi-temp web server at %q", *addr)
		log.Println("If you want to see more verbose log run with -debug")
	}

	go MonitorTemperature(ctx, debug, *path, *interval, temperature)
	go func() { log.Fatal(s.ListenAndServe()) }()

	// Run until we are interrupted
	<-ctx.Done()
	s.Shutdown(Context())
}

// Context returns a context that is cancelled automatically when a SIGINT,
// SIGQUIT or SIGTERM signal is received.
func Context() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		select {
		case <-sig:
			cancel()
		}
	}()

	return ctx
}

// Server creates the HTTP server that is used by Prometheus to scrape the temperature metric.
func Server(addr string) *http.Server {
	h := http.NewServeMux()
	h.Handle("/metrics", promhttp.Handler())

	return &http.Server{
		Addr:         addr,
		Handler:      h,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     log.New(os.Stderr, "HTTP: ", 0),
	}
}

// MonitorTemperature periodically reads the temperature from the path (e.g.
// "/sys/class/thermal/thermal_zone0/temp") and updates the given prometheus gauge.
func MonitorTemperature(ctx context.Context, debug *log.Logger, path string, interval time.Duration, temperature prometheus.Gauge) {
	debug.Printf("Checking temperature every %v from %q", interval, path)

	for {
		raw, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("Failed to read temperature from %q: %v", path, err)
		}

		cpuTempStr := strings.TrimSpace(string(raw))
		cpuTempInt, err := strconv.Atoi(cpuTempStr) // e.g. 55306
		if err != nil {
			log.Fatalf("%q does not contain an integer: %v", path, err)
		}

		cpuTemp := float64(cpuTempInt) / 1000.0
		debug.Printf("CPU temperature: %.3f°C", cpuTemp)
		temperature.Set(cpuTemp)

		select {
		case <-time.After(interval):
			continue
		case <-ctx.Done():
			return
		}
	}
}
