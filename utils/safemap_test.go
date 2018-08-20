// Copyright 2014 izigo Author. All Rights Reserved.
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

package utils

import "testing"

var safeMap *IZIMap

func TestNewIZIMap(t *testing.T) {
	safeMap = NewIZIMap()
	if safeMap == nil {
		t.Fatal("expected to return non-nil IZIMap", "got", safeMap)
	}
}

func TestSet(t *testing.T) {
	if ok := safeMap.Set("diepdt", 1); !ok {
		t.Error("expected", true, "got", false)
	}
}

func TestReSet(t *testing.T) {
	safeMap := NewIZIMap()
	if ok := safeMap.Set("diepdt", 1); !ok {
		t.Error("expected", true, "got", false)
	}
	// set diff value
	if ok := safeMap.Set("diepdt", -1); !ok {
		t.Error("expected", true, "got", false)
	}

	// set same value
	if ok := safeMap.Set("diepdt", -1); ok {
		t.Error("expected", false, "got", true)
	}
}

func TestCheck(t *testing.T) {
	if exists := safeMap.Check("diepdt"); !exists {
		t.Error("expected", true, "got", false)
	}
}

func TestGet(t *testing.T) {
	if val := safeMap.Get("diepdt"); val.(int) != 1 {
		t.Error("expected value", 1, "got", val)
	}
}

func TestDelete(t *testing.T) {
	safeMap.Delete("diepdt")
	if exists := safeMap.Check("diepdt"); exists {
		t.Error("expected element to be deleted")
	}
}

func TestItems(t *testing.T) {
	safeMap := NewIZIMap()
	safeMap.Set("diepdt", "hello")
	for k, v := range safeMap.Items() {
		key := k.(string)
		value := v.(string)
		if key != "diepdt" {
			t.Error("expected the key should be diepdt")
		}
		if value != "hello" {
			t.Error("expected the value should be hello")
		}
	}
}

func TestCount(t *testing.T) {
	if count := safeMap.Count(); count != 0 {
		t.Error("expected count to be", 0, "got", count)
	}
}
