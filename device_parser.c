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
#include "device_parser.h"

static void print_parsed_basic( fuota_manifest_basic *parsed_structure)
{
    printf(" Structure size: %lu bytes \r\n", sizeof(fuota_manifest_basic));
    printf("ManifestId: %d \r\n", parsed_structure->ManifestId);
    printf("ManifestVersion: %d \r\n", parsed_structure->ManifestVersion);
    printf("TargetHardwareCount: %d \r\n", parsed_structure->TargetHardwareCount);
    for (size_t i = 0; i < TARGET_HARDWARE_COUNT; i++)
    {
        printf("TargetHardwareUID[%zu] %d \r\n", i, parsed_structure->TargetHardwareUID[i]);
    }
    printf("FirmwareVersion: %d \r\n", parsed_structure->FirmwareVersion);
    printf("FirmwareImageType: %d \r\n", parsed_structure->FirmwareImageType);
    printf("FirmwareDataType: %d \r\n", parsed_structure->FirmwareDataType);
    printf("HashAlgorithmId: %d \r\n", parsed_structure->HashAlgorithmId);
    printf("SignatureAlgorithmId: %d \r\n", parsed_structure->SignatureAlgorithmId);
    printf("SignatureSize: %d \r\n", parsed_structure->SignatureSize);
    printf("SignatureValue: %s \r\n", parsed_structure->SignatureValue);

    return;
}


int main()
{
    FILE *f;
    fuota_manifest_basic r;
    f = fopen("example1.bin", "r");
    if (!f)
    {
        return 1;
    }
    printf("parsing example1.bin \r\n");
    fread(&r, sizeof(fuota_manifest_basic), 1, f);
    print_parsed_basic(&r);
    fclose(f);

    return 0;
}