package main

import (
	"fmt"
	"log"

	"github.com/chrislentz/snapraidcron/commands"
	"github.com/chrislentz/snapraidcron/utilities"
)

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	output := ""

	config, err := utilities.LoadConfigFile()

	if err != nil {
		output = fmt.Sprintln(err)
	}

	diffOutput := ""
	diffDetected := false

	syncOutput := ""
	dataSynced := false

	scrubNewOutput := ""
	scrubOutput := ""

	// Run DIFF command
	diffOutput, diffDetected = commands.Diff(config.SnapraidBin)
	output = output + diffOutput

	if diffDetected {
		// Run SYNC command
		syncOutput, dataSynced = commands.Sync(config.SnapraidBin)
		output = output + syncOutput

		if dataSynced {
			// Run SCRUB NEW command
			scrubNewOutput = commands.ScrubNew(config.SnapraidBin)
			output = output + scrubNewOutput
		}
	}

	// Run SCRUB  command
	scrubOutput = commands.Scrub(config.SnapraidBin)
	output = output + scrubOutput

	// SMTP Email
	var subject string

	if dataSynced {
		subject = "SnapRAID Cron - Sync Completed Successfully"
	} else {
		subject = "SnapRAID Cron - Nothing Synced"
	}

	utilities.SendSmtp(config.Smtp.Host, config.Smtp.Port, config.Smtp.Username, config.Smtp.Password, config.Smtp.To, config.Smtp.From, subject, output)
}
