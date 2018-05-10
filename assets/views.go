// Code generated by go-bindata.
// sources:
// assets/add.gohtml
// assets/footer.gohtml
// assets/header.gohtml
// assets/index.gohtml
// assets/view.gohtml
// assets/views.go
// DO NOT EDIT!

package assets

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

var _assetsAddGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x92\x41\x8b\xd6\x40\x0c\x86\xef\xfe\x8a\x90\x7b\xed\x1f\x98\x0e\xb8\x28\x22\x78\x52\x3c\x4b\xfa\x25\xb5\x83\xed\x64\x48\xd3\x75\x97\xa1\xff\x5d\xda\xed\x7e\xbb\x2a\x0a\xb2\xb7\xd0\xa4\xcf\xfb\xc0\xbc\xb5\x02\xcb\x90\xb2\x00\x12\x33\xc2\xb6\xbd\x02\xa8\x15\x5c\xe6\x32\x91\x0b\xe0\x28\xc4\x62\xe7\x26\x94\x18\x08\x46\x93\xa1\xc3\x5a\xe1\xf5\x0d\x2d\xf2\xe5\xd3\x47\xd8\x36\x8c\x21\xc1\x65\xa2\x65\xe9\x70\x20\x18\xa8\x21\x33\xfd\xd1\x4c\x32\x38\x02\x59\xa2\x66\x4c\xcc\x92\x3b\x74\x5b\x05\x63\x68\x53\x84\xf7\x0a\x3d\x5d\xbe\x83\x2b\xa4\xcc\x72\x17\x5a\x8a\xa1\x2d\x71\x8f\x1a\xd4\x66\x98\xc5\x47\xe5\x0e\x8b\x2e\x8e\xfb\x67\x80\xc0\xe9\xf6\x9a\xa4\x36\x37\xdf\x4c\xd7\x72\x2e\x01\xc2\x44\xbd\x4c\x30\xa8\x75\x98\x18\xe3\x8d\x92\x31\x7c\x78\x1b\xda\x63\x71\x3d\x4b\xb9\xac\x0e\x7e\x5f\xa4\x43\x97\x3b\xc7\x5f\x98\x17\xcd\x6e\x3a\x21\x24\x3e\x30\x90\x69\x96\x87\xa9\x4c\x74\x91\x51\x27\x16\xeb\xf0\x91\xfe\xe8\xd6\x72\xba\xfd\x6f\xcd\xc5\xc9\xfc\x2b\x93\x0b\xc6\xcf\xfb\x0c\xfb\xfc\x2f\xe1\xe3\xf6\xef\xc2\xcf\x80\xa7\xf8\xf3\x88\x17\x98\x4a\xe6\x13\xf2\x2e\xf3\x4b\x2d\xaf\xb0\xd3\xf1\x09\xfe\x87\x61\xbf\xba\x6b\x3e\xb1\xcb\xda\xcf\xe9\xe9\xbd\x7a\xcf\xd0\x7b\x6e\x8a\xa5\x99\xec\x1e\xe3\x1b\xe6\xd0\x3e\xfc\x71\x14\xa9\xdd\xa3\xe3\xef\xbd\x1e\x54\xfd\xec\x75\xad\x20\x99\x61\xdb\x7e\x06\x00\x00\xff\xff\xc6\xfe\x50\xc4\x0e\x03\x00\x00")

func assetsAddGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsAddGohtml,
		"assets/add.gohtml",
	)
}

func assetsAddGohtml() (*asset, error) {
	bytes, err := assetsAddGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/add.gohtml", size: 782, mode: os.FileMode(420), modTime: time.Unix(1496071034, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsFooterGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\xcf\xc1\x6a\x84\x30\x10\x06\xe0\xbb\x4f\x31\xe4\xae\x69\x0b\x6d\xb7\x25\xe6\xd0\x87\xe8\xb5\x8c\x66\xb2\xa6\x8d\x89\x24\x23\x61\x11\xdf\xbd\xa8\xcb\x2e\xec\x5c\xe6\xf0\xf3\x7f\xf0\x2f\x0b\x18\xb2\x2e\x10\x08\x1b\x23\x53\x12\xb0\xae\x15\x00\x80\xca\x7d\x72\x13\x43\x4e\x7d\x2b\x06\xe6\x29\x7f\x4a\x39\x67\x6a\x6c\x0c\x8c\x85\x72\x1c\xa9\xe9\xe3\x28\x4f\xcf\x2f\x6f\xaf\xef\x27\xf3\xf1\xd4\xfc\x66\xa1\x95\x3c\x8a\xfa\x50\x26\xfd\x75\x01\x85\x30\x24\xb2\x77\xa7\x94\xd2\xe4\x12\x93\xe9\x88\x71\x57\x04\x30\xa6\x33\x71\x2b\x7e\x3a\x8f\xe1\x4f\xe8\x5b\xae\x24\x6a\xa8\xe1\xf0\x1e\xa5\xb3\xe3\x61\xee\x76\xe2\x56\x90\x9c\xc8\xfb\x58\x77\x73\x0a\x26\x96\x70\xb5\xef\xf4\x4e\x6d\xf7\xed\xa8\x40\x0c\xa0\x1c\xf4\x1e\x73\x6e\x85\x45\xb0\x58\x1f\x6a\x8d\x9e\xb7\x41\xee\xba\x45\xa2\x56\x72\xd2\x95\x92\x5d\x34\x97\xed\x0f\x3c\x7a\x5d\x2d\x0b\x50\x30\xb0\xae\xff\x01\x00\x00\xff\xff\x73\x31\x31\x8b\x4e\x01\x00\x00")

func assetsFooterGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsFooterGohtml,
		"assets/footer.gohtml",
	)
}

func assetsFooterGohtml() (*asset, error) {
	bytes, err := assetsFooterGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/footer.gohtml", size: 334, mode: os.FileMode(420), modTime: time.Unix(1496071030, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsHeaderGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x44\x90\x4d\xcf\xd2\x40\x14\x85\xf7\xfd\x15\xd7\xd9\x9a\xb7\x13\x2d\x11\x62\xda\x26\xa0\x28\xca\x02\x02\x48\xc2\x72\x3a\x73\xe9\x8c\xcc\x07\x99\x3b\x15\x6b\xd3\xff\x6e\x80\xe8\xbb\xba\xb9\xcf\x39\x9b\xf3\x0c\x03\x28\x3c\x1b\x8f\xc0\x34\x0a\x85\x91\xc1\x38\x66\xe5\x9b\xcf\x9b\x4f\x87\xd3\x76\x09\x3a\x39\x5b\x67\xe5\xbf\x83\x42\xd5\x19\x40\x99\x4c\xb2\x58\x1f\x22\x5a\x1b\x60\xd1\x45\xaf\xc2\xcd\x97\xfc\x89\xef\x05\x87\x49\x80\xd4\x22\x12\xa6\x8a\x75\xe9\xfc\x32\x63\xc0\x1f\x91\x35\xfe\x02\x11\x6d\xc5\x28\xf5\x16\x49\x23\x26\x06\x3a\xe2\xb9\x62\x3a\xa5\x2b\x7d\xe4\xdc\x89\xdf\x52\xf9\xbc\x09\x21\x51\x8a\xe2\x7a\x7f\x64\x70\xfc\x3f\xe0\x45\x5e\xe4\x53\x2e\x89\x5e\x59\xee\x8c\xcf\x25\x11\xcb\x00\x00\x8c\x4f\xd8\x46\x93\xfa\x8a\x91\x16\xc5\x6c\xf2\xb2\x38\x9e\x8c\xd9\x7f\xfb\x82\xeb\x77\xea\xab\xfb\xbe\x9b\x5f\x7a\xd9\xad\xe6\xab\x5d\x5b\xbc\xdf\xb8\x1f\xf2\x76\x9b\x06\x5f\xec\x4e\xaa\x9d\x1c\xc5\xdb\xad\xdb\x1f\xe8\x0f\x5f\x7f\x98\xfd\x6a\xd4\xf2\xa7\x9e\x74\x0c\x64\x0c\x44\x21\x9a\xd6\xf8\x8a\x09\x1f\x7c\xef\x42\x47\xac\xce\x4a\xfe\x54\x53\x36\x41\xf5\x20\xad\x20\xaa\x98\x0c\x3e\x09\xe3\x31\xb2\xc7\xee\x26\xde\x05\x0c\x03\xa0\x57\x30\x8e\x7f\x03\x00\x00\xff\xff\x59\xe7\xf9\x64\x7a\x01\x00\x00")

func assetsHeaderGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsHeaderGohtml,
		"assets/header.gohtml",
	)
}

func assetsHeaderGohtml() (*asset, error) {
	bytes, err := assetsHeaderGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/header.gohtml", size: 378, mode: os.FileMode(420), modTime: time.Unix(1496071027, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsIndexGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\xdd\x8a\xdb\x3c\x10\xbd\xff\x9e\x62\x10\xdf\xad\xa3\x74\x2f\x7a\xe5\x18\xf6\xaf\x50\x28\xa5\xf4\xe7\x01\x26\x99\x71\x2d\x90\x25\x33\x9a\x6d\x0b\x46\xef\x5e\x64\x3b\xd9\x6c\x36\x1b\xd2\x1b\xa3\xf9\x39\x3a\x9e\x73\x24\x8d\x23\x10\xb7\x2e\x30\x18\x17\x88\xff\x18\xc8\xf9\x3f\x80\x71\x04\xe5\x7e\xf0\xa8\x0c\xa6\x63\x24\x96\xa5\x52\x0f\x4d\x8d\xd0\x09\xb7\x1b\x33\x8e\xb0\xba\xc3\xc4\x3f\xbe\x7e\x82\x9c\x91\xc8\x34\xb5\x83\x9d\xc7\x94\x36\xa6\x45\x68\xb1\x1a\xfc\x53\x32\x4d\x6d\x5d\x03\xb7\x44\xa0\xc2\xde\x47\xd8\x46\x14\xaa\x2d\x36\xb5\x1d\x9a\xb2\xa9\xe2\xd6\xf3\x1e\x39\x07\xd3\xb7\xda\x46\x21\x16\xa6\x25\x4c\x2a\x6e\x60\x32\x05\x54\x60\x32\x2f\xca\xb2\x6b\x6a\xab\xdd\x71\x7c\x57\x58\x20\x60\xcf\xa7\x95\x6f\x8a\xa2\x40\xa8\xaf\x2a\x8f\x81\xce\xe6\xef\x51\x28\xd5\x5b\x01\xdb\xdc\xc7\x7e\xf0\xac\x4c\x60\xe1\x7b\x54\xf4\xa7\xbd\x5f\xa2\x0b\x7a\x6d\xf3\xed\x4e\x5d\x0c\xe9\x39\x5d\xdb\xfd\x54\xe3\x08\x82\xe1\x27\xc3\x6a\x9a\x24\x4d\xfa\x9f\xcc\x5d\x02\x7a\xe1\xc8\xff\x47\x96\xfc\x72\xfc\xdb\x16\x97\x3e\x3e\x40\xce\xaf\xdd\xf1\x2e\x70\xb5\xeb\x50\xd4\x00\x8a\xc3\xaa\x73\x44\x1c\x36\x46\xe5\x89\x67\xdb\x66\x97\x94\x5e\x12\x96\x3d\x3f\x63\xcf\x90\xf3\xf9\xe2\x03\x2a\x4f\x2a\xaf\x3e\x44\xe9\x51\xc1\xac\x6f\xaa\xf5\xbb\xea\x66\xbd\x7e\x6f\x2e\xa2\x1e\x03\xfd\x13\x66\x72\xe6\x59\xe6\x9c\xc1\xc2\x21\xff\x26\x6a\xf6\xe8\x0c\x6c\x2e\x9c\xc5\x1d\x02\x80\xb7\x04\x17\x6e\x85\x53\x77\x49\xf3\xa5\xe5\xa2\xe0\x57\x30\x11\x97\x1f\xbf\x44\xa4\x82\xd7\xd2\x1c\xcf\x3a\x9d\x3f\x80\xfd\x09\xe4\x40\xcb\xc5\xb7\xd3\x15\x6c\x4e\x5f\x87\x36\x46\x5d\x5e\x87\x43\xfb\xdf\x00\x00\x00\xff\xff\xd6\x82\x22\xb2\x56\x04\x00\x00")

func assetsIndexGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsIndexGohtml,
		"assets/index.gohtml",
	)
}

func assetsIndexGohtml() (*asset, error) {
	bytes, err := assetsIndexGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/index.gohtml", size: 1110, mode: os.FileMode(420), modTime: time.Unix(1496071023, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsViewGohtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x4f\x6f\xe3\xb6\x13\xbd\xe7\x53\x0c\x74\xb1\x0c\xdb\x92\xe2\xdf\x2f\x2d\xa2\x58\x2e\xba\xc9\x6e\x51\xa0\x45\x83\x36\x7b\x28\x82\x1c\xc6\xe2\xc8\x62\x96\x26\x05\x92\x4a\x6c\x78\xf5\xdd\x0b\xea\x9f\xa5\x8d\x53\xf4\xb0\x07\x03\x34\xe7\xcd\x9b\xc7\xc7\x21\xa9\xe3\x11\x18\x65\x5c\x12\x78\x2f\x9c\x5e\x3d\xa8\xaa\x0b\x80\xe3\x11\x2c\xed\x0a\x81\x96\xc0\xcb\x09\x19\xe9\x36\xb2\x2a\xd6\x2b\x84\x5c\x53\x96\x78\xc7\x23\x04\x1f\xd0\xd0\xe7\x3f\x7f\x83\xaa\xf2\xd6\x2b\x0e\xa9\x40\x63\x12\x2f\x43\xc8\x70\x81\x5a\xab\xd7\x85\xa0\xcc\x7a\x80\x9a\xe3\x22\xe7\x8c\x91\x4c\x3c\xab\x4b\xf2\xd6\xab\x90\xaf\xe1\x17\x05\x1b\x4c\xbf\x80\x55\xc0\x25\xa3\xfd\x2a\xc4\xf5\x2a\x2c\xd6\xae\x14\xe3\x2f\xf0\xca\x99\xcd\x13\x6f\x19\x45\xc5\xde\x83\x9c\xf8\x36\xb7\xdd\x5f\x07\x02\x58\xa5\x28\x5f\xd0\x00\x67\x89\xb7\x29\xb5\x64\xea\x55\xde\xe6\xa8\xad\xab\xd0\xc4\x6a\xb6\x90\xf1\x97\x7a\x60\x52\xcd\x0b\x0b\x46\xa7\x89\x97\x5b\x5b\x98\x38\x0c\x53\x26\x9f\x4d\x90\x0a\x55\xb2\x4c\xa0\xa6\x20\x55\xbb\x10\x9f\x71\x1f\x0a\xbe\x31\x61\xcd\x17\x3c\x9b\x70\x19\xfc\x3f\x88\xda\xbf\x9b\x52\x32\x41\xc1\x8e\xcb\xe0\xd9\xb8\x62\x0d\xf1\xa0\x46\x23\xf0\x05\x35\xa4\x76\x0f\x09\x30\x95\x96\x3b\x92\x36\xd8\x92\xfd\x28\xc8\x0d\x3f\x1c\x7e\x65\xfe\x37\xba\xa7\x37\x7d\x9e\x55\x16\xc5\xbd\xe2\xd2\x1a\x48\xa0\x36\x5c\xa1\x66\x41\x3b\x55\x55\x27\x68\x51\x4f\xfd\x55\x90\xb4\xf7\xa4\xef\xf0\xe0\x12\xea\xa8\xdb\xcf\x05\x68\x94\x5b\xea\xf2\x6f\x51\xb3\x7b\xad\xb6\x9a\x8c\x81\x45\xbd\xb3\x1d\x12\x82\x3b\xb4\x14\x7c\x52\x7a\x87\x16\x9c\xd5\x3f\x2c\xa2\xcb\x45\xb4\x74\x1d\x10\xd7\x80\xbe\xfa\x7c\xc0\x4f\x92\xf5\x4c\x03\x59\x0c\x0f\x4e\xfa\xe3\x5b\x25\xae\xcc\xb8\x78\xdd\x51\xe7\x0b\x7b\xef\x95\x7a\x3a\x95\x1a\xb9\x08\x09\x48\x7a\x85\x7a\xec\xa7\x76\x3f\xef\xcd\xb0\x87\x82\x62\x98\x08\x2e\x69\xd2\xb1\x32\xb4\x18\xf7\x08\x00\x81\x1b\x12\x26\xae\xd5\xcf\xfb\x59\x87\x32\x64\x4d\xdc\x2f\xa7\x56\x34\x18\x03\x64\x5c\x88\x18\x5c\x87\xcf\x47\xf3\xae\xcb\xb7\x5a\x95\x92\xdd\x2a\xa1\x74\x0c\x9e\xde\x6e\xd0\xff\xf1\x6a\x7e\x79\xbd\xac\x7f\x51\xb0\x9c\x7a\xe3\xa4\x5a\x46\x0c\x93\x8f\xc6\xf2\x1d\x5a\x62\x93\x71\xbc\x91\xed\xba\xa9\x03\xdc\xa1\x45\x7f\x3a\x44\x55\xf3\xef\xa5\x75\x79\x75\x35\x87\xeb\xeb\x39\x5c\xfe\x6f\x39\x87\x7f\x91\xfb\x73\x6a\x4b\x14\xef\x69\x6d\xa2\x67\x84\xf6\xe3\xa7\x8b\x91\xf4\xaa\x3d\x10\x59\x29\x53\xcb\x95\x3c\xb3\xe0\x7e\x61\xd4\xcd\xbb\x9e\x6b\x7b\x03\xa0\xe8\x0e\xc4\xf0\x3c\x85\xf5\xee\x06\x82\xe4\xd6\xe6\x1d\xd2\x5d\x7b\x0f\x0e\x34\x06\x77\xe1\x4c\x69\xf0\x39\x24\x10\xdd\x00\x87\xd5\x88\x01\xf8\x6c\x36\x1d\x38\xdc\x4b\x79\xa0\x5d\xe1\xe8\x7a\xea\x45\x2b\xe8\xe6\x2d\xf6\x91\x3f\x41\xf2\x4d\xea\x0a\x22\xf8\x09\x22\x88\xc7\xf3\xa7\xec\x91\xe8\x77\xab\x74\xfe\x6a\xb2\xa5\x96\x27\xae\x26\x5e\xbd\xb1\x78\xb8\x4f\xfd\xb2\xb0\x9e\x1c\x99\xfb\x1d\x2d\x6b\xee\xb0\xd6\xaf\x37\x17\xda\xa3\x4b\x7d\xe4\x4f\x4f\xf0\xf5\x2b\x44\xa7\xe5\x37\x9a\x1a\xe7\x46\xcb\xef\xd9\xfe\x83\x55\x67\xb0\x3c\x03\x9f\x9b\x07\xc5\xf0\xe0\xb7\xa5\xa7\xd3\xd1\x11\xda\x68\xc2\x2f\xa7\x84\xea\xbc\xd3\x8d\xbe\xb3\x36\x9f\xe8\x2d\x9d\xa8\x9b\x7b\xd3\xd2\x1f\x9b\xe7\xf6\x1a\x73\x77\xa5\x3f\xbd\x19\xc4\x77\x4a\xda\x1c\x12\xf0\x27\xd1\x04\x66\xe0\xb7\x78\xf7\xb4\x7c\x7e\xb8\xfd\xdd\x45\xfd\x29\xcc\xe0\x72\x3a\x0d\x8c\xe0\x29\xf9\x8b\xe5\x88\x80\xd5\x07\xa2\x4d\x1f\x67\x37\xd5\xde\xc9\x3b\x10\x6a\xf7\x96\x8d\x32\x3e\x95\x42\xfc\x4d\xa8\x4f\x1a\xdb\xb5\x3b\x18\x24\x49\x93\x35\x03\x6f\xe1\xc1\xac\xd5\xde\xfd\x63\x5d\x8b\xd6\xdf\x17\x83\x57\x74\xf4\x15\x92\x29\x65\xdb\xaf\x90\xe3\xb1\x7e\x01\xaa\xea\xe2\x9f\x00\x00\x00\xff\xff\x22\x5a\xc6\xd5\xbe\x08\x00\x00")

func assetsViewGohtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsViewGohtml,
		"assets/view.gohtml",
	)
}

func assetsViewGohtml() (*asset, error) {
	bytes, err := assetsViewGohtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/view.gohtml", size: 2238, mode: os.FileMode(420), modTime: time.Unix(1496071751, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsViewsGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x9a\x5b\x6f\x1c\x49\x72\x85\x9f\xbb\x7f\x45\x2d\x81\x5d\x90\x86\x4c\xd5\xfd\x42\x40\x0f\xde\x9d\x59\x7b\x1e\x3c\x0b\xac\xc7\x7e\x71\x1a\x8b\xba\x64\x69\x1a\x16\xd9\x32\x49\xcd\xa4\x34\xd0\x7f\x37\xbe\x88\x53\x62\x8b\xe2\x5c\xa0\xd1\x43\x93\xcd\x62\x55\x66\x64\x64\xc4\x39\x27\x22\xeb\xf9\xf3\xec\x2f\xc7\x25\x66\x2f\xe3\x4d\xbc\x1d\xef\xe3\x92\x4d\x6f\xb3\x97\xc7\x7f\x9e\x0e\x37\xcb\x78\x3f\x5e\xee\x9f\x3f\xcf\xee\x8e\x6f\x6e\xe7\x78\x77\xc5\xf7\xf1\xee\x2e\xde\xdf\x3d\x1f\x97\xe5\xf2\xe5\xf1\xfb\xfb\xeb\x57\x27\x17\xd7\xe3\xf1\x3e\xde\x7e\x7a\xfd\xfb\x38\x2e\x4f\x5d\x3f\xdc\x2c\x31\x7d\x7a\xf9\x87\x43\xfc\xf1\xe9\xab\x77\x97\x2f\x8f\x5c\xfa\xea\x6f\xd9\xb7\x7f\xfb\x2e\xfb\xfa\xab\x6f\xbe\xfb\xc3\x7e\xff\x7a\x9c\xff\x77\x7c\x19\x75\xe3\x7e\x7f\xb8\x7e\x7d\xbc\xbd\xcf\xce\xf7\xbb\xb3\xe9\xed\x7d\xbc\x3b\xdb\xef\xce\xe6\xe3\xf5\xeb\xdb\x78\x77\xf7\xfc\xe5\xbb\xc3\x6b\x2e\xac\xd7\xf7\xfc\x3a\x1c\xfd\xe7\xf3\xc3\xf1\xcd\xfd\xe1\x15\x7f\x1c\xed\x81\xd7\xe3\xfd\xf7\xcf\xd7\xc3\xab\xc8\x17\x2e\xdc\xdd\xdf\x1e\x6e\x5e\xda\xff\xee\x0f\xd7\xf1\x6c\x7f\xb1\xdf\xaf\x6f\x6e\xe6\x4c\x9e\xfa\x7b\x1c\x97\x73\xbe\x64\xff\xfd\x3f\x4c\xfb\x2c\xbb\x19\xaf\x63\xe6\x8f\x5d\x64\xe7\xdb\xd5\x78\x7b\x7b\xbc\xbd\xc8\x7e\xda\xef\x5e\xbe\xb3\xbf\xb2\xab\x17\x19\x56\x5d\x7e\x1b\x7f\xfc\xbb\x79\xea\xdc\xcc\xe6\xef\x3f\xbf\x59\xd7\x78\x6b\xc3\x5e\x5c\xec\x77\x87\xd5\x1e\xf8\xc3\x8b\xec\xe6\xf0\x8a\x21\x76\xb7\xf1\xfe\xcd\xed\x0d\x7f\x3e\xcb\xd6\xeb\xfb\xcb\xaf\x19\x7d\x3d\x3f\x63\xa0\xec\x8f\xff\x77\x95\xfd\xf1\x87\x33\xb7\xc4\xe6\xba\xd8\xef\xde\xef\xf7\xbb\x1f\xc6\xdb\x6c\x7a\xb3\x66\x3e\x8f\x4f\xb2\xdf\xfd\xc3\xcd\x79\x91\x1d\x8e\x97\x7f\x39\xbe\x7e\x7b\xfe\xa7\xe9\xcd\xfa\x2c\x7b\xf9\xee\x62\xbf\x9b\x5f\x7d\xbd\x59\x7a\xf9\x97\x57\xc7\xbb\x78\x7e\xb1\xff\x52\xf6\x30\x8c\x8f\xff\x33\x03\xc5\xdb\x5b\xb7\x5b\x17\xa7\x37\xeb\xe5\x9f\x31\xfd\xfc\xe2\x19\x77\xec\xdf\xef\xf7\xf7\x6f\x5f\x2b\x02\x70\xf9\x9b\xf9\x9e\x51\x6c\x7d\xda\x8f\xfd\xee\x70\xb3\x1e\xb3\xec\x78\x77\xf9\xd7\xc3\xab\xf8\xcd\xcd\x7a\xfc\xf0\x9c\xb6\x70\xbb\x7e\x32\x82\xed\x61\x96\x69\x1b\xf7\xbb\xbb\xc3\x3b\xfb\xfb\x70\x73\xdf\xd6\xfb\xdd\x35\xa9\x93\x7d\x18\xf4\xdf\x8f\x4b\xb4\x8b\xdf\x1d\xae\x63\x46\x98\x5c\xf2\x8d\x79\x2c\x54\xce\xd7\xc3\xe3\xb9\x2e\xb2\x6f\xc7\xeb\x78\x7e\xa1\x19\x98\x53\xab\x5c\x0f\x97\xcc\xbe\x7f\xff\x0b\xcf\xfe\xc7\xe1\x1d\xcf\x9a\x35\x1f\x3f\x8a\xa1\xbf\xf8\x28\xb6\x9e\x5f\x9c\x5a\xfe\xf1\x00\x2c\xed\xd7\x06\x60\x71\xe7\x17\x0f\x0b\xfd\x64\x04\xad\xfe\xe7\x07\xf9\xe6\xee\xab\xc3\xed\xf9\x45\x36\x1d\x8f\xaf\x4e\x9f\x1e\x5f\xdd\xfd\xca\xca\xdf\xde\xf9\xc2\xe3\xed\x3a\xce\xf1\xa7\xf7\x27\x4f\x2b\x24\x88\xf2\x7f\x38\x28\xfc\xcb\xb2\xfc\xab\x01\x4a\xf6\x42\xe1\x70\x7e\x16\x52\xb1\x86\xd4\x4f\x21\xe5\x7d\x48\x79\xfe\xf4\x67\x5d\x43\x1a\xeb\x90\x86\x32\xa4\xba\xf0\xfb\x97\x36\xa4\x9a\xff\xcf\x21\xf5\x6d\x48\x71\x0d\x69\x8d\x21\xf5\x63\x48\x43\x1e\x52\x37\x85\x14\x17\x1f\x7f\x60\xec\x18\xd2\xd4\x87\x54\xf2\x29\x43\xea\xfa\x90\x9a\x32\xa4\x6a\x0e\xa9\x9e\x42\x5a\xc7\x90\xca\x26\xa4\xa9\x09\xa9\xaf\xfc\xd9\xb6\x0e\xa9\xee\x43\x5a\xaa\x90\xba\x26\xa4\xa1\x0b\x69\x2c\xdc\x9e\x66\x09\x69\x19\xfd\xbe\x8e\xb1\xa7\x90\xca\x31\xa4\x7c\x0c\x69\x2a\x43\x9a\xba\x90\x96\xdc\xed\x9e\xb1\x6d\x0a\x69\xce\x43\x9a\x66\x9f\x23\x2f\x43\x9a\x27\xb7\x95\xfb\x59\x67\x81\x3d\x55\x48\x33\xcf\xb7\x21\x4d\x8b\xdf\x37\xf6\x21\x15\x4d\x48\xcd\x1c\x52\x6c\x43\xaa\xca\x90\x86\x22\xa4\x9c\x35\xe6\xbe\xa6\xb9\x0e\xa9\x2d\x43\x8a\x5d\x48\x65\x1b\xd2\x50\x87\x54\xf4\xee\xd7\x9a\xbf\x2b\xb7\x7d\xae\x42\x6a\xb0\x9b\x75\x60\xc7\x12\x52\xb9\x84\xb4\xf2\x6c\xe3\x6b\xad\xb1\xbd\x0f\xa9\x6a\x43\xea\xe7\x90\x4a\x9e\x2b\x42\x6a\x9b\x90\xc6\xd2\x7f\xc7\x21\xa4\x8e\xb9\x73\x9f\x07\x1b\xb9\x0f\xfb\x57\x7c\x53\x84\x54\xcf\x6e\x6b\xd5\xfb\x3a\x9a\x41\xcf\xb7\xfe\xbf\x79\xf6\x3d\xad\xa6\x90\xba\x3a\xa4\x86\x38\x68\x42\x6a\xab\x90\x5a\xf6\xa7\x0a\xa9\xaf\x43\x5a\x3b\xf7\x6b\xb5\xb8\xdf\xa7\xe8\x7b\x54\x4e\xf2\xef\x1c\x52\x57\x86\x54\x74\xbe\x36\x62\x80\xb5\xb2\xae\x8e\x78\x59\x43\x2a\xc6\x90\x96\xda\xe7\x26\x1e\xe6\xc6\xd7\xc9\x9a\x89\x0d\x62\xaa\xe4\x77\xf4\xbd\x6a\xbb\x90\xfa\xdc\xf7\x8c\xb5\xae\x3c\x37\xfa\x7c\xf8\xa5\xc2\x47\xab\xaf\x63\xe9\x7c\x7e\x9e\xcf\x0b\xdf\xbf\xba\xf6\xfd\x33\x1f\xe4\xee\x1f\x8b\xa1\xde\xfd\x15\x59\xdb\xa2\xf5\x17\x21\x75\xb3\xc7\x64\x31\x79\x5c\xe1\x03\x6c\x67\xcd\xc4\xe7\x84\xef\x66\xb7\x95\x78\x6b\x94\x13\x75\xe5\x71\x89\x0f\xe7\xce\xaf\xdb\x1c\x5d\x48\x33\x31\x1c\x43\xaa\x46\xdf\x9b\x12\xdb\xa3\xdb\x43\xdc\xb5\x43\x48\x03\x7b\x4c\x5c\x0f\x6e\x2b\xfb\x40\x6c\x35\xdc\xdf\x85\x54\x90\x5b\xe4\x47\xee\xd7\xc9\xb1\xd8\x7b\xfe\xb1\xe6\x69\xf4\xfd\x6e\x57\x9f\x0f\xbf\xce\xf8\x6b\xf6\xfd\x21\xee\x88\xd3\xb9\x7d\xc8\x03\x72\x96\xdf\x76\xcf\xea\xb1\x88\x3f\x56\xe5\x31\xfe\xe3\x5e\xfc\x3f\x76\x21\xad\x7d\x48\x2b\xfe\xd2\xba\x58\x5f\x3d\x7a\x2e\x14\x95\xc7\x2e\xbe\xe7\x1e\x7c\xc5\xbe\x8f\x6b\x48\x53\xee\xf1\xbc\x92\x2b\x83\xf0\xa1\x0b\xa9\x2d\x42\x9a\x56\xb7\xbd\x9d\xdc\x27\x53\xe1\xf9\x88\xef\x99\x9b\x3d\xe7\xc3\x1e\x76\xba\xb6\x08\x57\x5a\xe1\xcc\xc8\x7e\x72\xdf\x1c\x52\x11\x7d\x0d\xec\x1f\x76\x71\x2f\xe3\xb2\x46\xf6\xb1\xa8\xdd\xc7\x0b\xb6\x55\x7e\x2f\x6b\x65\x7c\x9e\x6d\x6a\xf7\x23\x63\x11\x27\xe3\xe2\xb9\xc5\xf8\xd8\xbb\x4c\xbe\xef\x79\xfb\x31\x2e\xf2\xc1\xb7\xac\xad\xc9\x1d\x07\x88\x91\xbc\xda\xee\x3b\xdb\x44\xd2\x23\x18\x16\x7b\x3f\xa5\x8a\x36\x8e\x3f\x51\x55\xfb\xdd\xee\x31\x8e\x3f\xdb\xef\x76\x67\x9f\x68\xd1\xb3\x67\xfb\xdd\xc5\x07\xb6\x7d\xf4\x08\xd3\xfd\x93\x5d\x3b\x9d\xce\x24\xc2\x07\x1d\xf6\xb4\x99\xbf\x26\x74\x3e\xe8\x13\x53\x18\x57\x2f\x1e\xb3\xd5\x4f\xf0\xf8\x55\xf6\x94\xbd\x19\x3c\x7d\x95\x75\x7d\xf9\x2c\x83\x71\xaf\x4e\x09\xf9\xbc\x2e\xf3\x0b\xbb\x0e\x8f\x5e\x39\xcf\xfe\xe7\xcd\x21\x9d\x17\xf5\xd0\xe6\x5d\x91\x57\xf5\xb3\x2c\xbf\x78\xbf\xdf\x8d\x4c\xfb\x27\x9b\xe0\x27\x5b\xd2\x55\xa6\x95\x61\xd3\x95\xfd\x7c\xff\xc1\xb9\xe3\xb3\x27\x38\xf2\xaf\xa6\xde\x7f\x07\x4d\x42\x11\x84\xaa\xc1\xf5\xe8\xf0\x49\xba\x17\xb9\x87\x0e\x74\x01\x55\xd5\xab\xc3\x4e\x24\x2c\xa3\x43\x01\x69\xda\x2e\x4e\x5d\x50\xe1\x16\xc6\xa4\x0c\x29\x6f\xd4\x38\x3b\x84\x42\x59\x63\xeb\x10\xd6\x0f\x0e\x2d\x65\xe5\xa1\x5a\x14\x0e\x8d\x84\x36\xb0\x37\x0b\x5a\x09\xed\x8d\xc2\x80\x13\xd2\xb5\x5b\xfd\x3b\x30\xc0\xfc\x40\xe3\xa4\x74\x36\x9b\x05\x89\x8c\x0d\x1d\x40\x91\xa4\x35\x36\x43\x89\xac\xdb\x60\x7a\x0c\xa9\x5b\x44\x03\x95\xc3\x62\x1d\xfd\x1a\x30\x64\x6b\x6f\x43\x2a\x07\x9f\x13\xf8\xa8\x06\x87\x79\xfc\xd4\xce\x92\x12\x7c\x1a\x8d\x33\x7b\xca\x1a\x14\x54\x4e\xb3\xf8\x0d\xff\x62\x2f\x90\x07\xcc\x90\xc6\x40\xe5\x2a\xa8\x81\x66\x56\xf9\x09\x1a\x1a\x1a\x97\x19\x40\x06\x32\x83\x34\x85\xa2\x49\x77\x52\x1d\xca\x60\x4e\xf6\x09\x3f\xb2\xfe\xae\x73\xe8\x83\xc6\x97\xd2\xf7\x89\xb5\x03\xa7\xd8\x03\x14\x02\x2b\x0d\xd4\xa8\x3d\x66\x3f\x2a\xd1\x46\x29\xb8\x00\xfa\x7b\x41\x2c\xb6\xb3\x8f\xd0\x2c\xfe\x63\x2e\xd6\xce\x1e\x99\x1c\xc8\x1d\x8a\x80\xb5\x49\x10\xc5\x7e\x46\xfc\x88\xff\xa0\xf0\xd6\xa9\x63\x20\xce\x7a\x87\xf1\xa6\x77\x7b\xbb\xca\x69\xda\xe4\x47\xeb\xf2\x80\x98\x31\x69\x56\xfb\x7e\x10\x5f\xd0\x79\x54\x6c\x6c\x12\x6e\xcc\xdd\xdf\xdc\x87\x0d\x8c\x65\x30\x0b\x4d\x37\xbe\xe7\xcc\x43\x1e\x58\x5c\x43\x9d\xd1\x63\xb5\x96\x8d\xc0\x79\x2d\x89\xd2\x88\x9e\xf0\x1d\xfe\x87\x66\x91\x0e\x55\xed\x74\xc9\xfc\xf9\xea\xfb\x02\xbc\x37\x92\x41\xc4\x20\xfb\x83\x3f\xb7\x38\x23\xaf\xd8\xa3\xc7\xb0\x8b\x8d\xe4\x50\x25\x49\x5a\xc7\xd3\xfb\x1e\xc1\xee\x69\x66\x7f\x2e\xf2\x9e\x8e\x71\x0a\xbe\x1f\xd5\xfc\x4f\xe1\xef\xe9\x83\xbf\x1d\x82\x9f\x30\xf9\x0b\xa3\xf0\x23\xc3\x05\xc4\x15\x88\xfa\x39\x40\x9c\x7f\x39\x20\xfe\x37\x6b\x02\xfc\x0e\x20\x46\xff\x91\x28\xf5\x22\xed\x50\x7a\xb0\xa3\x03\x08\x6a\x92\x00\x00\x00\xc4\xd0\x8f\xcb\xe0\xfa\x92\x80\x06\xc0\x08\x46\x80\x14\x4d\xbf\x28\xa9\xc6\x4d\xeb\x8f\xae\xa7\xf9\x50\x9f\xa0\x97\x08\x76\x92\xbd\xdb\x00\x42\xfa\x3a\xef\x5c\x43\xa0\x0f\x99\x0b\xbd\x63\x9a\x68\x55\x82\xe5\x02\xf7\xc9\x13\x08\x9d\x89\xbd\x80\xe3\x30\x39\xa8\x91\xa0\xb9\x00\x90\x84\x01\x90\x01\x15\xb4\x31\x09\x45\xd2\xb3\xa6\x41\x75\x02\x7a\x1f\xf0\x43\x5b\x33\x06\xe3\x0d\x22\x1d\xc8\xc4\xea\xa9\xd6\x35\x59\x25\x20\x6e\xa4\xb9\x79\x06\x6d\x86\xbe\xaf\x59\x7b\x13\x52\x31\xb8\xef\x06\x69\x54\x00\x72\x03\x03\x6a\x39\x34\x3f\x36\xb5\xb9\xea\x8e\xc6\x81\xd9\x74\xe4\xe2\x09\x0f\x20\x43\x54\x80\x11\x35\x06\x76\xd4\x83\xaf\x1f\xc0\x66\x9c\x62\x23\xb5\xd1\x81\x39\x4a\xc7\x52\xc3\xa0\xc7\x59\xaf\xd5\x93\xac\xb1\x71\xcd\xc5\x1e\xd8\x3e\x49\xfb\xe2\x23\xea\x2a\x74\x33\xe3\x43\x5c\xec\x1d\x04\xc4\x5a\x01\x51\x7c\xdc\x6a\xbf\x00\x5b\x03\xeb\xc5\x41\x7e\x91\x9d\x10\x28\xf5\xe9\x3a\x78\xcd\x88\xaf\x4a\xe9\xf2\x5e\xe3\x00\x5c\x90\x19\x75\x2a\x60\x6b\xb6\xae\x4e\xf2\xf8\xa0\x89\x3e\x26\xc4\x69\xb5\xcd\xe0\xd7\xa8\xc3\x00\xcc\x5e\x62\x01\xb2\x2f\xa4\x7f\xb7\x18\xe6\x7f\xec\xd7\x72\x52\x3f\xae\xaa\xad\x58\x3f\xeb\x63\xbd\x10\x27\xfa\x7b\x52\x9d\x37\x68\xdc\x45\x44\x0b\x41\xf4\xa5\xd7\x10\x90\x44\x64\xcd\x93\x5f\xa7\x8e\x23\x2e\xa8\x59\x58\x0b\xcf\xa0\x7b\x21\x6a\xee\x01\x94\x89\x0b\xd6\x6f\x82\x42\x7b\xce\x7c\xc4\xcf\x20\x62\xe7\x79\xc4\x05\x00\x3c\x62\xdf\xe2\x64\x82\x7d\xe8\x66\xc6\xe4\x77\x21\x02\x04\xf8\x6d\x4e\xd5\x4a\x56\xbf\xa8\x3e\x64\x6c\x23\xde\xe8\xb5\x10\xb1\xdf\xea\x37\x39\xcd\x1e\x90\xa7\xd4\x3f\xd4\x25\xf8\x83\xbd\x61\x6c\x6a\x38\xe2\x8b\xb8\xa5\xb6\xa4\x5e\x9b\xa3\x8b\x8d\x55\x1a\x9f\xbd\xa0\x7e\xac\x55\x7f\xa3\xf1\x37\xad\x6f\xb6\xf7\x3e\x17\xb5\x43\x2e\xe1\x01\xc1\xe0\x5f\xea\x3b\xf6\x8d\x78\x5a\xa2\xc7\xf1\x96\x9b\x60\x03\x62\x00\xe2\xa2\x86\x65\x6d\x0f\x35\xc0\x03\x69\xd9\xde\x77\x1e\x57\x16\x37\xe3\x2f\x90\xd6\x29\x0a\x7e\x2e\x69\x9d\x8e\x71\x4a\x5a\x1f\x35\xa4\x9f\x22\xad\xd3\x07\x7f\x3b\x69\x3d\x61\xf2\x17\x26\xad\x47\x86\x6f\xa4\xd5\xf5\x9f\x45\x5a\x65\xf7\xe5\x48\xeb\x9b\x9b\x25\xa6\xdf\xc1\x59\x88\x4d\x62\x94\x1a\x95\xb8\x26\x67\x0c\xef\x73\x17\xf3\x76\x4f\x74\xec\xe2\x1a\x39\x39\xaa\x9e\x25\x57\x4a\xe5\x31\x18\x8e\x90\xa7\x9e\x07\x87\x1b\xf1\x16\x58\x87\xc0\x23\xfe\x88\x39\x13\x8a\x83\x04\xeb\xe2\x7c\x09\x16\x21\x62\xc9\xb1\x56\x82\xac\x56\x5f\x00\xec\x22\x66\xe1\x32\xb0\x02\xec\x21\x97\xac\x67\x52\xfa\x6f\x6c\x21\xb6\x2b\x71\xcb\x20\x31\x89\xd8\x45\x34\x82\xc7\xb6\x9e\xce\x31\xb1\x52\xff\xaa\xe8\x5c\x54\xb3\x46\xfe\x36\x71\x3b\x38\x9e\xc2\x11\xd8\x88\xd0\x66\x6d\x08\xeb\x55\x3d\x1d\x13\xaf\xad\xf3\x03\x73\x20\x7c\x59\x27\xd8\x0d\xce\xc0\xff\xcc\x4b\x21\x45\x3e\x1b\xbf\x4f\xbe\x46\xf2\x14\x91\x09\xef\xce\x95\xd7\xee\xcc\x37\xa9\xc7\x03\x6e\x22\xb4\x8d\x57\x7b\xe7\xda\xad\x3f\x09\xc6\xcd\x2a\x10\xf0\x01\x36\x80\xf9\x6d\xef\xbd\x8c\x62\x74\x3e\x30\xbc\x29\xdd\x8e\x56\xfd\x33\x70\xc1\x84\x73\xed\xb8\x01\x4f\x82\x27\x5b\xdf\x0f\x7f\x5b\xbf\x62\xf4\x7d\x61\x7c\xe6\x2d\x16\x69\x94\x52\x85\x51\x29\x1c\xac\x5c\xb4\xe3\xf3\x5c\xfc\x6e\xfd\xcf\xd1\xc7\x2a\xc5\x17\xd8\xc8\xfe\xc2\xe1\xe5\xe8\xfa\x03\xce\xc6\x3e\xb0\x8c\xd8\xdb\xfe\x26\x96\x66\xcd\x85\x66\x69\xc5\x1b\x8b\x7a\x7a\x60\x2d\x78\x87\x0e\x00\x3b\x79\x0e\x2e\x33\xec\x6e\xbc\x38\xb3\x1e\x8d\x74\x17\x7b\x45\x2c\x96\x2a\x88\xfa\xc2\xb1\x39\x6e\xb1\x55\x78\x8c\xa2\x37\xac\x0f\xa9\x9e\x0b\x7c\x8c\x9f\x3f\xec\xf9\xec\xf6\x33\x1f\x45\x53\x37\xa9\x77\x53\xfb\xdc\xe4\x89\xf5\xe6\x4a\x8f\xdd\x4e\x71\xbc\xaa\xa7\x0c\x47\xd9\x1e\xcc\xae\x2f\x4c\xfb\x88\x9f\xe0\x21\xc6\x02\xf3\xc9\x55\xe3\x4d\xf5\xe1\x88\x0f\x7c\x80\xff\x89\x33\xf6\x13\x3e\x1e\x56\xd7\x79\x8d\x7a\xc5\x9d\x7a\xbb\xb3\x62\x99\x3d\x25\x26\x57\xf5\x48\xf9\x6d\x5c\xd8\xaa\x6f\x18\xfd\x83\x8f\xf0\x07\x3e\xc2\xc7\x14\xb2\xe4\xc7\x87\x22\x4e\x85\x11\x1c\x69\xfa\x40\x31\x3b\x4a\x73\x12\x4b\x85\xf2\xb1\x96\xf6\xeb\x55\x60\x33\x5f\xab\xf9\x0c\x63\xa2\x8a\xd9\xd6\xfb\x9b\xc4\xbc\xe9\xab\xd5\x31\x60\x55\x0e\x13\x5f\xa6\x3b\x47\xd7\x22\xfc\xc6\x3e\xb3\xb7\x76\xff\x35\xd2\x9c\xec\x8b\x15\xe4\xa3\xf3\x2a\x7a\x16\xad\xd1\xaa\x8f\xce\xfe\xf1\x7f\xd6\xc4\xde\x14\xea\x91\x99\x06\xaf\xfc\x3e\xfc\x43\x3c\x18\x47\xb6\x2a\x72\xd5\xcf\x2d\xd5\xf7\x07\xb3\xd8\x07\xeb\x55\x2a\xff\xc1\x21\xc6\x1d\xa4\x39\x0a\xed\x45\x2f\x9c\x01\x37\xc0\x97\xad\xa0\x25\x27\x3b\xe9\x4f\xd6\x0a\x5e\x44\x15\xae\xf1\x84\xc7\xd1\x64\xf8\x0c\xfc\x9e\x54\x53\x8c\xb5\xc7\x85\xd5\x0c\x6a\x52\x98\x5e\x97\xee\xd9\xf0\x8a\x75\x99\x7e\xa8\xb5\xb7\x93\xdb\x36\x29\x3f\x0b\x69\xb7\x5a\xbd\x64\xf4\xf0\xb6\x77\xa6\xbb\xd4\x2b\xad\x2b\xd7\x6a\x60\xfd\x63\xbe\xe0\x43\xfe\x63\x4f\x29\x0d\x48\xe1\x9d\xd7\x3f\xa3\x27\x4e\x08\xea\x73\xe5\xc4\xc9\x10\xa7\x6a\xe2\xf4\x18\xfb\x29\x31\x71\xf2\xd8\x6f\xd7\x12\x9f\x9a\xfb\x85\xa5\xc4\xc7\x56\x4b\x49\x14\x45\x91\x7f\x9e\x94\xa8\xbe\x9c\x94\xf8\xaf\x43\xfc\xf1\xf7\x9c\xd6\xcd\x1e\x09\xb5\xda\x64\x20\xda\xa4\x6e\x3d\xd9\x88\x02\x68\x54\x59\xa2\x1c\x60\xac\x5c\xc8\x34\x94\x62\x16\xb5\xd9\xac\x9b\xaf\x8a\xaf\x54\x9b\x67\x1e\x3c\x9b\xc8\xfe\x51\xd5\x4e\x2f\x65\x0d\x22\x97\x42\x4f\xcb\x0e\xb5\xae\x40\xc4\x56\xa8\x63\x15\x99\xda\x42\xb5\xda\x7f\x9d\x2a\xb7\x45\xaa\x23\x0a\x5d\xc9\x4e\x18\x1c\x7b\x41\x79\x18\x1b\x96\x03\x3d\x2a\xa9\x19\xab\x32\x75\x02\xc3\x07\xb4\x18\xa5\xde\xc9\x62\x14\x45\xab\x13\xbc\x42\xa7\x8f\xac\x0d\xc4\x20\xd3\x2a\xb5\xc3\x40\x2f\xe6\xb6\xaa\x5c\xcf\x96\xb3\xda\x55\xa3\xb3\x16\xf6\xcf\xca\x68\xab\x3c\xd4\xa2\xac\xd4\x96\x2d\x55\x9d\x50\x55\xf6\xb5\xcf\x89\xed\xd6\x92\xeb\x65\x5f\x25\x54\xd2\xa9\x06\x1f\xd6\x51\xa9\x2a\xea\x85\xec\x6b\xf9\x30\x56\x1e\x75\x4a\x94\x3b\x5a\x17\x62\x69\x10\x02\xdf\xc2\x14\xd8\xd7\x6c\x4c\xdd\x39\xda\x99\xc2\x50\x9b\x8d\x75\xc1\x60\xd6\x4e\xd5\x69\x65\x2f\x85\x63\x27\x65\xb3\x3f\x6b\x27\xb2\x9a\x7b\xd4\x29\xe9\xa8\x36\x26\x48\xca\xde\xa1\xa2\xac\x1b\x20\xc4\x6b\x1a\xf7\x0b\x0a\x63\x14\xb2\xe3\x0b\x54\x95\x55\x5a\xa3\xfb\xd2\xc6\x8c\xae\x70\xcc\x0e\xb5\xa0\x51\x1b\xa8\xd1\x79\x63\x84\x55\x9d\x87\x46\x95\xa4\x4e\x1d\xd9\xb3\x55\xa7\xa3\x86\xa6\x62\x7c\x8b\x09\xc5\xe9\xa8\x4e\x40\x23\xff\x92\x17\x76\xa2\x28\x05\x58\x0e\xce\xaa\xa8\x59\xe2\xac\xd1\xf8\x56\xf9\xf6\x52\xd5\x93\x3f\x3b\xab\x4d\x3b\xa9\xf5\xba\x8a\x15\x4c\x59\xe4\xae\x34\x67\xc5\xec\xa6\x96\x47\xcd\xc5\x7d\xa8\x19\x3e\x54\x96\x95\x18\xab\x51\x57\x62\x90\x42\x65\xdc\x5c\x76\x10\x43\xb5\x5a\xb2\xa6\xd8\xe4\x5b\x18\x15\xe5\xcc\x33\x76\xa2\xb5\xea\xb4\x3b\x7a\x85\xda\x16\xf2\x5d\xe7\xfb\x48\x3e\x10\x13\xc5\xf0\xd0\x99\x40\x61\xa3\x0a\x27\x75\x84\xec\x04\x5e\xa7\xfa\x30\x2d\x8a\x78\xd6\xa9\x36\x0a\x7f\x52\xce\xe6\x52\x24\x96\xeb\xad\xc7\x2f\x7b\x46\x1c\x54\x8d\xb3\xa2\x75\x94\x56\x75\x95\x72\xb5\x7d\x5b\xaf\x16\xec\x94\xb5\x75\x75\xcb\x77\x8b\x0d\xa9\xfc\x5c\xac\x09\x56\x98\x4a\x50\x67\xa7\xea\x1c\x6b\xf0\x25\xdf\x3b\xb5\x9c\x9b\xe6\xa1\xfb\x61\x58\xa6\x36\x33\x8a\xbb\xee\xdd\x46\xe6\x22\xef\xc6\x42\xa7\xdd\x85\xdb\xc1\xb3\x7c\x50\x74\x76\xc2\xad\xb6\xbc\x29\xbf\x46\x6f\x03\xe8\xc4\x9b\x35\x45\xc5\x66\xd4\xdb\x01\x51\xed\xfa\x45\x5d\xb5\x5c\x8a\x7f\xd0\xe9\x79\x54\xce\xb6\x52\xb4\x53\xe5\x78\x48\x3c\x4d\x3a\xc1\xb6\x23\x0e\x75\xc9\xb0\x9f\x7b\x50\x2c\x8c\xdf\xeb\x34\x1e\x1b\xc8\x0f\x53\xd9\x52\x95\xbd\x4e\x83\x89\x7f\x54\x32\x3e\x23\x9f\x46\xbd\xc9\x60\xe3\xd5\x8e\xe7\x36\x5e\x2e\x0c\xc8\x3d\x46\x50\xab\x8c\x47\x8e\xb2\x66\xdb\x07\x55\x03\x60\x29\xaa\xcb\xba\x19\xab\xab\x92\x55\x38\xc1\xba\x2c\xa7\xa3\xe3\x08\x31\x01\x8e\x59\xde\x15\xaa\x26\x27\xbd\xf1\xb1\xfa\xbe\x63\xaf\x29\xa2\xc6\xf7\x65\x1a\xc4\x59\x83\xef\x51\xa7\x8e\xa5\xa9\xe5\xd9\xf7\x98\xf8\xa9\x74\x9c\x43\x1c\x4f\x3a\xfa\xc2\x27\xc6\x1b\x7a\x9b\xa2\xd4\x71\x19\x7b\x0e\x6f\x11\xbf\x93\xde\xb6\xb0\x37\x15\xc4\x8d\xf8\x18\x8c\x32\x9c\x14\xe6\xe0\xb7\xed\x14\x9c\xf5\x9a\xea\x2a\xf5\xa6\xc9\xe2\xd5\x46\xa9\x2a\x06\xfe\xac\xb6\xb7\x38\x2a\xbf\x87\x98\xea\xf5\x06\x81\x7d\xef\x7c\xbc\x5a\xc7\x4f\xa3\xf2\x67\xd2\x9b\x10\xdb\xd1\xc7\x22\x2c\x24\x1e\xb7\xae\x16\xb1\x0c\xe6\xc0\xa3\xcb\xa6\xe8\x74\x6c\x48\x1e\x74\x3a\xfd\x86\xcb\xbb\x2d\x9e\x74\xf4\xd3\x69\x6f\x88\xd1\x46\x55\x9e\x8d\x59\xa9\x33\xdc\x7a\x1c\x12\x07\xfc\xbf\xd5\xd1\x15\xb8\x31\x09\xef\xe1\x34\xcb\xaf\xc1\xd7\xb7\x28\x5f\x88\xa3\x4e\x8a\x16\xc5\x4b\x5e\xac\x3a\x2d\x07\x57\xbb\xc6\x6d\xeb\x94\xfb\x76\x3c\x39\xf9\xc7\xaa\xa9\xd9\xf3\xb7\x55\x17\x9b\xff\x77\xea\xda\xe2\xf7\x4e\x95\x14\x1c\x53\x28\xc7\xd8\xcf\x56\x3a\xc3\x70\xba\x96\xc2\xaf\x84\xa5\x93\x9e\x57\x55\x5c\x29\xbe\xd9\x43\xb0\x1c\x3c\xb5\xb7\x79\x74\xd4\x67\xc7\x67\xb9\xf0\x44\x47\x82\x16\x53\x93\x3a\x28\x8a\xd7\xad\x73\x17\xf5\x86\x92\x71\x55\xae\xe3\xc2\xc6\xf3\x23\xea\xd4\xbf\xd7\x31\xe0\xa2\xa3\x4c\x7c\xbb\xe8\x58\xcb\xc6\x5a\x7d\x0c\xeb\xc8\xb7\xea\xb6\x8a\x57\x59\x9f\x8d\xab\xa3\x5a\xcb\xcd\x45\x71\xda\x79\x5c\x12\x57\x54\xbd\x70\xdb\x87\xce\xca\xe2\x9c\xb1\xe8\x2d\x29\xfc\x6e\x9d\xc6\x45\x6f\x6d\x28\x97\xa2\x8e\xde\xd0\x09\x5c\xb3\xa3\xd6\x45\x6f\xcd\x14\x9e\x57\xd6\xf1\x54\x27\x22\xea\x4d\x29\x34\xc0\xaa\xbc\xe4\x7b\x2f\x7d\xb2\xca\xef\x71\x8b\x15\x1d\x31\x8e\x7a\x53\xc8\xf6\x77\xeb\x44\x08\xe7\xc1\xd4\x51\xdd\x51\x34\x95\x61\x5b\xe1\x58\x31\xf5\x0f\x38\x3b\xe9\x08\x99\x7d\x00\xf3\xe1\x7c\xe2\xa2\xd0\x11\x20\x7b\x13\xd5\x99\x59\xd4\xc1\x02\xff\x0b\x75\x34\x46\xbd\x65\xb2\x75\xec\x19\xb7\x16\x4e\xaf\xf2\x41\xad\x6e\x00\x55\x1d\xcf\x82\xbb\xe0\x10\x58\x45\xae\xcc\xc2\xaa\x65\x7b\x2b\xa5\xf2\x38\x27\xf6\x26\xed\x3f\xb1\x31\x88\x27\xd1\x16\xf0\xcd\xc6\x05\xe4\x37\xeb\x1e\x55\xcd\x2d\xdb\xb1\xfb\xe8\x7c\x30\x97\x0f\x38\xd2\x4b\x63\x58\x17\x5d\x6f\x1d\x81\xc7\x93\xf4\xf2\xa0\xae\xb8\x75\xe0\xa4\x43\x07\xc5\xa7\x8d\x5d\xea\xd4\x63\x52\xe7\xb8\xf5\xb5\x74\xaa\x62\x47\x71\x81\xbd\xc9\x32\x78\x0c\x5b\x75\x3b\xba\x3d\x85\x34\xaf\x1d\xf3\x77\x6e\x6b\xae\x37\xb1\xb6\xb5\x74\xc2\xce\x49\x9c\xd3\xe9\x6d\x3d\x7b\x93\x6c\xf1\xd8\x32\xce\x1e\x9c\x03\x2c\xb7\x74\x82\x02\x7f\xf7\xba\xc6\xff\xd7\xc1\xf3\x35\x2a\xa6\xed\x84\xa7\x71\x1f\x8e\x8a\x3f\x3b\x46\x16\xaf\x2f\xd2\x5e\xf3\x76\x4a\x96\xab\xc2\x57\xe7\x04\x7f\xd9\x5b\x6c\xd1\xf3\xbb\x56\x17\xa5\xde\x8e\xb0\xb7\xb7\xcb\x26\xd7\x74\x95\x5e\x63\xb0\x2e\xcd\xe0\xb8\x07\x4e\xe5\xd2\x94\xf1\xa4\x63\xb1\xcd\x6d\x27\x10\x95\xe7\xd5\xf6\x76\x23\xf5\x8f\x75\x25\x2b\xef\x08\x74\xea\x32\x15\xc2\x74\xd6\xda\x6e\x27\x1c\xea\x9c\x58\x9d\x55\xe8\x4d\xa1\xc2\xb1\x2a\xea\xb4\x68\xf8\x99\x0a\x7e\x3b\x41\x03\xb3\xb1\x87\xd8\x78\xa8\xef\x1e\x55\xf0\x0f\x75\xe1\xe7\x16\xf0\x0f\x23\x9c\xd6\xef\x27\xef\x9b\x3f\x55\xbe\x3f\x3c\xf4\xdb\xab\xf7\x4f\x4c\xfd\xc2\xc5\xfb\x47\x26\xab\x76\x2f\xcb\xea\x33\x8e\x01\xfe\x3f\x00\x00\xff\xff\xe6\x54\xaf\xe2\x00\x30\x00\x00")

func assetsViewsGoBytes() ([]byte, error) {
	return bindataRead(
		_assetsViewsGo,
		"assets/views.go",
	)
}

func assetsViewsGo() (*asset, error) {
	bytes, err := assetsViewsGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/views.go", size: 28672, mode: os.FileMode(420), modTime: time.Unix(1525989347, 0)}
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
	"assets/add.gohtml": assetsAddGohtml,
	"assets/footer.gohtml": assetsFooterGohtml,
	"assets/header.gohtml": assetsHeaderGohtml,
	"assets/index.gohtml": assetsIndexGohtml,
	"assets/view.gohtml": assetsViewGohtml,
	"assets/views.go": assetsViewsGo,
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
		"add.gohtml": &bintree{assetsAddGohtml, map[string]*bintree{}},
		"footer.gohtml": &bintree{assetsFooterGohtml, map[string]*bintree{}},
		"header.gohtml": &bintree{assetsHeaderGohtml, map[string]*bintree{}},
		"index.gohtml": &bintree{assetsIndexGohtml, map[string]*bintree{}},
		"view.gohtml": &bintree{assetsViewGohtml, map[string]*bintree{}},
		"views.go": &bintree{assetsViewsGo, map[string]*bintree{}},
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

