package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var size = flag.String("size", "", "width,height in pixels (e.g. 1024px,768px or 3,3)")

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing paths of images to cat")
		os.Exit(2)
	}
	for _, path := range os.Args[1:] {
		if err := cat(path); err != nil {
			fmt.Fprintf(os.Stderr, "could not cat %s: %v\n", path, err)
		}
	}
}

func widthAndHeight() (w, h string) {
	if *size != "" {
		sp := strings.SplitN(*size, ",", -1)
		if len(sp) == 2 {
			w = sp[0]
			h = sp[1]
		}
	}
	return
}

func cat(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	width, height := widthAndHeight()

	fmt.Print("\033]1337;")
	fmt.Printf("File=inline=1:")
	if width != "" || height != "" {
		if width != "" {
			fmt.Printf(";width=%s", width)
		}
		if height != "" {
			fmt.Printf(";height=%s", height)
		}
	}

	wc := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	defer wc.Close()
	_, err = io.Copy(wc, f)
	if err != nil {
		return err
	}
	fmt.Printf("\a\n")

	return err
}
