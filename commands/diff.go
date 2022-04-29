package commands

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/chrislentz/snapraidcron/utilities"
)

func Diff() (string, bool) {
	output := ""
	diffDetected := false

	output = utilities.AppendLabelToOutput(output, "START: DIFF Command")

	startTime := time.Now()

	cmd := exec.Command("/bin/bash", "snapraid diff")
	// cmd := exec.Command("/bin/bash", "./diff.sh")
	stdout, err := cmd.Output()

	if err != nil {
		output = utilities.AppendToOutput(output, fmt.Sprintf("Command Error: %s", err))
	} else {
		diffOutput := string(stdout)

		if !strings.Contains(diffOutput, "There are differences!") {
			output = utilities.AppendToOutput(output, "No differences detected.")
		} else {
			diffDetected = true

			equal := utilities.SingleMatchAndReplaceRegex(diffOutput, " equal", " +\\d+", true)
			removed := utilities.SingleMatchAndReplaceRegex(diffOutput, " removed", " +\\d+", true)
			added := utilities.SingleMatchAndReplaceRegex(diffOutput, " added", " +\\d+", true)
			moved := utilities.SingleMatchAndReplaceRegex(diffOutput, " moved", " +\\d+", true)
			copied := utilities.SingleMatchAndReplaceRegex(diffOutput, " copied", " +\\d+", true)
			updated := utilities.SingleMatchAndReplaceRegex(diffOutput, " updated", " +\\d+", true)

			length := 0

			for _, s := range [6]string{equal, removed, added, moved, copied, updated} {
				if len(s) > length {
					length = len(s)
				}
			}

			output = utilities.AppendToOutput(output, fmt.Sprintf("%s No Change", utilities.AddTrailingSpacesToString(length, equal)))
			output = utilities.AppendToOutput(output, fmt.Sprintf("%s No Removed", utilities.AddTrailingSpacesToString(length, removed)))
			output = utilities.AppendToOutput(output, fmt.Sprintf("%s No Added", utilities.AddTrailingSpacesToString(length, added)))
			output = utilities.AppendToOutput(output, fmt.Sprintf("%s No Moved", utilities.AddTrailingSpacesToString(length, moved)))
			output = utilities.AppendToOutput(output, fmt.Sprintf("%s No Copied", utilities.AddTrailingSpacesToString(length, copied)))
			output = utilities.AppendToOutput(output, fmt.Sprintf("%s No Updated", utilities.AddTrailingSpacesToString(length, updated)))
		}
	}

	executedIn := time.Since(startTime)

	output = utilities.AppendLabelToOutput(output, fmt.Sprintf("END: DIFF Command (took %f seconds)", executedIn.Seconds()))

	return output, diffDetected
}
