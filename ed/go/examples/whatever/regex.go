package main

import (
	"regexp"
	"strings"
)

func main() {
	println(r("1st test"))
	println(r("1st.test"))
	println(r("this is test # 3"))
	println(r("Last config, last but not least (conf 4)."))
}

func r(s string) string {
	var re = regexp.MustCompile(`[^\w\d]+`)
	str := re.ReplaceAllString(s, "-")
	return strings.ToLower(str)
}
