// Copyright 2021 by William Jones. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Routines and data struct for getting various system and cpu information.
package systemcpu

import (
	"os"
	"runtime"
	"sync"
)

// SysRun : Strucy to hold all the system variables
// Cant use a new function because of the mutexs
type SysRun struct {
	MU           sync.Mutex // Lock struct before any access
	Goroutines   int        // Goroutines : Current count of running go routines.
	CPUCount     int        // CPUCount : Number of CPUs
	Architecture string     // Architecture : intel, amd64, arm ...
	OS           string     // OS : Linux, Windows ...
	GoVersion    string     // GoVersion : 1.16.2
	PID          int        // PID : pid of the running program. Used to make unique name.
}

// Update : Update all the variables in the struct.
func (s *SysRun) Update() {
	s.GetCPU()
	s.GetOS()
	s.GetGo()
}

// GetCPU : Get the number of cpu(s) or cores and the architecture and update the struct.
func (s *SysRun) GetCPU() {
	s.MU.Lock() // lock
	s.CPUCount = runtime.NumCPU()
	s.Architecture = runtime.GOARCH
	s.MU.Unlock()
}

// GetOS : Get the OS and PID of the program and update the struct.
func (s *SysRun) GetOS() {
	s.MU.Lock() // lock
	s.OS = runtime.GOOS
	s.PID = os.Getpid()
	s.MU.Unlock()
}

// GetGo : Get the go version and the number of go routines and update the struct.
func (s *SysRun) GetGo() {
	s.MU.Lock()
	s.GoVersion = runtime.Version()
	s.Goroutines = runtime.NumGoroutine()
	s.MU.Unlock()
}
