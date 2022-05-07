package models

import "fmt"

var (
	Levels = []string{
		"Halqa", "Majlis", "Region", "Ilaqa", "National",
	}
	Tanzeem = map[string]string{
		"K": "Khuddam",
		"T": "Atfal",
		"A": "Ansar",
	}
)

func ColsMap(cols []string, keyPrefix string, keyPostfix string, valPrefix string, valPostfix string) map[string]string {
	result := make(map[string]string)
	for _, k := range cols {
		result[fmt.Sprintf("%s%s%s", keyPrefix, k, keyPostfix)] = fmt.Sprintf("%s%s%s", valPrefix, k, valPostfix)
	}

	return result
}
