package clipboard

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PauloLeal/mac-android-handoff-server/utils"
	"github.com/sirupsen/logrus"
)

const copyTextScript = `#!/usr/bin/osascript
on run args
	set the clipboard to (first item of args)
end
`

const copyFileScript = `#!/usr/bin/osascript
	on run args
		set the clipboard to POSIX file (first item of args)
end
`

const pasteFileScript = `#!/usr/bin/osascript
on run args
        tell application "System Events" to write (the clipboard as «class furl») to POSIX file (first item of args)
end
`

// write (the clipboard to POSIX file (first item of args))

var baseTempDir string

func init() {
	d, err := os.MkdirTemp("/tmp/", "mac-android-handoff-")
	if err != nil {
		log.Panic("Unable to create temp dir")
	}
	baseTempDir = d
	cleanup()
}

func cleanup() {
	go func() {

	}()
}

func AddToClipboard(data []byte, name string) error {
	fileType := http.DetectContentType(data)

	logrus.Debugln(fileType)
	if strings.HasPrefix(fileType, "text/plain") {
		err := utils.RunOsaScript(copyTextScript, string(data))
		if err != nil {
			return err
		}
	} else {
		if len(name) == 0 {
			return errors.New("name required")
		}

		f, err := os.Create(fmt.Sprintf("%s/%s", baseTempDir, name))
		if err != nil {
			return err
		}
		defer f.Close()
		f.Write(data)

		path, err := filepath.Abs(f.Name())
		if err != nil {
			return err
		}

		err = utils.RunOsaScript(copyFileScript, path)
		if err != nil {
			return err
		}
	}

	return nil
}

func ReadFromClipboard() ([]byte, error) {
	f, _ := os.CreateTemp(baseTempDir, "pasted-")
	logrus.Debugln(f.Name())
	utils.RunOsaScript(pasteFileScript, f.Name())
	return nil, nil
}
