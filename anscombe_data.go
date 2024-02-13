package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

// do the linear regression awesomeness
func linearRegression(x, y []float64) (float64, float64) {
	// Get the mean for x and y
	meanX, meanY := calculateMean(x), calculateMean(y)

	// Calculate slope (i.e. the x value)
	numerator, denominator := 0.0, 0.0
	for i := 0; i < len(x); i++ {
		numerator += (x[i] - meanX) * (y[i] - meanY)
		denominator += math.Pow(x[i]-meanX, 2)
	}
	slope := numerator / denominator

	// Calculate intercept (i.e. the y value)
	intercept := meanY - slope*meanX

	return slope, intercept
}

func calculateMean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

func main() {
	// Create a new file to store the results
	file, err := os.Create("Go_linreg_output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Some other stuff I found you need to do to write the file efficiently
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Record the start time
	startTime := time.Now()

	// bring in Anscombe's quartet dataset
	x := [][]float64{
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5},
		{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8},
	}

	y := [][]float64{
		{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68},
		{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74},
		{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73},
		{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89},
	}

	// Calculate linear regression coefficients for the four pairings
	for i := 0; i < len(x); i++ {
		slope, intercept := linearRegression(x[i], y[i])

		// Print the results to console
		fmt.Printf("For x%d and y%d:\n", i+1, i+1)
		fmt.Printf("Slope: %.4f\n", slope)
		fmt.Printf("Intercept: %.4f\n", intercept)
		fmt.Println("-------------")

		// Write the results to the file for inclusion in a test
		fmt.Fprintln(writer, slope, intercept)
	}

	// Record the end time
	endTime := time.Now()

	// Calculate and print the execution time
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Execution Time: %s\n", executionTime)

	// Save execution time to a text file
	executionTimeFile, err := os.Create("Go_Execution_Time.txt")
	if err != nil {
		fmt.Println("Error creating execution time file:", err)
		return
	}
	defer executionTimeFile.Close()

	// Write execution time to the file (without "ms" suffix)
	fmt.Fprintf(executionTimeFile, "%v\n", executionTime.Milliseconds())

	fmt.Println("Results written to Go_linreg_output.txt and Go_Execution_time.txt files")
}
