package vlc

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"

	"MeowTube/logger"

	exec2 "github.com/cli/safeexec"
)

// VideoLAN struct
type VideoLAN struct {
	vlc string
}

// New initialization VideoLAN
func New() *VideoLAN {
	v := "vlc"
	if runtime.GOOS == "darwin" {
		v = "VLC"
	}
	vlc, err := exec2.LookPath(v)
	if err != nil {
		logger.ThrowError(err)
	}
	return &VideoLAN{
		vlc: vlc,
	}
}

// Execute vlc command with args
func (v *VideoLAN) Execute(args ...string) (stdOut string, stdErr string, err error) {
	fmt.Println("Opening VLC...")
	cmd := exec.Command(v.vlc, args...)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		logger.ThrowError(err)
	}
	stdOut, stdErr = stdout.String(), stderr.String()
	return
}

// GetVlc to get value of VideoLAN.vlc
func (v *VideoLAN) GetVlc() string {
	return v.vlc
}
