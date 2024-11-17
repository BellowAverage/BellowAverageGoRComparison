package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"gonum.org/v1/gonum/stat"
)

// Function to calculate the median
func median(data []float64) float64 {
	n := len(data)
	sort.Float64s(data)
	if n%2 == 0 {
		return (data[n/2-1] + data[n/2]) / 2
	}
	return data[n/2]
}

// Function to perform bootstrapping
func bootstrapMedian(data []float64, resamples int) ([]float64, float64, float64) {
	n := len(data)
	medians := make([]float64, resamples)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < resamples; i++ {
		resample := make([]float64, n)
		for j := 0; j < n; j++ {
			resample[j] = data[rand.Intn(n)]
		}
		medians[i] = median(resample)
	}

	// Calculate mean and standard deviation of the bootstrap medians
	mean, std := stat.MeanStdDev(medians, nil)

	return medians, mean, std
}

func main() {
	// Generate sample data
	rand.Seed(123)
	sampleData := make([]float64, 1000)
	for i := range sampleData {
		sampleData[i] = rand.NormFloat64()*10 + 50
	}

	// Perform bootstrapping
	resamples := 2000
	medians, mean, std := bootstrapMedian(sampleData, resamples)

	// Calculate 95% confidence interval
	sort.Float64s(medians)
	lower := medians[int(0.025*float64(resamples))]
	upper := medians[int(0.975*float64(resamples))]

	fmt.Printf("Bootstrap Mean: %.2f\n", mean)
	fmt.Printf("Bootstrap Std Dev: %.2f\n", std)
	fmt.Printf("95%% Confidence Interval: [%.2f, %.2f]\n", lower, upper)
}
