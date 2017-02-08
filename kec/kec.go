package kec

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/log"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/state"
	"io"
)

const (
	driverName               = "aliyunecs"
	defaultRegion            = "cn-hangzhou"
	defaultInstanceType      = "ecs.t1.small"
	defaultRootSize          = 20
	internetChargeType       = "PayByTraffic"
	ipRange                  = "0.0.0.0/0"
	machineSecurityGroupName = "docker-machine"
	vpcCidrBlock             = "10.0.0.0/8"
	vSwitchCidrBlock         = "10.1.0.0/24"
	timeout                  = 300
	defaultSSHUser           = "root"
	maxRetry                 = 20
)

type Driver struct {
	*drivers.BaseDriver
	Id string
}

func NewDriver(hostName, storePath string) drivers.Driver {
	id := generateId()
	return &Driver{
		Id: id,
		BaseDriver: &drivers.BaseDriver{
			SSHUser:     defaultSSHUser,
			MachineName: hostName,
			StorePath:   storePath,
		}}
}

func generateId() string {
	rb := make([]byte, 10)
	_, err := rand.Read(rb)
	if err != nil {
		log.Errorf("Unable to generate id: %s", err)
	}

	h := md5.New()
	io.WriteString(h, string(rb))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Create a host using the driver's config
func (d *Driver) Create() error {

	return nil
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {

	return nil
}

// GetCreateFlags returns the mcnflag.Flag slice representing the flags
// that can be set, their descriptions and defaults.
func (d *Driver) GetCreateFlags() []mcnflag.Flag {

	return nil
}

// GetIP returns an IP or hostname that this host is available at
// e.g. 1.2.3.4 or docker-host-d60b70a14d3a.cloudapp.net
func (d *Driver) GetIP() (string, error) {

	return nil, nil
}

// GetMachineName returns the name of the machine
func (d *Driver) GetMachineName() string {

	return nil
}

// GetSSHHostname returns hostname for use with ssh
func (d *Driver) GetSSHHostname() (string, error) {

	return nil, nil
}

// GetSSHKeyPath returns key path for use with ssh
func (d *Driver) GetSSHKeyPath() string {

	return nil
}

// GetSSHPort returns port for use with ssh
func (d *Driver) GetSSHPort() (int, error) {

	return nil, nil
}

// GetSSHUsername returns username for use with ssh
func (d *Driver) GetSSHUsername() string {

	return nil
}

// GetURL returns a Docker compatible host URL for connecting to this host
// e.g. tcp://1.2.3.4:2376
func (d *Driver) GetURL() (string, error) {

	return nil, nil
}

// GetState returns the state that the host is in (running, stopped, etc)
func (d *Driver) GetState() (state.State, error) {

	return nil, nil
}

// Kill stops a host forcefully
func (d *Driver) Kill() error {

	return nil
}

// PreCreateCheck allows for pre-create operations to make sure a driver is ready for creation
func (d *Driver) PreCreateCheck() error {

	return nil
}

// Remove a host
func (d *Driver) Remove() error {

	return nil
}

// Restart a host. This may just call Stop(); Start() if the provider does not
// have any special restart behaviour.
func (d *Driver) Restart() error {

	return nil
}

// SetConfigFromFlags configures the driver with the object that was returned
// by RegisterCreateFlags
func (d *Driver) SetConfigFromFlags(opts drivers.DriverOptions) error {

	return nil
}

// Start a host
func (d *Driver) Start() error {

	return nil
}

// Stop a host gracefully
func (d *Driver) Stop() error {

	return nil
}
