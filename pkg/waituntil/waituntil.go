package waituntil

import (
	"log"
)

func Run(opts CmdOpts) {
	printOptions(opts)

	// Create new job
	job, err := BuildJob(opts)
	if err != nil {
		panic(err)
	}

	// Run job command
	err = job.Run()
	if err != nil {
		panic(err)

	}
}

func printOptions(opts CmdOpts) {
	log.Println("You specified options:")
	log.Printf("  Command:    %s", opts.Command)
	log.Printf("  Interval:   %d", opts.Interval)
	log.Printf("  Retries:    %d", opts.MaxRetries)
	log.Printf("  OK Message: %s", opts.OkText)
	log.Printf("  NG Message: %s", opts.NgText)
}
