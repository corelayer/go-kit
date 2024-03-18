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

package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/corelayer/go-kit/pkg/application"
	"github.com/corelayer/go-kit/pkg/timestamp"
)

var (
	VersionSemVer string
	VersionCommit string
	VersionDate   = timestamp.Now()
)

func PrintVersion(v *application.Version) error {
	fmt.Println("APP VERSION DETAILS")
	fmt.Println("Version:", v.SemVer)
	fmt.Println("Commit:", v.Commit)
	fmt.Println("Date:", v.Date)
	return nil
}

func main() {
	var err error
	root := &cobra.Command{
		Use:               "app",
		Short:             "App Short",
		Long:              "App Long",
		CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true, HiddenDefaultCmd: true},
	}

	version := application.Version{
		SemVer: VersionSemVer,
		Commit: VersionCommit,
		Date:   VersionDate,
	}

	version.SetRun(PrintVersion)

	app := application.NewApplication(root, version)
	if err = app.RegisterEnvironment("EXAMPLE", []string{"KEY"}); err != nil {
		panic(err)
	}

	if err = app.Run(); err != nil {
		fmt.Println(err)
	}
}
