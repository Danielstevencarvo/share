// Code generated by go-bindata. DO NOT EDIT.
// sources:
// frontend/app.js
// frontend/index.html

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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataFrontendAppjs = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x55\x51\x8f\xe2\x36\x10\x7e\xcf\xaf\x98\x07\x24\x3b\x57\x36\xf4\xd4" +
	"\x7b\x22\xe5\xa4\xea\x8e\xbb\x43\xda\x65\x2b\x96\x7d\xab\xba\x32\xce\x24\x71\xd7\xd8\xc8\x76\x1a\x56\x15\xff\xbd" +
	"\xb2\x49\x42\x02\x2c\xf7\x94\xc4\x33\xf3\xcd\xf8\x9b\x99\x2f\x12\x1d\xe4\x42\xe2\xf3\xea\x1e\x66\xa0\x2a\x29\xd3" +
	"\x88\x6b\x65\x1d\x3c\x3c\x7d\x7f\x59\x3e\xae\x5f\xbe\x3d\x3e\x2f\xbf\xc2\x0c\xc8\x37\x21\x11\x94\x76\x90\xeb\x4a" +
	"\x65\xe4\xdc\xef\xcf\xf9\xea\x61\xb1\x5e\xcf\x83\xef\x52\x3b\x60\x52\xea\x1a\x33\x70\x1a\x18\xe7\x68\x6d\x48\x34" +
	"\x88\xfb\x3e\x5f\xce\x57\x8b\x2f\x2f\xf3\xd5\xca\x47\x3d\xe9\x2d\xba\x52\xa8\x02\x6a\x54\x0e\x6a\xa3\x55\x41\xd2" +
	"\x28\x9a\x7c\x80\x42\xea\x0d\x93\x30\x9a\xe6\x4c\x5a\x84\x0f\x93\x28\xe2\x92\x59\x0b\x4f\x25\x33\x08\xff\x45\x00" +
	"\x42\x09\x47\xe3\xf0\x0a\x30\xa2\xb5\x50\x99\xae\xe3\x64\x23\x54\x46\x49\xc9\x6c\xc9\x4b\xa6\x0a\x24\x63\xa0\x18" +
	"\xc3\xec\x73\xe3\x09\xe0\x4a\x61\x13\xef\x70\xaf\x59\x46\xe3\x34\x1c\x1f\x9a\xe7\x15\xe3\x21\x8a\x00\x70\xbb\xc1" +
	"\xcc\x53\xb2\x50\xb9\xa6\xfe\x66\xfe\x06\xa4\xcd\x2f\x72\x68\x0e\x67\xfd\xe3\x0e\x4f\x65\x12\xe7\xc6\x68\xf3\x80" +
	"\xd6\xb2\x02\xe9\x80\xee\xb6\x86\x28\x3c\x4e\xfd\xf1\x6f\x47\xd3\x28\x61\xff\xb0\x7d\x48\x31\xee\xb0\x3d\x7b\x3a" +
	"\x9b\x02\xf9\x31\xff\xe3\x2b\x19\x37\xa7\xb6\x0a\xec\x4f\x81\x66\xcc\xb1\x31\x58\xc7\x5c\x65\xc7\xb0\x2f\xcd\x80" +
	"\x86\x61\x6d\xfe\x7a\x57\x02\xd2\xc6\xf9\xd0\xa2\xa3\xbf\xc4\x14\xe8\xbe\x34\xad\xe7\x0d\x54\xef\x3c\x70\x3d\xc3" +
	"\x3b\x9c\x08\xbe\x5d\x47\x83\xef\xc7\xd7\xbd\xed\x3c\xf7\xfb\xd2\x24\x05\xba\x15\xda\x9d\x56\x16\x7f\x20\xcb\xd0" +
	"\x50\xf2\x45\x2b\x87\xca\xdd\xad\xdf\x76\x48\xe2\x34\x6a\xa6\x83\x24\x5c\x2b\xc7\x84\x42\x43\xe2\xa4\x14\x19\xb6" +
	"\x8d\xf7\x36\x4f\xab\x62\x5b\x24\x71\xe2\x70\xef\x68\xd3\x81\xc4\x56\x1b\xeb\x8c\x50\x45\x77\x22\x99\x75\x0b\x95" +
	"\xe1\xfe\x31\xa7\x64\x42\x62\xf8\x05\x3e\xc6\x6d\x1a\x3f\x04\xbe\xba\x64\xcb\x1c\x2f\xe9\xe4\x6f\xb1\x65\x05\xfe" +
	"\x35\x99\xc4\xa7\x79\x68\xd3\x49\xa1\x5e\xef\xac\xe1\x24\x4e\x98\x73\x86\x12\xff\x3e\x6e\x7b\xdf\xf1\xe4\xdd\x6d" +
	"\xa9\xeb\xbb\x00\x45\xe2\xf0\x41\x3b\xb3\x41\x57\x19\xd5\x4e\xcf\xd5\x1a\xfe\x15\x19\xea\x61\x0d\x9e\x45\x6b\x38" +
	"\xcc\x3c\xfc\xef\x56\x57\x86\xe3\x67\xd2\x81\x5a\xc3\x6f\xd6\x14\xec\xbb\x1d\xaa\x6c\xad\xe9\x88\x92\x90\x81\xc4" +
	"\x97\x35\x37\x86\x9b\x35\x47\x17\xa4\x94\x06\xf3\x8e\x95\xf0\x71\x5e\x42\x97\xa0\x40\x85\x46\xf0\x41\x8a\xfe\x30" +
	"\x5d\x8c\x5f\x6f\x8c\xb6\xc7\x45\x0c\x5b\x7c\x84\xb5\xb5\x70\xbc\x0c\xa3\x9d\x0c\xfd\x01\x38\xb3\x08\x9f\x7e\xfd" +
	"\x34\xed\xc6\xfc\x14\x3f\x58\xe5\xb4\x73\xd8\x18\x64\xaf\xe9\x30\xfe\xb7\x5b\xf1\x9d\xa2\xbe\x83\x91\x61\xce\x2a" +
	"\xe9\xde\x83\xe8\x89\xeb\x55\x80\x86\xeb\xf7\x04\xa9\x01\xbb\xca\xe1\xb9\x4f\x27\xba\x24\x09\x7a\xd0\xee\x4d\x1f" +
	"\xe3\xe7\x5b\x17\x5a\xd8\xc6\x5f\x34\xb0\x55\xe0\x5e\xcb\x1a\xd5\x3d\x2a\x7d\x22\x35\x67\x4e\x68\x15\xd4\xba\xb7" +
	"\xaa\x1f\xfb\x4a\x7e\x29\xdb\x4d\x8e\x43\x14\x8d\x68\x5e\x29\xee\x21\x9a\x2c\x61\x2f\xc2\xff\x65\x06\x0a\xeb\xe3" +
	"\xbf\xe6\x58\x54\x38\x4e\x8e\xbf\x9c\x34\xf2\xb2\xf5\x7f\x00\x00\x00\xff\xff\x35\x81\xcc\x08\x4a\x07\x00\x00")

func bindataFrontendAppjsBytes() ([]byte, error) {
	return bindataRead(
		_bindataFrontendAppjs,
		"frontend/app.js",
	)
}



func bindataFrontendAppjs() (*asset, error) {
	bytes, err := bindataFrontendAppjsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "frontend/app.js",
		size: 1866,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataFrontendIndexhtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x96\x5d\x57\xdb\x38\x13\xc7\xaf\xe1\x53\xa8\x7e\x2e\x41\x56\xde\x08" +
	"\x81\xda\x39\x4f\x20\x65\x0b\x34\xa4\x81\xd2\x6d\x7b\xa7\x58\x63\x5b\x89\x2c\x19\x49\x8e\x13\x7a\xf6\xbb\xef\xb1" +
	"\x13\xe7\xa5\x2d\x14\xf6\xec\xe6\xc6\xd1\xcb\x8c\xfe\xbf\x99\x91\xc7\xde\x1b\xa6\x02\xbb\x48\x01\xc5\x36\x11\xdd" +
	"\x7d\xaf\x78\x20\x41\x65\xe4\x3b\x20\x9d\xee\x3e\x42\x5e\x0c\x94\x15\x7f\x10\xf2\xde\x60\x8c\x6e\xe1\x21\xe3\x1a" +
	"\x18\x4a\xc0\x52\x64\x69\x64\x10\xc6\xab\xf5\x72\x2a\x88\xa9\x36\x60\x7d\x27\xb3\x21\xee\x38\xdb\x4b\xb1\xb5\x29" +
	"\x2e\xec\x67\xbe\xf3\x05\xdf\xf7\xf0\xb9\x4a\x52\x6a\xf9\x58\x80\x83\x02\x25\x2d\x48\xeb\x3b\x97\xef\x7c\x60\x11" +
	"\xec\x58\x4a\x9a\x80\xef\xcc\x38\xe4\xa9\xd2\x76\x6b\x73\xce\x99\x8d\x7d\x06\x33\x1e\x00\x2e\x07\x87\x88\x4b\x6e" +
	"\x39\x15\xd8\x04\x54\x80\x5f\x3f\x44\x26\xd6\x5c\x4e\xb1\x55\x38\xe4\xd6\x97\xca\xe9\xee\x6f\x80\xce\x94\xb2\xc6" +
	"\x6a\x9a\xa2\xf3\xbb\xbb\x0d\x8b\xe0\x72\x8a\x34\x08\xdf\x31\x76\x21\xc0\xc4\x00\xd6\x41\xb1\x86\xd0\x77\x0a\x0e" +
	"\x73\x4a\x48\xc0\xe4\xc4\xb8\x81\x50\x19\x0b\x05\xd5\xe0\x06\x2a\x21\x74\x42\xe7\x44\xf0\xb1\x21\x36\xe7\xd6\x82" +
	"\xc6\xe3\xea\x04\xd2\x72\x6b\x6e\x8d\x04\xc6\x90\xf5\x9c\x9b\x70\xe9\x06\xc6\x38\xe5\xb1\xcb\x1f\x97\x16\x22\xcd" +
	"\xed\xc2\x77\x4c\x4c\x1b\x47\x6d\xfc\xa1\xd7\x39\x79\x3c\x98\x9e\x84\x93\x68\x70\x3d\x22\xd3\x87\xd6\x70\xd8\x18" +
	"\xe8\xb0\xf3\x59\xd8\xaf\x89\x20\x9f\xdf\xdd\x1d\xdc\x46\xb5\x30\x6e\xd4\x7c\x07\x05\x5a\x19\xa3\x34\x8f\xb8\xf4" +
	"\x1d\x2a\x95\x5c\x24\x2a\x33\x0e\x22\xff\x1e\x5b\xa9\x3f\xa7\x36\x88\x57\x50\x8c\xea\xa9\x58\xbc\x96\x6b\x1c\xf7" +
	"\x5b\xa4\xcf\xce\x7b\xe9\x4d\xef\xab\x18\xa6\xb3\x6b\x63\xee\x69\x23\xee\xd7\xfa\x59\x67\xae\xa3\xe1\xcc\x7e\x6d" +
	"\xe5\x8d\x23\x13\x0f\x9e\xe7\x7a\x39\x58\x66\xc0\x0d\x95\xb4\x34\x07\xa3\x92\x25\x97\x06\x01\xd4\x80\x21\xb3\x23" +
	"\xb7\xe6\xd6\x97\x29\xa2\x42\x3c\x0f\xd0\xec\xb4\xf0\x01\xab\x7d\xec\x34\xe5\xc9\x94\x8e\x06\xe7\xf9\xa4\x73\xd1" +
	"\xba\xbd\x3a\x6b\xb7\xed\xe3\x65\x3e\xbc\x4e\x34\x1b\xb7\xda\x07\xa9\xd2\x7d\x32\x9c\xe9\xab\x83\xe6\xf1\x9f\x0f" +
	"\x97\x83\xe3\x7b\x75\x66\xf3\xf7\xc3\xf6\x8d\x88\x9e\x84\xaa\x90\x2c\xb7\x02\xba\x77\x31\xd5\x80\xfe\x8f\x44\xf6" +
	"\xc8\x43\xd0\x2e\x57\x1e\x59\xae\xac\xb6\x95\xbc\xdd\x95\xd4\xb8\x8e\xbe\x23\x0b\x73\x8b\xa9\xe0\x91\x3c\x45\x01" +
	"\x48\x0b\xfa\x2d\x2a\xb8\xb1\xe1\x8f\x70\x8a\x8e\x8f\xd2\xf9\x5b\xf4\x57\x65\xd1\xf8\xbd\x45\xa3\xb5\x6d\xe1\x16" +
	"\x57\x90\x72\x09\x1a\x7d\x47\x09\xd5\x11\x97\xd8\xaa\xf4\x14\x35\x6b\x3b\xdb\x4c\xac\x72\x1c\x81\x04\xcd\x83\xc3" +
	"\xd5\x90\x27\x34\x82\x6a\x30\xe3\x0c\x54\x35\x00\xad\x55\xe1\x90\x71\x93\x0a\xba\x38\x45\x52\x49\xa8\xbc\x79\x64" +
	"\x4d\xe9\x91\xea\xb5\xe4\x8d\x15\x5b\x74\xf7\xf7\xf6\x3c\x49\x67\x28\x10\xd4\x18\xdf\x91\x74\x36\xa6\x1a\x2d\x1f" +
	"\x18\xe6\x29\x95\x0c\x8b\xa8\x9a\x28\x6a\x15\x8d\x23\x9c\x6a\x9e\x50\xbd\x70\xaa\xb8\x79\x74\xd7\x01\x1e\x6b\x2a" +
	"\x59\x55\x40\xff\x73\xba\x1e\xaf\x36\x84\xd4\xa0\x90\x62\x53\xe4\x05\x53\x61\xb1\x79\xc8\xa8\x06\xa7\xeb\x11\xde" +
	"\xdd\xaa\x98\x5f\x25\x8e\x96\x72\x89\xa4\xb3\x2a\x7b\x8c\xaf\xa5\x6f\xc2\x5a\x06\x44\x28\xca\xb8\x8c\x36\x12\xe3" +
	"\xfa\xc6\xfd\xcf\x72\x52\x2e\x0b\xd3\x90\xe2\x34\x13\x06\x1c\x44\x35\xa7\x38\xe6\x8c\x81\xf4\x1d\xab\xb3\x5d\x85" +
	"\x1e\xa9\xdc\x79\x84\xf1\x17\xc8\x59\x25\x72\x57\xce\xde\xde\xde\xde\x4f\x52\xca\x37\x07\x66\x2a\x97\x05\x42\x11" +
	"\xa2\x27\x4e\x2e\x9c\x34\xba\x1e\x5d\x47\x79\xed\x88\x0b\x28\x2e\x34\x2e\x16\x50\x31\x2a\xba\x80\xd3\xcd\x64\x59" +
	"\xf9\xac\x08\xa4\x47\xe2\xc6\xab\x00\xca\xd2\x5b\x56\xfa\xb2\xc4\xb7\x93\xff\xac\x02\x67\x3b\xee\x49\x84\x8c\x0e" +
	"\x7c\x67\xbd\x95\x27\x11\x0e\x45\xc6\x19\x5a\x1b\x19\xbd\x15\xa7\x22\xe9\xaf\x90\x59\x5e\x8a\x8d\xf5\xd6\x46\x48" +
	"\xc6\xc0\xb0\x06\x93\x2a\x69\xf8\x0c\xd0\x8f\x13\xb8\xde\x1e\x2f\x4e\xb6\xd5\x96\xce\xca\x96\xa9\x95\x30\x5b\x0b" +
	"\xa4\x5c\xd9\x68\x2c\xa5\xbd\x42\x65\x79\x5b\x5f\x50\x0b\x30\x0f\x04\x4d\xa8\xe5\x4a\xe2\x80\xeb\x40\x3c\x55\x85" +
	"\x65\x2d\xac\x49\x97\xde\x9f\x48\x71\xd1\xba\x27\xa3\x0c\xf4\x02\x85\x5c\x1b\x7b\x88\x6c\x0c\x12\x7d\x54\x69\x0a" +
	"\xda\x9d\x98\xd5\x78\xd3\xdd\xaf\xb6\x9a\xbb\x09\x34\x4f\xed\x32\x85\x2f\x6c\x78\x93\x87\xe2\x2c\xd2\x74\x9b\x6e" +
	"\x7d\x35\x28\x5b\xdc\x64\xa7\x41\xfc\xaa\xc7\x5d\x44\xe9\xf9\x98\x5c\x5f\x8d\xc4\x87\x9b\x70\x98\x9d\xd4\x2d\x6d" +
	"\x36\x14\xb9\x19\x7c\x9b\x0b\x9b\xdf\xaa\xce\xc8\x26\xd3\xc1\x2d\xeb\x65\x9d\xa7\x7b\x5c\xd7\x23\x4b\xd1\xff\x9c" +
	"\x20\xad\x42\x43\xea\x6e\xbd\xe5\xd6\xab\x89\x97\x51\xdc\x7d\x6a\x0c\x20\xd0\x5f\xf4\x15\xed\x99\x87\x24\x4c\xa7" +
	"\x27\x5f\x6e\x47\x97\x9f\x44\x5f\x2d\x06\xc9\xbd\x8d\xae\xcf\xde\x49\xd6\xe7\xe6\x2e\xf8\x4f\x29\x9e\xfa\xa8\x9a" +
	"\xfc\xf8\x4d\xf5\x7b\xa4\xa3\x83\x5a\xe3\x31\x3b\xba\xbf\xff\x30\x9a\x0e\x8f\xf3\xfa\x1f\x97\xba\x3d\x0b\xed\x79" +
	"\x34\x08\x2f\xd8\xb7\xa0\xf7\x1e\xfa\xf6\x42\x5e\x7f\x33\x67\xe6\x45\x48\x3f\x33\xd1\x34\x2d\x74\xec\x60\x7b\x64" +
	"\xd9\xb1\x3c\xb2\xfc\xe6\xfe\x3b\x00\x00\xff\xff\xe0\x13\xe8\x73\x84\x0b\x00\x00")

func bindataFrontendIndexhtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataFrontendIndexhtml,
		"frontend/index.html",
	)
}



func bindataFrontendIndexhtml() (*asset, error) {
	bytes, err := bindataFrontendIndexhtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "frontend/index.html",
		size: 2948,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"frontend/app.js":     bindataFrontendAppjs,
	"frontend/index.html": bindataFrontendIndexhtml,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"frontend": {Func: nil, Children: map[string]*bintree{
		"app.js": {Func: bindataFrontendAppjs, Children: map[string]*bintree{}},
		"index.html": {Func: bindataFrontendIndexhtml, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
