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
	"crypto/x509"
	"text/template"
	"time"

	"bytes"

	"github.com/linkedin/Burrow/core/protocol"
)

// executeTemplate provides a common interface for notifier modules to call to process a text/template in the context
// of a protocol.ConsumerGroupStatus and create a message to use in a notification.
func executeTemplate(tmpl *template.Template, extras map[string]string, status *protocol.ConsumerGroupStatus, eventID string, startTime time.Time) (*bytes.Buffer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Helper functions for templates
var helperFunctionMap = template.FuncMap{
	"jsonencoder":     templateJSONEncoder,
	"topicsbystatus":  classifyTopicsByStatus,
	"partitioncounts": templateCountPartitions,
	"add":             templateAdd,
	"minus":           templateMinus,
	"multiply":        templateMultiply,
	"divide":          templateDivide,
	"maxlag":          maxLagHelper,
	"formattimestamp": formatTimestamp,
}

// Helper function for the templates to encode an object into a JSON string
func templateJSONEncoder(encodeMe interface{}) string { _ = "STUB: not implemented"; return "" }

// Helper - recategorize partitions as a map of lists
// map[string][]string => status short name -> list of topics
func classifyTopicsByStatus(partitions []*protocol.PartitionStatus) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

// Template Helper - Return a map of partition counts
// keys are warn, stop, stall, rewind, unknown
func templateCountPartitions(partitions []*protocol.PartitionStatus) map[string]int {
	_ = "STUB: not implemented"
	return nil
}

// Appends supplied certificates to trusted certificate chain
func buildRootCAs(extraCaFile string, noVerify bool) *x509.CertPool {
	_ = "STUB: not implemented"
	return nil
}

// Template Helper - do maths
func templateAdd(a, b int) int { _ = "STUB: not implemented"; return 0 }

func templateMinus(a, b int) int { _ = "STUB: not implemented"; return 0 }

func templateMultiply(a, b int) int { _ = "STUB: not implemented"; return 0 }

func templateDivide(a, b int) int { _ = "STUB: not implemented"; return 0 }

func maxLagHelper(a *protocol.PartitionStatus) uint64 { _ = "STUB: not implemented"; return 0 }

func formatTimestamp(timestamp int64, formatString string) string {
	_ = "STUB: not implemented"
	return ""
}
