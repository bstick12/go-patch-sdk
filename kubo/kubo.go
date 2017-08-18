// Code generated by go-bindata.
// sources:
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/kubo.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/add-http-proxy.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/add-https-proxy.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/add-no-proxy.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/aws-cloud-provider.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/aws-lb-cloud-config.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/aws-lb.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/cf-routing.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/gcp-cloud-provider.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/k8s_haproxy_static_ips_vsphere.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/load_balancer_target_pools.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/master-haproxy-openstack.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/master-haproxy-vsphere.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/remove-haproxy.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/vsphere-cloud-provider.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/worker-haproxy-openstack.yml
// vendor/github.com/cloudfoundry-incubator/kubo-deployment/manifests/ops-files/worker-haproxy-vsphere.yml
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

var _kuboYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x57\xdd\x8e\x9b\x3a\x10\xbe\xe7\x29\x46\x5b\xa9\xea\x56\x35\xcb\x5f\x08\x70\x79\x5e\xa3\xaa\x90\x31\x43\xe2\x13\x83\xa9\x6d\xb2\x4d\xab\xbe\xfb\x91\x81\xf0\x93\x26\xdd\x74\xdb\xd3\x5c\xac\x36\xf3\xe7\x99\x6f\x66\x3e\x3b\x0d\xad\x31\x83\x77\xef\x4a\x6c\x85\x3c\xd5\xd8\x98\xdc\x8a\x1e\x1f\x1d\x47\xa1\x40\xaa\x51\x67\x0e\x81\xc1\xec\xd0\x15\x92\xa0\x61\xa5\x03\x70\x44\xa5\xb9\x6c\x32\x08\x1c\x80\x4e\x89\x0c\xf6\xc6\xb4\x3a\x7b\x7a\xd2\x46\x2a\xba\x43\x77\x27\xe5\x4e\x20\x6d\xb9\x76\x99\xac\x9f\x7a\xe7\xb6\x2b\x04\x67\x4f\x53\x20\x12\x90\xae\xe8\x1a\xd3\x11\xa3\x3a\x6d\x4e\x24\x8c\x02\xdf\xf5\x7d\x12\x78\xfe\xd6\xdb\x06\x3e\xf1\x52\x3f\xf6\x42\xb2\x49\x7d\x3f\x88\xb6\x49\x3a\x69\x7a\x45\xea\x9a\xdd\x57\x07\x40\xef\xa9\x9f\x41\xea\x57\x5b\x1a\x6f\xd2\x00\x31\x66\x9b\x20\x0a\x92\x4d\x54\x85\xf1\x26\xaa\x70\x9b\xc4\x9b\x6d\x84\x89\x1f\x57\x15\x5b\x15\xb4\xac\x45\x50\x83\xda\x4c\xea\x52\xb2\x03\xaa\x55\xb1\x89\xeb\xb9\xfe\x2b\x2a\x1e\x42\x91\xc1\xff\xe7\x35\x7b\xc4\x8f\xa3\xd0\x8f\x49\xe8\x85\xd1\x26\xde\xc6\xd1\xa4\xb1\x8a\x20\x5a\xd5\xec\x51\x96\x78\x95\xe7\x87\x8c\xc5\x49\xec\x45\x5b\x56\x96\x5b\xc6\x58\x14\x24\x34\xd9\x26\x11\xdb\x60\x9c\xfa\xc5\x5c\xf3\x9e\xb6\x4a\x7e\x39\xbd\xa2\x86\xd1\x93\x24\x6e\xe8\x7a\x77\xf4\x2d\x09\x7d\x12\x46\x49\x6a\x3b\x11\x2f\xfb\x96\x84\xfe\xaa\x06\x3f\xad\xb6\xde\xa6\x8c\xd2\x4d\x52\x04\x11\x8d\x58\x94\x62\x8a\x2c\xd9\x6e\xbd\x62\x53\x20\x46\x05\x8d\x36\x51\x81\x57\x5a\xe5\x68\x83\x35\x43\x21\xfa\x31\xa5\x82\x53\x9d\xc1\x90\x93\x03\x20\x75\x06\xab\x3c\x97\x11\x1e\xc6\x9c\x1f\x1c\x87\x37\xda\xd0\x86\x61\xbe\x53\xb2\x6b\x17\x13\x3f\x0e\xfb\x59\xaf\x33\x08\x1d\x80\x06\xcd\xb3\x54\x07\x9d\x39\x00\x67\xcb\xb7\xa3\x90\xd8\xaf\xab\x85\xd2\xf9\xa8\x7a\x7c\x74\x00\xe8\x57\x9d\xc1\xc7\xaf\xfe\x27\x07\xe0\x5f\x59\xac\x62\x8c\xa7\x01\x8c\xab\xb7\xde\x38\x80\x56\xc9\x16\x95\xe1\xd8\x7b\xd9\x8f\x55\x9d\xff\xb7\x6e\x9f\x3b\xae\x30\xd7\x5a\x64\x50\x51\xa1\x71\x52\xb5\x88\x2a\xbf\xaa\x3f\x03\xb8\x80\xed\x58\xe7\xe6\xd4\x62\x06\x4c\xd6\xb5\x6c\x1c\xeb\xae\x34\xd7\xc6\xd2\x43\xc9\xf5\x61\x54\x6f\xfc\xc0\x73\x26\xac\x6a\xaa\x4d\xbf\x2d\x0b\xb4\x82\x1b\x68\xbd\x5f\xa2\xf5\x02\x2a\x4c\xc8\xae\x24\xad\x92\x47\x5e\xf6\xf1\x2f\xf0\xb9\x84\x06\xbe\x7d\x5f\x78\x57\x82\x36\x0d\x8a\x2b\xb8\x2e\x8c\x0e\x5d\x81\xaa\x41\x83\x9a\xd0\x96\xbf\x7c\xc4\x08\x2b\x2d\x6b\xde\x90\x4e\xa3\x1a\xc2\xf4\xdf\x57\xba\x96\x6a\xfd\x2c\x55\x69\x29\xb6\xef\xe5\x5a\xdc\x4f\x84\xfd\xd8\x04\x04\x9a\x4b\xfb\x95\x6c\x32\x2e\x28\x3b\x60\x53\xe6\xad\x54\x26\x83\x24\x8a\x42\x78\x03\xff\x48\xbd\x07\xc1\x9b\x83\x86\x3d\x65\x87\xd1\x74\x30\x19\x62\x0d\x05\xe6\x43\x9b\x7a\xe7\x29\xa2\x11\x7a\x9e\xa2\xd9\xd6\x7a\x1a\xa1\xc9\x2c\xe9\x3d\x96\xa8\x31\xd9\x54\x7c\x77\x37\x62\x6b\xa0\x49\xcf\x40\x6f\xad\x7d\xde\x29\x01\x0f\x67\x32\xba\x96\xf0\x5e\x6a\xf3\xf8\x98\xdd\xae\xe5\xe1\x35\x58\xbe\xbe\xf2\xb1\x0c\x26\x1b\xa3\xa4\x10\xa8\x48\x4d\x1b\xba\xfb\xf5\x09\x5d\x04\xd3\x6c\x8f\x65\x27\x6e\xc4\xb8\xe6\x70\xb2\xeb\x4b\x74\x8b\x4c\xff\x56\x13\xde\x9f\x9b\xf0\x02\x23\x8c\x4b\x7e\xb1\xf4\x64\xbe\x54\x16\xcb\xef\xdf\x58\xfe\xdb\xe4\x08\x50\x62\x45\x3b\x61\x32\xf8\x58\x36\xfa\x03\xec\xa8\xc1\x67\x7a\xfa\x74\x41\x11\x77\x71\xd6\x25\x8f\xcc\x39\x2e\x60\x3a\x0b\x7b\x29\x93\x8d\xee\xea\x19\x28\xc3\xda\x7c\x5c\xb6\x0c\xbe\x41\xa5\x64\x7d\xae\x38\x1f\x1d\xe1\xbb\x73\x03\xe5\x3d\xcd\x7b\x8b\x79\xbc\x4a\xae\x69\x21\x30\xb7\x53\xde\x67\x3e\xd3\xb3\x3d\xc9\x2e\xef\xb8\xd3\x69\x9a\xa6\xf0\xe6\xcc\xe7\xe5\x07\x28\x3a\x03\x8d\x34\xd0\x69\x2c\x67\xf0\x2d\x6e\x97\x8c\x7b\xeb\x7e\xfa\x15\xc6\xbd\x8b\x33\xa7\xc7\xd1\xc2\x64\x21\xfb\x11\x90\x41\x39\xc3\x31\x9e\x72\x81\x04\x6f\x8d\x05\x49\x5f\x5e\x60\xbc\xb5\xdb\xfe\xf9\x52\x2c\xe4\x2e\x17\x78\xb4\x61\x50\x29\xa9\x26\xc5\xf8\x9e\xc9\x4b\xc5\x8f\xa8\x32\x90\x47\x54\x82\x9e\xce\xf7\x66\x73\xbc\xb2\x85\x7f\x80\xcc\x96\x7b\xf4\x97\xc8\xe8\xf7\xee\xc7\x31\x99\x3f\x5a\xf4\x0f\x75\x08\x34\xcb\x22\x04\x9a\x5b\x74\x7a\x65\x45\xff\x27\x26\x9b\x96\xe7\xfa\xdb\xc6\xf7\x82\xc8\x73\x9c\xae\x2d\xa9\x41\x7b\x1e\xa3\x0d\x55\xfc\xcc\x6b\x35\xfd\x92\xf3\x26\xaf\x04\xdf\xed\xcd\x20\xd2\xa8\x38\x9d\xc7\xb9\xb7\x3f\xe5\xcf\xd4\xb0\x7d\x6e\x78\xdd\xc7\xf4\x3c\x8f\x84\xf6\xaf\x67\xdf\xdf\x7d\xec\xdb\x16\xce\x91\x2a\x3e\x2c\xc3\xfa\x27\xd8\xfa\x11\xe1\x00\x0c\x29\x4f\x82\x8b\xde\xde\x67\x29\x73\x46\x27\x03\x66\x71\xae\x38\xa3\xc6\x96\x22\x5b\xc3\x65\x33\xa2\xce\x75\xce\xe8\x62\x69\x07\xbe\xcd\xc7\x59\xa4\x53\xc8\x45\xb3\xef\x09\x6b\x63\xce\x59\x5c\x84\xbd\xfd\x1e\xe8\x6d\xa9\x30\xf6\x09\x66\xf8\x11\x7b\x07\xcb\x6c\xb7\x5d\x3e\xfd\x90\xe2\x60\xf5\xf7\xb3\xec\xc5\xe4\x25\x3f\x02\xbe\xe7\xfa\x9e\xe7\x06\xde\xf0\xdb\xd3\x8a\x56\x69\xaf\x05\xee\x78\x87\xde\x54\xb8\xfa\xc8\x7e\xaa\x74\x99\xe8\x6c\x16\xae\x90\x8c\x0a\xe7\xbf\x00\x00\x00\xff\xff\x43\x33\x73\xc0\x24\x10\x00\x00")

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

	info := bindataFileInfo{name: "kubo.yml", size: 4132, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
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

	info := bindataFileInfo{name: "ops-files/add-http-proxy.yml", size: 121, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
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

	info := bindataFileInfo{name: "ops-files/add-https-proxy.yml", size: 123, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
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

	info := bindataFileInfo{name: "ops-files/add-no-proxy.yml", size: 117, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesAwsCloudProviderYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xce\xb1\x0a\x02\x31\x0c\x80\xe1\xbd\x4f\x91\xc9\x41\x38\xba\x17\xc4\x47\x29\xb1\x17\xb4\xda\x6b\x42\xd2\x5e\xf1\xed\x05\x39\x07\xc1\x45\x6e\xfd\xff\xe5\x9b\xa0\x3d\x85\x02\x28\x49\xc1\x44\x0e\x40\xb0\xdd\x02\xf8\x5c\xad\x61\x4d\x14\xaf\xca\x5d\xcc\x57\x5c\xe8\xb4\xa0\x35\x52\x7f\xe7\xcb\x16\x52\xe1\x3e\x4f\xa2\xbc\xe6\x99\xd4\x8b\xb2\x90\xb6\x4c\xe6\xbf\xcf\xd9\x01\xac\x58\x3a\x05\x38\xe0\xb0\xf8\xbe\xf1\x73\x1d\x00\x6c\x0e\x1c\xe6\xfe\x33\x0d\xd6\xc7\x5e\xd3\xf1\x87\xe9\x15\x00\x00\xff\xff\xc2\xbd\x17\xad\x1a\x01\x00\x00")

func opsFilesAwsCloudProviderYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesAwsCloudProviderYml,
		"ops-files/aws-cloud-provider.yml",
	)
}

func opsFilesAwsCloudProviderYml() (*asset, error) {
	bytes, err := opsFilesAwsCloudProviderYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/aws-cloud-provider.yml", size: 282, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesAwsLbCloudConfigYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcd\xb1\x0a\xc2\x40\x0c\x87\xf1\xfd\x9e\x22\x63\x3b\xc8\xed\x07\xe2\x83\x88\x1c\xb1\xfe\x51\x31\x67\x42\x2e\xad\xf8\xf6\x22\xdd\x3a\x75\xfe\x3e\xf8\x1d\x28\xbe\x86\x42\x0e\x13\x9e\x90\x88\x8c\xe3\x51\x28\x2f\xad\xfe\x4b\xcf\x6f\x6e\x38\x36\xee\x01\xcf\x93\xe8\x7c\xab\xe6\x6a\xf0\x78\xa2\x67\xc8\xb5\x9f\x12\xd1\xc2\x32\xa3\xd0\x79\x18\xd6\xb3\x06\xfb\x1d\x51\x4d\x55\xc6\xf1\x92\x76\x32\x1f\xf5\xd7\x3e\x66\x3d\xb7\xcc\x2f\x00\x00\xff\xff\xaa\xd8\xa7\xb8\xce\x00\x00\x00")

func opsFilesAwsLbCloudConfigYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesAwsLbCloudConfigYml,
		"ops-files/aws-lb-cloud-config.yml",
	)
}

func opsFilesAwsLbCloudConfigYml() (*asset, error) {
	bytes, err := opsFilesAwsLbCloudConfigYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/aws-lb-cloud-config.yml", size: 206, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesAwsLbYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xc8\x31\x0e\xc2\x30\x0c\x05\xd0\x3d\xa7\xf8\x63\x3a\xa0\xee\x5e\x18\x18\x39\x04\x32\xed\x17\x20\xb9\xa9\xe5\x38\x48\xb9\x3d\x0b\xeb\xbb\x20\xa7\x53\x10\x74\xd3\x8d\x05\x70\xcd\xb7\x60\x4d\x7d\xf5\x6b\x01\xbe\x6a\x83\x52\x00\xe0\x3e\x9e\x8c\xc6\x64\xbf\xd9\xe8\xc9\x10\xd4\xba\x7f\x82\x5b\x9e\xf1\x68\x7a\x70\x59\xd6\x5a\x77\xba\x9d\xf3\x60\xcb\xbf\x95\x5f\x00\x00\x00\xff\xff\x44\xa5\xab\x80\x65\x00\x00\x00")

func opsFilesAwsLbYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesAwsLbYml,
		"ops-files/aws-lb.yml",
	)
}

func opsFilesAwsLbYml() (*asset, error) {
	bytes, err := opsFilesAwsLbYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/aws-lb.yml", size: 101, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesCfRoutingYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\xcb\x6e\xe3\x30\x0c\xbc\xe7\x2b\x88\x9e\x92\x83\x60\xf4\x6a\x60\xbf\xa4\x28\x04\x56\x66\x12\x6d\x65\x49\x20\xe9\x74\xd3\xaf\x5f\xc8\x72\x1a\x3b\x75\x80\x14\x8b\x3d\x5a\x9c\xe1\x6b\xc6\x34\xa0\xe7\x4c\x2d\x30\xe5\x80\x8e\x36\x00\x19\xf5\xd8\x42\x73\x42\xf6\xf8\x16\x48\x9a\x88\x3d\xfd\xd2\x20\xe6\x7d\x78\xa3\x40\xda\xa4\xac\x3e\x45\x69\x5c\xea\xfb\x14\x6d\x89\x6f\x00\x4e\x18\x06\x6a\x61\xbb\x2d\x30\x8e\xa4\x24\xb6\x47\x51\x62\x7b\x4c\xa2\xbb\xdd\x66\xf3\x2f\xc5\x30\x28\x71\x44\xf5\x27\x1a\x2b\xca\xb5\xe4\xcb\xfd\x9a\xaf\x3f\x2f\x5a\xd3\xfc\x87\x21\x7d\x14\xc5\xe8\xc8\x1e\x38\x0d\x79\xaa\x5a\xc9\xcd\xef\xf4\x26\x8d\xf9\x2a\xb0\x01\x00\x28\xf1\x16\xae\x85\x0c\x66\x6f\x38\x0d\x4a\x86\xe9\xe0\x45\x19\x79\x04\x32\x05\x42\xa9\xd8\x34\xbe\x64\x4e\x99\x58\x3d\x49\x4d\x05\x40\x7f\xc6\xf5\x05\x5b\x30\x36\x27\xd6\xf5\x29\x4a\x64\xb7\x9b\x48\x2e\xa4\xa1\xb3\xfb\x34\xc4\x8e\xcf\x97\x4c\x00\x98\xbd\x1d\x38\x94\x04\xa5\x1d\x1f\x0f\xc6\xed\xc7\xee\x06\x0e\x5f\x64\x80\x01\x71\x05\x37\x20\xae\xe0\x5c\xf0\x14\xd5\xfa\xee\x06\x5d\xdf\x8d\xef\xd6\xf1\x42\x8e\x49\xd7\x39\x35\xf6\x23\x49\x56\x15\xa8\x2b\x97\x73\x74\xe3\xe3\x85\x23\x2d\x3c\x57\x14\xe9\x47\xe2\xf7\x69\xd5\x66\x62\x6d\xb7\x1d\xe5\x90\xce\x3d\x45\x15\x3b\x61\xa6\x21\xf0\x53\x5a\x78\xf9\x7c\x7e\x1d\xbf\x8a\xf6\x4b\x6e\x91\xc5\xa5\xb8\xf7\x87\x69\xe6\xef\x0a\xaf\x69\x0c\xb7\x66\x19\x97\xff\x74\x54\xcd\xd2\x36\xcd\x7d\xd3\xb6\xf7\x9d\xf0\xb4\xc8\x1d\x48\x4d\x46\x91\x8f\xc4\xdd\xc5\x3f\xf3\xb7\x99\x48\x1a\x66\x7d\xcd\x3b\x2b\xbc\xe5\xbf\x36\xb1\xcc\xfa\xbe\x1f\x9f\xfe\x8e\x5d\x1f\x37\x22\xc0\x14\xb6\x8f\x5a\x7c\xc9\xf8\x89\x89\x57\x99\x0f\xd9\xf9\x9a\x41\xde\x7d\xb6\x1a\xc4\x9e\x88\xfd\xde\x3b\x2c\x27\xab\x05\xe5\x81\x66\x28\xcc\xd9\x76\xa9\x47\x1f\xbf\x8d\x93\x4d\x0d\x98\xb2\xf9\x59\xea\x88\xba\x50\xaf\x47\x77\xf4\xb1\x6a\x37\x4b\x50\x60\xc6\xc7\x7a\x59\x8c\xcf\xb2\xe8\xee\x72\x64\x6e\x09\x8b\x13\x33\xca\x23\xc4\x6b\xc0\xf2\x7e\xd3\x58\xf9\x73\xaf\xf6\xfb\x96\x79\x69\x43\x51\xea\x1d\x85\x30\x6e\x44\xf4\x3c\x3e\x9e\x7a\x5b\x8f\x41\x3d\xec\x9b\xbf\x01\x00\x00\xff\xff\xb0\x40\xe3\x59\x01\x07\x00\x00")

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

	info := bindataFileInfo{name: "ops-files/cf-routing.yml", size: 1793, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesGcpCloudProviderYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x8f\xb1\x6a\x04\x31\x0c\x44\xfb\xfd\x0a\x95\xe7\xc2\xb8\x37\x84\x7c\x8a\x51\x6c\xe1\xf8\xb2\x67\x09\x59\xbb\x21\x7f\x1f\x9c\x35\x81\x94\x69\xaf\x12\x7a\x33\x30\x33\x1e\xec\x4b\x28\x82\x92\xec\x98\x69\x03\x10\xb4\xf7\x08\xa1\xf5\x61\xd8\x33\xa5\xaa\x7c\xc8\x08\x1d\x1f\xf4\xf2\xc9\xfa\x41\x1a\xee\xfc\xb6\x40\xde\xf9\x28\x5e\x94\xcf\x56\x48\x83\x28\x0b\xa9\x35\x1a\xe1\xaf\xf2\xba\x01\x9c\xb8\x1f\x14\x37\x00\x58\x99\xf5\x27\x0f\xe6\xbd\x30\x80\x28\xdf\x29\x9b\x6f\x25\xc2\xed\xb6\xbe\xd4\x8a\x73\xcb\xd0\xc9\x66\x09\x3f\xd3\xa7\x65\xfd\xbf\xfa\xd5\xd0\x77\x2e\xe4\x0d\xeb\xb4\x5c\x28\x4d\x94\x0c\xab\x73\xdb\xff\x46\x3f\x70\xd8\x13\x8c\xfe\x0e\x00\x00\xff\xff\x62\x19\xa5\x88\xea\x01\x00\x00")

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

	info := bindataFileInfo{name: "ops-files/gcp-cloud-provider.yml", size: 490, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesK8s_haproxy_static_ips_vsphereYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcc\x41\x0a\x83\x30\x10\x46\xe1\x7d\x4e\x31\x4b\xb3\x90\x74\x1d\xe8\x59\xc2\xa8\x3f\x28\xc6\x24\xcc\x4c\xda\x7a\xfb\x12\x70\xfb\x3e\x78\x33\xd9\xdd\x10\x49\xd0\x32\xaf\x70\x44\x8d\x6d\x8f\x14\x0a\xec\x5b\xe5\xd4\x30\xfc\x7d\x71\xe9\x9c\x83\xf6\xa5\xc0\x34\xbc\x82\x1a\xdb\xb1\x3a\xa2\x0f\xe7\x8e\xe8\x88\x66\x9a\xa6\xb3\x2f\x90\x02\x83\xa6\x8b\xd5\x20\x69\xaf\x6a\xde\x3f\x3c\x86\xa3\x71\x93\xfa\xbb\xd3\xd1\x12\x6f\x9b\x40\x15\xea\xbd\xfb\x07\x00\x00\xff\xff\x3a\x37\x7b\xbd\x8b\x00\x00\x00")

func opsFilesK8s_haproxy_static_ips_vsphereYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesK8s_haproxy_static_ips_vsphereYml,
		"ops-files/k8s_haproxy_static_ips_vsphere.yml",
	)
}

func opsFilesK8s_haproxy_static_ips_vsphereYml() (*asset, error) {
	bytes, err := opsFilesK8s_haproxy_static_ips_vsphereYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/k8s_haproxy_static_ips_vsphere.yml", size: 139, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesLoad_balancer_target_poolsYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xca\x31\x0a\xc2\x40\x10\x05\xd0\x7e\x4f\xf1\xcb\xa4\x90\xed\x17\xc4\xa3\x2c\x43\xfc\xa8\x30\xeb\x0c\xb3\x93\x80\xb7\x17\xb1\x49\xfd\xde\x05\xf9\x71\x36\x04\x5d\x65\x63\x01\x5c\xf2\xd9\x50\x8f\xd1\x7f\x32\xeb\x5b\x06\xaf\x43\x66\x32\xea\xa6\xb6\xdf\xbb\x87\x39\x23\x5f\x9c\x35\x25\x1e\xcc\xee\x66\x7a\x2b\xc0\x21\xba\xb3\x61\x59\xfe\xbf\x9f\x78\x5d\xcb\x37\x00\x00\xff\xff\x3f\xfc\x99\xbe\x6c\x00\x00\x00")

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

	info := bindataFileInfo{name: "ops-files/load_balancer_target_pools.yml", size: 108, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
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

	info := bindataFileInfo{name: "ops-files/master-haproxy-openstack.yml", size: 154, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
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

	info := bindataFileInfo{name: "ops-files/master-haproxy-vsphere.yml", size: 188, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
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

	info := bindataFileInfo{name: "ops-files/remove-haproxy.yml", size: 300, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesVsphereCloudProviderYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x91\xc1\x6a\xed\x20\x10\x86\xf7\x79\x8a\x59\x5d\x92\x0b\x41\xba\x15\x4a\x1f\x25\x58\xfd\x9b\x63\x9b\xa8\xcc\xa8\x87\xbe\x7d\x49\xea\x09\xc9\xa2\x8b\xd2\x9d\xf3\x7f\xdf\x8c\x03\x33\x52\xfe\x4c\xd0\xc4\x48\x8b\xb1\xe8\x88\x92\xc9\x37\x4d\xca\x07\xc9\x26\x58\x4c\x33\xc7\x92\x44\x05\xb3\xe2\x79\x35\x92\xc1\xea\x3d\xbe\xb6\xc0\x2e\xb1\xb8\x31\x71\xac\xde\x81\x55\xe2\x98\xc0\xd9\x43\xd4\x95\xbc\x74\x44\xd5\x2c\x05\x9a\xfe\x55\x49\x37\x30\xa6\xdd\x98\x1e\x46\x47\x44\x6d\x97\x26\xec\x49\x7b\xeb\xbd\x20\x2a\x02\xd6\xd4\xf7\xd5\x22\x64\xf0\xb4\xd5\xc3\xd0\x60\x32\x22\xf7\xc8\xee\x2c\x3c\xb2\x43\x12\x70\xbd\xce\xf0\xe9\x80\x3e\x08\x6c\x61\x8c\x6f\x8b\x99\x35\x3d\xb5\xd8\x99\x6c\xbe\xe5\x73\x9f\xb3\x47\xdf\x26\x48\x8e\x8c\x0b\x97\x83\xdf\x23\x7f\xf8\x30\x8f\xce\xb3\x26\x75\x1d\xa1\xea\x7a\x4a\xea\x2a\xc3\xa0\xfa\xde\x79\x86\xcd\x91\xa7\x52\xfc\xb6\xfc\xef\xee\xb4\xfd\xf7\xd7\x3b\xfd\xff\xe1\x4e\x5f\x01\x00\x00\xff\xff\x2f\x21\xd4\x13\x32\x02\x00\x00")

func opsFilesVsphereCloudProviderYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesVsphereCloudProviderYml,
		"ops-files/vsphere-cloud-provider.yml",
	)
}

func opsFilesVsphereCloudProviderYml() (*asset, error) {
	bytes, err := opsFilesVsphereCloudProviderYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/vsphere-cloud-provider.yml", size: 562, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
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

	info := bindataFileInfo{name: "ops-files/worker-haproxy-openstack.yml", size: 157, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesWorkerHaproxyVsphereYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xcd\x41\x6a\xc3\x30\x10\x85\xe1\xbd\x4e\xf1\x96\x12\xd4\x78\x2f\xe8\x49\x42\x10\x83\x35\x4d\x44\x15\x69\xd0\x8c\xeb\xea\xf6\xc5\x8d\xb7\xef\xfd\xf0\x2d\xb0\x29\x1c\x31\x58\x2a\x6d\xec\x00\x21\x7b\x46\xac\xa5\xa9\x51\xdb\x38\x3d\x46\xdf\x45\xd7\x46\x2f\xfe\x3c\xfa\xf8\xe6\xb1\x3c\x49\x46\xff\x9d\x6b\x63\x3b\x17\x75\xc0\x0f\xd5\x9d\xa3\x03\x16\x9c\x65\x84\xf7\x99\xa5\xf6\xf9\xe2\x66\x9a\xae\x32\x04\x07\x00\x99\xbf\x68\xaf\x16\x71\xcb\x4d\x3f\xf0\x20\xe3\x83\xe6\xfd\xff\x53\x23\x2b\x5b\x2a\xa2\x11\x37\x78\xff\x26\xd3\x45\xa6\x22\x89\x72\x1e\xac\xca\x1a\x02\xee\xee\x2f\x00\x00\xff\xff\xd8\xb2\x4a\xb1\xc1\x00\x00\x00")

func opsFilesWorkerHaproxyVsphereYmlBytes() ([]byte, error) {
	return bindataRead(
		_opsFilesWorkerHaproxyVsphereYml,
		"ops-files/worker-haproxy-vsphere.yml",
	)
}

func opsFilesWorkerHaproxyVsphereYml() (*asset, error) {
	bytes, err := opsFilesWorkerHaproxyVsphereYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "ops-files/worker-haproxy-vsphere.yml", size: 193, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _opsFilesWorkerHaproxyYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x92\xcd\x6e\xab\x40\x0c\x85\xf7\x3c\x85\xa5\xbb\x01\x29\x08\x65\xc9\x48\x57\x7d\x90\xa8\x1a\x0d\xe0\x04\xca\xfc\xd5\xe3\x49\x4a\x9e\xbe\x82\x0c\x69\x68\xbb\xc8\xa6\xde\x61\x1f\x9b\xef\x1c\x28\x81\x27\x8f\x02\x08\xbd\x56\x2d\x66\x00\x5e\x71\x2f\xa0\x1a\x6c\x60\x65\x5b\x94\x27\x72\xd1\x87\xaa\xcc\x00\xce\x4a\x47\x14\x19\x2c\x65\x95\x41\x01\x17\x47\x23\x52\xd9\x2b\x4f\xee\x63\x4a\xa3\x75\x37\x08\xd8\xaf\x6a\xe4\x59\x1a\xd6\xed\x32\xed\xe7\x79\x87\x5e\xbb\xc9\xa0\xe5\x20\x93\xaa\x28\x92\x0a\xa0\xc3\xa3\x8a\x9a\x05\x1c\x3a\x1b\x76\x70\x52\x8c\x17\x35\xbd\xa6\xb9\xba\x06\x01\x87\xeb\x7e\x7d\x0e\x8c\xa6\x45\xad\x05\x30\xc5\xc0\x2b\xcf\xd9\xc8\x9b\xcb\xd6\x19\xe3\x6c\xea\xbe\xb9\xe6\x07\xce\xd6\x07\x00\xa1\x46\x15\x7e\x19\xb4\xce\x86\x68\xf0\x7e\x60\x2e\x6e\xbd\x6c\x54\x3b\xa2\xed\x1e\xdb\x00\x47\x72\x66\x8d\x4a\x7e\xbf\xe4\xc9\x79\x24\x1e\xb6\xb7\x7a\x25\x17\xd9\xf6\x50\x37\x04\xd5\x68\x94\x3d\xb3\x5f\x3c\xe2\x66\x3c\x03\xe8\xc1\x8e\xd2\x3b\x62\x01\x75\x5d\xd7\xf0\x0f\x08\xdf\xe3\x40\xd8\xed\xa0\x89\x0c\xd6\x31\xc4\x80\x5d\x96\x3d\xfd\xe9\xe7\x68\xfe\xdf\xe8\xab\x39\xb4\x5b\x63\x8c\x0d\x6a\xe4\xea\x8b\xbf\x4a\xe6\x97\xd7\xbf\xdc\xff\x17\xc8\xf3\xad\x75\xf9\x10\xd4\xa2\x2d\x8a\xbf\xa0\x79\x86\xe2\x48\xce\xf2\x03\xc6\x67\x00\x00\x00\xff\xff\xa8\xe5\xf0\x1c\x0f\x03\x00\x00")

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

	info := bindataFileInfo{name: "ops-files/worker-haproxy.yml", size: 783, mode: os.FileMode(436), modTime: time.Unix(1503085548, 0)}
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
	"ops-files/aws-cloud-provider.yml": opsFilesAwsCloudProviderYml,
	"ops-files/aws-lb-cloud-config.yml": opsFilesAwsLbCloudConfigYml,
	"ops-files/aws-lb.yml": opsFilesAwsLbYml,
	"ops-files/cf-routing.yml": opsFilesCfRoutingYml,
	"ops-files/gcp-cloud-provider.yml": opsFilesGcpCloudProviderYml,
	"ops-files/k8s_haproxy_static_ips_vsphere.yml": opsFilesK8s_haproxy_static_ips_vsphereYml,
	"ops-files/load_balancer_target_pools.yml": opsFilesLoad_balancer_target_poolsYml,
	"ops-files/master-haproxy-openstack.yml": opsFilesMasterHaproxyOpenstackYml,
	"ops-files/master-haproxy-vsphere.yml": opsFilesMasterHaproxyVsphereYml,
	"ops-files/remove-haproxy.yml": opsFilesRemoveHaproxyYml,
	"ops-files/vsphere-cloud-provider.yml": opsFilesVsphereCloudProviderYml,
	"ops-files/worker-haproxy-openstack.yml": opsFilesWorkerHaproxyOpenstackYml,
	"ops-files/worker-haproxy-vsphere.yml": opsFilesWorkerHaproxyVsphereYml,
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
		"aws-cloud-provider.yml": &bintree{opsFilesAwsCloudProviderYml, map[string]*bintree{}},
		"aws-lb-cloud-config.yml": &bintree{opsFilesAwsLbCloudConfigYml, map[string]*bintree{}},
		"aws-lb.yml": &bintree{opsFilesAwsLbYml, map[string]*bintree{}},
		"cf-routing.yml": &bintree{opsFilesCfRoutingYml, map[string]*bintree{}},
		"gcp-cloud-provider.yml": &bintree{opsFilesGcpCloudProviderYml, map[string]*bintree{}},
		"k8s_haproxy_static_ips_vsphere.yml": &bintree{opsFilesK8s_haproxy_static_ips_vsphereYml, map[string]*bintree{}},
		"load_balancer_target_pools.yml": &bintree{opsFilesLoad_balancer_target_poolsYml, map[string]*bintree{}},
		"master-haproxy-openstack.yml": &bintree{opsFilesMasterHaproxyOpenstackYml, map[string]*bintree{}},
		"master-haproxy-vsphere.yml": &bintree{opsFilesMasterHaproxyVsphereYml, map[string]*bintree{}},
		"remove-haproxy.yml": &bintree{opsFilesRemoveHaproxyYml, map[string]*bintree{}},
		"vsphere-cloud-provider.yml": &bintree{opsFilesVsphereCloudProviderYml, map[string]*bintree{}},
		"worker-haproxy-openstack.yml": &bintree{opsFilesWorkerHaproxyOpenstackYml, map[string]*bintree{}},
		"worker-haproxy-vsphere.yml": &bintree{opsFilesWorkerHaproxyVsphereYml, map[string]*bintree{}},
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

