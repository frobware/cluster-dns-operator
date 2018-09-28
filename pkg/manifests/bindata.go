// Code generated by go-bindata.
// sources:
// manifests/00-cluster-role.yaml
// manifests/00-custom-resource-definition.yaml
// manifests/00-namespace.yaml
// manifests/01-cluster-role-binding.yaml
// manifests/01-service-account.yaml
// manifests/02-deployment.yaml
// manifests/image-references
// assets/dns/.daemonset.yaml.swp
// assets/dns/cluster-role-binding.yaml
// assets/dns/cluster-role.yaml
// assets/dns/configmap.yaml
// assets/dns/daemonset.yaml
// assets/dns/namespace.yaml
// assets/dns/service-account.yaml
// assets/dns/service.yaml
// DO NOT EDIT!

package manifests

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

var _manifests00ClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\xb1\x6e\xf3\x30\x0c\x84\x77\x3d\x05\x91\x7f\xb6\x7f\x64\x2b\xbc\x76\xe8\xde\xa1\x3b\x2d\xd1\x31\x11\x9b\x14\x48\x3a\x05\xfa\xf4\x45\x12\x3b\x28\xea\x74\x12\x4f\x3c\x7d\x94\x4e\xff\xe0\x75\x5a\x3c\xc8\xc0\x74\x22\x18\xd4\x20\x46\x02\xad\x64\x18\x6a\xc0\xe1\x34\x0d\x6d\x3a\xb3\x94\x6e\xf3\xbe\xeb\x44\x09\x2b\x7f\x90\x39\xab\x74\x60\x3d\xe6\x16\x97\x18\xd5\xf8\x0b\x83\x55\xda\xf3\x8b\xb7\xac\xff\x2f\xc7\x9e\x02\x8f\x69\xa6\xc0\x82\x81\x5d\x02\x10\x9c\xa9\x83\x7c\x67\x35\x45\xbc\xd9\xc6\xad\x4d\xaf\x98\xa9\xbb\x5e\x42\x7c\xe4\x21\x9a\xa7\x5e\x5b\x26\xf2\x2e\x35\x80\x95\xdf\x4c\x97\xea\x57\x78\x03\x45\xbc\x7d\x1c\x6d\x59\x13\x80\x91\xeb\x62\x99\x56\xc7\x8a\x2b\xe2\xe4\x09\xe0\x42\xd6\xaf\x9d\x89\x3d\x6e\xc5\x27\x46\x1e\xf7\x6c\xac\xd5\xf7\xbc\x82\x34\xab\x38\xc5\x2f\x5a\x36\xc2\xa0\x3d\xe5\x70\xb8\x2d\x7f\xa6\xb6\x9f\xf0\x48\xc5\x6f\xb2\x6a\xb9\x17\x24\xa5\x2a\x4b\xdc\x95\x93\x5d\x78\xf3\xac\x02\x73\xd6\x65\x33\x64\x95\x81\x4f\x33\x56\xff\x99\xc3\xf5\xe7\x77\x1b\x3d\x4b\x61\x39\x3d\x7f\xd1\x93\xa8\xbe\x03\x00\x00\xff\xff\x78\xca\xa1\x8e\x49\x02\x00\x00")

func manifests00ClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests00ClusterRoleYaml,
		"manifests/00-cluster-role.yaml",
	)
}

func manifests00ClusterRoleYaml() (*asset, error) {
	bytes, err := manifests00ClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/00-cluster-role.yaml", size: 585, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests00CustomResourceDefinitionYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x90\xb1\x4e\x03\x31\x0c\x86\xf7\x3c\x85\x9f\xe0\xd0\x6d\x28\x2b\xdd\x40\x1d\x40\x62\x37\x89\xdb\x5a\xcd\xd9\x56\xec\x9c\x78\x7c\x74\x57\x90\x80\x8e\xf6\xf7\xeb\xff\x2c\xa3\xf1\x3b\x75\x67\x95\x0c\x68\x4c\x9f\x41\xb2\x4d\x3e\x5d\x1f\x7d\x62\x7d\x58\xe7\x0f\x0a\x9c\xd3\x95\xa5\x66\x78\x1a\x1e\xba\xbc\x92\xeb\xe8\x85\x0e\x74\x62\xe1\x60\x95\xb4\x50\x60\xc5\xc0\x9c\x00\x04\x17\xca\x50\xda\xf0\xa0\x5e\xc5\xc9\xa7\x2a\x3e\xa9\x91\xf8\x85\x4f\x31\xb1\x26\x37\x2a\x5b\xf6\xdc\x75\x58\x86\x3b\x7e\x6b\xf1\x2d\x02\xf0\xed\xbe\x15\x1e\x8e\x6f\xfb\xb2\xb1\xc7\xf3\x3f\xf0\xc2\x1e\x3b\xb4\x36\x3a\xb6\xbf\x47\xec\xc0\x59\xce\xa3\x61\xff\x8d\x12\x80\x17\x35\xca\x70\xdc\x9c\x86\x85\x6a\x02\x58\x7f\xfe\xb2\xce\xd8\xec\x82\x73\xfa\x0a\x00\x00\xff\xff\x3a\xfe\x93\xd9\x2d\x01\x00\x00")

func manifests00CustomResourceDefinitionYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests00CustomResourceDefinitionYaml,
		"manifests/00-custom-resource-definition.yaml",
	)
}

func manifests00CustomResourceDefinitionYaml() (*asset, error) {
	bytes, err := manifests00CustomResourceDefinitionYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/00-custom-resource-definition.yaml", size: 301, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests00NamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x04\xc0\xb1\x11\x82\x31\x08\x05\xe0\x3e\x53\xb0\x40\x0a\xdb\x0c\x61\x69\xff\xee\xcf\xf3\xe4\x34\xc0\x01\x3a\xbf\xdf\x5b\x6d\x2f\xb9\xe3\xb0\x02\x17\x07\x42\x1f\xcc\x52\xb7\x25\xbf\xdb\x38\x6c\x6c\x34\xd6\x10\x11\x31\x1c\x2e\xf1\xa0\xd5\x4b\x9f\x3d\xaf\xcf\xb7\x9a\x39\xb7\xd5\xf4\x60\xa2\x3d\xc7\x3f\x00\x00\xff\xff\x93\x5a\xd8\xf8\x52\x00\x00\x00")

func manifests00NamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests00NamespaceYaml,
		"manifests/00-namespace.yaml",
	)
}

func manifests00NamespaceYaml() (*asset, error) {
	bytes, err := manifests00NamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/00-namespace.yaml", size: 82, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests01ClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x8f\x31\x4f\xc3\x40\x0c\x85\xf7\xfb\x15\x56\x99\x13\xd4\x0d\xdd\x06\xec\x0c\x45\x62\x77\x2e\x2e\x31\x4d\xec\xc8\x76\x2a\xc1\xaf\x47\xa7\x6b\x27\xaa\x4e\xac\x77\xcf\xef\x7d\xdf\x03\xbc\xb0\x8c\x0e\x31\x11\xe8\x4a\x86\xa1\x06\x65\xde\x3c\xc8\xc0\x74\x26\x08\x05\x0e\x87\x77\xb2\x33\x17\x82\xe7\x52\x74\x93\xe8\xd3\x89\x65\xcc\xf0\xda\xa2\x07\x9d\xa9\x16\xb1\x7c\x26\x5c\xf9\x83\xcc\x59\x25\x83\x0d\x58\x7a\xdc\x62\x52\xe3\x1f\x0c\x56\xe9\x4f\x4f\xde\xb3\x3e\x9e\xf7\x03\x05\xee\xd3\x42\x81\x23\x06\xe6\x04\x20\xb8\x50\xbe\xae\x77\xa3\x78\x77\x45\x4a\xbe\x0d\x5f\x54\xc2\x73\xea\xa0\x2d\x5f\x80\x2e\x3c\xf7\xaf\xdb\xa7\xaf\x58\x28\x57\x4d\xf1\x89\x8f\xd1\xdd\xcc\x56\xe9\x03\x1d\x2b\xcf\x1f\xc5\xff\x5b\xd9\x9c\xec\xad\xa6\xab\xd0\xce\xbf\x3d\x68\xc9\xde\x94\xb0\x29\xe5\xfb\x15\xf9\xd6\xe3\x2e\xfd\x06\x00\x00\xff\xff\xc3\xf2\x9a\xab\xd1\x01\x00\x00")

func manifests01ClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests01ClusterRoleBindingYaml,
		"manifests/01-cluster-role-binding.yaml",
	)
}

func manifests01ClusterRoleBindingYaml() (*asset, error) {
	bytes, err := manifests01ClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/01-cluster-role-binding.yaml", size: 465, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests01ServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xb1\x0a\xc3\x30\x0c\x44\x77\x7f\xc5\x41\xe6\x04\xba\x7a\xeb\xd8\xb9\xd0\xdd\xd8\x0a\x11\x4d\x2c\x57\x92\xf3\xfd\x25\x21\xdd\x3a\xdf\xe3\xdd\x1b\x70\xcf\x59\x7a\x75\xcc\xa2\xf0\x85\x20\x8d\x34\xb9\x28\xd8\x8d\xd6\x79\xc2\xc3\x61\x8b\xf4\xb5\x40\xe9\xd3\x59\x09\x35\x6d\x64\x2d\x65\x82\x65\x69\x54\xc2\x80\x46\xba\xb1\x19\x4b\xb5\x29\xbc\xb9\x96\x88\x27\xe9\xce\x99\x2e\x7f\x48\x8d\x5f\xa4\x07\x11\xb1\xdf\xc2\x46\x9e\x4a\xf2\x14\x03\x4e\x5f\x44\x5e\xbb\x39\xe9\x58\xaa\x8d\xbf\x88\x6b\x3c\xcf\xe2\x91\x56\x6d\xe1\xd9\xc7\xbf\xec\x37\x00\x00\xff\xff\x83\x8f\x49\xa7\xcc\x00\x00\x00")

func manifests01ServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests01ServiceAccountYaml,
		"manifests/01-service-account.yaml",
	)
}

func manifests01ServiceAccountYaml() (*asset, error) {
	bytes, err := manifests01ServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/01-service-account.yaml", size: 204, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifests02DeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x52\xcb\x4e\xc3\x40\x0c\xbc\xe7\x2b\xac\xde\x43\xd5\xeb\xde\xa2\x52\x04\x12\xb4\x51\xa9\xe0\x88\xcc\xc6\x69\x57\xec\x4b\x6b\x27\x52\xff\x1e\x85\x86\x36\x85\xa0\xe0\xa3\x67\x3c\x9e\xf5\x2c\x46\xf3\x42\x89\x4d\xf0\x0a\x30\x46\x9e\xb7\x8b\xec\xc3\xf8\x4a\xc1\x2d\x45\x1b\x8e\x8e\xbc\x64\x8e\x04\x2b\x14\x54\x19\x80\x47\x47\x0a\xb4\x6d\x58\x28\xe5\x95\xe7\x3c\x44\x4a\x28\x21\xf5\x20\x47\xd4\xa4\x20\x44\xf2\x7c\x30\xb5\xe4\xa3\x5c\x8e\xa4\x3b\xb9\x44\xd1\x1a\x8d\xac\x60\x91\x01\x30\x59\xd2\x12\x52\x87\x00\x38\x14\x7d\x78\xc4\x77\xb2\x7c\x6a\x4c\x6c\x17\x72\xd1\xa2\x50\x3f\x3d\x30\xdd\x95\xbd\x12\x9a\x90\x02\xf8\x36\xd8\x95\x0e\x5e\xd0\x78\x4a\x83\xf1\x7c\x4a\xe0\x54\xc6\xe1\x7e\x78\x8d\x79\x48\x66\x6f\xfc\xe8\x51\x54\x67\x9e\x65\x30\xad\x83\x73\xe8\x2b\x35\x68\xe5\xff\xda\x58\x36\xd6\x96\xc1\x1a\x7d\x54\xf0\x50\xaf\x83\x94\x89\xb8\x8b\xf2\xc2\x23\xdf\x0e\x75\x2f\x2f\x7a\x2d\x76\xcb\xfb\xb7\x75\xf1\xb4\x7a\x2e\x8b\xe5\xea\x8a\x03\xd0\xa2\x6d\xe8\x2e\x05\xa7\x7e\x00\x00\xb5\x21\x5b\x6d\xa9\xfe\x8d\xf4\x58\x89\x72\x50\xe7\x60\x6e\xce\xbf\x65\xd4\xc6\xa6\x5c\x6d\x8b\xdd\x66\xfb\xe5\x64\xcc\x84\x82\xd9\xd8\x29\x66\x3d\x97\x29\xb5\x46\x53\xa1\x75\x68\xbc\xac\xff\x0e\xeb\x33\x00\x00\xff\xff\x0c\xd4\x5d\x1f\x03\x03\x00\x00")

func manifests02DeploymentYamlBytes() ([]byte, error) {
	return bindataRead(
		_manifests02DeploymentYaml,
		"manifests/02-deployment.yaml",
	)
}

func manifests02DeploymentYaml() (*asset, error) {
	bytes, err := manifests02DeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/02-deployment.yaml", size: 771, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _manifestsImageReferences = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8d\xb1\x0e\xc2\x30\x0c\x44\xf7\x7c\x85\x95\x3d\xad\x58\x33\xb3\x30\x23\xb1\x5b\xad\x1b\xac\x36\x71\x64\x1b\xbe\x1f\xb5\x45\x4c\x4c\x77\xba\x93\xde\x5b\xb9\xcd\x19\x6e\x15\x0b\xdd\x5d\x09\x6b\xc0\xce\x0f\x52\x63\x69\x19\x78\xdf\x07\xe9\xd4\xec\xc9\x8b\x0f\x2c\xe3\xfb\x12\xac\xd3\x94\x03\x80\x63\xb1\x3d\x13\x34\xac\x94\x61\xda\x5e\xe6\xa4\x69\x6e\x96\xa4\x93\xa2\x8b\x06\x00\x80\x45\xa5\x66\x38\x2a\xc0\x69\x8c\x57\x99\x56\xd2\x43\x1c\xbf\xcf\x49\x89\x3f\xdd\x28\xca\x85\x5b\xfa\xc7\xcd\x1b\x3a\x99\xc7\xf0\x09\x00\x00\xff\xff\x37\x7c\x86\x26\xc1\x00\x00\x00")

func manifestsImageReferencesBytes() ([]byte, error) {
	return bindataRead(
		_manifestsImageReferences,
		"manifests/image-references",
	)
}

func manifestsImageReferences() (*asset, error) {
	bytes, err := manifestsImageReferencesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "manifests/image-references", size: 193, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsDaemonsetYamlSwp = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x9a\xbf\x8f\x1b\x45\x14\xc7\xbf\xa1\xe3\xec\xf5\x22\xe8\xf8\x21\x0d\x4e\x41\x41\xd6\x6b\xe7\x74\xc7\xb1\x1d\xca\x89\x10\x29\x77\x58\xca\x25\x8a\x94\x50\xcc\xcd\x3e\x7b\x87\x9b\xdd\x59\xcd\xcc\x5a\x32\x05\x54\x14\x50\x51\x50\x22\x0a\xaa\x34\xd4\x48\x54\xf7\x07\x50\x21\x41\x85\xae\xa0\xa4\x42\x22\xa2\xa0\x42\x7b\xf6\xd9\xbb\x28\x46\xd7\x40\x14\x69\x3e\xcd\xde\xdc\x7b\xef\xfb\x7e\xf8\xd9\x92\xe5\x39\x1e\xde\xbb\x75\xc0\xf6\x06\x23\x00\x78\x01\xf8\xea\xe5\x6f\x1f\x3c\x7c\xe5\x3e\xdc\x11\x60\xf8\x4c\xe2\x12\xa4\x99\x28\xa3\xd1\x6e\x74\x7d\x7b\x38\xb0\x1f\x8a\x81\xa1\x34\xe3\x6e\x20\x74\xde\xf6\xfb\xb8\x16\x8c\xed\x34\x9e\xea\xd8\x1a\x11\x4f\xa5\xcb\xaa\xe3\xda\x2f\xd6\x25\x15\x36\x93\x13\x17\x0b\x55\x59\x47\x26\x4a\x0b\x1b\xe9\x92\x0c\x77\xda\xc4\xdc\x5a\x72\x36\x4e\x0b\x1b\xa7\x9c\x72\x5d\x58\x72\x83\x39\xcf\xd5\x65\xca\xf3\x78\x3c\xe7\x04\xd8\xbe\x3e\x1a\xd6\x7f\x5d\xed\xbf\xce\x5e\x7a\xf1\xee\xd3\x2e\xc8\xe3\xf1\x78\x3c\x1e\x8f\xc7\xe3\xf1\xfc\x8f\xb8\xf2\x0a\x3e\x01\xf0\xdc\xf2\xbc\xbb\x7c\x5e\xf9\xc7\xd3\xe3\xf1\x78\x3c\x1e\x8f\xc7\xe3\xf1\x78\x3c\x1e\xcf\xb3\x0b\x4f\x81\x5f\x9f\x07\x7e\xde\x5a\xfc\xfe\x7f\xf1\xfd\xff\xf7\x10\xf8\x25\x04\x7e\x0c\x81\xef\x42\xe0\xcb\x10\xf8\x3c\x04\x6c\x08\xe4\x21\xc0\x43\x60\x1c\x02\x6f\x85\xc0\x9b\x21\xf0\x5a\x08\x84\x21\xf0\xb8\x07\xfc\xd6\x03\x7e\xe8\x01\xdf\xf7\x80\x47\x3d\x40\xf5\x80\x77\x7b\xc0\x1b\x3d\xe0\x71\x00\x9c\x05\xc0\x69\x00\x7c\x1d\x00\x9f\x06\xc0\x2c\x00\xee\x07\xc0\x7e\x00\x5c\x0b\x80\x5e\x00\xfc\xd9\x05\xce\xba\xc0\x69\x17\x78\xd4\x05\xbe\xe8\x02\xf3\x2e\x30\xe9\x02\xb7\xbb\xc0\x4e\x17\x78\xb5\x0b\xfc\xd5\x01\xce\x3a\xc0\x69\x07\xf8\xa6\x03\x7c\xd6\x01\x3e\xea\x00\x0f\x3b\xc0\xcd\x0e\xd0\xef\x00\x61\x07\xf8\x63\x6b\xd1\xe3\x4f\x5b\x4f\x79\xe0\x1e\x8f\xc7\xe3\xf1\x78\x3c\x9e\xff\x00\xd6\xa2\xe4\x2e\x4b\xd8\x0d\x6d\x68\x22\x15\xb5\x8c\x11\x3b\xa1\xf9\x06\x9b\x74\x94\xdb\xa4\xf5\xaf\xab\xec\x90\xe7\xc4\xa4\x65\x96\x1c\xe3\x8e\x99\xaa\x70\x32\x6f\xc6\x09\x5d\x4c\xe4\xf4\x80\x97\xeb\xc8\x88\x15\x3c\xa7\x64\x69\x8a\x66\x5a\x55\xab\x90\xc5\x61\x95\x26\x2d\xec\x58\x2b\x29\xe6\x09\xdb\xa7\x09\xaf\x94\x6b\x48\x4f\xb8\x54\x95\xa1\xa3\xcc\x90\xcd\xb4\x4a\x13\xb6\xd3\xb0\xda\x4a\x08\xb2\xb6\x61\x1d\x35\xac\x75\x95\xba\x72\x77\x48\xe8\x22\xb5\xed\x48\x59\x48\x27\xb9\xda\x27\xc5\xe7\x2b\x87\xdd\x61\xab\x73\x2b\x32\xaa\x7b\x78\xef\xe8\x68\xdc\x32\x94\xda\xb8\x84\xed\x0d\xf7\xda\xfe\x8b\x99\xc7\x19\x71\xe5\xb2\x86\x25\x73\xae\xbc\x49\x6e\x3d\x1c\x25\x67\x54\x90\xb5\x63\xa3\x8f\xa9\x39\xed\xd2\x68\xa7\x85\x56\x09\x3b\xba\xd1\x4c\xb9\x98\x65\x4e\xce\x48\x61\x1b\x33\x16\xba\x70\x5c\x16\x64\xc6\xe7\x15\xbd\x3d\xda\xd9\xbe\xb4\x5a\x5a\xd8\xc8\x89\x72\xb3\xda\xce\xf6\x06\xb5\xbb\xfb\x4f\x54\xbb\xa4\x52\x3d\xbc\xd6\x8a\x19\xe2\xe9\xfb\x85\x9a\x27\xcc\x99\xaa\xb9\x56\xb9\xae\x0a\x37\x5e\x0c\x95\x9c\x88\x85\x36\xd4\xce\xb3\x79\xc7\x2e\xb6\xec\xa0\xd6\x68\xa4\xe3\x66\x6a\x13\xf6\x80\xf5\xa3\x3a\xaa\x7f\x8d\xf5\x9b\xd2\xf1\xc5\x9b\xa2\xcf\x3e\x58\x85\x08\x9d\xe7\xbc\x48\xcf\xa3\x96\x7e\x4d\xb3\xcc\xf9\x94\xc6\x95\x52\x17\x4b\x7c\x6b\x72\xa8\xdd\xd8\x90\xa5\xc2\xb5\xbd\x12\xd6\x4f\xb5\x38\x21\x33\x90\xba\x71\xfd\x5b\x1b\x39\x95\x45\xb4\xd4\x4e\x66\xdb\x83\xd1\x68\x30\xec\xa3\xdd\xe5\xba\xf3\xd5\x74\x57\x7d\x59\x32\x33\x29\xe8\x1d\x21\xea\x76\x0f\x5b\xfe\xb6\x24\xb1\xee\x7f\x95\x35\xe2\x65\xd9\x14\x55\xfc\x98\xd4\x52\x30\x27\xc7\x53\xee\x78\x7d\x72\x94\x97\x8a\xbb\xd5\x9a\x6e\x10\xc8\xb9\x13\xd9\xed\x95\x86\x25\x45\xc2\x69\x93\x60\x9d\xfe\x89\x91\xeb\xb4\x75\x97\xb6\xe4\x82\x92\x86\x67\xe3\x66\x3c\xfe\xe5\xa3\x68\x5d\x30\x2f\xe5\x3d\x32\x56\xea\x22\x61\xbc\x2c\x6d\x3c\x1b\xe1\x44\xd6\xaf\xde\xfe\xf9\x25\xfa\x3b\xe4\xf0\x77\x00\x00\x00\xff\xff\x26\x49\xe8\x21\x00\x30\x00\x00")

func assetsDnsDaemonsetYamlSwpBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsDaemonsetYamlSwp,
		"assets/dns/.daemonset.yaml.swp",
	)
}

func assetsDnsDaemonsetYamlSwp() (*asset, error) {
	bytes, err := assetsDnsDaemonsetYamlSwpBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/.daemonset.yaml.swp", size: 12288, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x31\x4b\x04\x41\x0c\x85\xfb\xfc\x8a\x70\xfd\xae\xd8\x49\x3a\xb5\xb7\x38\xc1\x3e\x37\x93\xf3\xe2\xed\x66\x96\x49\x66\x41\x7f\xbd\x0c\x83\x20\xa8\x60\x97\xe2\xbd\xf7\x7d\xb9\xaa\x65\xc2\xc7\xa5\x79\x48\x3d\x96\x45\x1e\xd4\xb2\xda\x2b\xf0\xa6\x2f\x52\x5d\x8b\x11\xd6\x13\xa7\x99\x5b\x5c\x4a\xd5\x0f\x0e\x2d\x36\x5f\xef\x7c\xd6\x72\xb3\xdf\xc2\x2a\xc1\x99\x83\x09\x10\x11\x8d\x57\x21\x4c\x63\x6f\xca\xe6\x94\xcd\xc1\xdb\xe9\x4d\x52\x38\xc1\x84\x83\xf8\x2c\x75\xd7\x24\xf7\x29\x95\x66\x01\x5f\xc5\x1e\x1e\xb7\x6f\x9c\x84\xb0\x6c\x62\x7e\xd1\x73\x4c\xdf\x36\xa1\x96\x45\x8e\x72\xee\xc8\x1f\x0f\xc0\x5f\x12\xff\xd8\x6d\x2e\xf5\xa9\x87\xba\xe8\xc1\xdf\x3d\x64\x25\x1f\xaa\x3c\x54\xe9\xd7\x66\x07\x1c\xe0\x33\x00\x00\xff\xff\xdb\x7c\x21\xe3\x4d\x01\x00\x00")

func assetsDnsClusterRoleBindingYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleBindingYaml,
		"assets/dns/cluster-role-binding.yaml",
	)
}

func assetsDnsClusterRoleBindingYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role-binding.yaml", size: 333, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\x8d\x3d\x6b\xc4\x30\x0c\x40\x77\xff\x0a\x71\x7b\x52\xba\x15\xaf\x1d\xba\x77\xe8\xae\xb3\x05\x27\xe2\x48\x46\x92\x53\xe8\xaf\x2f\x89\xb3\xbd\xf7\xd0\xc7\xc6\x52\x33\x7c\xb6\xe1\x41\xf6\xad\x8d\x12\x76\xfe\x21\x73\x56\xc9\x60\x4f\x2c\x2b\x8e\x78\xa9\xf1\x1f\x06\xab\xac\xdb\x87\xaf\xac\x6f\xc7\x7b\xda\x29\xb0\x62\x60\x4e\x00\x82\x3b\x65\x28\xf3\xcc\x52\xc5\x73\x15\x4f\x36\x1a\x79\x4e\x0b\x60\xe7\x2f\xd3\xd1\xfd\x9c\x5d\xe0\xf1\x48\x00\x46\xae\xc3\x0a\xdd\x8d\xa4\x76\x65\x09\xbf\xcc\xc9\x0e\x2e\x34\xa5\x6b\x9d\x70\x7e\xf1\x8e\xb3\x1f\x64\xcf\x7b\xb7\xb1\xc7\x05\xbf\x18\xe5\x95\xfe\x03\x00\x00\xff\xff\x8e\xf7\xdc\x36\xd4\x00\x00\x00")

func assetsDnsClusterRoleYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsClusterRoleYaml,
		"assets/dns/cluster-role.yaml",
	)
}

func assetsDnsClusterRoleYaml() (*asset, error) {
	bytes, err := assetsDnsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/cluster-role.yaml", size: 212, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsConfigmapYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8c\x4f\x4b\x03\x31\x10\xc5\xef\xf9\x14\x0f\x3c\x77\x6b\x59\x56\x30\xd7\x9e\xf5\xe8\x7d\x4c\x66\x9b\xd0\x6c\x12\x66\x92\xa2\xa8\xdf\x5d\x6a\x75\xb5\xe0\x3b\xbd\x7f\xfc\x8e\x31\x7b\x8b\x7d\xc9\x73\x3c\x3c\x50\x35\x54\xe3\x13\x8b\xc6\x92\x2d\x4e\x3b\xb3\x70\x23\x4f\x8d\xac\x01\x6e\xf0\x48\x0b\x23\x2a\x94\x1b\xa8\x41\x7a\x6e\x71\x61\x03\x64\x5a\x58\x2b\x39\xb6\x28\x95\xb3\x86\x38\xb7\x8d\x4b\x5d\x1b\xcb\xc6\x67\x35\x3f\x8c\x7d\x11\x9e\x63\x62\x8b\x77\x03\x00\x83\x9d\xc6\x69\xc4\xdb\x57\x38\x8b\x45\x8a\xe8\x1a\x03\x53\x6a\x61\x8d\xc7\xfe\xcc\x92\xb9\xb1\xe2\x9b\x3e\xa4\xe2\x28\x21\xe6\x0d\x79\x2f\x03\x49\x25\xc4\x7a\x77\x31\xbf\xd8\xb3\x6a\xf1\x8a\x98\x95\x5d\x17\xbe\x5a\x7a\xd5\x26\x4c\xcb\x55\x39\x53\x4a\x2d\x48\xe9\x87\xf0\x3f\x7e\x7d\x7f\xac\xae\x4a\x59\xb8\x05\xee\x0a\x7b\xbf\x9b\xc6\xbf\xc3\xcb\x2b\x06\x6c\xb9\xb9\xad\xb0\x96\x74\x1a\x5c\xc9\xf3\x7a\x70\xe4\x02\x63\xbc\x5d\x0b\xe1\x54\xc8\x9b\x0b\xff\x33\x00\x00\xff\xff\x6c\x54\x72\x1f\xa6\x01\x00\x00")

func assetsDnsConfigmapYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsConfigmapYaml,
		"assets/dns/configmap.yaml",
	)
}

func assetsDnsConfigmapYaml() (*asset, error) {
	bytes, err := assetsDnsConfigmapYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/configmap.yaml", size: 422, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsDaemonsetYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x4d\x6b\xdb\x4c\x10\xbe\xfb\x57\x0c\x7e\xaf\xaf\xe2\x98\x90\x92\xee\xad\xc4\xd0\x16\x9a\x54\x10\xb7\x97\xd2\xc3\x64\x35\xb2\x96\xec\xee\x2c\x3b\x23\x83\xff\x7d\x91\x6d\x29\x2b\xea\x06\xa2\x93\xd8\xe7\x63\x66\x1e\xcd\xea\xc5\xc5\xc6\xc0\x06\x29\x70\x7c\x22\x5d\x60\x72\x3f\x29\x8b\xe3\x68\x00\x53\x92\xd5\x7e\xbd\x08\xa4\xd8\xa0\xa2\x59\x00\xfc\x07\x8f\x18\x08\x9c\x80\x90\x02\x2a\xe4\x3e\xaa\x0b\xb4\x00\x88\x18\x48\x12\x5a\x32\xc0\x89\xa2\x74\xae\xd5\xca\xfa\x5e\x94\x72\xd5\x44\x59\x00\x78\x7c\x26\x2f\x83\x0f\x14\x1c\x4c\xc9\xc0\x40\x90\x44\x76\x00\x85\x3c\x59\xe5\x7c\x22\x06\x54\xdb\x7d\x2b\x94\x17\xb5\x00\x4a\x21\x79\x54\x3a\xab\x8a\xa6\x87\xc7\xcf\x0c\xfe\x61\x01\x30\xb6\x70\x7c\xa7\xbc\x77\x96\x3e\x59\xcb\x7d\xd4\x61\xee\x57\x1e\x80\xe5\xa8\xe8\x22\xe5\xc9\xb4\x3a\x46\x50\x72\x00\x5c\xc0\x1d\x19\x58\x36\x6c\x5f\x28\x5f\x39\x5e\x4d\x85\x57\x9c\xdd\xce\xc5\xca\x72\xa6\x26\x8a\xd9\xdf\x5c\xad\xd7\x57\xd7\xcb\xb9\xb6\xee\xbd\xaf\xd9\x3b\x7b\x30\xf0\xb5\x7d\x64\xad\x33\x09\x45\x9d\x58\x96\x43\xc0\xe1\x23\xfe\x82\xe5\xd9\x6a\x09\xbf\x27\x18\xf3\x4e\x8e\x58\x65\x39\xb6\xcb\xff\x61\xb9\x22\xb5\xab\x33\x73\x75\xcf\x99\x5a\xe7\xa9\x94\xec\xd9\xf7\x81\x1e\x86\xa1\x8b\xc0\xc6\xe9\x06\x1b\xb7\xab\x4e\xa4\x09\x05\x08\x03\xbf\x46\xed\x0c\x94\x15\x0a\x46\x26\x6c\xbe\x47\x7f\x30\xa0\xb9\x7f\x95\x26\xce\xf3\x3a\x53\xb2\x35\x67\x35\x70\x7b\x73\x7b\x53\xb8\xfc\x9d\x31\x40\xca\xac\x6c\xd9\x1b\xf8\xb1\xa9\xdf\xef\x54\xa9\x4d\x17\xdd\xb6\xf7\x6f\xb8\x7d\x5c\x5f\x70\x0b\xa4\xd9\xd9\xcb\xbd\x95\x6e\xde\xed\x29\x92\x48\x9d\xf9\x99\x4c\x41\xef\x54\xd3\x67\xd2\xf2\x08\x20\x9d\x62\xed\x08\xbd\x76\x73\xe4\xd8\xca\xdd\xf5\xdd\xf5\xec\x58\x6c\x47\x43\x3b\x5f\xb6\xdb\xba\x00\x5c\x74\xea\xd0\x6f\xc8\xe3\xe1\x89\x2c\xc7\x46\x0c\x7c\x28\xa5\xc3\x5d\xe6\x5e\x27\xf0\xb6\xc0\xa4\xb7\x96\x44\xb6\x5d\x26\xe9\xd8\x37\x06\xd6\x05\xda\xa2\xf3\x7d\xa6\x02\x1d\xb5\x4d\x94\x71\x83\x37\xd4\x62\xef\xc7\xe5\x3d\xed\xd0\x3b\x76\xec\x74\xfe\x80\x69\x1e\xcf\x1b\x3f\xa5\x62\x76\xa5\x20\x73\x5d\x05\x2f\x74\x30\x30\xde\x81\x19\x36\x86\x3e\x81\x7f\x02\x00\x00\xff\xff\x6f\x48\xea\xb4\x2a\x05\x00\x00")

func assetsDnsDaemonsetYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsDaemonsetYaml,
		"assets/dns/daemonset.yaml",
	)
}

func assetsDnsDaemonsetYaml() (*asset, error) {
	bytes, err := assetsDnsDaemonsetYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/daemonset.yaml", size: 1322, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\xca\xb1\xae\x42\x31\x08\x06\xe0\x9d\xa7\x20\x67\xef\xbd\x71\xe5\x21\x1c\xdd\x49\xfb\x1b\x89\xa7\xd0\x14\xf4\xf9\x8d\x93\xfb\xf7\x34\x1f\xc2\x57\x9d\xc8\xa5\x1d\xa4\xcb\x6e\xd8\x69\xe1\xc2\xef\x0b\x4d\x94\x0e\x2d\x15\x62\x76\x9d\x10\x8e\x05\xcf\x87\xdd\xab\xf5\xf3\x95\x85\xdd\x86\x27\x31\xab\x7b\x94\x96\x85\xe7\x17\xf3\x0f\xfe\x59\xfc\x7b\x0c\xb4\xc4\x89\x5e\xb1\x85\x8f\x83\x3e\x01\x00\x00\xff\xff\xb5\x9f\xce\xf1\x79\x00\x00\x00")

func assetsDnsNamespaceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsNamespaceYaml,
		"assets/dns/namespace.yaml",
	)
}

func assetsDnsNamespaceYaml() (*asset, error) {
	bytes, err := assetsDnsNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/namespace.yaml", size: 121, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xca\x31\x0e\x82\x21\x0c\x06\xd0\x9d\x53\xf4\x02\x0c\xae\xdd\x3c\x83\x89\x7b\x53\x3e\x63\xa3\x14\x42\x0b\xe7\x37\x26\xff\xf6\x86\xf7\x31\x6f\x4c\x0f\xac\x63\x8a\xbb\xea\xd8\x9e\x45\xa6\x3d\xb1\xc2\x86\x33\x9d\x5b\xe9\x48\x69\x92\xc2\x85\xc8\xa5\x83\xa9\x79\x5c\x8e\x29\x0a\xa6\x31\xe1\xf1\xb6\x57\x56\xfd\xee\x48\xac\xfa\x2f\xbf\x00\x00\x00\xff\xff\x35\xeb\xbe\x6a\x5d\x00\x00\x00")

func assetsDnsServiceAccountYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceAccountYaml,
		"assets/dns/service-account.yaml",
	)
}

func assetsDnsServiceAccountYaml() (*asset, error) {
	bytes, err := assetsDnsServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service-account.yaml", size: 93, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsDnsServiceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8e\xbd\x4e\x03\x41\x0c\x84\xfb\x7d\x8a\x91\xa8\x0f\x09\x21\x9a\x6d\xa1\xa1\x41\x27\xf1\xd3\x3b\x7b\x43\x58\xe1\xfd\xd1\xda\x09\xe2\xed\x51\x2e\x10\xa0\x40\x94\x1e\x7f\xfe\xc6\xaf\xb9\x2e\x11\xf7\x1c\xfb\x9c\x18\xa4\xe7\x27\x0e\xcb\xad\x46\xec\x2f\x42\xa1\xcb\x22\x2e\x31\x00\x67\xb8\x93\x42\x64\x83\xd1\x21\x8e\xb1\xab\x9e\x0b\x03\x50\xa5\xd0\xba\x24\x46\xb4\xce\x6a\x2f\xf9\xd9\xa7\xa4\x3b\x73\x8e\x69\xa9\x16\x00\x95\x0d\xd5\x0e\x1e\xfc\x60\xa4\xf7\x88\x03\x60\x9d\xe9\x58\xf2\x79\x76\x3b\xe3\x2d\xab\x62\x43\xc8\xce\x5b\x11\xcf\x49\x54\xdf\x51\xa4\xca\x96\xcb\x79\x00\x8c\xca\xe4\x6d\xfc\x69\x05\x7a\x1b\xbe\xb6\x4e\xeb\x93\x5f\xf1\x71\x11\x71\x75\xb9\x0e\x2e\x63\x4b\x9f\xd7\xe8\x04\x8c\xe6\x2d\x35\x8d\x78\xbc\x99\x7f\x0b\x26\x4f\xfd\x5f\xc9\x37\x74\x12\x3d\x5c\xcf\xe1\x23\x00\x00\xff\xff\xd5\x5c\x70\x51\x6f\x01\x00\x00")

func assetsDnsServiceYamlBytes() ([]byte, error) {
	return bindataRead(
		_assetsDnsServiceYaml,
		"assets/dns/service.yaml",
	)
}

func assetsDnsServiceYaml() (*asset, error) {
	bytes, err := assetsDnsServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/dns/service.yaml", size: 367, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"manifests/00-cluster-role.yaml": manifests00ClusterRoleYaml,
	"manifests/00-custom-resource-definition.yaml": manifests00CustomResourceDefinitionYaml,
	"manifests/00-namespace.yaml": manifests00NamespaceYaml,
	"manifests/01-cluster-role-binding.yaml": manifests01ClusterRoleBindingYaml,
	"manifests/01-service-account.yaml": manifests01ServiceAccountYaml,
	"manifests/02-deployment.yaml": manifests02DeploymentYaml,
	"manifests/image-references": manifestsImageReferences,
	"assets/dns/.daemonset.yaml.swp": assetsDnsDaemonsetYamlSwp,
	"assets/dns/cluster-role-binding.yaml": assetsDnsClusterRoleBindingYaml,
	"assets/dns/cluster-role.yaml": assetsDnsClusterRoleYaml,
	"assets/dns/configmap.yaml": assetsDnsConfigmapYaml,
	"assets/dns/daemonset.yaml": assetsDnsDaemonsetYaml,
	"assets/dns/namespace.yaml": assetsDnsNamespaceYaml,
	"assets/dns/service-account.yaml": assetsDnsServiceAccountYaml,
	"assets/dns/service.yaml": assetsDnsServiceYaml,
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
	"assets": &bintree{nil, map[string]*bintree{
		"dns": &bintree{nil, map[string]*bintree{
			".daemonset.yaml.swp": &bintree{assetsDnsDaemonsetYamlSwp, map[string]*bintree{}},
			"cluster-role-binding.yaml": &bintree{assetsDnsClusterRoleBindingYaml, map[string]*bintree{}},
			"cluster-role.yaml": &bintree{assetsDnsClusterRoleYaml, map[string]*bintree{}},
			"configmap.yaml": &bintree{assetsDnsConfigmapYaml, map[string]*bintree{}},
			"daemonset.yaml": &bintree{assetsDnsDaemonsetYaml, map[string]*bintree{}},
			"namespace.yaml": &bintree{assetsDnsNamespaceYaml, map[string]*bintree{}},
			"service-account.yaml": &bintree{assetsDnsServiceAccountYaml, map[string]*bintree{}},
			"service.yaml": &bintree{assetsDnsServiceYaml, map[string]*bintree{}},
		}},
	}},
	"manifests": &bintree{nil, map[string]*bintree{
		"00-cluster-role.yaml": &bintree{manifests00ClusterRoleYaml, map[string]*bintree{}},
		"00-custom-resource-definition.yaml": &bintree{manifests00CustomResourceDefinitionYaml, map[string]*bintree{}},
		"00-namespace.yaml": &bintree{manifests00NamespaceYaml, map[string]*bintree{}},
		"01-cluster-role-binding.yaml": &bintree{manifests01ClusterRoleBindingYaml, map[string]*bintree{}},
		"01-service-account.yaml": &bintree{manifests01ServiceAccountYaml, map[string]*bintree{}},
		"02-deployment.yaml": &bintree{manifests02DeploymentYaml, map[string]*bintree{}},
		"image-references": &bintree{manifestsImageReferences, map[string]*bintree{}},
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

