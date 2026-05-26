// Copyright 2017 LinkedIn Corp. Licensed under the Apache License, Version
// 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

package helpers

// ValidateIP returns true if the provided string can be parsed as an IP address (either IPv4 or IPv6).
func ValidateIP(ipaddr string) bool { _ = "STUB: not implemented"; return false }

// ValidateHostname returns true if the provided string can be parsed as a hostname. In general this means:
//
// * One or more segments delimited by a '.'
// * Each segment can be no more than 63 characters long
// * Valid characters in a segment are letters, numbers, and dashes
// * Segments may not start or end with a dash
// * The exception is IPv6 addresses, which are also permitted.
// * An underscore is allowed to support Docker Swarm service names.
func ValidateHostname(hostname string) bool { _ = "STUB: not implemented"; return false }

// Try Docker Swarm service name

// Try as an IP address

// ValidateZookeeperPath returns true if the provided string can be parsed as a Zookeeper node path. This means that it
// starts with a forward slash, and contains one or more segments that are separated by slashes (but does not end with
// a slash).
func ValidateZookeeperPath(path string) bool { _ = "STUB: not implemented"; return false }

// Root node is OK

// ValidateTopic returns true if the provided string is a valid topic name, which may only contain letters, numbers,
// underscores, dashes, and periods.
func ValidateTopic(topic string) bool { _ = "STUB: not implemented"; return false }

// ValidateFilename returns true if the provided string is a sane-looking filename (not just a valid filename, which
// could be almost anything). Right now, this is defined to be the same thing as ValidateTopic.
func ValidateFilename(filename string) bool { _ = "STUB: not implemented"; return false }

// ValidateEmail returns true if the provided string is an email address. This is a very simplistic validator - the
// string must be of the form (something)@(something).(something)
func ValidateEmail(email string) bool { _ = "STUB: not implemented"; return false }

// ValidateURL returns true if the provided string can be parsed as a URL. We use the net/url Parse func for this.
func ValidateURL(rawURL string) bool { _ = "STUB: not implemented"; return false }

// ValidateHostList returns true if the provided slice of strings can all be parsed by ValidateHostPort
func ValidateHostList(hosts []string) bool { _ = "STUB: not implemented"; return false }

// ValidateHostPort returns true if the provided string is of the form "hostname:port", where hostname is a valid
// hostname or IP address (as parsed by ValidateIP or ValidateHostname), and port is a valid integer.
func ValidateHostPort(host string, allowBlankHost bool) bool {
	_ = "STUB: not implemented"
	// Must be hostname:port, ipv4:port, or [ipv6]:port. Optionally allow blank hostname
	return false
}

// Validate the port is a numeric (yeah, strings are valid in some places, but we don't support it)

// Listeners can have blank hostnames, so we'll skip validation if that's what we're looking for

// Only IPv6 can contain :

// If all the parts of the hostname are numbers, validate as IP. Otherwise, it's a hostname
