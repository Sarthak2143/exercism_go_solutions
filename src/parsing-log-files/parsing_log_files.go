package parsinglogfiles

import (
    "regexp"
    "fmt"
)

func IsValidLine(text string) bool {
    re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
    re, _ := regexp.Compile(`<[~*=-]*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    count := 0
    re := regexp.MustCompile(`(?i)".*password.*"`)
    for _, i := range lines {
        if re.MatchString(i) {
            count++
        } 
    }
    return count
}

func RemoveEndOfLineText(text string) string {
    re, _ := regexp.Compile(`end-of-line\d*`)
    return re.ReplaceAllString(text, "")
}
func TagWithUserName(lines []string) []string {
    re := regexp.MustCompile(`User\s+(\w+)`)
    ret := []string{}
    for _, i := range lines {
        founds := re.FindStringSubmatch(i)
        if founds != nil {
            i = fmt.Sprintf("[USR] %s %s", founds[1], i)
        }
        ret = append(ret, i)
    }
    return ret
}

