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
	"github.com/spf13/viper"
)

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
