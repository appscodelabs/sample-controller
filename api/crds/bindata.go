// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// wardle.example.com_fischers.yaml
// wardle.example.com_flunders.yaml
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

var _wardleExampleCom_fischersYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x94\x4f\x6f\x1c\x37\x0c\xc5\xef\xf3\x29\x08\xf4\x90\x4b\x77\x16\x46\x2f\xc5\xdc\x0a\xa7\x06\x82\xfe\x41\x11\xa7\xb9\x73\x25\xee\x0e\x6b\x0d\xa5\x92\xd4\x3a\xee\xa7\x2f\xa4\xd9\x7f\x46\x63\xf4\x94\xb9\xcd\x1b\xee\x23\xf5\xf8\x5b\x0d\x9b\xcd\x66\xc0\xc2\x9f\x49\x8d\xb3\x4c\x80\x85\xe9\x8b\x93\xb4\x37\x1b\x9f\x7e\xb4\x91\xf3\xf6\x78\xb7\x23\xc7\xbb\xe1\x89\x25\x4e\x70\x5f\xcd\xf3\xf2\x91\x2c\x57\x0d\xf4\x9e\xf6\x2c\xec\x9c\x65\x58\xc8\x31\xa2\xe3\x34\x00\x04\x25\x6c\xe2\x27\x5e\xc8\x1c\x97\x32\x81\xd4\x94\x06\x00\xc1\x85\x26\xd8\xb3\x85\x99\xd4\xc6\x67\xd4\x98\x68\xa4\x2f\xb8\x94\x44\x63\xc8\xcb\x60\x85\x42\xf3\x38\x68\xae\x65\x82\xaf\x54\xac\x2e\xd6\x8a\x00\xd6\xa9\x1e\x56\xc3\xae\x24\x36\xff\xe5\x56\xfd\x95\xcd\xfb\x97\x92\xaa\x62\xba\xb6\xef\xa2\xb1\x1c\x6a\x42\xbd\xc8\x03\x40\x51\x32\xd2\x23\xfd\x29\x4f\x92\x9f\xe5\x81\x29\x45\x9b\x60\x8f\xc9\x68\x00\xb0\x90\x0b\x4d\xf0\x7b\x9b\xa2\x60\xa0\x38\x00\x1c\x31\x71\xec\x87\x5e\xe7\xca\x85\xe4\xa7\x3f\x3e\x7c\xfe\xe1\x31\xcc\xb4\xe0\x2a\x36\xe7\x5c\x48\x9d\xcf\xe3\xb7\xe7\x66\x03\x17\x0d\x20\x92\x05\xe5\xd2\x1d\xe1\x5d\xb3\x5a\x6b\x20\xb6\xcc\xc9\xc0\x67\x82\xe3\xaa\x51\x04\xeb\x6d\x20\xef\xc1\x67\x36\x50\xea\x67\x10\xef\x23\xdd\xd8\x42\x2b\x41\x81\xbc\xfb\x8b\x82\x8f\xf0\xd8\xce\xa9\x06\x36\xe7\x9a\x22\x84\x2c\x47\x52\x07\xa5\x90\x0f\xc2\xff\x5c\x9c\x0d\x3c\xf7\x96\x09\x9d\x4e\x79\x9e\x1f\x16\x27\x15\x4c\x2d\x84\x4a\xdf\x03\x4a\x84\x05\x5f\x40\xa9\xf5\x80\x2a\x37\x6e\xbd\xc4\x46\xf8\x2d\x2b\x01\xcb\x3e\x4f\x30\xbb\x17\x9b\xb6\xdb\x03\xfb\x99\xb9\x90\x97\xa5\x0a\xfb\xcb\x36\x64\x71\xe5\x5d\xf5\xac\xb6\x8d\x74\xa4\xb4\x35\x3e\x6c\x50\xc3\xcc\x4e\xc1\xab\xd2\x16\x0b\x6f\xfa\xe0\xe2\x1d\xdc\x25\x7e\xa7\x27\x40\xed\xdd\xcd\xa4\xfe\xd2\xd6\x66\xae\x2c\x87\x8b\x1c\xd9\x30\xa5\xfc\x4c\xf1\x21\x55\x89\xa4\xf6\xe6\x16\xde\xff\xa7\x14\xe6\x9c\xa2\x01\x76\xe8\x5a\xb2\xa7\x0f\x63\x67\x03\x7c\x46\x07\x54\xba\x69\x32\xde\x98\xb3\xd3\xf2\xaa\xdb\x1b\x23\x9e\x65\x54\xc5\x97\x8b\xda\xc9\x7f\x13\x98\xf6\x0f\x00\x6e\xa3\xad\x66\x6b\xf0\x57\x2e\x9a\xd4\xd6\xf9\xf1\xe7\xc7\x4f\x70\x4e\xab\xb3\xf3\x1a\x96\x8e\xc9\xf5\x67\x76\x25\xa6\x6d\x98\x65\x4f\xba\x12\xb7\xd7\xbc\x74\x47\x92\x58\x32\x8b\xf7\x97\x90\x98\xe4\x35\x2d\x56\x77\x0b\x7b\x43\xf4\xef\x4a\xe6\x0d\xac\x11\xee\x51\x24\x3b\xec\x08\x6a\x89\xe8\x14\x47\xf8\x20\x70\x8f\x0b\xa5\x7b\x34\xfa\xe6\xbc\xb4\x80\x6d\xd3\x22\xfd\x7f\x62\x6e\x6f\xba\xd7\x85\x6b\x5a\xc3\x57\xa5\xe3\xf9\x92\x3d\xde\x61\x2a\x33\xde\x5d\xb5\xce\xc0\xe6\x74\x37\xde\x7c\x06\xe8\xb7\x50\x9c\xc0\xb5\xd2\x2a\x78\x56\x3c\xd0\x49\xf9\x37\x00\x00\xff\xff\x0a\x79\xca\xf8\xc1\x05\x00\x00")

func wardleExampleCom_fischersYamlBytes() ([]byte, error) {
	return bindataRead(
		_wardleExampleCom_fischersYaml,
		"wardle.example.com_fischers.yaml",
	)
}

func wardleExampleCom_fischersYaml() (*asset, error) {
	bytes, err := wardleExampleCom_fischersYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wardle.example.com_fischers.yaml", size: 1473, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _wardleExampleCom_flundersYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\x4d\x6f\xe4\x36\x0c\xbd\xfb\x57\x10\xe8\x61\x2f\x3b\x1e\x04\xbd\x14\xbe\x2d\xd2\x2e\xb0\xe8\x07\x8a\x24\xdd\x3b\x47\xa6\xc7\x6c\x64\x4a\x15\xa9\x49\xd2\x5f\x5f\x48\xb6\x67\x26\x83\xd9\x6c\x2e\xf5\xc9\x7e\x26\x1f\x9f\x1e\x29\xa9\xd9\x6c\x36\x0d\x46\xfe\x4a\x49\x39\x48\x07\x18\x99\x9e\x8d\xa4\x7c\x69\xfb\xf8\x93\xb6\x1c\xb6\x87\x9b\x1d\x19\xde\x34\x8f\x2c\x7d\x07\xb7\x59\x2d\x4c\x77\xa4\x21\x27\x47\x3f\xd3\xc0\xc2\xc6\x41\x9a\x89\x0c\x7b\x34\xec\x1a\x00\x97\x08\x0b\xf8\xc0\x13\xa9\xe1\x14\x3b\x90\xec\x7d\x03\x20\x38\x51\x07\x83\xcf\xd2\x53\xd2\xf6\x09\x53\xef\xa9\xa5\x67\x9c\xa2\xa7\xd6\x85\xa9\xd1\x48\xae\x70\xec\x53\xc8\xb1\x83\x2b\x11\x33\x8b\x96\x20\x80\x59\xd5\xe7\x99\xb0\x22\x9e\xd5\x7e\x3d\x47\x7f\x63\xb5\xfa\x27\xfa\x9c\xd0\x9f\xca\x57\x50\x59\xf6\xd9\x63\x3a\xc2\x0d\x40\x4c\xa4\x94\x0e\xf4\x97\x3c\x4a\x78\x92\xcf\x4c\xbe\xd7\x0e\x06\xf4\x4a\x0d\x80\xba\x10\xa9\x83\x3f\x8a\x8a\x88\x8e\xfa\x06\xe0\x80\x9e\xfb\xba\xe8\x59\x57\x88\x24\x9f\xfe\xfc\xf2\xf5\xc7\x7b\x37\xd2\x84\x33\x08\xd0\x93\xba\xc4\xb1\xc6\xad\xfa\x80\x15\x50\x60\x59\x22\xd8\x4b\x24\x78\x62\x1b\x01\xa1\x98\x01\x28\x7d\x79\x35\xb4\xac\xed\xc2\x13\x53\x88\x94\x8c\x57\x1b\xca\x73\xd6\xc9\x23\x76\x51\xf1\x43\x91\x34\xc7\x40\x5f\x7a\x47\x0a\x36\x12\x1c\x66\x8c\x7a\xd0\x2a\x17\xc2\x00\x36\xb2\x42\xa2\xea\x85\x58\x5d\xda\x19\x2d\x94\x10\x14\x08\xbb\xbf\xc9\x59\x0b\xf7\xc5\xaf\xa4\xa0\x63\xc8\xbe\x07\x17\xe4\x40\xc9\x20\x91\x0b\x7b\xe1\x7f\x8f\xcc\x0a\x16\x6a\x49\x8f\x46\x4b\x5f\xd6\x87\xc5\x28\x09\xfa\x62\x66\xa6\x8f\x75\xe1\x13\xbe\x40\xa2\x52\x03\xb2\x9c\xb1\xd5\x10\x6d\xe1\xf7\x90\x08\x58\x86\xd0\xc1\x68\x16\xb5\xdb\x6e\xf7\x6c\xeb\xec\xba\x30\x4d\x59\xd8\x5e\xb6\x2e\x88\x25\xde\x65\x0b\x49\xb7\x3d\x1d\xc8\x6f\x95\xf7\x1b\x4c\x6e\x64\x23\x67\x39\xd1\x16\x23\x6f\xaa\x70\xb1\xba\x01\xa6\xfe\x87\xb4\x0c\xba\x7e\x38\x53\x5a\x3a\xd4\x81\x5a\x62\xd9\x1f\xe1\x3a\x88\xdf\xf4\xbd\x0c\x64\xed\xf3\x92\x36\xeb\x3f\xd9\x5b\xa0\xe2\xca\xdd\x2f\xf7\x0f\xb0\x16\xad\x2d\x78\xed\x79\x75\xfb\x94\xa6\x27\xe3\x8b\x51\x2c\x03\xa5\xb9\x71\x43\x0a\x53\x65\x24\xe9\x63\x60\xb1\xfa\xe1\x3c\x93\xbc\x36\x5d\xf3\x6e\x62\x2b\x9d\xfe\x27\x93\x5a\xe9\x4f\x0b\xb7\x28\x12\x0c\x76\x04\x39\xf6\x68\xd4\xb7\xf0\x45\xe0\x16\x27\xf2\xb7\xa8\xf4\xbf\xdb\x5e\x1c\xd6\x4d\xb1\xf4\xfb\xc6\x9f\x1f\x3c\xaf\x03\x67\xb7\x8e\xf0\x7a\xb2\x5c\xed\xd0\xb2\x17\xef\xcb\x86\xe3\x79\x53\x94\x78\x1e\xd8\xd5\xd1\xaf\xe3\xbe\x46\xb5\x67\x2c\xd7\x76\x62\x79\x06\x2e\x03\x9f\xee\x68\xa0\x44\xe2\xe8\xf5\xdf\x8b\xea\x9f\xea\x91\x36\xd7\x58\x12\x3f\xc2\x94\x2d\xa3\xf7\x2f\x40\xcf\xce\x67\xe5\x03\xad\x7b\x67\x91\x71\xe4\x6e\x2f\xb8\xaf\xfa\x54\x45\x5d\x24\xbe\x57\x94\x04\x1b\x29\xad\xe9\x6f\x4a\xbb\x58\xf6\xbb\xa5\xa5\x35\xe3\xa1\x44\xbc\xa5\xeb\x61\xa4\x53\x74\x25\x7c\x67\x91\xeb\x53\x51\xcf\xd5\xef\xce\x45\x8d\x3a\x4e\xc6\xfc\xf5\xad\x91\xb8\x52\xe7\x02\x3a\xac\x37\xee\xe1\x06\x7d\x1c\xf1\xe6\x84\x55\x29\x9b\xe5\xa2\x3c\xfb\x0d\x50\xaf\xa4\xbe\x03\x4b\x99\x66\xc0\x42\xc2\x3d\x9d\x2e\xa6\x53\xda\x7c\x65\xbf\x9d\x55\x91\xff\x02\x00\x00\xff\xff\xe8\x1f\x18\x54\x04\x08\x00\x00")

func wardleExampleCom_flundersYamlBytes() ([]byte, error) {
	return bindataRead(
		_wardleExampleCom_flundersYaml,
		"wardle.example.com_flunders.yaml",
	)
}

func wardleExampleCom_flundersYaml() (*asset, error) {
	bytes, err := wardleExampleCom_flundersYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "wardle.example.com_flunders.yaml", size: 2052, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"wardle.example.com_fischers.yaml": wardleExampleCom_fischersYaml,
	"wardle.example.com_flunders.yaml": wardleExampleCom_flundersYaml,
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
	"wardle.example.com_fischers.yaml": {wardleExampleCom_fischersYaml, map[string]*bintree{}},
	"wardle.example.com_flunders.yaml": {wardleExampleCom_flundersYaml, map[string]*bintree{}},
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