package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/tkuchiki/faketime"
)

func main() {
	f := faketime.NewFaketime(2020, time.February, 10, 23, 0, 0, 0, time.UTC)
	defer f.Undo()
	f.Do()

	var stdout, stderr bytes.Buffer
	var cmd *exec.Cmd

	if len(os.Args) == 2 {
		cmd = exec.Command(os.Args[1])
	} else {
		cmd = exec.Command(os.Args[1], os.Args[2:]...)
	}

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		panic("exec.Command(\"" + strings.Join(cmd.Args, "\", \"") + "\").Run() Exception: " + err.Error() + "\n\n" + stderr.String())
	}

	fmt.Println(stdout.String())

}
