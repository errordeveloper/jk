// Code generated by vfsgen; DO NOT EDIT.

// +build !dev

package lib

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Assets statically implements the virtual filesystem provided to vfsgen.
var Assets = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2018, 11, 29, 13, 15, 48, 178094962, time.UTC),
		},
		"/std.js": &vfsgen۰CompressedFileInfo{
			name:             "std.js",
			modTime:          time.Date(2018, 11, 29, 13, 43, 24, 178432097, time.UTC),
			uncompressedSize: 20092,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5c\x6f\x73\xa3\x38\xd2\x7f\xff\x7c\x0a\x92\x17\x29\xb4\x60\x0a\x88\xc7\x9b\x35\x96\xa7\x66\x76\x93\xbb\x5c\xcd\x65\xb6\x92\xcc\x5e\xdd\xa5\x52\x2e\xc5\x88\x44\xcf\x12\xe1\x05\x31\xb9\x39\x9b\xe7\xb3\x3f\x25\x89\x3f\x02\x84\xff\x64\xb3\x77\xf7\x66\x26\x80\xfa\xd7\xad\x56\x77\xab\xd5\x92\xfc\x15\xa5\x06\x83\x6c\xb3\x59\x17\x01\x73\x2e\x92\xf4\x19\x31\xb8\xfe\x90\xb3\x64\xea\xda\x7f\xb9\xf9\x7c\x35\xf5\xec\xbf\x7f\xf8\xeb\xa7\xa9\x6f\x5f\xa3\x97\xe9\x29\x6f\xf5\xb7\x94\x30\xfc\x21\x7d\xcc\x60\x94\xd3\x25\x23\x09\x35\xc1\x9a\x3d\x91\xcc\x79\x78\x80\x34\x8f\xe3\xa0\x7c\x58\xac\x92\x0c\xba\x2d\x12\x67\x95\x26\x2c\x61\xdf\x56\xd8\x59\x2c\x08\x25\xac\xc1\x60\x36\xae\x61\x04\x25\xab\x70\x20\x0e\x52\xcc\xf2\x94\x1a\xfc\x45\x1b\xef\x11\xb3\xeb\x24\x61\x1f\x32\x8d\x54\xd8\x4e\xc1\x5a\x52\x9a\xe9\x66\x43\xf1\x8b\xa1\x90\x82\x52\x02\x13\x3b\x29\x46\xe1\x25\x65\xa7\xbe\x89\x9d\x55\x92\x11\xd9\x27\x60\xa9\x4f\x36\x06\x43\x3d\x59\x21\xf6\xa4\xf4\x03\xac\xb9\x56\x31\x2c\xa5\x77\x16\x8b\x24\x8a\x32\xcc\x4c\xa5\x73\xf6\x18\x54\x7d\xc2\xef\x9b\x86\x19\x4b\x09\x7d\x54\x1b\x5a\xd8\x66\x60\xca\xb5\x3a\xc4\xfd\x2b\x8a\x73\x7c\x28\xfb\xc9\x9b\xb1\xe7\xff\xa8\x86\xb0\x07\xf3\x33\x0d\xf3\x72\x08\xce\xda\xcc\xc1\xb4\x32\x4a\x87\xdb\xe4\x90\x0c\x84\x86\x98\xb2\xae\x14\x6c\xbb\x14\x9e\x5b\x8b\xc1\xb6\x8b\xc1\xc0\xb4\x63\xc5\x19\x43\x29\xd3\x58\x1c\x03\x6b\x26\x3f\x7e\x7e\xf8\x5f\xbc\x64\xe6\xb8\x63\x34\x28\x0c\x7f\x6e\x5b\x8b\xb0\x7a\xfe\xfe\x82\xe0\x38\xfc\x2c\xe5\x74\x6d\x6c\xbb\x7d\xd2\x5f\x3a\x43\xad\xa5\xf5\xf4\xb4\xb7\xad\x71\x12\xae\x81\x6b\x52\xd1\x65\xdf\x4e\xed\x96\xba\xfb\x20\x97\x1d\x45\x77\x24\x10\x30\xa7\x1a\xfe\x98\x86\x7a\x6d\x95\xd6\xc2\x1b\x94\x1a\x6b\x6c\xa3\x08\xe4\x57\x2c\xa2\x13\x76\x04\xf1\xfa\xea\xf3\xd5\xf9\xd4\xb5\x6b\xb8\xa9\x67\x5f\x63\x14\x8a\x3f\x7d\xfb\x47\x44\x97\x38\x16\x0f\xa7\x9c\xe6\xaf\x38\xcb\xd0\x23\xde\x37\x52\xd5\x04\x6f\x12\xa7\x1a\xb4\x3a\x4a\xf5\xe4\x61\xfd\x18\x55\x93\xd5\x11\x8a\x29\x11\x8a\xb5\x22\x94\xfa\x64\x33\xa0\xef\x01\x4a\x1f\xb3\x5b\x8d\x9b\xee\x70\x90\xf1\x80\x7f\x7c\x21\x1a\x07\x91\xc3\xe3\xf0\xd1\x19\x16\xe2\x6d\xa2\x54\x4e\xa5\xe2\x3a\x91\x42\x06\xa9\x86\xb5\x70\xc3\xbe\xbe\xbb\x1e\xea\xb7\x94\x86\x42\x61\x49\x6d\x6d\x89\x31\xea\x58\xb9\x6b\xa7\xb6\xd2\x67\x1d\xc8\xbe\x8e\xda\x10\x62\x1a\xea\x04\xde\xed\x24\x29\x5c\x17\x41\xea\xdc\x5c\xfe\xe3\xfc\xf3\xc5\xe2\xe6\xcf\x9f\xaf\x6f\xa1\xdf\xbc\xb8\xbc\xba\x85\xe3\x20\x75\x2e\x2e\x3f\x9d\x2f\x2e\x7f\x3a\xbf\xba\xbd\xbc\xb8\x3c\xbf\x5e\x7c\x3a\xbf\xfa\xd3\xed\x9f\xc5\xa7\x73\xba\x4c\x42\x42\x1f\xe1\xfa\xcb\xed\xc5\xd9\xe2\xe3\xdf\x6f\xcf\x6f\xa6\x9e\xfd\xe5\xf6\xc2\x9b\x2c\x6e\x6e\xaf\x2f\xaf\xfe\x34\xf5\x39\x0f\xc2\xad\x10\x72\x3b\x15\xf6\xf8\x21\x4d\xd1\x37\xd3\x07\x41\xea\x44\x71\x82\xaa\x6f\x17\xf2\x6f\xf9\xb5\x24\x72\x1e\xf2\x28\xc2\x69\xdd\x74\x32\x6e\x9a\x4e\xc6\x43\x4d\x49\xf6\x89\x30\x16\xe3\x73\x1a\x12\x44\x05\x05\x37\x40\x6f\x22\x09\xaa\xe7\x33\xf9\x78\xe7\xd9\xee\x3d\xa8\xc8\xef\xdc\x7b\x08\xa1\x17\xa4\xce\xa7\x84\x3e\xea\x5c\x38\x4e\x5e\x20\xdb\xb8\xd2\x83\x9f\xc8\xe3\x13\xc4\x1b\xb7\x28\x09\x9c\x65\x8a\x11\xeb\x06\xdc\xca\x27\x20\x74\x4f\x4e\x30\x84\xee\xfb\xb2\xf5\x3f\xce\xaf\x3f\x4f\xb9\x3c\xf2\x59\xb4\xae\xa1\x94\x09\x33\x29\x7b\xac\xba\x63\x19\x01\x2a\x99\xe6\xf3\xb9\x0b\xac\x5a\xa8\xef\xc6\xfe\x0f\xe3\x1f\x26\xdf\xfb\x3f\x4c\x34\x80\xf8\xb7\x1c\xc5\x6d\xdf\x52\xc2\x90\xe8\x22\x64\xfc\xbf\x93\x93\xa6\x9b\x90\x89\xff\x6b\x38\x2e\x3c\x54\x84\x77\x6d\x3e\x3d\x3a\x1f\x73\x12\x87\x38\x6d\x81\x93\xc8\x3c\xaa\xad\xd2\x73\xfd\x71\x81\xe3\x0c\x57\x56\x5a\x54\xc1\x30\x75\x3e\x7e\x63\xf8\xa3\x18\x09\x07\xc5\x71\xb2\x44\x0c\x9b\x18\x48\x5d\x67\x2b\xb4\xc4\x10\xcb\x87\x67\x42\x51\x4c\x1e\x29\xf4\xe4\xf3\x57\x86\x1e\x62\xac\x44\x68\xf9\x62\x41\xe8\x22\xcf\x30\x2c\x87\x8b\x64\x57\x38\x63\x38\x84\x11\x8a\xb3\x12\x29\x11\xfe\xb1\x10\x0e\x5e\xb5\x93\xc4\x19\xbc\xbb\x2f\x9f\xf1\x92\x25\xe9\x82\xe6\xcf\x0b\x1c\xe3\xe7\xac\x6a\x17\x25\xe9\x12\x2f\x42\x1c\xa1\x3c\x66\x99\x44\x2d\x1a\x1d\x28\x1a\x17\x2d\x7f\xaa\x1b\xaa\xb1\x45\x03\xc4\xf4\x20\x21\x62\x48\x6a\xa7\x6f\x08\x46\xa9\x44\x3d\x25\xca\x1a\x8b\x1f\xa6\x75\x1e\xbe\x31\x9c\x99\xc0\xc9\xf2\x07\x24\x9c\xa3\xfa\xa0\x4e\x19\xbd\x57\xd2\xea\xca\x60\x0c\xc0\x90\x04\x82\xf9\x2e\xf1\x6b\x11\xa4\x3f\x3a\x59\x4c\x96\xf8\x0d\xc5\x58\xa5\x78\xd5\xf1\x4e\x12\x99\x6c\xde\x32\xaa\x72\x50\x6a\x1b\x63\x05\xb7\x54\x02\xff\xaf\x16\x64\x89\x56\x68\x49\xd8\x37\x13\x8c\x1a\xdb\xb4\x30\xb0\xbc\x13\x36\xf2\x82\x97\x27\x12\x97\x52\x8b\x2f\x33\x62\x31\x0b\x4b\x07\xa0\xb0\x0f\x12\x28\x1e\x50\x8a\xfc\x98\x26\x2f\x8d\x37\x54\x8c\x55\x57\xb0\x34\x40\x23\x2a\x9d\x69\x85\x42\x93\x0c\xa9\x00\x85\x2d\x03\x8c\x92\xd4\x94\x9e\xe8\x06\x78\xc6\x02\x6c\x59\x75\xb6\xe2\xbc\xf0\xa4\x49\x4c\x5f\x23\xa5\xa7\x7c\x12\xd2\x83\xd7\xed\xfb\x36\xde\x42\x6b\xb0\x46\xd0\x13\x69\xc8\x36\x34\x6f\xb2\x15\xce\x9b\xb4\xf0\xfc\x9d\x78\xa7\xfe\x56\x3c\x9e\x33\x29\x78\xe3\x9d\x78\x6a\x64\xd6\xe0\x4d\xc6\x2d\xbc\xb3\xed\x78\xe5\x44\x38\x8c\x58\x36\x38\x44\xc6\xde\xfc\xa1\xc5\xdc\x5b\x4e\x91\xd7\xeb\xc6\x98\xbb\x97\xe9\xf1\x89\x40\x3c\x2a\xe3\xbd\x15\x49\x37\xbe\x02\xca\xef\x41\xf1\xb1\xde\x8a\xa5\xd3\x9c\xc0\x1a\xf7\xb0\xb8\x0e\xb7\x62\xe9\x34\x26\xb0\xce\x7a\x58\x5c\x77\xc3\x58\x83\x63\xaa\x93\xac\x1e\xdf\x1d\x78\xfb\x4a\x57\x8f\xed\x16\xbc\x2a\x4b\x6d\xc5\x46\x9e\xc4\xf2\xe8\xd8\x9f\xa1\x36\x1b\x7c\x04\xd3\x92\x67\x69\x0d\xcd\x44\x1d\x27\x8c\x33\xdb\xc9\xad\x35\xec\x07\xb2\xf3\x26\x87\xf3\x6b\x0d\xc0\x81\xfc\x4e\xfd\xc3\xf9\xb5\x06\x68\x3b\xbf\xa3\x2a\x29\x33\x53\xd0\xe2\x3b\x19\x1f\xc8\xb7\x6f\x6a\x87\xf4\xb4\xb2\xbd\x57\xf0\x3c\xa0\xb7\x1a\x9e\x07\xf7\x53\x2e\x89\x5e\xc9\xb2\x5c\x4f\x1d\xc6\xf1\x86\xa5\xf9\x52\xcb\x51\x01\xa7\x22\xc1\xdc\x13\x99\x96\xd9\x68\x3b\x55\x66\x47\xb0\x9d\xca\xac\xd9\x53\x9a\xbc\x18\x3c\xd7\x3e\x4f\xd3\x24\x35\x8f\x2f\x62\xc4\x64\x66\x90\x4d\x8d\x4c\xc8\x65\x3c\xe7\x19\x33\x1e\xb0\x91\xe1\x94\xa0\x98\xfc\x0b\x87\x06\xa1\x31\xa1\xd8\x39\x1e\xe4\x9f\xb0\xab\x8e\x08\x8d\x06\xab\x64\x79\x17\x7b\x99\x40\xd7\x6c\x11\x87\x91\xc2\xd0\x44\x08\x24\x7b\x39\x2c\x05\xd7\x51\x3f\x92\xc9\x1c\xfc\x8e\xdd\xb7\x95\xa1\x87\x48\x3a\xc6\xd0\x4b\x2b\xb5\x09\x9b\x8a\xd5\xce\xb7\xb4\x8b\x69\x25\x5f\xe3\x63\x7e\x72\xea\xfb\x9e\xef\xbf\x1b\x7f\xef\xef\x52\xd1\x12\x51\xae\x0b\xce\xc3\x90\x89\xad\xf1\x80\xbf\x25\x34\x34\x7c\xe3\x91\x3c\x22\x91\xf4\x72\x05\xc9\x3c\x13\xcf\x66\x5e\x20\x73\x45\xfd\x92\x88\x80\x80\x3a\x19\x66\x3f\x57\xc9\x2f\x19\x61\xfe\xaa\xce\xdf\x31\x33\x59\xf5\x64\x8b\x8f\xa5\x3e\xe8\xa0\x89\xf7\xfc\x49\x9d\x51\xd4\xba\x80\x76\x1a\x55\x87\x68\xc4\x2c\xb5\xfd\xc0\x90\x29\x15\x95\x3e\xd3\xda\x30\xa5\xae\xd5\x25\x9e\x58\xe3\xb5\x6c\x04\xde\xdd\x17\x9a\x35\x1f\x0b\x86\x33\xdb\xd2\xb8\xf0\x3d\x74\x8b\xf6\xca\x90\xa5\xb9\x6e\x61\xb8\x87\x0d\xd6\xb5\x16\x9d\x33\xa9\xc2\x6f\x36\x47\x07\x39\x58\x0d\x6c\x2c\x51\x1c\xe3\xd0\x78\x21\xec\x29\xc9\x99\xa1\xa8\xf0\x18\x14\xed\xc9\xca\x05\x81\x52\xb0\xab\x24\x0f\x94\x12\x5a\x4b\x5b\x23\x4f\x68\x2b\xc0\x73\xe8\x96\x8b\xfe\x46\x47\x5c\x7d\xa3\x11\x58\x57\xe6\x69\x29\x8d\xe5\x97\xf6\xc4\xdc\x26\x3f\x82\xee\x7b\x36\x6a\xbf\x9b\xba\xd2\xd6\x29\xf4\x83\x0e\xed\xa8\xa7\x7c\x29\x75\x06\x4d\x62\x51\xf0\x5d\xbb\x66\xd5\xa1\xce\x64\xdb\x04\xba\xe2\x7f\x04\x1b\x67\x0f\xd8\x94\x0b\x5d\xda\x82\xb2\xca\x77\x62\x4c\x1f\xd9\x93\xb4\x0e\x4e\x15\xe9\x16\x56\x2a\xc5\x1d\xbe\xe7\x56\x99\x41\xa8\x16\x39\xa5\x04\x11\x68\x56\x54\x39\xec\x08\x9b\xcf\xb2\x20\xb7\x3a\x6f\x1b\x23\x69\x01\x21\x2b\x07\x47\x3a\x06\x56\x0e\xc0\x7a\x99\x50\x46\x68\x8e\x0d\x56\x14\x09\xec\x0a\xf7\x90\x62\xf4\x6b\x51\x90\xc8\x4c\xca\xb1\x91\xc5\x13\x5d\xbf\x82\xad\xab\x1f\x3b\x19\x31\x20\xab\x35\x2d\x9d\xad\xf2\xec\xa9\xed\xf5\x60\x10\xa8\xc3\xd0\xee\x04\x0b\x50\xe8\x6a\x33\x55\xfc\x1e\x28\xa7\x10\x4a\xb2\xee\x9e\x08\x0f\xca\x72\x04\x49\x55\x26\x12\xc1\xab\xb5\x98\xb7\xd5\xd0\x64\x0d\x15\x38\x45\xd4\x21\xa5\x65\x1c\xc1\xc1\x66\x3b\x3c\x37\x22\x31\x36\x48\x88\x29\x23\x11\xc1\x69\x3d\x45\x4b\x5c\xe3\x78\x98\x7f\x51\xd9\x10\x1d\x64\x3e\xf2\x02\xca\x1d\x90\xd6\x0e\xd8\x2c\xb7\x88\xb3\x7c\x42\xe9\x8f\x49\x88\x3f\x30\x93\x02\x50\x14\xfb\xa8\x03\x04\x9d\xfc\x88\x35\x83\xaa\xce\x36\x8d\x79\x0c\x84\xc2\x14\xff\x96\x93\x14\xcb\xb4\xa9\x33\x4a\xb2\xf8\xac\xb5\x44\x39\x74\xe9\xa8\x63\xf4\xa7\xbe\x99\x82\xa0\x5d\x3b\x69\xfc\x81\x58\x18\x1c\x41\x97\x0f\xd8\x11\xdd\x3d\x20\x38\x0e\x8d\x63\x0b\x5b\xc7\x4a\xbe\xc4\x86\x53\x13\x1e\x7f\x7e\x11\xd5\xbf\x76\xde\x47\x34\x13\x95\x74\x90\x6e\xa9\x50\xb5\xc4\xd6\x34\xca\xbe\xab\x92\x44\xf1\x8d\x88\x17\x83\x53\x4b\x57\x88\xce\x90\x57\xbe\xd6\x65\x0f\xd4\xfd\xa5\x1d\x13\x98\xac\x5d\xdf\x88\x1d\xdd\x5e\x4e\x6a\x10\x9a\x31\x44\x97\x38\x89\x94\xea\x79\x9d\x1c\xa9\x05\xdd\xbb\xfb\x72\x87\xc1\x2d\x0b\x61\xe9\x8c\x95\xae\x54\xba\x67\x35\x98\xaa\x95\xa6\x96\x25\xbc\x8e\xce\xde\xbd\xf3\x7f\x98\x6c\x36\x74\x0e\xdf\x4d\x4e\x7d\x17\xac\x09\xa4\x0d\x7e\xa6\x23\x83\x26\x9d\xcd\x3c\x17\x58\x99\x65\x4e\xde\xbd\x3b\x9d\x8c\x4c\x81\x22\x5e\x8e\x24\x0c\x0f\x87\x64\xe6\xf9\x67\x60\x8d\x65\xfc\x22\x65\x64\x13\x1f\x7c\x77\xac\x7c\x99\xcf\x27\x27\xa7\xde\xc6\xfb\xc1\x57\xdb\x08\x68\xb5\x91\xe7\x9f\x78\xef\x36\xbe\x3f\x2e\x5b\x29\x5f\xce\x4e\xbe\xdf\xf8\x63\xd7\x96\xad\x26\xa7\x1b\xce\xb9\x50\xf1\xbb\xef\xea\x17\x45\xd1\x5a\x57\x57\xf9\x96\x62\x8c\xa6\x67\xe3\x52\xa5\xb6\xb7\xcb\x49\x47\xb0\x6a\x0b\xea\x94\x28\x85\xae\x9d\x28\xb3\xa3\x8d\x6a\xd7\x2a\xd3\xc6\x20\x9d\x55\x64\x01\x57\xf2\x1a\xdd\x25\x96\x75\x0f\xf1\x5d\x7a\x5f\xa8\x46\x55\x1b\xe7\x76\xbb\xd2\xec\xa7\x94\x28\xad\xad\x93\x7a\x0f\x64\x20\x19\x97\x52\x72\x11\x17\xd5\x36\x6a\x55\x02\x5e\x40\xb7\x4d\x58\x27\xcc\xba\x9d\x0e\xb9\x73\xa1\x94\x58\x3b\xfb\x42\x0c\x74\xc4\x50\x7a\x24\xf8\x0f\xae\x36\x84\x70\x83\xb4\x95\xb4\x43\xe4\x75\x6f\x06\x11\x94\x51\xd6\xa4\xeb\xb5\x32\xd8\x20\x40\x15\x74\xb7\xf7\xa0\x1c\xfc\x41\x94\xea\xe0\xc3\xe0\x36\x92\xb2\xf3\x0b\x66\x33\x7f\x3c\x9f\xfb\xe3\xad\x68\xa2\xf1\x20\x9c\x94\xea\x8e\xdd\xef\x92\xa8\x53\x4d\xd4\x89\x24\xaa\x88\xb3\x99\x37\x99\xcf\xbd\xc9\x4e\x99\xb6\x00\xd6\x42\x6d\x5a\x8f\x96\x77\x3f\x9b\x9d\xed\x12\xb4\x53\x10\xdc\x1b\xb7\xf3\xce\xbf\xe7\x3d\xe9\xbc\x3c\xbd\xe7\x2a\xdf\xd9\xb5\x2d\x22\x28\x07\x08\xc0\x7c\x3e\xef\xfa\x56\xaf\x37\x9d\x72\x64\xcb\xcd\xe4\xee\x66\x17\xd5\xee\xbe\xb1\xc6\x5b\x9c\xae\x12\x79\x7f\x3e\xb2\x83\x2d\x46\xd5\xab\x9d\x9c\x74\x15\xdb\x72\xab\xf9\xce\x2d\x6b\x12\x6a\x5f\x82\x3a\x98\x95\x9b\xda\x77\xee\x76\x43\xd5\x95\x70\x2b\x06\xdd\x3d\xec\xf7\xee\xd4\xd3\xf2\x1c\x6a\xef\x4d\x35\x32\x5a\xe3\xae\x94\x93\xf1\x36\x29\x75\x7b\x39\xcd\x49\x96\xca\x40\x21\xde\x0e\xd0\x75\xea\xc3\x11\xfa\x95\xe2\x3e\x42\xd0\x75\x13\x88\xe7\xf3\x61\x07\xac\x25\xfb\x43\x80\xfb\xa5\xe6\x7d\x71\x83\xae\x67\xf3\xb7\xde\x24\xe8\xfa\x36\x7f\xbd\xc5\xbb\xeb\xee\xfd\x17\x88\xd1\xaf\x83\x6b\x52\x57\x9e\xcd\x24\x2f\x9a\xba\x92\x35\xb6\xb1\x38\x76\x30\xec\xad\x75\x67\xb7\xb0\xa9\xdc\xbe\xcf\xa7\x09\x08\xfb\x31\xd2\x95\xd7\xb9\xe3\x36\x5e\x5f\x29\xb5\xd5\xbf\x26\x74\xec\x83\xdf\xeb\x89\xea\xaf\x5b\xf1\xb5\x91\x43\xaf\xd7\xad\xc1\x63\x58\xca\x47\x5c\x2e\xa7\x2e\xeb\x75\xad\xae\xe0\xd5\xca\x21\x66\xed\xd4\xc4\xda\x6f\x0d\xbe\x63\x2d\xd7\x48\x67\x90\xcc\x60\x49\x62\x64\x4f\x49\xca\x0c\x96\x18\xcb\x84\x32\x44\xa8\x81\xa8\xb2\xfa\xae\xea\xab\x0c\x1e\x1f\x77\x4a\x83\x43\x22\x94\x15\x43\x0b\xca\x45\x91\x13\xa5\xc9\xf3\x8f\xe5\xda\xa3\x35\x9d\x95\x1b\xda\x03\x3d\xc4\x00\x14\xad\x7a\x86\x4e\xaf\xd5\x71\x36\xfd\x72\x79\x34\x38\xe1\x94\xc5\xac\x66\x41\x9c\x82\xf7\xdd\x37\x16\x16\xe7\x51\x07\x39\x8b\xf3\x71\xbd\xc3\x67\xe5\x61\x45\x6c\x75\x78\xf3\x65\xab\xf3\xf0\x20\xe6\x97\x60\x8f\x7e\x65\xdd\x15\xa5\x80\xb7\x34\xf3\x99\x2c\x00\xe8\xdf\x53\x3e\x6c\x72\x01\xe8\x06\x4c\x29\xa2\x5d\x5e\xdd\x8a\x7a\x3c\x84\xb0\x39\x94\xe6\x34\x67\xd2\xb4\xb9\x6d\x73\xc2\xc5\x66\x16\x01\x85\x5c\xa8\x66\x33\x22\x35\x9e\xa8\xe5\x43\x25\x8f\xb5\xb2\x72\x95\x8a\x66\x7c\x69\xb8\x4e\x20\x6a\x56\xa6\xd1\xb6\xe6\x7c\x8d\xb8\x4e\xa0\x89\x4e\x4e\x3d\x30\x9b\x4d\x36\xd1\xc9\xe4\xb4\xa1\xcd\xb7\xd2\x8e\xdd\x92\xd6\x7b\xc7\xb3\x56\x7f\x63\x72\x6a\x01\x93\xb7\x60\xc2\x01\x18\x41\xfc\x3d\xa7\x3d\xab\x69\x39\x4c\x5e\xc1\x84\x1c\x46\x56\x0c\xab\x55\x2e\xd5\x9b\x7d\x52\x2e\x75\x93\x11\x14\x0d\x83\x81\x76\x66\x32\x9f\xf3\x15\xb9\x58\x87\xdb\x66\x72\x62\x7a\x72\x35\xee\x01\xab\x5c\x90\x17\xad\x4d\x09\xbd\xf1\x10\x1a\x92\xb4\xbb\x47\x50\x0d\x68\xd7\x32\xd9\x70\xd8\x5a\x2c\xbe\x76\xcb\x38\xdb\x70\x5a\x0e\xbc\x13\x74\x11\x63\xba\x4f\x1e\xdd\xe7\xb3\x4d\xe0\x27\x94\x2d\x88\x26\xc8\x96\x15\x99\xff\x86\xea\xe4\x3e\xe1\x53\xc8\xaa\x54\x6c\x70\x55\xd6\xde\x37\x76\x56\xda\x94\xc7\xeb\xea\x48\x9a\xe6\xc3\x79\xe3\x6b\x4a\x0e\xcb\x84\x66\xcc\x20\x50\x6e\xa8\x38\x51\x8a\xf1\xbf\xb0\x59\x1d\x6f\x07\x41\x05\x64\x50\x93\xda\x19\x3c\x3e\xb6\xd7\x91\xf8\x34\x4d\x20\x11\xa7\xdf\x6d\x79\xab\x60\x8a\xa0\x5f\xc0\x75\x21\xaa\xf4\x19\x33\xa2\xf2\x94\x64\x59\x19\x31\x3d\xd7\x1f\x83\x92\x5d\x0e\x13\x08\x89\x73\x8d\x5e\xde\x53\x87\x25\xd2\x91\x4c\x30\xfd\xcb\xcd\xe7\x2b\x47\x06\x4e\x12\x7d\x33\x69\xd5\x3e\x84\x51\xab\x4e\x67\xe6\xd5\x97\xa7\xee\x97\x0c\x6c\xb9\x80\x60\x46\x40\x7b\x55\xc0\x8c\xec\xb0\xf7\xe5\x67\xc4\x9e\xcc\xc8\x7e\xea\x7d\xb8\xfd\xb6\xe2\x14\x49\xef\x83\x3c\xf5\x6f\x46\x36\xaa\xc4\x7b\x80\x43\xa7\xfb\xb9\x28\xfa\x33\xd8\xed\x2f\xca\x41\x6b\x33\xaa\x4e\x52\x37\x77\x73\xfa\x2d\xcd\xc8\x7e\xa8\xd8\xaf\xa0\xee\xcc\x34\x67\x10\x95\x9b\x09\xe6\x0a\x04\xbf\x9c\xfd\x2d\x49\x7f\xc5\xa9\xef\x64\x98\x86\x66\xd4\x3e\xb5\x68\x02\x50\xd4\x46\x90\x49\x7b\xe2\x66\x75\x7c\xcc\xed\x47\xee\x3b\x25\xe2\xe6\x41\xe2\xfc\x84\x23\x9c\xa6\x38\xbc\xc6\xec\x2b\x8a\xeb\x3b\x08\xd5\xeb\xa9\x67\x0b\x8f\x9c\xfa\xbc\xf1\x45\x1e\x47\x24\x7e\xc6\x94\xc9\xeb\x1a\x75\x6b\xc4\x10\x6f\x49\xc3\xcf\xd1\x0d\x4b\x31\x7a\x9e\xfa\x25\xdd\x29\xa7\x6b\xae\x2e\xec\x7b\x5b\x41\xa5\x79\x93\x0b\x0b\x2d\xc0\xfa\xce\x82\x4e\x30\xc5\xfd\x4c\x2c\xaf\x2d\xa8\xc4\xaf\xb9\xb9\x30\xd0\x1b\x79\x12\xe1\x0d\xaf\x2e\x54\xe7\xe2\x9a\xbb\x0b\xf5\xf6\x45\x1d\x67\xc4\xf1\xe7\x8e\x48\xc2\x96\xb5\xba\xe8\x5e\x29\xf0\xba\xa4\x28\x0c\x6f\x3a\xdd\xe8\xdf\x9c\x99\x8c\xc5\xb5\x1f\xd6\x15\xa3\x0b\x86\x69\x38\x20\xc5\xf6\x7b\x02\x8d\x19\xef\x6f\x5f\x15\xc5\x1b\x59\x57\x0d\x57\xdb\x56\x5f\x24\xad\x65\x55\xcd\x5e\x67\x57\x9a\x5e\xfc\x67\xad\xaa\x16\x48\x58\x8d\x46\x07\x7a\x8b\xaa\xc9\x7e\xbf\x3d\xd5\x50\x98\x86\x5a\xfe\xfb\xda\xd2\x35\xce\x56\x09\xcd\xf6\xbe\x61\xd5\xa7\x7c\x63\xdb\xaa\x61\x7b\x36\xd6\x17\x75\xab\xad\x55\xcd\x7f\x9f\xcd\x69\x7a\x99\x8a\x59\xe4\x8f\xbe\x90\xd5\x9d\xb5\xca\xab\x59\x7b\x48\xf6\x6f\xb8\xa5\xa5\x91\xa2\xe5\x0a\x9a\xa1\xd2\xdd\xdb\xd2\xc0\xa0\xb0\xec\x6f\xf7\x0a\x57\xff\xa2\x22\xf7\x0e\xad\x96\x76\x20\xef\x7b\xaf\x4b\x83\xa1\x78\x9b\xb6\x8b\x3b\xbd\x0e\x31\x74\x80\xa7\x21\x86\xde\xca\xbb\x38\x54\xe3\x51\x2d\x31\xf4\x5e\x84\x18\x7a\xa5\xe7\xb4\xa5\xee\xec\xce\xfd\x71\xb7\xab\x4b\xce\xd2\x0e\xdb\x1d\x1c\x08\xc7\xbc\x39\x0a\xc3\x8f\x1d\x09\xb7\xde\xe8\x2d\xc9\xb8\x21\x74\x79\xec\x1a\x7c\x25\x73\xdc\xdf\x06\x14\xa2\x37\x32\x05\x15\xb1\xb6\x08\xad\x6c\x5a\xc3\x50\x5a\xbe\xce\x3e\x54\xfe\x62\x50\xf4\xbc\xbb\x63\xe6\xf6\x88\x31\x0d\x87\x48\x77\x0d\x45\x93\xe3\xef\x3f\x12\x0d\xcd\x1b\x0d\x84\x02\x58\x8f\x83\x4e\x30\xed\x30\x34\x0d\x5f\x37\x0a\xda\xde\xfc\x67\xb3\x2a\xad\x48\xe2\xd7\x11\x5e\x31\xd7\x4e\x0e\x9b\x6b\x3b\x8b\xbe\x7a\xb2\x1d\x96\xe9\xd0\x90\xa6\xfb\xd1\x84\x9d\xb3\xac\xc2\x5e\xf8\x82\xd6\x3c\xba\x9e\x72\xda\xd5\xe6\xef\x4f\x37\xdb\x60\xbf\xf4\xc7\x44\x37\x3f\x7b\x62\x7e\xd6\x69\x76\x08\x73\x57\x10\xf6\xeb\x20\xac\x10\x63\x1a\x0e\xa8\x65\x7b\x14\x90\xa5\x08\x04\xd7\x45\x53\x4c\x8a\x38\x99\xfc\x80\xe1\x1e\x67\x52\xea\x32\xd5\x2e\x6f\x36\x71\xd5\x96\x42\x52\xfa\x99\x09\x9a\xbb\xc4\x26\x08\x62\xcc\x8c\x4c\xfc\x1b\x05\xd9\x0b\x61\xcb\x27\x93\x34\xe6\x6f\x02\xb0\x5e\xa2\x0c\x1b\x7d\x8d\x8a\xc2\xc4\xda\x5c\x87\xfc\xff\xac\x80\xe8\x8e\xde\x57\xdc\x18\x6c\xa6\xf3\xa0\x84\x33\x19\x08\x22\x58\x9f\x48\x2f\xcf\xa7\x2a\xe0\x39\x51\xe1\x65\x89\x63\x6d\xae\xb1\xf8\x63\x80\x81\x68\xd5\xe6\xd0\x54\x3a\x99\x23\x68\x4d\xa0\x61\xd6\x62\xa5\xd4\x56\xcc\x35\xa6\x61\xc3\x4d\xd0\x05\xe5\x0d\x92\x69\xaf\x94\xfa\x85\xfe\x4a\x93\x17\x6a\x3c\xcb\x52\x92\x91\xe2\x25\xc1\x5f\x71\x68\x44\x69\xf2\x6c\xa4\x39\x65\xe4\x19\x1f\x8b\x53\x67\x19\x84\x30\xa7\x21\x8e\x08\xc5\x61\x15\x58\x8b\xcc\x8c\x94\x7a\x52\x2e\xaf\x93\xd8\x04\xac\xd1\x1d\xbb\x87\x52\xb7\xd8\x96\x2a\x48\x6d\x2e\x1a\x29\x9a\xf6\x61\x53\x66\x36\x01\x9c\xf7\x4a\xbd\x0c\x28\x8d\x9f\x54\x23\xe3\x36\x59\x19\xd1\x2e\x83\xc3\xa0\x31\xa2\x43\x16\x4c\x26\x01\x95\x45\x51\x65\xf5\xa2\x98\x54\x27\x89\x2e\x87\x5c\x37\xc2\x15\x80\xb2\xe7\xf5\x73\x9a\x3c\x93\x8c\xaf\x3e\x84\x8f\xe9\x86\x1d\x14\x7a\x46\x75\x05\xae\xc3\xab\x7a\xdf\x62\x57\xab\x4c\xef\x3f\xca\xd1\x93\x52\x22\x53\xfc\xf2\x03\x9c\xaf\x6b\xd5\x13\xae\xfa\x10\xc7\x98\x61\x03\xdd\xe1\x7b\x75\x4a\xe6\x83\x94\x9b\xd8\x26\xce\x03\xa1\x7c\x40\xab\xbf\xd2\xfa\xaf\x90\x1b\x1a\xfe\xe7\x0a\x2f\x19\x0e\x0d\xc5\x5a\x8d\x28\x49\x8d\x95\xe0\x4a\x22\x82\x43\x23\x2c\x3b\x70\xcc\xfb\x0e\x8a\xca\x6e\x77\x69\xec\xf8\x02\x91\x18\x87\x06\x4b\x8c\x10\x2f\x93\x90\x1b\xb2\x1c\xc3\xd2\x90\xf1\x6f\x39\xce\xd8\x31\x00\x45\xd1\x94\x47\x53\xbc\xfc\x6a\x46\x72\x63\xee\x01\x3e\x88\xa2\xe7\x83\x53\xfd\xaa\xca\xbe\x19\x4e\x43\xf1\x26\xf9\x8d\x02\x57\x1b\x66\x5f\x24\x4d\x6e\xd3\x10\xbe\x26\xb3\xd1\xf6\x22\x4f\x0f\x5e\x14\xbf\x66\x05\xa2\xe5\xcd\x03\x4f\x92\x1f\xfa\xe3\x46\xdb\xd2\x97\xe6\x24\xbf\xf2\xeb\x46\x0a\x6f\x91\x0b\x68\x54\xad\x5b\x81\x2b\x64\x28\x0c\xbf\xa4\x3b\x17\xc7\xd5\x3a\xa8\x4d\x78\xdb\xed\x64\x3f\x1f\x38\xf5\xeb\x85\xb5\x42\x8b\x69\xa8\x95\x74\xfb\xd4\x5d\xfb\xf3\xaa\x09\xa5\xa4\xb3\xaf\xf3\xce\xf3\xd5\x29\xb7\xb5\x19\xc3\xc0\xa0\xbe\x78\xac\xec\x29\xc5\x24\x76\xbd\xe7\x93\x41\xbd\xf8\x9c\x70\x60\xe3\x84\x0c\x6d\x9c\x90\x6a\xe3\xa4\xb6\x78\xcd\xbe\x09\xb1\xb3\x8a\x77\xa2\xdf\x37\x21\x20\x20\xd5\xbe\x49\x52\xeb\xe9\xc9\xe4\x53\x51\x67\x13\x85\xf4\x36\x51\xe4\x5e\x49\x0c\xd7\x71\xf2\x38\xcd\x6c\xb9\xbf\x36\x25\xb6\x38\x9f\x32\xa5\x36\xb7\xba\xe9\xaa\x08\xf0\x3f\x57\x49\xca\x8c\x32\x92\x19\x71\xf0\x3f\xff\x1f\x00\x00\xff\xff\xc1\x19\x6f\xd0\x7c\x4e\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/std.js"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
