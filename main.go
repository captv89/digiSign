package main

import (
	"os"
)

// Init function
func init() {
	//	Check if output directory exists
	if _, err := os.Stat("output"); os.IsNotExist(err) {
		// create output directory
		err := os.Mkdir("output", 0755)
		if err != nil {
			panic(err)
		}
	}
}

// main function
func main() {

}
