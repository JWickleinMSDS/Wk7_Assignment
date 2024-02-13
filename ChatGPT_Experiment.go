package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func runCommand(cmdName string, args ...string) (time.Duration, error) {
	start := time.Now()
	cmd := exec.Command(cmdName, args...)
	err := cmd.Run()
	elapsed := time.Since(start)
	return elapsed, err
}

func main() {
	files := []string{"R_execution_time.txt", "Python_execution_time.txt", "Go_Execution_Time.txt"} // Updated file names

	for _, file := range files {
		fmt.Printf("Executing %s...\n", file)
		elapsed, err := runCommand("bash", file) // Use "cmd" on Windows
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error executing %s: %v\n", file, err)
			continue
		}
		fmt.Printf("%s executed in %s\n", file, elapsed)
	}
}
