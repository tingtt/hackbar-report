package markdown

import (
	"hackbar-report/internal/interface-adapter/markdown/components"
	"reflect"
	"strings"
)

const (
	TAG_TYPE          = "mdblk-type"
	TAG_DEFAULT       = "mdblk-default"
	TAG_SEPARATE_WITH = "mdblk-list-separate-with"
	TAG_FORMAT        = "mdblk-format"
)

func lookup(tag reflect.StructTag, key string, defaultValue string) string {
	label, ok := tag.Lookup(key)
	if !ok {
		label = defaultValue
	}
	return label
}

func lookupDefault(tag reflect.StructTag, defaultValue string) (string, bool) {
	value := lookup(tag, TAG_DEFAULT, defaultValue)
	return value, defaultValue != value
}

func label(tag reflect.StructTag, defaultValue string) string {
	return lookup(tag, "label", defaultValue /* default */)
}

func isSkippable(tag reflect.StructTag) bool {
	return strings.HasSuffix(lookup(tag, TAG_TYPE, ""), ",omitempty")
}

func isList(tag reflect.StructTag) (bool, []components.ListOptionApplier) {
	if !strings.HasPrefix(lookup(tag, TAG_TYPE, ""), "list") {
		return false, nil
	}

	optionAppliers := make([]components.ListOptionApplier, 0, 1)

	if separator := lookup(tag, TAG_SEPARATE_WITH, ""); separator != "" {
		optionAppliers = append(optionAppliers, components.WithSeparateBy(separator))
	}

	return true, optionAppliers
}

func hasFormat(tag reflect.StructTag) (bool, []components.TextOptionApplier) {
	format := lookup(tag, TAG_FORMAT, "")
	if format == "" {
		return false, nil
	}
	return true, []components.TextOptionApplier{components.WithFormat(format)}
}
