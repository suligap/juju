// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"labix.org/v2/mgo/bson"
	"labix.org/v2/mgo/txn"
)

// NetworkInterface represents the state of a machine network
// interface.
type NetworkInterface struct {
	st  *State
	doc networkInterfaceDoc
}

// networkInterfaceDoc represents a network interface for a machine on
// a given network.
type networkInterfaceDoc struct {
	MACAddress string `bson:"_id"`
	// InterfaceName is the network interface name (e.g. "eth0").
	InterfaceName string
	NetworkName   string
	MachineId     string
}

func newNetworkInterface(st *State, doc *networkInterfaceDoc) *NetworkInterface {
	return &NetworkInterface{st, *doc}
}

// MACAddress returns the MAC address of the interface.
func (ni *NetworkInterface) MACAddress() string {
	return ni.doc.MACAddress
}

// InterfaceName returns the name of the interface.
func (ni *NetworkInterface) InterfaceName() string {
	return ni.doc.InterfaceName
}

// NetworkName returns the machine network name of the interface.
func (ni *NetworkInterface) NetworkName() string {
	return ni.doc.NetworkName
}

// MachineId returns the machine id of the interface.
func (ni *NetworkInterface) MachineId() string {
	return ni.doc.MachineId
}

func removeNetworkInterfacesOps(st *State, machineId string) ([]txn.Op, error) {
	var doc struct {
		MACAddress string `bson:"_id"`
	}
	ops := []txn.Op{}
	sel := bson.D{{"machineid", machineId}}
	iter := st.networkInterfaces.Find(sel).Select(bson.D{{"_id", 1}}).Iter()
	for iter.Next(&doc) {
		ops = append(ops, txn.Op{
			C:      st.networkInterfaces.Name,
			Id:     doc.MACAddress,
			Remove: true,
		})
	}
	if err := iter.Err(); err != nil {
		return nil, err
	}
	return ops, nil
}
