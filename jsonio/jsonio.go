// Copyright 2021 by William Jones. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Collection of functions to work with json files.
package jsonio

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

// LoadJSon : Loads a slice of a struct data from JSon file.
func LoadJSon(filename string, js interface{}) error {

	if !fileExists(filename) {
		return errors.New("file does not exist")
	}

	fileptr, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileptr.Close()
	decoder := json.NewDecoder(fileptr)
	err = decoder.Decode(&js)
	if err != nil {
		return err
	}

	return nil
}

// SaveJSon : Saves a struct slice to a file. Creating it if needed.
func SaveJSon(filename string, js interface{}) error {
	file, _ := json.MarshalIndent(js, "", " ")
	err := ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

// fileExists checks if a file exists and is not a directory
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
