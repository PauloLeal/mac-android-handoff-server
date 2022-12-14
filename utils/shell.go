package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func shouldDeleteFiles() bool {
	return os.Getenv("ANDROID_HANDOFF_DELETE_FILES") == ""
}

func runScript(command string, script string, args ...string) error {
	f, err := ioutil.TempFile("/tmp/", "android-handoff")
	if err != nil {
		return err
	}
	defer f.Close()

	f.Write([]byte(script))
	path, err := filepath.Abs(f.Name())
	if err != nil {
		return err
	}
	defer func() {
		if shouldDeleteFiles() {
			os.Remove(path)
		}
	}()

	newArgs := append([]string{}, path)
	newArgs = append(newArgs, args...)

	cmd := exec.Command(command, newArgs...)
	logrus.Debugln(cmd.String())
	err = cmd.Run()
	if err != nil {
		return err
	}

	c := cmd.ProcessState.ExitCode()
	if c != 0 {
		return errors.New("failed to run command")
	}

	return nil
}

func RunOsaScript(script string, args ...string) error {
	return runScript("/usr/bin/osascript", script, args...)
}

func RunShellScript(script string, args ...string) error {
	return runScript("/bin/sh", script, args...)
}
