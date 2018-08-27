package main

import (
	"fmt"
	"regexp"
	"strings"
	"sort"
)

type SubM map[string]string

type M map[string]SubM

func main() {
	println(r("1st test"))
	println(r("1st.test"))
	println(r("this is test # 3"))
	println(r("Last config, last but not least (conf 4)."))
}

func r(s string) string {
	re := regexp.MustCompile(`[^\w\d]+`)
	str := re.ReplaceAllString(s, "-")
	return strings.ToLower(str)
}

func tagsFromHTML(s string) string {
	re := regexp.MustCompile(`(?imsU)<(\w+?).*>`)
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
	re := regexp.MustCompile(`(?imsU)<a\b.*href="(.*)".*>(([^<].*[^<])(<.*>)?)?<\/a`)
	matches := re.FindAllStringSubmatch(s, -1)

	for _, m := range matches {
		fmt.Printf("\n %s 🔵 %s", m[1], strings.Trim(m[3], " "))
	}
}

func match(s string) string {
	re := regexp.MustCompile(`(?i)^[_.]\d+[a-zA-Z]*[_]?$`)
	if re.MatchString(s) {
		return "VALID"
	} else {
		return "INVALID"
	}
}

func addKeyIntoMap(m M, k string) {
    _, imMap := m[k]
    if !imMap {
        m[k] = make(SubM)
    }
}

func addValueIntoMap(m M, k string, v string) {
    m[k][v] = v
}

func printMap(m M) {
    ml := []string{}
    for v := range m {
        ml = append(ml, v)
    }
    sort.Strings(ml)

    for _, k := range ml {
        v := m[k]
        fmt.Printf("%s:", k)
        s := []string{}
        for sk, _ := range v {
            s = append(s, sk)
        }
        sort.Strings(s)
        fmt.Printf("%s\n", strings.Join(s, ","))
    }
}

func tagsAndAttributes(m M, s string) {
    re1 := regexp.MustCompile(`<([a-z0-9]+)\s?[^>]*>`)
    re2 := regexp.MustCompile(`([a-z]+)=["'].*?["']`)

    matches1 := re1.FindAllStringSubmatch(s, -1)
    for _, mv := range matches1 {
        tagName := mv[1]
        tag := mv[0]
        addKeyIntoMap(m, tagName)

        matches2 := re2.FindAllStringSubmatch(tag, -1)

        fmt.Printf("\n ❇️ %v ||| %v \n", tag, matches2)
        for _, mv2 := range matches2 {
            attr := mv2[1]
            addValueIntoMap(m, tagName, attr)
        }
    }
}

func removeWWW() {
  url := "https://www.thethirty.byrdie.com/july-horoscopes-2018"
  regex := regexp.MustCompile(`^(https?://)(www.)(.*)$`)
  result := regex.ReplaceAllString(url, `$1$3`)
  fmt.Printf("> %+v", result)
}
