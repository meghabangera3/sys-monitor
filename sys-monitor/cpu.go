package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getCPUUsage() (float64, error) {
	read := func() (idle, total uint64, err error) {
		data, err := os.ReadFile("/proc/stat")
		if err != nil {
			return 0, 0, fmt.Errorf("cannot read /proc/stat: %w", err)
		}

		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if !strings.HasPrefix(line, "cpu ") {
				continue
			}

			fields := strings.Fields(line)[1:]
			var vals []uint64
			for _, f := range fields {
				v, err := strconv.ParseUint(f, 10, 64)
				if err != nil {
					continue
				}
				vals = append(vals, v)
			}

			if len(vals) < 4 {
				return 0, 0, fmt.Errorf("unexpected /proc/stat format")
			}

			idle = vals[3]
			for _, v := range vals {
				total += v
			}
			return idle, total, nil
		}
		return 0, 0, fmt.Errorf("cpu line not found")
	}

	idle1, total1, err := read()
	if err != nil {
		return 0, err
	}

	time.Sleep(200 * time.Millisecond)

	idle2, total2, err := read()
	if err != nil {
		return 0, err
	}

	idleDelta := float64(idle2 - idle1)
	totalDelta := float64(total2 - total1)

	if totalDelta == 0 {
		return 0, nil
	}

	return (1 - idleDelta/totalDelta) * 100, nil
}