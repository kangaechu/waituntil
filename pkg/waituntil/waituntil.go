package waituntil

import "log"

func Run(opts CmdOpts) {
	log.Println("You specified options:")
	log.Printf("  Command:    %s", opts.Command)
	log.Printf("  Interval:  %5d", opts.Interval)
	log.Printf("  Retries:   %5d", opts.MaxRetries)
	log.Printf("  OK Message: %s", opts.OkText)
	log.Printf("  NG Message: %s", opts.NgText)
}
