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

var __1_initial_schema_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _1_initial_schema_down_sql() ([]byte, error) {
	return bindata_read(
		__1_initial_schema_down_sql,
		"1_initial_schema.down.sql",
	)
}

var __1_initial_schema_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x58\x4d\x53\xdb\x3a\x14\x5d\xbf\xfc\x0a\xed\x70\x66\x58\x10\xde\x83\x61\x86\x61\x61\x12\x85\xe7\x79\xc1\xe6\x39\x4e\xa7\x74\xa3\x11\xb6\x70\xd5\xda\x72\xc6\x56\x42\xe1\xd7\x77\x24\x27\x8e\x24\x7f\xe0\x14\x4a\x4b\xb3\x8a\xaf\x8e\xae\xaf\xa5\x73\xce\x95\x3d\x98\xf8\xde\x0d\x08\xec\xcb\x19\x04\xce\x14\xc0\x8f\xce\x3c\x98\x83\x55\x41\xf2\xe2\x7c\x30\xf6\xa1\x1d\xc0\xcd\xa8\x8c\x01\x6b\x00\x00\x8d\x80\xf1\xbb\xa3\x31\x65\x7c\x7b\xe5\x7a\x01\x70\x17\xb3\x19\xb0\x17\x81\x87\x1c\x77\xec\xc3\x6b\xe8\x06\x87\x03\x00\xa2\xda\xdc\x35\xce\xc3\xcf\x38\xb7\x8e\x4f\x4e\x87\xbb\xa9\x02\x2b\x6e\xc8\x70\x4a\xea\xd8\x91\x80\x6a\x58\x92\x62\x9a\x34\xe6\x3d\xfd\x47\x60\x27\x70\x6a\x2f\x66\x3b\xbc\x9e\x57\xc5\xff\x7d\x6c\xe6\xbe\xa3\x59\x4b\xcd\x27\x23\x01\x36\x73\xd3\x34\xee\x7c\x46\x13\xbf\xcc\x69\x48\x10\xcf\x50\x4a\x8a\x02\xc7\xc4\x5c\x4e\x13\x1f\xe6\x04\x73\x12\x3d\xbb\xfc\x72\x0d\x97\x51\x27\xd6\xcc\x7d\xe3\x3b\xd7\xb6\x7f\x0b\xfe\x83\xb7\xc0\xa2\xd1\x50\xc4\x16\xae\xf3\xff\x02\xca\x90\xd8\x3e\x2b\xaa\xc7\x45\xb8\x1e\xad\x36\xd0\xda\xfe\x1b\x0e\x86\x00\xba\x57\x8e\x0b\x2f\x1c\xc6\xb2\xc9\xa5\x41\x91\x8b\xa3\xaa\xa2\xf1\xbf\xb6\x3f\x87\xc1\xc5\x8a\xdf\x9f\x81\xb1\x37\x9b\xd9\x01\x94\x17\x28\x26\x8c\xe4\x38\x41\x21\x3d\x1f\x34\xd3\x37\xcc\xd2\x74\xc5\x28\xa7\xc4\x24\xb1\x32\xf2\x52\x2a\x3f\xed\x41\xe5\x76\xba\x8d\x8e\xcf\x0c\x6c\xf6\xc0\x48\x8e\x54\xa1\xb4\xe7\x2d\xb1\xaa\x50\xda\x25\x12\x91\x22\xcc\xe9\x92\xd3\x8c\x35\xd2\x58\x93\x53\x11\xe6\xd9\x03\xc2\x69\xb6\xda\xac\x45\x3b\xc5\xda\xe9\xde\x2c\x8f\x04\x17\x1c\xe1\x90\xd3\x35\xe9\x45\xc9\x52\x1e\x4b\x92\x57\xfa\x68\xaf\xa5\x92\xd2\x97\x8c\xb2\x67\xea\x5e\xae\xee\x12\x1a\x6a\x5b\x9f\x65\x09\xc1\xcc\xa8\x63\xf4\xb3\x35\x17\x91\x84\x18\xf8\xc6\x42\x8e\x7a\x08\xb4\x59\x88\xa5\x08\xa5\x00\x8d\x11\x41\x62\xeb\x69\x33\x63\xec\xb9\xf3\xc0\xb7\x1d\x37\x50\x75\x22\xe9\x25\xf8\x28\xff\x14\x68\x42\x23\x74\x9f\xe5\x84\xc6\x0c\x4c\x3d\x1f\x3a\x57\x6e\x59\x4d\x45\xdc\x21\xf0\xe1\x14\xfa\xd0\x1d\xc3\xf9\xb6\x6d\xc8\xb0\xe7\x82\x09\x9c\xc1\x00\x02\x1f\xce\x03\xdf\x19\x07\x22\xb4\xb8\x99\xd8\x4a\xe8\xad\x2c\x82\xad\x49\x5e\x60\x21\x87\xba\x49\x28\x63\x0d\x36\xf1\x83\x16\xd1\xdd\x05\xb6\x2b\xfe\x88\xc4\xac\xee\xae\xb8\x73\x87\x76\x1c\x27\xdf\x78\xdf\x7b\x27\x94\x7d\xed\x8b\xd5\xe5\xde\x8d\x5d\xd3\x88\x64\x3d\xb1\xba\x1c\x3b\xa4\x58\x02\x91\x94\x7a\x9b\xb4\x0e\x8e\x0e\xea\xaa\xed\xab\xd8\x7d\xd4\xfa\xda\x4a\xd5\xf4\x28\x02\x3a\x2d\x2c\xed\xb2\x2e\x5a\x85\xb7\x48\x83\x22\x55\xd0\x9f\xda\x14\xac\x67\x57\x55\xac\xf5\xcd\xa7\xde\x5a\xee\x28\xb0\xb7\xad\x6c\x71\xef\xc1\x55\x38\x8e\x4d\x33\x11\x21\xe9\x21\x1c\xc7\xe0\x83\xed\x8b\xf4\xf2\x50\x5a\xf9\x86\x42\x94\x5a\x95\xaf\x75\x1a\x22\x8c\x37\x1d\x85\x44\xf8\x2d\xcf\x41\x1a\x01\xe4\xc4\xbe\x46\xb7\x8f\xd9\xed\x6b\x78\xcf\xe3\xdf\x61\xfb\x2f\xcd\xc3\x58\x6e\xcb\x8c\x54\xd0\x6a\xb5\x77\x7a\x6b\x38\x12\x08\xbe\x20\x33\x07\xd2\x7d\xa7\xc3\x5d\x8c\x7b\xeb\x06\xa3\xf5\xdc\x17\x58\xcc\xa6\xca\x3f\xd1\x5e\x76\xfe\xdc\x60\x34\xfa\xa0\x54\xb5\x79\xaa\x00\x95\x01\x95\x34\xd7\x78\x5b\x3a\x94\xfa\x53\xdd\x4a\x03\xff\x05\xfa\xf6\x0d\xb5\x71\x54\x7d\xcd\x98\xcd\x71\x3c\x54\xee\xaa\xce\x16\x4f\x22\xc7\xcb\x69\x1b\xa6\xd7\x1b\xa3\x78\x64\xa3\x8c\x43\x31\xf9\x8d\xdf\x37\x1f\x51\xd3\x87\x13\x63\xb4\xd1\x6f\x35\x8b\xe8\x30\x5b\x73\x3f\x35\xdb\xea\x72\xcf\x56\xa0\x78\x53\x22\x11\x12\x9e\xd5\x5e\x89\xf4\x4c\x72\xcf\x77\x30\x03\x58\x33\x58\x01\xce\x09\x2e\x36\x2f\x9b\xda\xd7\x98\xae\x8f\x0d\x1d\x5e\xb4\xff\x71\x48\x5b\xf7\xdf\xf3\x40\xa4\x97\xf8\x8b\x4d\x6b\x1f\x2d\x0c\xbe\x07\x00\x00\xff\xff\xdd\x02\x28\x6e\x3a\x14\x00\x00")

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

var __2_stored_procedure_get_communities_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x93\x4d\x6e\xdb\x30\x10\x85\xd7\xe2\x29\x66\x27\x0b\x30\x7c\x81\xa2\x0b\xd9\x66\x8c\x18\x89\x64\xc8\xf2\x22\x2b\x82\xa5\x58\x81\x85\xfe\x20\x92\x0d\xea\xd3\x17\x1c\xb2\x36\xfd\xd3\x64\x35\xf3\xde\x47\x86\x7a\x33\xde\x54\x34\xaf\x29\x1c\xaa\x72\x43\xb7\xa7\x8a\x42\x2b\x0d\x13\x63\xdf\xdb\x41\x19\x25\xf5\x22\x23\x6b\xba\x7b\x2d\x08\x39\xd2\x37\xba\xa9\x81\x24\xfb\x63\x59\xb0\xbc\xaa\xf2\x8f\x7c\xb7\x5b\x60\x57\xae\xf7\x74\x53\x2f\x48\x92\xa4\xaa\x49\x97\x20\x56\xaa\x59\x02\x81\xf0\x97\x0e\xbc\x97\x28\xbb\x22\x36\xce\x01\x3f\xdf\xf2\xe3\xe7\x20\xe7\x6d\xf0\xb0\x61\xcd\x13\xe2\xa4\xe5\x7c\xb9\xda\x63\x36\x48\x31\xdb\x48\x2d\x66\x35\x19\x35\x0e\x48\x46\x7d\x8c\x39\x75\xfc\xcc\xfb\xd1\x0e\x06\x39\x2f\x30\x8e\x4a\x4c\xaa\xbe\xf5\x1f\xd9\xb7\xb1\xdc\x71\x6d\xb8\x30\xea\xb7\x7f\x90\x6b\x99\xef\x63\x6a\x9a\x95\x90\x07\x39\xbf\x4b\xad\x79\xeb\x51\xd4\xd8\x24\x67\xd6\x7b\xf5\xe1\x40\x3d\xee\x47\x35\x44\xb0\x19\xd9\xaf\x51\xdd\x7c\xc0\x64\x7f\x74\x4a\x78\x06\xcb\xd8\x14\xb3\xe4\x46\xfa\x44\x43\xbd\x04\x37\x30\x3b\x35\x17\x23\xd4\xb7\xe1\x75\xf2\x9f\x1d\xea\xd8\x36\xbc\xd5\xe9\x12\xcc\xca\x15\xb1\xe1\x06\xe1\x1c\xbb\xc2\x2a\x23\x49\x06\x5c\x43\x5a\x49\x6d\x3b\x93\x92\x97\xaa\x7c\x07\x92\x44\xab\x06\x82\xbc\xd1\x97\x1a\xf6\xe5\x6b\x01\x0b\x92\x5c\x16\xee\x02\xfd\x61\x7e\x4f\x92\xbb\x1d\x14\xf8\xef\x33\xc8\x8f\xe0\x9e\x41\x12\xbc\xfc\x7a\xca\x89\x20\x0c\x49\x76\x55\x79\x3a\xc0\xfa\x03\x84\x59\xdd\x5c\x4a\x32\x30\x50\x16\x70\x27\xc3\x77\xbf\x9b\xff\x79\xd8\xc3\xbb\xbe\xfe\x69\xb8\x20\xc2\x52\xfb\x54\x70\xa9\xaf\x91\xb9\x79\xca\x66\xcb\x8d\xf4\x88\xef\x99\x1b\x49\x44\x75\xf2\xa7\xb9\x32\xae\x7b\x46\x54\x92\x6b\xbf\xee\x81\x99\x51\xc0\x31\xe4\x47\xc0\x99\xe0\x89\xbb\xa8\xd0\x00\x61\xd1\xbb\xc6\x65\x1f\xe2\xb2\x2e\x2e\xfb\x3c\xae\x6f\x84\xd0\x62\x4b\xfe\x06\x00\x00\xff\xff\xf4\xe0\xcf\x06\x5c\x04\x00\x00")

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

var __3_stored_procedure_get_users_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xcf\x6e\x83\x30\x0c\x87\xcf\xf1\x53\xf8\x16\x90\xd0\x5e\x60\x27\x4a\x3d\xb4\x6a\x85\x29\x65\x87\x9d\x50\xda\x58\x28\x52\x33\x2a\x42\xde\x7f\x5a\xc8\xba\xac\x39\xfd\xfc\x7d\x71\xfe\xb8\x51\x54\x0f\x84\xef\xaa\x6f\x68\xff\xa1\x08\x27\x5e\xc7\xe0\x79\xf1\x45\x09\x3b\x6a\x5f\x3b\x80\x13\xbd\x51\x33\x20\x88\xc3\xa9\xef\xc6\x5a\xa9\xfa\xb3\x6e\xdb\x22\x56\xfd\xee\x40\xcd\x50\x80\x10\xd2\x1a\x59\x61\x78\xb2\xa6\x42\xc0\xb4\xa4\x49\xd4\xfc\xc7\x3f\x37\x7c\x69\xc7\xd1\xfd\x16\xd5\x9f\x67\xa7\xed\x35\xca\x98\x32\x73\xef\x7a\xe8\x38\xdb\x39\xe2\xb3\x9d\x33\x6a\xdd\xb4\x3d\xca\x4d\x19\xbd\x2d\xf6\xc2\xc3\x7c\x64\xef\xf5\xb4\x9d\x16\xd1\xb8\xce\xa3\xdb\x60\xb6\xfb\xb2\xb0\x5e\x79\xfb\x46\xca\x99\x0d\x37\x73\xb7\x29\x97\x20\x4a\xd4\x1e\xa5\x62\x1f\xae\xab\x84\x17\xd5\x1f\x11\x44\x1c\x2b\x86\x67\x00\xea\xf6\xdf\x01\x00\x00\xff\xff\xde\x59\xe8\x0a\x79\x01\x00\x00")

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

var __4_stored_procedure_get_conversations_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x93\x41\xaf\xa2\x30\x14\x85\xd7\xbd\xbf\xe2\xee\x80\xc4\xf8\x07\x26\xb3\x40\xac\x44\xe3\x80\xa9\xb8\x70\x36\xc4\x29\x8d\x69\x06\xa8\x81\x62\xde\xf3\xd7\xbf\x50\xc0\xd7\xa2\x79\xac\xce\x3d\xe7\x2e\xb8\xe7\x4b\x23\x46\xc3\x8c\xe2\x81\xa5\x11\x5d\x9f\x18\xc5\xab\xd0\x39\x57\xf5\x5d\x34\xed\x45\x4b\x55\xb7\x7e\x00\x2b\x1a\x6f\x13\x80\x23\xdd\xd3\x28\x43\x20\xbb\x63\x9a\xe4\x21\x63\xe1\x39\x8c\x63\xdf\x4c\xe9\x6a\x47\xa3\xcc\x07\x42\x3c\x59\x78\x0b\xe4\x4b\x59\x2c\x10\x70\xfc\xbc\xc7\xe8\x3e\x64\xb1\xf8\x76\xb9\xaa\xaa\xae\x96\xfa\xf3\xef\x18\x3f\x8d\xdc\x5d\xec\x5a\xd1\xac\xc7\x9d\x5e\xe7\x85\x13\x6b\xf1\xa1\x4d\xd6\x0b\xcb\x2f\x65\xfd\xdf\xf8\xbd\xb0\x7c\x59\x5d\x87\x7f\xac\xae\x96\x7b\x97\x85\x50\xc6\x37\xca\x4a\x6e\xdd\xbf\x52\x72\x13\x0d\xf2\x25\x3b\x34\x92\x0b\x6b\x21\xbf\xf5\x86\x7d\x6a\x23\x2e\x5a\x8c\x57\x0e\xda\xbe\xef\x56\x3c\xd3\x51\x5b\x69\x21\x4a\x31\xa5\xa3\x9e\x95\x28\x6a\xdd\xf6\xb1\x5a\x4e\x53\x00\x24\xc0\x4b\x8b\x1e\x13\x6d\x57\x6a\x0f\x36\x2c\xfd\x83\x40\x1c\xb6\xc8\x61\x4f\x37\x19\xee\xd2\x6d\x82\x3e\x90\x27\x61\x67\xcd\xb0\xe8\xcd\x1f\xb9\x4f\xe0\x95\x4b\xde\xa6\xaf\x5c\xfc\x73\xb2\xea\x0d\x5a\x07\xaf\x9a\xf1\x75\x18\xab\x19\xe4\x79\xeb\xea\xb5\xf6\x79\xf5\xea\xb5\xfb\x79\xff\x6a\x02\x10\x00\x21\x01\x86\x47\x9c\x0a\x07\x62\x1a\x9e\x46\xe4\x0a\x48\xcc\xd2\xd3\x01\x57\xe7\x01\x8c\xdb\x28\x04\xc8\x15\xa6\xc9\xbb\x0c\x7f\x0f\x4f\xe5\x17\x00\x4d\xd6\x5f\x01\x00\x00\xff\xff\x96\x4d\x72\x64\xa4\x03\x00\x00")

func _4_stored_procedure_get_conversations_up_sql() ([]byte, error) {
	return bindata_read(
		__4_stored_procedure_get_conversations_up_sql,
		"4_stored_procedure_get_conversations.up.sql",
	)
}

var __5_payment_table_down_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _5_payment_table_down_sql() ([]byte, error) {
	return bindata_read(
		__5_payment_table_down_sql,
		"5_payment_table.down.sql",
	)
}

var __5_payment_table_up_sql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x54\xdf\x6f\x9b\x30\x10\x7e\xf7\x5f\x71\x6f\x4d\xaa\x55\x6d\xd3\x6d\x9a\x14\x65\x12\x05\xa7\x43\xa3\xd0\x39\x8e\xb4\x3e\x21\x17\x2e\x99\x57\x30\xc8\x36\x95\xd2\xbf\x7e\x32\x85\xfc\x6e\xb7\xa9\x3c\x71\x1f\xdf\xf7\xf9\xee\xb8\x33\xf1\x22\x4e\x19\x70\xef\x3a\xa2\xd0\x18\xd4\x86\x00\x78\x41\x00\xa2\xac\x1a\x65\x61\x51\x54\xc2\x42\x3c\x8f\x22\x08\xe8\xd4\x9b\x47\x1c\x4e\x2e\x4e\x08\x00\x80\x37\x75\x4a\x25\x4a\x1c\x93\x43\x1b\xb8\x4d\x82\x70\x7a\x0f\xb5\x96\x19\xa6\xb6\x4a\x4b\x34\x46\x2c\xb1\x77\x4c\x78\xeb\x3a\xde\x51\x66\x55\x59\x36\x4a\x5a\x89\x7b\xfa\x1a\xf5\xfb\x0c\x6c\x95\xfe\xae\xa4\x3a\x10\x13\x12\xb0\xe4\xae\x53\x87\x53\xa0\x3f\xc3\x19\x9f\x41\x2d\x56\x25\x2a\x6b\xc6\xc4\x67\xd4\xe3\xb4\x23\xf4\x30\x0c\x08\x80\xcc\xe1\xc8\xf3\x20\x97\x52\xd9\x3e\xea\x0f\x02\x6f\xce\x93\x34\x8c\x7d\x46\x6f\x69\xcc\x3f\x10\x80\xe7\xa3\xfa\x27\xa1\xb3\x5f\x42\x0f\x46\x9f\x3e\x0f\x37\x72\xc7\xd7\x98\xc9\x5a\xa2\xb2\x69\xbe\xa5\x7c\x9d\x6f\x50\xe5\xa8\x77\xc8\x6f\xf2\xbb\x1f\xbe\xf7\xbc\xf4\x6b\xbf\x1c\xc7\xb7\xab\x1a\x0f\xf3\x37\xa5\x28\x8a\x75\x03\xd6\x23\x73\x79\xd2\xa6\x64\x85\x6d\xcc\x7f\x49\x34\x16\xf8\x24\x94\x4d\x77\x8e\xeb\xab\xb8\x1a\x0d\x5d\x56\xeb\x0e\x75\xdc\x9d\xd6\xee\x55\x7c\xd0\x4d\x55\xe5\x98\xd6\xcd\xc3\x23\xae\xfe\xa9\xfb\x1a\x0b\xb1\x4a\x1b\x5d\xbc\xd9\xcd\x4c\xa3\xb0\xb8\xff\x87\x8f\x0f\x87\xe3\xdf\xb1\xf0\xd6\x63\xf7\xf0\x9d\xde\xc3\x40\xe6\x43\x87\xcd\xe3\xf0\xc7\x9c\xb6\x90\xcc\x8f\xa1\xae\xce\xc1\x73\x87\xfb\x49\x3c\xe3\xcc\x0b\x63\xde\x8f\x69\xba\x33\x32\x69\xbb\x96\xed\xdb\xa2\xd2\x28\x97\x0a\xa6\x09\xa3\xe1\x4d\xfc\x72\xe8\x0e\x79\x08\x8c\x4e\x29\xa3\xb1\x4f\x67\xdd\x3e\x0f\x5a\x38\x89\x21\xa0\x11\xe5\x14\x18\x9d\x71\x16\xfa\xdc\x41\xf3\xbb\xc0\xdb\x82\x5e\x49\x67\x33\x91\x7f\xcb\x65\xc3\x7c\x5f\x22\x64\x08\x34\xbe\x09\x63\x3a\x09\x95\xaa\x82\xeb\xbd\x35\x9c\x5c\xac\xc7\xcd\xff\xe6\xb1\x19\xe5\x93\xc6\x2e\xbe\x80\x9f\x44\x91\xc7\x69\x1b\xa4\x4b\x54\xa8\x45\x91\x66\x72\x4c\xc8\xf9\x29\x21\x67\xed\xe8\x1b\x68\x2f\xc1\x33\xb8\x84\xc9\x57\x78\xa8\x2a\x63\x3b\x60\xe4\x80\x7a\x54\x77\xe1\x95\x0b\xfb\x7b\x69\xd5\xde\x42\xdd\x97\x8f\xee\x8b\xb1\xe2\x11\x9d\xe9\xcb\x72\xa0\xd9\xb6\xad\x51\xe5\x52\x2d\xb7\x8d\xb3\xaa\xac\x0b\xb4\x98\x6f\xdb\x2f\x84\x2c\x30\x27\xe4\xf4\x9c\xfc\x09\x00\x00\xff\xff\xb9\x3e\x79\x08\xd1\x05\x00\x00")

func _5_payment_table_up_sql() ([]byte, error) {
	return bindata_read(
		__5_payment_table_up_sql,
		"5_payment_table.up.sql",
	)
}

var _bindata_go = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x9a\x49\x6f\x25\xc7\x72\x85\xd7\xbc\xbf\xe2\x9a\xc0\x33\x48\xa0\xd1\xae\x79\x78\x80\x36\xef\xc1\x5b\x2f\xbc\x75\x1a\x44\x56\x65\x56\x9b\x70\x73\xd0\x25\xa9\x97\x92\xa1\xff\x6e\x7c\x11\xa7\xcc\x76\x83\x94\x48\xb5\x80\x5e\x74\x5f\xde\xa1\x32\x23\x63\x38\x71\x22\x22\xef\xe3\xfa\xdf\xf1\x53\x3e\xde\x5c\x7f\x3a\xc5\xc7\xeb\xbb\xdb\x87\xc3\xe1\xfa\xe6\xfe\xee\xf4\x78\xbc\x38\x9c\x9d\x2f\x3f\x3f\xe6\x87\xf3\xc3\xd9\xf9\x7a\x77\x73\x7f\xca\x0f\x0f\xff\xf2\xe9\x97\xeb\x7b\x3e\xd8\x6e\x1e\x79\xb9\xbe\xe3\xff\x87\xc7\xd3\xf5\xed\xa7\x87\xf3\xc3\xe5\xe1\xb0\x3d\xdd\xae\xc7\xe5\xfa\x36\xc5\xc7\x78\x75\xca\x31\x5d\xf0\xd7\xf1\x3f\xfe\x93\xb5\x3e\x1c\x6f\xe3\x4d\x3e\xfa\xef\x2f\x8f\x17\xfb\xa7\xf9\x74\xba\x3b\x5d\x1e\xff\xe7\x70\xf6\xe9\x17\x7b\x77\xfc\xeb\x0f\x47\xb6\xfa\xf8\x6f\xf9\x1f\xff\x9e\x63\xca\xa7\x0b\x93\x85\xf7\x7f\x7b\xda\xb6\x7c\xb2\x65\x2f\x2f\x0f\x67\xd7\x9b\x3d\xf0\x4f\x3f\x1c\x6f\xaf\x3f\xb3\xc4\xd9\x29\x3f\x3e\x9d\x6e\x79\xfb\xe1\xb8\xdd\x3c\x7e\xfc\x57\x56\xdf\x2e\xce\x59\xe8\xf8\x97\x1f\xff\x7a\xfc\xcb\x4f\xe7\x2e\x89\xed\x75\x79\x38\xfb\xf5\x70\x38\xfb\x29\x9e\x8e\xcb\xd3\x76\xf4\x7d\x7c\x93\xc3\xd9\x95\x8b\xf3\xc3\xf1\xfa\xee\xe3\xdf\xef\xee\x7f\xbe\xf8\xe7\xe5\x69\xfb\x70\xfc\xf4\xcb\x25\xb2\x7e\xfc\xfb\xe7\xbb\x87\x7c\x71\x79\xf8\xf3\xc4\xd0\x53\xcb\xd3\xf6\xf1\x6f\x48\x72\x71\xf9\x81\x25\x0e\xbf\x1e\x0e\x48\x78\x75\x55\x5f\x5d\xdf\x5e\x3f\x5e\xc7\xcf\x57\x0f\xeb\x7f\xe5\x9b\x78\x95\xee\xfe\x71\x7b\xf5\xf0\xe3\xe7\xe3\x0f\xd2\xf2\xc5\x79\x28\xf5\x16\xca\xb4\x84\x52\x4d\xa1\x54\xd5\xcb\xff\xb6\x2d\x94\xaa\xfe\xff\xef\xed\xb3\x57\x7e\xef\xff\xce\x77\x23\xbf\x2e\xc9\xc5\x8b\xa6\xdd\x4f\xf6\xa5\x6f\x1c\xce\xce\x7e\xe3\x44\x1f\x0e\x67\x67\xe7\x5f\x7f\xfb\x91\x6f\x3f\x3e\xfc\xf8\xf9\xfc\xc3\xe1\xec\xf2\x37\xf4\xf2\x74\xff\x87\xb5\x92\xd6\x50\xfa\x29\x94\x2e\x85\xd2\xb7\xa1\xa4\x25\x94\x36\x86\x52\x77\xa1\xf4\x29\x94\x05\x4d\xad\xa1\x54\x31\x94\x9c\x42\x19\xab\x50\x86\xc1\x9f\xa9\xab\x50\x52\x0e\x65\x6a\x43\x19\xea\x50\xa6\xc1\x5f\xf9\x57\x37\xa1\x4c\x7d\x28\x79\x0c\x65\x9c\x43\x59\xeb\x50\xf2\x10\x4a\x3b\x87\xd2\xe5\x50\x22\x9f\x77\xa1\xc4\x36\x94\xba\x0e\x65\x19\x7c\xed\xd4\x87\x92\x62\x28\x63\x13\xca\xca\x3e\x43\x28\x5d\x13\x4a\xae\x43\x49\x3c\x33\x86\xd2\x74\xa1\x34\x63\x28\x53\xf6\xbf\xc7\x2d\x94\x5c\xb9\xcc\x5d\x0c\xa5\x5b\x42\x59\xda\x50\xa6\x18\x4a\xdc\xfc\x77\x31\xfb\xdf\xb1\x0f\x65\x6c\x43\x59\x73\x28\x73\x1f\x4a\x9b\x42\x99\xa7\x50\xb6\xc9\xcf\x52\x25\xd7\x59\x5e\x43\x59\x97\x50\xea\x39\x94\xaa\xf3\xdf\xb3\xfe\x5a\x85\x32\x6d\xfe\xbe\x5d\xfd\x59\xce\xdf\xf7\xa1\x74\x75\x28\x1b\xb2\x36\xa1\x8c\x6b\x28\x2d\x3a\x1e\x42\x89\xe8\x23\xf9\xb3\x6b\x0a\x25\xf2\xcc\xea\x1e\x39\x2c\xcf\xf6\x98\x52\x28\x13\xcf\xa0\x8f\xc5\x75\xd3\xf2\x9b\xde\xd7\x1b\x97\x50\x32\x7f\x47\x7f\x16\x7d\xd5\xa3\x9f\x15\x39\x17\xbd\x9f\xb0\xc5\x18\x4a\xbd\xba\xbe\xec\x2c\xad\xef\x95\xf1\x87\xc1\xbf\xaf\x5a\xdf\x37\x36\xae\x73\x7c\xa1\xed\xfd\x6c\xfc\x7e\xdd\x42\x69\xa7\x50\x16\xe9\xba\xdb\xdc\x76\x3c\x6b\xf2\xa1\x9b\x26\x94\x66\x09\x65\xe0\x99\xc9\x65\xc2\x06\x99\xf5\xa6\x50\xe6\xda\xcf\x34\x44\xf7\x99\xb9\x09\x65\x68\x42\x99\x63\x28\x6d\x27\x9f\x48\xa1\x6c\x29\x94\x6e\x0c\x65\xa8\xdc\xb6\xa3\x9e\x69\x36\xf7\xb7\x16\x9b\xa2\xef\x1c\x4a\x3f\x86\xb2\xf6\xee\xbf\x63\x0a\x65\x58\x7d\x9d\x25\xbb\xce\xfa\xd9\x7d\x00\x5d\xb3\x56\xd3\xba\xbe\xda\xc1\x6d\x9f\x5a\xed\x9d\x5d\xaf\xdd\x10\x4a\xdd\xba\xaf\xaf\x6b\x28\x03\xcf\xcb\xbf\xd1\x41\xaf\x33\xe1\x4f\x9c\x77\x65\x7f\xe2\x63\x75\x7d\xf0\x2c\xf1\x86\x0c\xf8\x0a\x7b\xe0\xff\x9c\x0b\x3d\x11\x3f\xf8\x33\xbe\x35\x8f\xa1\xf4\xb5\xcb\x95\x06\xdf\x8f\x33\xe4\xd6\xcf\x88\x8f\x5b\x2c\x80\x4e\x3c\xab\xf8\xc2\x06\xeb\x6e\xab\xda\x65\x42\x86\x7a\x70\xdf\xde\x5a\x21\x1a\xf2\xc5\x50\xe6\xca\xf5\xdf\x66\xb7\x4f\x8c\x2e\x77\xd7\x87\xb2\x70\x36\x74\x95\x5c\x27\xa9\x72\xfb\xb3\x57\x8d\x3f\x66\x8f\x65\x7c\x63\x89\xae\x6f\xec\x8f\x3c\xd8\x03\xbf\x22\x36\x17\xf6\x9a\x3d\xa6\xf0\x7d\xb3\xb5\xfc\x96\xfd\x90\xd5\x7c\x79\xf3\x33\xb5\x9b\xfb\x0e\x58\x80\xfd\xd0\x67\xda\x42\x99\x37\xf7\xd9\x05\x1b\x8d\xa1\xcc\xe8\x61\x76\x9b\xcd\x9d\xfb\x74\x33\x85\xd2\x0c\x1e\x37\xb9\x73\xbf\x64\xcf\xa6\x76\x3d\xa3\x7f\x6c\x6a\xb6\x1d\x5d\xaf\xa9\x71\x8c\x61\x2f\x9e\x8f\xb2\x1b\x98\x61\x7b\x35\x1e\xaf\x3d\x3e\x1c\x5d\x3e\x30\x8e\x75\x2d\xce\x06\x3f\xbf\xc5\x65\x76\x5f\xa8\x56\xf7\x37\xe2\x1a\x7b\xe0\x27\x7c\xd7\x77\x1e\xe7\x4b\xe7\x98\xd7\x24\xdf\x03\x3d\xf4\xd2\x6f\xd3\xbb\x5f\xa0\xa7\xa6\x71\x19\xf3\xec\x31\x81\xdc\x93\xe2\x0f\xb9\x2d\x56\x5a\xc7\xe4\x5a\xd8\x69\x3a\x21\x2e\x1b\xf7\x51\x74\xca\xba\x71\x75\xbb\xe2\x3f\xab\xf0\x93\x75\xc1\xb3\x66\x75\xbc\xaa\xa4\x47\x6c\x07\x96\x9a\x7f\xb4\xbe\x1f\xbf\xe5\x79\x6c\x89\x2e\xf0\x0d\xe2\x07\xb9\xb0\xfd\x16\x43\x19\xa6\x67\xfc\x34\x79\x1b\xf7\x65\x64\xc6\x0f\x86\xd1\x63\xbe\x17\xd6\x12\x5b\x9c\xb5\x8e\x8a\xfd\xc5\xed\x0c\x96\x55\xca\x07\x9c\x1f\x2c\x1c\xf0\xdd\xce\x71\x0c\x79\x90\xd3\x7c\x09\x99\x85\xcd\xab\xf2\x00\xba\x36\xfc\x9b\x5c\xd7\xc4\xfa\x34\xb9\x5c\xd5\xfe\x2a\x2c\x45\x6f\xd8\x12\xbc\xc9\x8b\xeb\xac\x6d\x7d\x3f\x30\x71\x99\x3d\x6e\xb0\x23\xf1\x05\x46\xe3\x3b\xd8\x6b\xec\xdd\x46\xe8\x06\xbb\x21\x03\x3e\x44\x8c\x91\x07\xd0\x07\xf2\x80\x2d\xf8\x18\x67\xc3\x67\x90\x19\x59\xf1\x93\x6e\xfd\xc2\x37\x15\x47\xe0\xa4\xe1\x53\x72\x9d\x83\xb9\xf8\xf0\x56\xb9\x5d\x58\x1b\x7d\x13\x93\x96\x33\xd0\x8d\x6c\x3d\x24\xc7\x58\x70\x1e\x1f\x9b\x1a\xd7\xe5\xbc\xba\x3e\x2d\x8f\x34\xbe\x1f\x18\x8e\xaf\x10\x5f\x96\xc3\x46\x3f\x0f\x3a\x23\xfe\xe7\xd9\xf1\xc1\xb0\x7a\x72\x2c\x30\x5f\x69\xfc\xef\x6e\x76\x5b\x1a\x1e\x2b\x9f\x10\x0b\x7c\xde\xc8\x66\x95\xf0\x83\x18\x07\xa3\xc0\x22\xb0\x29\x91\x43\x7b\x97\xb9\xd9\x31\x6c\xf2\xd8\xc3\x57\xb3\x7c\x04\x39\xc0\x4a\xd6\x26\xd6\xf0\x4d\xc3\x45\x70\x7d\xd2\xeb\xe2\x9f\x81\x03\x86\x93\xc9\x31\x06\x9e\x80\xbe\x8d\x7f\x24\xe7\x2a\x16\x43\xc8\xd6\x09\x7b\x6b\x97\x99\x75\x89\x8d\xd8\x79\x5c\x55\x8a\x41\x5b\x33\x3a\xde\xa0\x17\xc3\xbd\xec\xaf\xf8\x72\x14\xe7\x89\x8b\xeb\x8a\xfd\xd3\x6e\x9f\xe8\xba\x43\x0e\x7c\xb7\xd1\xba\xd8\xb9\x91\x2e\x79\x96\xd8\x68\xb0\x79\x74\x3e\x84\x9f\x75\xe2\x0c\xd8\x03\xde\x80\xad\xf1\x9f\x24\xde\x85\x8e\x88\x27\xfc\x85\xef\x79\xbf\xf5\xe2\x0e\xa3\x9f\x99\xf5\xd0\x63\x23\x7f\x60\xdd\xa8\x5c\x38\x82\x17\x9b\xfb\x04\xcf\xb7\xc2\x28\x7c\xae\x15\x9e\x60\x53\xf3\xed\x5a\x78\x3e\x3a\x87\x42\xae\x5d\xc7\x9c\xcb\xf0\xb1\xf2\xfd\xf0\xc3\x49\x71\x88\xee\xf8\x0c\x0e\x05\x2e\x10\x2f\xc4\x3a\x67\x32\x1c\x53\x3c\x71\xe6\xae\x53\xfc\x24\xc7\x50\x62\xb2\x16\x87\x01\xcb\xf1\x0b\x72\x37\x79\xd8\xb8\xc6\xe6\x58\x87\x1c\x70\x16\x7c\x81\xe7\xb2\xfc\x19\x7e\xb9\x29\x17\xad\x8a\x11\xec\x6a\x1c\x64\xf1\xfd\xc1\xe3\x24\x0e\xb1\xf4\xce\x83\xb0\x35\xeb\xa1\x0f\xb0\xc9\x30\x04\x1d\x56\xae\x17\x62\xc4\x72\x86\xf4\x6c\xeb\x26\x5f\xcf\xb8\xd7\xe2\xdc\x89\xfc\xb6\x29\x76\xc1\x2a\xf6\x36\x2c\x1b\x3c\xf7\x92\x47\x2c\x47\x67\xf7\x6b\xfc\x9f\xdf\x12\x4b\xad\xfc\x27\x56\x2e\xa3\xc9\x36\xbb\xaf\x11\xbf\xe8\xd3\xf2\x5e\xaf\x7c\x36\xfa\xf9\x7b\xf1\xf2\x59\xcf\x1b\x57\x4f\xee\x83\xd5\x9e\xfb\x7b\xff\x8e\xdc\x44\x6c\xda\xda\xb5\xef\xc3\x99\x56\x71\x5d\x5e\xf9\xbc\x95\x9e\x2c\x1e\xa3\x38\xee\xe0\x3c\xaf\xd9\x79\x8e\x70\x13\x5e\xc0\xf9\x90\xa7\x56\x8e\x1f\x15\x6b\x51\xb9\xab\x13\x5f\x35\x3f\x49\x6e\x2b\xf3\x83\xcd\x39\x4f\x52\xac\xd7\xc2\x60\x62\x0e\x3d\x0d\x9b\xe3\xfe\x24\x7e\x0a\x6f\xe6\xbc\x60\x43\xab\xef\xa8\x4b\xf2\x5e\x17\x6c\x5e\x17\xa0\x43\xd3\x8d\x62\xa9\x56\xdd\x01\x47\x24\x2e\xd1\xff\x34\xbb\x1f\x77\xeb\xf3\xb3\xf8\x02\x7b\x54\xf2\x9b\x56\xaf\xe8\xdd\x70\x73\xf6\x98\x05\xdb\xf0\x31\xe3\x23\xd1\xcf\x8e\xfe\xc0\x2e\xcb\xa9\xc9\x7d\x9c\x7d\x38\x8b\xf1\xd1\xda\xe3\x1b\xf9\x37\x61\x19\x39\xa6\xab\x1c\x7f\xc8\x99\xc4\x02\x67\xed\xe4\xdf\xe8\x87\xd8\x20\x67\xe0\x8f\xf8\xe1\xd7\x95\xaa\x61\x6a\x23\xec\xcd\xcf\x35\xda\xef\x56\xaa\x5e\x1b\x7e\x6b\x9d\xea\xab\xbc\x5c\xa5\x3e\xdd\xbf\x50\xa3\x36\x57\x0f\x8f\x77\xa7\x9c\xae\xee\x4f\x77\x6b\x4e\x4f\xa7\x7c\xf5\x29\x3f\x5e\xad\x77\x37\x37\x4f\x3c\x9e\x1f\xbe\x5f\x35\xff\x0e\xd9\xde\xa9\xb7\x77\xac\x6c\xba\xfc\xfd\xdf\xbf\xd6\x03\x78\xc3\x4e\xdf\xd0\x15\xa0\x2a\x82\x7d\x9a\x87\x66\x75\x05\x2a\xaf\x88\x40\x74\x22\x81\x2a\xb7\x99\xbd\x3a\x23\x43\x57\xfa\x0d\xcf\xc2\xe6\x61\xff\x95\x98\xfc\x20\xe6\x0e\x9b\x23\x22\x41\x0d\x98\x0c\xd9\xa5\x51\x05\x69\xcc\xa3\xf7\x48\xb4\x0a\xb6\x77\xe4\x04\x9d\x67\x55\x4f\x56\x55\xb6\xcf\x4c\x96\xec\x6a\x0c\x44\xd9\xdf\x32\x7f\xef\x11\x47\x96\xa4\xa2\x24\x63\x81\xfc\xad\x3e\x03\xf1\x40\x56\xb2\x16\xf2\xb3\x4e\x14\xb2\x74\xaa\xb4\x40\x7f\xaa\x93\x4e\xb2\x59\xa4\x0b\xb5\x88\x68\xf4\x61\x59\x6c\x67\x87\x62\x8b\x8d\xaa\x6c\x58\x87\x31\x17\x75\x13\xda\xd9\x33\x4b\x2b\x16\x45\xe5\xcc\x19\xa9\x44\xc8\x64\x83\xaa\x56\xcb\xe4\x92\x07\xdd\x4c\xb2\x05\x32\x81\x28\x64\x10\x58\x81\x55\x7c\xea\xa0\x18\xfa\x6d\x8e\x3a\xe8\x09\x94\x89\x62\x6a\xa0\x0d\xfa\x33\xb4\x8c\xbe\x07\x28\x32\x89\x71\xf2\x3d\x6c\xc7\x2a\x6b\x55\x59\x54\xac\x8d\xf4\x6a\x68\x58\x2b\xa3\x6d\xea\x6e\xb4\xaa\x70\xc4\x42\x61\xa5\xac\x85\xec\x64\xe6\x55\xba\xca\x92\x25\xab\x7b\x01\xdb\x06\x45\x39\x2b\x6b\x23\x2f\xd9\xc2\x74\x29\x76\x60\x95\x7e\xef\x7a\xe5\x15\xd9\xf8\x0e\x66\x30\x2a\x13\xa3\x4b\xd8\x06\x55\x35\x59\x17\xfd\x91\x45\xf8\x6d\xa5\x0e\x13\xfa\xc0\x17\x8d\xf5\x44\x47\x7f\xce\x0f\xba\x62\xab\x5a\x0c\x03\xbd\xc2\x12\x93\x58\x13\x67\x59\x94\x81\xf0\x35\xbe\xa7\x12\xa7\x5a\xc2\xae\xb0\x66\x98\x15\x67\x80\x59\xcc\x62\x40\x7c\x0f\x6a\xb3\x37\xd5\x2e\xd5\x13\x6c\x8c\xcc\x0e\x6b\xb0\xec\xd7\x39\x8b\xc4\xe6\x83\xaa\x7f\x7c\x04\x79\x60\xfe\x64\x08\x58\x6b\x96\xdf\xb6\x62\x4e\xd6\xd5\x99\x5d\xde\xb8\x67\xd2\xe4\x3a\xa5\x72\x19\xd4\x01\x1b\xd5\xfd\x1a\xd5\xe1\x59\xc5\x5a\x6b\x55\x56\x30\x2d\xf4\x08\x43\x60\x5d\xf3\xf5\xe4\x95\x61\xab\x8c\x08\x4b\xde\x2b\x30\x7c\x19\x99\xd1\x01\xf2\x99\xce\x1b\x67\xaa\x75\x7a\xee\xf0\x60\x4b\xfc\x27\xa9\x92\xac\x55\xd9\xc2\x84\xea\xde\xf5\xcb\xf3\xc8\x94\xc5\x36\xf1\xb1\x55\x55\x37\xfb\xe0\xab\x7c\xdf\x8b\x45\xf4\xf2\x5f\xec\x81\x4c\x60\xd1\xac\x8a\xd3\x62\x75\x75\x7d\xf3\x19\x67\x89\x62\x66\x60\x88\xc9\x52\xfb\x6f\x6b\xb1\x21\x7c\xd0\x3a\x6d\xb5\x63\x92\x55\xb6\xab\xb3\x1d\xf4\x4d\x4c\x80\x21\xf8\xef\xa8\x4e\x17\xbf\x41\x4e\xce\x5a\x8b\xc1\x62\x7b\x70\x04\xec\x9a\xd4\x75\x84\x6d\x74\x62\x1e\xec\x6d\xdd\x9e\xe8\xe7\x07\xeb\x88\x13\xab\x06\x67\x67\x83\xbd\x18\x1f\x19\x1c\xfb\x6d\x62\x70\x93\xaa\x59\x6c\x6b\x7a\x88\xde\xa1\x30\x36\x54\xb9\xac\xa3\xba\x63\x30\x53\x58\xcc\x22\x76\x62\x98\xad\x8a\x8e\xef\x90\xdb\x7c\x3d\xab\x53\x30\xb9\x1c\xab\xba\x1d\xc4\x8c\xb1\xc2\xca\xb1\x08\xbf\xa2\x22\x22\x46\xad\x5b\x1b\x3d\x16\x8d\x4d\x27\x67\x50\x93\xce\x3f\x8b\x0d\x72\x1e\x74\xbf\xfb\x05\x98\xc9\xf3\xc8\x0c\x3e\xe1\xdb\xf8\x1b\xfb\x58\xd7\x6f\x71\x79\xd1\x37\xf8\xc3\x3a\xf8\x32\x3a\x68\x55\x09\x74\xea\x72\x59\x9c\x55\xae\x3f\xeb\x34\x4e\xce\xe6\x2b\x61\x4c\x52\x87\x94\x4a\x7b\x50\x3c\x83\x37\x7b\xb5\xc5\x3f\xce\x08\xc3\x45\x16\x18\xe6\xb4\x57\x04\x8d\xc7\xe7\x26\xbc\xf8\x9a\x45\x70\xd6\xac\xce\xdf\xee\x93\xd5\x0b\x4c\xeb\xcd\x19\xf8\xcf\xe7\x10\x5f\xb0\xb1\x37\x30\x88\x17\xf9\x59\xfb\xf2\x73\x4f\x0f\xf9\xf4\x1d\x99\xd9\x9b\xa4\x7a\xa7\x3e\xdf\xb4\xa6\xe9\xf2\xb7\x7e\xf9\x1a\x0f\xfb\xcd\xd5\xbf\x81\x81\xe1\x75\xb3\xbc\x70\x50\x0f\xb8\x15\x92\x58\x0f\x75\xf3\xbe\x0b\xd9\x1f\x64\x02\x1d\xac\xa7\x57\x7d\xd1\xb7\x19\x85\x22\xc9\xeb\x93\x41\x3d\x08\x63\x6d\xbd\xaf\x33\xa7\xe7\xfe\x64\xaf\x1e\x17\x08\x01\xca\x35\x62\x40\xb0\x26\x32\x0c\xa8\x6c\xfd\x27\xf5\x1c\xc9\x76\xa0\xe6\xa6\x1e\x36\x19\x74\x53\x5f\xa7\x57\xdd\x59\x29\xf2\x40\x0e\x90\x9b\x48\x1c\xd4\x47\x8b\xb5\xeb\x01\x39\x91\x79\xef\x41\x80\x40\x9c\xad\x53\xaf\x90\x0c\x60\x3d\xc4\xcd\xff\x9e\x2a\xf5\xeb\x93\xef\x63\xb2\x56\xea\x31\xb5\x8e\x68\xd6\xff\x18\x5c\x66\xde\x83\xb2\x64\xc1\x9d\xcd\x5a\xef\x45\xcc\xd1\x32\xae\x66\x33\xe8\x62\x12\xd3\x4d\xea\x5f\x82\xca\x96\x19\x26\x47\x98\xa8\x1e\x88\xf5\x0d\x3b\x67\x39\xa0\xa1\xf5\x86\xd4\x4b\x69\x55\xfb\x8e\x9a\x27\xf0\x19\xa8\xc5\x7e\xd6\x27\xeb\xd5\x0f\x15\xdb\xcc\x49\xb3\x97\xe8\x35\x26\x28\x4a\xfd\xcb\x39\x46\xf5\xe8\xda\x49\x4c\x7c\x16\xda\xb5\x6e\xbf\x56\x6c\x24\xc9\xce\xd8\x9b\x75\x60\xef\x36\x17\x4a\xce\xae\xac\x67\xdd\x28\x7b\xac\xce\x1c\x58\x1f\x06\xc2\x73\x73\x7a\xee\xe1\x61\xc3\x55\xf3\x94\x24\xe6\xb9\x0c\x8e\xaa\xe8\x60\x91\x9f\x99\xad\x84\xf4\x26\xfb\xac\xbe\xe0\xe8\xf2\xc3\x1e\x9a\x2f\x98\xa8\xf5\xcc\x3a\xcd\x1f\x7a\x47\x63\x62\xc2\xe6\x73\x8b\x7a\xbe\xa3\xeb\x87\xcf\xad\x87\xdb\xa9\x37\xb8\x68\xc6\xa8\xfa\x38\x2b\xbb\xa7\x57\x50\x27\xa9\xd7\x6a\x33\xaf\xe8\xb2\x3e\xff\xee\x8d\xa8\xf3\x87\x30\xfc\x0d\x2b\xfe\x3e\xe2\xbc\x88\xdb\xdd\x6b\x78\x7f\xfb\x53\x3e\x3d\xf8\xc5\x86\xef\x87\xdf\xef\x92\xee\x9d\x3a\x7d\xd7\xda\xa6\xdd\xb7\x3c\xf1\x1a\xae\xbf\x69\xb7\x3f\xa3\xc2\xae\x35\x57\x11\xd3\x83\x69\xee\x15\xf6\xa2\xf9\x3b\xb1\x6e\xac\x49\x33\x3f\xf0\xde\xfa\xf9\x83\x63\x80\xf5\x96\x56\x8f\x13\xaa\xb3\x49\xac\x6e\xd1\x3c\xd6\x2a\xb9\xce\xe3\x10\x36\x37\x88\x6d\x47\x55\xda\x83\x30\xde\xaa\x66\xb1\x4a\x9b\x2f\x56\xfe\x1e\xd9\xac\x47\xb9\x6a\xee\x9d\x54\x01\x0a\xeb\xf9\x1b\x46\x65\x4c\x7e\x70\x7c\xb1\x39\x9d\x2a\x4d\xe2\x9c\x38\xee\x93\xaa\x87\xc9\xe7\x54\xd6\x57\xaf\x1c\xcf\x6c\xce\xd4\xfb\x6f\x6c\x36\xab\xf9\x08\xeb\xf6\xea\x9b\xc3\xcc\xd1\x1f\x38\x50\x2b\x97\xd4\x3a\x6f\xa3\x7e\x68\x6a\x3d\x87\x75\xca\x09\x8b\xaa\x6b\x2a\x35\x18\x66\x53\x6b\xfe\x53\xfb\xbe\xc8\xb9\x57\xf5\x9d\xe6\x71\x60\x29\x4c\x14\xec\x5b\x55\x69\x80\xf7\xd6\x2f\x9d\x55\x61\x2d\x62\xac\x8b\x7f\xd6\xac\x9e\x33\xd0\x37\x39\x60\xd9\x7b\xea\x9a\xc1\x0e\x9a\x39\x62\x3b\x9b\x2b\xaa\xe7\x6b\xff\x72\x28\xb3\x2a\xb3\x4d\x98\x8f\x8e\x6c\x5e\x9a\x74\x47\x23\xf9\xcc\x8a\x7c\x06\x3e\x47\xad\x3f\xab\xff\x68\xf3\xa0\x5e\x5d\x8a\xc5\xf3\x27\xf9\x95\x1c\x40\xa5\x89\xbc\xab\x7a\xe6\xe4\xfe\x4d\xfd\x59\xe4\x31\x3f\x53\xe7\x86\xb3\x58\xcf\x5b\x95\x6f\x5c\x9f\xe5\x83\xd1\xcf\xda\xc7\xfc\x63\xd0\xac\x55\xb3\x10\xcb\xad\xf2\x59\x38\xc4\xd8\xb8\x3c\xe4\x03\xeb\xf4\x68\x46\xd9\xaa\x22\x40\x26\x64\x6d\x34\x23\xc6\xfe\x93\x2a\xfd\x41\x95\xe1\xff\xcd\x4e\x56\xf7\x11\xaa\xfc\xa4\x4a\x3e\xab\x0f\x8c\x5d\xb0\xbb\x71\x0e\xf5\xab\x59\xd3\xb8\x8f\xfc\x0f\x5b\xce\xba\xf3\x61\x3d\x54\xe5\x4b\xf2\x07\x76\xb1\x8a\x33\x3b\x9f\x32\x1f\xd3\x5d\x0f\xec\x4b\x3c\xd7\x9a\xf3\x53\x91\xda\x1c\x2f\x3a\xb7\xb1\x1e\xf1\xea\xfa\x9c\x14\x8b\x56\xe1\x0e\xce\x95\xe0\x0d\xdd\xde\x07\x9f\x3d\x8e\xe1\x15\xac\x43\xb5\x8a\x6f\xcc\xe2\x6a\x76\x77\x66\x74\x1e\x62\x33\x3e\xf5\xea\xc1\x14\xfc\xa5\xdb\x2b\x4e\xc5\x05\x31\x0b\x17\x89\xe2\x8b\xc6\xc3\x5a\xcf\xcf\xe4\x46\xf4\xce\xb9\xf0\x5f\xab\xf0\xa2\xfb\x85\xcd\x6a\x26\x8f\x3f\x93\xa9\x7b\xee\x54\x64\xcd\x4b\xd0\x8b\xe5\xf6\xfe\x8b\xcf\x16\x7d\xb6\xe9\xfc\x8d\xfb\x7c\xa5\x19\x61\xa5\x3b\x40\xc4\xce\xac\x3b\x43\xec\x3d\x88\x47\xcd\x9a\xf1\xb0\x1f\xdf\x75\xd3\xf3\x9c\x3e\xe9\x8e\x06\xba\xc5\x97\x6d\xad\x55\x5c\x6d\xd2\x1d\x9c\xc9\x3b\x0a\x36\x67\x51\x65\x08\x1f\xb6\xbb\x12\x9b\xf4\xa3\xee\x05\x32\xed\xbe\xdf\xbf\x92\xd5\xf0\x69\x7e\x83\x9f\x0e\xe2\x71\xfb\x9d\x98\x77\x67\xb5\x3f\xc4\x13\xde\xb1\xf2\xdb\x33\xda\x8b\xbc\xa1\xbf\xba\x8f\x3f\xdf\xe4\xdb\xc7\xab\xc7\xb8\x7c\xce\xdf\x8f\x22\xbc\x26\xc8\x3b\x35\xf7\xda\x32\xa6\xa6\xaf\xbe\x7c\x2d\xc7\x7f\xbd\xc6\x37\xa4\x73\x1b\x2d\x77\x1a\x4f\xaa\x29\xb3\x37\xcc\x47\x35\x6b\x7b\x8d\x87\xf8\x1e\xb7\xb3\x86\x6c\xef\x50\x90\x04\x2d\xb3\xc6\x3a\x83\xae\x8a\x54\xbd\x97\x07\xa4\x33\xa3\xe2\x4a\x97\x93\x46\x4e\x76\xbd\xa7\x56\xc3\x2e\xab\xa9\xa2\x66\xde\xaa\xe6\xdc\xac\x2b\x25\x40\x2c\xb2\xb4\xba\x9a\x07\x5c\x00\xd9\xd6\xf0\x9e\x1d\x46\x91\xaf\xde\x9b\xbd\xa3\x8f\xf2\xb2\xd2\x3c\x65\xc6\xa6\xab\x2b\x76\xed\x69\xf6\x74\x97\x75\xbd\xc1\xae\x26\x55\x0e\x2d\x50\xfc\xfd\x0a\xd1\xa8\xab\x3a\xfb\x55\x33\x52\x8a\xa5\xef\xde\x21\x6f\x2f\x15\x29\x4d\xf6\x34\x3b\x69\x04\x69\x57\xf1\xb4\x1e\x3a\xed\xd4\xf8\xef\xd4\xe4\xe1\x75\xb7\x0d\x34\xc0\xca\x8e\xde\xa1\xbf\xd1\x6b\xbd\x53\xad\xd6\x1b\x57\x8b\xae\x3b\x02\xbf\xc8\x3f\x0a\xca\x81\x38\x20\x81\x73\x01\x31\x40\xf8\x7e\x9d\xb0\x55\xf3\x14\xa8\x27\x5d\x8c\xa3\x46\x8a\x93\xc3\x26\xe7\xb7\x72\xa3\xd6\xf5\xc8\xde\x53\x69\xab\x35\x7a\x35\x03\xad\xe1\x1b\xd5\x98\x14\xa4\xee\x57\x8a\x66\x0d\x3c\xac\x14\xea\x34\xe2\x53\xa9\x5b\xab\xe4\x02\x3a\xa3\xd2\x14\x29\x9a\xd7\x58\x39\x55\x58\x75\x0d\x70\xd6\x88\x9d\x92\x8f\xf4\x88\x1e\x1a\x8d\xbe\x57\x51\x41\x6b\x40\x8b\x2e\xe2\x47\x8d\xae\xe0\xec\x2d\x06\x74\x6a\x74\x53\xd7\x85\x6c\x58\x23\x4a\x65\x57\xd8\x1a\x6f\x48\xb7\xba\x2a\x51\xe9\xea\xd9\xa0\xab\x27\xb3\xae\xd8\x4d\x6a\x0d\xa0\x3b\x2b\xfd\x67\x5f\x13\x79\x6b\xd9\x2c\x8b\x6a\xa1\x17\xd2\xa0\x5d\xa3\x6a\x3d\x5e\xf0\x07\x83\xf6\xc1\xcf\xc5\x7e\xd5\x5e\x5a\x6b\x44\xbd\x8f\x65\xed\xfa\x92\xfc\x70\xd9\xaf\x48\xa9\xcc\x34\xba\x9d\x5c\x06\xa8\x0b\x69\x02\xfa\x8b\x1d\x2d\x4d\x6b\xb8\xb1\x68\x1c\xbb\x28\x66\xad\xa1\xdb\x3d\x37\x2f\xf7\xab\x12\x16\x33\xa2\xa1\x95\x46\xbc\xc8\x14\x45\x2f\xa6\xe8\x7e\x68\x57\x70\x06\x5d\x13\x6a\xd5\xa4\x55\x0b\x61\x10\xd5\xb7\x6b\x3d\x4a\x63\xdd\xec\x7a\xa3\x0c\xdf\x34\x2c\x98\x3b\xa5\x29\x5d\xc3\xc3\xdf\xa2\xfc\xb8\xde\xd3\xbc\x6c\x9d\x75\x7d\x6f\xd4\xf5\x84\x59\x29\x90\x18\xee\xa5\x87\xa4\xd4\xd8\xeb\x0a\x0d\xba\xb3\xab\x88\xba\x4a\x81\xac\x51\xcd\xe3\xa8\xf2\xde\xf6\x59\xdc\xf7\x2d\xa6\x16\x5d\x0f\x5e\xdd\x0f\x78\xd6\x86\x2c\xad\x86\x58\x6a\xf1\x4c\xba\xce\x51\xe9\xef\xfd\x0a\xc7\xa2\x01\x14\xfb\xda\xf5\xaf\x51\x03\xa9\xe4\x71\xb4\x8a\x4a\xe5\xc1\xe3\x76\x12\x2d\xe2\x79\x7e\xb7\x8a\x62\x2c\x6a\x7b\xad\xc2\x0c\x7c\x07\xd9\xed\x0a\x6b\xfd\x7c\xbd\x68\x5b\xfe\x37\x00\x00\xff\xff\xcf\x49\xdf\xc9\x00\x30\x00\x00")

func bindata_go() ([]byte, error) {
	return bindata_read(
		_bindata_go,
		"bindata.go",
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
	"5_payment_table.down.sql": _5_payment_table_down_sql,
	"5_payment_table.up.sql": _5_payment_table_up_sql,
	"bindata.go": bindata_go,
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
	"5_payment_table.down.sql": &_bintree_t{_5_payment_table_down_sql, map[string]*_bintree_t{
	}},
	"5_payment_table.up.sql": &_bintree_t{_5_payment_table_up_sql, map[string]*_bintree_t{
	}},
	"bindata.go": &_bintree_t{bindata_go, map[string]*_bintree_t{
	}},
}}
