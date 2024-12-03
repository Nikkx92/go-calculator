package calculator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Calc(str string) {

	text, _ := regexp.MatchString(`^("(.{1,12})")\s(\+|-)\s("(.{1,12})")$|^("(.{1,12})")\s(\*|\/)\s([1-9]|10)$`, str)
	if text != true {
		panic("Недопустимый ввод")
	}
	var result string
	withoutQuotes := strings.ReplaceAll(str, `"`, "")
	findZnak := regexp.MustCompile(` \/ | \* | \+ | - `).FindAllString(str, -1)
	znak := strings.Join(findZnak, "")
	part := strings.Split(withoutQuotes, znak)
	if znak == " + " {
		result = part[0] + part[1]
	} else if znak == " - " {
		result = regexp.MustCompile(part[1]).ReplaceAllString(part[0], "")
	} else if znak == " / " {
		nd := strings.Split(part[0], "")
		d, _ := strconv.Atoi(part[1])
		l := len(nd) / d
		for i, v := range nd {
			result += v
			if i == l-1 {
				break
			}
		}
	} else if znak == " * " {
		n, _ := strconv.Atoi(part[1])
		result = strings.Repeat(part[0], n)
	}
	if len(result) > 40 {
		fmt.Printf("%.40s ...", fmt.Sprintf("%#v", result))
	} else {
		fmt.Printf("%#v", result)
	}
}
