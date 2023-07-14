package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type SubM map[string]string

type M map[string]SubM

func main() {
	// println(r("1st test"))
	// println(r("1st.test"))
	// println(r("this is test # 3"))
	// println(r("Last config, last but not least (conf 4)."))
	// noCapture()
	// checkIsValidGeneralID()
	// r2()
	allFindFuncs()
}

func checkIsValidGeneralID() {
	fmt.Printf("[GeneralID] 1=%v \n", isValidGeneralID("-1"))
	fmt.Printf("[GeneralID] 2=%v \n", isValidGeneralID(""))
	fmt.Printf("[GeneralID] 2=%v \n", isValidGeneralID("_"))
	fmt.Printf("[GeneralID] 2=%v \n", isValidGeneralID("_x"))
	fmt.Printf("[GeneralID] 2=%v \n", isValidGeneralID("x37dc"))
}

func isValidGeneralID(s string) bool {
	re := regexp.MustCompile(`(?i)^[\w]+[\w\d-]+$`)
	return re.MatchString(s)
}

func noCapture() {
	s := "http-host-url"
	re := regexp.MustCompile(`(?:http-)(.*)`)
	fmt.Printf("%+v \n", re.FindAllStringSubmatch(s, -1))
}

func r(s string) string {
	re := regexp.MustCompile(`[^\w\d]+`)
	str := re.ReplaceAllString(s, "-")
	return strings.ToLower(str)
}

func r2() {
	re := regexp.MustCompile(`(\*pkg\.|Type|_)`)
	fmt.Printf("%+v \n", re.ReplaceAllString("*pkg.Type_Name", "")) // Name
}

func instagram(ss string) string {
	s := "<blockquote class=\"instagram-media\" data-instgrm-permalink=\"https://www.instagram.com/p/BmTmwgXlJsh/?utm_source=ig_embed\" data-instgrm-version=\"9\" style=\" background:#FFF; border:0; border-radius:3px; box-shadow:0 0 1px 0 rgba(0,0,0,0.5),0 1px 10px 0 rgba(0,0,0,0.15); margin: 1px; max-width:540px; min-width:326px; padding:0; width:99.375%; width:-webkit-calc(100% - 2px); width:calc(100% - 2px);\"><div style=\"padding:8px;\"> <div style=\" background:#F8F8F8; line-height:0; margin-top:40px; padding:62.4537037037037% 0; text-align:center; width:100%;\"> <div style=\" background:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACwAAAAsCAMAAAApWqozAAAABGdBTUEAALGPC/xhBQAAAAFzUkdCAK7OHOkAAAAMUExURczMzPf399fX1+bm5mzY9AMAAADiSURBVDjLvZXbEsMgCES5/P8/t9FuRVCRmU73JWlzosgSIIZURCjo/ad+EQJJB4Hv8BFt+IDpQoCx1wjOSBFhh2XssxEIYn3ulI/6MNReE07UIWJEv8UEOWDS88LY97kqyTliJKKtuYBbruAyVh5wOHiXmpi5we58Ek028czwyuQdLKPG1Bkb4NnM+VeAnfHqn1k4+GPT6uGQcvu2h2OVuIf/gWUFyy8OWEpdyZSa3aVCqpVoVvzZZ2VTnn2wU8qzVjDDetO90GSy9mVLqtgYSy231MxrY6I2gGqjrTY0L8fxCxfCBbhWrsYYAAAAAElFTkSuQmCC); display:block; height:44px; margin:0 auto -44px; position:relative; top:-22px; width:44px;\"></div></div><p style=\" color:#c9c8cd; font-family:Arial,sans-serif; font-size:14px; line-height:17px; margin-bottom:0; margin-top:8px; overflow:hidden; padding:8px 0 7px; text-align:center; text-overflow:ellipsis; white-space:nowrap;\"><a href=\"https://www.instagram.com/p/BmTmwgXlJsh/?utm_source=ig_embed\" style=\" color:#c9c8cd; font-family:Arial,sans-serif; font-size:14px; font-style:normal; font-weight:normal; line-height:17px; text-decoration:none;\" target=\"_blank\">A post shared by Jennifer Lopez (@jlo)</a> on <time style=\" font-family:Arial,sans-serif; font-size:14px; line-height:17px;\" datetime=\"2018-08-10T17:13:28+00:00\">Aug 10, 2018 at 10:13am PDT</time></p></div></blockquote> <script async defer src=\"//www.instagram.com/embed.js\"></script>"
	re := regexp.MustCompile(`(?imsU)data-instgrm-permalink="(.*)"`)
	r := re.FindStringSubmatch(s)
	fmt.Printf("%+v \n", r)
	return ""
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

// DeleteAttrsFromTag deletes html tag's attrs.
// @see: https://regex101.com/r/cjBrK8/1
func DeleteAttrsFromTag(s string) (res string) {
	res = regexp.MustCompile(`<(\w+)(\s*[\w-]+=["][^"]*["])*(\s*[/])?>`).ReplaceAllString(s, "<$1$3>")
	return res
}

func allFindFuncs() {
	pattern := `-([a-z0-9]+)-`
	s := `style1-season2-color3-test3-test4`
	b := []byte(s)

	re, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}

	m01 := re.FindAllStringSubmatch(s, -1)
	m02 := re.Find(b)
	m03 := re.FindAll(b, -1)
	m04 := re.FindAllIndex(b, -1)
	m05 := re.FindIndex(b)
	m06 := re.FindAllString(s, -1)
	m07 := re.FindAllStringIndex(s, -1)
	m08 := re.FindAllStringSubmatch(s, -1)
	m09 := re.FindAllStringSubmatchIndex(s, -1)
	m10 := re.FindAllSubmatch(b, -1)
	m11 := re.FindAllSubmatchIndex(b, -1)
	m12 := re.FindIndex(b)
	m13 := re.FindString(s)
	m14 := re.FindStringIndex(s)
	m15 := re.FindStringSubmatch(s)
	m16 := re.FindStringSubmatchIndex(s)
	m17 := re.FindSubmatch(b)
	m18 := re.FindSubmatchIndex(b)

	fmt.Printf("Matches01: %v\n", m01) // Matches01: [[-season2- season2] [-test3- test3]]
	fmt.Printf("Matches02: %s\n", m02) // Matches02: -season2-
	fmt.Printf("Matches03: %s\n", m03) // Matches03: [-season2- -test3-]
	fmt.Printf("Matches04: %v\n", m04) // Matches04: [[6 15] [21 28]]
	fmt.Printf("Matches05: %v\n", m05) // Matches05: [6 15]
	fmt.Printf("Matches06: %v\n", m06) // Matches06: [-season2- -test3-]
	fmt.Printf("Matches07: %v\n", m07) // Matches07: [[6 15] [21 28]]
	fmt.Printf("Matches08: %v\n", m08) // Matches08: [[-season2- season2] [-test3- test3]]
	fmt.Printf("Matches09: %v\n", m09) // Matches09: [[6 15 7 14] [21 28 22 27]]
	fmt.Printf("Matches10: %s\n", m10) // Matches10: [[-season2- season2] [-test3- test3]]
	fmt.Printf("Matches11: %v\n", m11) // Matches11: [[6 15 7 14] [21 28 22 27]]
	fmt.Printf("Matches12: %v\n", m12) // Matches12: [6 15]
	fmt.Printf("Matches13: %v\n", m13) // Matches13: -season2-
	fmt.Printf("Matches14: %v\n", m14) // Matches14: [6 15]
	fmt.Printf("Matches15: %v\n", m15) // Matches15: [-season2- season2]
	fmt.Printf("Matches16: %v\n", m16) // Matches16: [6 15 7 14]
	fmt.Printf("Matches17: %s\n", m17) // Matches17: [-season2- season2]
	fmt.Printf("Matches18: %v\n", m18) // Matches18: [6 15 7 14]

}
