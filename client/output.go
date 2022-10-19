package client

import (
	"encoding/json"
	"strings"

	"gopkg.in/yaml.v2"
)

type OutputFormat interface {
	//ToConsole(color bool) string
	ToConsole() string
}

func Output(content OutputFormat, formatType string) (string, error) {
	switch strings.ToLower(formatType) {
	case "json":
		sj, err := json.Marshal(content)
		if err != nil {
			return "", err
		}
		return string(sj), nil
	case "yaml":
		sy, err := yaml.Marshal(content)
		if err != nil {
			return "", err
		}
		return string(sy), nil
	// case "console-color":
	// 	return content.ToConsole(true), nil
	default:
		return content.ToConsole(), nil
	}
}
