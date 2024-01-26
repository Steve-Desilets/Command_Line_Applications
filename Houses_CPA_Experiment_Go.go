package main

// import the packages that will be needed for this project
import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	//	"reflect"
)



func main() {
// Execute 100 iterations of the experimental trials
// Append the runtime for each experimental trial to a slice called runtime_slice

	// Set up variables
	var runtimeSlice []int64
	loopCounter := 0

	// Create output file to which we will write the outputs
	file, err := os.Create("housesOutputGo.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	for i := 0; i < 100; i++ {
		startTime := time.Now()

		// Open the CSV file
		csvFile, err := os.Open("housesInput.csv")
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
		defer csvFile.Close()

		// Read the CSV content
		reader := csv.NewReader(csvFile)
		records, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Error reading CSV:", err)
			return
		}

		// Assume the first row contains headers
		headers := records[0]
		data := records[1:]

		// Initialize maps to store summary statistics for each column
		// Maps from column name to statistics
		columnStatistics := make(map[string]map[string]float64)

		// Loop through each column
		for colIndex, header := range headers {
			// Initialize statistics for the current column
			colStats := make(map[string]float64)
			colValues := make([]float64, 0)

			// Loop through each row in the current column
			for _, row := range data {
				// Parse the value as a float64
				val, err := strconv.ParseFloat(row[colIndex], 64)
				if err != nil {
					fmt.Printf("Error parsing value in column %s: %v\n", header, err)
					continue
				}

				// Store the value in the current column's values slice
				colValues = append(colValues, val)
			}

			// Calculate summary statistics for the current column
			colStats["mean"] = calculateMean(colValues)
			colStats["min"] = calculateMin(colValues)
			colStats["max"] = calculateMax(colValues)

			// Store the statistics for the current column
			columnStatistics[header] = colStats
		}

		// Print the summary statistics
		for header, stats := range columnStatistics {
			fmt.Fprintf(file, "Column: %s\n", header)
			fmt.Fprintf(file, "  Mean: %f\n", stats["mean"])
			fmt.Fprintf(file, "  Min: %f\n", stats["min"])
			fmt.Fprintf(file, "  Max: %f\n", stats["max"])
			fmt.Fprintf(file, "\n ")
		}

		endTime := time.Now()
		executionTime := endTime.Sub(startTime)
		runtimeSlice = append(runtimeSlice, executionTime.Microseconds())

		loopCounter += i
	}
	printSlice(runtimeSlice)

	// Calculate the total and average runtime across each of the experimental trials. Print these statistics to the output .txt file
	var runtimeSum int64 = 0

	for i := 0; i < 100; i++ {
		runtimeSum += runtimeSlice[i]
	}
	runtimeSumString := strconv.FormatInt(runtimeSum, 10)

	avgRuntime := (float64(runtimeSum)) / (float64(100))
	avgRuntimeString := fmt.Sprintf("%f", avgRuntime)

	fmt.Fprintf(file, "Summary Statistics For Experimental Trial Runtimes \n")

	fmt.Fprintf(file, "Runtime Sum in Microseconds \n")
	fmt.Fprintf(file, runtimeSumString)

	fmt.Fprintf(file,"\nAverage Trial Runtime in Microseconds \n")
	fmt.Fprintf(file, avgRuntimeString)

}

// Define a function that will print out the slice of runtimes from each experimental trial
func printSlice(s []int64) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


// Define a function to calculate the mean of a slice of float64 values
func calculateMean(values []float64) float64 {
	sum := 0.0
	for _, v := range values {
		sum += v
	}
	return sum / float64(len(values))
}


// Define a function to calculate the minimum value in a slice of float64 values
func calculateMin(values []float64) float64 {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

// Define a function to calculate the maximum value in a slice of float64 values
func calculateMax(values []float64) float64 {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}




