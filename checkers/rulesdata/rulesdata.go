// Code generated by go-bindata. DO NOT EDIT.
// sources:
// rules/rules.go

package rulesdata


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

var _bindataRulesRulesGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\x61\x73\xdb\x36\x93\xfe\x1c\xff\x8a\x2d\x87\x4d\x28\x45\xa6\x1c\x4f\xdb\xe9\xc4\x51\x6e\xd2\xf8\xd2\xf3\x4c\x92\x66\x1c\xa7\xed\x4c\xea\x86\x10\xb9\xa2\x71\x06\x01\x16\x00\x2d\xf1\x52\xff\xf7\x1b\x00\x14\x4d\xd2\x22\x23\xeb\x8d\xdf\xbe\x1f\x3a\xb5\x88\xc5\xf3\x2c\x76\xb1\x8b\x5d\x20\x39\x89\x2f\x49\x8a\x90\x0a\x59\x30\x54\x7b\x7b\x34\xcb\x85\xd4\x10\xec\x3d\xf0\x52\xaa\x2f\x8a\x79\x18\x8b\x6c\xfa\x57\x41\x14\x65\xa5\xc6\x69\x2a\xf6\x8d\x64\x5a\x10\x99\x4c\x13\xc5\xbc\xbd\xd1\xde\xde\x74\x9a\x88\xf8\xa9\x2a\xb2\x8c\xc8\x12\x8e\x51\x63\xac\x15\x24\xb8\x40\x29\x31\x81\x45\xc1\x63\x4d\x05\x07\x46\x35\x4a\xc2\x14\xe8\x0b\xa2\x21\x26\x1c\xe6\x08\x8a\x66\x39\xa3\x0b\x8a\x49\x85\xa3\x49\xaa\x00\x00\x94\x2e\x19\x02\xae\x72\x94\x34\x43\xae\x09\xab\x04\xe6\xb8\x10\x12\xc1\x11\x58\xf4\x60\x04\x9f\x61\x11\x8c\xe0\x3a\x18\x55\x42\x64\xa1\x51\x42\x2d\x14\x8c\xf6\x8c\xa0\xfb\xf9\x81\x33\x92\xcd\x13\x12\x64\x90\x28\x16\xbe\x21\x3a\xbe\x40\x39\x82\xcf\x7b\x0f\x32\xf7\x2b\x88\x3a\xe0\xfe\x22\xf0\xc7\x44\xa6\xca\x72\x44\xa3\x70\xef\xc1\x83\xdf\x2e\x50\x62\x90\x7d\xf4\x16\xde\x79\xf8\x56\x24\x18\x9e\xa8\x20\x3a\x49\x90\xeb\x68\x04\x0f\x1f\x42\x35\x74\x86\x2b\x0d\xdf\xcc\xc0\xcb\x09\xa7\xb1\xb7\x69\x44\x62\x2c\xae\x50\xae\xc7\x0c\x91\x77\x1e\xbe\x14\x5c\x69\x4b\x75\x8a\xc6\x2d\x81\x67\x6c\x26\x71\x29\xa9\x46\x20\x0a\x2a\x2d\x8d\x72\x56\xb7\xc8\x1b\xed\x0d\xac\x21\xbf\x4c\xc3\x5d\x17\xd2\x54\xa9\xfa\x94\x5f\xa6\xde\x79\xf8\xcb\xfc\x7f\x31\xd6\x76\xc6\xbb\xcb\xf4\x2d\xc9\x30\x1a\x6d\xa3\xf3\x5a\x99\x5a\xf1\xeb\xfe\x8d\x94\x4b\x8c\x89\xc6\x04\xa8\x98\x52\x51\x68\xca\x60\xbd\x73\x0b\x45\x52\x54\x77\xdd\x3a\x0e\x24\x3c\x45\x92\xbc\x60\x2c\x90\xdd\x5d\x43\x45\x73\xcc\xee\x1c\x37\xe5\xb8\xd6\x65\x68\xf3\x74\xe0\xfd\x4f\x95\xa1\x2b\x93\x74\xc6\x81\x36\xd7\x38\x81\x42\x61\x43\x01\xa0\x5c\x69\x24\x49\xd4\x72\x6d\x03\xe1\x15\x65\x38\x48\x61\x04\x36\x71\x08\xd5\x18\x1e\x20\xf9\xcd\xb8\xae\x62\x99\x80\xfb\x6f\x33\x5b\x2d\xd9\x43\xd7\x18\xff\xc2\xa2\x8e\xa9\x1c\x5c\xd3\x31\x95\x03\x4b\xb2\xa3\x03\x0c\x6f\x45\xfe\x92\x09\x85\xfd\x1c\xb5\x44\x8f\x73\x1a\xe3\x03\x3c\xc7\x54\xc5\x44\x26\x1b\x19\xaa\xb1\x1e\xfc\x7a\xb4\x46\xef\x0d\x0f\x55\xa8\x9c\xc6\x54\x14\x0a\xb2\x42\xe3\x0a\x98\x88\x2f\xa7\x05\x37\xff\x03\x91\xa3\x24\x26\xf9\x76\x43\x24\xa1\x24\xe5\x42\x69\x1a\x0f\xc5\x49\x56\x84\xaf\x45\x7c\x19\x8c\x8e\xcc\x9f\x1f\x2c\xe6\xad\x14\xdb\x10\x72\xb1\xdd\x14\xb5\xb1\x33\x27\x89\x95\xb8\x1d\x32\xd3\x29\x44\x59\xf1\x24\x02\xc2\x13\xf3\xd7\x61\x04\x44\x22\x90\x24\xc1\x04\xb4\x80\x8c\x5c\x22\xe4\x42\x29\x3a\x67\x08\xd2\x9a\x10\x08\x30\xca\x11\x96\x26\x6d\x41\x94\x1d\x46\xc6\x88\x85\xc2\x04\x82\x25\xd5\x17\x40\x9c\x1e\xc6\x23\xd3\x29\x70\xe1\x7e\x36\xdc\xe3\x67\xc5\x93\x5a\x67\x3f\x2b\x0e\x6b\x7d\x3b\x09\x31\x2b\x9e\xac\xd3\xf4\x6c\x06\xf6\xc3\x61\xf5\xa1\xe5\x53\xb7\x6e\xaa\x20\xa3\x4a\x51\x9e\x4e\x2a\x57\x18\xbd\x2c\xb2\xc9\x5d\x59\x86\x09\x25\x1a\x59\xe9\x58\x5e\xe8\x60\x8d\xd8\xda\x3c\x56\xbb\xd3\x96\x7a\xa7\xff\x01\xfa\x4d\xa7\x90\xd0\xc5\x02\x25\x72\x0d\xdd\xcd\xd5\x63\xdb\x2a\xd7\x7f\x8d\x25\x34\xb6\xb9\x53\x79\x02\x19\x29\xe7\x08\x0e\x17\x96\x44\x01\xe5\x1a\x79\x82\xc9\x7f\xdd\xc5\xc0\x0d\x1d\xef\x4b\xc5\xd3\xad\x75\x34\x46\x16\x85\xd9\xeb\x66\xc2\x36\x76\x7d\xfd\xaf\x68\xec\xd4\xab\xd1\x9e\xd4\x16\xf8\x2a\xe6\x3c\xfd\xba\xba\x9d\x6e\xad\x5c\x6f\xae\xe4\x94\x55\x15\x03\x50\x0e\x17\x5a\xe7\xe1\x5b\x5c\x9e\xe2\x5f\x05\x2a\x53\x91\x32\xa6\x26\xa0\x8a\x34\x45\xa5\x29\x4f\x2b\x09\xf1\x93\x48\x4a\x53\xbb\x10\x0e\x84\x69\x94\x9c\x68\x7a\x85\x77\x2d\x3a\x3a\x74\x81\xf7\xf3\x7f\x9f\x79\x13\x28\x24\x9b\x18\xc5\xba\x69\x75\x48\xbc\xa1\x57\x95\x62\xcd\x17\xf7\x61\xa0\x30\xf1\xba\x98\x7e\x86\xfa\x42\x24\x13\xf0\x2d\xac\x6f\xd4\xf0\xda\xee\xe2\x94\x35\xdc\x65\x7f\x5a\x81\xf7\xce\x4a\x5f\xc4\x6c\xaa\xea\xb5\x0a\xc3\xa6\x71\xd5\x85\x28\x58\x62\xfa\x81\x5c\xae\xfb\x06\x2d\x40\x5f\xa0\xf5\x99\xac\x3c\x34\x17\x49\x39\x54\x2a\xe2\x2a\x97\xa8\x94\x49\x49\xc0\xe8\x25\xc2\xc7\x73\x59\x70\x0c\xd4\xe8\xe3\xc1\xb9\x6b\x3b\x32\x52\x42\x4c\xcc\x31\x5b\xf0\x25\xe1\xa6\xaa\x34\x22\xa0\x18\x8d\x11\x08\x63\x22\xb6\x39\xad\xe3\xdd\x1c\xe5\x42\xc8\x8c\xf0\x78\xd0\xc7\x12\x9e\xce\x5a\xa4\x1d\xa7\xca\x09\x7c\x32\x22\x85\x5e\xfc\x18\x1e\x63\x2c\x12\x3c\x2d\x38\x9e\xf0\xf7\x5a\x52\x9e\x06\x6a\xbd\x0b\xb8\xd0\x68\x78\xdf\x23\xc2\xcf\x02\xa8\x52\x05\xc2\x42\x48\x48\x50\x13\xca\xd4\x53\x6b\x58\xf5\x74\x3a\x6d\xf4\x69\xa9\x60\x84\xa7\xd3\x54\x4c\xad\xbc\x9a\x7e\xf7\xfd\xe1\x0f\x07\x6e\x83\x38\xbb\xde\x50\x0e\xd5\xaf\xd5\x02\x7c\xbb\x82\x4e\xf8\x9a\x2e\xe0\xac\xcc\x5d\x8f\xa0\xac\xd6\xed\x82\x3f\x8a\x05\x57\x34\x41\x69\x8e\x6b\x46\x62\x13\x47\xbe\x0f\xf6\x64\xee\x5b\xb6\xaf\x46\x43\x25\x8e\x0d\x59\x10\x0b\x88\x18\xf2\xc8\x1c\xfc\xa6\xa5\x50\x05\xd3\xe6\x04\x13\xf3\x2b\x9b\x73\x8d\x71\x04\x2a\xfe\x48\xbb\xb2\x41\x21\x57\x1b\x83\xb4\xe3\x33\x86\x3c\x20\x52\x8e\xe0\xd9\x0c\x0e\x3a\xfe\xaa\xc7\x66\x66\xcc\x1a\x52\x31\x91\xe7\xe5\x6b\xe4\x43\x16\x34\xf3\xfc\x4f\x23\x78\x3e\x83\x83\x68\x14\xae\x4d\xe3\xfb\x46\x61\xc2\x96\xa4\x54\xa0\x65\x81\xd1\x68\xc3\xa4\x67\xfd\x73\x16\x84\xa9\x0d\x93\x56\x4e\xf9\xf6\xac\xaa\xc1\x5e\x0b\xcc\xac\x40\xbf\x91\xaf\x08\x2b\x10\xd4\x92\xe4\xb9\x71\x99\xf1\x91\x8b\x18\x53\x90\x71\xa1\xa1\x30\x95\x03\xe4\x44\x12\xc6\x90\x01\x51\x8a\xa6\xdc\x44\xc1\x16\x36\x1e\xeb\x2c\x87\x19\x8c\x57\x47\x30\x5e\x99\x3f\xca\x23\x18\x97\xe6\x0f\x9d\xe5\x1d\x9b\x8f\x57\x93\x6a\xac\x9c\xc0\x78\xe5\xac\x7e\x45\xd8\xfb\x25\xc9\x87\x6c\xee\x1b\x8e\xa7\x33\xf0\xcb\x23\xf0\xcd\x7c\x7f\x75\x04\xbe\x61\x33\x23\xd1\x86\xa6\x74\xff\xa6\x2b\xf5\xcb\x49\x25\xbb\x9a\x80\x5f\x0e\x36\xa4\x6a\x49\x75\x7c\xb1\x6f\x3a\xf6\xfd\xb9\x10\x0c\x94\x26\x1a\x8d\x25\xaa\xab\x0d\x93\x5d\x70\x95\x33\x1a\x53\x0d\x91\xf5\x33\x68\x92\x3a\x13\x6f\x61\x2d\x47\x60\x37\x08\x7c\x0e\xc3\xf0\xba\x63\xa1\x6a\xdc\x0d\xb9\x4d\x69\xbf\x9c\xc9\x62\x30\xae\x5b\xb8\xe0\x8f\x3f\xc1\x75\xbb\x37\x71\x11\x8b\xf0\xa8\x25\x79\xfd\xc8\x45\xef\xfa\xeb\xe7\xeb\x47\xad\x2d\x58\x7d\x36\xe6\xde\x1e\xfa\x46\xba\x0b\x6f\x46\x1c\x45\xaf\x07\xea\x5a\x15\x12\xb4\x79\x0d\xb9\xcd\x33\x26\x45\x2c\x18\x49\xa3\xfa\x7e\x20\x17\xa6\x50\x90\xfd\xed\x4f\xc7\xf4\x73\xb3\x83\xc6\x06\x23\xfc\x49\x08\x16\x78\x73\x6f\xe2\xa2\x6e\x02\xde\x1c\x12\x11\x2b\xaf\x7b\x52\x5f\x11\x09\x73\x30\x3b\xe1\x08\xea\x99\xbf\x12\x19\x3c\x9c\x9b\x49\x9b\x00\xac\xd3\x8c\xec\xb1\xd1\x7f\xc8\x67\x0d\x5d\xfc\xb1\xed\x56\xeb\x46\xb2\x6d\x04\x53\xd1\xb8\x5c\x91\x09\xa5\xed\xf1\xc7\x4a\x53\xb2\xa0\x94\x42\x1e\x41\x9d\x96\x5d\x1c\x37\x15\x6d\x79\xd3\x11\x1e\x17\xae\xb4\xbf\x07\xd2\x35\xf4\x66\xe2\x57\x4c\x10\xfd\xc3\x77\xf7\xc0\x5b\x21\x6f\xa6\x3d\xe1\xfa\x1e\x28\x4f\xb8\xee\xa5\xbb\x97\x35\x5a\xdc\xcd\x94\xeb\xa3\xf6\xab\x73\x3a\xe0\xcd\xa4\x1f\xe8\xbd\xd8\xd5\xc0\xf6\x13\xde\x8b\x65\x1d\xb0\x23\xed\xaf\x3f\xb3\x5c\x97\xe0\xaa\x22\x88\x2f\x30\xbe\x6c\x5f\x76\x9b\xf3\x46\x23\x87\xcc\xe4\x1a\x9a\x50\x91\x11\x4d\x4d\xdb\x51\xde\xb5\x99\x30\xa7\xba\xaa\xca\x92\xce\xe1\x60\xeb\x74\xcf\xa5\x18\xab\x90\x73\xd0\x99\x29\xce\xbf\x54\xb2\xa8\x11\x7c\xe3\x0a\x89\x3b\xd5\x7c\xde\x3a\xbb\x47\xbe\x1f\xb9\x6c\x1e\xf9\xca\x40\xfd\xe1\xfd\xe1\x75\xae\xa5\xd7\x44\xb3\xaf\x47\x34\xbb\x21\xea\xf5\x8d\xc4\xa4\xe0\x09\xe1\xda\xb8\xf7\x0a\xa5\xeb\x12\xe6\xa8\x97\x88\x7c\xed\x33\xc2\x13\xf8\x78\x3e\x2f\xf5\x36\xc7\x74\x2c\xf2\x32\x98\x4f\xaa\x09\x81\x1a\x75\x0f\x86\xb5\x80\xaa\x12\xbe\x23\xf9\xdd\x48\xab\x21\x57\xd8\x79\xfe\xa7\x1a\xd9\x57\xa3\xc6\x6e\xb6\xc5\x4b\xf5\x6c\x52\x42\x74\x23\x13\x99\x96\x29\xf2\xd5\x70\xf1\x62\x75\x50\xe1\x09\x4f\x70\xe5\x5a\xde\xde\xc6\xc8\x76\x43\xdd\xa3\xb3\xd1\x09\x75\xcb\x96\x26\x74\xe0\x7e\x05\xab\xd1\x04\xca\xae\x61\xac\x05\x2a\xb9\x55\xbd\xcc\x72\xf4\x75\xbb\x9f\xc3\xef\x7f\xfc\xe1\xbb\xea\xf6\xde\x50\xbd\x30\xcb\x19\xac\x8f\x36\x2d\xc0\x37\x2b\xf0\xcb\xee\x35\xc6\xca\x3b\x0f\xdf\x15\x12\xab\xc7\x90\xb2\xfa\xb9\x6d\x27\xd4\xb4\x80\x7f\x63\x02\xbf\x1c\x0d\x76\x42\xf5\x5b\x5a\xc3\x71\x55\x6e\xa9\x22\x23\x71\xf8\x76\x8f\x73\x8a\xa6\x5f\x5d\x4a\x92\xe7\xb7\x6b\xa0\x4d\x5b\x7a\x99\x86\x2f\x92\x24\xd8\x7f\xd2\xf5\xd8\x32\x0d\x8f\x05\xc7\xf5\x8d\x6e\x05\xf9\xaa\xe0\x83\x06\xf5\x6f\xf0\x3a\xe6\x5b\xa6\xad\x50\x2f\x79\x1c\xfe\x46\xa8\xfe\x59\x8a\x22\xef\x84\xbc\xd9\x92\xf5\x98\xd5\x02\xdc\x15\x83\x49\xe3\x26\x0f\xb4\xd3\x8b\x3f\x2f\x16\xe1\x99\x2c\x78\x4c\x34\x06\x07\x5d\xe2\x79\xb1\x68\x32\x3b\x3f\xfc\x54\x2c\x16\x28\x37\xf0\xba\x81\xf0\x14\x15\xea\x41\x56\x7b\x99\xf1\x3f\x84\x27\xac\xb2\x4a\x75\xbb\xa1\x5f\x89\x82\x27\xcd\xe8\x35\xb0\xad\xc1\x6a\xd6\x20\xfc\x7a\x67\xbe\xcf\x19\xd5\x6f\xeb\x87\x18\x67\xd7\x26\x72\x4b\xf0\x36\xe4\x6d\xc4\x53\xb7\x6d\x1a\x6f\x3b\x43\xb0\x95\xf4\x0b\xc6\xb6\xc1\x7e\x43\xf2\xa0\xe0\xd4\xb4\x92\xe1\x99\x38\xa3\x9a\xe1\xfa\xf1\x68\x13\x78\x25\x32\x68\x08\xe7\xaf\xca\x0c\x37\xb9\xd7\x0b\xbd\xd1\x26\xbd\x1b\xe2\x43\x0a\x3b\xb1\xed\x4c\xd1\x92\xfd\x82\x21\x9c\x6c\xdb\x0c\x1f\x4c\xe0\x6c\x32\x83\x13\xae\x04\xee\x8a\xfa\x5a\x2c\x87\x51\xad\xc0\x5d\x51\x7b\x5d\xb6\x46\xfd\xb2\xc3\x12\x49\x96\xe1\xb1\x24\xcb\x37\x44\x5d\xb6\x4c\x6b\xfe\xe3\x94\x4d\x80\x66\x24\xc5\xf0\x9d\xe9\xd4\x3e\x5f\xdf\x7e\x5d\xb4\x84\x35\xcc\x6d\xb2\xde\x64\x19\x49\x4c\x71\x95\x87\x2f\x45\x96\x53\x86\xe3\x68\x20\x5f\xae\x65\xdf\x14\x4a\xd7\xf2\x5b\x24\x4c\x89\xd5\x8d\x5e\x9b\x2b\xf0\x62\xfb\x52\x9e\x13\xad\x51\xf2\x5b\x2d\xa3\xc4\xc6\x9c\x06\xe7\xed\x79\x36\xdd\x3a\x41\x23\x37\x94\x6d\x3b\x2a\xf8\x39\xd1\xdd\xec\x97\x13\xbd\xe9\x1f\x16\x44\xe6\x6c\x6d\x31\x57\xf7\xa7\x06\xc3\x3d\x4a\xde\xd6\xb5\xfd\xe4\xd9\x26\x7f\xf7\xcb\xfb\x93\xdf\xef\x5d\x03\xcb\xb2\xe5\xdb\x68\xfb\xe4\xdc\xf6\x3a\xa0\x9b\x25\xd5\x04\x16\x52\x64\x13\xd0\x62\x02\x07\x5d\xaf\x0e\x4a\x9b\x33\x75\xfd\x1a\xfa\x92\x30\xb6\x4d\x21\xb2\x29\x23\xf9\xff\x87\x52\x74\xad\x6a\xbe\x79\xe7\xe1\xaf\x84\x15\x68\x7b\x59\x57\x5e\xf7\x3d\x4f\x11\x99\xc2\xc1\x04\x72\x29\xe6\x64\xce\x4a\xc8\xd0\xd4\xc5\xfb\x4f\xa2\x51\xe8\x1e\x52\x1c\xde\x76\x29\xf2\xdf\xaa\xd0\x97\xcf\xc4\x7f\xd4\x3e\xff\xa4\x2e\x0d\x65\x4c\x71\xc6\x93\xa0\x95\xb7\x23\x2e\xf6\x45\x0e\x6e\xc8\x46\x41\x13\xd4\xbd\x0b\x1b\xb2\xc2\x5e\x64\x0e\x45\xd5\xcd\xcd\x6f\xdf\x3f\xe5\x82\x79\x59\xf5\xce\x37\xb2\xd5\x73\xb1\xd8\xaa\x0e\x5d\xc1\x0c\x56\x30\x86\xc3\x4e\x84\xad\x60\x3c\x83\x43\x17\x47\x0e\xfa\x97\xe1\x3b\x61\x77\xa5\x0b\x8f\xc1\xd8\xeb\x56\xf1\x7e\x73\xa8\x6d\xea\x2b\x57\x8f\x1f\xb7\x8f\xc8\x35\xda\xfe\x4e\x68\xfb\xfb\xdd\x4a\xb5\x56\xce\x2f\x77\xc0\x83\xc7\xb3\xea\xa6\x7a\x93\x86\xbb\x41\xee\xaf\x21\x37\x60\x8e\x77\xc4\x1c\xf7\xab\x39\xdd\x11\x72\xda\x0f\xf9\xed\x8e\x90\xdf\xf6\x43\x3e\xdc\x11\xf2\x61\x3f\xe4\xdf\x3b\x42\xfe\xdd\x0f\xf9\xe7\x8e\x90\x7f\xf6\x43\x3e\x7b\xb6\x23\xe6\xb3\x67\xfd\xa0\xcf\x9f\xef\x08\xfa\xfc\xf9\x80\x8b\x76\x5d\xfd\xc3\x7a\xf9\xd7\x7b\xff\x1f\x00\x00\xff\xff\x6d\x73\x01\x11\x0e\x2b\x00\x00")

func bindataRulesRulesGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataRulesRulesGo,
		"rules/rules.go",
	)
}



func bindataRulesRulesGo() (*asset, error) {
	bytes, err := bindataRulesRulesGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "rules/rules.go",
		size: 11022,
		md5checksum: "",
		mode: os.FileMode(436),
		modTime: time.Unix(1622980644, 0),
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
	"rules/rules.go": bindataRulesRulesGo,
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
	"rules": {Func: nil, Children: map[string]*bintree{
		"rules.go": {Func: bindataRulesRulesGo, Children: map[string]*bintree{}},
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
