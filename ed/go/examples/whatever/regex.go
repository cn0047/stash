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

func tagsFromHTML(s string) string {
	var re = regexp.MustCompile(`(?imsU)<(\w+?).*>`)
	matches := re.FindAllStringSubmatch(s, -1)

	tags := make(map[string]uint)
	for _, m := range matches {
		tags[m[1]] = 1
	}

	// sort
	a := make([]string, 0)
	for t := range tags {
		a = append(a, t)
	}
	sort.Strings(a)

	return strings.Join(a, ";")
}

func hrefToLink(s string) {
	var re = regexp.MustCompile(`(?imsU)<a\b.*href="(.*)".*>(([^<].*[^<])(<.*>)?)?<\/a`)
	matches := re.FindAllStringSubmatch(s, -1)

	for _, m := range matches {
		fmt.Printf("\n %s 🔵 %s", m[1], strings.Trim(m[3], " "))
	}
}

func match(s string) string {
	var re = regexp.MustCompile(`(?i)^[_.]\d+[a-zA-Z]*[_]?$`)
	if re.MatchString(s) {
		return "VALID"
	} else {
		return "INVALID"
	}
}
