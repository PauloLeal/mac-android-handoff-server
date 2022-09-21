package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/PauloLeal/mac-android-handoff-server/clipboard"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	targetFile = kingpin.Flag("file", "File to copy").Short('f').String()
	text       = kingpin.Flag("text", "Text to copy").String()
	paste      = kingpin.Flag("paste", "paste from clipboard").Short('p').Bool()
)

func initializeLogs() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	kingpin.Parse()

	initializeLogs()

	if *paste {
		clipboard.ReadFromClipboard()
	}

	if *targetFile != "" {
		logrus.Debugln(*targetFile)
		f, err := os.Open(*targetFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		b, err := io.ReadAll(bufio.NewReader(f))
		if err != nil {
			log.Fatal(err)
		}

		fileName := filepath.Base(f.Name())
		err = clipboard.AddToClipboard(b, fileName)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *text != "" {
		err := clipboard.AddToClipboard([]byte(*text), "")
		if err != nil {
			log.Fatal(err)
		}
	}

	// AddToClipboard(io.ReadAll(os.Ope))
}
