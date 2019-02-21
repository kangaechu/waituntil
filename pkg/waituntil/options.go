package waituntil

type CmdOpts struct {
	Command    string `short:"c" long:"command" description:"Check command that this tool repeatedly isseud." required:"true"`
	Interval   uint   `short:"i" long:"interval" description:"Check command retry interval (sec) default:60"`
	MaxRetries uint   `short:"r" long:"retries" description:"Check command number of retries default:20"`
	OkText     string `short:"o" long:"ok" description:"Text recognized as OK (normal end)" required:"true"`
	NgText     string `short:"n" long:"ng" description:"Text recognized as NG (with error)" required:"true"`
}
