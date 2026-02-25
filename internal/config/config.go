package config

import (
	"encoding/json"
	"os"
	"strings"
)

type Config interface {
	GetString(name string) (value string, found bool)
	GetStringDefault(name, defVal string) string

	GetInt(name string) (value int, found bool)
	GetIntDefault(name string, defVal int) int

	GetBool(name string) (value bool, found bool)
	GetBoolDefault(name string, defVal bool) bool

	GetFloat(name string) (value float64, found bool)
	GetFloatDefault(name string, defVal float64) float64

	GetSection(name string) (value Config, found bool)
}

type DefaultConfig struct {
	configData map[string]interface{}
}

func New(filename string) (Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	m := map[string]interface{}{}
	err = decoder.Decode(&m)
	if err != nil {
		return nil, err
	}
	cfg := &DefaultConfig{configData: m}
	return cfg, nil
}

func (c *DefaultConfig) get(name string) (value interface{}, found bool) {
	data := c.configData
	for _, key := range strings.Split(name, ":") {
		value, found = data[key]
		// if section
		if newSection, ok := value.(map[string]interface{}); ok && found {
			data = newSection
			//if line
		} else {
			return
		}
	}
	return
}

func (c *DefaultConfig) GetSection(name string) (section Config, found bool) {
	value, found := c.get(name)
	if sectionData, ok := value.(map[string]interface{}); ok && found {
		return &DefaultConfig{sectionData}, true
	}
	return nil, false
}

// GetBool implements [Config].
func (c *DefaultConfig) GetBool(name string) (bool, bool) {
	if value, found := c.get(name); found {
		return value.(bool), found
	}
	return false, false
}

// GetBoolDefault implements [Config].
func (c *DefaultConfig) GetBoolDefault(name string, defVal bool) bool {
	if value, found := c.GetBool(name); found {
		return value
	}
	return defVal
}

// GetFloat implements [Config].
func (c *DefaultConfig) GetFloat(name string) (value float64, found bool) {
	if value, found := c.get(name); found {
		return value.(float64), found
	}
	return 0, false
}

// GetFloatDefault implements [Config].
func (c *DefaultConfig) GetFloatDefault(name string, defVal float64) float64 {
	if value, found := c.GetFloat(name); found {
		return value
	}
	return defVal
}

// GetInt implements [Config].
func (c *DefaultConfig) GetInt(name string) (value int, found bool) {
	if value, found := c.get(name); found {
		return value.(int), found
	}
	return 0, false
}

// GetIntDefault implements [Config].
func (c *DefaultConfig) GetIntDefault(name string, defVal int) int {
	if value, found := c.GetInt(name); found {
		return value
	}
	return defVal
}

// GetString implements [Config].
func (c *DefaultConfig) GetString(name string) (value string, found bool) {
	if value, found := c.get(name); found {
		return value.(string), found
	}
	return "", false
}

// GetStringDefault implements [Config].
func (c *DefaultConfig) GetStringDefault(name string, defVal string) string {
	if value, found := c.GetString(name); found {
		return value
	}
	return defVal
}
