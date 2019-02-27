package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
)

var (
	comName  string
	version  string
	name     string
	terminal string
	out      bytes.Buffer
	stdErr   bytes.Buffer
)

func commit() error {
	ps := fmt.Sprintf("docker ps noja/xlab:%s", comName)
	contName, _ := exec.Command(ps).Output()
	hashName := string(contName[:3])
	commitArgs := []string{
		"-e",
		"docker",
		"commit",
		hashName,
		"noja/xlab:latest"}
	err := exCmd(commitArgs)
	return err
}

func run() error {
	pwd, _ := exec.Command("pwd").Output()
	fmtDir := string(pwd[:])
	mTarget := "/root/mount"
	mStatement := fmtDir[:len(fmtDir)-1] + ":" + mTarget
	image := fmt.Sprintf("n0ja/xlab:%s", version)
	runArgs := []string{
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
	err := exCmd(runArgs)
	return err
}

func exCmd(cmdArgs []string) error {
	xlab := exec.Command(terminal, cmdArgs...)
	xlab.Stdout = &out
	xlab.Stderr = &stdErr
	err := xlab.Run()
	return err
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stdErr.String())
		log.Fatal(err)
	}
}

func init() {
	flag.StringVar(&comName, "c", "", "specify by NAME, a container to commit to latest")
	flag.StringVar(&version, "v", "latest", "specify the target image version")
	flag.StringVar(&name, "n", "xlab-container", "specify a name for the container")
	flag.StringVar(&terminal, "t", "urxvt", "terminal in which to execute container initialization")
}

func main() {
	flag.Parse()
	if len(comName) > 0 {
		handleErr(commit())
	} else {
		handleErr(run())
	}
}
