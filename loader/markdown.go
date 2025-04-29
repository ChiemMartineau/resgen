package loader

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/goccy/go-yaml"
)

func readMarkdownFile[K any](ctx context.Context, path string) (*K, error) {
	config := ctx.Value(ctxConfigKey{}).(*Config)
	langs := config.Languages

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	text := string(content)
	cursor := 0

	takeUntil := func(delim string) (string, bool) {
		idx := strings.Index(text[cursor:], delim)
		if idx == -1 {
			return "", false
		}
		part := text[cursor : cursor+idx]
		cursor += idx + len(delim)
		return part, true
	}

	formatError := func(msg string) error {
		return fmt.Errorf("unable to read MD at %s: %s", path, msg)
	}

	// Start of file is invalid
	if read, ok := takeUntil("---\n"); read != "" || !ok {
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

	// Get all sections
	var sections []string
	for {
		segment, ok := takeUntil("\n---\n")
		if !ok {
			break
		}
		sections = append(sections, strings.TrimSpace(segment))
	}
	sections = append(sections, strings.TrimSpace(text[cursor:]))

	// Make section(s) into a LocalizedString
	var body LocalizedString
	if len(sections) == 1 {
		body = NewLocalizedStringFomString(langs, sections[0])
	} else {
		body, err = NewLocalizedStringFromSlice(langs, sections)
		if err != nil {
			return nil, formatError(err.Error())
		}
	}

	// Unsafe reflection hack
	bodyField := reflect.ValueOf(&data).Elem().FieldByName("Body")
	bodyField.Set(reflect.ValueOf(body))

	return &data, nil
}

func readMarkdownDir[K any](ctx context.Context, path string) ([]K, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	data := make([]K, len(entries))
	for i, en := range entries {
		el, err := readMarkdownFile[K](ctx, filepath.Join(path, en.Name()))
		if err != nil {
			return nil, err
		}
		data[i] = *el
	}
	return data, nil
}
