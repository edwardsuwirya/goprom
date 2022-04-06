package metric

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"runtime"
)

type memoryUsageCollector struct {
	vec *prometheus.GaugeVec
}

func (m *memoryUsageCollector) Collect() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	m.vec.With(prometheus.Labels{"stats": "Alloc"}).Set(bToMb(ms.Alloc))
	m.vec.With(prometheus.Labels{"stats": "Sys"}).Set(bToMb(ms.Sys))
	//metric.MemoryUsage.With(prometheus.Labels{"stats": "Garbage Collector"}).Set(float64(m.NumGC))
	fmt.Printf("Alloc = %v MiB", bToMb(ms.Alloc))
	//fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB\n", bToMb(ms.Sys))
	//fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func NewMemoryUsageCollector(vec *prometheus.GaugeVec) Collector {
	return &memoryUsageCollector{
		vec: vec,
	}
}
