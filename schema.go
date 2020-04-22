package main

import (
	"encoding/binary"
	"os"
)

type manifest struct {
	Manifest_id                         uint8    `json:"manifest_id"`
	Manifest_version                    uint8    `json:"manifest_version"`
	Number_target_hardware              uint8    `json:"number_target_hardware"`
	Target_hardware_versions            []uint32 `json:"target_hardware_versions"`
	Firmware_version                    uint32   `json:"firmware_version"`
	Firmware_image_type                 string   `json:"firmware_image_type"`
	Diff_type                           string   `json:"diff_type"`
	Previous_image_size                 uint32   `json:"previous_image_size"`
	Target_firmware_version             uint32   `json:"target_firmware_version"`
	Previous_firmware_signature_size    uint8    `json:"previous_firmware_signature_size"`
	Previous_firmware_payload_signature string   `json:"previous_firmware_payload_signature"`
	Firmware_data_type                  string   `json:"firmware_data_type"`
	Compression_type                    string   `json:"compression_type"`
	Hash_algorithm_id                   string   `json:"hash_algorithm_id"`
	Signature_algorithm_id              string   `json:"signature_algorithm_id"`
	Signature_size                      uint8    `json:"signature_size"`
	Payload_signature                   string   `json:"payload_signature"`
	App_info_size                       uint8    `json:"app_info_size"`
	App_info_bytes                      string   `json:"app_info_bytes"`
}

type manifestEnum uint8

const (
	full  manifestEnum = iota
	delta manifestEnum = iota
)

const (
	compressed   manifestEnum = iota
	uncompressed manifestEnum = iota
)

const (
	MD5     manifestEnum = iota
	SHA_256 manifestEnum = iota
)

const (
	DSA_PSS manifestEnum = iota
	ECDSA   manifestEnum = iota
	HMAC    manifestEnum = iota
	AES_CBC manifestEnum = iota
)

const (
	gzip    manifestEnum = iota
	bzip2   manifestEnum = iota
	deflate manifestEnum = iota
	lz4     manifestEnum = iota
	lzma    manifestEnum = iota
)
const (
	jojodiff manifestEnum = iota
	RMTD     manifestEnum = iota
	R3diff   manifestEnum = iota
	MoRE     manifestEnum = iota
)

var enumToValue = map[string]manifestEnum{
	"full":         full,
	"delta":        delta,
	"uncompressed": uncompressed,
	"compressed":   compressed,
	"MD5":          MD5,
	"SHA_256":      SHA_256,
	"DSA_PSS":      DSA_PSS,
	"ECDSA":        ECDSA,
	"HMAC":         HMAC,
	"AES_CBC":      AES_CBC,
	"gzip":         gzip,
	"bzip2":        bzip2,
	"deflate":      deflate,
	"lz4":          lz4,
	"lzma":         lzma,
	"jojodiff":     jojodiff,
	"RMTD":         RMTD,
	"R3diff":       R3diff,
	"MoRE":         MoRE,
}

// Todo improve the function to avoid processing each element alone
func writeBinary(input *manifest, file *os.File) {
	binary.Write(file, binary.BigEndian, input.Manifest_id)
	binary.Write(file, binary.BigEndian, input.Manifest_version)
	binary.Write(file, binary.BigEndian, input.Number_target_hardware)
	for _, u := range input.Target_hardware_versions {
		binary.Write(file, binary.BigEndian, u)
	}
	binary.Write(file, binary.BigEndian, input.Firmware_version)
	binary.Write(file, binary.BigEndian, enumToValue[input.Firmware_image_type])
	if input.Diff_type != "" {
		binary.Write(file, binary.BigEndian, enumToValue[input.Diff_type])
	}
	if input.Target_firmware_version != 0 {
		binary.Write(file, binary.BigEndian, input.Target_firmware_version)
	}
	if input.Previous_image_size != 0 {
		binary.Write(file, binary.BigEndian, input.Previous_image_size)
	}
	if input.Previous_firmware_signature_size != 0 {
		if input.Previous_firmware_payload_signature != "" {
			binary.Write(file, binary.BigEndian, input.Previous_firmware_signature_size)
			binary.Write(file, binary.BigEndian, []byte(input.Previous_firmware_payload_signature))
		}
	}
	binary.Write(file, binary.BigEndian, enumToValue[input.Firmware_data_type])
	if input.Compression_type != "" {
		binary.Write(file, binary.BigEndian, enumToValue[input.Compression_type])
	}
	binary.Write(file, binary.BigEndian, enumToValue[input.Hash_algorithm_id])
	binary.Write(file, binary.BigEndian, enumToValue[input.Signature_algorithm_id])
	binary.Write(file, binary.BigEndian, input.Signature_size)
	binary.Write(file, binary.BigEndian, []byte(input.Payload_signature))
	if input.App_info_size != 0 {
		if input.App_info_bytes != "" {
			binary.Write(file, binary.BigEndian, input.App_info_size)
			binary.Write(file, binary.BigEndian, []byte(input.App_info_bytes))
		}
	}
}
