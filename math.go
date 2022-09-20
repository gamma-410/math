package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
  "strings"
)

func main() {
	flag.Parse()
	name := flag.Arg(0)
	f, _ := os.Open(name)
	scanner := bufio.NewScanner(f)

	lenNum := 0
	capNum := 0
	cutText := make([]string, lenNum, capNum) // 1
  splitLen := make([]int, lenNum, capNum) // 2

  // 1
	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		cutText = append(cutText, text)
		i++
	}
  
  // 2
  l := 0
  for i > l {
    cut := cutText[l]
    cutCut := strings.Split(cut, " ")
    splitLen = append(splitLen, len(cutCut))
    l++
  }
  fmt.Println(cutText)
  fmt.Println(splitLen)
}
