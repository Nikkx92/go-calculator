package main

import (
	"bufio"
	"fmt"
	"github.com/Nikkx92/go-calculator/calculator"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите текст:")
	for sc.Scan() {
		txt := sc.Text()
		calculator.Calc(txt)
	}
}
