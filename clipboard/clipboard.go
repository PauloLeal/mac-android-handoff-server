package clipboard

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/PauloLeal/mac-android-handoff-server/utils"
)

const copyFileScript = `#!/usr/bin/osascript
	on run args
		set the clipboard to POSIX file (first item of args)
end
`

func AddToClipboard(data []byte) error {
	fileType := http.DetectContentType(data)

	if strings.HasPrefix(fileType, "image/") {
		exts, err := mime.ExtensionsByType(fileType)
		if err != nil {
			return err
		}
		f, err := ioutil.TempFile("/tmp/", fmt.Sprintf("android-handoff*%s", exts[0]))
		if err != nil {
			return err
		}
		defer f.Close()
		f.Write(data)

		path, err := filepath.Abs(f.Name())
		if err != nil {
			return err
		}
		defer os.Remove(path)

		script := copyFileScript
		err = utils.RunOsaScript(script, path)
		if err != nil {
			return err
		}
	} else {
		script := fmt.Sprintf("#/bin/sh\necho \"%s\" | /usr/bin/pbcopy", string(data))
		err := utils.RunShellScript(script)
		if err != nil {
			return err
		}
	}

	return nil
}
