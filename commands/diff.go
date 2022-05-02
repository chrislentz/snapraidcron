package commands

import (
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/chrislentz/snapraidcron/utilities"
)

func Diff(snapraidBin string) (string, bool) {
	output := ""
	diffDetected := false

	output = utilities.AppendLabelToOutput(output, "START: DIFF Command")

	startTime := time.Now()

	cmd := exec.Command("/bin/bash", "snapraid diff")
	// cmd := exec.Command("/bin/bash", "./diff.sh")
	stdout, err := cmd.Output()
	diffCommandOutput, err := cmd.CombinedOutput()
	diffExitCode := 0

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			diffExitCode = exitError.ExitCode()
		}

		if diffExitCode != 2 {
			output = utilities.AppendToOutput(output, fmt.Sprintf("Command Error: %s", err))
		}
	}

	if err == nil || diffExitCode == 2 {
		diffOutput := string(diffCommandOutput)

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

			output = utilities.AppendToOutput(output, fmt.Sprintf("%s Changed", utilities.AddTrailingSpacesToString(length, equal)))
			output = utilities.AppendToOutput(output, fmt.Sprintf("%s Removed", utilities.AddTrailingSpacesToString(length, removed)))
			output = utilities.AppendToOutput(output, fmt.Sprintf("%s Added", utilities.AddTrailingSpacesToString(length, added)))
			output = utilities.AppendToOutput(output, fmt.Sprintf("%s Moved", utilities.AddTrailingSpacesToString(length, moved)))
			output = utilities.AppendToOutput(output, fmt.Sprintf("%s Copied", utilities.AddTrailingSpacesToString(length, copied)))
			output = utilities.AppendToOutput(output, fmt.Sprintf("%s Updated", utilities.AddTrailingSpacesToString(length, updated)))
		}
	}

	executedIn := time.Since(startTime)

	output = utilities.AppendLabelToOutput(output, fmt.Sprintf("END: DIFF Command (took %f seconds)", executedIn.Seconds()))

	return output, diffDetected
}
