package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	name := flag.Arg(0)
	f, _ := os.Open(name)
	scanner := bufio.NewScanner(f)

	lenNum := 0
	capNum := 0
	cutText := make([]string, lenNum, capNum)

	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		cutText = append(cutText, text)
		i++
	}
	fmt.Println(cutText)
}
