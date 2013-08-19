// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package dock

import (
	"fmt"

	"launchpad.net/juju-core/instance"
)

type lxcInstance struct {
	id string
	dockerid string
}

var _ instance.Instance = (*lxcInstance)(nil)

// Id implements instance.Instance.Id.
func (lxc *lxcInstance) Id() instance.Id {
	return instance.Id(lxc.id)
}

// Id implements instance.Instance.Id.
func (lxc *lxcInstance) DockerId() instance.Id {
	return instance.Id(lxc.dockerid)
}

func (lxc *lxcInstance) Addresses() ([]instance.Address, error) {
	logger.Errorf("lxcInstance.Addresses not implemented")
	return nil, nil
}

// DNSName implements instance.Instance.DNSName.
func (lxc *lxcInstance) DNSName() (string, error) {
	return "", instance.ErrNoDNSName
}

// WaitDNSName implements instance.Instance.WaitDNSName.
func (lxc *lxcInstance) WaitDNSName() (string, error) {
	return "", instance.ErrNoDNSName
}

// OpenPorts implements instance.Instance.OpenPorts.
func (lxc *lxcInstance) OpenPorts(machineId string, ports []instance.Port) error {
	return fmt.Errorf("not implemented")
}

// ClosePorts implements instance.Instance.ClosePorts.
func (lxc *lxcInstance) ClosePorts(machineId string, ports []instance.Port) error {
	return fmt.Errorf("not implemented")
}

// Ports implements instance.Instance.Ports.
func (lxc *lxcInstance) Ports(machineId string) ([]instance.Port, error) {
	return nil, fmt.Errorf("not implemented")
}

// Add a string representation of the id.
func (lxc *lxcInstance) String() string {
	return fmt.Sprintf("lxc:%s", lxc.id)
}