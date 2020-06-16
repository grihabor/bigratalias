package linkedlist

import (
	"io"

	"github.com/clipperhouse/typewriter"
)

func init() {
	err := typewriter.Register(NewBigRatAliasWriter())
	if err != nil {
		panic(err)
	}
}

type BigRatAliasWriter struct{}

func NewBigRatAliasWriter() *BigRatAliasWriter {
	return &BigRatAliasWriter{}
}

func (sw *BigRatAliasWriter) Name() string {
	return "bigratalias"
}

func (sw *BigRatAliasWriter) Imports(t typewriter.Type) (result []typewriter.ImportSpec) {
	result = append(result, )
	return result
}

func (sw *BigRatAliasWriter) Write(w io.Writer, t typewriter.Type) error {
	tag, found := t.FindTag(sw)

	if !found {
		// nothing to be done
		return nil
	}

	tmpl, err := templates.ByTag(t, tag)

	if err != nil {
		return err
	}

	if err := tmpl.Execute(w, t); err != nil {
		return err
	}

	return nil
}
