# Rowenta robot vacuum cleaner - Prometheus Exporter

Exports prometheus metrics on some metrics exposed from the Rowenta robot vacuum cleaner.

Tested on Rowenta X-PLORER SERIE 80.

## Exposed metrics

```
# HELP rowenta_battery_level Battery Percentage
# TYPE rowenta_battery_level gauge
rowenta_battery_level 0.89
# HELP rowenta_battery_voltage_volts Battery Voltage in volts
# TYPE rowenta_battery_voltage_volts gauge
rowenta_battery_voltage_volts 16.576
# HELP rowenta_charging Charging state
# TYPE rowenta_charging gauge
rowenta_charging 0
# HELP rowenta_clean_time_seconds Time spent cleaning time in seconds
# TYPE rowenta_clean_time_seconds counter
rowenta_clean_time_seconds 646020
# HELP rowenta_distance_driven_meters Distance travelled in meters
# TYPE rowenta_distance_driven_meters counter
rowenta_distance_driven_meters 88578.16
# HELP rowenta_runs_total Total number of runs
# TYPE rowenta_runs_total counter
rowenta_runs_total 516
# HELP rowenta_uptime_seconds Robot uptime in seconds
# TYPE rowenta_uptime_seconds gauge
rowenta_uptime_seconds 1.253267660219385e+06
```

## Running with Docker

`docker run -ti --rm -p 9101:9101 -e ROWENTA_ENDPOINT=http://rowenta.internal:8080 chadiem/rowenta-robot-vacuum-exporter:master`

The endpoint can be found by doing a nmap on the network.

It might be interesting later on to auto-discover it. Contributions are welcome!