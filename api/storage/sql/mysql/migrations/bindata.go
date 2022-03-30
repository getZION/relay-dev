package migrations

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var __1_initial_schema_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\x51\x6f\xab\x36\x14\x7e\x8e\x7f\xc5\x79\x03\xa4\x28\x4a\xb7\xf6\x6a\x53\xd5\x07\x42\x9c\x8c\x2e\x85\xce\x81\x49\xbd\x2f\x0e\x05\xdf\xcc\x5b\x02\x15\x98\xde\x35\xbf\x7e\xb2\x69\x02\x98\xa4\xb4\x5b\x34\x35\x2f\x31\x3e\x9f\x7d\x3e\x8e\xbf\xf3\x61\x87\x60\x3b\xc0\x10\xd8\x93\x05\x86\x55\x9c\x6d\xb7\x65\xca\x05\x67\xc5\x0a\x4c\x04\xb0\xe2\xc9\x0a\x1e\xf9\x9a\xa7\x02\x3c\x3f\x00\x2f\x5c\x2c\xc0\x0e\x03\x9f\xba\x9e\x43\xf0\x1d\xf6\x82\xa1\x84\xed\x24\xee\x39\xca\xe3\x3f\xa2\xdc\xfc\xe1\xea\xca\x3a\xa0\x55\x38\x8d\xb6\xac\x8e\x5f\x5c\x8d\xb5\x78\xf6\x3d\x65\x39\x4d\xde\xde\xa4\x02\x95\x05\xcb\xdb\xdb\x75\x91\x09\x2b\xe2\x9c\x3f\x09\x9e\xa5\x4d\x98\x9e\x55\xa2\xb2\xef\x34\xda\x66\x65\x2a\x3a\xaf\xa9\x20\x7c\xbb\xd6\x12\x4d\xf1\xcc\x0e\x17\x0d\xc8\x26\x2a\x04\x8d\x62\xc1\x9f\xd9\x61\x8f\x0e\xe8\x29\xe7\x31\xa3\x4f\x2c\xa7\x5b\x56\x14\xd1\x9a\x1d\x4f\x57\xc1\x44\x46\xff\xcc\x78\x7a\x02\x52\x3e\x6e\x78\xbc\x02\xc1\xd3\x17\x9e\x0a\xf3\xa2\xa6\x64\x5c\x18\x0a\x12\xe7\x2c\x12\xac\x7b\x70\x2a\x58\x3e\x25\xad\x60\x87\x6a\xc2\x36\x4c\x01\x8e\x65\x18\xab\x0c\xf7\xc4\xbd\xb3\xc9\x03\xfc\x8a\x1f\xc0\x94\x12\xb1\xe4\x6c\xe8\xb9\xbf\x85\x58\x4d\x2a\xd9\x1c\x8f\x54\x67\x67\x56\xff\x9d\xe8\xae\x5a\x28\xff\x2c\x64\x01\xf6\xe6\xae\x87\x6f\xdc\x34\xcd\xa6\x13\x4d\x77\x37\xe3\x03\x2d\xe7\x17\x9b\x2c\x71\x70\x53\x8a\x6f\x3f\x6d\x1f\x2f\xc1\xf1\x17\x0b\x3b\xc0\xfb\x67\x3a\xfe\x79\x3c\xa6\x11\xa7\x31\x47\x48\x17\x7c\xfa\xcc\xf2\x22\x92\x52\xf9\xcf\x92\xef\x94\x72\xdf\x4e\x2f\xb4\xa7\x3f\x04\xfb\x5b\xf4\xea\x8c\xa7\x7f\xf5\x61\xde\x21\xd7\x67\x9e\xb0\xac\x0f\xf4\x96\xc8\x2a\x09\xbc\x42\xa8\x52\x6c\x47\x4c\x7b\xd0\xa7\x56\x62\x53\x6b\x32\xa6\x26\xb5\x23\x33\xb5\x09\x05\x74\x7c\x6f\x19\x10\xdb\xf5\x02\x4d\x40\xb4\x05\xa6\x0d\x37\xa5\x5f\x79\x42\xbf\x65\x39\xe3\xeb\x74\x05\x33\x9f\x60\x77\xee\xbd\xb2\xd6\x32\x00\xc1\x33\x4c\xb0\xe7\xe0\xa5\x6e\xc8\x15\x57\xf0\x3d\x98\xe2\x05\x0e\x30\x10\xbc\x0c\x88\xeb\x04\x72\x2a\xbc\x9f\xda\x8d\xa9\xff\xa3\x7b\xa4\x17\x7f\xb0\x6b\x7a\x3c\xfe\x3d\xee\xce\xb6\x11\xdf\xf4\x29\xb8\x6f\x93\x47\xde\xdb\x04\xef\x68\xa6\x83\x5f\xeb\xae\xde\xb5\x02\xad\x13\x3a\x80\xbe\x6e\xe8\x57\x7a\x52\x09\x36\xf9\x58\x17\xd4\x15\x37\xeb\x71\xaf\xf7\x5e\xfe\x0b\xf5\x5c\xeb\xf2\x11\xd1\x7a\xaf\x1e\x11\x35\x6a\xfd\xe5\xd2\x52\xdb\xda\x4e\x80\x09\x2c\x71\x00\x72\xb3\xfd\xce\xea\x81\xae\x59\xca\xf2\x68\x43\x63\xde\x3a\xd7\x76\x95\xe4\xae\xdd\x57\x39\x03\x73\xd9\x96\x2c\x15\xe7\xbd\x24\x35\xad\xa4\xef\x83\x21\x4f\xaa\xef\xbe\x74\xae\x8f\xca\xa7\x33\xf1\x57\xa3\xd6\xab\x65\x76\xe7\x6a\x78\x5d\x30\xb3\x1e\x1f\x71\xf3\xea\x5c\xa9\xbe\x13\x6d\xfb\xfc\x9b\x6e\xae\x73\xd0\x0c\xbd\x7d\xe1\xf8\x88\xa5\x9f\x64\xbb\x7f\x21\x35\x28\xe8\xf4\x34\xbb\xfa\xd5\x5b\xac\xf6\x46\x5e\x99\xc7\x79\x3e\x30\x5f\xce\xd5\x68\xd5\xb7\xb1\x36\x8b\xee\xc5\x0a\xc0\xfb\xdd\x26\x32\x4d\x25\x61\xa8\xdb\xb0\x52\xd1\xb0\x5a\xa6\x5c\xa6\xf9\x3b\x2c\x93\x96\x73\x6c\xd9\xa0\x55\xbe\x56\xda\x56\x05\xa1\xf1\xa1\x36\x65\xb0\xca\xd8\x5a\x2d\xa2\xb5\xd5\xc8\xdd\x5c\x2d\x5f\x4e\xc5\x91\x55\x97\xe0\x9e\xf8\x0e\x9e\x86\x04\xc3\x6a\xcd\x04\x6d\x5f\x05\x2c\x34\xc1\x73\xd7\x43\x68\x89\x17\xd8\x09\x00\x0d\x6e\x97\xbe\x47\x6d\x42\xec\x07\x7b\x3e\x37\xd5\x93\x3f\xb9\xc5\x4e\x60\x1a\x3c\x31\x86\x10\x8f\x78\x32\x04\x43\xfa\xbb\x7a\x92\x83\x21\x18\xbb\xd7\xe0\x4e\x45\x25\x11\x63\x08\x62\x24\x07\x43\x30\x94\x30\x8c\x21\x94\x23\x35\xb2\x2c\x88\x0a\x30\x08\x2b\xca\x8d\x30\xd0\x8c\xf8\x77\x80\x06\x39\xdb\x44\x2f\x3f\x8e\x1a\x04\x21\x46\x0b\x3c\x0b\xe0\xd6\x77\x3d\x30\xd1\xe0\x40\x72\x50\x97\x50\xe6\x43\x03\x8d\x76\xac\x32\x5b\x60\x2f\x55\x4d\xd0\x40\xa5\xd0\x12\x54\x62\x80\x58\xa0\xc1\x9c\xf8\xe1\x3d\x4c\x1e\x20\x16\xa3\xe6\xd6\xc8\x02\x21\x35\xdb\x9e\x85\x1b\x55\x85\x53\xdc\x9c\x3d\xf4\xeb\x71\x6e\xe5\x28\x2c\x58\x3e\x95\x47\x6f\x2f\x41\x15\xe4\x14\x41\x15\x84\xb8\x6c\x32\x2c\x47\xcd\x04\xc8\x82\x52\x32\x6c\xcf\x2a\x86\x3b\x9e\x5c\x23\x84\xbd\xe9\x35\xfa\x27\x00\x00\xff\xff\xc7\x4e\xcd\x9d\x9f\x0f\x00\x00")

func _1_initial_schema_down_sql() ([]byte, error) {
	return bindata_read(
		__1_initial_schema_down_sql,
		"1_initial_schema.down.sql",
	)
}

var __1_initial_schema_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x4d\x73\x9b\x3a\x14\x5d\x3f\x7e\x85\x76\xc1\x33\x59\xc4\xf9\x9a\xcc\x64\xb2\x20\xb6\x9c\xc7\x3c\x07\xe7\x61\xfc\xe6\xa5\x1b\x45\x01\x85\xaa\xc5\xe0\x01\xd9\x69\xf2\xeb\x3b\x12\x36\x91\x04\xc2\xb8\x75\xda\xd4\x2b\xb8\x3a\xba\x5c\xa4\x73\x8e\x2e\x1e\xf8\xd0\x09\x20\x18\x3a\x81\x73\xed\x4c\x21\x70\x47\xc0\x9b\x04\x00\xfe\xef\x4e\x83\x29\x78\xc8\x49\x82\x5f\x4e\x1e\x2e\xad\x65\x41\x40\x79\x73\x69\x59\x43\x7f\x72\x07\x02\xe7\x7a\x2c\xf0\x6b\xec\xb2\x20\x79\x71\x69\xad\x13\x96\xa3\x22\x06\x6c\x0b\x00\x1a\x01\xed\xf7\x48\x63\x9a\xb2\xcd\x1d\x7f\xa8\x37\x1b\x8f\x81\x33\x0b\x26\xc8\xf5\x06\x3e\xbc\x85\x5e\x70\x68\x01\x10\xd5\xe6\xae\x70\x1e\x7e\xc6\xb9\x7d\x7c\x76\xde\x7b\x9b\xca\xb1\xfc\x81\x29\x9e\x93\x3a\xb6\xcf\xa1\x0a\x96\xcc\x31\x4d\x1a\xf3\x9e\x9f\x72\xec\x10\x8e\x9c\xd9\xf8\x0d\xaf\xe6\x95\xf1\x27\xc7\x7a\xee\x47\x9a\x19\x6a\x3e\xeb\x73\xb0\x9e\x9b\xce\xe3\xd6\x77\xd4\xf1\x8b\x9c\x86\x04\xb1\x0c\xcd\x49\x51\xe0\x98\xe8\xcb\xa9\xe3\xc3\x9c\x60\x46\xa2\xad\xcb\x2f\xd6\x70\x11\xb5\x62\xf5\xdc\x77\xbe\x7b\xeb\xf8\xf7\xe0\x1f\x78\x0f\x6c\x1a\xf5\x78\x6c\xe6\xb9\xff\xce\xa0\x08\xf1\xed\xb3\xa3\x7a\x9c\x87\xeb\xd1\x6a\x03\xed\xcd\x55\xcf\xea\x01\xe8\xdd\xb8\x1e\xbc\x72\xd3\x34\x1b\x5e\x6b\x14\xb9\x3a\xad\x2a\x1a\xfc\xed\xf8\x53\x18\x5c\x2d\xd9\xd3\x05\x18\x4c\xc6\x63\x27\x80\xe2\x06\xc5\x24\x25\x39\x4e\x50\x48\x4d\xf4\x0d\xb3\xf9\x7c\x99\x52\x46\x89\x4e\x62\x69\xe4\x67\xa9\xfc\xba\x03\x95\xcd\x74\xeb\x1f\x5f\x68\xd8\xec\x39\x25\x39\x92\x85\x62\xce\x5b\x62\x65\xa1\x98\x25\x12\x91\x22\xcc\xe9\x82\xd1\x2c\x6d\xa4\xb1\x22\xa7\x22\xcc\xb3\x67\x84\xe7\xd9\x72\xbd\x16\x66\x8a\x99\xe9\xde\x2c\x8f\x04\x17\x0c\xe1\x90\xd1\x15\xe9\x44\xc9\x52\x1e\x0b\x92\x57\xfa\x30\xd7\x52\x49\xe9\x4b\x46\xd3\x2d\x75\x2f\x96\x8f\x09\x0d\x95\xad\xcf\xb2\x84\xe0\x54\xab\xa3\xff\xde\x9a\x8b\x48\x42\x34\x7c\x63\x21\x47\x1d\x04\xda\x2c\xc4\x52\x84\x42\x80\xda\x08\x27\xb1\xfd\xba\x9e\x31\x98\x78\xd3\xc0\x77\x5c\x2f\x90\x75\x22\xe8\xc5\xf9\x28\x2e\x0a\x34\xa4\x11\x7a\xca\x72\x42\xe3\x14\x8c\x26\x3e\x74\x6f\xbc\xb2\x9a\x8a\xb8\x3d\xe0\xc3\x11\xf4\xa1\x37\x80\xd3\xcd\xb1\x21\xc2\x13\x0f\x0c\xe1\x18\x06\x10\xf8\x70\x1a\xf8\xee\x20\xe0\xa1\xd9\xdd\xd0\x91\x42\xdb\x2c\xe2\x68\x4f\x16\x91\xae\x48\x5e\x60\x2e\x87\xba\x49\x48\x63\x0d\x36\xf1\x83\x16\xd1\x7e\x0a\x6c\x56\xfc\x05\xf1\x59\x66\xc9\x33\xf2\x8d\x75\xcd\x99\xd0\xf4\x6b\x57\xac\x2a\xe3\x76\xec\x8a\x46\x24\xeb\x88\x55\x65\xd6\x22\xb1\x12\x88\x84\x84\x4d\x92\x39\x38\x3a\xa8\xab\xb1\xab\x12\x77\x51\xe1\xbe\x15\xa8\xe8\x8c\x07\xd4\xed\xb6\x95\xdb\xba\x18\x25\x3e\x22\x05\x8a\x64\xa1\x7e\x32\x29\x53\xcd\x2e\xab\x53\x39\x0f\x5f\x3f\x98\x46\x19\x8e\x75\x69\xf2\x90\x50\x24\xc3\x31\xf8\xcf\xf1\x79\x7a\xd1\xe2\x55\x2a\x94\xb6\xa7\x56\xe5\xbe\x7a\x0b\x92\xb2\xa6\xc6\x82\x87\x7f\x65\x57\x21\xf3\x42\xd0\x68\x4b\x33\xad\x34\x15\x9d\x2d\x66\x57\x9b\xe9\x60\x75\x7f\xde\x61\x5a\x4a\x56\x5b\x6e\x5b\x8f\x54\xd0\x6a\xb5\xed\xcd\x55\xd3\x01\xcb\xf9\x82\xf4\x1c\x48\x55\x7b\x8b\xa6\xb5\x67\xab\xb2\x56\x4e\xb0\xee\xc2\x36\x55\xd9\xb5\x07\xa8\x5e\xf7\x5d\x5b\x80\xf3\xbd\x7e\x25\xbc\xa0\x06\xa3\x51\x07\x85\xaa\xf5\x33\x1a\x54\x06\x54\xd2\x5c\xe1\x6d\xe9\x50\xf2\x4f\x76\x2b\x05\xfc\x17\xe8\xea\xd6\xb2\x5d\x57\xa7\x89\x36\x9b\xe1\xb8\x27\x3d\x55\x9e\xcd\xdf\x44\x8c\x97\xd3\xd6\x4c\xaf\x1f\x47\xfc\x95\xb5\x32\x0e\xf9\xe4\xad\x5f\x6f\xfb\xde\x97\xa6\xbf\x21\xb4\xd1\x46\xbf\x55\x2c\xa2\xc5\x6c\xf5\xfd\x54\x6c\xab\xcd\x3d\x8d\x40\xfe\xdd\x41\x22\xc4\x3d\xcb\x5c\x89\xf0\x4c\xf2\xc4\xde\x60\x1a\xb0\xcb\x57\x79\x8b\xcd\xec\xde\x5f\x28\x4b\xfa\x21\x3a\x8c\x2d\x25\xfe\x66\x3f\xda\x85\xe6\xd6\xf7\x00\x00\x00\xff\xff\x85\x94\x74\xa9\x97\x13\x00\x00")

func _1_initial_schema_up_sql() ([]byte, error) {
	return bindata_read(
		__1_initial_schema_up_sql,
		"1_initial_schema.up.sql",
	)
}

var __2_stored_procedure_get_communities_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _2_stored_procedure_get_communities_down_sql() ([]byte, error) {
	return bindata_read(
		__2_stored_procedure_get_communities_down_sql,
		"2_stored_procedure_get_communities.down.sql",
	)
}

var __2_stored_procedure_get_communities_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x93\x3f\x6f\xdb\x30\x10\xc5\x67\xde\xa7\xb8\x4d\x36\x60\x68\xe9\x18\x74\x90\x1d\xc6\x88\x91\x4a\x81\xec\x0c\x99\x08\x95\xba\x0a\x2c\xf4\x0f\x22\xd9\xc0\xf9\xf4\x05\x49\xd9\xa1\xff\xb4\x9e\xee\xde\xfb\x59\x3a\xbf\x3b\x5b\x4d\x38\x51\x5b\x1d\xbf\x3d\x00\x6c\x4a\x9e\x1d\x38\xbe\x96\xc5\x86\x3f\xbe\x95\x1c\x1b\x32\x42\x0e\x5d\x67\x7b\x65\x14\xe9\xc5\x12\xd6\x7c\xfb\x9c\x03\xec\xf9\x0b\xdf\x1c\x10\xd8\x6e\x5f\xe4\x22\x2b\xcb\xec\x3d\xdb\x6e\x17\xbe\x2b\xd6\x3b\xbe\x39\x2c\x80\xb1\x44\xd5\xc9\x0a\x65\xaa\xea\x15\x02\xce\x9f\xa4\xaf\x3a\xf2\xb2\x2b\x62\xe3\x73\xc6\x3f\x2f\xf9\xe1\xa3\xa7\xa9\x9e\x3d\xdf\x88\xfa\x0e\x61\x35\x4d\xe7\x47\x07\xec\x24\xc5\x6c\x4d\x5a\x4e\x6a\x34\x6a\xe8\x3d\x19\xf5\x31\xc6\xb5\x9c\x86\x8f\xac\x1b\x6c\x6f\x3c\x47\x5e\x10\x95\x57\x62\x52\x75\x4d\xf8\x91\x5d\x13\xcb\x6d\xa5\x4d\x25\x8d\xfa\x13\x06\x72\xad\x08\x7d\x4c\x8d\x93\x92\x34\xd2\xd4\x91\xd6\x55\x13\x50\xaf\x89\x91\x26\x31\xab\x37\x5f\x30\xc3\xef\x41\xf5\x11\x6c\x06\xe1\x94\x0b\xd0\xfe\x6c\x95\x0c\x8c\x2f\x63\x53\x4e\x54\x19\x0a\x89\xce\xf5\x0a\xdd\xc2\xec\x58\x9f\x8d\xb9\xbe\x0c\xaf\xa5\x93\x3d\xd7\xb1\x6d\xaa\x46\x27\x2b\x34\xa9\x2b\x62\xc3\x2d\xc2\x39\x36\xf5\xd5\x12\xd8\x12\x2b\x8d\x49\x49\xda\xb6\x26\x81\xa7\xb2\xf8\x81\xc0\xc2\x21\xa6\xd1\xc5\xa1\x84\x17\xfe\x74\xc0\x5d\xf1\x9c\xe3\x02\xd8\xf9\xee\xd8\x09\x3a\x8a\x70\x2e\xec\xea\x14\xa5\x9f\x62\x89\xd9\x1e\xdd\x34\xc0\xfc\x3b\xae\xde\x70\x14\xce\x43\x69\x80\x6d\xcb\xe2\xed\x15\xd7\xef\x28\x4d\x7a\xf1\x6c\x58\xa2\xc1\x22\xc7\x2b\x19\xbf\x87\x4b\xfd\xc7\x7c\x37\xe3\xfd\xff\x8f\xe2\x62\x11\xf3\x8d\x87\x90\xfc\x8d\x7f\x25\xe8\xd6\x4b\xb5\x70\x1b\x09\x4c\x24\x44\x58\x4b\xbf\x4c\x04\x9d\x5b\x9f\x78\xb6\x47\x1f\xbf\xc7\xef\xc7\xe1\x7d\x94\xd6\x23\x5f\x91\xd8\x9b\x48\xac\x8b\xc4\xde\x8f\xe4\x01\x80\xe7\x8f\xf0\x37\x00\x00\xff\xff\x77\x29\x0d\x55\x5b\x04\x00\x00")

func _2_stored_procedure_get_communities_up_sql() ([]byte, error) {
	return bindata_read(
		__2_stored_procedure_get_communities_up_sql,
		"2_stored_procedure_get_communities.up.sql",
	)
}

var __3_stored_procedure_get_users_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _3_stored_procedure_get_users_down_sql() ([]byte, error) {
	return bindata_read(
		__3_stored_procedure_get_users_down_sql,
		"3_stored_procedure_get_users.down.sql",
	)
}

var __3_stored_procedure_get_users_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xbf\x6e\x83\x30\x10\x87\x67\xee\x29\x6e\x03\x24\xd4\xa5\x63\x26\xfe\x5c\x11\x51\x03\x91\x03\x43\x27\xe4\x86\x13\xb2\x04\x25\xb2\xf1\xd0\xb7\xaf\x62\xa7\xad\x5b\x4f\xbf\xfb\x3e\x9f\xe5\x3b\x6b\x18\x35\x2f\xf2\xf3\xf9\x00\x50\x0a\xca\x7b\xc2\xb3\xe8\x4a\xaa\x06\x41\x38\xf3\x3e\x5a\xc3\xda\x24\x29\x14\x54\x37\x2d\xc0\x85\x5e\xa9\xec\x11\xa2\xe3\xa5\x6b\xc7\x5c\x88\xfc\x2d\xaf\xeb\xc4\x55\x5d\x71\xa4\xb2\x4f\x20\x8a\xe2\x66\x8a\x33\xb4\x4f\x6a\xca\x10\xf0\x71\xe2\x4a\x79\x3a\xfd\xc5\x83\x61\xfd\x21\x57\x76\xce\x3e\x8a\xec\xd7\xd3\x2a\xd5\xe2\x24\xdf\x53\x60\xda\xef\xae\x7f\x1d\x85\xda\x1c\x7e\x57\x5b\x40\x9b\x75\xf6\x9f\x5a\xe7\x80\x9e\xb5\xba\x72\xbf\x9d\xd8\x18\x39\xfb\xd7\x6e\x77\x34\xee\xdb\xb8\x7a\x18\xdc\x2e\x35\xcb\x9d\xfd\x18\x57\x9f\x03\x3b\xdc\xa6\x1f\x6b\x7d\x4e\x21\x4a\x51\x1a\x8c\x05\x1b\xbb\xec\x31\xbc\x88\xee\x84\x10\xf9\xa5\xbb\x71\x0d\xda\x03\x00\xb5\xd5\x57\x00\x00\x00\xff\xff\x44\x42\xdc\xcd\x8d\x01\x00\x00")

func _3_stored_procedure_get_users_up_sql() ([]byte, error) {
	return bindata_read(
		__3_stored_procedure_get_users_up_sql,
		"3_stored_procedure_get_users.up.sql",
	)
}

var __4_stored_procedure_get_conversations_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _4_stored_procedure_get_conversations_down_sql() ([]byte, error) {
	return bindata_read(
		__4_stored_procedure_get_conversations_down_sql,
		"4_stored_procedure_get_conversations.down.sql",
	)
}

var __4_stored_procedure_get_conversations_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x93\x41\x6f\x9c\x30\x14\x84\xcf\x7e\xbf\xe2\xdd\x58\xa4\xd5\x5e\x7a\x8c\x7a\x60\x89\x83\xb2\x4a\x21\xf2\x92\x43\x4e\x88\xda\xd6\xca\x2a\xe0\x15\x98\xa8\xc9\xaf\xaf\xb0\x31\x35\x50\x75\x4f\xf3\x66\xe6\xb0\x9e\x4f\x8c\x83\xc4\x5e\x36\xf5\xe7\xb7\x07\x80\x94\xd1\xa4\xa4\xf8\xca\x8a\x94\x3e\xbe\x31\x8a\x37\x69\x2a\xae\xbb\x0f\xd9\x0f\xb5\x51\xba\x1b\x0e\x31\x9c\x69\xf6\x9c\x03\x5c\xe9\x0b\x4d\x4b\x04\x72\xb9\x16\x79\x95\x30\x96\xbc\x27\x59\x76\xb0\x57\x71\xbe\xd0\xb4\x3c\x00\x21\x91\x12\xd1\x11\xf9\x49\x89\x23\x02\xce\xbf\xe8\x6b\x76\xbf\x94\x38\xfe\x75\xb9\x6e\xdb\xb1\x53\xe6\xd3\xc7\x8b\x51\xad\x8b\x46\xfe\x36\xb6\x30\x89\xc0\x6f\x54\xf7\xcb\xfa\x93\x08\x7c\xd5\xde\xdc\x9f\x68\x6f\x81\xfb\xa1\x84\xd4\xd6\xb7\x2a\x48\xee\xe3\xcf\x46\x71\x1b\x39\xb9\xcb\xee\xbd\xe2\x32\x28\x54\xd6\x08\xdf\xd2\xcb\xda\xc8\xf9\x19\x4e\x07\xe9\x78\x17\x4b\x3a\xeb\x20\x15\xb2\x91\x3e\x9d\xf5\x66\x25\xd9\x99\x61\x8a\xf5\xc9\x5f\x31\x90\x18\xeb\x01\x23\x26\x87\xb1\x31\x11\x3c\xb1\xe2\x07\x02\x71\x68\x4f\x2b\x86\xc8\xe1\x85\x3e\x95\x78\x29\x9e\x73\x3c\x00\x59\x48\x92\xb0\x66\x37\x9f\xcc\xff\xf2\xf5\x80\xf5\x9a\x70\x48\x59\xaf\x31\xbb\x01\x06\xd9\x0b\x1f\x4f\x47\x25\x76\x1d\x4f\x59\x6f\x30\xaf\x50\xeb\x0d\xeb\xed\xf8\x7a\xbf\xfe\x96\x80\xde\x23\xd8\x62\xd0\x9e\x43\x0c\x84\xc4\x98\x5c\xd1\xef\x0e\xc4\x0e\xbd\xcc\xec\x5c\xe4\x1a\x48\xc6\x8a\xb7\x57\x3c\xbf\x3b\x4c\xeb\x61\x21\x46\xae\xb1\xc8\xff\x95\xe1\x77\xf7\x65\x3c\x00\xd0\xfc\xf1\x4f\x00\x00\x00\xff\xff\x46\xfb\x35\x5d\xa0\x03\x00\x00")

func _4_stored_procedure_get_conversations_up_sql() ([]byte, error) {
	return bindata_read(
		__4_stored_procedure_get_conversations_up_sql,
		"4_stored_procedure_get_conversations.up.sql",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"1_initial_schema.down.sql": _1_initial_schema_down_sql,
	"1_initial_schema.up.sql": _1_initial_schema_up_sql,
	"2_stored_procedure_get_communities.down.sql": _2_stored_procedure_get_communities_down_sql,
	"2_stored_procedure_get_communities.up.sql": _2_stored_procedure_get_communities_up_sql,
	"3_stored_procedure_get_users.down.sql": _3_stored_procedure_get_users_down_sql,
	"3_stored_procedure_get_users.up.sql": _3_stored_procedure_get_users_up_sql,
	"4_stored_procedure_get_conversations.down.sql": _4_stored_procedure_get_conversations_down_sql,
	"4_stored_procedure_get_conversations.up.sql": _4_stored_procedure_get_conversations_up_sql,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"1_initial_schema.down.sql": &_bintree_t{_1_initial_schema_down_sql, map[string]*_bintree_t{
	}},
	"1_initial_schema.up.sql": &_bintree_t{_1_initial_schema_up_sql, map[string]*_bintree_t{
	}},
	"2_stored_procedure_get_communities.down.sql": &_bintree_t{_2_stored_procedure_get_communities_down_sql, map[string]*_bintree_t{
	}},
	"2_stored_procedure_get_communities.up.sql": &_bintree_t{_2_stored_procedure_get_communities_up_sql, map[string]*_bintree_t{
	}},
	"3_stored_procedure_get_users.down.sql": &_bintree_t{_3_stored_procedure_get_users_down_sql, map[string]*_bintree_t{
	}},
	"3_stored_procedure_get_users.up.sql": &_bintree_t{_3_stored_procedure_get_users_up_sql, map[string]*_bintree_t{
	}},
	"4_stored_procedure_get_conversations.down.sql": &_bintree_t{_4_stored_procedure_get_conversations_down_sql, map[string]*_bintree_t{
	}},
	"4_stored_procedure_get_conversations.up.sql": &_bintree_t{_4_stored_procedure_get_conversations_up_sql, map[string]*_bintree_t{
	}},
}}
