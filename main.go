package main

import (
	"fmt"
	"log"
	"time"

	"github.com/edfan0930/dlltool/printer"
	"golang.org/x/sys/windows"
)

func main() {

	//imagePath := "C:\\path\\to\\your\\image.jpg"
	imagePath := ".\\R.jpg"
	err := printer.ShellExecute("print", imagePath, "Microsoft Print to PDF", "", windows.SW_SHOW)
	//	err := printer.ShellExecute("printaaa", imagePath, "", "", 5)
	if err != nil {
		log.Fatalf("Failed to print image: %v", err)
		return
	}

	fmt.Println("over printer")
	t := time.NewTimer(30 * time.Second)
	<-t.C
}
