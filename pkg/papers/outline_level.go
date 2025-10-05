package papers

import (
	"fmt"
	"regexp"
)

//	If I ever need to actually convert roman numerals to ints: "github.com/brandenc40/romannumeral"

const unitIndent = ` {4}`
const arabic = `[0-9]+`
const latinCap = `[A-Z]+`
const latinLower = `[a-z]+`
const romanCap = `[IVXLCDM]+`
const romanLower = `[ivxlcdm]+`

func outlinePointRe() *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(
		`^\s*(?:%s\.|%s\.|%s\.|%s\)|\(%s\)|\(%s\)|%s\))\s+`,
		romanCap,
		latinCap,
		arabic,
		latinLower,
		arabic,
		latinLower,
		romanLower,
	))
}

func indent(lvl int) string {
	return fmt.Sprintf("(?:%s){%d}", unitIndent, lvl)
}

func l0re() *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(`^(%s)\.\s+`, romanCap))
}

func l1re() *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(`^%s(%s)\.\s+`, indent(1), latinCap))
}

func l2re(short bool) *regexp.Regexp {
	if short {
		return regexp.MustCompile(fmt.Sprintf(`^(%s)\.\s+`, arabic))
	}
	return regexp.MustCompile(fmt.Sprintf(`^%s(%s)\.\s+`, indent(2), arabic))
}

func l3re(short bool) *regexp.Regexp {
	if short {
		return regexp.MustCompile(fmt.Sprintf(`^%s(%s)(?:\.|\))\s+`, indent(1), latinLower))
	}
	return regexp.MustCompile(fmt.Sprintf(`^%s(%s)(?:\.|\))\s+`, indent(3), latinLower))
}

func l4re(short bool) *regexp.Regexp {
	if short {
		return regexp.MustCompile(fmt.Sprintf(`^%s\((%s)\)\s+`, indent(2), arabic))
	}
	return regexp.MustCompile(fmt.Sprintf(`^%s\((%s)\)\s+`, indent(4), arabic))
}

func l5re(short bool) *regexp.Regexp {
	if short {
		return regexp.MustCompile(fmt.Sprintf(`^%s\((%s)\)\s+`, indent(3), latinLower))
	}
	return regexp.MustCompile(fmt.Sprintf(`^%s\((%s)\)\s+`, indent(5), latinLower))
}

func l6re(short bool) *regexp.Regexp {
	if short {
		return regexp.MustCompile(fmt.Sprintf(`^%s(%s)\)\s+`, indent(4), romanLower))
	}
	return regexp.MustCompile(fmt.Sprintf(`^%s(%s)\)\s+`, indent(6), romanLower))
}

func lvlRe(level int, short bool) *regexp.Regexp {
	switch level {
	case 0:
		if short {
			return l2re(true)
		}
		return l0re()
	case 1:
		if short {
			return l3re(true)
		}
		return l1re()
	case 2:
		if short {
			return l4re(true)
		}
		return l2re(false)
	case 3:
		if short {
			return l5re(true)
		}
		return l3re(false)
	case 4:
		if short {
			return l6re(true)
		}
		return l4re(false)
	case 5:
		if short {
			return nil
		}
		return l5re(false)
	case 6:
		if short {
			return nil
		}
		return l6re(false)
	default:
		return nil
	}
}

func determineLevel(line string, short bool) int {
	if short {
		if l2re(true).MatchString(line) {
			return 0
		}
		if l3re(true).MatchString(line) {
			return 1
		}
		if l4re(true).MatchString(line) {
			return 2
		}
		if l5re(true).MatchString(line) {
			return 3
		}
		if l6re(true).MatchString(line) {
			return 4
		}
	} else {
		if l0re().MatchString(line) {
			return 0
		}
		if l1re().MatchString(line) {
			return 1
		}
		if l2re(false).MatchString(line) {
			return 2
		}
		if l3re(false).MatchString(line) {
			return 3
		}
		if l4re(false).MatchString(line) {
			return 4
		}
		if l5re(false).MatchString(line) {
			return 5
		}
		if l6re(false).MatchString(line) {
			return 6
		}
	}

	return -1
}
