/*
 * Copyright 2024 CoreLayer BV
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */

package application

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

var Config *Configuration

func RegisterEnvironment(prefix string, keys []string) error {
	var (
		err     error
		eConfig EnvConfiguration
		eViper  *viper.Viper
	)

	if eConfig, err = NewEnvConfiguration(prefix, keys); err != nil {
		return err
	}

	if eViper, err = eConfig.GetViper(); err != nil {
		return err
	}
	return Config.SetEnvironment(eViper)
}

func RegisterConfiguration(name string, filename string, searchPaths []string) error {
	var (
		err    error
		fViper *viper.Viper
	)
	if Config.FileExists(name) {
		return ErrFileConfigurationExists
	}
	c := NewFileConfiguration(filename, searchPaths)
	if fViper, err = c.GetViper(); err != nil {
		return err
	}
	return Config.SetFile(name, fViper)
}

type ConfigFileFlags struct {
	Name  string
	Paths []string
}

type ConfigFileParams struct {
	File  StringVar
	Paths StringSliceVar
}

func NewConfiguration() *Configuration {
	return &Configuration{
		files: make(map[string]*viper.Viper),
		mux:   sync.Mutex{},
	}
}

type Configuration struct {
	env   *viper.Viper
	files map[string]*viper.Viper
	mux   sync.Mutex
}

func (c *Configuration) GetEnv() (*viper.Viper, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	if c.env == nil {
		return nil, ErrEnvConfigurationNotLoaded
	}
	return c.env, nil
}

func (c *Configuration) GetFile(name string) (*viper.Viper, error) {
	if !c.FileExists(name) {
		return nil, ErrFileConfigurationNotFound
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.files[name], nil
}

func (c *Configuration) SetEnvironment(v *viper.Viper) error {
	if c.env != nil {
		return ErrEnvConfigurationAlreadyLoaded
	}
	c.mux.Lock()
	defer c.mux.Unlock()
	c.env = v
	return nil
}

func (c *Configuration) SetFile(name string, v *viper.Viper) error {
	if c.FileExists(name) {
		return ErrFileConfigurationExists
	}
	c.mux.Lock()
	defer c.mux.Unlock()

	c.files[name] = v
	return v.ReadInConfig()
}

func (c *Configuration) FileExists(name string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	_, found := c.files[name]
	return found
}

func NewEnvConfiguration(prefix string, keys []string) (EnvConfiguration, error) {
	if len(keys) == 0 {
		return EnvConfiguration{}, ErrEnvConfigurationIsEmpty
	}
	return EnvConfiguration{
		prefix: prefix,
		keys:   keys,
	}, nil
}

type EnvConfiguration struct {
	prefix string
	keys   []string
}

func (e EnvConfiguration) GetViper() (*viper.Viper, error) {
	var (
		err error
		v   *viper.Viper
	)

	if !e.HasKeys() {
		return nil, ErrEnvConfigurationIsEmpty
	}

	v = viper.New()
	v.SetEnvPrefix(e.prefix)

	for _, k := range e.keys {
		err = v.BindEnv(k)
		if err != nil {
			return nil, err
		}
	}
	return v, nil
}

func (e EnvConfiguration) HasKeys() bool {
	if len(e.keys) == 0 {
		return false
	}
	return true
}

func (e EnvConfiguration) Keys() []string {
	return e.keys
}

func (e EnvConfiguration) Prefix() string {
	return e.prefix
}

func NewFileConfiguration(file string, searchPaths []string) FileConfiguration {
	// Clean search paths
	paths := make([]string, 0)
	for _, p := range searchPaths {
		if !strings.Contains(p, "..") {
			paths = append(paths, filepath.Clean(p))
		}
	}

	// Split file into path and filename
	path, filename := filepath.Split(file)

	// Make sure to clean path, this also causes the path to either be "." or a full path
	if path != "" {
		path = filepath.Clean(path)
	}
	c := FileConfiguration{
		filename: filename,
		path:     path,
		paths:    paths,
	}

	return c
}

type FileConfiguration struct {
	filename string
	path     string
	paths    []string
}

func (c FileConfiguration) GetViper() (*viper.Viper, error) {
	v := viper.New()

	// If a full path is specified, set the config file to that path
	if c.path != "" {
		fullPath := filepath.Join(c.path, c.filename)
		v.SetConfigFile(fullPath)
	} else {
		configName, configType := c.getViperConfig()
		v.SetConfigName(configName)
		v.SetConfigType(configType)

		for _, path := range c.paths {
			v.AddConfigPath(path)
		}
	}
	return v, v.ReadInConfig()
}

func (c FileConfiguration) getViperConfig() (string, string) {
	var (
		configName string
		configType string
	)

	fileExtension := filepath.Ext(c.filename)
	if fileExtension == "" {
		configType = "yaml"
	} else {
		configType = strings.TrimPrefix(fileExtension, ".")
	}

	configName = strings.TrimSuffix(c.filename, fileExtension)

	return configName, configType
}
