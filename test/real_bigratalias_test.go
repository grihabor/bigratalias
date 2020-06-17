// Generated by: main
// TypeWriter: bigratalias
// Directive: +gen on *Real

package test

import "math/big"

type RealTime struct {
	rat *big.Rat
}

func NewRealTime(rat *big.Rat) *RealTime {
	return &RealTime{rat: rat}
}

// Le returns true if x is less than y
func (x *RealTime) Le(y *RealTime) bool {
	return x.rat.Cmp(y.rat) < 0
}

// Leq returns true if x is less than or equal to y
func (x *RealTime) Leq(y *RealTime) bool {
	return x.rat.Cmp(y.rat) <= 0
}

// Ge returns true if x is greater than y
func (x *RealTime) Ge(y *RealTime) bool {
	return x.rat.Cmp(y.rat) > 0
}

// Geq returns true if x is greater than or equal to y
func (x *RealTime) Geq(y *RealTime) bool {
	return x.rat.Cmp(y.rat) >= 0
}

// Eq returns true if x is equal to y
func (x *RealTime) Eq(y *RealTime) bool {
	return x.rat.Cmp(y.rat) == 0
}

// Neq returns true if x is not equal to y
func (x *RealTime) Neq(y *RealTime) bool {
	return x.rat.Cmp(y.rat) != 0
}

func (x *RealTime) Add(y *RealDuration) *RealTime {
	return NewRealTime(new(big.Rat).Add(x.rat, y.rat))
}

func (x *RealTime) Sub(y *RealTime) *RealDuration {
	return NewRealDuration(new(big.Rat).Sub(x.rat, y.rat))
}

func MaxRealTime(x *RealTime, y *RealTime) *RealTime {
	if x.Ge(y) {
		return x
	}
	return y
}

func MinRealTime(x *RealTime, y *RealTime) *RealTime {
	if x.Le(y) {
		return x
	}
	return y
}

func (x *RealTime) Inv() *RealTime {
	return NewRealTime(new(big.Rat).Inv(x.rat))
}

type RealDuration struct {
	rat *big.Rat
}

func NewRealDuration(rat *big.Rat) *RealDuration {
	return &RealDuration{rat: rat}
}

func (x *RealDuration) Mul(y *RealDuration) *RealDuration {
	return NewRealDuration(new(big.Rat).Mul(x.rat, y.rat))
}

// Le returns true if x is less than y
func (x *RealDuration) Le(y *RealDuration) bool {
	return x.rat.Cmp(y.rat) < 0
}

// Leq returns true if x is less than or equal to y
func (x *RealDuration) Leq(y *RealDuration) bool {
	return x.rat.Cmp(y.rat) <= 0
}

// Ge returns true if x is greater than y
func (x *RealDuration) Ge(y *RealDuration) bool {
	return x.rat.Cmp(y.rat) > 0
}

// Geq returns true if x is greater than or equal to y
func (x *RealDuration) Geq(y *RealDuration) bool {
	return x.rat.Cmp(y.rat) >= 0
}

// Eq returns true if x is equal to y
func (x *RealDuration) Eq(y *RealDuration) bool {
	return x.rat.Cmp(y.rat) == 0
}

// Neq returns true if x is not equal to y
func (x *RealDuration) Neq(y *RealDuration) bool {
	return x.rat.Cmp(y.rat) != 0
}

func (x *RealDuration) Neg() *RealDuration {
	return NewRealDuration(new(big.Rat).Neg(x.rat))
}

func (x *RealDuration) Div(y *big.Rat) *RealDuration {
	return NewRealDuration(new(big.Rat).Mul(x.rat, new(big.Rat).Inv(y)))
}

func MaxRealDuration(x *RealDuration, y *RealDuration) *RealDuration {
	if x.Ge(y) {
		return x
	}
	return y
}

func MinRealDuration(x *RealDuration, y *RealDuration) *RealDuration {
	if x.Le(y) {
		return x
	}
	return y
}