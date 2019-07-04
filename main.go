package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/galijot/cloner"
)

func main() {
	src, dst := getPaths()
	cloner.Clone(src, dst)
}

// Reads the paths FROM and TO which a users wants to clone.
func getPaths() (string, string) {

	reader := bufio.NewScanner(os.Stdin)
	space := " "

	fmt.Println("Enter the path FROM which you want to clone:")
	reader.Scan()
	fromPath := strings.Trim(reader.Text(), space)

	fmt.Println("Enter the path TO which you wan to clone:")
	reader.Scan()
	toPath := strings.Trim(reader.Text(), space)

	// return fromPath, toPath
	fmt.Println(fromPath, toPath)

	return "/Users/mariogalijot/Desktop/fromPath", "/Users/mariogalijot/Desktop/toPath"
}
