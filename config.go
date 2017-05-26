package goconf

type Option func(o *opts)

type opts struct {
	yaml      bool
	yamlPath  string
	yamlBytes []byte
	env       bool
	envPrefix string
}

func WithYaml(path string) Option {
	return func(o *opts) {
		o.yaml = true
		o.yamlPath = path
	}
}

func WithYamlFromBytes(yaml []byte) Option {
	return func(o *opts) {
		o.yaml = true
		o.yamlBytes = yaml
	}
}

func WithEnv(prefix string) Option {
	return func(o *opts) {
		o.env = true
		o.envPrefix = prefix
	}
}
