// Copyright 2017 LinkedIn Corp. Licensed under the Apache License, Version
// 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

package notifier

import (
	"net/http"
	"regexp"
	"text/template"
	"time"

	"go.uber.org/zap"

	"crypto/tls"

	"github.com/linkedin/Burrow/core/protocol"
)

// HTTPNotifier is a module which can be used to send notifications of consumer group status via outbound HTTP calls to
// another server. This is useful for informing another system, such as an alert system, when there is a problem. One
// HTTP call is made for each consumer group that matches the allowlist/denylist and the status threshold (though
// keepalive connections will be used if configured).
type HTTPNotifier struct {
	// App is a pointer to the application context. This stores the channel to the storage subsystem
	App *protocol.ApplicationContext

	// Log is a logger that has been configured for this module to use. Normally, this means it has been set up with
	// fields that are appropriate to identify this coordinator
	Log *zap.Logger

	name           string
	groupAllowlist *regexp.Regexp
	groupDenylist  *regexp.Regexp
	extras         map[string]string
	urlOpen        string
	urlClose       string
	methodOpen     string
	methodClose    string
	templateOpen   *template.Template
	templateClose  *template.Template
	sendClose      bool

	httpClient *http.Client
}

// Configure validates the configuration of the http notifier. At minimum, there must be a url-open specified, and if
// send-close is set to true there must also be a url-close. If these are missing or incorrect, this func will panic
// with an explanatory message. It is also possible to configure a specific method (such as POST or DELETE) to be used
// with these URLs, as well as a timeout and keepalive for the HTTP smtpClient.
func (module *HTTPNotifier) Configure(name, configRoot string) {
	_ = "STUB: not implemented"

	// Validate and set defaults for profile configs
	return
}

// Set defaults for module-specific configs if needed

func buildHTTPTLSConfig(extraCaFile string, noVerify bool) *tls.Config {
	_ = "STUB: not implemented"
	return nil
}

// Start is a no-op for the http notifier. It always returns no error
func (module *HTTPNotifier) Start() error {
	_ = "STUB: not implemented"

	// Stop is a no-op for the http notifier. It always returns no error
	return nil
}

func (module *HTTPNotifier) Stop() error {
	_ = "STUB: not implemented"

	// GetName returns the configured name of this module
	return nil
}

func (module *HTTPNotifier) GetName() string {
	_ = "STUB: not implemented"

	// GetGroupAllowlist returns the compiled group allowlist (or nil, if there is not one)
	return ""
}

func (module *HTTPNotifier) GetGroupAllowlist() *regexp.Regexp {
	_ = "STUB: not implemented"
	return nil
}

// GetGroupDenylist returns the compiled group denylist (or nil, if there is not one)
func (module *HTTPNotifier) GetGroupDenylist() *regexp.Regexp {
	_ = "STUB: not implemented"
	return nil
}

// GetLogger returns the configured zap.Logger for this notifier
func (module *HTTPNotifier) GetLogger() *zap.Logger {
	_ = "STUB: not implemented"

	// AcceptConsumerGroup has no additional function for the http notifier, and so always returns true
	return nil
}

func (module *HTTPNotifier) AcceptConsumerGroup(status *protocol.ConsumerGroupStatus) bool {
	_ = "STUB: not implemented"

	// Notify makes a single outbound HTTP request. The status, eventID, and startTime are all passed to the template for
	// compiling the request body. If stateGood is true, the "close" template and URL are used. Otherwise, the "open"
	// template and URL are used.
	return false
}

func (module *HTTPNotifier) Notify(status *protocol.ConsumerGroupStatus, eventID string, startTime time.Time, stateGood bool) {
	_ = "STUB: not implemented"
	return
}

// Send request to HTTP endpoint

// Add basic auth using the provided username and password
