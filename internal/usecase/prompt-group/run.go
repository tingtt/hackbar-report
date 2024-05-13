package promptgroup

import (
	"fmt"
	"hackbar-report/internal/domain/prompt"
	"io"
	"reflect"

	"github.com/fatih/color"
)

func Run[T comparable](out io.Writer, in io.Reader, prompt *T) error {
	v := reflect.ValueOf(prompt).Elem()

	return iterateFields(v, func(fieldName string, tag reflect.StructTag, value reflect.Value) error {
		if value.Kind() != reflect.Struct {
			return nil
		}

		heading := label(fieldName /* default */, tag)
		_, err := fmt.Fprintln(out, color.GreenString("%s", heading))
		if err != nil {
			return err
		}

		return iterateFields(value, promptYield(out, in))
	})
}

func iterateFields(rv reflect.Value, yield func(fieldName string, tag reflect.StructTag, value reflect.Value) error) error {
	rt := rv.Type()

	for i := range make([]interface{}, rt.NumField()) {
		field := rt.Field(i)
		err := yield(field.Name, field.Tag, rv.FieldByName(field.Name))
		if err != nil {
			return err
		}
	}

	return nil
}

func promptYield(out io.Writer, in io.Reader) func(fieldName string, tag reflect.StructTag, value reflect.Value) error {
	return func(fieldName string, tag reflect.StructTag, value reflect.Value) error {
		if value.Kind() != reflect.String {
			return nil
		}

		message := label(fieldName /* default */, tag)
		suffix := suffix(tag)
		p := prompt.New(fmt.Sprintf("  %s%s:", message, suffix))
		answer, err := p.Run(out, in)
		if err != nil {
			return err
		}

		value.SetString(answer)
		return nil
	}
}

func label(defaultLabel string, tag reflect.StructTag) string {
	label, ok := tag.Lookup("label")
	if !ok {
		label = defaultLabel
	}
	return label
}

func suffix(tag reflect.StructTag) string {
	suffix, ok := tag.Lookup("suffix")
	if !ok {
		return ""
	}
	return suffix
}
