main package
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
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

// create a new function called "igloo" that prints odd numbers from 1 to 10.
func igloo() {
	// add your code here
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

	