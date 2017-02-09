package ksyunkec

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

var (
	dockerPort = 2376
	swarmPort  = 3376
)

const (
	driverName          = "ksyunkec"
	defaultSSHUser      = "root"
	defaultInstanceType = "I1.1A"
)

type Driver struct {
	*drivers.BaseDriver
	InstanceId         string
	Id                 string
	AccessKey          string
	SecretKey          string
	ImageId            string
	InstanceType       string
	DataDiskGb         int
	MaxCount           string
	MinCount           string
	SubnetId           string
	InstancePassword   string
	ChargeType         string
	PurchaseTime       string
	SecurityGroupId    string
	PrivateIpAddress   string
	InstanceName       string
	InstanceNameSuffix string
	SriovNetSupport    bool
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
	return driverName
}

// GetCreateFlags returns the mcnflag.Flag slice representing the flags
// that can be set, their descriptions and defaults.
func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{
		mcnflag.StringFlag{
			Name:   "ksyunkec-access-key-id",
			Usage:  "KEC Access Key ID",
			Value:  "",
			EnvVar: "KEC_ACCESS_KEY_ID",
		},
		mcnflag.StringFlag{
			Name:   "ksyunkec-access-key-secret",
			Usage:  "KEC Access Key Secret",
			Value:  "",
			EnvVar: "KEC_ACCESS_KEY_SECRET",
		},
		mcnflag.StringFlag{
			Name:   "ksyunkec-image-id",
			Usage:  "KEC machine image",
			EnvVar: "KEC_IMAGE_ID",
		},
		mcnflag.StringFlag{
			Name:   "ksyunkec-instance-type",
			Usage:  "KEC instance type",
			Value:  defaultInstanceType,
			EnvVar: "KEC_INSTANCE_TYPE",
		},
		mcnflag.IntFlag{
			Name:   "ksyunkec-disk-size",
			Usage:  "Data disk size for instance in GB",
			Value:  0,
			EnvVar: "KEC_DISK_SIZE",
		},
		mcnflag.StringFlag{
			Name:   "ksyunkec-subnet-id",
			Usage:  "Subnet id for instance association",
			EnvVar: "KEC_SUBNET_ID",
		},
		mcnflag.StringFlag{
			Name:   "ksyunkec-ssh-password",
			Usage:  "set the password of the ssh user",
			EnvVar: "KEC_SSH_PASSWORD",
		},
		mcnflag.StringFlag{
			Name:   "ksyunkec-charge-type",
			Usage:  "set the charge type of instance",
			Value:  "Monthly",
			EnvVar: "KEC_CHARGE_TYPE",
		},
		mcnflag.IntFlag{
			Name:   "ksyunkec-purchase-time",
			Usage:  "set the purchase time of instance",
			Value:  0,
			EnvVar: "KEC_PURCHASE_TIME",
		},
		mcnflag.StringFlag{
			Name:   "ksyunkec-security-group",
			Usage:  "KEC security group",
			EnvVar: "KEC_SECURITY_GROUP",
		},
		mcnflag.StringFlag{
			Name:   "ksyunkec-private-ip",
			Usage:  "KEC instance private IP",
			Value:  "",
			EnvVar: "KEC_PRIVATE_IP",
		},
		mcnflag.StringFlag{
			Name:   "ksyunkec-instance-name",
			Usage:  "Name for instance",
			Value:  "",
			EnvVar: "KEC_INSTANCE_NAME",
		},
	}
}

// GetIP returns an IP or hostname that this host is available at
// e.g. 1.2.3.4 or docker-host-d60b70a14d3a.cloudapp.net
func (d *Driver) GetIP() (string, error) {

	return nil, nil
}

// GetMachineName returns the name of the machine
func (d *Driver) GetMachineName() string {
	return d.MachineName
}

// GetSSHHostname returns hostname for use with ssh
func (d *Driver) GetSSHHostname() (string, error) {

	return nil, nil
}

// GetSSHKeyPath returns key path for use with ssh
func (d *Driver) GetSSHKeyPath() string {
	if d.SSHKeyPath == "" {
		d.SSHKeyPath = d.ResolveStorePath("id_rsa")
	}
	return d.SSHKeyPath
}

// GetSSHPort returns port for use with ssh
func (d *Driver) GetSSHPort() (int, error) {
	if d.SSHPort == 0 {
		d.SSHPort = drivers.DefaultSSHPort
	}

	return d.SSHPort, nil
}

// GetSSHUsername returns username for use with ssh
func (d *Driver) GetSSHUsername() string {
	if d.SSHUser == "" {
		d.SSHUser = drivers.DefaultSSHUser
	}
	return d.SSHUser
}

// GetURL returns a Docker compatible host URL for connecting to this host
// e.g. tcp://1.2.3.4:2376
func (d *Driver) GetURL() (string, error) {
	ip, err := d.GetIP()
	if err != nil {
		return "", err
	}
	if ip == "" {
		return "", nil
	}
	return fmt.Sprintf("tcp://%s:%d", ip, dockerPort), nil
}

// GetState returns the state that the host is in (running, stopped, etc)
func (d *Driver) GetState() (state.State, error) {

	return nil, nil
}

// Kill stops a host forcefully
func (d *Driver) Kill() error {
	log.Debugf("%s | Killing instance ...", d.MachineName)

	//TODO send stop vm http request.
	//if err := d.getClient().StopInstance(d.InstanceId, true); err != nil {
	//	return fmt.Errorf("%s | Unable to kill instance %s: %s", d.MachineName, d.InstanceId, err)
	//}
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
	d.AccessKey = opts.String("ksyunkec-access-key-id")
	d.SecretKey = opts.String("ksyunkec-access-key-secret")
	d.ImageId = opts.String("ksyunkec-image-id")
	d.InstanceType = opts.String("ksyunkec-instance-type")
	d.DataDiskGb = opts.Int("ksyunkec-disk-size")
	d.SubnetId = opts.String("ksyunkec-subnet-id")
	d.InstancePassword = opts.String("ksyunkec-ssh-password")
	d.ChargeType = opts.String("ksyunkec-charge-type")
	d.PurchaseTime = opts.String("ksyunkec-purchase-time")
	d.SecurityGroupId = opts.String("ksyunkec-security-group")
	d.PrivateIpAddress = opts.String("ksyunkec-private-ip")
	d.InstanceName = opts.String("ksyunkec-instance-name")
	d.InstanceNameSuffix = ""
	d.SriovNetSupport = false
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
