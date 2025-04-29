package loader

import (
	"context"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
)

type Date time.Time

func (d *Date) UnmarshalYAML(ctx context.Context, node ast.Node) error {
	config := ctx.Value(ctxConfigKey{}).(*Config)

	switch node := node.(type) {
	case *ast.StringNode:
		parsed, err := time.Parse(config.DateFormat, node.Value)
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
