package pkg

func Abs[T SignedNumber] (ts []T) []T {
    out := make([]T, len(ts))
    for i,t := range ts {
	if t < 0 {
	    t *= -1
	}
	out[i] = t
    }
    return out
}


func Min[T SignedNumber | UnsignedNumber] (ts []T) T {
    if len(ts) == 0 {
	return 0 // should this be min by type?
    }

    min := ts[0]
    for _,t := range ts[1:] {
	if t < min {
	    min = t
	}
    }
    return min
}

func Max[T SignedNumber | UnsignedNumber] (ts []T) T {
    if len(ts) == 0 {
	return 0
    }

    max := ts[0]

    for _,t := range ts[1:] {
	if t > max {
	    max = t
	}
    }
    return max
}
