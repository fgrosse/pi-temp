# pi-temp

[![Build Status](https://secure.travis-ci.org/fgrosse/pi-temp.png?branch=master)](http://travis-ci.org/fgrosse/pi-temp)
[![License](https://img.shields.io/badge/license-MIT-4183c4.svg)](https://github.com/fgrosse/pi-temp/blob/master/LICENSE)

Small utility to export the CPU temperature of a Raspberry Pi as Prometheus metric.

### Usage

```shell
# Cross compile and upload to your raspberry:
$ GOARCH=arm64 go build && scp pi-temp raspberry:/usr/local/bin/pi-temp
pi-temp                                                    100% 9619KB   9.4MB/s   00:01

# Log into your Pi:
$ ssh raspberry

# Helpful help output is helpful:
Usage of pi-temp:
  -addr string
    	The address to listen on for HTTP requests. (default ":8080")
  -debug
    	Enable debug output.
  -file string
    	Path to the thermal zone file. (default "/sys/class/thermal/thermal_zone0/temp")
  -interval duration
    	The interval at which the temperature is checked. (default 10s)

# Run it:
$ pi-temp -debug
Checking temperature every 10s from "/sys/class/thermal/thermal_zone0/temp"
CPU temperature: 55.306°C
CPU temperature: 54.768°C
CPU temperature: 54.768°C
CPU temperature: 54.768°C
…
```

### Why?

Node exporter doesn't contain this metric for my Pi and this was super easy
to do by myself.

### Systemd

I use the [`pi-temp.service`](pi-temp.service) file to install the cross compiled
binary as simple systemd service into my Raspberry Pi like this:

```shell
# Upload systemd unit file:
$ scp pi-temp.service raspberry:/etc/systemd/system/pi-temp.service

# Log into your Pi:
$ ssh raspberry

# Enable service to start on boot:
$ systemctl enable pi-temp.service

# Start pi-temp service now:
$ systemctl start pi-temp

# Check its running:
$ systemctl status pi-temp
● pi-temp.service - Prometheus Temperature Monitor
   Loaded: loaded (/etc/systemd/system/pi-temp.service; enabled; vendor preset: disabled)
   Active: active (running) since Sun 2018-08-12 23:20:18 CEST; 1s ago
 Main PID: 7886 (pi-temp)
    Tasks: 7 (limit: 1053)
   Memory: 1.3M
   CGroup: /system.slice/pi-temp.service
           └─7886 /usr/local/bin/pi-temp -interval=10s -addr=:9101

Aug 12 23:24:49 rasputin systemd[1]: Started Prometheus Temperature Monitor.
Aug 12 23:24:49 rasputin pi-temp[7953]: Starting pi-temp web server at ":9101"
Aug 12 23:24:49 rasputin pi-temp[7953]: If you want to see more verbose log run with -debug
```

### License

© Friedrich Große 2018, distributed under [MIT License](LICENSE).
