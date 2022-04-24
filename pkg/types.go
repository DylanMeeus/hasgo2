package pkg

type SignedNumber interface {
    int | int8 | int16 | int32 | int64 | float32 | float64  
}

type UnsignedNumber interface {
    uint16 | uint32 | uint64
}

type ComplexNumber interface {
    complex64 | complex128
}
