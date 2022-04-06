package metric

import "time"

type Collector interface {
	Collect()
}

func ExecuteCollector(collectors ...Collector) {
	for _, collector := range collectors {
		c := collector
		go func() {
			for {
				c.Collect()
				time.Sleep(1 * time.Second)
			}
		}()
	}
}
