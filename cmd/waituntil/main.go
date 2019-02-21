package main

//HOGE_ID=$(aws hoge create | jq -r ".HogeID")
//waituntil --interval 10 --retry 100 --command "aws hoge get --id $HOGE_ID" --ok="SUCCEEDED|READY" --ng="FAILED"
import (
	"github.com/jessevdk/go-flags"
	"github.com/kangaechu/waituntil/pkg/waituntil"
	"os"
)

var opts waituntil.CmdOpts

func main() {

	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}
	waituntil.Run(opts)
}
