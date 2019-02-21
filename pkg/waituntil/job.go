package waituntil

import (
	"github.com/mattn/go-shellwords"
	"log"
	"os/exec"
)

type Job struct {
	Command    []string // ジョブ実行に使用するコマンド
	Interval   uint     // 実行間隔
	MaxRetries uint     // リトライの回数
	OkText     string   // OKとなるメッセージ
	NgText     string   // NGとなるメッセージ
}

func BuildJob(opts CmdOpts) (job Job, err error) {
	job = Job{}

	// コマンドをParse (クオート文字を考慮)
	commandList, err := shellwords.Parse(opts.Command)
	if err != nil {
		return job, err
	}

	job.Command = commandList
	job.Interval = opts.Interval
	job.MaxRetries = opts.MaxRetries
	job.OkText = opts.OkText
	job.NgText = opts.NgText
	return job, nil
}

func (j Job) Run() (err error) {
	command := j.Command[0]
	args := j.Command[1:]
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		return err
	}

	result := string(out)
	log.Println("output: ", result)
	return nil
}
