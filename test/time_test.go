package test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

// +gen * bigratalias
type Time struct {
	rat *big.Rat
}

func NewTime(r *big.Rat) *Time {
	return &Time{rat: r}
}

func TestTime(t *testing.T) {
	fr := &Time{rat: big.NewRat(1, 2)}
	assert.Equal(t, big.NewRat(-1, 2), fr.SubInt64(1).rat)
}
