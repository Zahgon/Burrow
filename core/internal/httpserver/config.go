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

func (hc *Coordinator) configMain(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_ = "STUB: not implemented"
	// Build JSON structs for config
	return
}

func (hc *Coordinator) writeModuleListResponse(w http.ResponseWriter, r *http.Request, coordinator string, modules []string) {
	_ = "STUB: not implemented"
	return
}

func (hc *Coordinator) configStorageList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_ = "STUB: not implemented"
	return
}

func (hc *Coordinator) configConsumerList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_ = "STUB: not implemented"
	return
}

func (hc *Coordinator) configClusterList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_ = "STUB: not implemented"
	return
}

func (hc *Coordinator) configEvaluatorList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_ = "STUB: not implemented"
	return
}

func (hc *Coordinator) configNotifierList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	_ = "STUB: not implemented"
	return
}

func (hc *Coordinator) configStorageDetail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_ = "STUB: not implemented"
	return
}

func (hc *Coordinator) configConsumerDetail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_ = "STUB: not implemented"
	return
}

func (hc *Coordinator) configEvaluatorDetail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_ = "STUB: not implemented"
	return
}

func (hc *Coordinator) configNotifierHTTP(w http.ResponseWriter, r *http.Request, configRoot string) {
	_ = "STUB: not implemented"
	return
}

func (hc *Coordinator) configNotifierSlack(w http.ResponseWriter, r *http.Request, configRoot string) {
	_ = "STUB: not implemented"
	return
}

func (hc *Coordinator) configNotifierEmail(w http.ResponseWriter, r *http.Request, configRoot string) {
	_ = "STUB: not implemented"
	return
}

func (hc *Coordinator) configNotifierNull(w http.ResponseWriter, r *http.Request, configRoot string) {
	_ = "STUB: not implemented"
	return
}

func (hc *Coordinator) configNotifierDetail(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	_ = "STUB: not implemented"
	return
}

// Return the right profile structure
