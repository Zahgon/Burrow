// Copyright 2017 LinkedIn Corp. Licensed under the Apache License, Version
// 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

package core

import (
	"os"

	"go.uber.org/zap"
)

// CheckAndCreatePidFile takes a single argument, which is the path to a PID file (a file that contains a single
// integer, which is the process ID of a running process). If this file exists, and if the PID is that of a running
// process, return false as that indicates another copy of this process is already running. Otherwise, create the
// file and write this process's PID to the file and return true. Any error doing this (such as not having permissions
// to write the file) will return false.
//
// This func should be called when Burrow starts to prevent multiple copies from running.
func CheckAndCreatePidFile(filename string) bool {
	_ = "STUB: not implemented"
	// Check if the PID file exists
	return false
}

// The file exists, so read it and check if the PID specified is running

// This could happen inside a docker
// container, e.g. the pid of Burrow could be
// equal to 1 each time the container is
// restarted.

// Try sending a signal to the process to see if it is still running

// The process exists, so we're going to assume it's an old Burrow and we shouldn't start

// Create a PID file, replacing any existing one (as we already checked it)

// RemovePidFile takes a single argument, which is the path to a PID file. That file is deleted. This func should be
// called when Burrow exits.
func RemovePidFile(filename string) { _ = "STUB: not implemented"; return }

// ConfigureLogger returns a configured zap.Logger which can be used by Burrow for all logging. It also returns a
// zap.AtomicLevel, which can be used to dynamically adjust the level of the logger. The configuration for the logger
// is read from viper, with the following defaults:
//
// logging.level = info
//
// If logging.filename (path to the log file) is provided, a rolling log file is set up using lumberjack. The
// configuration for that log file is read from viper, with the following defaults:
//
// logging.maxsize = 100
// logging.maxbackups = 10
// logging.maxage = 30
// logging.use-localtime = false
// logging.use-compression = false
func ConfigureLogger() (*zap.Logger, *zap.AtomicLevel) { _ = "STUB: not implemented"; return nil, nil }

// Set config defaults for logging

// Create an AtomicLevel that we can use elsewhere to dynamically change the logging level

// If a filename has been set, set up a rotating logger. Otherwise, use Stdout

// OpenOutLog takes a single argument, which is the path to a log file. This process's stdout and stderr are redirected
// to this log file. The os.File object is returned so that it can be managed.
func OpenOutLog(filename string) *os.File {
	_ = "STUB: not implemented"
	// Move existing out file to a dated file if it exists
	return nil
}

// Redirect stdout and stderr to out file
