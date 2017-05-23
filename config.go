package jsonconfig

import (
	"encoding/json"
	"path/filepath"
)

// Config defines a set of methods for operating on Configs
type Config interface {
	Merge(interface{})
	MergeFromFile(string) error
	Dump() string
}

// DumpConfigs returns a formatted JSON string representation of the config
func DumpConfigs(conf Config) string {
	j, _ := json.MarshalIndent(conf, "", "\t")
	return string(j)
}

// LoadJsonConfigs loads all the .json configs
func LoadJsonConfigs(srcDir string, conf Config) (err error) {
	for _, f := range readDirectory(srcDir, "*.json") {
		err = conf.MergeFromFile(f) // ignoring errors?!
		if err != nil {
			break
		}
	}
	return
}

func readDirectory(srcDir, pattern string) []string {
	// We can skip this error, since our pattern is fixed and known-good.
	files, _ := filepath.Glob(srcDir + pattern)
	return files
}
