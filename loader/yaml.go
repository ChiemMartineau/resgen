package loader

import (
	"context"
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type YamlErrorSource struct {
	Path       string
	LineOffset int
	CharOffset int
}

type YamlError struct {
	YamlError yaml.Error
	Source    YamlErrorSource
}

func (e YamlError) Error() string {
	pos := e.YamlError.GetToken().Position
	return fmt.Sprintf(
		"unable to read YAML at %s:%d:%d: %s",
		e.Source.Path,
		e.Source.LineOffset+pos.Line,
		e.Source.CharOffset+pos.Column,
		e.YamlError.GetMessage(),
	)
}

func readYamlFile[K any](ctx context.Context, path string) (*K, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var data K
	if err := yaml.UnmarshalContext(ctx, content, &data, yaml.Strict()); err != nil {
		return nil, YamlError{
			YamlError: err.(yaml.Error),
			Source:    YamlErrorSource{Path: path},
		}
	}

	return &data, nil
}
