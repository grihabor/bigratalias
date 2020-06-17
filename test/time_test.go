package test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

// +gen * bigratalias
type Real struct {
	rat *big.Rat
}

func NewTime(r *big.Rat) *Real {
	return &Real{rat: r}
}

func TestRealTime(t *testing.T) {
	ts := NewRealTime(big.NewRat(10, 1))
	dt := NewRealDuration(big.NewRat(1, 2))
	result := ts.Add(dt)
	assert.True(t, result.Eq(NewRealTime(big.NewRat(21, 2))))
}

func TestDurationAdd(t *testing.T) {
	dt := NewRealDuration(big.NewRat(1, 3))
	assert.True(t, dt.Add(dt).Eq(NewRealDuration(big.NewRat(2, 3))))
}
