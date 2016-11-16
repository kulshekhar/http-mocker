// Code generated by go-bindata.
// sources:
// script/interceptor.ts
// script/script.min.js
// DO NOT EDIT!

package main

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

var _scriptInterceptorTs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x4d\x6f\xe3\x36\x10\xbd\xeb\x57\xcc\x29\xa2\x00\x97\x76\x7b\x2a\x24\x28\x40\x11\xb4\x48\x0b\xa7\x5e\xd8\x09\x10\x20\xc8\x81\x91\xc6\xb6\x02\x9a\xe4\x92\xd4\x26\x46\xd6\xff\x7d\x41\x91\xfa\xb0\xe2\x18\x39\xec\x61\x6f\xb6\x38\xf3\xe6\xcd\x7b\x33\x94\x0a\xce\x8c\x81\xfb\xeb\xe5\xbf\xc2\xa2\x2e\x50\x59\xa9\xe1\x2d\x02\x50\xba\xfa\xc6\x2c\x82\xb1\xcc\x56\x05\x68\x64\xa5\x14\x7c\x0f\xaf\x0b\x85\x02\x72\xb8\xbf\x99\x5f\x5b\xab\x96\xf8\xb5\x46\x63\xa9\xd2\xd2\x4a\xbb\x57\x48\xa5\x42\x91\x9d\x03\x58\xa1\x28\xcf\x01\x18\x14\xe5\x09\x00\x87\xbb\x58\x42\x0e\x6b\xc6\x0d\x9e\x08\x70\x79\xc3\x80\x41\xc4\x8e\xd9\x62\x8b\xda\xa4\x70\xe3\x7f\x3d\x3c\x42\x0e\x0f\x8f\x4d\x50\x21\x85\xb1\xba\x2e\xac\xd4\xa4\x4d\xb0\x4c\x6f\xd0\xde\x2d\xe7\x29\x18\xab\x2b\xb1\x81\x1c\xe2\xad\xb5\x2a\x9d\x4e\xb9\x2c\x18\xdf\x4a\x63\xd3\xdf\xff\x9c\xcd\x66\xd3\x38\x81\x37\x38\x38\xa4\xa6\x0c\x09\xc5\xba\x5a\x49\xa3\x27\x80\xdd\x56\x86\xb6\x4c\xa8\xaa\x4d\x17\x9a\xb8\x6e\x7a\x84\xbf\x38\x27\x27\x18\x9f\xc1\xa1\xb4\x7b\xd0\x63\x69\x34\x56\x6a\x24\x6d\xde\x59\xc7\x9c\x21\x47\x43\x40\x1b\xa3\xb3\xf3\x99\x26\x58\x39\xca\x5c\x05\x07\x1b\x1a\x9d\xff\xbf\x04\x9b\xaa\x3d\x21\x47\x72\x76\x8f\x5d\x19\x72\x96\x5b\x92\x9d\x48\x73\x35\x3e\x4e\x73\xc4\x7a\x5f\xda\x19\x3b\xae\xe9\xa0\x53\xf8\xa7\x16\x85\xad\xa4\x68\xc9\x55\x6b\x20\xa3\x7e\xfc\x1e\xb4\x01\xce\x66\x5b\xeb\x20\x4d\x83\x1f\x26\x1a\x0c\xf2\x35\xe4\x0d\xd1\x2c\xfa\x94\xe6\xeb\x50\x1d\xc8\x0e\xed\x56\x96\xed\xf0\x4f\xa0\xd6\xbc\xfd\xd3\x57\x6e\x24\x90\xba\xda\x54\x82\xf1\xbb\xe5\x1c\x72\x17\x97\x85\x53\x0f\xe1\xf6\xe6\xcb\x62\x75\x1b\x07\x0a\xd0\xec\x31\x2d\x18\xe7\xc4\xe5\x4f\x42\xdc\xa4\xa1\x4b\xbb\xbd\x0b\x2a\x1f\x3e\x56\xad\x91\xdc\x29\xfb\x29\xd5\xfc\xe5\xf0\x73\x55\x0b\xf3\xd6\xab\xe6\xe0\x03\xbe\x07\x53\x4c\x1b\x2c\xbd\x34\x02\x5f\xe0\x6e\x39\x27\x63\xd5\x92\xec\x38\x87\x43\xee\xb5\xe8\x56\x7c\x5d\x71\x8b\x9a\x90\x5d\x02\xf9\x65\xd7\x81\x6f\xd3\x11\x91\x6b\xd8\x51\x7c\x55\x1a\x8d\x71\x3c\xf2\x1c\x62\xef\x55\x9c\x0c\xc2\x61\x14\xd5\x30\x5a\xe2\xe6\xef\x57\x45\x86\x27\x49\xd6\xa5\x1c\xa2\xee\xa7\x27\xe7\x39\x19\xc8\xfb\xd6\xa8\x62\x76\x2b\xd8\x0e\x69\xb8\xfe\x8e\xa1\x3a\x00\xaf\x38\x90\x21\x1f\x8f\x46\x39\x8a\x8d\xdd\xc2\x25\xcc\xe0\xe2\xe2\xfd\xf9\xc3\xec\xb1\x0d\xc9\x4f\x16\xf6\x87\x5d\x62\xd7\xc0\x61\xac\xad\x62\x7b\x2e\x99\x33\xed\xcd\x5d\x48\x4a\x0a\x83\x29\xc4\x3a\x9e\x34\xd7\x54\x6d\xae\x64\x89\x29\xfc\x31\x9b\xb9\xc9\x8b\x7a\x99\xf9\x80\xe3\x50\xd3\x00\x48\xfb\x6c\xc8\x81\x3b\xc2\xfd\x93\xec\x5d\x74\x5b\xba\x8d\x6d\xff\x77\xbc\x87\xa5\x83\xc3\x1d\xf5\x1c\x62\xf9\xf4\x8c\x85\x8d\x07\xe3\x16\x96\xd1\xa0\x0d\x63\x7a\x8d\xac\x44\x4d\xe2\x2b\x29\x2c\x0a\xfb\xdb\xed\x5e\x61\x3c\x81\x98\x29\xc5\xab\x82\xb9\x79\x9d\x3e\x1b\x29\xe2\x81\xdb\x6e\x9e\x87\x9b\xf9\xdf\x6a\xf1\x3f\xf5\x83\x54\xad\xf7\x24\x30\x48\x7a\x51\x0f\x80\xdc\xe0\x40\x8e\x31\x42\x9b\xd2\x37\x36\x58\xea\x43\x14\x91\x97\x4a\x94\xf2\x05\x98\x01\x26\xf6\x09\x1d\x7d\x87\x8c\xef\xf2\x2c\x8a\xfc\xf7\x4a\x78\x27\x36\x95\xfb\x69\x4b\xc3\x34\xc3\xf7\x70\x55\x65\xfe\x2d\x18\x6c\x66\x62\x9f\x85\xf7\x51\x6b\xb4\xa8\x77\x4f\xa8\xb3\xe8\x10\xfd\x08\x00\x00\xff\xff\x92\x49\x78\x64\x08\x09\x00\x00")

func scriptInterceptorTsBytes() ([]byte, error) {
	return bindataRead(
		_scriptInterceptorTs,
		"script/interceptor.ts",
	)
}

func scriptInterceptorTs() (*asset, error) {
	bytes, err := scriptInterceptorTsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "script/interceptor.ts", size: 2312, mode: os.FileMode(420), modTime: time.Unix(1479319937, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scriptScriptMinJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x54\x51\x6b\xdb\x30\x10\xfe\x2b\xa9\x1e\x82\x04\x9a\xe3\xee\x69\x24\x68\x63\x94\x41\x37\xda\x65\xa4\x2d\x14\x4a\x1f\x34\xe7\x12\xab\xb8\x3a\x4d\xba\xb4\x0d\xc6\xff\x7d\xc8\x76\x6a\xbb\x9b\xc3\x5e\xf6\xe6\xf3\xdd\x77\xf7\xdd\xdd\x77\x7a\xd2\x7e\x72\x7b\xbe\xfa\x6a\x09\x7c\x06\x8e\xd0\xab\xcd\xce\x66\x64\xd0\x72\x51\x1e\x3e\x27\xc4\x49\x94\x4f\x68\xd6\x93\x54\x29\x45\xd3\x29\x27\xc5\x72\x22\x37\x9f\xcd\x0a\xcc\x74\x91\x63\xa0\xf9\xe9\x87\x34\x4d\x67\x4c\x48\xca\x4d\x48\x48\xfb\x2d\xd0\xcd\xea\x42\x51\xf3\xe3\x51\x53\x96\x83\x0f\xea\xee\xbe\xf2\x40\x3b\x6f\x27\x94\x38\x8f\x84\xb4\x77\xd0\xb8\xbb\xea\x24\xca\x01\x2a\x71\xbb\x90\x73\x12\x95\xfc\x03\xf4\xb9\x28\x06\x38\x0e\x6a\x00\x15\x35\x36\xd1\xce\x15\x7b\x0e\x92\xc4\x22\xb6\x0d\xc3\x4c\x1e\x02\xa1\x87\x7e\xfb\xb7\x97\x17\xe7\x44\x6e\x05\xbf\x76\x10\xfa\xb1\xe8\xc0\x2a\x4a\x5e\x96\x0e\xac\x1c\x8d\x0a\x60\xd7\x31\xea\x0a\xec\x3a\xd6\xfa\xff\x15\x3a\x8f\x39\x2c\xb4\x5f\xad\x1e\xca\xab\x27\xa6\xe6\x47\x09\xb4\x8b\x7c\x45\xc4\x32\xe3\x88\x48\x46\x8c\xb0\x88\xb5\x3a\x26\x20\x4a\xb3\xe1\x27\x54\x17\x59\xae\x44\x19\xd7\x61\xeb\x9d\x2d\x8e\x4f\xa4\xdb\xb2\x74\x6d\x3f\xe8\xcd\xd6\x58\x5d\x44\xa1\x39\x49\x8a\xfd\x58\x5e\x5d\x33\x09\x49\xa6\x8b\x82\xc7\x10\x49\xd2\x76\x6a\x14\x55\x35\x42\x32\xb6\xf7\x17\x92\xb1\xaf\x7f\x25\x19\x06\x29\x1a\x0c\x29\x0b\xcf\x93\x9b\xd5\x05\x7f\xcb\x57\x48\xa7\x6c\x27\xf0\x8d\x29\x08\x3c\xef\x33\x60\x81\xbc\xb1\x5b\xa6\x54\x4c\x8f\x9b\x09\x24\xf0\xe2\x3c\x84\x60\xd0\x4e\xa7\xbc\x6f\xd6\x65\x56\xb0\xfd\xf2\xe2\x06\xff\x45\x23\xf7\x28\x27\xa7\x29\xb7\xfa\xb1\x3d\x9b\x61\xd4\xa2\x3d\x49\x9b\x14\x60\xb7\x94\x7f\x4c\xa7\x53\x7b\x97\xde\xb7\xa6\xea\xc3\x9b\x5f\x95\x90\xa8\x4a\x0f\xc1\xa1\x0d\x30\x67\x9e\xc9\x40\x9a\x76\xe1\x0c\xd7\x30\x7f\x9f\xa6\xd5\xc2\xf5\x92\x71\x4c\x3a\xb7\x72\x31\x75\x67\x4b\x4c\x0e\x89\x1a\xd7\xc1\x12\x92\xe1\xcf\x07\xc8\xa8\x9b\x01\x7e\x6a\x26\x19\x80\xda\x0d\x9c\x83\x5e\x83\xe7\xec\x0c\x2d\x81\xa5\x77\xd7\x7b\x07\x4c\xb2\x78\xf0\x26\xd3\x71\x98\xb3\x87\x80\x96\x89\x81\x2c\xbe\x5d\x2d\xbf\x27\xcd\x80\xcd\x66\xcf\x51\x08\x31\xef\xfb\xb1\x95\x4a\x7d\x84\xea\xa8\x30\x65\x7b\x84\xe3\x51\x51\x19\xf2\xa0\x78\x75\x72\x2a\x0f\xc2\xaa\xbf\x2b\x2e\x16\xcf\xc6\xae\xf1\x39\x79\xf3\x1a\x0f\xcd\x7a\x93\x97\x8d\x60\x46\x1e\x6a\x51\xbe\x3e\xae\x31\xeb\xef\x00\x00\x00\xff\xff\x69\xc6\xbe\xf5\xe1\x05\x00\x00")

func scriptScriptMinJsBytes() ([]byte, error) {
	return bindataRead(
		_scriptScriptMinJs,
		"script/script.min.js",
	)
}

func scriptScriptMinJs() (*asset, error) {
	bytes, err := scriptScriptMinJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "script/script.min.js", size: 1505, mode: os.FileMode(420), modTime: time.Unix(1479319939, 0)}
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
	"script/interceptor.ts": scriptInterceptorTs,
	"script/script.min.js": scriptScriptMinJs,
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
	"script": &bintree{nil, map[string]*bintree{
		"interceptor.ts": &bintree{scriptInterceptorTs, map[string]*bintree{}},
		"script.min.js": &bintree{scriptScriptMinJs, map[string]*bintree{}},
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
