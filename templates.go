package linkedlist

import "github.com/clipperhouse/typewriter"

var templates = typewriter.TemplateSlice{
	list,
}

var list = &typewriter.Template{
	Name: "BigRatAlias",
	Text: `
// Le returns true if x is less than y
func (x {{.Pointer}}{{.Name}}) Le(y {{.Pointer}}{{.Name}}) bool {
	return x.Cmp(y) < 0
}

// Leq returns true if x is less than or equal to y
func (x {{.Pointer}}{{.Name}}) Leq(y {{.Pointer}}{{.Name}}) bool {
	return x.Cmp(y) <= 0
}

// Ge returns true if x is greater than y
func (x {{.Pointer}}{{.Name}}) Ge(y {{.Pointer}}{{.Name}}) bool {
	return x.Cmp(y) > 0
}

// Geq returns true if x is greater than or equal to y
func (x {{.Pointer}}{{.Name}}) Geq(y {{.Pointer}}{{.Name}}) bool {
	return x.Cmp(y) >= 0
}

// Eq returns true if x is equal to y
func (x {{.Pointer}}{{.Name}}) Eq(y {{.Pointer}}{{.Name}}) bool {
	return x.Cmp(y) == 0
}

// Neq returns true if x is not equal to y
func (x {{.Pointer}}{{.Name}}) Neq(y {{.Pointer}}{{.Name}}) bool {
	return x.Cmp(y) != 0
}

func NeqInt64(x {{.Pointer}}{{.Name}}, y int64) bool {
	return Neq(x, Int64(y))
}

func Int64(i int64) {{.Pointer}}{{.Name}} {
	return big.NewRat(i, 0)
}

func Seconds(d time.Duration) {{.Pointer}}{{.Name}} {
	return big.NewRat(d.Nanoseconds(), int64(time.Second))
}

func UInt64(i uint64) {{.Pointer}}{{.Name}} {
	return big.NewRat(int64(i), 0)
}

func (x {{.Pointer}}{{.Name}}) Mul(y {{.Pointer}}{{.Name}}) {{.Pointer}}{{.Name}} {
	return new(big.Rat).Mul(x, y)
}

func Neg(x {{.Pointer}}{{.Name}} ) {{.Pointer}}{{.Name}} {
	return new(big.Rat).Neg(x)
}

func Unix(t time.Time) {{.Pointer}}{{.Name}} {
	return big.NewRat(t.UnixNano(), 1e9)
}

func (x {{.Pointer}}{{.Name}}) Add(y {{.Pointer}}{{.Name}}) {{.Pointer}}{{.Name}} {
	return new(big.Rat).Add(x, y)
}

func (x {{.Pointer}}{{.Name}}) Sub(y {{.Pointer}}{{.Name}}) {{.Pointer}}{{.Name}} {
	return new(big.Rat).Sub(x, y)
}

func (x {{.Pointer}}{{.Name}}) Div(y {{.Pointer}}{{.Name}}) {{.Pointer}}{{.Name}} {
	return Mul(x, Inv(y))
}

func AddInt64(x {{.Pointer}}{{.Name}}, y int64) {{.Pointer}}{{.Name}} {
	return Add(x, Int64(y))
}

func SubInt64(x {{.Pointer}}{{.Name}}, y int64) {{.Pointer}}{{.Name}} {
	return Sub(x, Int64(y))
}

func (x {{.Pointer}}{{.Name}}) Max(y {{.Pointer}}{{.Name}}) {{.Pointer}}{{.Name}} {
	if Ge(x, y) {
		return x
	}
	return y
}

func Inv(x {{.Pointer}}{{.Name}}) {{.Pointer}}{{.Name}} {
	return new(big.Rat).Inv(x)
}

`}
