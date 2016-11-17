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

var _scriptInterceptorTs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x56\x5f\x6f\xdb\x36\x10\x7f\xf7\xa7\xb8\x87\xb9\x92\x31\x97\x71\xfb\x34\x48\x50\x87\xa2\xdb\x90\x0d\xe9\x3c\x38\x0d\x50\x20\x0b\x60\x46\x3a\xd9\x2a\x68\x92\x23\xa9\x25\x5e\xec\xef\x3e\x50\xa4\x24\x4a\x71\x8c\x02\xdd\xc3\x5e\x0c\x99\x77\xf7\xbb\xdf\xfd\x25\x73\x46\xb5\x86\xcf\x97\xab\x5f\xb9\x41\x95\xa3\x34\x42\xc1\xd3\x04\x40\xaa\xea\x6f\x6a\x10\xb4\xa1\xa6\xca\x41\x21\x2d\x04\x67\x7b\x78\x5c\x4a\xe4\x90\xc1\xe7\x8f\x57\x97\xc6\xc8\x15\xfe\x55\xa3\x36\x44\x2a\x61\x84\xd9\x4b\x24\x42\x22\x4f\xcf\x01\x5c\x23\x2f\xce\x01\x68\xe4\xc5\x09\x00\x8b\xbb\x5c\x41\x06\x25\x65\x1a\x4f\x28\x58\xbb\x50\x21\xd0\xd8\x51\x93\x6f\x51\xe9\x04\x3e\xba\xaf\xdb\x3b\xc8\xe0\xf6\xae\x51\xca\x05\xd7\x46\xd5\xb9\x11\x2a\x6e\x0d\x0c\x55\x1b\x34\x37\xab\xab\x04\xb4\x51\x15\xdf\x40\x06\xd1\xd6\x18\x99\x5c\x5c\x30\x91\x53\xb6\x15\xda\x24\x6f\x7e\x58\x2c\x16\x17\xd1\x0c\x9e\xe0\x68\x91\x1a\x37\xb1\x77\xd6\xf9\x9a\x35\xf9\x04\x30\xdb\x4a\x93\x96\x09\x91\xb5\xee\x54\x67\x36\x9a\x1e\xe1\x3d\x63\xf1\x09\xc6\x67\x70\x08\xe9\x0e\x7a\x2c\x85\xda\x08\x85\x71\x6b\x77\xb6\x62\xb6\x20\x83\x26\x20\x4d\xa1\xd3\xf3\x96\xda\x97\x72\x64\x79\xed\x2b\xd8\xd0\xe8\xea\xff\xbf\x60\x53\xb5\x92\x78\x90\xce\xee\xd8\xba\x89\xcf\x72\x9b\xa5\x27\xcc\xac\x8f\x97\xcd\x2c\xb1\xbe\x2e\x6d\x8f\x0d\x7d\x5a\xe8\x04\x7e\xa9\x79\x6e\x2a\xc1\x5b\x72\x55\x09\xf1\x28\x1e\x37\x07\xad\x82\x2d\xb3\xa9\x95\x4f\x4d\x83\xef\x3b\x1a\x34\xb2\x12\xb2\x86\x68\x3a\xf9\xaa\x9c\x97\xde\x3b\xc4\x3b\x34\x5b\x51\xb4\xcd\x3f\x87\x5a\xb1\xf6\x4f\xef\xb9\x49\x81\x50\xd5\xa6\xe2\x94\xdd\xac\xae\x20\xb3\x7a\xa9\x97\x3a\x08\x3b\x37\x7f\x2c\xaf\x3f\x45\x9e\x02\x34\x73\x4c\x72\xca\x58\x6c\xed\xe7\x5e\x6f\xde\xd0\x25\xdd\xdc\xf9\x2c\x1f\x5f\xce\x5a\x93\x72\x9b\xd9\xaf\xca\x9a\x5b\x0e\xff\x6d\xd6\x7c\xbf\xf5\x59\x2b\xa8\xa1\x09\x50\xbe\xb7\x7e\xbc\x23\x87\x2a\xa9\xd2\x58\xb8\x1c\x71\x7c\x80\x9b\xd5\x55\x3c\x4e\xdf\x2c\x1d\xda\x30\xc8\x5c\x52\xba\x59\x2f\x2b\x66\x50\xc5\xf1\x6e\x06\xd9\xbb\x2e\x14\x17\xaf\x65\x24\x4a\xd8\x11\x7c\x94\x0a\xb5\xb6\x84\xb2\x0c\x22\x57\xb4\x68\x16\xa8\xc3\x48\xab\x61\xb4\xc2\xcd\xcf\x8f\x32\x0e\x25\xb3\xb4\x33\x39\x4e\xba\x4f\x47\xce\x71\xd2\x90\xf5\xa1\x11\x49\xcd\x96\xd3\x1d\x12\xbf\x07\x87\x50\x3d\x82\xcb\x3d\xc4\x21\x21\x0f\xf7\xea\xd5\xf3\x43\xc2\x90\x6f\xcc\x16\xde\xc1\xe2\xa4\xfc\x76\x71\xd7\xaa\x64\x27\xe9\x38\x61\x67\xd8\x85\x75\x1c\x67\x5c\xd2\x3d\x13\xd4\xd6\xf4\xc9\xee\x2b\x29\xb8\xc6\xc4\x7d\xd7\xcc\x24\x10\x15\x58\xd2\x9a\x75\x8a\x11\x1c\xe7\xcd\x7e\xab\xf5\x07\x51\x60\x02\x6f\x17\x0b\xdb\xb2\x93\xbe\x2c\x2c\x60\x1f\xd6\xc0\x23\x90\xde\x1a\x32\x60\x36\x94\xe0\xe4\x70\xb0\x88\xe9\x33\xa3\x96\x5b\x6b\xd2\xfd\x3f\x1c\xce\xb3\xed\x22\x07\x64\x1a\x03\x3a\x2e\x7e\xfb\x2b\x18\xbe\x57\x1b\xdd\xb4\x71\x77\x4b\xb6\x6a\x0c\x0d\xec\x50\x6b\xba\xb1\xbe\xd7\xd3\xdc\x76\x34\x17\x06\xb6\x94\x17\x0c\x8b\x04\xa6\xf9\x77\x4f\xe3\xc6\x3e\xae\xd3\x41\xab\xda\x31\x19\xf5\xa3\xc7\xfc\x3e\x83\xf5\x9f\x7c\x9a\xff\xd4\x0c\xd2\x34\x0f\x0c\x07\x4a\xbe\xd9\x2d\x50\xd8\xe4\xf0\x23\xac\xa7\x7a\x0d\x09\xac\xa7\xcb\x81\x6d\x10\x99\xbb\x31\x03\x19\xc0\xba\x14\xdc\xbc\xd6\x66\xcf\x30\x81\xca\x50\x56\xe5\x29\xe4\x82\x09\x95\x80\xc2\x22\x05\x27\xaf\xfe\xc1\x04\xde\xbc\x95\x8f\xeb\xf9\xcb\xe6\x5c\xa8\x1d\x65\x9d\xf9\x46\x21\xf2\x67\x00\xe9\x08\xc1\x06\x12\x9e\x0c\x26\xef\x54\x0c\x35\xd7\xdb\xaa\x34\x61\x18\xdf\x14\xc4\xb7\x85\x10\xd0\xf5\x24\x09\x13\x9b\xd8\x17\x6c\x0e\x84\x90\x80\x7c\x3f\x7f\xe1\xa0\xf8\x92\x76\x23\x98\x41\x24\xee\xbf\x60\x6e\xa2\x60\x99\xfa\x3b\x47\xa3\xf1\xdb\xf8\x12\x69\x81\x2a\x8e\x3e\x08\x6e\x90\x9b\xd7\x9f\xf6\x12\xa3\x39\x44\x54\x4a\x56\xe5\xd4\xae\xe5\x8b\x2f\x5a\xf0\x28\xa0\x68\xd7\x76\x78\x01\xfd\x76\xbd\xfc\x9d\xb8\x0e\xaa\xca\x7d\xec\x19\xcc\xfa\xe5\xf0\x6c\x5a\xc6\x08\xad\x49\x1f\x18\xf4\x77\xd7\x71\x32\x89\x1f\x2a\x5e\x88\x07\xa0\xba\xb9\x1c\xc8\xe8\xb9\x3d\x7e\xb2\xa4\x93\x89\x7b\x96\xfb\xa7\x5f\xe3\x39\x7c\xa5\x36\xf0\xb2\xbe\x67\x55\x0e\xfd\x8e\x4d\xfc\x0e\x87\x43\x7b\x6d\x87\x7a\xfd\x3e\xa3\x7c\x3f\x90\x84\x0b\x8c\xd7\xbb\x7b\x54\x56\xec\x1e\xb5\xc7\xc9\xbf\x01\x00\x00\xff\xff\xf1\xa7\x59\xf5\x21\x0c\x00\x00")

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

	info := bindataFileInfo{name: "script/interceptor.ts", size: 3105, mode: os.FileMode(420), modTime: time.Unix(1479347375, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scriptScriptMinJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x55\x51\x6f\xe4\x26\x10\xfe\x2b\x3e\xa4\x5d\x19\x1d\x47\x9c\x3c\x55\xb6\xe8\xa9\xba\x56\x4a\xab\x5c\xb7\xda\x5c\xa4\x93\xd2\x3c\x50\x7b\xbc\xe6\x44\x80\xc2\x6c\x92\xed\xc6\xff\xbd\xc2\xf6\xae\xed\xa4\x1b\xf5\xa5\x7d\x63\x98\xf9\xe6\x1b\x18\xbe\xe1\x41\xfa\xe4\xeb\xe5\xfa\x67\x83\xe0\x4b\x70\x68\xbd\xa8\xb7\xa6\x44\x65\x4d\x4a\xf7\x87\x65\x82\x29\xd2\xfd\x83\x55\x55\x92\x09\x21\x70\xb9\x4c\x51\x90\x06\xd1\xe5\x67\x67\xda\x96\x52\x37\x36\x60\x7e\xfe\x5d\x96\x65\x67\x84\x32\x6c\x54\xe0\x28\xfd\x06\xf0\x66\x7d\x25\xb0\xdf\xb8\x97\x58\x36\xe0\x83\xb8\xbd\x6b\x3d\xe0\xd6\x9b\x04\xb9\xf3\x16\x2d\xee\x1c\xf4\xee\x91\x1d\xe9\x7e\x86\xe2\x6e\x1b\x9a\x14\x69\xcb\x5e\x81\x7e\xd0\x7a\x86\x4b\x41\xcc\xa0\xb4\xc3\x72\xe9\x9c\xde\xa5\xc0\x90\x16\xf1\xd8\x30\xcf\xe4\x21\xa0\xf5\x30\x3d\xfe\xd7\xcf\x57\x97\x88\x6e\x0d\x7f\x6e\x21\x4c\x63\xad\x03\x23\x90\x3f\xad\x1c\x18\x76\x32\x2a\x80\xa9\x62\xd4\x35\x98\x2a\x72\xfd\xf7\x0c\xa3\x47\x1d\x1a\x3a\x65\xeb\x2e\xe5\xe8\x89\xa9\xd3\x37\x0b\x18\x1a\x79\x44\x44\x9a\xd3\x88\x58\x0c\x3d\x51\x45\xe4\x1a\x2b\x01\xba\x57\x75\xfa\x0e\x3b\x92\xd5\x9a\xee\x63\x3b\x6c\xd7\xb3\xe2\xed\x1b\x19\xbb\xcc\xcc\x70\x1e\xeb\xd5\x46\x19\xa9\xe3\x43\x33\x0c\x05\xf9\x6d\x75\xfd\x85\x30\xe0\xa5\xd4\x3a\x8d\x21\x0c\x99\x1d\x5f\x23\x6d\xdb\x13\x45\xc6\xe3\xfd\x43\x91\xf1\x5c\xff\xb6\xc8\x30\x4b\x81\x3d\xc8\x08\x03\x8f\xc9\xcd\xfa\x2a\x7d\x59\x30\x65\x4e\xd8\xf1\x85\xd7\x4a\x23\xf8\x74\x8a\x27\x01\xbd\x32\x1b\x22\x44\xcc\x6f\xeb\x04\x39\x3c\x39\x0f\x21\x28\x6b\xa2\x0c\x27\x66\x47\xb3\x86\xcd\x4f\x4f\x6e\xb6\x4f\x87\xf7\x2e\x0c\x77\x12\x1b\x23\xef\x07\xdd\xcc\xa3\x8a\x41\x93\xb0\x5c\x02\xd7\x60\x36\xd8\x7c\x9f\x2d\x97\x70\x9b\xdd\x0d\xa6\x98\x66\xe8\xb7\x5a\xca\x82\xd8\x7b\x08\xce\x9a\x00\x79\x5c\x6d\x35\xe6\xa4\x82\x5a\x6e\x35\x26\x4e\xee\xb4\x95\x15\x69\x59\x40\x89\xdb\xf0\xc9\x56\x90\x5f\x64\x59\x5b\xa8\x3a\x75\x47\x1a\x1a\xf8\xe8\x17\x2e\x52\x8e\xf6\xf3\xf3\x45\x96\xb1\xc0\x0f\x2c\xbd\xff\x60\x3d\x3f\x9f\x26\x2d\x40\x07\xe8\x7a\xe0\xc5\xed\x1d\x53\x82\x2c\xca\x9b\xf5\x55\x62\x2c\x26\x8d\x34\x95\x86\x2a\x4f\x16\x25\x79\xff\xb2\x31\x45\x1c\x71\xea\xbd\x20\xbf\x9b\x45\xf9\xa3\x44\xd9\x85\xb1\xb8\xf3\xaa\x21\x1f\xc9\x22\x90\x9c\x2c\x56\x84\xf9\x7e\x48\x91\xda\x1a\xfc\x10\x70\xa7\x21\x4f\x14\x4a\xad\xca\x22\x29\xad\xb6\x3e\x4f\x3c\x54\x45\xd2\xfb\xd5\x5f\x90\x27\xe7\x17\xee\x89\xb0\x19\xc2\x58\x7f\x2f\xf5\x11\xb1\xf1\x00\xe6\x15\xa6\x20\x0c\x29\x65\x9e\x6f\x4d\x68\x54\x8d\xff\x13\x29\x65\xa5\x35\xc1\x6a\xe0\xda\x6e\x86\x89\x3a\xec\xb0\x5b\x75\xc7\x4b\x6b\x4a\x89\xa9\xa7\xb4\x25\xf6\x8f\x6f\x50\xe2\x78\x53\xe1\x63\x2f\x80\x00\x38\x28\xe7\x12\x64\x05\x3e\x25\x9f\xac\x41\x30\xf8\xe1\xcb\xce\x01\x61\x24\xa6\x55\xa5\x8c\x1a\x38\xfb\x16\xac\x21\x74\x26\xe7\x5f\xae\x57\xbf\xf2\xbe\x0d\xaa\xde\xa5\x81\x52\x9a\x4f\xfd\x61\x90\x78\x37\x3c\xc5\x9b\x03\x85\x0d\xc3\xf3\x74\x54\x54\x34\x3b\x4c\x2a\xf1\xee\x9c\x1d\x06\x42\xb7\x6e\x53\x5a\x3c\x2a\x53\xd9\x47\xfe\xe2\x17\x9d\x9b\x9d\x00\x3f\xf7\x3a\x3f\xf5\xc1\x32\x60\x76\x98\x69\x13\x4d\x0f\x7f\xe7\xf1\xed\x43\x6f\x4f\xd4\x62\x8f\x9f\x69\xac\xe6\xef\x00\x00\x00\xff\xff\x06\xb0\x28\x65\xd1\x07\x00\x00")

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

	info := bindataFileInfo{name: "script/script.min.js", size: 2001, mode: os.FileMode(420), modTime: time.Unix(1479347376, 0)}
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

