package goconf

import (
	"io/ioutil"
	"testing"

	"time"

	"github.com/stretchr/testify/require"
)

func TestWatch(t *testing.T) {
	var bytes2 = []byte(`
a: foo1
camelcase: false
b:
  c: 2
  d:
  - 3
  - 2
  - 1
manualoverride1: "foobar1"
splitword1: "hello world1"
id: 1234567
`)

	err := ioutil.WriteFile("config3.yml", bytes, 0777)
	require.NoError(t, err)

	sample := &SampleA{DefaultValue: "default"}
	err = Parse(sample, WithYaml("config3.yml"))
	require.NoError(t, err)
	require.EqualValues(t, expSample, sample)

	WatchYamlFile("cofig3.yml", sample, func() error {
		require.EqualValues(t, &SampleA{
			A: "foo1",
			B: &SampleB{
				C: 2,
				D: []int{3, 2, 1},
			},
			ManualOverride1: "foobar1",
			SplitWord1:      "hello world1",
			ID:              "1234567",
			DefaultValue:    "default",
		}, sample)

		return nil
	})

	time.Sleep(time.Second * 1)
	err = ioutil.WriteFile("config3.yml", bytes2, 0777)
	require.NoError(t, err)
}
