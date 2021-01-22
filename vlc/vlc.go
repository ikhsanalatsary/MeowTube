package vlc

import (
	"bytes"
	"log"
	"os/exec"

	exec2 "github.com/cli/safeexec"
)

// VideoLAN struct
type VideoLAN struct {
	vlc string
}

var executableFile string

// New initialization VideoLAN
func New() *VideoLAN {
	vlc, err := exec2.LookPath("vlc")
	if err != nil {
		log.Fatal(err)
	}
	executableFile = vlc
	return &VideoLAN{
		vlc: vlc,
	}
}

//  Execute asd
func (v *VideoLAN) Execute(args ...string) (stdOut string, stdErr string, err error) {
	cmd := exec.Command(v.vlc, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Start()
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}
