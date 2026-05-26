// Copyright 2017 LinkedIn Corp. Licensed under the Apache License, Version
// 2.0 (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.

package httpserver

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (hc *Coordinator) handleClusterList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_ = "STUB: not implemented"
	// Fetch cluster list from the storage module
	return
}

func getTLSProfile(name string) *httpResponseTLSProfile { _ = "STUB: not implemented"; return nil }

func getSASLProfile(name string) *httpResponseSASLProfile { _ = "STUB: not implemented"; return nil }

func getClientProfile(name string) httpResponseClientProfile {
	_ = "STUB: not implemented"
	return *new(httpResponseClientProfile)
}

func (hc *Coordinator) handleClusterDetail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_ = "STUB: not implemented"
	// Get cluster config
	return
}

func (hc *Coordinator) handleTopicList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_ = "STUB: not implemented"
	// Fetch topic list from the storage module
	return
}

func (hc *Coordinator) handleTopicDetail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_ = "STUB: not implemented"
	// Fetch topic offsets from the storage module
	return
}

func (hc *Coordinator) handleTopicConsumerList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_ = "STUB: not implemented"
	// Fetch topic offsets from the storage module
	return
}

func (hc *Coordinator) handleConsumerList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_ = "STUB: not implemented"
	// Fetch consumer list from the storage module
	return
}

func (hc *Coordinator) handleConsumerDetail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_ = "STUB: not implemented"
	// Fetch consumer data from the storage module
	return
}

func (hc *Coordinator) handleConsumerStatus(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_ = "STUB: not implemented"
	// Fetch consumer data from the storage module
	return
}

func (hc *Coordinator) handleConsumerStatusComplete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_ = "STUB: not implemented"
	// Fetch consumer data from the storage module
	return
}

func (hc *Coordinator) handleConsumerDelete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_ = "STUB: not implemented"
	// Delete consumer from the storage module
	return
}
