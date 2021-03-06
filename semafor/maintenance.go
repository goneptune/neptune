// Copyright 2019 Alexey Yanchenko <mail@yanchenko.me>
//
// This file is part of the Neptune library.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package semafor

import (
	"net/http"
)

func Maintenance(w http.ResponseWriter, r *http.Request) {

	var data = HTMLData{}
	data.HeaderToHTML("Maintenance") //Title
	var bd = []string{"Site under Maintenance"}
	data.BodyToHTML(bd) //Content
	data.ShowPage(w, r, "maintenance")
}
