package main

import (
	"os"
	"runtime"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
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

	// if *paste {
	// 	clipboard.ReadFromClipboard()
	// }

	// if *targetFile != "" {
	// 	logrus.Debugln(*targetFile)
	// 	f, err := os.Open(*targetFile)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	defer f.Close()

	// 	b, err := io.ReadAll(bufio.NewReader(f))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fileName := filepath.Base(f.Name())
	// 	err = clipboard.AddToClipboard(b, fileName)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// if *text != "" {
	// 	err := clipboard.AddToClipboard([]byte(*text), "")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// AddToClipboard(io.ReadAll(os.Ope))

	runtime.LockOSThread()

	cocoa.TerminateAfterWindowsClose = false
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		url := "https://icons-for-free.com/iconfiles/png/24/clock+time+timer+watch+icon-1320183419333964306.png"
		imageUrlRef := core.NSURL_URLWithString_(core.NSString_FromString(url))

		icon := cocoa.NSImage_InitWithURL(imageUrlRef)

		obj := cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength)
		obj.Retain()
		obj.Button().SetImage(icon)

		// obj.SetAction(objc.Sel("mainButtonClicked:"))
		// cocoa.DefaultDelegateClass.AddMethod("mainButtonClicked:", func(_ objc.Object) {
		// 	logrus.Infoln("QEQWEQWEW")
		// })

		itemQuit := cocoa.NSMenuItem_New()
		itemQuit.SetTitle("Quit")
		itemQuit.SetAction(objc.Sel("terminate:"))

		menu := cocoa.NSMenu_New()
		menu.AddItem(itemQuit)
		obj.SetMenu(menu)
	})

	// pasteboard := cocoa.NSPasteboard_GeneralPasteboard()
	// pasteboard.Retain()

	// server start
	app.Run() // blocks
	// server stop
}
