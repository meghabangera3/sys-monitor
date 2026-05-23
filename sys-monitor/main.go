package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	interval := flag.Int("interval", 5, "Check interval in seconds")
	threshold := flag.Float64("threshold", 80.0, "CPU alert threshold percentage")
	flag.Parse()

	fmt.Println("🖥  System Monitor Starting...")
	fmt.Printf("   Interval: %ds | CPU Alert Threshold: %.0f%%\n", *interval, *threshold)
	startPrometheusServer("9090")

	for {
		cpu, err := getCPUUsage()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading CPU: %v\n", err)
		}

		mem, err := getMemoryUsage()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading memory: %v\n", err)
		}

		disk, err := getDiskUsage()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading disk: %v\n", err)
		}

		printMetrics(cpu, mem, disk)

		if cpu > *threshold {
			fmt.Printf("⚠️  ALERT: CPU usage %.2f%% exceeds threshold %.0f%%\n", cpu, *threshold)
		}

		time.Sleep(time.Duration(*interval) * time.Second)
	}
}

func printMetrics(cpu, mem, disk float64) {
	fmt.Printf("[%s]\n", time.Now().Format("15:04:05"))
	fmt.Printf("  CPU  : %.2f%%\n", cpu)
	fmt.Printf("  MEM  : %.2f%%\n", mem)
	fmt.Printf("  DISK : %.2f%%\n\n", disk)
}
