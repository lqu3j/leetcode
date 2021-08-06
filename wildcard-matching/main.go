package main

import "log"

func main() {
	log.Println(isMatch("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))
}

func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)
	sIndex := 0
	pIndex := 0

	for pIndex < pLen {
		if p[pIndex] == '*' {
			for pIndex < pLen && p[pIndex] == '*' {
				pIndex++
			}
			for sIndex < sLen {
				if pIndex == pLen && sIndex == sLen {
					return true
				}
				if sIndex < sLen && pIndex < pLen && isMatch(s[sIndex:], p[pIndex:]) {
					return true
				}
				sIndex++
			}
		} else {
			if (sIndex < sLen && s[sIndex] == p[pIndex]) || p[pIndex] == '?' {
				sIndex++
				pIndex++
			} else {
				return false
			}
		}
	}
	return pIndex == pLen && sIndex == sLen
}
