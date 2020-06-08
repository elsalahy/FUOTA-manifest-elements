# FUOTA-manifest-elements

> This repository is a work in progress, and currently it should not be used for production purposes.

## Introduction
This repo represents an effort to develop a standard manifest file for LoRaWAN Firmware Updates Over the Air (FUOTA), building on the latest LoRa Alliance specs and recommendations.

## Usage
- The manifest schema is written in `JSON` and input examples are validated in `Go`.

- The binary is generated using `Go` if the input example `1` or `2` is valid.

- The `device_parser.c` simulates a device parsing the binary manifest.

To use the validation schema, run this command once:

```
go get github.com/xeipuuv/gojsonschema
```
and then to validate the examples:

```
go run *.go
```

For the C parser, just compile:

```
gcc device_parser.c -o parser
```

and run it
```
./parser
```

Feel free to edit and test various examples.
