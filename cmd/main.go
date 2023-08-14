package main

import "github.com/tf63/go_workspace/internal"

func main() {
	/* internal.varSample() */ // キャメルケースではダメ
	// internal.VarSample()       // パスカルケースじゃないといけない

	// internal.FuncInputSample()
	// internal.NilSample()
	// internal.StringSample()

	// internal.PrintSample()
	// internal.ForSample()
	// internal.SwitchSample()
	// internal.DeferSample()

	// internal.StructSample()
	// internal.StructPointerSample()
	// internal.StructFuncSample()
	// internal.StructMethodSample()
	// internal.StructExtendSample()

	// internal.ArrayTypeSample()
	// internal.ArraySample()
	// internal.ArraySizeSample()

	internal.MapSample()
}
