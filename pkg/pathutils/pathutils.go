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

package pathutils

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func CreateDirectory(path string, permissions os.FileMode) error {
	var (
		err          error
		extendedPath string
		src          os.FileInfo
	)

	if extendedPath, err = GetExpandedPath(path); err != nil {
		return err
	}

	src, err = os.Stat(extendedPath)
	if os.IsNotExist(err) {
		return os.MkdirAll(extendedPath, permissions)
	} else if src.Mode().IsRegular() {
		return os.ErrExist
	}
	return nil
}

// GetExpandedPath sourced from https://github.com/spf13/viper/blob/master/util.go#L108
func GetExpandedPath(path string) (string, error) {
	if path == "$HOME" || strings.HasPrefix(path, "$HOME"+string(os.PathSeparator)) {
		path = userHomeDir() + path[5:]
	}

	path = os.ExpandEnv(path)

	if filepath.IsAbs(path) {
		return filepath.Clean(path), nil
	}

	return filepath.Abs(path)
}

func GetFilenames(path string, extensions []string) ([]string, error) {
	return walkPath(path, extensions)
}

func isValidFile(e os.DirEntry, extensions []string) bool {
	ext := filepath.Ext(e.Name())
	for _, v := range extensions {
		if v == ext {
			return true
		}
	}
	return false
}

// userHomeDir sourced from https://github.com/spf13/viper/blob/master/util.go#L140
func userHomeDir() string {
	switch runtime.GOOS {
	case "windows":
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	default:
		return os.Getenv("HOME")
	}
}

func walkPath(path string, extensions []string) ([]string, error) {
	var (
		err          error
		extendedPath string
		dirEntries   []os.DirEntry
		files        []string
	)

	extendedPath, err = GetExpandedPath(path)
	if err != nil {
		return nil, err
	}

	dirEntries, err = os.ReadDir(extendedPath)
	if err != nil {
		return files, err
	}

	for _, file := range dirEntries {
		if !file.IsDir() {
			if isValidFile(file, extensions) {
				files = append(files, extendedPath+"/"+file.Name())
			}
		} else {
			var subDirFiles []string
			subDirFiles, err = walkPath(extendedPath+"/"+file.Name(), extensions)
			if err != nil {
				return files, err
			}
			files = append(files, subDirFiles...)
		}
	}
	return files, nil
}
