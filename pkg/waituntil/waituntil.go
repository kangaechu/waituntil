package waituntil

func Run(opts CmdOpts) {
	opts.Print()

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
