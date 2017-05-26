goconf
---
[![GoDoc](https://godoc.org/plimble/goconf?status.svg)](https://godoc.org/github.com/plimble/goconf)
[![Build Status](https://travis-ci.org/plimble/goconf.svg?branch=master)](https://travis-ci.org/plimble/goconf?branch=master)
[![Coverage Status](https://coveralls.io/repos/plimble/goconf/badge.svg?branch=master&service=github&foo)](https://coveralls.io/github/plimble/goconf?branch=master)
[![Go Report Card](https://goreportcard.com/badge/plimble/goconf)](https://goreportcard.com/report/plimble/goconf)

Combine yaml and environment config

## Features
- [Parse yaml](gopkg.in/yaml.v2)
- [Parse env](github.com/kelseyhightower/envconfig)
- Watch yaml config file

## Installation

```
go get gopkg.in/plimble/goconf.v1
```

## Env format

```go
type SampleA struct {
	A               string // PREFIX_A
	CamelCase       bool // PREFIX_CAMELCASE
	ManualOverride1 string `envconfig:"manual_override_1"` // PREFIX_MANUAL_OVERRIDE_1
	SplitWord1      string `split_words:"true"` // SPLIT_WORD1
	ID              string // PREFIX_ID
	DefaultValue    string `envconfig:"DEFAULT_VALUE"` // PREFIX_DEFAULT_VALUE
}
```
## Yaml format

```go
type SampleA struct {
	A               string `json:"abc"` // abc
	CamelCase       bool `yaml:"cc"` // cc
	ManualOverride1 string // manualoverride1
	SplitWord1      string // splitword1
	ID              string // id
	DefaultValue    string // defaultvalue
}
```

## Example

```go

type Sample struct {
  Value string
}

var bytes = `
value: v1
`

sample := &Sample{}
// Parse Order: yaml bytes -> yaml file -> env
err = goconf.Parse(sample,
    WithEnv("prefix"),
    WithYamlFromBytes(bytes),
    WithYaml("path.yml"),
)
```

## Watch Config

```go

goconf.WatchYamlFile("path.yml", sample, func() error {
  fmt.Println("event on chane")
  return nil
})

// or ignore event
goconf.WatchYamlFile("path.yml", sample, nil)

```
