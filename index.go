package main

// Convert decimal to new base.
// eg - 5 (base 10) => 10 (base 5)
// eg =
//      34 (base 10) => 114 (base 5)
//      114 (base 5) => 34 (base 10)

type Integer interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}


func decimalTo(v int, base int) []int {
    if v <= 0 {
        return []int{0}
    }
	var r int
	outs := make([]int, 0, 1e6)
	for v > 0 {
		v, r = divmod(v, base)
		outs = append([]int{r}, outs...)
	}
	return outs
}

// 98 (10) => 1202 (4)
// 1202 (4) => 98 (10)
// Convert number from old base to base 10.
func toDecimal[N Integer](n N, oldBase N) N {
    var o N
    var digits []N
    digits = getDigits(n)
    for i, digit := range digits {
        o += (digit * pow(oldBase, i))
    }
    return o
}

func divmod[N Integer](x, y N) (N, N) {
	return x / y, x % y
}

func getDigits[N Integer](v N) []N {
	o := make([]N, 0, 1e6)
	for v > 0 {
        // For digits in "correct" order, use first append.
		// o = append([]N{v % 10}, o...)
        o = append(o, v%10)
		v /= 10
	}
	return o
}

func pow[N Integer](b N, expon int) N {
    var o N
    o = 1
    for expon > 0 {
        o *= b
        expon--
    }
    return o
}
