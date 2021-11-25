package config

import (
	"io/ioutil"
	"os"

	corev1 "k8s.io/api/core/v1"
)

const (
	DefaultMilvusVersion   = "v2.0.0-rc8-20211104-d1f4106"
	DefaultMilvusBaseImage = "milvusdb/milvus"
	DefaultMilvusImage     = DefaultMilvusBaseImage + ":" + DefaultMilvusVersion
	DefaultImagePullPolicy = corev1.PullIfNotPresent
	MilvusConfigTpl        = "milvus.yaml.tmpl"
)

const (
	TemplateRelativeDir = "config/assets/templates"
	ChartDir            = "config/assets/charts"
	ProviderName        = "milvus-operator"
)

var (
	workDir       = ""
	defaultConfig *Config
)

func Init() error {
	c, err := NewConfig(workDir)
	if err != nil {
		return err
	}
	defaultConfig = c
	if os.Getenv("DEBUG") == "true" {
		defaultConfig.debugMode = true
	}

	return nil
}

func IsDebug() bool {
	return defaultConfig.debugMode
}

func GetMilvusConfigTemplate() string {
	return defaultConfig.GetTemplate(MilvusConfigTpl)
}

type Config struct {
	debugMode bool
	templates map[string]string
}

func NewConfig(workDir string) (*Config, error) {
	config := &Config{
		templates: make(map[string]string),
	}

	templateDir := workDir + TemplateRelativeDir

	tmpls, err := ioutil.ReadDir(templateDir)
	if err != nil {
		return nil, err
	}
	for _, tmpl := range tmpls {
		data, err := ioutil.ReadFile(templateDir + "/" + tmpl.Name())
		if err != nil {
			return nil, err
		}

		config.templates[tmpl.Name()] = string(data)
	}

	return config, nil
}

func (c Config) GetTemplate(name string) string {
	return c.templates[name]
}
