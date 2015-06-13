// Copyright 2015 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License. See the AUTHORS file
// for names of contributors.
//
// Author: Peter Mattis (peter@cockroachlabs.com)

package client_test

import (
	"testing"

	"github.com/cockroachdb/cockroach/proto"
)

func TestCreateTable(t *testing.T) {
	s, db := setup()
	defer s.Stop()

	if err := db.CreateTable(proto.TableSchema{
		Table: proto.Table{
			Name: "users",
		},
		Columns: []proto.Column{
			{Name: "id", Type: proto.Column_BYTES},
			{Name: "name", Type: proto.Column_BYTES},
			{Name: "title", Type: proto.Column_BYTES},
		},
		Indexes: []proto.TableSchema_IndexByName{
			{Index: proto.Index{Name: "primary", Unique: true},
				ColumnNames: []string{"id"}},
		},
	}); err != nil {
		t.Fatal(err)
	}
}
