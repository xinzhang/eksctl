// Code generated by go-bindata.
// sources:
// assets/aws-node.yaml
// assets/coredns.yaml
// DO NOT EDIT!

package defaultaddons

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

var _awsNodeYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x56\xdd\x6e\xe3\x36\x13\xbd\xd7\x53\x0c\xfc\x5d\x2c\xf0\xa1\x92\xe3\x6c\x9a\x06\xba\xf3\x3a\x6e\xba\x68\xe2\x35\xe2\x6e\x16\xc5\x22\x30\x68\x6a\x22\xb3\xa6\x48\x96\x1c\xca\xf1\x3e\x7d\x41\xc9\x3f\x52\x2c\x7b\xd1\xa2\x40\x79\x63\x99\x33\x73\x86\x73\x78\xc8\x61\x1c\xc7\x11\x33\xe2\x09\xad\x13\x5a\xa5\x60\x17\x8c\x27\xcc\xd3\x52\x5b\xf1\x8d\x91\xd0\x2a\x59\xdd\xb8\x44\xe8\x7e\x39\x88\xfe\x07\x2b\xbf\x40\xab\x90\xd0\x41\x59\x87\x38\x58\xe0\x8b\xb6\x08\x83\xe4\x26\xb9\x00\xb7\xd4\x5e\x66\xe0\x1d\x9e\x85\x5a\x20\xb1\x41\xb4\x12\x2a\x4b\x61\x24\xbd\x23\xb4\x8f\x5a\x62\x54\x20\xb1\x8c\x11\x4b\x23\x00\xc5\x0a\x4c\x81\xad\x5d\xac\x74\x86\x91\xf5\x12\x5d\x1a\xc5\xc0\x8c\xb8\xb3\xda\x1b\x17\x9c\x62\xe0\x36\xab\x70\x59\xc1\xbe\x69\xc5\xd6\x2e\xe1\xba\x88\x00\x2c\x3a\xed\x2d\xc7\xad\x5b\xef\xff\xbd\xea\x37\xa0\x3a\xc3\x38\xba\x08\x42\x0d\x8b\x86\xbd\x89\x0d\x5f\x7b\xbd\xe7\x63\x18\xa3\x33\x57\xe3\xe8\x0c\xdd\x29\x44\xf8\xda\x93\xc2\x51\xef\x07\xe8\xad\x19\xf1\x65\xf8\xc8\x91\x7a\xcf\x6f\x53\xe0\x2b\xa1\xaa\x68\xec\x4a\x96\x31\x2c\xb4\x72\x48\x67\x90\x9f\xa3\xb7\x5b\x58\xee\x88\x9d\xa1\x2d\x05\xc7\x21\xe7\xda\x2b\x3a\xc7\x2d\x1c\x8a\x48\xab\x3d\x8e\xdd\xc6\x11\x16\x47\xd8\xff\xad\x3c\x3e\x08\x95\x09\x95\x9f\x55\x89\x96\xf8\x88\x2f\xc1\xb2\x23\xfa\xcc\xaa\x23\x80\x63\x0d\x1e\x61\x3a\xbf\xf8\x03\x39\x55\xe2\xeb\x64\xf6\xef\xf1\x59\x43\xdc\x56\x7b\x3b\x43\x6a\xf1\x7b\xd0\xc3\x9e\x85\x7f\xb2\x6d\x00\x92\x2d\x50\x56\x32\x02\x58\xdd\xb8\x98\x19\xd3\xac\xc8\x20\x0f\x36\x6f\x32\x46\x38\x23\xcb\x08\xf3\x4d\xed\x4d\x1b\x83\x29\x3c\x6a\x29\x85\xca\x3f\x57\x0e\x11\x80\x43\x89\x9c\xb4\xad\x7d\x8a\x20\xbd\xfb\x46\x8a\xae\x24\x00\x84\x85\x91\x8c\x70\x1b\xd4\x28\x24\x0c\xd9\x8a\xef\x46\x08\x83\x29\xa5\xa9\xda\xb5\x86\xb3\xe3\x4b\xcc\xbc\x44\x9b\x30\x69\x96\x2c\x39\x28\x2f\x28\x88\x5b\x41\x82\x33\x19\x1b\x9d\xa5\xf0\xee\x5d\x15\xb6\x2b\xba\xfa\x6e\x6d\xe0\xe4\x2d\xad\x61\x2c\xb5\xa3\x09\xd2\x5a\xdb\x55\x0a\x64\xfd\x6e\x9e\xb4\x44\xdb\x5e\x4e\x0c\xda\x84\x39\x6d\x53\x18\xbf\x0a\x57\x9d\xd7\x30\xb8\x56\xc4\x84\x42\xdb\x70\x15\x05\xcb\x31\x85\xeb\x8b\xcb\xab\x8b\xc1\xe0\xea\xfd\xd5\x8f\x97\x49\xb6\xb2\x09\x72\x9b\x78\x17\xaf\xd1\x51\x7c\xd9\xbe\xcd\xfa\xf5\xbf\x38\x30\xc4\x95\x48\xcb\x41\xf2\x3e\xb9\xdc\x73\x51\x21\x4e\xbd\x94\x53\x2d\x05\xdf\xa4\x30\x94\x6b\xb6\x71\x7b\xbb\xd1\x96\x1a\xd4\xc5\x87\x65\x4d\xb5\xa5\x14\xae\x07\xd7\x3f\xdd\xec\xcd\x3b\x95\x15\x48\x56\xf0\x03\xca\x91\xf6\xea\x81\xaa\x4c\x1b\xb1\xf1\xd6\x6f\xf8\x65\x36\x7f\x9a\x8e\xe6\xbf\xde\xcc\xe6\xa3\xc9\xc7\xf9\xfd\xa7\xbb\xfb\xf1\xd3\xf8\xbe\xe1\x0a\x50\x32\xe9\x31\x85\xdb\xf1\x87\xcf\x77\x1d\x18\x0f\xbf\xcf\x27\x9f\x6e\xc7\xf3\xc9\xf0\x61\x7c\x1c\xf7\xb3\xd5\x45\xda\x9a\x06\x78\x11\x28\xb3\xed\xf1\xef\xb0\x4c\x19\x2d\xd3\x4a\x07\x49\xa8\x21\x6c\x7b\x47\xda\x2f\xc3\xdf\x46\xbf\x54\x49\x67\xd3\xe1\xe8\xdf\xcc\xbc\x3b\x01\xc9\xfe\xd8\xee\xbd\x5b\x37\xff\x61\xf2\x4f\x8f\x8e\x5c\x1b\x94\x1b\x9f\xc2\xe0\xa2\x38\x9c\x05\xe4\xde\x0a\xda\x8c\xb4\x22\x7c\xa5\xa6\xb7\xb1\xa2\x14\x12\x73\xcc\x5a\x1a\x06\x28\xb5\xf4\x05\x3e\x04\xf5\xb7\xa4\x51\x84\x99\x7a\xb5\xfd\x70\x02\xfa\xda\x50\x9f\x2b\xd1\x5f\x08\x75\x24\x11\xae\x44\xbc\x10\x2a\xce\x84\x3d\x07\x81\xc4\x2b\x08\x85\x94\x64\x9d\x20\x0a\xe9\x7b\x20\x25\xb3\x7d\xa9\xf3\xa3\x70\xa9\xf3\x33\xa1\x21\xca\x7a\xd5\xcf\x34\x5f\xa1\x4d\x9c\xe6\xab\x23\x84\xda\xd6\x30\xd5\xdc\x34\x8e\xec\xe9\x6a\xc3\xd2\xaa\x54\x4d\xce\xeb\xd4\xc7\xc4\xc5\x67\x2a\x3e\x03\xd4\x45\x5f\x7c\xa2\xfa\x33\x30\x6d\x02\xe3\x53\xc5\x7f\x17\xe3\x2d\x9d\x6f\x9f\x08\xcc\x88\x43\x17\x3b\xd1\xd2\xbd\x23\x5d\x3c\x6e\x25\x7f\x8b\x2f\x42\x89\x70\xa1\x76\xf4\x3a\x54\x82\x6b\xf5\x22\x72\x97\x74\x3f\xf4\x76\xb7\xba\xe3\x3a\xf4\xad\x6d\x23\x8f\x00\xf2\xba\xf7\x9f\x7a\x1e\x96\x87\xe7\x52\xd5\x41\x06\xbb\x6e\x5a\x97\x6d\xa4\xb7\x4c\x36\xf3\xd7\x3d\x44\xa8\xdc\x4b\x66\x1b\x86\xba\xc5\x56\x75\x8d\x27\x1f\x47\xf5\x5c\x14\xfd\x15\x00\x00\xff\xff\x5c\x3e\x17\xa3\x59\x0b\x00\x00")

func awsNodeYamlBytes() ([]byte, error) {
	return bindataRead(
		_awsNodeYaml,
		"aws-node.yaml",
	)
}

func awsNodeYaml() (*asset, error) {
	bytes, err := awsNodeYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "aws-node.yaml", size: 2905, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _corednsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x5b\x6f\xdb\x3a\x12\x7e\xf7\xaf\x20\xf4\xbc\x52\xec\x26\xe9\x66\xf9\xd6\xda\xde\xae\x81\xd6\x35\xe2\xa4\x2f\x8b\x45\x30\xa6\xc6\x16\xd7\x14\xc9\xe5\xc5\xb5\xbb\xe7\xfc\xf7\x03\xea\x66\x4a\x71\x8a\x14\x2d\x70\x8e\x9e\x24\x72\xf8\xcd\x70\x2e\xdf\x8c\x40\xf3\x2f\x68\x2c\x57\x92\x92\xc3\x64\xb4\xe7\x32\xa7\x64\x8d\xe6\xc0\x19\xbe\x63\x4c\x79\xe9\x46\x25\x3a\xc8\xc1\x01\x1d\x11\x22\xa1\x44\x4a\x98\x32\x98\x4b\xdb\x7c\x5b\x0d\x0c\x29\xd9\xfb\x0d\xa6\xf6\x64\x1d\x96\x23\x42\x04\x6c\x50\xd8\x70\x84\x10\xdc\xdb\x0c\x4a\xf8\xa6\x24\x7c\xb5\x19\x53\xe5\x15\x53\xa5\x56\x12\xa5\x8b\xb1\x08\xd9\xdf\xd9\x14\xb4\x6e\xb0\xc2\x6a\x9a\xa6\xa3\xd8\x46\xb3\x01\x96\x81\x77\x85\x32\xfc\x1b\x38\xae\x64\xb6\xbf\xb3\x19\x57\x57\x87\xc9\x06\x1d\xb4\x57\x98\x0a\x6f\x1d\x9a\x7b\x25\xb0\x67\x7f\x6c\x56\x50\x62\x24\x3a\xac\xce\x6f\x94\x72\xd6\x19\xd0\x9a\xcb\x5d\xad\x28\xcd\x71\x0b\x5e\x38\xfb\xb3\xb7\x68\xfd\x56\x7b\x87\xb6\xc2\xc6\x0b\xb4\x74\x94\x12\xd0\xfc\x83\x51\x5e\x57\x86\xa5\x24\x49\x46\x84\x18\xb4\xca\x1b\x86\xcd\x1a\xca\x5c\x2b\x2e\x2b\x5b\x52\x62\xeb\x08\xd5\x1f\x5a\xe5\xf5\x4b\x17\x8c\xf0\x79\x40\xb3\x69\xce\x0a\x6e\x5d\xf5\xf2\x15\x1c\x2b\x5e\xa7\x4f\xaa\x7c\x08\xb3\x43\xf7\x2b\xe2\xf1\x9e\xcb\x9c\xcb\x5d\x2f\x2c\x20\xa5\x72\xd5\xf1\x26\x36\x97\x70\x7b\xe1\x02\xef\x94\xd7\x39\x38\xa4\x24\x71\xc6\x63\xf2\x97\x8b\xae\x12\x78\x8f\xdb\xea\x7a\x8d\xbf\xbf\xe3\xaf\x11\x21\xcf\x33\xf7\x05\x64\xeb\x37\xff\x45\xe6\xaa\xd4\xb9\x58\xb1\xaf\xae\xd3\x61\x38\x3b\x0a\x98\x2a\xb9\xe5\xbb\x4f\xa0\xff\xd4\xea\x6f\xf5\x4e\x95\xc1\x2d\x17\x48\xc9\x6f\x95\x64\x46\x6f\xaf\xc9\xff\xab\xd7\x4a\x83\x31\xca\xd8\xee\xb3\x40\x10\xae\xe8\x3e\xcf\x89\x40\x58\xed\xdb\x4c\x28\x06\x22\x02\x20\x55\x0d\x11\x2e\x2d\x32\x6f\x30\x5a\xf7\xda\x3a\x83\x50\x46\x4b\x5b\x10\xc2\x15\x46\xf9\x5d\x41\xb8\x4c\x21\xcf\x4d\x06\x46\x03\xe1\xfa\x6d\xf5\xd2\xc9\xfe\xde\xbd\x69\xa3\x4a\x74\x05\x7a\x4b\xe8\x3f\x26\xb7\xd7\xf1\xc6\xf1\x44\x32\x72\x85\x8e\x5d\x85\x12\x14\x87\x8c\x29\xb9\xed\x04\x18\xb0\x02\xc9\xf5\x78\x54\x03\x0e\x03\x86\x47\x87\x32\xbc\xda\x41\xc1\xcd\x50\x0b\x75\x2a\xf1\x57\xf0\xf7\xa5\x94\x1f\x16\x58\x0d\x9c\x84\x48\xcd\x96\xeb\xe4\xf5\x91\xb7\x1a\x19\xad\xf8\x47\x0b\xce\xc0\x52\xf2\x66\x44\x48\xa8\x55\x87\xbb\x53\x6d\x80\x3b\x69\xa4\xe4\x5e\x09\xc1\xe5\xee\xb1\xaa\xfa\x9a\x25\xe2\x15\xda\xf8\xac\x84\xe3\xa3\x84\x03\x70\x01\x9b\x90\x32\x93\x00\x87\x02\x99\x53\xa6\x96\x29\x03\x0d\x7e\x8c\x2e\xf8\xd2\x15\x5f\x9d\xbc\x0e\x4b\x2d\x3a\x1b\x62\x87\x87\x47\xf4\x54\xbd\xac\xec\x07\x6a\xa5\xf5\x5a\xf5\xde\x2b\xfe\xe5\x20\xc2\x75\x96\x71\x65\xb8\x3b\x4d\x05\x58\xbb\x8c\x28\x25\x0d\x34\x9f\x32\xc3\x1d\x67\x20\x1a\x69\xd8\x6e\xb9\xe4\xee\x74\x36\x38\x48\xbd\x7b\xb6\x1a\x62\xf6\x3f\xcf\x0d\xe6\x33\x6f\xb8\xdc\xad\x59\x81\xb9\x0f\x01\x59\xec\xa4\xea\x96\xe7\x47\x64\x3e\x30\x5d\x7c\xb2\xc6\x5c\x37\x61\x79\x40\x53\xda\xfe\x76\x5a\x47\x69\x7e\xd4\x06\xad\x3d\x37\x86\x58\x62\x8f\x27\x4a\x92\x90\xf5\x83\xe6\xa0\x6c\x32\x10\x26\x44\x69\x34\x10\x52\x80\x2c\xe4\xb3\xcd\x03\x08\x8f\xcf\x34\xd4\xbd\x53\xfa\xe3\xeb\x35\x83\x61\xc5\x2f\xd3\x0d\x65\xfe\xf6\xa6\x59\x77\x4a\x04\x8c\xd8\x11\xad\x19\xd3\x26\x7c\xef\xf2\x5c\x49\xfb\x59\x8a\xd3\xd9\x82\xb3\xe6\x64\x7e\xe4\xd6\x75\x8e\x61\x4a\x3a\xe0\x12\x4d\x04\x37\x24\x87\xfa\xe1\x25\xec\x90\x92\xb7\xe3\x37\x37\xe3\xc9\xe4\xe6\xfa\xe6\xf6\x4d\x96\xef\x4d\x86\xcc\x64\xf7\xf3\x0f\x8b\xcf\xcb\x41\xca\xe2\xde\x5e\x35\x20\xf4\x30\xc9\x26\xd9\x75\x1f\x6b\xe5\x85\x58\x29\xc1\xd9\x89\x92\xc5\x76\xa9\xdc\xca\xa0\xc5\xaa\x6d\xd5\x4f\x6f\x14\x69\x1f\xc1\x4b\xee\x06\x6e\x2a\xb1\x54\xe6\x44\xc9\xe4\xef\xe3\x4f\x7c\x90\x97\x68\x87\xd2\x4c\x7b\x4a\x26\xe3\x71\x79\x11\xa3\x07\x01\x66\x67\x29\xf9\x37\x49\xd2\x40\xc6\xc9\xdf\x48\x52\x11\x74\x73\xab\xab\xb6\x1f\x25\xe4\x3f\xdd\x91\x83\x12\xbe\xc4\x4f\xa1\x04\x23\xbd\x67\xa7\x86\x7e\x9a\xd6\x42\x91\xfe\x32\xc8\xaf\xc0\x15\x94\xc4\x1a\x7a\x77\x81\x3c\xc4\x94\x92\x30\xe5\x9c\x1b\x87\x32\x7d\x3d\x5d\x40\x57\xca\x38\x4a\xa2\x1e\xd3\xb2\x7e\x1f\x57\x1b\xe5\x14\x53\x82\x92\xc7\xd9\xea\x47\x71\x52\xc7\xf4\x45\xac\x87\xe9\x77\xb0\x7a\x9d\xaf\x45\x2b\xd1\x19\xce\x2e\x5b\x16\xa3\x55\xad\x39\x70\x98\x92\x0e\x8f\x2e\x0e\x2d\x08\xa1\xbe\xae\x0c\x3f\x70\x81\x3b\x9c\x5b\x06\xa2\xaa\x14\x1a\x7a\xb5\x8d\xdd\xcd\x40\xc3\x86\x0b\xee\xf8\xb0\xe2\x20\xcf\x87\x04\xb4\x9c\x3f\x3c\xbd\x5f\x2c\x67\x4f\xeb\xf9\xfd\x97\xc5\x74\xde\xdb\xce\x8d\xd2\xc3\x03\x20\xc4\x85\xc0\xdd\x2b\xe5\xfe\xc9\x05\x36\x43\x5c\x3f\x8c\x82\x1f\x50\xa2\xb5\x2b\xa3\x36\x18\xe3\x15\xce\xe9\x0f\xe8\xfa\x2a\x74\x9d\x28\x83\x01\x87\x34\xe9\x40\xc9\xdd\xf8\x6e\xdc\x5b\xb6\xac\xc0\xe0\xe4\x7f\x3d\x3c\xac\xa2\x8d\x40\xe4\x1c\xc4\x0c\x05\x9c\xd6\xc8\x94\xcc\x6d\x28\xf0\x48\xc2\xf1\x12\x95\x77\xdd\xe6\x6d\xb4\x67\x3d\x63\x68\xed\x43\x61\xd0\x16\x4a\xe4\x75\x8b\x6d\x9f\x2d\x70\xe1\x0d\x46\xbb\xed\xd9\x5c\xda\xb6\xec\x67\xf5\xe8\xdd\x6c\xd4\x55\xf1\x03\x55\xc3\xda\xe9\x74\xd0\x52\x2e\xf2\x57\x75\x61\x87\xcf\x1b\x4c\xc5\x9e\x6d\x29\x0f\xe8\xb7\xf6\x74\xb7\xf9\xe2\x9c\xdc\x0c\xde\x17\x66\xac\xc1\xef\xc1\xc5\x21\xeb\xd9\x6f\xcf\x79\x4e\x0c\xcd\xa4\x0e\x6a\x12\xca\x26\xb9\xb0\x6d\x99\x01\xfd\xe2\xef\xcf\x2b\x66\xb6\x66\x1c\x4e\x9b\x01\x22\x42\xfa\xe9\xe9\xae\xd3\xda\x0e\x2a\xfd\x09\xec\x92\x75\x8d\x35\x8b\x15\x25\xb3\xe5\xfa\x69\xfa\xf1\x71\xfd\x30\xbf\x7f\x5a\x84\xc4\xed\xd8\x2e\x1d\x70\x99\x8e\x49\x6a\x48\x69\xe9\x05\xc2\x7a\xe1\x40\x60\x9a\x3f\x02\x00\x00\xff\xff\x17\xfe\xaa\x4f\x0c\x11\x00\x00")

func corednsYamlBytes() ([]byte, error) {
	return bindataRead(
		_corednsYaml,
		"coredns.yaml",
	)
}

func corednsYaml() (*asset, error) {
	bytes, err := corednsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "coredns.yaml", size: 4364, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
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
	"aws-node.yaml": awsNodeYaml,
	"coredns.yaml": corednsYaml,
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
	"aws-node.yaml": &bintree{awsNodeYaml, map[string]*bintree{}},
	"coredns.yaml": &bintree{corednsYaml, map[string]*bintree{}},
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
