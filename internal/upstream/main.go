// Copyright 2020 Envoyproxy Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var (
		upstreamPort uint
		hello        string
	)

	flag.UintVar(&upstreamPort, "upstream", 18080, "Upstream HTTP/1.1 port")
	flag.StringVar(&hello, "message", "Default message", "Message to send in response")
	flag.Parse()

	log.Printf("upstream listening HTTP/1.1 on %d\n", upstreamPort)
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := w.Write([]byte(hello)); err != nil {
			log.Println(err)
		}
	})
	// Ignore: G114: Use of net/http serve function that has no support for setting timeouts
	// nolint:gosec
	if err := http.ListenAndServe(fmt.Sprintf(":%d", upstreamPort), nil); err != nil {
		log.Println(err)
	}
}
