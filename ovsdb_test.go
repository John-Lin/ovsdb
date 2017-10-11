// Modifications copyright (C) 2017 Che Wei, Lin
// Copyright 2014 Cisco Systems Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package ovsdb

import (
	"fmt"
	"testing"
	"time"
)

var ovsDriver *OvsDriver

func TestCreateDeleteBridge(t *testing.T) {

	ovsDriver = NewOvsDriverWithUnix("ovsbr10")
	time.After(100 * time.Millisecond)

	// Test delete
	err := ovsDriver.DeleteBridge("ovsbr11")
	if err != nil {
		fmt.Printf("Error deleting the bridge. Err: %v", err)
		t.Errorf("Failed to delete a bridge")
	}
}

func TestCreateDeleteMultipleBridge(t *testing.T) {
	bridgeSize := 5
	// Test create
	ovsDrivers := make([]OvsDriver, 10)
	for i := 0; i < bridgeSize; i++ {
		brName := "ovsbr2" + fmt.Sprintf("%d", i)
		fmt.Println(brName)
		ovsDrivers[i] = *(NewOvsDriverWithUnix(brName))
	}

	time.After(100 * time.Millisecond)

	// Test delete
	for i := 0; i < bridgeSize; i++ {
		brName := "ovsbr2" + fmt.Sprintf("%d", i)
		fmt.Println(brName)
		err := (ovsDrivers[i]).DeleteBridge(brName)
		if err != nil {
			fmt.Printf("Error deleting the bridge. Err: %v", err)
			t.Errorf("Failed to delete a bridge")
		}
	}
}

/*
func TestCreateDeletePort(t *testing.T) {
	// Create a port
	err := ovsDriver.CreatePort("port12", "internal", 11)
	if err != nil {
		fmt.Printf("Error creating the port. Err: %v", err)
		t.Errorf("Failed to create a port")
	}

	// HACK: wait a little so that interface is visible
	time.Sleep(time.Second * 1)

	ovsDriver.PrintCache()

	if ovsDriver.IsPortNamePresent("port12") {
		fmt.Printf("Interface exists\n")
	} else {
		fmt.Printf("Interface does not exist\n")
	}

	// Delete port
	err = ovsDriver.DeletePort("port12")
	if err != nil {
		fmt.Printf("Error Deleting the port. Err: %v", err)
		t.Errorf("Failed to delete a port")
	}
}

func TestCreateVtep(t *testing.T) {
	// Create a port
	err := ovsDriver.CreateVtep("vtep1", "10.10.10.10")
	if err != nil {
		fmt.Printf("Error creating the VTEP. Err: %v", err)
		t.Errorf("Failed to create a port")
	}

	time.After(100 * time.Millisecond)

	isPresent, vtepName := ovsDriver.IsVtepPresent("10.10.10.10")
	if (!isPresent) || (vtepName != "vtep1") {
		t.Errorf("Unable to find the VTEP. present: %v, name: %s", isPresent, vtepName)
	}
}

func TestAddController(t *testing.T) {
	// Create a port
	err := ovsDriver.AddController("127.0.0.1", 6666)
	if err != nil {
		fmt.Printf("Error adding controller. Err: %v", err)
		t.Errorf("Failed to add controller")
	}
}
*/
