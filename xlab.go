package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var (
	com      string
	version  string
	name     string
	terminal string
	image    string
	stdOut   bytes.Buffer
	stdErr   bytes.Buffer
)

const (
	hashLen     int    = 12
	imageName   string = "n0ja/xlab"
	mountTarget string = "/root/mount"
)

func commit() error {
	contID := docPsGrep()
	commitArgs := []string{
		"commit",
		contID,
		image}
	fmt.Println(fmt.Sprintf("Committing %s @ %s to %s", com, contID, image))
	return exCmd("docker", commitArgs)
}

func run() error {
	pwd, _ := exec.Command("pwd").Output()
	fmtDir := string(pwd[:])
	mStatement := fmt.Sprintf("%s:%s", fmtDir[:len(fmtDir)-1], mountTarget)
	runArgs := []string{
		"-e",
		"docker",
		"run",
		"-it",
		"--rm",
		"-v",
		mStatement,
		"-w",
		mountTarget,
		"--name",
		name,
		image}
	return exCmd(terminal, runArgs)
}

func docPsGrep() string {
	psArgs := []string{
		"ps",
		"-f",
		"name=" + com}
	psOut, _ := exec.Command("docker", psArgs...).Output()
	psString := string(psOut[:])
	psArr := strings.Split(psString, "\n")
	return psArr[1][:hashLen]
}

func exCmd(name string, cmdArgs []string) error {
	xlab := exec.Command(name, cmdArgs...)
	xlab.Stdout = &stdOut
	xlab.Stderr = &stdErr
	return xlab.Run()
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(fmt.Sprintf("%v: %s", err, stdErr.String()))
		log.Fatal(err)
	}
}

func init() {
	flag.StringVar(&com, "c", "", "specify by NAME, a container to commit to latest. The default container is named xlab.")
	flag.StringVar(&version, "v", "latest", "specify the target image version")
	flag.StringVar(&name, "n", "xlab", "specify a name for the container")
	flag.StringVar(&terminal, "t", "urxvt", "terminal in which to execute container initialization")
}

func main() {
	flag.Parse()
	image = fmt.Sprintf("%s:%s", imageName, version)
	if com != "" {
		handleErr(commit())
	} else {
		handleErr(run())
	}
}
