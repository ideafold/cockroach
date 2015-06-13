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

package client

import (
	"encoding/json"

	"github.com/cockroachdb/cockroach/proto"
)

func prettyJSON(v interface{}) string {
	pretty, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(pretty)
}

// CreateTable ...
func (db *DB) CreateTable(schema proto.TableSchema) error {
	return nil
}

// DescribeTable ...
func (db *DB) DescribeTable(name string) (proto.TableSchema, error) {
	// TODO(pmattis): Read table descriptor. Convert descriptor to table schema.
	s := proto.TableSchema{}
	return s, nil
}

// RenameTable ...
func (db *DB) RenameTable(oldName, newName string) error {
	panic("TODO(pmattis): unimplemented")
}

// DeleteTable ...
func (db *DB) DeleteTable(oldName, newName string) error {
	panic("TODO(pmattis): unimplemented")
}

// ListTables ...
func (db *DB) ListTables() ([]string, error) {
	panic("TODO(pmattis): unimplemented")
	// TODO(pmattis): Scan namespace keys, extract table names.
}
