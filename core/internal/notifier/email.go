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
	"crypto/tls"
	"net/smtp"
	"regexp"
	"text/template"
	"time"

	"go.uber.org/zap"
	"gopkg.in/gomail.v2"

	"github.com/linkedin/Burrow/core/protocol"
)

// EmailNotifier is a module which can be used to send notifications of consumer group status via email messages. One
// email is sent for each consumer group that matches the allowlist/denylist and the status threshold.
type EmailNotifier struct {
	// App is a pointer to the application context. This stores the channel to the storage subsystem
	App *protocol.ApplicationContext

	// Log is a logger that has been configured for this module to use. Normally, this means it has been set up with
	// fields that are appropriate to identify this coordinator
	Log *zap.Logger

	name           string
	groupAllowlist *regexp.Regexp
	groupDenylist  *regexp.Regexp
	extras         map[string]string
	templateOpen   *template.Template
	templateClose  *template.Template

	to   string
	from string

	smtpDialer   *gomail.Dialer
	sendMailFunc func(message *gomail.Message) error
}

// Configure validates the configuration of the email notifier. At minimum, there must be a valid server, port, from
// address, and to address. If any of these are missing or incorrect, this func will panic with an explanatory message.
// It is also possible to specify an auth-type of either "plain" or "crammd5", along with a username and password.
func (module *EmailNotifier) Configure(name, configRoot string) {
	_ = "STUB: not implemented"

	// Abstract the SendMail call so we can test
	return
}

// Set up dialer and extra TLS configuration

func buildEmailTLSConfig(extraCaFile string, noVerify bool, smtpHost string) *tls.Config {
	_ = "STUB: not implemented"
	return nil
}

// Builds authentication profile for smtp client
func (module *EmailNotifier) getSMTPAuth(configRoot string) smtp.Auth {
	_ = "STUB: not implemented"

	// Set up SMTP authentication
	return *new(smtp.Auth)
}

// Start is a no-op for the email notifier. It always returns no error
func (module *EmailNotifier) Start() error {
	_ = "STUB: not implemented"

	// Stop is a no-op for the email notifier. It always returns no error
	return nil
}

func (module *EmailNotifier) Stop() error {
	_ = "STUB: not implemented"

	// GetName returns the configured name of this module
	return nil
}

func (module *EmailNotifier) GetName() string {
	_ = "STUB: not implemented"

	// GetGroupAllowlist returns the compiled group allowlist (or nil, if there is not one)
	return ""
}

func (module *EmailNotifier) GetGroupAllowlist() *regexp.Regexp {
	_ = "STUB: not implemented"
	return nil
}

// GetGroupDenylist returns the compiled group denylist (or nil, if there is not one)
func (module *EmailNotifier) GetGroupDenylist() *regexp.Regexp {
	_ = "STUB: not implemented"
	return nil
}

// GetLogger returns the configured zap.Logger for this notifier
func (module *EmailNotifier) GetLogger() *zap.Logger {
	_ = "STUB: not implemented"

	// AcceptConsumerGroup has no additional function for the email notifier, and so always returns true
	return nil
}

func (module *EmailNotifier) AcceptConsumerGroup(status *protocol.ConsumerGroupStatus) bool {
	_ = "STUB: not implemented"

	// Notify sends a single email message, with the from and to set to the configured addresses for the notifier. The
	// status, eventID, and startTime are all passed to the template for compiling the message. If stateGood is true, the
	// "close" template is used. Otherwise, the "open" template is used.
	return false
}

func (module *EmailNotifier) Notify(status *protocol.ConsumerGroupStatus, eventID string, startTime time.Time, stateGood bool) {
	_ = "STUB: not implemented"
	return
}

// Put the from and to lines in without the template. Template should set the subject line, followed by a blank line

// Process template headers and send email

// sendEmail uses the gomail smtpDialer to send a constructed message. This function is mocked for testing purposes
func (module *EmailNotifier) sendEmail(m *gomail.Message) error {
	_ = "STUB: not implemented"
	return nil
}

// createMessage organizes all relevant email message content into a structure for easy use
func (module *EmailNotifier) createMessage(messageContent string) (*gomail.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Go doesn't support regex lookaheads yet

func getKeywordContent(header, subjectDelimiter string) string {
	_ = "STUB: not implemented"
	return ""
}
