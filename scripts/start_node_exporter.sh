#!/bin/bash
nohup /opt/promonitor &
node_exporter --collector.vmstat --collector.tcpstat --collector.processes