package main

import (
	"fmt"
	"syscall"
)

func getDiskUsage() (float64, error) {
	var stat syscall.Statfs_t

	err := syscall.Statfs("/", &stat)
	if err != nil {
		return 0, fmt.Errorf("cannot read disk usage: %w", err)
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bfree * uint64(stat.Bsize)
	used := total - free

	if total == 0 {
		return 0, nil
	}

	return float64(used) / float64(total) * 100, nil
}