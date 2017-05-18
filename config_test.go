package jsonconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

type BenchConfig struct {
	PGDBs      []PGConfig
	Benchmarks []BenchmarkConfig
}

func (c *BenchConfig) Merge(conf interface{}) {
	c.PGDBs = append(c.PGDBs, conf.(BenchConfig).PGDBs...)
	c.Benchmarks = append(c.Benchmarks, conf.(BenchConfig).Benchmarks...)
}

func (c *BenchConfig) Dump() string {
	return DumpConfigs(c)
}

func (c *BenchConfig) MergeFromFile(filePath string) error {

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error reading config file '%s': %s\n", filePath, err)
	}

	var newConf BenchConfig

	err = json.Unmarshal(buf, &newConf)
	if err != nil {
		return fmt.Errorf("Error parsing JSON in config file '%s': %s\n", filePath, err)
	}

	c.Merge(newConf)
	return nil

}

type BenchmarkConfig struct {
	Name string
	DS   string
}

type PGConfig struct {
	Name string
	Host string
}

func TestBenchConfigIsAConfig(t *testing.T) {

	var c Config
	c = &BenchConfig{} // This will bomb if BenchConfig is not a Config
	_ = c.Dump()

}
