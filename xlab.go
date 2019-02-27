package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

var (
	comName  string
	version  string
	name     string
	terminal string
	stdOut   bytes.Buffer
	stdErr   bytes.Buffer
)

const hashLen int = 12

func commit() error {
	contID := docPsGrep()
	commitArgs := []string{
		"commit",
		contID,
		"noja/xlab:latest"}
	err := exCmd("docker", commitArgs)
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
	err := exCmd(terminal, runArgs)
	return err
}

func docPsGrep() string {
	psOut, _ := exec.Command("docker", "ps").Output()
	psString := string(psOut[:])
	psArr := strings.Split(psString, "\n")
	var contInfo string
	for _, line := range psArr {
		if match, _ := regexp.MatchString(comName, line); match {
			contInfo = line
		}
	}
	return contInfo[:hashLen]
}

func exCmd(name string, cmdArgs []string) error {
	xlab := exec.Command(name, cmdArgs...)
	xlab.Stdout = &stdOut
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
