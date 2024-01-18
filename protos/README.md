## Protos
For protobuf based API development. Buf will take care of protobuf generation for each language.
  - Buf modules
  - Buf plugins

### Basic
  1. buf mod init
  2. buf mod update
  3. buf build
  4. buf generate proto
  5. buf lint proto
  6. buf curl --schema proto http://localhost:8080/missingstudio.v1.PingService/Ping