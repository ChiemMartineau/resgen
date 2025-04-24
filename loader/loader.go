package loader

import (
	"context"
	"path/filepath"
)

type Config struct {
	Languages  []string `yaml:"languages"`
	DateFormat string   `yaml:"date_format"`
}

type Job struct {
	Body    LocalizedString
	Start   Date            `yaml:"start"`
	End     Date            `yaml:"end"`
	Link    string          `yaml:"link"`
	Title   LocalizedString `yaml:"title"`
	Program LocalizedString `yaml:"program"`
}

func loadConfig(rootPath string) (*Config, error) {
	path := filepath.Join(rootPath, "resgen.yaml")

	cfg, err := readYamlFile[Config](context.Background(), path)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadData(rootPath string) (*Job, error) {
	cfg, err := loadConfig(rootPath)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	ctx = context.WithValue(ctx, ctxLanguagesKey{}, cfg.Languages)
	ctx = context.WithValue(ctx, ctxDateFormatKey{}, cfg.DateFormat)

	job, err := readMarkdownFile[Job](ctx, filepath.Join(rootPath, "education", "bsc.md"))
	if err != nil {
		return nil, err
	}

	return job, nil
}
