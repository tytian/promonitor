# 介绍
>这是一个利用prometheus为gin框架提供的一些通用的metric监控工具

# customize metrics list
```shell
# HELP service_http_request_count_total Total number of HTTP requests made
# TYPE service_http_request_count_total counter
service_http_request_count_total{endpoint="",method="GET",origin="",server="localhost:8090",status="200"} 2
# HELP service_http_request_duration_seconds HTTP request latencies in seconds
# TYPE service_http_request_duration_seconds histogram
service_http_request_duration_seconds_bucket{endpoint="",method="GET",origin="",server="localhost:8090",status="200",le="0.005"} 2
service_http_request_duration_seconds_bucket{endpoint="",method="GET",origin="",server="localhost:8090",status="200",le="0.01"} 2
service_http_request_duration_seconds_bucket{endpoint="",method="GET",origin="",server="localhost:8090",status="200",le="0.025"} 2
service_http_request_duration_seconds_bucket{endpoint="",method="GET",origin="",server="localhost:8090",status="200",le="0.05"} 2
service_http_request_duration_seconds_bucket{endpoint="",method="GET",origin="",server="localhost:8090",status="200",le="0.1"} 2
service_http_request_duration_seconds_bucket{endpoint="",method="GET",origin="",server="localhost:8090",status="200",le="0.25"} 2
service_http_request_duration_seconds_bucket{endpoint="",method="GET",origin="",server="localhost:8090",status="200",le="0.5"} 2
service_http_request_duration_seconds_bucket{endpoint="",method="GET",origin="",server="localhost:8090",status="200",le="1"} 2
service_http_request_duration_seconds_bucket{endpoint="",method="GET",origin="",server="localhost:8090",status="200",le="2.5"} 2
service_http_request_duration_seconds_bucket{endpoint="",method="GET",origin="",server="localhost:8090",status="200",le="5"} 2
service_http_request_duration_seconds_bucket{endpoint="",method="GET",origin="",server="localhost:8090",status="200",le="10"} 2
service_http_request_duration_seconds_bucket{endpoint="",method="GET",origin="",server="localhost:8090",status="200",le="+Inf"} 2
service_http_request_duration_seconds_sum{endpoint="",method="GET",origin="",server="localhost:8090",status="200"} 0.001763333
service_http_request_duration_seconds_count{endpoint="",method="GET",origin="",server="localhost:8090",status="200"} 2
# HELP service_http_request_size_bytes HTTP request sizes in bytes.
# TYPE service_http_request_size_bytes summary
service_http_request_size_bytes_sum{endpoint="",method="GET",origin="",server="localhost:8090",status="200"} 1618
service_http_request_size_bytes_count{endpoint="",method="GET",origin="",server="localhost:8090",status="200"} 2
# HELP service_http_response_size_bytes HTTP request sizes in bytes.
# TYPE service_http_response_size_bytes summary
service_http_response_size_bytes_sum{endpoint="",method="GET",origin="",server="localhost:8090",status="200"} 1146
service_http_response_size_bytes_count{endpoint="",method="GET",origin="",server="localhost:8090",status="200"} 2
# HELP service_uptime HTTP service uptime
# TYPE service_uptime counter
service_uptime 33
# HELP process_uptime_seconds HTTP service uptime seconds
# TYPE process_uptime_seconds gauge
process_uptime_seconds 33
```
# go metrics list
```shell
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 2.8709e-05
go_gc_duration_seconds{quantile="0.25"} 2.8709e-05
go_gc_duration_seconds{quantile="0.5"} 2.8709e-05
go_gc_duration_seconds{quantile="0.75"} 2.8709e-05
go_gc_duration_seconds{quantile="1"} 2.8709e-05
go_gc_duration_seconds_sum 2.8709e-05
go_gc_duration_seconds_count 1
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 9
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.21.4"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 2.480128e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 3.889624e+06
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 1.446061e+06
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 15604
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 3.400216e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 2.480128e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 2.94912e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 4.816896e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 8708
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 2.162688e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 7.766016e+06
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.701782993041445e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 24312
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 9600
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 15600
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 111216
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 114072
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 4.194304e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 1.677243e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 622592
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 622592
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 1.50418e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 11
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 1
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```