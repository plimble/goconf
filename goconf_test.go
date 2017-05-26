package goconf

import (
	"os"
	"testing"

	"io/ioutil"

	"github.com/stretchr/testify/require"
)

type SampleA struct {
	A               string
	B               *SampleB
	CamelCase       bool
	ManualOverride1 string `envconfig:"manual_override_1"`
	SplitWord1      string `split_words:"true"`
	ID              string
	DefaultValue    string
}

type SampleB struct {
	C int   `envconfig:"GO_C"`
	D []int `envconfig:"GO_D"`
	E []int
}

var expSample = &SampleA{
	A: "foo",
	B: &SampleB{
		C: 9,
		D: []int{1, 2, 3},
	},
	CamelCase:       true,
	ManualOverride1: "foobar",
	SplitWord1:      "hello world",
	ID:              "123456",
	DefaultValue:    "default",
}

var bytes = []byte(`
a: foo
camelcase: true
b:
  c: 9
  d:
  - 1
  - 2
  - 3
manualoverride1: "foobar"
splitword1: "hello world"
id: 123456
`)

func TestEnv(t *testing.T) {
	os.Setenv("GO_A", "foo")
	os.Setenv("GO_CAMELCASE", "true")
	os.Setenv("GO_ID", "123456")
	os.Setenv("GO_D", "1,2,3")
	os.Setenv("GO_C", "9")
	os.Setenv("GO_SPLIT_WORD1", "hello world")
	os.Setenv("GO_MANUAL_OVERRIDE_1", "foobar")

	sample := &SampleA{A: "baz", DefaultValue: "default"}
	err := Parse(sample, WithEnv("go"))
	require.NoError(t, err)
	require.EqualValues(t, expSample, sample)
}

func TestYaml(t *testing.T) {
	sample := &SampleA{DefaultValue: "default"}
	err := Parse(sample, WithYamlFromBytes(bytes))
	require.NoError(t, err)
	require.EqualValues(t, expSample, sample)
}

func TestYamlFromFile(t *testing.T) {
	err := ioutil.WriteFile("config1.yml", bytes, 0777)
	require.NoError(t, err)

	sample := &SampleA{DefaultValue: "default"}
	err = Parse(sample, WithYaml("config1.yml"))
	require.NoError(t, err)
	require.EqualValues(t, expSample, sample)
}

func TestCombile(t *testing.T) {
	os.Setenv("GOO_A", "baz")

	cbytes := []byte(`
a: foo
b:
  e:
  - 3
  - 3
  - 3
manualoverride1: "foobar"
`)

	cfile := []byte(`
a: "bar"
b:
  e:
  - 4
  - 4
  - 4
`)

	cexp := &SampleA{
		A: "baz",
		B: &SampleB{
			C: 9,
			D: []int{1, 2, 3},
			E: []int{4, 4, 4},
		},
		ManualOverride1: "foobar",
		DefaultValue:    "default",
	}

	sample := &SampleA{
		A: "000",
		B: &SampleB{
			C: 9,
			D: []int{0, 0, 0},
			E: []int{0, 0, 0},
		},
		ManualOverride1: "111",
		DefaultValue:    "default",
	}

	err := ioutil.WriteFile("config2.yml", cfile, 0777)
	require.NoError(t, err)

	err = Parse(sample,
		WithYamlFromBytes(cbytes),
		WithYaml("config2.yml"),
		WithEnv("goo"),
	)
	require.NoError(t, err)
	require.EqualValues(t, cexp, sample)
}
