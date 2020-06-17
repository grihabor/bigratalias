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
	return x.rat.Cmp(y.rat) < 0
}

// Leq returns true if x is less than or equal to y
func (x {{.Pointer}}{{.Name}}) Leq(y {{.Pointer}}{{.Name}}) bool {
	return x.rat.Cmp(y.rat) <= 0
}

// Ge returns true if x is greater than y
func (x {{.Pointer}}{{.Name}}) Ge(y {{.Pointer}}{{.Name}}) bool {
	return x.rat.Cmp(y.rat) > 0
}

// Geq returns true if x is greater than or equal to y
func (x {{.Pointer}}{{.Name}}) Geq(y {{.Pointer}}{{.Name}}) bool {
	return x.rat.Cmp(y.rat) >= 0
}

// Eq returns true if x is equal to y
func (x {{.Pointer}}{{.Name}}) Eq(y {{.Pointer}}{{.Name}}) bool {
	return x.rat.Cmp(y.rat) == 0
}

// Neq returns true if x is not equal to y
func (x {{.Pointer}}{{.Name}}) Neq(y {{.Pointer}}{{.Name}}) bool {
	return x.rat.Cmp(y.rat) != 0
}

func (x {{.Pointer}}{{.Name}}) Mul(y {{.Pointer}}{{.Name}}) {{.Pointer}}{{.Name}} {
	return New{{.Name}}(new(big.Rat).Mul(x.rat, y.rat))
}

func (x {{.Pointer}}{{.Name}}) Neg() {{.Pointer}}{{.Name}} {
	return New{{.Name}}(new(big.Rat).Neg(x.rat))
}

func (x {{.Pointer}}{{.Name}}) Add(y {{.Pointer}}{{.Name}}) {{.Pointer}}{{.Name}} {
	return New{{.Name}}(new(big.Rat).Add(x.rat, y.rat))
}

func (x {{.Pointer}}{{.Name}}) Sub(y {{.Pointer}}{{.Name}}) {{.Pointer}}{{.Name}} {
	return New{{.Name}}(new(big.Rat).Sub(x.rat, y.rat))
}

func (x {{.Pointer}}{{.Name}}) Div(y {{.Pointer}}{{.Name}}) {{.Pointer}}{{.Name}} {
	return x.Mul(y.Inv())
}

func (x {{.Pointer}}{{.Name}}) AddInt64(y int64) {{.Pointer}}{{.Name}} {
	return x.Add(New{{.Name}}(big.NewRat(y, 1)))
}

func (x {{.Pointer}}{{.Name}}) SubInt64(y int64) {{.Pointer}}{{.Name}} {
	return x.Sub(New{{.Name}}(big.NewRat(y, 1)))
}

func (x {{.Pointer}}{{.Name}}) Max(y {{.Pointer}}{{.Name}}) {{.Pointer}}{{.Name}} {
	if x.Ge(y) {
		return x
	}
	return y
}

func (x {{.Pointer}}{{.Name}}) Inv() {{.Pointer}}{{.Name}} {
	return New{{.Name}}(new(big.Rat).Inv(x.rat))
}
`}
