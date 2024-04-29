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
	"github.com/spf13/cobra"
)

func NewApplication(c *cobra.Command, v Version) *Application {
	if v.RunE != nil {
		c.AddCommand(v.Command())
		c.Version = ""
	} else {
		c.Version = v.String()
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

func (a *Application) RegisterCommands(c []Commander, f func(cmd *cobra.Command)) {
	for _, cmdr := range c {
		a.Root.AddCommand(cmdr.Initialize(f))
	}
}

func (a *Application) Run() error {
	return a.Root.Execute()
}

func init() {
	Config = NewConfiguration()
}
