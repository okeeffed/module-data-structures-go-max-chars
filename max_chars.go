package maxchars

import (
	"regexp"
)

func maxChars(str string) string {
	cs := make(map[string]int)

	for _, c := range str {
		key := string(c)

		re := regexp.MustCompile(`[a-zA-Z]`)
		valid := re.MatchString(key)

		if val, ok := cs[key]; ok && valid {
			cs[key] = val + 1
		} else if valid {
			cs[key] = 1
		}
	}

	mc := ""
	max := 0

	for key, value := range cs {
		if cs[key] > max {
			mc = key
			max = value
		}
	}

	return mc
}
