package commands

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/chrislentz/snapraidcron/utilities"
)

func ScrubNew() string {
	output := "\n\n\n"

	output = utilities.AppendLabelToOutput(output, "START: SCRUB NEW Command")

	startTime := time.Now()

	cmd := exec.Command("/bin/bash", "snapraid scrub -p new")
	// cmd := exec.Command("/bin/bash", "./scrub.sh")
	stdout, err := cmd.Output()

	if err != nil {
		output = utilities.AppendToOutput(output, "Command Error: "+fmt.Sprintf("%s", err))
	} else {
		diffOutput := string(stdout)

		if strings.Contains(diffOutput, "Nothing to do") {
			output = utilities.AppendToOutput(output, "Nothing new needs scrubbing.")
		} else {
			output = utilities.AppendToOutput(output, "Scrub new completed successfully.")
		}
	}

	executedIn := time.Since(startTime)

	output = utilities.AppendLabelToOutput(output, fmt.Sprintf("END: SCRUB NEW Command (took %f seconds)", executedIn.Seconds()))

	return output
}
