package waituntil

import "log"

type CmdOpts struct {
	Command    string `short:"c" long:"command" description:"Check command that this tool repeatedly isseud." required:"true"`
	Interval   uint   `short:"i" long:"interval" description:"Check command retry interval (sec) default:60"`
	MaxRetries uint   `short:"r" long:"retries" description:"Check command number of retries default:20"`
	OkText     string `short:"o" long:"ok" description:"Text recognized as OK (normal end)" required:"true"`
	NgText     string `short:"n" long:"ng" description:"Text recognized as NG (with error)" required:"true"`
}

func (opts CmdOpts) Print() {
	log.Println("options:")
	log.Printf("  Command:    %s", opts.Command)
	log.Printf("  Interval:   %d", opts.Interval)
	log.Printf("  Retries:    %d", opts.MaxRetries)
	log.Printf("  OK Message: %s", opts.OkText)
	log.Printf("  NG Message: %s", opts.NgText)
	log.Println()
}
