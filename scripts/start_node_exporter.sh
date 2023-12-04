#!/bin/bash
nohub /opt/promonitor &
node_exporter --collector.vmstat --collector.tcpstat --collector.processes