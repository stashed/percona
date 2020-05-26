// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// installer.stash.appscode.com_stashperconaxtradbs.v1.yaml
// installer.stash.appscode.com_stashperconaxtradbs.yaml
package crds

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _installerStashAppscodeCom_stashperconaxtradbsV1Yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x4b\x8f\x1b\x37\x0c\xbe\xfb\x57\x10\xe8\x21\x97\x7a\x8c\x6d\x2e\xc5\xdc\xd2\x4d\x0f\x41\x5f\xc1\x3a\x08\x7a\xa5\x35\xf4\x98\x5d\x8d\xa4\x92\x94\x91\xed\xaf\x2f\xa4\x19\x3f\xd6\xaf\x6c\xb7\x88\x4e\x23\xbe\xf5\x7d\x24\x07\x13\x7f\x26\x51\x8e\xa1\x05\x4c\x4c\x5f\x8c\x42\xb9\x69\xf3\xf8\xa3\x36\x1c\x17\xdb\xbb\xd9\x23\x87\xae\x85\xfb\xac\x16\x87\x07\xd2\x98\xc5\xd1\x7b\x5a\x73\x60\xe3\x18\x66\x03\x19\x76\x68\xd8\xce\x00\x9c\x10\x16\xe1\x27\x1e\x48\x0d\x87\xd4\x42\xc8\xde\xcf\x00\x3c\xae\xc8\x6b\xb1\x01\xc0\x94\x5a\x50\x43\xdd\xcc\x00\x02\x0e\x34\xdd\x12\x89\x8b\x01\xbf\x98\x60\xb7\xd2\x86\x83\x1a\x7a\x4f\xd2\x54\x6d\x83\x29\xa9\x8b\x1d\x35\x2e\x0e\x33\x4d\xe4\x4a\xb4\x5e\x62\x4e\x2d\xdc\xb4\x1d\x93\x4c\xc9\x1d\x1a\xf5\x51\x78\x77\x9f\xef\x2b\x29\xdf\x3b\xbf\x7a\x1d\x1f\xbe\x2c\xea\x8f\x63\x69\x7f\x9a\xe0\xfb\x9f\xaa\xd2\xb3\xda\x2f\x57\x0c\x7e\x65\xb5\x6a\x94\x7c\x16\xf4\x17\x9f\x57\xf5\xca\xa1\xcf\x1e\xe5\x92\xc5\x0c\x40\x5d\x4c\xd4\xc2\xef\xa5\xfa\x84\x8e\xba\x19\xc0\x76\xa4\xab\x56\x3f\x9f\xd0\xdb\xde\xa1\x4f\x1b\xbc\x1b\x63\xba\x0d\x0d\x38\x3e\x0e\x20\x26\x0a\xef\x3e\x7e\xf8\xfc\x76\xf9\x4c\x0c\x90\x24\x26\x12\xdb\xe3\x30\x9e\xa3\x7e\x38\x92\x02\x74\xa4\x4e\x38\x59\x6d\x94\x37\x25\xe0\x68\x05\x5d\x69\x04\x52\xb0\x0d\xed\x4a\xa3\x6e\xaa\x01\xe2\x1a\x6c\xc3\x0a\x42\x49\x48\x29\x58\x6d\x8e\x67\x81\xa1\x18\x61\x80\xb8\xfa\x8b\x9c\x35\xb0\x24\x29\x61\x40\x37\x31\xfb\x0e\x5c\x0c\x5b\x12\x03\x21\x17\xfb\xc0\xff\xec\x63\x2b\x58\xac\x49\x3d\x1a\x4d\x68\x1f\x0e\x07\x23\x09\xe8\x61\x8b\x3e\xd3\xf7\x80\xa1\x83\x01\x9f\x40\xa8\x64\x81\x1c\x8e\xe2\x55\x13\x6d\xe0\xb7\x28\x04\x1c\xd6\xb1\x85\x8d\x59\xd2\x76\xb1\xe8\xd9\x76\x73\xe0\xe2\x30\xe4\xc0\xf6\xb4\x70\x31\x98\xf0\x2a\x5b\x14\x5d\x74\xb4\x25\xbf\x50\xee\xe7\x28\x6e\xc3\x46\xce\xb2\xd0\x02\x13\xcf\x6b\xe9\xc1\xea\x30\x0d\xdd\x77\x32\x4d\x8e\xbe\x79\x56\xab\x3d\xa5\xda\xfe\xc2\xa1\x3f\x52\xd4\xd6\xbb\xc1\x40\xe9\x3c\x60\x05\x9c\x5c\xc7\x57\x1c\x80\x2e\xa2\x82\xce\xc3\xcf\xcb\x4f\xb0\x4b\x5d\xc9\x38\x45\xbf\xe2\x7e\x70\xd4\x03\x05\x05\x30\x0e\x6b\x92\x91\xc4\xb5\xc4\xa1\xc6\xa4\xd0\xa5\xc8\xc1\xea\xc5\x79\xa6\x70\x0a\xbf\xe6\xd5\xc0\x56\x78\xff\x3b\x93\x5a\xe1\xaa\x81\x7b\x0c\x21\x1a\xac\x08\x72\xea\xd0\xa8\x6b\xe0\x43\x80\x7b\x1c\xc8\xdf\xa3\xd2\x37\x27\xa0\x20\xad\xf3\x02\xec\xcb\x28\x38\xde\x6b\xa7\xc6\x23\x6a\x47\x8a\xdd\x3a\xba\xc2\xd7\xf9\x8a\x58\x26\x72\x85\xc0\x82\xe1\x34\x2d\xeb\x28\xa3\x21\x4c\x96\x70\xb4\x6e\x0e\x67\xec\x57\x58\xb3\xa7\x67\x9a\xcb\x03\x5d\xce\x0a\xdd\x63\x4e\xa7\xd2\x5b\x1e\xe5\xa0\xf4\x17\xe5\x57\x01\xdb\x63\x11\x1d\xda\x03\x99\x3c\x5d\x76\x5f\x47\x19\xd0\xca\xce\xb6\xb7\x3f\xdc\x48\x50\x86\xb8\x27\x39\xb3\xb8\x42\x41\x0d\x9d\xbd\x2f\x2b\xf1\x8f\x2d\x89\x70\x47\xe7\x05\xdc\xa8\x9d\x07\xec\x2f\x78\xdc\x46\x49\xa8\x67\xbd\xfa\xd4\xaf\x20\x25\x94\xa2\xb2\xc5\x57\xba\x1b\xf6\xaf\xf0\x2b\x43\xc9\x42\xdd\xb9\xeb\x7c\xff\x9a\x8b\xaa\x5d\xad\x17\x94\x86\xe7\x79\x6e\xd0\xf4\x6a\x8a\x84\xd4\xa2\xfc\x67\x92\x5e\xdd\xca\x86\xd2\x93\xbd\x4b\xe9\x81\x92\x67\x87\x57\xa2\xfc\xdf\x8e\xbe\x45\xc9\x59\x09\x2f\x07\xfa\x72\xd8\xf9\xb4\x0f\x4e\x84\xb5\xfb\x4f\x64\x13\xdc\x5f\x5f\x80\x67\x42\x2d\xff\x90\xae\x05\x93\x3c\xba\x97\x40\x65\xbc\x46\xc9\xbf\x01\x00\x00\xff\xff\xcb\xb5\x2b\x1c\x77\x0a\x00\x00")

func installerStashAppscodeCom_stashperconaxtradbsV1YamlBytes() ([]byte, error) {
	return bindataRead(
		_installerStashAppscodeCom_stashperconaxtradbsV1Yaml,
		"installer.stash.appscode.com_stashperconaxtradbs.v1.yaml",
	)
}

func installerStashAppscodeCom_stashperconaxtradbsV1Yaml() (*asset, error) {
	bytes, err := installerStashAppscodeCom_stashperconaxtradbsV1YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "installer.stash.appscode.com_stashperconaxtradbs.v1.yaml", size: 2679, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _installerStashAppscodeCom_stashperconaxtradbsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x56\x5f\x6f\x1b\x45\x10\x7f\xf7\xa7\x18\x09\xa4\x02\xe2\xce\x84\x48\x08\xee\x05\x95\x14\xa4\x8a\x02\x55\x52\x2a\xa4\xa8\x48\xe3\xdd\xf1\x79\xc9\xde\xee\x32\x33\x67\x92\x7c\x7a\xb4\x7b\x77\x8e\xe3\x9c\x43\x1b\xd4\x7d\xf2\xfe\xe6\xff\x6f\x66\xd6\x87\xc9\xbd\x25\x16\x17\x43\x03\x98\x1c\x5d\x2b\x85\x7c\x93\xfa\xea\x5b\xa9\x5d\x5c\x6e\x4f\x56\xa4\x78\xb2\xb8\x72\xc1\x36\x70\xd6\x8b\xc6\xee\x9c\x24\xf6\x6c\xe8\x05\xad\x5d\x70\xea\x62\x58\x74\xa4\x68\x51\xb1\x59\x00\x18\x26\xcc\xe0\x1b\xd7\x91\x28\x76\xa9\x81\xd0\x7b\xbf\x00\xf0\xb8\x22\x2f\x59\x07\x00\x53\x6a\x40\x14\x65\xb3\x00\x08\xd8\xd1\x78\x4b\xc4\x26\x06\xbc\x56\x46\xbb\x92\xda\x05\x51\xf4\x9e\xb8\x2e\xd2\x1a\x53\x12\x13\x2d\xd5\x26\x76\x0b\x49\x64\xb2\xb7\x96\x63\x9f\x1a\x78\x54\x77\x08\x32\x06\x37\xa8\xd4\x46\x76\xd3\xbd\xda\x65\x92\x7f\x4f\x76\xe5\x3a\x14\x7e\x91\xc5\xaf\x87\xd4\xfe\x50\xc6\x17\x3f\x14\xa1\x77\xa2\x3f\x1f\x51\x78\xe5\x44\x8b\x52\xf2\x3d\xa3\x9f\x2d\xaf\xc8\xc5\x85\xb6\xf7\xc8\x73\x1a\x0b\x80\xc4\x24\xc4\x5b\xfa\x3d\x5c\x85\xf8\x4f\xf8\xc9\x91\xb7\xd2\xc0\x1a\xbd\xe4\x0c\xc5\xc4\x44\x0d\xfc\x9a\x8b\x4b\x68\xc8\x2e\x00\xb6\xe8\x9d\x2d\x3d\x18\xca\x8b\x89\xc2\xf3\xd7\x2f\xdf\x9e\x5e\x98\x0d\x75\x38\x80\xd9\x73\x4c\xc4\xba\x63\x61\x68\xcb\x6e\x20\x76\x18\x80\x25\x31\xec\x52\xf1\x08\xcf\xb2\xab\x41\x07\x6c\x1e\x01\x12\xd0\x0d\xc1\x76\xc0\xc8\x82\x94\x30\x10\xd7\xa0\x1b\x27\xc0\x54\x6a\x08\x5a\x52\xda\x73\x0b\x59\x05\x03\xc4\xd5\x5f\x64\xb4\x86\x8b\x5c\x27\x0b\xc8\x26\xf6\xde\x82\x89\x61\x4b\xac\xc0\x64\x62\x1b\xdc\xed\xce\xb3\x80\xc6\x12\xd2\xa3\xd2\xc8\xf2\x74\x5c\x50\xe2\x80\x3e\x93\xd0\xd3\x97\x80\xc1\x42\x87\x37\xc0\x94\x63\x40\x1f\xf6\xbc\x15\x15\xa9\xe1\x97\xc8\x04\x2e\xac\x63\x03\x1b\xd5\x24\xcd\x72\xd9\x3a\x9d\x56\xc0\xc4\xae\xeb\x83\xd3\x9b\xa5\x89\x41\xd9\xad\x7a\x8d\x2c\x4b\x4b\x5b\xf2\x4b\x71\x6d\x85\x6c\x36\x4e\xc9\x68\xcf\xb4\xc4\xe4\xaa\x92\x78\xd0\xb2\x47\x9d\xfd\x84\xc7\x7d\x91\x67\x7b\x99\xea\x4d\x2a\x23\xcf\x2e\xb4\x3b\xb8\x0c\xdb\x51\xde\xf3\xa4\x81\x13\xc0\xd1\x6c\xc8\xff\x8e\xde\x0c\x65\x56\xce\x7f\xbc\x78\x03\x53\xd0\xd2\x82\xfb\x9c\x17\xb6\xef\xcc\xe4\x8e\xf8\x4c\x94\x0b\x6b\xe2\xa1\x71\x6b\x8e\x5d\xf1\x48\xc1\xa6\xe8\x82\x96\x8b\xf1\x8e\xc2\x7d\xd2\xa5\x5f\x75\x4e\x73\xa7\xff\xee\x49\x34\xf7\xa7\x86\x33\x0c\x21\x2a\xac\x08\xfa\x64\x51\xc9\xd6\xf0\x32\xc0\x19\x76\xe4\xcf\x50\xe8\xa3\xd3\x9e\x19\x96\x2a\x53\xfa\xdf\xc4\xef\xbf\x5f\xd3\x99\x5b\x8f\x7c\xca\x63\x75\x0f\x01\xe8\xf0\xfa\x15\x85\x56\x37\x0d\x7c\x73\x7a\x20\x4b\xa8\x79\x24\x1b\xf8\xf3\x12\xab\xdb\x77\x9f\x5d\x56\x58\xdd\x7e\x55\x7d\xf7\xee\x8b\xcb\xf1\xc7\xe7\xdf\x7f\x7a\x60\x33\x9b\xe4\x04\x0f\x0d\xdc\xc1\xd3\x2b\x38\x3b\x34\x0f\xdf\xa5\x8b\x44\x26\x4f\x51\x6e\xe5\xb8\xa8\xeb\xc8\x83\x22\x8c\x9a\xb0\xf7\xc6\x4d\x67\x58\x16\x58\x3b\x4f\xef\x41\xd2\x0a\xcd\x55\x9f\x0e\x69\x3a\xa6\x9d\x0f\x72\x3b\x83\x1e\xa5\x62\xac\x3d\x1a\xd4\x73\x52\xbe\x99\x33\x5d\x47\xee\x50\xf3\x1f\x83\x9e\x7e\x7d\xd4\x75\x7e\x31\x5a\xe2\xd9\x0e\x1c\x50\x5d\x9c\xf6\xde\xe7\x11\xf8\x6d\x4b\xcc\xce\x3e\x18\x85\xa3\xf9\xba\x0e\xdb\x07\xda\x8f\x31\xc2\xd4\x3a\x39\x52\xda\xa3\xac\x30\xa5\x28\x4e\xe3\x13\x4c\x15\xdb\x0f\xb4\xc9\x3b\xef\x98\xec\xa1\x59\xb5\xcb\x7f\x46\x30\xe5\xf7\x40\xa4\x78\xe8\xff\x68\x23\x9e\xd4\x04\x26\xd1\xc8\x1f\xd4\x86\x27\x0d\xa6\x22\xb7\xa4\xcf\x53\x3a\xa7\xe4\x9d\xc1\x59\x0f\xff\x67\x3e\x8f\xd3\xfe\x20\xf4\xfb\x11\x3a\xe7\xb0\x1a\xf7\xf8\x1e\x54\xe6\xf8\x1e\x32\x92\xfa\xf8\x43\x75\x00\x6d\xa7\xcf\xce\xed\x09\xfa\xb4\xc1\x93\x3b\xac\x70\x55\x8d\x1f\x85\x7b\x62\x80\xf2\x21\x64\x1b\x50\xee\x87\x68\x39\x6e\xde\xaa\x01\xf9\x37\x00\x00\xff\xff\x5f\xa5\xf7\x30\xce\x0a\x00\x00")

func installerStashAppscodeCom_stashperconaxtradbsYamlBytes() ([]byte, error) {
	return bindataRead(
		_installerStashAppscodeCom_stashperconaxtradbsYaml,
		"installer.stash.appscode.com_stashperconaxtradbs.yaml",
	)
}

func installerStashAppscodeCom_stashperconaxtradbsYaml() (*asset, error) {
	bytes, err := installerStashAppscodeCom_stashperconaxtradbsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "installer.stash.appscode.com_stashperconaxtradbs.yaml", size: 2766, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"installer.stash.appscode.com_stashperconaxtradbs.v1.yaml": installerStashAppscodeCom_stashperconaxtradbsV1Yaml,
	"installer.stash.appscode.com_stashperconaxtradbs.yaml":    installerStashAppscodeCom_stashperconaxtradbsYaml,
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
	"installer.stash.appscode.com_stashperconaxtradbs.v1.yaml": &bintree{installerStashAppscodeCom_stashperconaxtradbsV1Yaml, map[string]*bintree{}},
	"installer.stash.appscode.com_stashperconaxtradbs.yaml":    &bintree{installerStashAppscodeCom_stashperconaxtradbsYaml, map[string]*bintree{}},
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
