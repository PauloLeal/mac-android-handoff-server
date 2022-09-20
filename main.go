package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/PauloLeal/mac-android-handoff-server/clipboard"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	targetFile = kingpin.Flag("file", "File to copy").Short('f').String()
	text       = kingpin.Flag("text", "Text to copy").String()
)

func main() {
	kingpin.Parse()

	if *targetFile != "" {
		fmt.Println(*targetFile)
		f, err := os.Open(*targetFile)
		defer f.Close()
		if err != nil {
			log.Fatal(err)
		}

		b, err := io.ReadAll(bufio.NewReader(f))
		if err != nil {
			log.Fatal(err)
		}

		err = clipboard.AddToClipboard(b)
		if err != nil {
			log.Fatal(err)
		}
	}

	if *text != "" {
		err := clipboard.AddToClipboard([]byte(*text))
		if err != nil {
			log.Fatal(err)
		}
	}

	// AddToClipboard(io.ReadAll(os.Ope))
}
