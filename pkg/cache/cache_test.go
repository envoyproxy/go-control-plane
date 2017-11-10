// Copyright 2017 Envoyproxy Authors
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

package cache

import "testing"

func TestWatchCancel(t *testing.T) {
	called := 0
	w := Watch{stop: func() { called++ }}
	w.Cancel()
	w.Cancel()
	if called != 1 {
		t.Errorf("got count %d; stop function should be called once by Cancel()", called)
	}
}
