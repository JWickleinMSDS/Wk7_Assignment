package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"testing"
	"time"
)

const epsilon = 1e-4

// Use this to check if the results are nearly equal (meaning account for rounding)
func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestLinearRegressionResultsFromFile(t *testing.T) {
	// Expected values from the results of Anscombe in R and Python
	// These values are hard coded that I pulled from R and Python by running
	// them in R Studio and a Jupyter notebook, respectively.
	expectedValues := []struct {
		Slope     float64
		Intercept float64
	}{
		{0.5001, 3.0001},
		{0.500, 3.001},
		{0.4997, 3.0025},
		{0.4999, 3.0017},
	}

	// Read the content of the output file produced from anscombe_data
	content, err := os.ReadFile("Go_linreg_output.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Split file content into lines
	lines := bufio.NewScanner(strings.NewReader(string(content)))
	var actualValues []struct {
		Slope     float64
		Intercept float64
	}

	// Iterate over lines and parse the values
	for lines.Scan() {
		var slope, intercept float64
		n, _ := fmt.Sscanf(lines.Text(), "%f %f", &slope, &intercept)
		if n == 2 {
			actualValues = append(actualValues, struct {
				Slope     float64
				Intercept float64
			}{slope, intercept})
		}
	}

	// Compare the results from the output file (i.e. anscombe_data.go results) and what R and Python produced.
	for i, actual := range actualValues {
		expected := expectedValues[i]
		if !almostEqual(actual.Slope, expected.Slope) || !almostEqual(actual.Intercept, expected.Intercept) {
			t.Errorf("Test case %d failed. Got Slope=%.4f, Intercept=%.4f. Expected Slope=%.4f, Intercept=%.4f",
				i+1, actual.Slope, actual.Intercept, expected.Slope, expected.Intercept)
		}
	}
}

func readExecutionTime(filename string) (time.Duration, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	// Parse the execution time from the file content
	executionTimeStr := strings.TrimSpace(string(content))

	// Check if the duration string contains a unit, if not, assume milliseconds
	if !strings.ContainsAny(executionTimeStr, "smh") {
		executionTimeStr += "ms"
	}

	executionTime, err := time.ParseDuration(executionTimeStr)
	if err != nil {
		return 0, err
	}

	return executionTime, nil
}

func compareExecutionTimes(t *testing.T, goTime, pythonTime, rTime time.Duration) {
	// Compare execution times
	fmt.Printf("Go Execution Time: %s\n", goTime)
	fmt.Printf("Python Execution Time: %s\n", pythonTime)
	fmt.Printf("R Execution Time: %s\n", rTime)

	// Write the comparison to a file
	file, err := os.Create("Execution_Time_Comparisons.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	fmt.Fprintf(writer, "Go Execution Time: %s\n", goTime)
	fmt.Fprintf(writer, "Python Execution Time: %s\n", pythonTime)
	fmt.Fprintf(writer, "R Execution Time: %s\n", rTime)
}

func TestExecutionTimeComparisons(t *testing.T) {
	// Read execution times from files
	goTime, err := readExecutionTime("Go_Execution_Time.txt")
	if err != nil {
		t.Fatal(err)
	}

	pythonTime, err := readExecutionTime("Python_execution_time.txt")
	if err != nil {
		t.Fatal(err)
	}

	rTime, err := readExecutionTime("R_execution_time.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Compare and write to file
	compareExecutionTimes(t, goTime, pythonTime, rTime)
}
