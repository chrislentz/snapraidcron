package commands

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/chrislentz/snapraidcron/utilities"
)

func Scrub(snapraidBin string) string {
	output := "\n"

	output = utilities.AppendLabelToOutput(output, "START: SCRUB Command")

	startTime := time.Now()

	cmd := exec.Command("/bin/bash", "snapraid scrub")
	// cmd := exec.Command("/bin/bash", "./scrub.sh")
	stdout, err := cmd.Output()

	if err != nil {
		output = utilities.AppendToOutput(output, "Command Error: "+fmt.Sprintf("%s", err))
	} else {
		diffOutput := string(stdout)

		if strings.Contains(diffOutput, "Nothing to do") {
			output = utilities.AppendToOutput(output, "Nothing historical needs scrubbing.")
		} else {
			output = utilities.AppendToOutput(output, "Historical scrub completed successfully.")
		}
	}

	executedIn := time.Since(startTime)

	output = utilities.AppendLabelToOutput(output, fmt.Sprintf("END: SCRUB Command (took %f seconds)", executedIn.Seconds()))

	return output
}
