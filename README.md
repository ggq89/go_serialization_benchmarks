# Benchmarks of Go serialization methods

[![Gitter chat](https://badges.gitter.im/alecthomas.png)](https://gitter.im/alecthomas/Lobby)

This is a test suite for benchmarking various Go serialization methods.

## Serialization Libraries Tested

- [encoding/json](https://pkg.go.dev/encoding/json) - Go standard library JSON
- [encoding/gob](https://pkg.go.dev/encoding/gob) - Go standard library Gob
- [gotiny](https://github.com/niubaoshu/gotiny) - Reflect-based binary serialization
- [msgp](https://github.com/tinylib/msgp) - Code-generated MessagePack
- [msgpack](https://github.com/vmihailenco/msgpack) - Reflect-based MessagePack
- [jsoniter](https://github.com/json-iterator/go) - Fast drop-in JSON replacement
- [easyjson](https://github.com/mailru/easyjson) - Code-generated JSON
- [bson](https://gopkg.in/mgo.v2/bson) - MongoDB BSON (mgo)
- [mongobson](https://go.mongodb.org/mongo-driver/mongo) - MongoDB BSON (official driver)
- [davecgh/xdr](https://github.com/davecgh/go-xdr) - XDR serialization
- [calmh/xdr](https://github.com/calmh/xdr) - XDR with code generation
- [ugorji/codec](https://github.com/ugorji/go/tree/master/codec) - MsgPack & Binc codec
- [sereal](https://github.com/Sereal/Sereal/tree/master/Go/sereal) - Sereal serialization
- [alecthomas/binary](https://github.com/alecthomas/binary) - Binary serialization
- [flatbuffers](https://github.com/google/flatbuffers/tree/master/go) - Google FlatBuffers
- [capnproto](https://github.com/glycerine/go-capnproto) - Cap'n Proto (legacy)
- [capnproto v3](https://github.com/capnproto/go-capnp) - Cap'n Proto v3
- [hprose](https://github.com/hprose/hprose-go) - Hprose v1
- [hprose2](https://github.com/hprose/hprose-golang) - Hprose v2
- [dedis/protobuf](https://go.dedis.ch/protobuf) - DEDIS Protobuf (reflect)
- [pulsar](https://github.com/cosmos/cosmos-proto) - Cosmos Protobuf (Pulsar)
- [protobuf-go](https://github.com/protocolbuffers/protobuf-go) - Google Protocol Buffers
- [gogo/protobuf](https://github.com/gogo/protobuf) - GoGo Protobuf
- [colfer](https://github.com/pascaldekloe/colfer) - Colfer serialization
- [gencode](https://github.com/andyleap/gencode) - Code-generated binary
- [goavro](https://gopkg.in/linkedin/goavro.v1) - Apache Avro v1
- [goavro v2](https://github.com/linkedin/goavro) - Apache Avro v2
- [ikea](https://github.com/ikkerens/ikeapack) - IkeaPack
- [shamaton/msgpack](https://github.com/shamaton/msgpack) - Shamaton MessagePack
- [ssz](https://github.com/prysmaticlabs/go-ssz) - Simple Serialize (Ethereum 2.0)
- [200sc/bebop](https://github.com/200sc/bebop) - Bebop serialization
- [wellquite/bebop](https://wellquite.org/bebop) - Bebop (alternative)
- [fastjson](https://github.com/valyala/fastjson) - Fast JSON manipulation
- [benc](https://github.com/deneonet/benc) - Binary encoding
- [mus](https://github.com/mus-format/mus-go) - MUS format
- [idr](https://github.com/chmike/ditp) - IDR encoding
- [fastape](https://github.com/nazarifard/fastape) - High-performance serialization
- [SBE](https://github.com/aeron-io/simple-binary-encoding) - Simple Binary Encoding
- [fory](https://github.com/apache/fory) - Apache Fory
- [thrift](https://github.com/apache/thrift/tree/master/lib/go) - Apache Thrift

# Current Serialization Results

https://alecthomas.github.io/go_serialization_benchmarks

## Running the benchmarks

To benchmark and validate, without cloning the repository, replace the `.` from the commands below with `github.com/alecthomas/go_serialization_benchmarks@latest`.

```bash
go run .
```

To validate the correctness of the serializers:
```bash
go run . --validate
```

To update the benchmark report:
```bash
go run . --genreport
```

To update the benchmark report with a longer benchmark run (to get more accurate results):
```bash
go test -tags genreport -run TestGenerateReport -benchtime 10s -timeout 1h #--validate
```

## Recommendation

If correctness and interoperability are the most
important factors [JSON](http://golang.org/pkg/encoding/json/) or [Protobuf](https://google.golang.org/protobuf) are your best options.

But as always, make your own choice based on your requirements.

## Adding New Serializers

Review the following instructions _before_ opening the PR to add a new
serializer:

- Create all the required serializer code in `internal/<short serializer name>`.
- Add an entry to the serializer in [benchmarks.go](benchmarks.go).
- If the serializer supports both reusing/not reusing its marshalling buffer:
  - Add both a `serializer` and `serializer/reuse` entries, each one
    respectively reusing/not reusing the resulting marshalling buffer. Set the
    `BufferReuseMarshal` flag accordingly.
- If the serializer supports both safe and unsafe string unmarshalling:
  - Add both a `serializer` and `serializer/unsafe` entries, each one
    respectively unmarshalling into safe and unsafe strings. Set the
    `UnsafeStringUnmarshal` flag accordingly.
- If the serializer supports both marshalling buffer reuse and unsafe string
  unmarshalling, merge both options into a single `serializer/unsafe_reuse`
  entry (check the baseline serializer for an example).
- Regenerate the report by running:

```
go run . --genreport
```

- **Include the updated report data in your PR**

## Data

The data being serialized is the following structure with randomly generated values:

```go
type A struct {
    Name     string
    BirthDay time.Time
    Phone    string
    Siblings int
    Spouse   bool
    Money    float64
}
```

