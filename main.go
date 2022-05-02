package main

import (
	"fmt"
	"log"

	"github.com/chrislentz/snapraidcron/commands"

	"github.com/joho/godotenv"
)

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Load .env variables
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	output := ""
	diffOutput := ""
	diffDetected := false

	syncOutput := ""
	dataSynced := false

	scrubNewOutput := ""
	scrubOutput := ""

	// Run DIFF command
	diffOutput, diffDetected = commands.Diff()
	output = output + diffOutput

	if diffDetected {
		// Run SYNC command
		syncOutput, dataSynced = commands.Sync()
		output = output + syncOutput

		if dataSynced {
			// Run SCRUB NEW command
			scrubNewOutput = commands.ScrubNew()
			output = output + scrubNewOutput
		}
	}

	// Run SCRUB  command
	scrubOutput = commands.Scrub()
	output = output + scrubOutput

	// Print the output
	fmt.Println(output)
}
