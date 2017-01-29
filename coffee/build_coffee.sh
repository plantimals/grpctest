#!/bin/bash
protoc --go_out=plugins=grpc:. *.proto

#protoc --go_out=plugins=grpc,import_path=coffee:. coffee/*.proto
