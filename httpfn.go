// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hellohttp

import (
	"fmt"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/otakakot/research-cloud-functions-metadata/config"
	"github.com/otakakot/research-cloud-functions-metadata/env"
)

func init() {
	functions.HTTP("HelloHTTP", HelloHTTP)
}

// HelloHTTP is an HTTP Cloud Function.
func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	env.Init()

	proj := config.NewProject(env.Get())

	pid, err := proj.GetID(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	fmt.Fprintf(w, "Hello, %s!", pid)
}
