package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMemoryUsage() (float64, error) {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return 0, fmt.Errorf("cannot read /proc/meminfo: %w", err)
	}

	values := make(map[string]uint64)
	for _, line := range strings.Split(string(data), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		key := strings.TrimSuffix(fields[0], ":")
		val, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			continue
		}
		values[key] = val
	}

	total, ok1 := values["MemTotal"]
	available, ok2 := values["MemAvailable"]

	if !ok1 || !ok2 {
		return 0, fmt.Errorf("could not find memory values in /proc/meminfo")
	}

	if total == 0 {
		return 0, nil
	}

	used := total - available
	return float64(used) / float64(total) * 100, nil
}