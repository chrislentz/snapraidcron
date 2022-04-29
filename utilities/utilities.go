package utilities

import (
	"fmt"
	"regexp"
	"strings"
)

func SingleMatchAndReplaceRegex(s string, replaceString string, regex string, reverse bool) string {
	var r *regexp.Regexp
	if !reverse {
		r = regexp.MustCompile(fmt.Sprintf("(?m)(%s)(%s)+", replaceString, regex))
	} else {
		r = regexp.MustCompile(fmt.Sprintf("(?m)(%s)+(%s)", regex, replaceString))
	}

	match := r.FindString(s)

	return strings.TrimSpace(strings.Replace(match, replaceString, "", 1))
}

func AppendToOutput(output string, stringToAdd string) string {
	return output + "\n" + stringToAdd
}

func AppendLabelToOutput(output string, label string) string {
	if output != "" {
		output = output + "\n"
	}

	output = output + "--------------------------------------------------" + "\n"
	output = output + label + "\n"
	output = output + "--------------------------------------------------"

	return output
}

func AddTrailingSpacesToString(length int, s string) string {
	if len(s) < length {
		return s + strings.Repeat(" ", length-len(s))
	}

	return s
}
