package metrics

import (
	"runtime"
)

const (
	// The amount of memory allocated by the application.
	Alloc = "Alloc"

	// The amount of memory used by Go's internal hash table implementation
	BuckHashSys = "BuckHashSys"

	// The amount of memory freed by the application.
	Frees = "Frees"

	// The fraction of CPU time used by the garbage collector
	GCCPUFraction = "GCCPUFraction"

	// The amount of memory used by the garbage collector.
	GCSys = "GCSys"

	// The amount of memory allocated to the heap by the application
	HeapAlloc = "HeapAlloc"

	// The amount of idle (unused) heap memory.
	HeapIdle = "HeapIdle"

	// The amount of heap memory currently in use.
	HeapInuse = "HeapInuse"

	// The number of objects currently in the heap.
	HeapObjects = "HeapObjects"

	// The amount of heap memory that has been returned to the operating system[
	HeapReleased = "HeapReleased"

	// The total size of the heap memory.
	HeapSys = "HeapSys"

	// The time (in nanoseconds) since the last garbage collection.
	LastGC = "LastGC"

	// The number of pointer lookups performed by the garbage collector
	Lookups = "Lookups"

	// The amount of memory used by runtime mcache structures that are currently being used
	MCacheInuse = "MCacheInuse"

	// The total size of the memory reserved for runtime mcache structures
	MCacheSys = "MCacheSys"

	// The amount of memory used by runtime mspan structures that are currently being used
	MSpanInuse = "MSpanInuse"

	// The total size of the memory reserved for runtime mspan structures
	MSpanSys = "MSpanSys"

	// The number of heap memory allocations made by the application.
	Mallocs = "Mallocs"

	// The estimated heap size (in bytes) when the next garbage collection will occur.
	NextGC = "NextGC"

	// The number of garbage collections that have been forced by the application.
	NumForcedGC = "NumForcedGC"

	// The number of garbage collections that have been performed by the application.
	NumGC = "NumGC"

	// The amount of memory used by other system-level activities (e.g. network I/O
	OtherSys = "OtherSys"

	// The total time (in nanoseconds) spent by the garbage collector in performing pauses (i.e. when the application is stopped).
	PauseTotalNs = "PauseTotalNs"

	// The amount of stack memory currently in use by the application.
	StackInuse = "StackInuse"

	// The total size of the stack memory.
	StackSys = "StackSys"

	// The total memory used by the application (including heap, stack, and other system-level activities).
	Sys = "Sys"

	// The total amount of memory allocated by the application, including memory that has been freed.
	TotalAlloc = "TotalAlloc"
)

// Container that holds all mem stats metrics
type MemStats map[string]float64

// Read updates memory stats from runtime package
//
// Pre-cond:
//
// Post-cond: metrics updated using package runtime
// Reflect could be used here
func (m MemStats) Read() {
	var runtimeMemStat runtime.MemStats
	runtime.ReadMemStats(&runtimeMemStat)
	m[Alloc] = float64(runtimeMemStat.Alloc)
	m[BuckHashSys] = float64(runtimeMemStat.BuckHashSys)
	m[Frees] = float64(runtimeMemStat.Frees)
	m[GCCPUFraction] = float64(runtimeMemStat.GCCPUFraction)
	m[GCSys] = float64(runtimeMemStat.GCSys)
	m[HeapAlloc] = float64(runtimeMemStat.HeapAlloc)
	m[HeapIdle] = float64(runtimeMemStat.HeapIdle)
	m[HeapInuse] = float64(runtimeMemStat.HeapInuse)
	m[HeapObjects] = float64(runtimeMemStat.HeapObjects)
	m[HeapReleased] = float64(runtimeMemStat.HeapReleased)
	m[HeapSys] = float64(runtimeMemStat.HeapSys)
	m[LastGC] = float64(runtimeMemStat.LastGC)
	m[Lookups] = float64(runtimeMemStat.Lookups)
	m[MCacheInuse] = float64(runtimeMemStat.MCacheInuse)
	m[MCacheSys] = float64(runtimeMemStat.MCacheSys)
	m[MSpanInuse] = float64(runtimeMemStat.MSpanInuse)
	m[MSpanSys] = float64(runtimeMemStat.MSpanSys)
	m[Mallocs] = float64(runtimeMemStat.Mallocs)
	m[NextGC] = float64(runtimeMemStat.NextGC)
	m[NumForcedGC] = float64(runtimeMemStat.NumForcedGC)
	m[NumGC] = float64(runtimeMemStat.NumGC)
	m[OtherSys] = float64(runtimeMemStat.OtherSys)
	m[PauseTotalNs] = float64(runtimeMemStat.PauseTotalNs)
	m[StackInuse] = float64(runtimeMemStat.StackInuse)
	m[StackSys] = float64(runtimeMemStat.StackSys)
	m[Sys] = float64(runtimeMemStat.Sys)
	m[TotalAlloc] = float64(runtimeMemStat.TotalAlloc)
}
