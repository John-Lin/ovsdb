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
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var ovsDriver *OvsDriver

func TestCreateDeleteBridge(t *testing.T) {
	bridgeName := "br100"

	ovsDriver = NewOvsDriverWithUnix(bridgeName)
	time.Sleep(300 * time.Millisecond)

	exist := ovsDriver.IsBridgePresent(bridgeName)
	assert.True(t, exist)
	exist = ovsDriver.IsBridgePresent(bridgeName + "111")
	assert.False(t, exist)

	// Test delete
	err := ovsDriver.DeleteBridge(bridgeName)
	assert.NoError(t, err)
}

func TestCreateDeleteMultipleBridge(t *testing.T) {
	bridgeSize := 5
	// Test create
	ovsDrivers := make([]OvsDriver, 10)
	for i := 0; i < bridgeSize; i++ {
		brName := "ovsbr2" + fmt.Sprintf("%d", i)
		ovsDrivers[i] = *(NewOvsDriverWithUnix(brName))
	}

	time.Sleep(300 * time.Millisecond)

	// Test delete
	for i := 0; i < bridgeSize; i++ {
		brName := "ovsbr2" + fmt.Sprintf("%d", i)
		err := (ovsDrivers[i]).DeleteBridge(brName)
		assert.NoError(t, err)

	}
}
func TestCreateDeletePort(t *testing.T) {
	bridgeName := "ovsbr10"
	portName := "port10"
	// Create a Bridge
	ovsDriver = NewOvsDriverWithUnix(bridgeName)
	// Create a port
	err := ovsDriver.CreatePort(portName, "internal", 11)
	assert.NoError(t, err)

	// HACK: wait a little so that interface is visible
	time.Sleep(time.Second * 1)

	exist := ovsDriver.IsPortNamePresent(portName)
	assert.True(t, exist)
	// Delete port
	err = ovsDriver.DeletePort(portName)
	assert.NoError(t, err)

	err = ovsDriver.DeleteBridge(bridgeName)
	assert.NoError(t, err)
}

func TestCreateVtep(t *testing.T) {
	bridgeName := "ovsbr10"
	vethName := "vetp1"
	vethIP := "10.10.10.10"
	unknownIP := "1.2.3.4"
	ovsDriver = NewOvsDriverWithUnix(bridgeName)
	// Create a port
	err := ovsDriver.CreateVtep(vethName, vethIP)
	assert.NoError(t, err)

	time.Sleep(300 * time.Millisecond)

	isPresent, vtepName := ovsDriver.IsVtepPresent(vethIP)
	assert.True(t, isPresent)
	assert.Equal(t, vtepName, vethName)

	isPresent, vtepName = ovsDriver.IsVtepPresent(unknownIP)
	assert.False(t, isPresent)
	assert.Equal(t, vtepName, "")

	err = ovsDriver.DeleteVtep(vethName)
	assert.NoError(t, err)
	err = ovsDriver.DeleteBridge(bridgeName)
	assert.NoError(t, err)
}

func TestAddController(t *testing.T) {
	bridgeName := "ovsbr10"
	ovsDriver = NewOvsDriverWithUnix(bridgeName)
	// Create a port
	err := ovsDriver.AddController("127.0.0.1", 6666)
	assert.NoError(t, err)
	// HACK: wait a little so that interface is visible
	time.Sleep(time.Second * 1)
	exist := ovsDriver.IsControllerPresent("127.0.0.1", 6666)
	assert.True(t, exist)
	exist = ovsDriver.IsControllerPresent("127.0.0.1", 5555)
	assert.False(t, exist)
	err = ovsDriver.DeleteBridge(bridgeName)
	assert.NoError(t, err)
}
