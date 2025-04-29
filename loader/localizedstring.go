package loader

import (
	"context"
	"errors"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
)

type LocalizedString map[string]string

func NewLocalizedStringFomString(langs []string, value string) LocalizedString {
	out := LocalizedString(make(map[string]string, len(langs)))
	for _, lang := range langs {
		out[lang] = value
	}
	return out
}

func NewLocalizedStringFromSlice(langs []string, values []string) (LocalizedString, error) {
	if len(values) != len(langs) {
		return nil, errors.New("cannot constructot localized string, mismatched number of values and languages")
	}
	out := LocalizedString(make(map[string]string, len(langs)))
	for i, lang := range langs {
		out[lang] = values[i]
	}
	return out, nil
}

func (s *LocalizedString) UnmarshalYAML(ctx context.Context, node ast.Node) error {
	config := ctx.Value(ctxConfigKey{}).(*Config)
	langs := config.Languages

	switch n := node.(type) {
	case *ast.StringNode:
		*s = NewLocalizedStringFomString(langs, n.String())

	case *ast.SequenceNode:
		values := make([]string, len(langs))

		for i, v := range n.Values {
			values[i] = v.String()
		}

		var err error
		*s, err = NewLocalizedStringFromSlice(langs, values)

		if err != nil {
			return &yaml.SyntaxError{
				Message: err.Error(),
				Token:   n.End,
			}
		}

	default:
		return &yaml.UnexpectedNodeTypeError{
			Actual:   node.Type(),
			Expected: ast.StringType,
			Token:    node.GetToken(),
		}
	}

	return nil
}
