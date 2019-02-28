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
	psName := fmt.Sprintf("name=%s", com)
	contID := docPsFilter(psName)
	commitArgs := []string{
		"commit",
		contID,
		image}
	fmt.Println(fmt.Sprintf("Committing %s @ %s to %s...", com, contID, image))
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

func docPsFilter(s string) string {
	psArgs := []string{
		"ps",
		"-f",
		s}
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
	flag.StringVar(&com, "c", "", "specify NAME of a container to commit to a version (-v)")
	flag.StringVar(&version, "v", "latest", "specify the target image VERSION")
	flag.StringVar(&name, "n", "xlab", "specify a NAME for the container")
	flag.StringVar(&terminal, "t", "urxvt", "TERMINAL in which to execute container initialization")
}

func main() {
	flag.Parse()
	image = fmt.Sprintf("%s:%s", imageName, version)
	if com != "" {
		err := commit()
		switch err {
		case nil:
			fmt.Println("Success!")
			break
		default:
			handleErr(err)
		}
	} else {
		handleErr(run())
	}
}
