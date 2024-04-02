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

import "github.com/spf13/cobra"

func AddConfigFileFlag(cmd *cobra.Command, f *ConfigFileFlags, p ConfigFileParams) {
	cmd.PersistentFlags().StringVarP(&f.Name, p.File.Name, p.File.Shorthand, p.File.Value, p.File.Usage)
}

func AddConfigFilePathsFlag(cmd *cobra.Command, f *ConfigFileFlags, p ConfigFileParams) {
	cmd.PersistentFlags().StringSliceVarP(&f.Paths, p.Paths.Name, p.Paths.Shorthand, p.Paths.Value, p.Paths.Usage)
}

func AddLogLevelFlag(cmd *cobra.Command, f *LogFlags, p LogParams) {
	cmd.PersistentFlags().StringVarP(&f.Level, p.Level.Name, p.Level.Shorthand, p.Level.Value, p.Level.Usage)
}

func AddLogFormatFlag(cmd *cobra.Command, f *LogFlags, p LogParams) {
	cmd.PersistentFlags().StringVarP(&f.Format, p.Format.Name, p.Format.Shorthand, p.Format.Value, p.Format.Usage)
}

func AddTuiInteractiveFlag(cmd *cobra.Command, f *TuiFlags, p TuiParams) {
	cmd.PersistentFlags().BoolVarP(&f.Interactive, p.Interactive.Name, p.Interactive.Shorthand, p.Interactive.Value, p.Interactive.Usage)
}
