// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../gmailctl.libsonnet
// ../../../default-config.jsonnet
package cmd

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

var _GmailctlLibsonnet = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x93\xcb\x8e\xdb\x3c\x0c\x85\xf7\x7e\x8a\x83\x59\x4d\x30\xfe\xe3\xbf\xdb\x0c\xa6\x40\x81\xb6\xcb\xae\xba\x0b\xb2\x50\x6c\xda\x26\xa0\x48\x01\x49\x4f\x26\x28\xe6\xdd\x0b\x5d\x9c\xb9\x00\xdd\x18\x96\xc8\x73\x48\x7e\x92\xba\x0e\xd3\xc9\xb1\xef\xcd\x43\xcd\x85\xc1\xc9\x00\xcf\x47\x71\x72\x6d\xba\xae\xe9\x3a\xfc\x9e\x59\xd1\xc7\x60\x8e\x83\x62\x31\xf6\x6c\x4c\x0a\x9b\x9d\x41\xf9\x74\xf6\x3c\x5e\xd1\xc7\xd3\x29\x06\x8c\xec\x8d\x04\xf1\x4c\xe2\x8c\x63\xd0\x6d\x93\x3c\xfa\xd9\x71\xf8\x99\x63\x0a\x56\x38\x8c\x4b\xe8\x53\x42\xb6\x69\x31\xf1\x33\x05\x38\x78\x56\x43\x1c\x21\x8b\x27\x6d\x93\x54\xc8\x16\x09\x49\x12\xe8\x52\xe2\x97\x99\x84\x60\x33\x95\x34\x38\xa1\x52\x81\x06\x58\x9c\x66\x32\x92\xac\xbd\xcc\xdc\xcf\x38\x91\x0b\xb5\x5d\x9b\xe9\x8a\xde\x05\x1c\x09\x1c\x8c\xe4\x2c\x64\x34\xc0\x25\xff\x6c\x91\x64\x71\xc4\x1d\x8f\x20\xaf\xeb\xf7\x6e\x5b\x48\x10\x84\x74\xf1\x96\x66\x38\x3a\xe5\xde\x79\x7f\xfd\xdc\x76\xed\x8f\x5c\x3f\x57\x1e\x49\xcc\x8a\x53\x1c\x78\x64\x1a\x70\xbc\xc2\x0d\x03\x87\x09\x2e\xe0\xdb\xaf\xef\xb8\xb0\xcd\x79\xa0\x40\x53\xe6\x96\xcc\x9c\xf7\x69\x2f\x89\xcf\x42\xcf\x1c\x17\xad\x7e\xba\x6d\x7c\xec\x9d\xff\xc0\xf5\x7e\xd4\x0d\x9e\x1a\xa0\xeb\xea\x29\x5d\xcb\xd0\x2b\xdb\xd4\x1c\xd8\x6e\x44\xd9\xb4\xd4\xa3\xa1\xfa\x6e\x1b\xa0\x18\x97\xfd\x7b\xd9\xe0\x09\x7f\x42\xb4\x1d\x64\x5b\x72\x5e\x1f\x4b\x09\xa1\x7e\x11\xe5\x67\xaa\x45\x62\xba\x12\x12\x97\x69\xce\x8d\x93\xa7\x13\x05\xd3\x3c\x88\xc8\xcd\xd8\x2d\x2f\xf7\x4e\xa4\x05\xb7\x6b\xf1\x16\xb2\x84\xc0\x61\x2a\xed\x03\x3c\x82\xf1\xf5\x09\x6a\xc3\xd6\x53\x98\x6c\x4e\x92\x4d\x82\x11\x72\x02\x56\x45\x5e\x91\x57\xaa\xdb\x5d\x57\x29\x5e\xea\xb4\xba\x12\x8e\xe3\xee\x2d\xe7\xbf\x7f\xb1\x7e\x03\x9d\x4f\xf2\xb3\xa2\x5f\x44\x28\x58\x0e\xd6\xd8\x8a\xeb\x22\x89\x54\xdd\x44\xe5\xb9\x7b\xb7\x03\xb8\x30\xec\x6e\xc0\x1f\xb0\x77\x22\x7b\x3e\x54\xac\x87\x5b\xe2\x6b\x7b\xfb\x75\xf9\x89\xe8\x0e\x35\xb5\xae\xd7\x84\x7c\x14\x39\xef\x06\x15\x0f\xf8\xd2\xbe\x2f\x52\x0f\xb2\x18\x6c\x0e\x37\xd6\x25\x76\x91\xc3\x06\xe6\xd8\xab\x09\xf7\xf6\xd8\x34\xc5\x6c\xd4\x36\xf9\xac\xea\x51\xf7\xff\x67\xf1\x3e\xff\x1d\x36\x8f\xf9\x55\xff\x78\x39\x47\xc9\xb7\xa7\x3e\x66\x6d\xd2\xbc\xef\x2f\xe5\xee\xc3\xaa\x6d\x5e\x9b\xbf\x01\x00\x00\xff\xff\x43\x03\x21\x1e\x72\x04\x00\x00")

func GmailctlLibsonnetBytes() ([]byte, error) {
	return bindataRead(
		_GmailctlLibsonnet,
		"../../../gmailctl.libsonnet",
	)
}

func GmailctlLibsonnet() (*asset, error) {
	bytes, err := GmailctlLibsonnetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../../gmailctl.libsonnet", size: 1138, mode: os.FileMode(420), modTime: time.Unix(1549574432, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _DefaultConfigJsonnet = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\x41\x6b\xdb\x4e\x10\xc5\xef\xfa\x14\xef\xcf\xff\xa0\x4b\xb0\x68\x8f\x0a\x81\x04\x62\x68\xa0\x8e\x43\xea\x1c\x4a\xc9\x61\x24\x8f\xac\xa5\xab\x1d\x31\x3b\x72\x63\x82\xbf\x7b\x59\xad\x93\x34\x21\xa7\xf5\xec\x3e\xff\xde\xbc\xa7\xaa\xc2\xed\x7a\xb3\xac\xb1\xe9\x5d\x84\x8b\x20\x44\x37\x8c\x9e\xc1\x4f\x94\xce\x45\x51\x55\xb8\xf3\x4c\x91\xa1\xdc\xb1\xc2\x04\xbd\xd9\x18\xeb\xaa\xda\x39\xeb\xa7\x66\xd1\xca\x50\x0d\x8d\x5a\xb5\x1b\xc8\xf9\xd6\xfc\xff\xad\x84\xce\xed\x26\x25\x73\x12\xd0\x89\x62\x2b\x6d\x04\x35\x32\x59\x02\x5a\xcf\xc8\x9a\xf4\x38\x90\x2d\x70\x2d\xa1\xb4\x34\xed\xd8\x92\x47\xdb\x53\xd8\xf1\x3f\xca\x17\x5a\xc3\x9d\x28\x27\x09\x8d\xa3\x3f\xc0\x65\xa2\xe0\x20\x93\x42\xfe\x04\xb8\xd0\xc8\xd3\x7f\x45\xba\xbe\x19\x46\x51\x9b\x29\xd1\x28\x6c\x49\xb7\xf0\xae\x51\xd2\x43\xe1\xa5\x25\x9f\x26\x5c\xc0\x65\x5d\xf9\x92\x60\xe1\x5d\x13\x25\x04\xb6\xf2\x7c\x06\xfd\x90\x81\x31\x45\xee\x26\x8f\x3d\xa9\xa3\xc6\x73\x84\x04\x98\x8c\xe9\x7d\xb3\xbe\x5e\xd7\xb8\x9b\x2c\xaf\xc1\x89\x83\x9e\x95\x4f\x36\x03\xe3\x02\xe5\xcf\xf5\xc3\xfd\x62\xb9\xba\xba\xf9\x7e\x39\x3b\xa5\xea\xca\xf3\x93\xc4\x64\x95\x44\xcf\x26\x35\x06\x3e\x9e\x17\xb3\xf1\xa6\x67\x50\x6b\x13\xf9\xf7\x3d\x14\xcf\x05\x50\x55\x58\xa5\x54\x26\x7a\x40\xcf\xb4\x65\x2d\x80\x3d\x6b\x74\x12\x6a\x94\xfb\x2f\xe4\xc7\x9e\xbe\x96\x67\x05\x40\x93\xf5\xa2\x35\xd2\x1f\x81\x40\x03\xd7\x79\x23\xdc\x5e\xad\x96\xf8\xb6\xbc\x5f\xce\x3a\xe4\xed\xd3\x12\x69\x3c\x9e\x15\xd9\x29\x47\x7c\x88\xfc\xd6\xb4\x4e\xa9\x85\x39\x26\xf2\x50\xe3\xd7\x8c\xc8\x26\x40\xe7\xbc\xb1\xd6\x73\xb8\xb3\xd3\x1d\xb5\x29\x40\xac\x5f\x45\xc0\x40\xfa\x3b\x7f\x2b\x0a\x56\xc3\x74\x7a\x55\x1f\xf3\x8f\xd3\xf1\x91\xfb\x86\xe8\x54\x86\x1a\x65\x43\x7a\x79\xa0\x5e\x64\xee\xf6\x03\xe4\x53\x6f\xd2\xb6\x77\x7b\x7e\xef\x0a\x78\x6a\xd8\xa7\x3c\x65\x27\x52\x3e\x7e\xb6\xce\x63\x71\x2c\xfe\x06\x00\x00\xff\xff\x4d\x13\xcd\x30\x42\x03\x00\x00")

func DefaultConfigJsonnetBytes() ([]byte, error) {
	return bindataRead(
		_DefaultConfigJsonnet,
		"../../../default-config.jsonnet",
	)
}

func DefaultConfigJsonnet() (*asset, error) {
	bytes, err := DefaultConfigJsonnetBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../../../default-config.jsonnet", size: 834, mode: os.FileMode(420), modTime: time.Unix(1549577715, 0)}
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
	"../../../gmailctl.libsonnet": GmailctlLibsonnet,
	"../../../default-config.jsonnet": DefaultConfigJsonnet,
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
	"..": &bintree{nil, map[string]*bintree{
		"..": &bintree{nil, map[string]*bintree{
			"..": &bintree{nil, map[string]*bintree{
				"default-config.jsonnet": &bintree{DefaultConfigJsonnet, map[string]*bintree{}},
				"gmailctl.libsonnet": &bintree{GmailctlLibsonnet, map[string]*bintree{}},
			}},
		}},
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

