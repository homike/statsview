package main

import (
	"math/rand"
	"statsview"
)

func generateValues() []float64 {
	values := make([]float64, 0)
	for i := 0; i < 7; i++ {
		values = append(values, float64(rand.Intn(300)))
	}
	return values
}

func main() {
	statsview.SetConfiguration(
		statsview.WithAddr("192.168.0.1:8088"),
		statsview.WithInterval(10000))

	views := []statsview.Viewer{
		statsview.NewBasicViewer("Goroutine", nil, func() []float64 {
			return generateValues()
		}),
		statsview.NewBasicViewer("QPS", []string{"QPS", "MaxQPS"}, func() []float64 {
			return generateValues()
		}),
	}
	statsview.Startup(views)
}
