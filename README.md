# pi-temp

Small utility to export the CPU temperature of a Raspberry Pi as Prometheus metric.

### Usage

```
Usage of pi-temp:
  -addr string
    	The address to listen on for HTTP requests. (default ":8080")
  -debug
    	Enable debug output.
  -file string
    	Path to the thermal zone file. (default "/sys/class/thermal/thermal_zone0/temp")
  -interval duration
    	The interval at which the temperature is checked. (default 10s)
```

### Why?

Node exporter doesn't contain this metric for my Pi and this was super easy
to do by myself.
