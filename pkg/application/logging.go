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
	"log/slog"
	"strings"
)

const (
	TextLogFormat LogFormat = iota
	JsonLogFormat
)

var logFormat = map[string]LogFormat{
	"text": TextLogFormat,
	"json": JsonLogFormat,
}

type LogFlags struct {
	Level  string
	Format string
}

type LogFormat int

func (f LogFormat) String() string {
	return [...]string{"none", "text", "json"}[f]
}

type LogParams struct {
	Level  StringVar
	Format StringVar
	Output string
}

func ParseLogFormat(format string) (LogFormat, bool) {
	f, ok := logFormat[strings.ToLower(format)]
	return f, ok
}

func ParseLevel(level string) (slog.Level, bool) {
	switch level {
	case "error":
		return slog.LevelError, true
	case "warn":
		return slog.LevelWarn, true
	case "info":
		return slog.LevelInfo, true
	case "debug":
		return slog.LevelDebug, true
	default:
		return slog.LevelError, false
	}
}
