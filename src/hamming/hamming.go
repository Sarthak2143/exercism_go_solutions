package hamming

import "errors"

func Distance(a, b string) (int, error) {
    rune_a := []rune(a)
    rune_b := []rune(b)
    if len(rune_a) != len(rune_b) {
        return 0, errors.New("Different length.")
    }
    count := 0
    for i := 0; i < len(rune_a); i++ {
        if rune_a[i] != rune_b[i] {
            count++
        }
    }
        
    return count, nil
}
