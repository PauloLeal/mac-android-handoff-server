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

const copyFileScript = `#!/usr/bin/osascript
	on run args
		set the clipboard to POSIX file (first item of args)
end
`

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
		script := fmt.Sprintf("#/bin/sh\necho \"%s\" | /usr/bin/pbcopy", string(data))
		err := utils.RunShellScript(script)
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

		script := copyFileScript
		err = utils.RunOsaScript(script, path)
		if err != nil {
			return err
		}
	}

	return nil
}
