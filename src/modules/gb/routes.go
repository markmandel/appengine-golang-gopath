//Copyright 2015 Google Inc.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

//Package gb implements the App Engine Module using GB
package gb

import (
	"lib/reverse"
	"net/http"

	//this is my external dependency (which is on the ./vendored GOPATH)
	uuid "github.com/nu7hatch/gouuid"
)

//init sets up the single route we have
func init() {
	//set up template
	initTemplate()

	//create http handler
	http.HandleFunc("/", handler)
}

//handler is a http handler that displays a UUID from
//our external package, and the result of `GiveMeASymbol`
//which comes from our custom function for this module.
//and uses the shared package `Reverse` to reverse the uuid.
func handler(w http.ResponseWriter, r *http.Request) {
	u, err := uuid.NewV4() //call our third party dependency

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	d := data{
		UUID:         u,
		Symbol:       GiveMeASymbol(),             //call our module level code
		ReversedUUID: reverse.Reverse(u.String()), //call our shared library package
	}

	//Execute our template and display to the screen
	err = display.Execute(w, d)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
