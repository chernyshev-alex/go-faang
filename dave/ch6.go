package dave

import (
	"bufio"
	"fmt"
	"os"
)

func do_ch6() {
	b, err := os.ReadFile("ch1.go")
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(b))
}

func readLine() {
	f, err := os.OpenFile("ch1.go", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		fmt.Printf("%s\n", line)
	}
}
