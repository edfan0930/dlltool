package printer

import (
	"log"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	shell32           = windows.NewLazySystemDLL("shell32.dll")
	procShellExecuteW = shell32.NewProc("ShellExecuteW")
)

// ShellExecute
func ShellExecute(verb, file, args, dir string, showCmd int) error {
	verbPtr := windows.StringToUTF16Ptr(verb)
	filePtr := windows.StringToUTF16Ptr(file)
	argsPtr := windows.StringToUTF16Ptr(args)
	dirPtr := windows.StringToUTF16Ptr(dir)

	_, _, err := procShellExecuteW.Call(
		0,
		uintptr(unsafe.Pointer(verbPtr)),
		uintptr(unsafe.Pointer(filePtr)),
		uintptr(unsafe.Pointer(argsPtr)),
		uintptr(unsafe.Pointer(dirPtr)),
		uintptr(showCmd),
	)

	if err != windows.ERROR_SUCCESS {
		return err
	}

	return nil
}

func main() {
	// Print an image
	//imagePath := "C:\\path\\to\\your\\image.jpg"
	imagePath := ".\\R.jpg"
	err := ShellExecute("print", imagePath, "", "", 1)
	if err != nil {
		log.Fatalf("Failed to print image: %v", err)
	}
}
