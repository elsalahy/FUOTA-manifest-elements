#include <stdio.h>

struct manifest
{
u_int8_t ManifestId;
u_int8_t ManifestVersion;
u_int8_t NumberTargetHardware;
u_int32_t TargetHardwareVersions[2];
u_int32_t FirmwareVersion;
u_int8_t FirmwareImageType;
u_int8_t DiffType;
u_int32_t PreviousImageSize;
u_int32_t TargetFirmwareVersion;
u_int8_t PreviousFirmwareSignatureSize;
u_int8_t PreviousFirmwareSignatureValue[10];
u_int8_t FirmwareDataType;
u_int8_t CompressionType;
u_int8_t HashAlgorithmId;
u_int8_t SignatureAlgorithmId;
u_int8_t SignatureSize;
u_int8_t SignatureValue[10];
u_int8_t AppInfoSize;
u_int8_t AppInfoBytes[5];
};

int main()
{
    FILE *f;
    struct manifest r;
    f=fopen("example1.bin","r");
    if (!f){
        return 1;
    }
    fread(&r,sizeof(struct manifest),1,f);
    printf("%d\n",r.ManifestId);
    printf("%d\n",r.ManifestVersion);
    fclose(f);
    printf("\n");
    return 0;
}