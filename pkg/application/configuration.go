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
	"sync"

	"github.com/spf13/viper"
)

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
