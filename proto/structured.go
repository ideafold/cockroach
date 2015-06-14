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

package proto

import (
	"fmt"
	"strings"
)

func validateName(name, typ string) error {
	if len(name) == 0 {
		return fmt.Errorf("empty %s name", typ)
	}
	// TODO(pmattis): Do we want to be more restrictive than this?
	if strings.Index(name, "/") != -1 {
		return fmt.Errorf("\"%s\" may not contain \"/\"", name)
	}
	return nil
}

// ValidateTableDesc validates that the table descriptor is well formed. Checks
// include validating the table, column and index names, verifying that column
// names and index names are unique and verifying that column IDs and index IDs
// are consistent.
func ValidateTableDesc(desc TableDescriptor) error {
	if err := validateName(desc.Name, "table"); err != nil {
		return err
	}

	if len(desc.Columns) == 0 {
		return fmt.Errorf("table must contain at least 1 column")
	}

	columnsByName := map[string]struct{}{}
	columnsByID := map[uint32]struct{}{}
	for _, column := range desc.Columns {
		if err := validateName(column.Name, "column"); err != nil {
			return err
		}

		if _, ok := columnsByName[column.Name]; ok {
			return fmt.Errorf("duplicate column name: \"%s\"", column.Name)
		}
		columnsByName[column.Name] = struct{}{}

		if _, ok := columnsByID[column.ID]; ok {
			return fmt.Errorf("column \"%s\" duplicate ID: %d", column.Name, column.ID)
		}
		columnsByID[column.ID] = struct{}{}

		if column.ID >= desc.NextColumnID {
			return fmt.Errorf("column \"%s\" invalid ID (%d) > next column ID (%d)",
				column.Name, column.ID, desc.NextColumnID)
		}
	}

	// TODO(pmattis): Check that the indexes are unique. That is, no 2 indexes
	// should contain identical sets of columns.

	if len(desc.Indexes) == 0 {
		return fmt.Errorf("table must contain at least 1 index")
	}

	indexesByName := map[string]struct{}{}
	indexesByID := map[uint32]struct{}{}
	for _, index := range desc.Indexes {
		if err := validateName(index.Name, "index"); err != nil {
			return err
		}

		if _, ok := indexesByName[index.Name]; ok {
			return fmt.Errorf("duplicate index name: \"%s\"", index.Name)
		}
		indexesByName[index.Name] = struct{}{}

		if _, ok := indexesByID[index.ID]; ok {
			return fmt.Errorf("index \"%s\" duplicate ID: %d", index.Name, index.ID)
		}
		indexesByID[index.ID] = struct{}{}

		if index.ID >= desc.NextIndexID {
			return fmt.Errorf("index \"%s\" invalid index ID (%d) > next index ID (%d)",
				index.Name, index.ID, desc.NextIndexID)
		}

		if len(index.ColumnIDs) == 0 {
			return fmt.Errorf("index \"%s\" must contain at least 1 column", index.Name)
		}

		for _, columnID := range index.ColumnIDs {
			if _, ok := columnsByID[columnID]; !ok {
				return fmt.Errorf("index \"%s\" contains unknown column ID %d", index.Name, columnID)
			}
		}
	}
	return nil
}

// TableDescFromSchema initializes a TableDescriptor from a TableSchema. The
// TableSchema is expected to be valid. An invalid table schema will result in
// an invalid table descriptor. Call validateTableDesc on the resulting
// descriptor to check for validity.
func TableDescFromSchema(schema TableSchema) TableDescriptor {
	desc := TableDescriptor{
		Table: schema.Table,
	}
	desc.Name = strings.ToLower(desc.Name)

	columnsByName := map[string]uint32{}
	for _, column := range schema.Columns {
		columnDesc := ColumnDescriptor{
			ID:     desc.NextColumnID,
			Column: column,
		}
		columnDesc.Name = strings.ToLower(columnDesc.Name)
		columnsByName[columnDesc.Name] = columnDesc.ID

		desc.Columns = append(desc.Columns, columnDesc)
		desc.NextColumnID++
	}

	for _, index := range schema.Indexes {
		indexDesc := IndexDescriptor{
			ID:    desc.NextIndexID,
			Index: index.Index,
		}
		indexDesc.Name = strings.ToLower(indexDesc.Name)

		for _, columnName := range index.ColumnNames {
			indexDesc.ColumnIDs = append(indexDesc.ColumnIDs, columnsByName[columnName])
		}

		desc.Indexes = append(desc.Indexes, indexDesc)
		desc.NextIndexID++
	}

	return desc
}

// TableSchemaFromDesc initializes a TableSchema from a TableDescriptor. The
// TableDescriptor is expected to be valid. An invalid table descriptor will
// result in an invalid table schema.
func TableSchemaFromDesc(desc TableDescriptor) TableSchema {
	schema := TableSchema{
		Table: desc.Table,
	}

	columnsByID := map[uint32]string{}
	for _, column := range desc.Columns {
		schema.Columns = append(schema.Columns, column.Column)
		columnsByID[column.ID] = column.Name
	}

	for _, index := range desc.Indexes {
		i := TableSchema_IndexByName{
			Index: index.Index,
		}
		for _, columnID := range index.ColumnIDs {
			i.ColumnNames = append(i.ColumnNames, columnsByID[columnID])
		}
		schema.Indexes = append(schema.Indexes, i)
	}
	return schema
}
