package mmio

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"
)

// NewPprof starts a new pprof and outputs to file cpu.pprof
func NewPprof() error {
	file, err := os.Create("cpu.pprof")
	if err != nil {
		return fmt.Errorf("Cannot create file: %v", err)
	}
	pprof.StartCPUProfile(file)
	return nil
}

// EndPprof ends pprof (set as a defer statement after NewPprof)
func EndPprof() {
	pprof.StopCPUProfile()
}

// NewTrace starts a new pprof and outputs to file cpu.trace
func NewTrace() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
}

// EndTrace ends trace (set as a defer statement after NewTrace)
func EndTrace() {
	trace.Stop()
}

// HeapDump writes a heap profiling file
func HeapDump(fp string) {
	if len(fp) == 0 {
		fp = "heap.out"
	}
	f, err := os.Create(fp)
	if err != nil {
		log.Fatalf("failed to create heap output file: %v", err)
	}
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatalf("failed to create heap profilling: %v", err)
	}
}
