package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	Addr string `json:"addr" env-default:":8080"`
}

func GetConfig(path string) (*Config, error) {
	c := &Config{}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		return nil, err
	}

	return c, nil

}
