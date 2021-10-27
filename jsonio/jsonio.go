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
	"strings"
)

// LoadJSonFromString : Load a stuct or a slice of structs from string s
func LoadJSonFromString(s string, js interface{}) error {
	if len(s) <= 0 {
		return errors.New("string is empty")
	}

	r := strings.NewReader(s)
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&js)
	if err != nil {
		return (err)
	}

	return nil
}

// LoadJSon : Loads a struct or a slice of a struct data from JSon file.
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

// SaveJSonToString : Saves struct or a slice to a string
func SaveJSonToString(js interface{}) (string, error) {
	rbytes, err := json.MarshalIndent(js, "", " ")
	if err != nil {
		return "", err
	}
	return string(rbytes), nil
}

// SaveJSon : Saves a struct or a struct slice to a file. Creating it if needed.
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
