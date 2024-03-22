#!/usr/bin/env bash
set -eu
set -o pipefail

url="https://raw.githubusercontent.com/honahuku/k8s-hands-on/main/manifest/k8s-scale-test/condition.txt"

while true; do
    condition=$(curl -s -H "Cache-Control: no-cache" -H "Pragma: no-cache" $url)
	current_time=$(date '+%Y-%m-%d %H:%M:%S')

	if [ "$condition" = "0" ]; then
		echo "[$current_time] Condition is 0. waiting for 10 seconds..."
		sleep 10
	elif [ "$condition" = "1" ]; then
		echo "[$current_time] Condition is 1. Running stress-ng..."
		stress-ng --cpu 8 --timeout 10s
	else
		echo "Unexpected condition value: $condition"
	fi
done
