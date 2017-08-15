// Code generated by go-bindata.
// sources:
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/kubo.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/add-http-proxy.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/add-https-proxy.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/add-no-proxy.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/aws_load_balancers.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/cf-routing.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/gcp-cloud-provider.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/k8s_master_static_ip_vsphere.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/load_balancer_target_pools.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/master-haproxy-openstack.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/master-haproxy-vsphere.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/remove-haproxy.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/worker-haproxy-openstack.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/worker-haproxy.yml
// DO NOT EDIT!

package kubo

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _kuboYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\xeb\x8e\xda\x3a\x10\xfe\x9f\xa7\x18\xb5\x52\xb5\x5b\x35\x10\x20\xcb\x26\x91\x8e\x8e\x74\x5e\xa3\xaa\xa2\x89\x3d\x01\x1f\x1c\x3b\xb5\x1d\xb6\xb4\xea\xbb\x1f\x39\x77\x58\xd0\x6e\x2f\xa7\xfc\x82\xb9\x7c\x9e\xf9\xe6\x62\xa3\xb0\xa2\x0c\xee\xee\x38\xd5\x52\x9f\x2a\x52\x2e\xf7\xa2\xfb\xfb\x20\x30\x24\x09\x2d\xd9\x2c\x08\xa1\x33\x3b\x34\x85\x0e\xc9\x31\x1e\x00\x1c\xc9\x58\xa1\x55\x06\xeb\x00\xa0\x31\x32\x83\xbd\x73\xb5\xcd\x96\xcb\x9d\x70\xfb\xa6\x58\x30\x5d\x2d\x6b\x71\xd4\x0e\x65\xc8\xca\x90\xbe\xd4\x64\x84\x3f\x00\xe5\x72\x04\x5a\x0e\x87\x2c\xb9\x7e\x52\x52\x23\x5f\x1e\xd7\x93\x7a\xb1\x5e\xb8\xdd\xd7\x00\xc0\xee\x71\x95\x01\x52\xfa\x40\xdb\xed\x8a\xf1\x94\x97\x9b\x82\xb3\x87\x94\x68\x93\x14\x65\x1a\x73\x9e\x26\xb4\x2e\xd7\x49\xc4\xe3\xf2\x2c\xe0\x79\xac\x12\x1d\x59\x37\xaa\xb9\x66\x07\x32\x67\xc9\x24\x8b\x68\xb1\xba\xcc\xa8\xd0\x76\xbf\x10\x7a\xc9\xe7\xb9\xb1\x32\xac\x25\xba\x52\x9b\x2a\x24\xb5\x5b\x76\x60\xa1\xb7\xed\x93\xfa\xfb\xf8\xd7\x88\xd7\x25\x10\xc7\x09\x21\xae\xcb\xf8\x31\xe1\x2c\x79\x4c\xe3\x74\xb3\x79\x4c\x56\x45\xfc\x98\x94\x48\xd1\x1a\x31\x8e\x89\x6f\x8b\x31\xc2\x3d\xd6\x46\x7f\x39\xbd\x2e\x20\xa9\x1b\x5e\xea\x46\x71\x73\x0a\x99\xae\xaa\x46\x09\x77\x5a\xf6\x10\x17\x81\x25\x8b\xcd\x22\x1a\xe3\x8a\xb6\x09\x2f\x10\xa3\x2d\x6e\xa2\x34\x21\x16\xc5\x51\xc1\xb7\xf4\xf0\xb8\x4a\x69\xbb\x5e\x3f\xb0\x78\xb5\x79\xc4\x6d\x7a\x85\xcb\xc0\x3a\xaa\x18\x49\xd9\xf6\x09\x4a\x81\x36\x03\x67\x1a\xeb\x7c\xd4\xda\x66\xd0\x14\x8d\x72\x4d\x38\xca\x46\x84\xbb\xbb\xc1\x37\xef\x65\xbe\xed\x84\xb2\x0e\x15\xa3\x7c\x67\x74\x53\xcf\xba\xaf\x6f\xbc\x41\x6f\x33\xd8\x04\x00\x8a\xdc\x93\x36\x07\x9b\x05\x00\x83\xe5\xbb\x5e\x18\xfa\x9f\x67\xcd\x6d\xf3\x5e\x75\x7f\x1f\x00\xe0\x57\x9b\xc1\xc7\xaf\xab\x4f\x01\xc0\xbf\xba\x38\xc3\xe8\x4f\x03\xe8\x39\x3b\xef\x7e\x80\xda\xe8\x9a\x8c\x13\xd4\x7a\xf9\x8f\x57\x0d\xdf\xbd\xdb\xe7\x46\x18\xca\xad\x95\x19\x94\x28\x2d\x8d\xaa\x9a\xc8\xe4\x57\xf5\x03\x1f\x33\x06\x8f\x55\xee\x4e\x35\x65\xe0\x2b\xaa\x55\xe0\xdd\x8d\x15\xd6\xf9\x51\xe5\xc2\x1e\x7a\xf5\xc3\x6a\x1d\x05\x23\x57\x15\x5a\xd7\x76\xf6\x8c\xad\xf5\x0d\xb6\xde\xcf\xd9\x7a\x81\x95\xb6\xc7\xc2\xda\xe8\xa3\xe0\x2d\xfe\x05\x3f\x97\xd4\xc0\xb7\xef\x33\xef\x52\xa2\x52\x24\xaf\xf0\x3a\x33\x3a\x34\x05\x19\x45\x8e\x6c\x88\xb5\x78\xf9\x88\x9e\x56\xe4\x95\x50\x61\x63\xc9\x74\x30\xed\xef\x33\x5d\x8d\xd6\x3e\x69\xc3\x7d\xe3\xb5\xb5\x3c\x17\xb7\x1d\xe1\x3f\x3e\x00\x49\xee\xd2\xfe\x4c\x36\x1a\x17\xc8\x0e\xa4\x78\x5e\x6b\xe3\x32\x48\xe2\x78\x03\x6f\xe1\x1f\x6d\xf7\x20\x85\x3a\x58\xd8\x23\x3b\xf4\xa6\x9d\x49\x87\xd5\x25\x98\x77\x65\x6a\x9d\x47\x44\x27\xed\xd4\x45\x93\xad\xf7\x74\xd2\x86\x93\xa4\xf5\x98\xb3\xc6\xb4\x2a\xc5\xee\xd5\x8c\x9d\x13\x1d\xb6\x0b\xe6\x9d\xb7\xcf\x1b\x23\xe1\xcd\xb0\x6b\xae\x05\xbc\xd7\xd6\xdd\xdf\x67\xb7\x73\x79\xf3\x33\x5c\xfe\x7c\xe6\x7d\x1a\x4c\x2b\x67\xb4\x94\x64\xc2\x0a\x15\xee\x7e\xbc\x43\x67\x60\x96\xed\x89\x37\xf2\x06\xc6\x35\x87\x93\x1f\xdf\xd0\xd6\xc4\xec\x2f\x15\xe1\xfd\x50\x84\x17\x36\x42\x3f\xe4\x17\x43\x1f\x4e\x77\xc6\x6c\xf8\x57\x37\x86\xff\xf6\x72\x04\xe0\x54\x62\x23\x5d\x06\x1f\xb9\xb2\x1f\x60\x87\x8e\x9e\xf0\xf4\xe9\x62\x45\xbc\x6a\x67\x5d\xee\x91\x29\xc6\x19\x4d\x83\xb0\x95\x32\xad\x6c\x53\x4d\x44\x39\x56\xe7\xfd\xb0\x65\xf0\x0d\x4a\xa3\xab\x21\xe3\xbc\x77\x84\xef\xc1\x0d\x96\xf7\x98\xb7\x16\x53\x7b\x79\x38\x3f\xa1\xfd\xe0\xa6\x69\x9a\xc2\xdb\x61\x69\xf3\x0f\x50\x34\x0e\x94\x76\xd0\x58\xe2\x13\xc3\x9e\x9c\xcb\xb5\x7a\xeb\x12\xfa\x91\xb5\xfa\xaa\xc5\x38\xbe\x56\x66\x26\x33\xd9\xf3\xac\x3b\xe5\x94\x73\x7f\x4a\x5b\xa8\xe9\x36\x12\xb5\xc3\x42\xfa\x4c\xce\x6f\x29\x51\xfb\x91\xfe\x7c\x29\x96\x7a\x97\x4b\x3a\x7a\x18\x32\x46\x9b\x51\x61\x9d\x36\xb8\xa3\x9c\x1b\x71\x24\x93\x81\x3e\x92\x91\x78\x1a\x2e\x47\x75\xbc\x32\x6a\xbf\x61\x63\xcd\x87\xe5\x0f\x6d\x9c\x5f\xbb\x04\xfb\x60\x7e\x6b\xd2\xcf\xf2\x90\xe4\xe6\x49\x48\x72\xb7\x76\xe6\x95\x39\xfc\x9f\xd6\xd5\x38\x3c\xd7\x1f\x30\xab\x68\x1d\x47\x41\xd0\xd4\x1c\x1d\xf9\xf3\x18\x2a\x34\x62\x58\x5e\x15\x7e\xc9\x85\xca\x4b\x29\x76\x7b\xd7\x89\x2c\x19\x81\x53\x3b\xb7\xf6\xa7\xfc\x09\x1d\xdb\xe7\x4e\x54\x2d\x66\x14\x85\x9b\x28\x8a\xfc\x53\xb7\x43\xbe\xa5\x0f\x8e\x68\x44\x37\x08\xe7\xff\x77\xce\x5f\x09\x01\x40\x17\xee\x28\xb8\xa8\xeb\xeb\x2c\x75\xce\x70\x34\x60\x9e\xe3\x52\x30\x74\x3e\x0d\x5d\x3b\xa1\x55\xcf\xb8\xb0\x39\xc3\xd9\xc0\x76\x0b\x35\xef\xfb\x10\x47\xc8\x59\xa1\x5f\x03\xeb\x31\xa7\x28\x2e\x60\x6f\x5f\xf8\xad\x2d\x4a\xe7\xdf\x58\x4e\x1c\xa9\x75\xf0\x5b\xed\xb6\xcb\xa7\x67\x21\x76\x56\x7f\x3e\xca\x56\x1c\xbe\xe4\x17\xc2\x2a\x5a\xac\xa2\x68\xb1\x8e\xfc\x1f\xb7\xff\x02\x00\x00\xff\xff\xca\x9e\x0d\x5f\x1f\x0f\x00\x00")

func kuboYmlBytes() ([]byte, error) {
	return bindataRead(
		_kuboYml,
		"kubo.yml",
	)
}

func kuboYml() (*asset, error) {
	bytes, err := kuboYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kubo.yml", size: 3871, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesAddHttpProxyYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xc8\xc1\x0a\xc2\x30\x0c\x06\xe0\x7b\x9f\x22\xc7\xed\x20\xb9\x17\xc4\x47\x19\xb1\xfe\xb8\xe9\x6c\x42\x9a\x4d\xfb\xf6\x22\x08\xbb\x7e\x27\x8a\x6e\xc8\xe4\xb0\x55\x0a\x12\x91\x49\xcc\x99\x78\xa9\x2d\xa4\x16\x4c\x77\xd7\xcd\x1a\x57\x79\xe1\xfc\x56\x7f\xc2\xf9\xa1\xd7\x3f\xdc\xb4\xfc\xc0\x5c\x0d\x1e\x0b\x1a\xa3\xee\x3c\x47\xd8\x64\xae\x9f\x7e\x49\x44\xbb\xac\x1b\x32\x0d\xc3\xc1\xe3\x98\xbe\x01\x00\x00\xff\xff\x65\xb3\x91\x20\x79\x00\x00\x00")

func opsFilesAddHttpProxyYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesAddHttpProxyYml,
		"ops-files/add-http-proxy.yml",
	)
}

func opsFilesAddHttpProxyYml() (*asset, error) {
	bytes, err := opsFilesAddHttpProxyYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/add-http-proxy.yml", size: 121, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesAddHttpsProxyYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xc8\x31\x0e\xc2\x30\x0c\x05\xd0\x3d\xa7\xf0\xd8\x0e\xc8\x7b\x24\xc4\x51\x2a\x13\xbe\x68\xa1\xc4\x96\xe3\x16\x72\x7b\x84\xc4\xd0\xf5\x9d\x28\xba\x21\x93\xc3\x56\x29\x48\x44\x26\x31\x67\xe2\xa5\xb6\x90\x5a\x30\xdd\x5d\x37\x6b\x5c\xe5\x85\xf3\x5b\xfd\x09\xe7\x87\x5e\xff\x70\xd3\xf2\x03\x73\x35\x78\x2c\x68\x8c\xba\xf3\x1c\x61\x6d\x32\xd7\x4f\xbf\x24\xa2\x5d\xd6\x0d\x99\x86\xe1\xe0\xe3\x98\xbe\x01\x00\x00\xff\xff\x20\x18\x45\xe7\x7b\x00\x00\x00")

func opsFilesAddHttpsProxyYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesAddHttpsProxyYml,
		"ops-files/add-https-proxy.yml",
	)
}

func opsFilesAddHttpsProxyYml() (*asset, error) {
	bytes, err := opsFilesAddHttpsProxyYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/add-https-proxy.yml", size: 123, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesAddNoProxyYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xc8\xb1\x0a\xc2\x40\x0c\x06\xe0\xfd\x9e\x22\x63\x3b\x48\xf6\x03\xf1\x51\x4a\x3c\x7f\xb4\x5a\x93\x90\x4b\xab\x7d\x7b\x11\x74\xfd\x0e\x94\xbb\xa3\x52\xc0\x17\x69\x28\x44\x2e\x79\xab\xc4\xb3\xf6\x14\x6d\x98\xae\x61\xab\x77\x56\x79\xe2\xf8\xb2\x78\x20\xf8\x6e\xe7\x1f\x5c\xac\x7d\xc1\xc3\x1c\x91\x33\x3a\x43\x37\x56\x9b\x3c\xec\xbd\x9f\x0a\xd1\x26\xcb\x8a\x4a\xc3\xf0\xc7\x71\x2c\x9f\x00\x00\x00\xff\xff\x87\x67\xd3\xc3\x75\x00\x00\x00")

func opsFilesAddNoProxyYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesAddNoProxyYml,
		"ops-files/add-no-proxy.yml",
	)
}

func opsFilesAddNoProxyYml() (*asset, error) {
	bytes, err := opsFilesAddNoProxyYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/add-no-proxy.yml", size: 117, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesAws_load_balancersYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\xb1\x0a\xc2\x40\x0c\x87\xf1\xfd\x9e\x22\x63\x3b\xc8\xed\x07\xe2\x83\x88\x1c\xb1\xfe\x51\x31\x67\x42\x2e\xad\xf8\xf6\x22\xdd\x3a\x75\xfe\x3e\xf8\x1d\x28\xbe\x86\x42\x0e\x13\x9e\x90\x88\x8c\xe3\x51\x28\x2f\xad\xfe\x4b\xcf\x6f\x6e\x38\x36\xee\x01\xcf\x93\xe8\x7c\xab\xe6\x6a\xf0\x78\xa2\x67\xc8\xb5\x9f\x12\xd1\xc2\x32\xa3\xd0\x79\x18\xd6\xb3\x06\xfb\x1d\x51\x4d\x55\xc6\xf1\x92\x76\x32\x1f\xf5\xd7\x3e\x66\x3d\xb7\xcc\x2f\x00\x00\xff\xff\xaa\xd8\xa7\xb8\xce\x00\x00\x00")

func opsFilesAws_load_balancersYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesAws_load_balancersYml,
		"ops-files/aws_load_balancers.yml",
	)
}

func opsFilesAws_load_balancersYml() (*asset, error) {
	bytes, err := opsFilesAws_load_balancersYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/aws_load_balancers.yml", size: 206, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesCfRoutingYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x55\x41\x6e\xe3\x3a\x0c\xdd\xe7\x14\x44\x57\xc9\x42\x75\xf2\x97\x06\xfe\x49\x8a\x42\x60\x65\x26\xd1\x54\x96\x04\x91\x4e\x9b\x9e\x7e\x20\xcb\x99\xd8\x89\xdd\x69\x31\xc0\x8c\x77\x12\x1f\xc9\x47\x91\x8f\x56\x20\xe7\x48\x35\x24\x8a\x0e\x0d\xad\x00\x22\xca\xb1\x86\xea\x84\xc9\xe2\x8b\x23\xae\x3c\xb6\xf4\xbf\x38\x56\xaf\xdd\x0b\x39\x92\x2a\x44\xb1\xc1\x73\x65\x42\xdb\x06\xaf\xb3\x7d\x05\x70\x42\xd7\x51\x0d\xeb\x75\x86\x25\x4f\x42\xac\x5b\x64\xa1\xa4\x8f\x81\x65\xb3\x59\xad\xfe\x24\x19\x3a\xa1\xe4\x51\xec\x89\xfa\x8c\x7c\x4d\xf9\xb4\x9c\xf3\xf9\xfb\x49\x4b\x98\xbf\x57\xe4\x4d\xbe\xe5\x3a\x57\x00\x0a\xd6\x6b\xb3\x57\x62\xa2\x4a\xa1\x13\x4a\x2a\x43\x36\x9b\xc1\xb4\x48\x2a\x9b\x77\xdb\xc7\xdd\x76\xfb\xf8\xdf\x76\xfb\xb8\xfb\x84\xa5\xf5\x2c\xe8\x0d\xe9\x43\x0a\x5d\x1c\xb8\x96\x68\xd5\x8f\xf0\xc2\x95\x1a\x13\x02\xc8\xf6\x1a\xae\x99\x15\x46\x5b\xc8\xa9\x44\x07\xcb\x92\x30\xf5\xc0\x44\x8e\x90\x0b\x36\xf4\x37\x31\x85\x48\x49\x2c\x71\x09\x05\x40\xef\x7d\xf1\x4e\x67\x8c\x8e\x21\xc9\xfc\x5b\x67\x4b\x5f\x56\xfe\x8c\x0b\x5d\xa3\xf7\xa1\xf3\x4d\x3a\x5f\x22\x01\x60\xb4\xba\x4b\x2e\x07\xc8\x74\xac\x3f\x28\xb3\xef\xd9\x75\xc9\xfd\x72\x06\xe8\x10\x67\x70\x1d\xe2\x0c\xce\x38\x4b\x5e\xb4\x6d\x6e\xd0\xe5\x5e\xd9\x66\x1e\xcf\x64\x12\xc9\xbc\x4f\xb1\x7d\x3a\x38\xb7\x2d\x99\xed\x40\x4c\xe1\xfd\xdc\x9f\x2f\x70\xae\x61\x57\x00\x24\x6f\x21\xbd\x0e\xaf\xac\x06\x87\xf5\xba\xa1\xe8\xc2\xb9\x25\x2f\xac\x07\xcc\xc0\x1f\x3f\xb8\x86\xa7\x8f\xdd\x73\x7f\xca\x6d\x9f\xfa\xfa\x83\xf5\xef\x43\xa5\xf7\x7d\xbd\xa0\xf6\x0e\xbd\x27\xd7\xfc\x16\x98\x1b\x6c\x82\xdf\xdb\xc3\x22\x74\x6e\x5a\xe0\x76\xec\xfa\x36\x3e\x1c\x45\x22\xd7\x55\xb5\xac\x87\x7a\x79\xa6\x1e\x26\xb1\x1d\x89\x8a\xc8\xfc\x16\x52\x73\x99\xc4\xf1\xdd\xa8\xdd\xe2\x46\xbc\xc6\xcc\xb2\xdf\x54\xeb\x83\x97\xba\xd7\xce\xb5\x89\xff\xfe\x0d\x2e\xf4\x8a\x96\xf9\xec\xcd\x37\x89\x2d\xe8\xf2\xeb\x8a\x03\x18\xcc\xfa\xab\x5a\x9e\x7a\x7c\x47\xad\xb3\x9e\x5f\xd2\xed\x35\x02\xbf\xda\xa8\xc5\xb1\x3e\x51\xb2\x7b\x6b\x30\x6f\xf4\x1a\x24\x75\x34\x42\x61\x8c\xba\x09\x2d\x5a\x7f\x57\x4e\x54\xc5\x70\xdd\xea\xe5\xf3\x28\x93\xe1\x6a\xd1\x1c\xad\x2f\xa3\x35\x0a\x90\x61\xca\xfa\xb2\x42\x95\x8d\x3c\x61\x77\xd9\xa6\xb7\x0e\x93\x5d\xda\xb7\x87\x29\xcd\x01\xf3\xfd\x0d\xb1\xbc\xa2\xae\xea\xb8\x8b\x3c\x55\x09\x0b\xb5\x86\x9c\xeb\x5f\x84\xa5\xcc\xf9\xa9\xd5\x65\xeb\x95\xff\xec\xea\x67\x00\x00\x00\xff\xff\x86\xe0\x9a\x91\x90\x08\x00\x00")

func opsFilesCfRoutingYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesCfRoutingYml,
		"ops-files/cf-routing.yml",
	)
}

func opsFilesCfRoutingYml() (*asset, error) {
	bytes, err := opsFilesCfRoutingYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/cf-routing.yml", size: 2192, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesGcpCloudProviderYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x8f\xb1\xce\xc3\x20\x10\x83\x77\x9e\xc2\x63\x18\x10\x3b\xd2\xaf\xff\x51\x22\x0a\xa7\x94\x34\x81\xd3\x41\x52\xf5\xed\x2b\x5a\x96\x8e\xdd\x3a\x9d\xec\xb3\xf4\xd9\x06\xed\xc1\xe4\x20\xc4\x9b\x0f\xa4\x00\xf6\xed\xea\x60\x53\xae\xcd\xe7\x40\xf3\x22\xe5\xe0\x6a\xb3\xdf\xe9\xef\x5e\xe4\x46\x62\xd7\x72\x19\x46\xd8\xca\x11\x0d\x4b\x39\x53\x24\xb1\x2c\x85\x49\x5a\xa2\x6a\x3f\x3f\xff\x0a\x38\xfd\x76\x90\x53\x00\x06\x73\x79\xf1\xd0\xef\xdb\x06\x58\xca\x4a\xa1\x99\x14\x1d\xa6\x69\xa8\x39\x45\xad\x47\x20\x53\xeb\x25\x4c\xa7\xf7\xc8\xd0\x5a\xab\xef\x96\xec\xbe\xb6\x5f\x5d\xf2\x0c\x00\x00\xff\xff\x46\xcc\x73\xf1\x94\x01\x00\x00")

func opsFilesGcpCloudProviderYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesGcpCloudProviderYml,
		"ops-files/gcp-cloud-provider.yml",
	)
}

func opsFilesGcpCloudProviderYml() (*asset, error) {
	bytes, err := opsFilesGcpCloudProviderYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/gcp-cloud-provider.yml", size: 404, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesK8s_master_static_ip_vsphereYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x14\xc9\x41\x0a\x83\x30\x10\x05\xd0\x7d\x4e\xf1\x97\xba\x90\xe9\x3a\xd0\xb3\xc8\x28\x1f\x2c\xc6\x18\x32\x3f\x2d\xbd\x7d\xe9\xfa\x2d\xd0\xb7\x31\xa3\xb3\x15\xdf\x99\x80\xe6\x3a\x32\xac\x52\x9f\xbb\x9f\x61\x7f\x7f\x5e\x5e\x87\x17\x8b\xb1\x55\x2a\xec\x61\x21\xd7\x6b\x4f\xc0\xdb\xcb\x60\x4e\xc0\x82\x69\x3a\xc7\xc6\x5e\x29\xc6\x7a\x79\x88\x7d\x3d\xee\xd0\x3c\xa7\x5f\x00\x00\x00\xff\xff\x5c\xe5\xe8\xa5\x67\x00\x00\x00")

func opsFilesK8s_master_static_ip_vsphereYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesK8s_master_static_ip_vsphereYml,
		"ops-files/k8s_master_static_ip_vsphere.yml",
	)
}

func opsFilesK8s_master_static_ip_vsphereYml() (*asset, error) {
	bytes, err := opsFilesK8s_master_static_ip_vsphereYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/k8s_master_static_ip_vsphere.yml", size: 103, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesLoad_balancer_target_poolsYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xcd\xb1\x0a\x02\x31\x0c\x87\xf1\xbd\x4f\x91\xf1\x6e\x90\xee\x05\xf1\x51\x4a\x38\xff\xa8\x98\x9a\x90\xe6\x4e\x7c\x7b\x91\x2e\xe2\xe4\xcd\xdf\x07\xbf\x03\xc5\xcb\x50\xc8\x61\xc2\x0b\x12\x91\x71\x5c\x0b\xe5\xad\xd5\x4f\xe9\xf9\xc1\x0d\xc7\xc6\x3d\xe0\x79\x11\x5d\xcf\xd5\x5c\x0d\x1e\x37\xf4\x1c\xec\x17\x44\x35\x55\x39\x25\xa2\x8d\x65\x45\xa1\x69\x1a\x7f\xfd\xca\xf3\x9c\xfe\xb4\x9e\xea\xf7\x3d\xd6\xf8\x7f\xac\x77\x00\x00\x00\xff\xff\x52\x7a\x0e\x9c\xd8\x00\x00\x00")

func opsFilesLoad_balancer_target_poolsYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesLoad_balancer_target_poolsYml,
		"ops-files/load_balancer_target_pools.yml",
	)
}

func opsFilesLoad_balancer_target_poolsYml() (*asset, error) {
	bytes, err := opsFilesLoad_balancer_target_poolsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/load_balancer_target_pools.yml", size: 216, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesMasterHaproxyOpenstackYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\xcd\x41\xaa\xc3\x30\x0c\x84\xe1\xbd\x4f\x31\xcb\x64\x11\xbc\x37\xbc\x93\x3c\x8a\x51\x83\x68\x4c\x12\x5b\x48\x72\xda\xdc\xbe\xb4\xa1\xcb\x0f\x86\xf9\x27\xf8\x29\x9c\xa0\x2c\x1b\xcd\x1c\x00\x21\x5f\x12\x62\xa9\xe6\x54\x67\xce\x0f\x6d\x5d\x2c\x56\xda\xf9\x6f\x27\x73\xd6\x69\x21\xd1\xf6\x3a\x63\x65\x7f\x36\x5d\x2d\x4e\x01\x38\x68\xeb\x9c\x02\x00\x7c\xb6\x09\x47\x91\xaf\xae\xc0\x4f\xe6\xe4\x65\xce\x45\x2c\xe1\x7f\x18\xd6\x7e\x67\xad\xec\x6c\xf9\x3a\xcf\x4b\x33\x1f\xc7\x5b\x78\x07\x00\x00\xff\xff\x30\xb4\x30\x48\x9a\x00\x00\x00")

func opsFilesMasterHaproxyOpenstackYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesMasterHaproxyOpenstackYml,
		"ops-files/master-haproxy-openstack.yml",
	)
}

func opsFilesMasterHaproxyOpenstackYml() (*asset, error) {
	bytes, err := opsFilesMasterHaproxyOpenstackYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/master-haproxy-openstack.yml", size: 154, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesMasterHaproxyVsphereYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcd\xcd\x6a\x85\x30\x10\x47\xf1\x7d\x9e\xe2\xbf\x4c\xa0\xe2\x3e\xd0\x27\x11\x09\x53\x9d\xaa\x18\x93\x90\x99\xd4\xe6\xed\x2f\xf7\x63\x7d\x7e\x70\x06\x68\x2f\xec\x51\xb9\x44\x5a\xd8\x00\x85\x74\xf7\x18\x8f\x24\x4a\x69\xe1\xb0\xd5\xdc\x8a\x8c\x89\x2e\xfe\xbe\x48\x94\xeb\xb0\x53\xa9\xf9\xbf\x8f\x89\xf5\xce\xf5\x14\x03\xfc\x51\x6c\xec\x0d\x30\xe0\x29\x3d\xac\x5d\xb9\xc4\xdc\x2f\x4e\x2a\xe1\x23\x9d\x33\x00\xb0\xf2\x2f\xb5\xa8\x1e\xd3\x9a\xe4\x0b\x1b\x29\xdf\xd4\xe7\x57\x13\x25\x3d\x96\x70\x14\xf1\x98\x60\xed\xd9\x7e\xb8\x26\x56\x96\xf0\xbe\x87\x3d\x8b\x3a\x87\xd9\x3c\x02\x00\x00\xff\xff\x72\x22\x7d\x0f\xbc\x00\x00\x00")

func opsFilesMasterHaproxyVsphereYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesMasterHaproxyVsphereYml,
		"ops-files/master-haproxy-vsphere.yml",
	)
}

func opsFilesMasterHaproxyVsphereYml() (*asset, error) {
	bytes, err := opsFilesMasterHaproxyVsphereYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/master-haproxy-vsphere.yml", size: 188, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesRemoveHaproxyYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x8d\x41\x0e\xc2\x30\x0c\x04\xef\x7d\x45\x3e\x10\xe5\x5e\x89\xb7\x54\x4e\x59\xd1\x50\x62\x5b\xb6\x8b\xe8\xef\x91\x50\xaf\xc0\xa9\xd7\xd5\xce\x4c\x4e\xb1\x2b\xc6\x64\xe8\xf2\xc4\x90\x92\x52\x2c\x63\x2a\x8d\x3d\x88\x67\x4c\x37\x93\x4d\xbd\x30\x75\x5c\x3a\x79\xc0\xf2\x42\x6a\xf2\xda\x87\x6f\xac\xe1\x01\x72\x1c\xd0\xbf\xf7\x8f\x52\xb9\x4b\x3d\x86\x75\xab\x30\x46\xc0\x33\x69\x2b\x6a\xa2\xb0\x68\xf0\x52\x69\x5e\xc1\xd7\x49\xc5\xe2\xb4\xc8\x47\xfe\x0e\x00\x00\xff\xff\x98\xff\x9b\x33\x2c\x01\x00\x00")

func opsFilesRemoveHaproxyYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesRemoveHaproxyYml,
		"ops-files/remove-haproxy.yml",
	)
}

func opsFilesRemoveHaproxyYml() (*asset, error) {
	bytes, err := opsFilesRemoveHaproxyYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/remove-haproxy.yml", size: 300, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesWorkerHaproxyOpenstackYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x34\x8d\x41\x0e\x83\x30\x0c\x04\xef\x79\xc5\x1e\xe1\x80\x72\x8f\xd4\xb7\x44\x56\xb0\x4a\x54\x9a\x58\xb6\xa1\xe5\xf7\x15\xa5\x3d\xce\x6a\x35\x33\xc1\x0f\xe1\x04\x65\x59\xa9\x70\x00\x84\x7c\x49\x88\xb5\x99\x53\x2b\x9c\xef\xda\x37\xb1\xd8\xe8\xc9\xb7\x57\xd7\x07\xeb\xb4\x90\x68\x7f\x1f\xb1\xb1\x9f\x8b\xc5\x29\x00\x3b\xad\x1b\xa7\x00\x00\xe7\x37\x61\xaf\xf2\xa5\x2b\xf0\x27\x73\xf2\x5a\x72\x15\x4b\x18\x86\x4b\x98\x7f\xc2\x5c\x25\xd3\x3c\x2b\x9b\xb1\x8d\x63\xf8\x04\x00\x00\xff\xff\x3a\xdb\xe0\xf6\x9d\x00\x00\x00")

func opsFilesWorkerHaproxyOpenstackYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesWorkerHaproxyOpenstackYml,
		"ops-files/worker-haproxy-openstack.yml",
	)
}

func opsFilesWorkerHaproxyOpenstackYml() (*asset, error) {
	bytes, err := opsFilesWorkerHaproxyOpenstackYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/worker-haproxy-openstack.yml", size: 157, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesWorkerHaproxyYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\xcf\x8e\xa3\x30\x0c\xc6\xef\x3c\x85\xa5\xbd\x80\x54\x84\x7a\x24\xd2\x6a\x1f\xa4\x5a\x45\x21\xb8\x2d\x4b\xfe\xad\xe3\xb4\x43\x9f\x7e\x04\x0d\x9d\x32\x33\x87\x5e\x26\x37\xec\xcf\x5f\x7e\x9f\x49\x0d\x3c\x05\x14\x40\x18\x8c\xd2\x58\x00\x04\xc5\x67\x01\xcd\xe0\x22\x2b\xa7\x51\x9e\xc8\xa7\x10\x9b\xba\x00\xb8\x28\x93\x50\x14\xb0\x1c\xa7\x2c\x0a\xb8\x7a\x1a\x91\xea\xb3\x0a\xe4\xdf\xa6\xdc\x5a\x67\xa3\x80\xfd\xaa\x46\x9e\xa5\x71\x9d\xae\xf3\x7c\x59\xf6\x18\x8c\x9f\x2c\x3a\x8e\x32\xab\xaa\x2a\xab\x00\x7a\x3c\xaa\x64\x58\xc0\xa1\x77\x71\x07\x27\xc5\x78\x55\xd3\xdf\xdc\x57\xb7\x28\xe0\x70\xdb\xaf\xdf\x91\xd1\x6a\x34\x46\x00\x53\x8a\xbc\xf2\x5c\xac\xbc\xa7\xd4\xde\x5a\xef\x72\xf5\x9f\xef\xbe\xe0\x6c\x73\x00\x10\x1a\x54\xf1\x9b\x86\xf6\x2e\x26\x8b\x0f\x83\xf9\xb0\x0e\xb2\x53\x7a\x44\xd7\x3f\x97\x01\x8e\xe4\xed\xba\x2a\xf9\xd9\x29\x90\x0f\x48\x3c\x6c\xbd\xce\x4a\x2e\xb2\xad\xd1\x7c\x83\x19\xdc\x28\x83\x27\x16\xd0\xb6\x6d\x0b\xbf\x80\xf0\x7f\x1a\x08\xfb\x1d\x74\x89\xc1\x79\x86\x14\xb1\x2f\x8a\x97\xff\xed\x9c\xfd\xf7\x1d\xaf\x99\xb7\x72\x2f\x8c\xa9\x43\x83\xdc\x7c\x00\x36\x39\xdd\x72\xfd\x9f\xc7\x83\x80\xb2\xdc\x66\x93\x4f\x9b\x58\xb4\x55\xf5\x13\x34\xaf\x50\x1c\xc9\x3b\x7e\xc2\x78\x0f\x00\x00\xff\xff\xea\xce\xe8\x8b\xf0\x02\x00\x00")

func opsFilesWorkerHaproxyYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesWorkerHaproxyYml,
		"ops-files/worker-haproxy.yml",
	)
}

func opsFilesWorkerHaproxyYml() (*asset, error) {
	bytes, err := opsFilesWorkerHaproxyYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/worker-haproxy.yml", size: 752, mode: os.FileMode(436), modTime: time.Unix(1502744098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"kubo.yml": kuboYml,
	"ops-files/add-http-proxy.yml": opsFilesAddHttpProxyYml,
	"ops-files/add-https-proxy.yml": opsFilesAddHttpsProxyYml,
	"ops-files/add-no-proxy.yml": opsFilesAddNoProxyYml,
	"ops-files/aws_load_balancers.yml": opsFilesAws_load_balancersYml,
	"ops-files/cf-routing.yml": opsFilesCfRoutingYml,
	"ops-files/gcp-cloud-provider.yml": opsFilesGcpCloudProviderYml,
	"ops-files/k8s_master_static_ip_vsphere.yml": opsFilesK8s_master_static_ip_vsphereYml,
	"ops-files/load_balancer_target_pools.yml": opsFilesLoad_balancer_target_poolsYml,
	"ops-files/master-haproxy-openstack.yml": opsFilesMasterHaproxyOpenstackYml,
	"ops-files/master-haproxy-vsphere.yml": opsFilesMasterHaproxyVsphereYml,
	"ops-files/remove-haproxy.yml": opsFilesRemoveHaproxyYml,
	"ops-files/worker-haproxy-openstack.yml": opsFilesWorkerHaproxyOpenstackYml,
	"ops-files/worker-haproxy.yml": opsFilesWorkerHaproxyYml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"kubo.yml": &bintree{kuboYml, map[string]*bintree{}},
	"ops-files": &bintree{nil, map[string]*bintree{
		"add-http-proxy.yml": &bintree{opsFilesAddHttpProxyYml, map[string]*bintree{}},
		"add-https-proxy.yml": &bintree{opsFilesAddHttpsProxyYml, map[string]*bintree{}},
		"add-no-proxy.yml": &bintree{opsFilesAddNoProxyYml, map[string]*bintree{}},
		"aws_load_balancers.yml": &bintree{opsFilesAws_load_balancersYml, map[string]*bintree{}},
		"cf-routing.yml": &bintree{opsFilesCfRoutingYml, map[string]*bintree{}},
		"gcp-cloud-provider.yml": &bintree{opsFilesGcpCloudProviderYml, map[string]*bintree{}},
		"k8s_master_static_ip_vsphere.yml": &bintree{opsFilesK8s_master_static_ip_vsphereYml, map[string]*bintree{}},
		"load_balancer_target_pools.yml": &bintree{opsFilesLoad_balancer_target_poolsYml, map[string]*bintree{}},
		"master-haproxy-openstack.yml": &bintree{opsFilesMasterHaproxyOpenstackYml, map[string]*bintree{}},
		"master-haproxy-vsphere.yml": &bintree{opsFilesMasterHaproxyVsphereYml, map[string]*bintree{}},
		"remove-haproxy.yml": &bintree{opsFilesRemoveHaproxyYml, map[string]*bintree{}},
		"worker-haproxy-openstack.yml": &bintree{opsFilesWorkerHaproxyOpenstackYml, map[string]*bintree{}},
		"worker-haproxy.yml": &bintree{opsFilesWorkerHaproxyYml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
