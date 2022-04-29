package commands

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/chrislentz/snapraidcron/utilities"
)

func Sync() (string, bool) {
	output := "\n\n\n"
	dataSynced := false

	output = utilities.AppendLabelToOutput(output, "START: SYNC Command")

	startTime := time.Now()

	cmd := exec.Command("/bin/bash", "snapraid sync")
	// cmd := exec.Command("/bin/bash", "./sync.sh")
	stdout, err := cmd.Output()

	if err != nil {
		output = utilities.AppendToOutput(output, "Command Error: "+fmt.Sprintf("%s", err))
	} else {
		diffOutput := string(stdout)

		if strings.Contains(diffOutput, "Nothing to do") {
			output = utilities.AppendToOutput(output, "Nothing to sync.")
		} else {
			dataSynced = true
			output = utilities.AppendToOutput(output, "Synced completed successfully.")
		}
	}

	executedIn := time.Since(startTime)

	output = utilities.AppendLabelToOutput(output, fmt.Sprintf("END: SYNC Command (took %f seconds)", executedIn.Seconds()))

	return output, dataSynced
}
