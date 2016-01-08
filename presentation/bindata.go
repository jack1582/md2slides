// Code generated by go-bindata.
// sources:
// templates/remark.tmpl
// templates/reveal.tmpl
// templates/shower.tmpl
// DO NOT EDIT!

package presentation

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

var _templatesRemarkTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x93\xdf\x6f\xd3\x30\x10\xc7\xdf\xf7\x57\xdc\xfa\xd2\xb1\xb6\xf1\xf8\x21\x81\xda\x64\x80\xd6\x21\x55\x94\x16\x8d\x22\xb4\xc7\x6b\x72\x49\x0c\x8e\x5d\xd9\x17\x4a\xa9\xf2\xbf\xe3\xd6\x49\x61\xc0\xdb\x1e\x2c\x3b\xe7\xf3\xc7\xdf\x3b\x7f\x13\x9f\x4f\x97\x37\xab\xfb\x8f\xb7\x50\x72\xa5\xae\xcf\xe2\x30\x01\xc4\x25\x61\x76\x58\xf8\x65\x45\x8c\x90\x96\x68\x1d\x71\xd2\xab\x39\x1f\xbd\xea\xb5\x5b\x8e\x77\x8a\xc2\x1a\xe0\x8d\xac\x36\xc6\x32\xd4\x56\x5d\x94\xcc\x1b\x37\x16\x22\x37\x9a\x5d\x54\x18\x53\x28\xc2\x8d\x74\x51\x6a\x2a\x91\x3a\xf7\x3a\xc7\x4a\xaa\x5d\x72\x8f\xda\x68\x1a\xbc\xc7\x3c\x27\x72\xc8\x3f\x9f\x4c\x1e\x41\x9b\x5a\x23\xb3\xc1\x27\xb2\x32\x1f\xbf\xb8\xba\x1a\xbe\xf4\xc3\xcf\x92\x51\xc9\xf4\x51\xe4\xcf\xeb\x5a\x73\x3d\xf8\x60\xb4\xf9\x2f\xb9\x45\xaf\x4d\xb6\x83\x3d\x1c\x68\xa3\x70\x72\x0c\xfd\xa3\x2a\x38\xaa\xea\x4f\xa0\x69\x53\xcb\xa7\x43\x28\x9f\xf9\xf1\x1c\xf6\x6d\x08\xfe\x3a\x18\x9a\x03\xbf\x9b\xd3\x9f\x3c\xcc\xdc\x92\x2c\x4a\x1e\x83\x36\xb6\x42\xd5\x6d\x76\x37\x44\x96\x2a\xb4\xdf\x46\xa9\xc9\x68\x78\xfa\x92\x5a\x49\x4d\xc7\xe0\x3f\x4a\x43\x95\x70\xa8\xf2\xa4\x34\x16\xa7\x47\x8e\x45\xe7\x8a\xf8\x50\x68\xeb\x01\xa6\x1f\x8c\x96\x10\x64\x96\xf4\x9c\xa9\x6d\x4a\xde\x1e\xfb\xbd\xb8\x84\xe9\x12\x16\xcb\x15\xcc\x16\xd3\xdb\xc5\x0a\x96\x77\x30\x5b\xc1\x97\xd9\x7c\x0e\xef\xde\xce\xe6\xe7\x70\x29\x9a\xc6\x27\x46\x77\xb8\xbd\xf1\x3a\x48\x73\xd3\xdd\xd9\x41\x3b\x9f\xa5\x56\x6e\x18\x9c\x4d\x93\x5e\xf7\x64\x85\xc6\x75\x54\x48\x2e\xeb\x75\x24\x8d\x08\xe5\x89\xcc\x6c\xb5\x32\x98\xb9\x36\x30\x52\xc8\xe4\x38\xaa\xa4\x8e\xbe\xba\xce\xb7\x22\x00\x1f\xd0\xaf\xcf\xbe\xa3\x05\xa7\x64\x46\xae\x34\x5b\x48\x20\x10\xa2\xd4\xeb\x60\xba\xf0\x42\xbd\xca\x5c\x16\x4d\xd3\x5a\xe9\x4f\x4c\x2c\x42\x4b\x7c\x8f\x8e\xbf\xd0\xaf\x00\x00\x00\xff\xff\x6f\x49\x9f\x7e\x5a\x03\x00\x00")

func templatesRemarkTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRemarkTmpl,
		"templates/remark.tmpl",
	)
}

func templatesRemarkTmpl() (*asset, error) {
	bytes, err := templatesRemarkTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/remark.tmpl", size: 858, mode: os.FileMode(420), modTime: time.Unix(1452273219, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesRevealTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x93\xb1\x4f\xeb\x30\x10\xc6\xf7\xfe\x15\x7e\x9e\xde\x1b\x9e\x8d\xca\x82\xc0\xc9\x52\x90\x18\x40\x20\xe8\xc2\xe8\x24\xd7\xf8\x8a\xe3\x54\xf6\xb5\xa8\x89\xf2\xbf\xe3\x26\x69\x54\x06\x26\xe8\xf6\xc5\x5f\xce\xbf\xef\xee\x12\xf5\xe7\xf6\x69\xb1\x7c\x7b\xbe\x63\xf7\xcb\xc7\x87\x74\xa6\x0c\x55\x96\x59\xed\xca\x84\x83\xe3\xe9\x8c\x31\x65\x40\x17\x07\x11\x65\x05\xa4\x59\x6e\xb4\x0f\x40\x09\xdf\xd2\xea\xff\x15\x1f\x2d\x8b\xee\x9d\x79\xb0\x09\x0f\xb4\xb7\x10\x0c\x00\x71\x66\x3c\xac\x12\x6e\x88\x36\xe1\x5a\x4a\x5d\xd6\xae\xd1\x16\x1a\x5f\x8b\x12\xc9\x6c\x33\x81\xb5\xac\x8a\x79\xb0\x58\x40\x90\x16\xb3\x20\x3d\xec\x40\x5b\x79\x29\xe6\xe2\x42\xe6\xe1\x78\x20\xa2\x3c\x3f\x8b\x0c\x54\x20\x3f\x0c\x12\x9c\x1f\x18\x4f\x7a\x68\x03\x2e\xdb\x7a\x37\x01\x95\x1c\x46\x7e\x90\x59\x5d\xec\xc7\x10\x05\xee\x58\x6e\x75\x08\x09\x1f\xae\x19\xd3\x7d\xb5\x06\xda\x64\x31\xd6\xb6\x3e\xae\x13\x98\x78\xed\x9d\xae\x9b\x1c\x15\x20\x27\xac\x5d\xda\xb6\xa2\xeb\x94\x3c\x3e\x9e\x94\x82\x2b\xa6\x02\x25\x23\x66\xcc\x32\xc8\x41\x87\xdc\xe3\x86\x58\xf0\xf9\x4f\x87\xb1\x0e\x7d\xe7\xa2\x42\x27\xd6\xb1\x87\x18\xa9\xbf\x3b\xfd\x4d\xd0\x7a\xfa\xa2\xbe\x45\xa4\x2f\xc3\x0b\xe8\x90\x50\x5b\x6c\xe0\x6f\x1c\xd1\xa2\x76\x2b\x2c\xbb\xee\xdf\xcd\x69\x91\x92\xc3\x8a\xe2\xd2\xe2\xaf\x93\xce\x3e\x03\x00\x00\xff\xff\x90\x35\x17\xbc\x52\x03\x00\x00")

func templatesRevealTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRevealTmpl,
		"templates/reveal.tmpl",
	)
}

func templatesRevealTmpl() (*asset, error) {
	bytes, err := templatesRevealTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/reveal.tmpl", size: 850, mode: os.FileMode(420), modTime: time.Unix(1452273177, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesShowerTmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x52\xcb\x6e\xdb\x30\x10\x3c\x27\x5f\xb1\x65\x72\xb4\x44\xd9\x85\x9b\x54\x15\x0d\x14\x6d\x80\x1e\xfa\x08\x10\x5f\x7a\xa4\xa8\xb5\xb4\x2d\x45\x2a\x24\xfd\x48\x0c\xfd\x7b\xa9\x47\xd2\xde\x7b\x10\xa0\x9d\xd9\x99\x5d\xed\xa8\x78\xf3\xf9\xc7\xa7\xed\xcf\xfb\x3b\xf8\xb2\xfd\xf6\x75\x73\x59\x34\xa1\xd5\xa0\xa5\xa9\x05\x43\xc3\x06\x00\x65\xb5\xb9\xbc\x28\x5a\x0c\x12\x54\x23\x9d\xc7\x20\xd8\x3e\xec\x92\x5b\xf6\x8a\x1b\xd9\xa2\x60\x07\xc2\x63\x67\x5d\x60\xa0\xac\x09\x68\x62\xdf\x91\xaa\xd0\x88\x9b\xf7\xab\x05\xec\x3d\xba\xc4\x2b\xa9\x65\xa9\x51\x18\xfb\x57\xdd\x84\xd0\x25\xf8\xb8\xa7\x83\x60\xa7\x64\x2f\x13\x65\xdb\x4e\x06\x8a\x7d\xff\x58\x11\x0a\xac\x6a\x1c\x65\x9a\xcc\x6f\x70\xa8\x05\xf3\xe1\x49\xa3\x6f\x10\xe3\xd4\xc6\xe1\x4e\xb0\xc1\xcd\xe7\x9c\xcb\xda\x9a\x67\xa9\xf1\xd9\xd9\xb4\xa6\xd0\xec\xcb\x94\x2c\x6f\xab\x95\xd7\x54\xa1\xe7\x9a\x4a\xcf\x7d\x63\x8f\xe8\xf8\x2a\xcd\xd2\x15\x0f\x0d\xb6\x91\x70\x54\x96\xd6\xf0\xc9\x99\x7b\xe5\x10\x4d\xb2\x7c\x77\x5a\x66\xa9\xf2\x7e\xb8\x09\x9f\x8e\x52\x94\xb6\x7a\x02\xa5\xa5\xf7\x71\x91\xd1\x09\x34\xf9\x10\x5b\x00\xce\x67\x17\xaf\x88\x70\x4d\xa6\xc2\xd3\x02\xae\xc7\xb1\x90\x0b\x48\x1f\xc6\x05\xfa\x3e\x76\x15\x1e\x55\x20\x6b\x5e\x5d\x06\x8a\x45\x31\xd0\x0e\xf0\x71\x56\x43\x06\x7d\x4f\x95\x60\xca\x1e\xd0\xb1\xc8\xa2\xa9\x22\xb4\x39\x9f\xd3\xbe\x2f\xf8\x6c\x32\x8d\x8d\xd4\x6c\x3d\x7c\xc0\x80\x01\x5c\x8d\x42\x68\x96\x70\x1e\x6b\x80\x56\xba\x9a\x4c\xfe\x36\xeb\x06\xf7\xec\xc3\x0c\x2b\xab\xad\xcb\xe1\x6a\x7d\xbb\xfe\xb8\xbe\x7b\x41\x03\x9e\x42\x22\x35\xd5\x26\x57\x31\x0b\x74\x2f\xc4\x2e\x86\x93\xc3\x4d\x96\xc5\xa7\x3b\xf1\x25\xdc\x6f\xe1\x41\x1a\x0f\xdf\xa5\x73\xf6\xb8\xf0\xf1\x3d\x89\xb9\xd3\x6e\x52\x8c\x7b\xf1\x79\xb1\x8b\x22\xde\x96\xba\x00\xde\xa9\xff\x89\x6d\x2a\xd2\x96\x4c\xfa\x2b\xc6\x53\xf0\xc9\x76\xc8\x69\x08\x68\xcc\x2b\xfe\xd5\x9b\xcb\x3f\x01\x00\x00\xff\xff\xed\xc3\xc7\xf8\xed\x02\x00\x00")

func templatesShowerTmplBytes() ([]byte, error) {
	return bindataRead(
		_templatesShowerTmpl,
		"templates/shower.tmpl",
	)
}

func templatesShowerTmpl() (*asset, error) {
	bytes, err := templatesShowerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/shower.tmpl", size: 749, mode: os.FileMode(420), modTime: time.Unix(1450172087, 0)}
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
	"templates/remark.tmpl": templatesRemarkTmpl,
	"templates/reveal.tmpl": templatesRevealTmpl,
	"templates/shower.tmpl": templatesShowerTmpl,
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
	"templates": &bintree{nil, map[string]*bintree{
		"remark.tmpl": &bintree{templatesRemarkTmpl, map[string]*bintree{}},
		"reveal.tmpl": &bintree{templatesRevealTmpl, map[string]*bintree{}},
		"shower.tmpl": &bintree{templatesShowerTmpl, map[string]*bintree{}},
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

