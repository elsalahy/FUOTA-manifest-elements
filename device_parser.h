// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#ifndef DEVICE_PARSER
#define DEVICE_PARSER

#include <stdio.h>

#define TARGET_HARDWARE_COUNT 0x02U
#define RANDOM_DATA "canbeanythingmatchingtheabovesize"

typedef struct __attribute__((__packed__))
{
    u_int8_t ManifestId;
    u_int8_t ManifestVersion;
    u_int8_t TargetHardwareCount;
    u_int32_t TargetHardwareUID[TARGET_HARDWARE_COUNT];
    u_int32_t FirmwareVersion;
    u_int8_t FirmwareImageType;
    u_int8_t FirmwareDataType;
    u_int8_t HashAlgorithmId;
    u_int8_t SignatureAlgorithmId;
    u_int8_t SignatureSize;
    u_int8_t SignatureValue[sizeof(RANDOM_DATA)];
} fuota_manifest_basic;

typedef struct __attribute__((__packed__))
{
    u_int8_t ManifestId;
    u_int8_t ManifestVersion;
    u_int8_t TargetHardwareCount;
    u_int32_t TargetHardwareUID[TARGET_HARDWARE_COUNT];
    u_int32_t FirmwareVersion;
    u_int8_t FirmwareImageType;
    u_int8_t DiffType;
    u_int32_t PreviousImageSize;
    u_int32_t TargetFirmwareVersion;
    u_int8_t PreviousFirmwareSignatureSize;
    u_int8_t PreviousFirmwareSignatureValue[sizeof(RANDOM_DATA)];
    u_int8_t FirmwareDataType;
    u_int8_t CompressionType;
    u_int8_t HashAlgorithmId;
    u_int8_t SignatureAlgorithmId;
    u_int8_t SignatureSize;
    u_int8_t SignatureValue[sizeof(RANDOM_DATA)];
    u_int8_t AppInfoSize;
    u_int8_t AppInfoBytes[sizeof(RANDOM_DATA)];
} fuota_manifest_extended;

#endif