package loader

import (
	"context"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
)

type ctxDateFormatKey struct{}

type Date time.Time

func (d *Date) UnmarshalYAML(ctx context.Context, node ast.Node) error {
	dateFormat := ctx.Value(ctxDateFormatKey{}).(string)

	switch node := node.(type) {
	case *ast.StringNode:
		parsed, err := time.Parse(dateFormat, node.Value)
		if err != nil {
			return &yaml.SyntaxError{
				Token:   node.GetToken(),
				Message: err.Error(),
			}
		}
		*d = Date(parsed)
	default:
		return &yaml.UnexpectedNodeTypeError{
			Actual:   node.Type(),
			Expected: ast.StringType,
			Token:    node.GetToken(),
		}
	}

	return nil
}
