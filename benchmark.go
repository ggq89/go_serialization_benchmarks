// This file is used to generate the report and is ignored by default on tests.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type reportLine struct {
	Name                  string `json:"name"`
	MarshalIterCount      int    `json:"marshal_iter_count"`
	UnmarshalIterCount    int    `json:"unmarshal_iter_count"`
	TotalIterCount        int    `json:"total_iter_count"`
	UnsafeStringUnmarshal bool   `json:"unsafe_string_unmarshal"`
	BufferReuseMarshal    bool   `json:"buffer_reuse_marshal"`
	Deterministic         bool   `json:"deterministic"`
	MarshalNsOp           int64  `json:"marshal_ns_op"`
	UnmarshalNsOp         int64  `json:"unmarshal_ns_op"`
	TotalNsOp             int64  `json:"total_ns_op"`
	SerializationSize     int64  `json:"serialization_size"`
	MarshalAllocBytes     int64  `json:"marshal_alloc_bytes"`
	UnmarshalAllocBytes   int64  `json:"unmarshal_alloc_bytes"`
	TotalAllocBytes       int64  `json:"total_alloc_bytes"`
	MarshalAllocs         int64  `json:"marshal_allocs"`
	UnmarshalAllocs       int64  `json:"unmarshal_allocs"`
	TotalAllocs           int64  `json:"total_allocs"`
	TimeSupport           string `json:"time_support"`
	APIKind               string `json:"api_kind"`
	URL                   string `json:"url"`
	Notes                 string `json:"notes"`
}

// 创建一个固定的测试数据
var testData = goserbench.SmallStruct{
	Name:     "Test Name",
	BirthDay: time.Now(),
	Phone:    "123-456-7890",
	Siblings: 2,
	Spouse:   true,
	Money:    123.45,
}

// CheckSerializerDeterminism 检查序列化器的确定性
// 返回是否确定以及任何错误
func CheckSerializerDeterminism(serializer goserbench.Serializer) (bool, error) {
	// 第一次序列化
	firstResult, err := serializer.Marshal(&testData)
	if err != nil {
		return false, err
	}

	// 后续9次序列化并比较
	for i := 1; i < 10; i++ {
		data, err := serializer.Marshal(&testData)
		if err != nil {
			return false, err
		}

		if string(firstResult) != string(data) {
			return false, nil
		}
	}

	return true, nil
}

func BenchAndReportSerializers(generateReport bool, validate bool, namesRe *regexp.Regexp) error {
	data := make([]reportLine, len(benchmarkCases))
	for i, bench := range benchmarkCases {
		if namesRe != nil && !namesRe.MatchString(bench.Name) {
			continue
		}

		marshalRes := testing.Benchmark(func(b *testing.B) {
			goserbench.BenchMarshalSmallStruct(b, bench.New())
		})
		fmt.Printf("%10s -   Marshal - %s %s\n", bench.Name, marshalRes.String(),
			marshalRes.MemString())

		deterministic, err := CheckSerializerDeterminism(bench.New())
		if err != nil {
			fmt.Printf("\nDeterminism check of %q failed: %v\n", bench.Name, err)
			fmt.Printf("Test with go test -validate -run Bench -bench BenchmarkSerializers/determinism/%s\n\n", bench.Name)
			return fmt.Errorf("benchmark %s failed", bench.Name)
		}
		if !deterministic {
			fmt.Printf("\nSerializer %q is not deterministic!\n", bench.Name)
		}

		unmarshalOk := false
		unmarshalRes := testing.Benchmark(func(b *testing.B) {
			goserbench.BenchUnmarshalSmallStruct(b, bench.New(), validate)
			unmarshalOk = true
		})
		if !unmarshalOk {
			fmt.Printf("\nUnmarshal benchmark of %q did not complete successfully\n", bench.Name)
			fmt.Printf("Test with go test -validate -run Bench -bench BenchmarkSerializers/unmarshal/%s\n\n", bench.Name)
			return fmt.Errorf("benchmark %s failed", bench.Name)
		} else {
			fmt.Printf("%10s - Unmarshal - %s %s\n", bench.Name, unmarshalRes.String(),
				unmarshalRes.MemString())
		}

		data[i] = reportLine{
			Name:                  bench.Name,
			MarshalIterCount:      marshalRes.N,
			UnmarshalIterCount:    unmarshalRes.N,
			Deterministic:         deterministic,
			TotalIterCount:        marshalRes.N + unmarshalRes.N,
			MarshalNsOp:           marshalRes.NsPerOp(),
			UnmarshalNsOp:         unmarshalRes.NsPerOp(),
			TotalNsOp:             marshalRes.NsPerOp() + unmarshalRes.NsPerOp(),
			UnsafeStringUnmarshal: bench.UnsafeStringUnmarshal,
			BufferReuseMarshal:    bench.BufferReuseMarshal,
			SerializationSize:     int64(marshalRes.Extra["B/serial"]),
			MarshalAllocBytes:     marshalRes.AllocedBytesPerOp(),
			UnmarshalAllocBytes:   unmarshalRes.AllocedBytesPerOp(),
			TotalAllocBytes: marshalRes.AllocedBytesPerOp() +
				unmarshalRes.AllocedBytesPerOp(),
			MarshalAllocs:   marshalRes.AllocsPerOp(),
			UnmarshalAllocs: unmarshalRes.AllocsPerOp(),
			TotalAllocs:     marshalRes.AllocsPerOp() + unmarshalRes.AllocsPerOp(),
			TimeSupport:     string(bench.TimeSupport),
			APIKind:         string(bench.APIKind),
			URL:             bench.URL,
			Notes:           strings.Join(bench.Notes, "\n"),
		}
	}

	if !generateReport {
		return nil
	}

	bytes, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	f, err := os.Create("report/data.js")
	if err != nil {
		return err
	}
	if _, err = f.Write([]byte("var data = ")); err != nil {
		return err
	}
	if _, err = f.Write(bytes); err != nil {
		return err
	}
	if _, err = f.Write([]byte(";")); err != nil {
		return err
	}

	fmt.Printf("\nSaved report to report/data.js !\n\n")
	return f.Close()
}
