>大家好，我是Tia。这是一个简单版的基于prometheus和grafana的Go程序的监控系统。内容包括：
> 基于prometheus获取服务的指标
> docker部署prometheus和grafana
> grafana展示dashboard
> 等等

# 监控系统架构
![img.png](img.png)

# 命令操作
## 启动prometheus和grafana
```shell
# 编译server
sh build.sh
# 启动容器
docker-compose -f docker-compose.yml up -d
# 关闭
docker-compose -f docker-compose.yml down
# 查看服务启动情况
docker ps
```
# 监控指标
```
# HELP go_gc_duration_seconds A summary of the pause duration of garbage collection cycles.
# TYPE go_gc_duration_seconds summary
go_gc_duration_seconds{quantile="0"} 8.4167e-05
go_gc_duration_seconds{quantile="0.25"} 0.00090025
go_gc_duration_seconds{quantile="0.5"} 0.001295749
go_gc_duration_seconds{quantile="0.75"} 0.002141583
go_gc_duration_seconds{quantile="1"} 0.002418042
go_gc_duration_seconds_sum 0.013329457
go_gc_duration_seconds_count 10
# HELP go_goroutines Number of goroutines that currently exist.
# TYPE go_goroutines gauge
go_goroutines 9
# HELP go_info Information about the Go environment.
# TYPE go_info gauge
go_info{version="go1.21.4"} 1
# HELP go_memstats_alloc_bytes Number of bytes allocated and still in use.
# TYPE go_memstats_alloc_bytes gauge
go_memstats_alloc_bytes 4.659776e+06
# HELP go_memstats_alloc_bytes_total Total number of bytes allocated, even if freed.
# TYPE go_memstats_alloc_bytes_total counter
go_memstats_alloc_bytes_total 2.1296488e+07
# HELP go_memstats_buck_hash_sys_bytes Number of bytes used by the profiling bucket hash table.
# TYPE go_memstats_buck_hash_sys_bytes gauge
go_memstats_buck_hash_sys_bytes 5343
# HELP go_memstats_frees_total Total number of frees.
# TYPE go_memstats_frees_total counter
go_memstats_frees_total 54537
# HELP go_memstats_gc_sys_bytes Number of bytes used for garbage collection system metadata.
# TYPE go_memstats_gc_sys_bytes gauge
go_memstats_gc_sys_bytes 3.998248e+06
# HELP go_memstats_heap_alloc_bytes Number of heap bytes allocated and still in use.
# TYPE go_memstats_heap_alloc_bytes gauge
go_memstats_heap_alloc_bytes 4.659776e+06
# HELP go_memstats_heap_idle_bytes Number of heap bytes waiting to be used.
# TYPE go_memstats_heap_idle_bytes gauge
go_memstats_heap_idle_bytes 4.923392e+06
# HELP go_memstats_heap_inuse_bytes Number of heap bytes that are in use.
# TYPE go_memstats_heap_inuse_bytes gauge
go_memstats_heap_inuse_bytes 7.135232e+06
# HELP go_memstats_heap_objects Number of allocated objects.
# TYPE go_memstats_heap_objects gauge
go_memstats_heap_objects 11988
# HELP go_memstats_heap_released_bytes Number of heap bytes released to OS.
# TYPE go_memstats_heap_released_bytes gauge
go_memstats_heap_released_bytes 2.768896e+06
# HELP go_memstats_heap_sys_bytes Number of heap bytes obtained from system.
# TYPE go_memstats_heap_sys_bytes gauge
go_memstats_heap_sys_bytes 1.2058624e+07
# HELP go_memstats_last_gc_time_seconds Number of seconds since 1970 of last garbage collection.
# TYPE go_memstats_last_gc_time_seconds gauge
go_memstats_last_gc_time_seconds 1.7017397910890296e+09
# HELP go_memstats_lookups_total Total number of pointer lookups.
# TYPE go_memstats_lookups_total counter
go_memstats_lookups_total 0
# HELP go_memstats_mallocs_total Total number of mallocs.
# TYPE go_memstats_mallocs_total counter
go_memstats_mallocs_total 66525
# HELP go_memstats_mcache_inuse_bytes Number of bytes in use by mcache structures.
# TYPE go_memstats_mcache_inuse_bytes gauge
go_memstats_mcache_inuse_bytes 4800
# HELP go_memstats_mcache_sys_bytes Number of bytes used for mcache structures obtained from system.
# TYPE go_memstats_mcache_sys_bytes gauge
go_memstats_mcache_sys_bytes 15600
# HELP go_memstats_mspan_inuse_bytes Number of bytes in use by mspan structures.
# TYPE go_memstats_mspan_inuse_bytes gauge
go_memstats_mspan_inuse_bytes 130368
# HELP go_memstats_mspan_sys_bytes Number of bytes used for mspan structures obtained from system.
# TYPE go_memstats_mspan_sys_bytes gauge
go_memstats_mspan_sys_bytes 130368
# HELP go_memstats_next_gc_bytes Number of heap bytes when next garbage collection will take place.
# TYPE go_memstats_next_gc_bytes gauge
go_memstats_next_gc_bytes 9.642248e+06
# HELP go_memstats_other_sys_bytes Number of bytes used for other system allocations.
# TYPE go_memstats_other_sys_bytes gauge
go_memstats_other_sys_bytes 1.067737e+06
# HELP go_memstats_stack_inuse_bytes Number of bytes in use by the stack allocator.
# TYPE go_memstats_stack_inuse_bytes gauge
go_memstats_stack_inuse_bytes 524288
# HELP go_memstats_stack_sys_bytes Number of bytes obtained from system for stack allocator.
# TYPE go_memstats_stack_sys_bytes gauge
go_memstats_stack_sys_bytes 524288
# HELP go_memstats_sys_bytes Number of bytes obtained from system.
# TYPE go_memstats_sys_bytes gauge
go_memstats_sys_bytes 1.7800208e+07
# HELP go_threads Number of OS threads created.
# TYPE go_threads gauge
go_threads 7
# HELP process_cpu_seconds_total Total user and system CPU time spent in seconds.
# TYPE process_cpu_seconds_total counter
process_cpu_seconds_total 0.67
# HELP process_max_fds Maximum number of open file descriptors.
# TYPE process_max_fds gauge
process_max_fds 1.048576e+06
# HELP process_open_fds Number of open file descriptors.
# TYPE process_open_fds gauge
process_open_fds 11
# HELP process_resident_memory_bytes Resident memory size in bytes.
# TYPE process_resident_memory_bytes gauge
process_resident_memory_bytes 1.7788928e+07
# HELP process_start_time_seconds Start time of the process since unix epoch in seconds.
# TYPE process_start_time_seconds gauge
process_start_time_seconds 1.7017386998e+09
# HELP process_virtual_memory_bytes Virtual memory size in bytes.
# TYPE process_virtual_memory_bytes gauge
process_virtual_memory_bytes 1.269518336e+09
# HELP process_virtual_memory_max_bytes Maximum amount of virtual memory available in bytes.
# TYPE process_virtual_memory_max_bytes gauge
process_virtual_memory_max_bytes 1.8446744073709552e+19
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 37
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
# HELP server_handle_request_seconds 
# TYPE server_handle_request_seconds histogram
server_handle_request_seconds_bucket{method="GET",path="/metric",status="404",type="http",le="0.005"} 4
server_handle_request_seconds_bucket{method="GET",path="/metric",status="404",type="http",le="0.01"} 4
server_handle_request_seconds_bucket{method="GET",path="/metric",status="404",type="http",le="0.025"} 4
server_handle_request_seconds_bucket{method="GET",path="/metric",status="404",type="http",le="0.05"} 4
server_handle_request_seconds_bucket{method="GET",path="/metric",status="404",type="http",le="0.1"} 4
server_handle_request_seconds_bucket{method="GET",path="/metric",status="404",type="http",le="0.25"} 4
server_handle_request_seconds_bucket{method="GET",path="/metric",status="404",type="http",le="0.5"} 4
server_handle_request_seconds_bucket{method="GET",path="/metric",status="404",type="http",le="1"} 4
server_handle_request_seconds_bucket{method="GET",path="/metric",status="404",type="http",le="2.5"} 4
server_handle_request_seconds_bucket{method="GET",path="/metric",status="404",type="http",le="5"} 4
server_handle_request_seconds_bucket{method="GET",path="/metric",status="404",type="http",le="10"} 4
server_handle_request_seconds_bucket{method="GET",path="/metric",status="404",type="http",le="+Inf"} 4
server_handle_request_seconds_sum{method="GET",path="/metric",status="404",type="http"} 0.0006793330000000001
server_handle_request_seconds_count{method="GET",path="/metric",status="404",type="http"} 4
server_handle_request_seconds_bucket{method="GET",path="/metrics",status="200",type="http",le="0.005"} 21
server_handle_request_seconds_bucket{method="GET",path="/metrics",status="200",type="http",le="0.01"} 34
server_handle_request_seconds_bucket{method="GET",path="/metrics",status="200",type="http",le="0.025"} 36
server_handle_request_seconds_bucket{method="GET",path="/metrics",status="200",type="http",le="0.05"} 37
server_handle_request_seconds_bucket{method="GET",path="/metrics",status="200",type="http",le="0.1"} 37
server_handle_request_seconds_bucket{method="GET",path="/metrics",status="200",type="http",le="0.25"} 37
server_handle_request_seconds_bucket{method="GET",path="/metrics",status="200",type="http",le="0.5"} 37
server_handle_request_seconds_bucket{method="GET",path="/metrics",status="200",type="http",le="1"} 37
server_handle_request_seconds_bucket{method="GET",path="/metrics",status="200",type="http",le="2.5"} 37
server_handle_request_seconds_bucket{method="GET",path="/metrics",status="200",type="http",le="5"} 37
server_handle_request_seconds_bucket{method="GET",path="/metrics",status="200",type="http",le="10"} 37
server_handle_request_seconds_bucket{method="GET",path="/metrics",status="200",type="http",le="+Inf"} 37
server_handle_request_seconds_sum{method="GET",path="/metrics",status="200",type="http"} 0.21039695999999997
server_handle_request_seconds_count{method="GET",path="/metrics",status="200",type="http"} 37
# HELP server_handle_request_status 
# TYPE server_handle_request_status gauge
server_handle_request_status{method="GET",path="/metric",status="404",type="http"} 4
server_handle_request_status{method="GET",path="/metrics",status="200",type="http"} 37
# HELP server_handle_request_total Number of server requests received in total
# TYPE server_handle_request_total counter
server_handle_request_total{method="GET",path="/metric",type="http"} 4
server_handle_request_total{method="GET",path="/metrics",type="http"} 38
```