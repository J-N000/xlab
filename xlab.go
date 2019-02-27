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
	com      bool
	version  string
	name     string
	terminal string
	stdOut   bytes.Buffer
	stdErr   bytes.Buffer
)

const (
	hashLen   int    = 12
	imageName string = "n0ja/xlab"
)

func commit() error {
	contID := docPsGrep()
	image := fmt.Sprintf("%s:%s", imageName, version)
	commitArgs := []string{
		"commit",
		contID,
		image}
	err := exCmd("docker", commitArgs)
	return err
}

func run() error {
	pwd, _ := exec.Command("pwd").Output()
	fmtDir := string(pwd[:])
	mTarget := "/root/mount"
	mStatement := fmtDir[:len(fmtDir)-1] + ":" + mTarget
	image := fmt.Sprintf("%s:%s", imageName, version)
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
		if match, _ := regexp.MatchString(version, line); match {
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
	flag.BoolVar(&com, "c", false, "specify by version, a container to commit. Default is latest.")
	flag.StringVar(&version, "v", "latest", "specify the target image version")
	flag.StringVar(&name, "n", "xlab", "specify a name for the container")
	flag.StringVar(&terminal, "t", "urxvt", "terminal in which to execute container initialization")
}

func main() {
	flag.Parse()
	switch com {
	case true:
		handleErr(commit())
		break
	default:
		handleErr(run())
	}
}
