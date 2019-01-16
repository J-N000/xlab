package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

var version string
var name string
var terminal string

func init() {
	flag.StringVar(&version, "v", "latest", "specify the target image version")
	flag.StringVar(&name, "n", "xlab-container", "specify a name for the container")
	terminal = "urxvt"
}

func main() {
	flag.Parse()
	pwd, _ := exec.Command("pwd").Output()
	mTarget := "/root/mount"
	mStatement := "src=" + string(pwd[:]) + ",dst=" + mTarget
	image := fmt.Sprintf("j-n000/xlab:%s", version)
	cmd := terminal
	cmdArgs := []string{
		"-e",
		"docker",
		"run",
		"-it",
		"--rm",
		"--mount",
		mStatement,
		"-w",
		mTarget,
		"--name",
		name,
		image}

	fmt.Println(mStatement)

	var out bytes.Buffer
	var stdErr bytes.Buffer
	xlab := exec.Command(cmd, cmdArgs...)
	xlab.Stdout = &out
	xlab.Stderr = &stdErr
	err := xlab.Run()

	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stdErr.String())
		log.Fatal(err)
	}
}

// example command
// urxvt -e docker run -it --rm -v /home/user/fs:/root/mount -w /root/mount --name xlab-container j-n000/xlab:latest
