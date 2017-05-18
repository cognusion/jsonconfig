package jsonconfig

import (
	"encoding/json"
	"path/filepath"
)

type Config interface {
	Merge(interface{})
	Dump() string
	MergeFromFile(string) error
}

// Return a formatted JSON string representation of the config
func DumpConfigs(conf Config) string {
	j, _ := json.MarshalIndent(conf, "", "\t")
	return string(j)
}

// Given a directory, load all the configs
func LoadJsonConfigs(srcDir string, conf Config) {
	for _, f := range readDirectory(srcDir, "*.json") {
		_ = conf.MergeFromFile(f) // ignoring errors?!
	}
}

func readDirectory(srcDir, pattern string) []string {
	// We can skip this error, since our pattern is fixed and known-good.
	files, _ := filepath.Glob(srcDir + pattern)
	return files
}
