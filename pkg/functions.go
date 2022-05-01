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


func Tail[T any] (ts []T) []T {
    if len(ts) <= 1 {
	return make([]T, 0)
    }


    out := make([]T, len(ts)-1)
    for i,v := range ts[1:] {
	out[i] = v
    }
    return out
}

// Head :: [a] -> a
func Head[T any] (ts []T) Maybe[T] {
    if len(ts) == 0 {
	return Nothing[T]()
    }
    return Just[T](ts[0])
}


func Reverse[T any](ts []T) []T {
    if len(ts) == 0 {
	return ts
    }

    out := make([]T, len(ts))

    for i, v := range ts {
	out[len(ts) - 1 - i] = v
    }
    return out
}
