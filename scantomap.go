// Copyright 2013 sigu-399 ( https://github.com/sigu-399 )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// author  			sigu-399
// author-github 	https://github.com/sigu-399
// author-mail		sigu.399@gmail.com
// 
// repository-name	gosqlmf
// repository-desc  Fetches results from SQL rows into a map - Go language
// 
// description		Main and only file.
// 
// created      	21-02-2013

package gosqlmf

import (
	"database/sql"
)

// QueryOne
//
// Database query function on top of FetchOne
//
func QueryOne(db *sql.DB, sql string) (bool, map[string]interface{}, error) {

	rows, err := db.Query(sql)
	if err != nil {
		return false, nil, err
	}

	ok, rowMap, err := FetchOne(rows)
	if err != nil {
		return false, nil, err
	}

	return ok, rowMap, nil
}

// QueryAll
//
// Database query function on top of FetchAll
//
func QueryAll(db *sql.DB, sql string) ([]map[string]interface{}, error) {

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	rowsMap, err := FetchAll(rows)
	if err != nil {
		return nil, err
	}

	return rowsMap, nil
}

// FetchOne
//
// Fetches one row into a map[string]interface{}, each key is the column name
//
// @return bool - next ( = rows.Next() )
// @return map[string]interface{}
// @return error - error or nil if everything is ok
//
func FetchOne(rows *sql.Rows) (bool, map[string]interface{}, error) {

	if !rows.Next() {
		return false, nil, nil
	}

	columns, err := rows.Columns()

	if err != nil {
		return true, nil, err
	}

	columnCount := len(columns)

	assoc, err := scanOne(rows, columnCount, columns)

	if err != nil {
		return true, nil, err
	}

	return true, assoc, nil
}

// FetchAll
//
// Fetches all rows into a []map[string]interface{}, each key is the column name
//
// @return []map[string]interface{}
// @return error - error or nil if everything is ok
//
func FetchAll(rows *sql.Rows) ([]map[string]interface{}, error) {

	var columns []string
	var columnCount int
	var err error

	assocArray := make([]map[string]interface{}, 0)

	processedRows := 0

	for rows.Next() {

		// Read columns on first row only
		if processedRows == 0 {

			columns, err = rows.Columns()

			if err != nil {
				return nil, err
			}

			columnCount = len(columns)
		}

		assoc, err := scanOne(rows, columnCount, columns)

		if err != nil {
			return nil, err
		}

		assocArray = append(assocArray, assoc)
		processedRows++
	}

	/// No rows - returns nil instead of empty []map
	if processedRows == 0 {
		return nil, nil
	}

	return assocArray, nil
}

// scanOne
//
// Private utility function that scans a row into a map
//
func scanOne(rows *sql.Rows, columnCount int, columns []string) (map[string]interface{}, error) {

	scanFrom := make([]interface{}, columnCount)
	scanTo := make([]interface{}, columnCount)

	for i, _ := range scanFrom {
		scanFrom[i] = &scanTo[i]
	}

	err := rows.Scan(scanFrom...)

	if err != nil {
		return nil, err
	}

	assoc := make(map[string]interface{})

	// Build the associative map from values and column names
	for i, _ := range scanTo {
		assoc[columns[i]] = scanTo[i]
	}

	return assoc, nil

}
