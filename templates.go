package linkedlist

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	list,
}

var list = &typewriter.Template{
	Name: "bigratalias",
	Text: `
type {{.Name}}Time struct {
	rat *big.Rat
}

func New{{.Name}}Time(rat *big.Rat) *{{.Name}}Time {
	return &{{.Name}}Time{rat: rat}
}

// Le returns true if x is less than y
func (x *{{.Name}}Time) Le(y *{{.Name}}Time) bool {
	return x.rat.Cmp(y.rat) < 0
}

// Leq returns true if x is less than or equal to y
func (x *{{.Name}}Time) Leq(y *{{.Name}}Time) bool {
	return x.rat.Cmp(y.rat) <= 0
}

// Ge returns true if x is greater than y
func (x *{{.Name}}Time) Ge(y *{{.Name}}Time) bool {
	return x.rat.Cmp(y.rat) > 0
}

// Geq returns true if x is greater than or equal to y
func (x *{{.Name}}Time) Geq(y *{{.Name}}Time) bool {
	return x.rat.Cmp(y.rat) >= 0
}

// Eq returns true if x is equal to y
func (x *{{.Name}}Time) Eq(y *{{.Name}}Time) bool {
	return x.rat.Cmp(y.rat) == 0
}

// Neq returns true if x is not equal to y
func (x *{{.Name}}Time) Neq(y *{{.Name}}Time) bool {
	return x.rat.Cmp(y.rat) != 0
}

func (x *{{.Name}}Time) Add(y *{{.Name}}Duration) *{{.Name}}Time {
	return New{{.Name}}Time(new(big.Rat).Add(x.rat, y.rat))
}

func (x *{{.Name}}Time) Sub(y *{{.Name}}Time) *{{.Name}}Duration {
	return New{{.Name}}Duration(new(big.Rat).Sub(x.rat, y.rat))
}

func Max{{.Name}}Time(x *{{.Name}}Time, y *{{.Name}}Time) *{{.Name}}Time {
	if x.Ge(y) {
		return x
	}
	return y
}

func Min{{.Name}}Time(x *{{.Name}}Time, y *{{.Name}}Time) *{{.Name}}Time {
	if x.Le(y) {
		return x
	}
	return y
}

func (x *{{.Name}}Time) Inv() *{{.Name}}Time {
	return New{{.Name}}Time(new(big.Rat).Inv(x.rat))
}

type {{.Name}}Duration struct {
	rat *big.Rat
}

func New{{.Name}}Duration(rat *big.Rat) *{{.Name}}Duration {
	return &{{.Name}}Duration{rat: rat}
}

func (x *{{.Name}}Duration) Mul(y *{{.Name}}Duration) *{{.Name}}Duration {
	return New{{.Name}}Duration(new(big.Rat).Mul(x.rat, y.rat))
}

// Le returns true if x is less than y
func (x *{{.Name}}Duration) Le(y *{{.Name}}Duration) bool {
	return x.rat.Cmp(y.rat) < 0
}

// Leq returns true if x is less than or equal to y
func (x *{{.Name}}Duration) Leq(y *{{.Name}}Duration) bool {
	return x.rat.Cmp(y.rat) <= 0
}

// Ge returns true if x is greater than y
func (x *{{.Name}}Duration) Ge(y *{{.Name}}Duration) bool {
	return x.rat.Cmp(y.rat) > 0
}

// Geq returns true if x is greater than or equal to y
func (x *{{.Name}}Duration) Geq(y *{{.Name}}Duration) bool {
	return x.rat.Cmp(y.rat) >= 0
}

// Eq returns true if x is equal to y
func (x *{{.Name}}Duration) Eq(y *{{.Name}}Duration) bool {
	return x.rat.Cmp(y.rat) == 0
}

// Neq returns true if x is not equal to y
func (x *{{.Name}}Duration) Neq(y *{{.Name}}Duration) bool {
	return x.rat.Cmp(y.rat) != 0
}

func (x *{{.Name}}Duration) Neg() *{{.Name}}Duration {
	return New{{.Name}}Duration(new(big.Rat).Neg(x.rat))
}

func (x *{{.Name}}Duration) Div(y *big.Rat) *{{.Name}}Duration {
	return New{{.Name}}Duration(new(big.Rat).Mul(x.rat, new(big.Rat).Inv(y)))
}

func Max{{.Name}}Duration(x *{{.Name}}Duration, y *{{.Name}}Duration) *{{.Name}}Duration {
	if x.Ge(y) {
		return x
	}
	return y
}

func Min{{.Name}}Duration(x *{{.Name}}Duration, y *{{.Name}}Duration) *{{.Name}}Duration {
	if x.Le(y) {
		return x
	}
	return y
}
`}
