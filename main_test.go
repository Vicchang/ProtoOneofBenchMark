package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"testing"

	"github.com/Vicchang/ProtoOneofBenchmark/protoobj"
	"google.golang.org/protobuf/proto"
)

func BenchmarkNormalMarshal(b *testing.B) {
	obj := &protoobj.Object{
		Text: "vvv",
		Kind: "kind",
		Type: "Wind",
	}

	for i := 0; i < b.N; i++ {
		proto.Marshal(obj)
	}
}

func BenchmarkOneofMarshal(b *testing.B) {
	obj := &protoobj.Object{
		Text: "vvv",
		Kind: "kind",
		Type: "Wind",
	}

	oneof := &protoobj.OneOfObject{
		Output: &protoobj.OneOfObject_Obj{
			Obj: obj,
		},
	}

	for i := 0; i < b.N; i++ {
		proto.Marshal(oneof)
	}
}

func BenchmarkNormalUnmarshal(b *testing.B) {
	obj := &protoobj.Object{
		Text: "vvv",
		Kind: "kind",
		Type: "Wind",
	}

	bs, _ := proto.Marshal(obj)

	t := protoobj.Object{}

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(bs, &t)
	}
}

func BenchmarkOneofUnmarshal(b *testing.B) {
	obj := &protoobj.Object{
		Text: "vvv",
		Kind: "kind",
		Type: "Wind",
	}

	oneof := &protoobj.OneOfObject{
		Output: &protoobj.OneOfObject_Obj{
			Obj: obj,
		},
	}

	bs, _ := proto.Marshal(oneof)

	t := protoobj.OneOfObject{}
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(bs, &t)
	}
}

func Helper() {
	f, err := os.Create("profile_oneof.prof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	err = pprof.StartCPUProfile(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pprof.StopCPUProfile()

	obj := &protoobj.Object{
		Text: "vvv",
		Kind: "kind",
		Type: "Wind",
	}

	oneof := &protoobj.OneOfObject{
		Output: &protoobj.OneOfObject_Obj{
			Obj: obj,
		},
	}

	// Call proto.Marshal in a loop to generate more samples
	for i := 0; i < 10000000; i++ {
		proto.Marshal(oneof)
	}

	// Use pprof to print the call stack of the fibonacci function
	//pprof.Lookup("cpu").WriteTo(os.Stdout, 1)
}
