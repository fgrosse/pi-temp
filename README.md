# pi-temp

[![Build Status](https://secure.travis-ci.org/fgrosse/pi-temp.png?branch=master)](http://travis-ci.org/fgrosse/pi-temp)
[![License](https://img.shields.io/badge/license-MIT-4183c4.svg)](https://github.com/fgrosse/pi-temp/blob/master/LICENSE)

Small utility to export the CPU temperature of a Raspberry Pi as Prometheus metric.

[![Build Status][travis_badge]][travis]

### Usage

```shell
# Cross compile and upload to your raspberry:
$ GOARCH=arm64 go build && scp pi-temp raspberry:/usr/local/bin/pi-temp
pi-temp                                                    100% 9619KB   9.4MB/s   00:01

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

### License

© Friedrich Große 2018, distributed under [MIT License](LICENSE).
