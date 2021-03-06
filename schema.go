package main

import (
	"encoding/binary"
	"os"
)

type manifest struct {
	ManifestId                     uint8    `json:"manifestId"`
	ManifestVersion                uint8    `json:"manifestVersion"`
	TargetHardwareCount            uint8    `json:"targetHardwareCount"`
	TargetHardwareUID              []uint32 `json:"targetHardwareUID"`
	FirmwareVersion                uint32   `json:"firmwareVersion"`
	FirmwareImageType              string   `json:"firmwareImageType"`
	DiffType                       string   `json:"diffType"`
	PreviousImageSize              uint32   `json:"previousImageSize"`
	TargetFirmwareVersion          uint32   `json:"targetFirmwareVersion"`
	PreviousFirmwareSignatureSize  uint8    `json:"previousFirmwareSignatureSize"`
	PreviousFirmwareSignatureValue string   `json:"previousFirmwareSignatureValue"`
	FirmwareDataType               string   `json:"firmwareDataType"`
	CompressionType                string   `json:"compressionType"`
	HashAlgorithmId                string   `json:"hashAlgorithmId"`
	SignatureAlgorithmId           string   `json:"signatureAlgorithmId"`
	SignatureSize                  uint8    `json:"signatureSize"`
	SignatureValue                 string   `json:"signatureValue"`
	AppInfoSize                    uint8    `json:"appInfoSize"`
	AppInfoBytes                   string   `json:"appInfoBytes"`
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
	binary.Write(file, binary.LittleEndian, input.ManifestId)
	binary.Write(file, binary.LittleEndian, input.ManifestVersion)
	binary.Write(file, binary.LittleEndian, input.TargetHardwareCount)
	for _, u := range input.TargetHardwareUID {
		binary.Write(file, binary.LittleEndian, u)
	}
	binary.Write(file, binary.LittleEndian, input.FirmwareVersion)
	binary.Write(file, binary.LittleEndian, enumToValue[input.FirmwareImageType])
	if input.DiffType != "" {
		binary.Write(file, binary.LittleEndian, enumToValue[input.DiffType])
	}
	if input.TargetFirmwareVersion != 0 {
		binary.Write(file, binary.LittleEndian, input.TargetFirmwareVersion)
	}
	if input.PreviousImageSize != 0 {
		binary.Write(file, binary.LittleEndian, input.PreviousImageSize)
	}
	if input.PreviousFirmwareSignatureSize != 0 {
		if input.PreviousFirmwareSignatureValue != "" {
			binary.Write(file, binary.LittleEndian, input.PreviousFirmwareSignatureSize)
			binary.Write(file, binary.LittleEndian, []byte(input.PreviousFirmwareSignatureValue))
		}
	}
	binary.Write(file, binary.LittleEndian, enumToValue[input.FirmwareDataType])
	if input.CompressionType != "" {
		binary.Write(file, binary.LittleEndian, enumToValue[input.CompressionType])
	}
	binary.Write(file, binary.LittleEndian, enumToValue[input.HashAlgorithmId])
	binary.Write(file, binary.LittleEndian, enumToValue[input.SignatureAlgorithmId])
	binary.Write(file, binary.LittleEndian, input.SignatureSize)
	binary.Write(file, binary.LittleEndian, []byte(input.SignatureValue))
	if input.AppInfoSize != 0 {
		if input.AppInfoBytes != "" {
			binary.Write(file, binary.LittleEndian, input.AppInfoSize)
			binary.Write(file, binary.LittleEndian, []byte(input.AppInfoBytes))
		}
	}
}
