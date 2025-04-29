package loader

import (
	"context"
	"path/filepath"
)

type Config struct {
	Languages  []string `yaml:"languages"`
	DateFormat string   `yaml:"date_format"`
}
type ctxConfigKey struct{}

type Experience struct {
	Start    Date            `yaml:"start"`
	End      Date            `yaml:"end"`
	Location LocalizedString `yaml:"location"`
	Link     string          `yaml:"link"`
	Role     LocalizedString `yaml:"role"`
	Company  LocalizedString `yaml:"company"`
	Body     LocalizedString
}

type Education struct {
	Start        Date              `yaml:"start"`
	End          Date              `yaml:"end"`
	Location     LocalizedString   `yaml:"location"`
	Institution  LocalizedString   `yaml:"institution"`
	Program      LocalizedString   `yaml:"program"`
	Grade        LocalizedString   `yaml:"grade"`
	Distinctions []LocalizedString `yaml:"distinctions"`
	Body         LocalizedString
}

type Project struct {
	Start    Date              `yaml:"start"`
	End      Date              `yaml:"end"`
	Location LocalizedString   `yaml:"location"`
	Link     string            `yaml:"link"`
	Title    LocalizedString   `yaml:"title"`
	Tools    []LocalizedString `yaml:"tools"`
	Body     LocalizedString
}

type Data struct {
	Experiences []Experience
	Education   []Education
	Projects    []Project
}

func loadConfig(rootPath string) (*Config, error) {
	path := filepath.Join(rootPath, "resgen.yaml")

	cfg, err := readYamlFile[Config](context.Background(), path)

	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func LoadData(rootPath string) (*Data, error) {
	cfg, err := loadConfig(rootPath)
	if err != nil {
		return nil, err
	}

	ctx := context.WithValue(context.Background(), ctxConfigKey{}, cfg)

	data := &Data{}

	data.Education, err = readMarkdownDir[Education](ctx, filepath.Join(rootPath, "education"))
	if err != nil {
		return nil, err
	}

	data.Experiences, err = readMarkdownDir[Experience](ctx, filepath.Join(rootPath, "experiences"))
	if err != nil {
		return nil, err
	}

	data.Projects, err = readMarkdownDir[Project](ctx, filepath.Join(rootPath, "projects"))
	if err != nil {
		return nil, err
	}

	return data, nil
}
