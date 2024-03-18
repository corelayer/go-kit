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
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Config Configuration

func NewApplication(c *cobra.Command, v Version) *Application {
	if v.runE != nil {
		c.AddCommand(v.Command())
	}

	if c.Version != "" {
		c.Version = ""
	}

	return &Application{
		Root:    c,
		Version: v,
	}
}

type Application struct {
	Root    *cobra.Command
	Version Version
}

func (a *Application) RegisterEnvironment(prefix string, keys []string) error {
	var (
		err     error
		eConfig EnvConfiguration
		eViper  *viper.Viper
	)

	eConfig = NewEnvConfiguration(prefix, keys)
	if !eConfig.HasKeys() {
		return fmt.Errorf("environment does not define keys")
	}

	if eViper, err = eConfig.GetViper(); err != nil {
		return err
	}
	return Config.SetEnvironment(eViper)
}

func (a *Application) RegisterCommands(c []Commander, f func(cmd *cobra.Command)) {
	for _, cmdr := range c {
		a.Root.AddCommand(cmdr.Initialize(f))
	}
}

// RegisterConfiguration TODO Add flags to cobra command for dynamic configuration name
// RegisterConfiguration TODO Add custom error type?
func (a *Application) RegisterConfiguration(name string, filename string, searchPaths []string) error {
	var (
		err    error
		fViper *viper.Viper
	)
	if Config.FileExists(name) {
		return fmt.Errorf("config name already exists")
	}
	c := NewFileConfiguration(filename, searchPaths)
	if fViper, err = c.GetViper(); err != nil {
		return err
	}
	return Config.SetFile(name, fViper)
}

func (a *Application) Run() error {
	return a.Root.Execute()
}

func init() {
	Config = NewConfiguration()
}
