package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

var (
	version string
	name string
	terminal string
	out bytes.Buffer
	stdErr bytes.Buffer
)

func init() {
	flag.StringVar(&version, "v", "latest", "specify the target image version")
	flag.StringVar(&name, "n", "xlab-container", "specify a name for the container")
	terminal = "urxvt"
}

func main() {
	flag.Parse()
	pwd, _ := exec.Command("pwd").Output()
	fmtDir := string(pwd[:])
	mTarget := "/root/mount"
	mStatement := fmtDir[:len(fmtDir) - 1] + ":" + mTarget
	image := fmt.Sprintf("j-n000/xlab:%s", version)
	cmdArgs := []string{
		"-e",
		"docker",
		"run",
		"-it",
		"--rm",
		"-v",
		mStatement,
		"-w",
		mTarget,
		"--name",
		name,
		image}

	xlab := exec.Command(terminal, cmdArgs...)
	xlab.Stdout = &out
	xlab.Stderr = &stdErr
	err := xlab.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stdErr.String())
		log.Fatal(err)
	}
}
