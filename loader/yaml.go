package loader

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"

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
	if err := yaml.UnmarshalContext(ctx, content, &data); err != nil {
		return nil, YamlError{
			YamlError: err.(yaml.Error),
			Source:    YamlErrorSource{Path: path},
		}
	}

	return &data, nil
}

func readMarkdownFile[K any](ctx context.Context, path string) (*K, error) {
	langs := ctx.Value(ctxLanguagesKey{}).([]string)

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	text := string(content)
	cursor := 0

	takeUntil := func(token string) (string, bool) {
		idx := strings.Index(text[cursor:], token)
		if idx == -1 {
			return "", false
		}
		part := text[cursor : cursor+idx]
		cursor += idx + len(token)
		return part, true
	}

	formatError := func(msg string) error {
		return fmt.Errorf("unable to read MD at %s: %s", path, msg)
	}

	if _, ok := takeUntil("---\n"); !ok {
		return nil, formatError("misformatted frontmatter (missing start)")
	}

	frontmatter, ok := takeUntil("\n---\n")
	if !ok {
		return nil, formatError("misformatted frontmatter (missing end)")
	}

	var data K
	if err := yaml.UnmarshalContext(ctx, []byte(frontmatter), &data, yaml.Strict()); err != nil {
		return nil, YamlError{
			YamlError: err.(yaml.Error),
			Source: YamlErrorSource{
				Path:       path,
				LineOffset: 1,
			},
		}
	}

	var sections []string
	for {
		segment, ok := takeUntil("\n---\n")
		if !ok {
			break
		}
		sections = append(sections, strings.TrimSpace(segment))
	}

	if rest := strings.TrimSpace(text[cursor:]); rest != "" {
		sections = append(sections, rest)
	}

	var body LocalizedString
	if len(sections) == 1 {
		body = NewLocalizedStringFomString(langs, sections[0])
	} else {
		body, err = NewLocalizedStringFromSlice(langs, sections)
		if err != nil {
			return nil, err
		}
	}

	bodyField := reflect.ValueOf(&data).Elem().FieldByName("Body")
	bodyField.Set(reflect.ValueOf(body))

	return &data, nil
}
