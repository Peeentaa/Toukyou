package analysis

import (
	"fmt"
	"os"
	"os/exec"
)

func RunAnalysis(analysis string) error {
	// Get the working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return err
	}

	// Define the path to the C++ executable
	exePath := wd + "/analysis/algorithms/algorithms.exe"
	dbPath := wd + "/data/db/data.db"

	// Check if the executable exists
	if _, err := os.Stat(exePath); os.IsNotExist(err) {
		fmt.Println("Executable does not exist:", err)
		return err
	}

	// Execute the C++ executable
	cmd := exec.Command(exePath, analysis, dbPath)

	// Run the command and capture the output
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running command:", err)
		return err
	}
	fmt.Println("Output:", string(out))

	return nil
}
