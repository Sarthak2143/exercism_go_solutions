package raindrops

import "strconv"

func Convert(number int) string {
    re := ""
    if number % 3 == 0 { re += "Pling" }
    if number % 5 == 0 { re += "Plang" }
    if number % 7 == 0 { re += "Plong" }
    if number % 3 != 0 && number % 5 != 0 && number % 7 != 0 {
        re = strconv.Itoa(number)
    }
    return re
}
