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

	"github.com/spf13/viper"
)

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
	var (
		err error
		v   *viper.Viper
	)
	v = viper.New()

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
	err = v.ReadInConfig()
	return v, err
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
