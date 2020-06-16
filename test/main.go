package test

import "math/big"

// +gen BigRatAlias
type FrameRate struct {
    *big.Rat
}
