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

const (
	ErrEnvConfigurationAlreadyLoadedMessage = " environment is already loaded"
	ErrEnvConfigurationIsEmptyMessage       = "environment configuration does not define keys"
	ErrEnvConfigurationNotLoadedMessage     = "environment not loaded"
	ErrFileConfigurationExistsMessage       = "configuration exists"
	ErrFileConfigurationNotFoundMessage     = "configuration not found"
)

var (
	ErrEnvConfigurationAlreadyLoaded = EnvConfigurationAlreadyLoadedError{message: ErrEnvConfigurationAlreadyLoadedMessage}
	ErrEnvConfigurationIsEmpty       = EnvConfigurationIsEmptyError{message: ErrEnvConfigurationIsEmptyMessage}
	ErrEnvConfigurationNotLoaded     = EnvConfigurationNotLoadedError{message: ErrEnvConfigurationNotLoadedMessage}
	ErrFileConfigurationExists       = FileConfigurationExistsError{message: ErrFileConfigurationExistsMessage}
	ErrFileConfigurationNotFound     = FileConfigurationNotFoundError{message: ErrFileConfigurationNotFoundMessage}
)

type EnvConfigurationAlreadyLoadedError struct {
	message string
}

func (e EnvConfigurationAlreadyLoadedError) Error() string {
	return e.message
}

type EnvConfigurationIsEmptyError struct {
	message string
}

func (e EnvConfigurationIsEmptyError) Error() string {
	return e.message
}

type EnvConfigurationNotLoadedError struct {
	message string
}

func (e EnvConfigurationNotLoadedError) Error() string {
	return e.message
}

type FileConfigurationExistsError struct {
	message string
}

func (e FileConfigurationExistsError) Error() string {
	return e.message
}

type FileConfigurationNotFoundError struct {
	message string
}

func (e FileConfigurationNotFoundError) Error() string {
	return e.message
}
