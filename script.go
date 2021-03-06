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

var _scriptInterceptorTs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x57\x51\x6f\xdb\x36\x17\x7d\xf7\xaf\xb8\x0f\x75\x25\xe3\x73\x15\xb7\x4f\x1f\x64\xa8\x45\xd1\x6d\xc8\x86\xb4\x1e\x9c\x06\x28\x90\x06\x30\x23\x5d\xdb\x2a\x68\x52\x23\xaf\x96\x78\x89\xff\xfb\x40\x91\xa2\x28\xc5\x71\x07\x74\x8f\x7b\x31\x64\xf1\xde\xc3\x73\x0f\x0f\x2f\xa9\x9c\x33\xad\xe1\xcb\xf9\xf2\x57\x41\xa8\x72\xac\x48\x2a\x78\x18\x01\x54\xaa\xfc\x93\x11\x82\x26\x46\x65\x0e\x0a\x59\x21\x05\xdf\xc3\xfd\xa2\x42\x01\x19\x7c\xf9\x78\x71\x4e\x54\x2d\xf1\x8f\x1a\x35\x25\x95\x92\x24\x69\x5f\x61\x22\x2b\x14\xf3\x53\x00\x97\x28\x8a\x53\x00\x1a\x45\x71\x04\xc0\xe0\x2e\x96\x90\xc1\x9a\x71\x8d\x47\x02\x4c\x5e\x18\x10\x44\xec\x18\xe5\x5b\x54\x3a\x85\x8f\xf6\xe9\xfa\x06\x32\xb8\xbe\xb1\x41\xf5\x2d\x2f\x73\xa0\xad\x92\x77\x0b\xf1\x49\x36\x21\x29\xdc\x4a\xc9\x91\x99\x52\x49\xd5\x16\x2e\x97\x42\x93\xaa\x73\x92\x2a\x6e\xa1\x89\xa9\x0d\xd2\xd5\xf2\x22\x05\x4d\xaa\x14\x1b\xc8\x20\xda\x12\x55\xe9\xd9\x19\x97\x39\xe3\x5b\xa9\x29\x7d\xfd\xff\xd9\x6c\x76\x16\x4d\xe0\x01\x0e\x06\xa9\x21\x14\x3b\x5a\x9e\xd5\xa4\x51\x1e\x80\xb6\xa5\x4e\x5a\xce\x49\x55\x6b\x1f\x3a\x31\x75\x77\x08\xef\x39\x8f\x8f\xd4\x76\x02\x27\x49\xfc\x8b\x0e\x2b\xe7\xc8\x54\x7c\x34\xcb\xc9\xe4\x02\x15\x6a\x92\x0a\x7d\xe8\x49\x13\x98\x35\xee\xf9\x2a\x69\xbc\xd3\x24\x02\xbc\x3b\x35\x98\x7e\xcf\x5e\xa7\x67\xd7\xce\x61\x83\x09\x8c\xf1\x9e\x9d\x3d\x18\x3c\x31\x7b\xeb\xcd\x46\x0d\xef\xec\xff\x44\xe9\x44\x29\xdb\xfc\xbe\x9f\xfc\x6b\x53\x50\x7c\xb2\x90\xc6\x97\xc3\x34\xc3\xe4\xf9\x34\xc3\xa0\xb3\x73\xbb\x35\xfb\x73\x1a\xe8\x14\x7e\xa9\x45\x4e\xa5\x14\x2d\xb9\x72\x0d\xf1\xa0\x6a\xdb\x68\xda\x00\x63\x7a\xaa\x95\xb0\xa4\x0e\x56\xe4\xa6\x11\x80\x46\xbe\x36\xed\x61\x5b\xea\xef\x89\xef\x96\x7e\xed\x66\x87\x78\x87\xb4\x95\x45\xdb\x33\xa6\x50\x2b\xde\xfe\xe9\x66\x6e\x24\x90\xaa\xdc\x94\x82\xf1\xab\xe5\x05\x64\x26\x6e\xee\x46\x2d\x84\x69\x37\xbf\x2f\x2e\x3f\x47\x8e\x02\x34\x8d\x32\xc9\x19\xe7\xb1\xc9\x9f\xba\xb8\x69\x43\x37\xf1\xed\xca\xa9\x7c\x78\x5e\xb5\x46\x72\xa3\xec\x3f\x52\xcd\x76\xdf\x7f\x57\x35\x67\xd9\x4e\xb5\x82\x11\x4b\x81\x89\xbd\x99\xc7\x4d\x64\x51\x2b\xa6\x34\x16\x56\x23\x81\x77\x70\xb5\xbc\x88\x87\xf2\x4d\xe6\xfd\x1c\x0e\x99\x15\xc5\xb7\xc8\x75\xc9\x09\x55\x1c\xef\x26\x90\xbd\xf5\xa5\xd8\x7a\x0d\x23\xb9\x86\x5d\x82\xf7\x95\x42\xad\x0d\xa1\x2c\x83\xc8\x2e\x5a\x34\x09\xc2\x61\x10\xd5\x30\x5a\xe2\xe6\xe7\xfb\x2a\x0e\x47\x26\x73\x9f\x72\x18\xf9\x47\x4b\xce\x72\x32\xfd\xd7\x97\x96\x54\x8c\xb6\x82\xed\x30\x71\xc7\x47\x1f\xaa\x43\xb0\xda\x43\x1c\x12\x72\x70\x2f\x5f\x3e\x7d\x99\x70\x14\x1b\xda\xc2\x5b\x98\x1d\x1d\xbf\x9e\xdd\xb4\x21\xd9\x51\x3a\x76\xd0\x27\xfa\xb2\x0e\x43\xc5\x2b\xb6\xe7\x92\x99\x35\x7d\x30\x6d\xb3\x92\x42\x63\x6a\x9f\x6b\x4e\x29\x44\x05\xae\x59\xcd\x7d\x60\x04\x87\x69\xd3\x66\x6b\xfd\x41\x16\x98\xc2\x9b\xd9\xcc\x58\x76\xd4\x2d\x0b\x0f\xd8\x87\x6b\xe0\x10\x92\x2e\x1b\x32\xe0\xa6\x94\xe0\xcd\xe3\xa3\x41\x9c\x3f\x49\x6a\xb9\xb5\x29\xfe\xff\xe3\xe3\x69\xb6\xbe\x72\x40\xae\x31\xa0\x63\xeb\x37\xbf\x92\xe3\x7b\xb5\xd1\x8d\x8d\xfd\x35\xa4\x0d\xe3\x48\xb0\x43\xad\xd9\xc6\xcc\xbd\x1a\xe7\xc6\xd1\x42\x12\x6c\x99\x28\x38\x16\x29\x8c\xf3\x17\x0f\x43\x63\x1f\x56\xf3\x9e\x55\xcd\x36\x19\xf8\xd1\x61\xfe\x2f\x83\xd5\x57\x31\xce\x7f\x6a\x36\xd2\x38\x0f\x12\x7b\x41\xce\xec\x06\x28\x34\x39\xbc\x83\xd5\x58\xaf\x20\x85\xd5\x78\xd1\xcb\x0d\x2a\xb3\x17\x8d\x60\x0c\x60\xb5\x96\x82\x5e\x69\xda\x73\x4c\xa1\x24\xc6\xcb\x7c\x0e\xb9\xe4\x52\xa5\xa0\xb0\x98\x83\x1d\x2f\xff\xc2\x14\x5e\xbf\xa9\xee\x57\xd3\xe7\xd3\x85\x54\x3b\xc6\x7d\xfa\x46\x21\x8a\x27\x00\xf3\x01\x82\x29\x24\x7c\xd3\xdb\x79\xc7\x6a\xa8\x85\xde\x96\x6b\x0a\xcb\xf8\xa1\x22\x7e\xac\x84\x80\xae\x23\x99\x70\xb9\x89\xdd\x82\x4d\x21\x49\x92\x80\x7c\xb7\xeb\xac\x1d\x6c\xe3\xef\x5d\x6f\xfb\xee\x68\xc6\x60\xf5\xc4\x6b\x47\x9c\xf6\x55\x58\xeb\xbc\x78\x30\x92\x86\xc6\x6b\x75\x3c\x84\xbb\xd3\xf9\xc8\xef\xfb\x0c\x22\x79\xfb\x0d\x73\x8a\x82\x0e\xee\x0e\x3a\x8d\xe4\x8e\x80\x73\x64\x05\xaa\x38\xfa\x20\x05\xa1\xa0\x57\x9f\xf7\x15\x46\x53\x88\x58\x55\xf1\x32\x67\xe6\x2c\x38\xfb\xa6\xa5\x88\x02\x5d\xcc\x59\x11\x9e\x7a\xbf\x5d\x2e\x3e\x25\xd6\xb6\xe5\x7a\x1f\x3b\x06\x93\x4e\x9b\x27\x5b\x74\x88\xd0\xa6\xf8\x3d\x1d\x1e\x98\x87\xd1\x28\xbe\x2b\x45\x21\xef\x80\xe9\xe6\x44\x4a\x06\x1f\x51\xc3\xab\xd6\x7c\x34\xb2\x1f\x5b\xee\x9a\xde\xcc\x1c\x7e\x51\x34\xf0\xee\x63\xa4\x6b\xec\xa9\x3b\x38\xe0\xb1\xbd\x2b\x84\x71\x5d\x13\x65\x62\xdf\x1b\x09\xbb\xa6\xa8\x77\xb7\xa8\xcc\xb0\xfd\x00\x39\x8c\xfe\x0e\x00\x00\xff\xff\x9a\x23\xfc\x1c\xf7\x0d\x00\x00")

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

	info := bindataFileInfo{name: "script/interceptor.ts", size: 3575, mode: os.FileMode(420), modTime: time.Unix(1479841470, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _scriptScriptMinJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x56\x51\x6f\xe3\x36\x0c\xfe\x2b\xae\x80\x06\x16\xaa\x73\xdd\x3e\x0d\x0e\xb4\x62\xb8\x0d\xe8\x86\xf6\x32\xa4\x57\xe0\x80\x5e\x1f\x34\x9b\x89\x75\x50\x45\x4d\x62\xae\xcd\x52\xff\xf7\x41\xb6\x13\xc7\xcd\x52\xec\x69\xc0\x9e\x62\x8a\xfc\xc8\x4f\xa2\xf8\x45\xdf\x95\x4f\xbe\x5c\xcf\x7f\xb5\x04\xbe\x04\x47\xe8\xe5\x62\x65\x4b\xd2\x68\x53\xbe\xd9\x7e\x26\x94\x12\xdf\x7c\x47\x5d\x25\xb9\x94\x92\x26\x93\x94\x24\xab\x89\x5c\x71\x7e\x6e\xb0\x54\xa6\xc6\x40\xc5\xc5\x0f\x79\x9e\x9f\x33\x2e\xa8\xd6\x21\x23\xe5\x97\x40\xf7\xf3\x1b\x49\xdd\xc2\x93\xa2\xb2\x06\x1f\xe4\xc3\x63\x1f\x51\x7b\x7c\x9e\xd9\x4f\x78\x1b\x3d\xf2\x24\x6f\x3c\xd0\xca\xdb\x84\x32\xe7\x91\x90\xd6\x0e\x3a\xd4\x40\x8a\xf8\x66\x94\x2c\x73\xab\x50\xa7\xc4\x1b\x71\x00\xfa\xc9\x98\x11\x2e\x05\x39\x82\xf2\x16\x9b\x29\xe7\xcc\x3a\x05\x41\x7c\x1a\x4f\x03\xc6\x99\x4a\x03\x6a\x74\x26\x6f\xb7\x32\x0e\xf7\x10\x08\x3d\xec\x03\xbe\xdc\xde\x5c\x13\xb9\x39\xfc\xb9\x82\xb0\x1f\x8b\x0e\xac\xa4\xec\x65\xe6\xc0\x5e\xf5\xbf\xc5\xbb\xd1\xe2\xa8\x37\x80\xad\x62\xae\x3b\xb0\xd5\x55\xff\x7b\x3c\x57\x8c\x8e\xbc\xff\x5f\x6c\x87\x15\xbd\xbd\xae\x07\x8d\xd9\x79\x22\xbd\xf4\x5d\x7a\xfd\x35\xdd\x21\x22\x89\xe3\x88\x48\x82\x1f\x61\x11\x6b\x0d\x4c\x80\x6f\xf4\x22\x3d\xa1\xb6\xc8\x6c\xce\x37\xf1\x56\x61\x7b\xf5\xa6\xef\x9f\xee\x70\x59\x85\xed\xf7\x83\x5e\x2f\xb5\x55\x26\x8e\x91\x15\x24\xd9\xef\xb3\xbb\xcf\x4c\x40\x56\x2a\x63\xd2\x18\x22\x48\xe0\x30\x6b\xbc\x69\x8e\x90\x8c\xdb\xfb\x07\x92\x71\x5f\xff\x96\x64\x18\xa5\xa0\x0e\x64\xa5\x85\xe7\xe4\x7e\x7e\x93\xbe\x25\xcc\x85\x93\x38\x0c\xea\x42\x1b\x02\x9f\xee\xe3\x59\x20\xaf\xed\x92\x49\x19\xf3\xe3\x22\xa1\x0c\x5e\x9c\x87\x10\x34\xda\x28\x32\x7b\x66\x5b\x66\x0e\xcb\x5f\x5e\xdc\x68\x9d\xf7\x63\x2b\x6d\xe6\x14\xd5\x56\x3d\xf5\xe3\x3f\x8e\x9a\xf6\xd2\x02\x93\x09\x64\x06\xec\x92\xea\x1f\xf3\xc9\x04\x1e\xf2\xc7\xde\x94\xfb\x19\xba\xa5\x86\x8b\x20\x37\x1e\x82\x43\x1b\xa0\x88\x5f\x2b\x43\x05\xab\x60\xa1\x56\x86\x12\xa7\xd6\x06\x55\xc5\x1a\x11\x48\xd1\x2a\x7c\xc4\x0a\x8a\xcb\x3c\x6f\xa6\x7a\x91\xba\x5d\x19\x1e\xb2\xc1\x2f\x5d\x2c\x39\xd8\xaf\xaf\x97\x79\x2e\x42\xb6\xad\xd2\xf9\xb7\xd6\xeb\xeb\xf1\xa2\x53\x30\x01\xda\x1e\xf8\x28\xa9\x5a\xb2\xd3\xf2\x7e\x7e\x93\x58\xa4\xa4\x56\xb6\x32\x50\x15\xc9\x69\xc9\xce\xde\x36\x26\xb2\x8b\x1a\xae\xcf\x24\xfb\x6a\x4f\xcb\x9f\x15\xa9\x36\x52\xc4\x95\x83\x9e\x5c\xb1\xd3\xc0\x0a\x76\x3a\x63\xc2\x77\x72\xcb\x16\x68\xe9\x43\xa0\xb5\x81\x22\xd1\xa4\x8c\x2e\xa7\x49\x89\x06\x7d\x91\x78\xa8\xa6\x49\xe7\xd7\x7f\x41\x91\x5c\x5c\xba\x17\x26\x46\x08\x8b\xfe\x49\x99\x1d\x62\xe9\x01\xec\x01\x66\xca\x04\x71\x2e\x7c\xb6\xb2\xa1\xd6\x0b\xfa\x8f\x8a\x72\x51\xa2\x0d\x68\x20\x33\xb8\xec\xff\x1b\xfa\x15\xf1\xa0\x1f\xb3\x12\x6d\xa9\x28\xf5\x9c\xc7\xb1\x1b\xfd\x81\xf1\xd6\x64\x07\x2d\x38\x6c\xc0\x19\xfb\x6a\xbb\x43\x67\x67\xd4\x30\xfc\xe3\x1b\x94\x34\x9c\x78\xb8\xea\x66\x29\x00\xf5\x43\x78\x0d\xaa\x02\x9f\xb2\x8f\x68\x09\x2c\x7d\xf8\xbc\x76\xc0\x04\x8b\xf4\x74\xa9\xe2\x38\x9d\x7f\x0b\x68\x19\x1f\x29\xc3\x6f\x77\xb3\x4f\x59\xd7\x4e\xbd\x58\xa7\x81\x73\x5e\xec\xfb\x43\xaf\x16\xad\x96\xcb\xf7\xb5\xbc\x57\xe9\xe3\x51\x51\x1c\xc4\x56\xf4\xe4\xc9\x85\xd8\x6a\x4b\xfb\xdd\xa4\x7c\xfa\xac\x6d\x85\xcf\xd9\x9b\xe7\xc6\xd8\x6c\x67\xf9\xb6\x93\x8c\x63\x2f\x11\x01\x02\x7b\x79\xdc\x93\x87\xfe\x91\xb1\x1b\x23\xe8\xec\xbd\xc1\xc3\xdd\xf3\x22\xb2\xf9\x3b\x00\x00\xff\xff\xc6\xf2\x79\x8d\xfa\x08\x00\x00")

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

	info := bindataFileInfo{name: "script/script.min.js", size: 2298, mode: os.FileMode(420), modTime: time.Unix(1479841473, 0)}
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

