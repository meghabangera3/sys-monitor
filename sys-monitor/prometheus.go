package main

import (
	"fmt"
	"net/http"
)

func startPrometheusServer(port string) {
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		cpu, err := getCPUUsage()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		mem, err := getMemoryUsage()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		disk, err := getDiskUsage()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "# HELP sys_cpu_usage_percent CPU usage percentage\n")
		fmt.Fprintf(w, "# TYPE sys_cpu_usage_percent gauge\n")
		fmt.Fprintf(w, "sys_cpu_usage_percent %.2f\n\n", cpu)

		fmt.Fprintf(w, "# HELP sys_memory_usage_percent Memory usage percentage\n")
		fmt.Fprintf(w, "# TYPE sys_memory_usage_percent gauge\n")
		fmt.Fprintf(w, "sys_memory_usage_percent %.2f\n\n", mem)

		fmt.Fprintf(w, "# HELP sys_disk_usage_percent Disk usage percentage\n")
		fmt.Fprintf(w, "# TYPE sys_disk_usage_percent gauge\n")
		fmt.Fprintf(w, "sys_disk_usage_percent %.2f\n", disk)
	})

	go http.ListenAndServe(":"+port, nil)
	fmt.Printf("   Prometheus metrics at http://localhost:%s/metrics\n\n", port)
}