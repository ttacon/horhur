package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/playwright-community/playwright-go"
	"github.com/ttacon/horhur"
)

var (
	filename = flag.String("f", "", "name of file to open as browser window")
)

func main() {
	flag.Parse()

	err := playwright.Install()
	if err != nil {
		fmt.Println("failed to install dependencies, err: ", err)
		os.Exit(1)
	}

	info, isRedirected, err := horhur.DoesItRedirect(*filename)
	if err != nil {
		fmt.Println("failed to check for redirect, err: ", err)
		os.Exit(1)
	} else if !isRedirected {
		fmt.Println("no redirect determined")
		return
	}

	fmt.Println("redirect detected")
	fmt.Println("new url detected: ", info.NewURL)
	fmt.Println("same domain: ", info.IsSameDomain)
}
