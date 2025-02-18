# Generate pb

## 1. Install protoc
Download the latest version of protoc from the official website: https://github.com/protocolbuffers/protobuf/releases
## 2. Create protobuf/internal/xxx/entity
```
mkdir -p protobuf/app/go/entity
mkdir -p protobuf/app/pytho/entity
```
## 3. Generate go pb
```
protoc \
  --go_out=/protobuf \
  --proto_path=/protobuf/proto \
  --go_opt=module=protobuf \
  --go_opt=Mgame.proto=protobuf/app/go/entity \
  /protobuf/proto/game.proto 
```
## 4. Generate python pb
### generate
```
 protoc \
  --python_out=/protobuf/app/python/entity \
  --proto_path=/protobuf/proto \
  /protobuf/proto/game.proto
```
### currently both proto2 and proto3 use _pb2.py for their generated files
https://protobuf.dev/reference/python/python-generated/