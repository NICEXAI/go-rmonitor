package mem

type MemoryStat struct {
	// Total amount of RAM on this system
	Total uint64 `json:"total"`
	// RAM available for programs to allocate
	Available uint64 `json:"available"`
	// RAM used by programs
	Used uint64 `json:"used"`
	// Percentage of RAM used by programs
	UsedPercent float64 `json:"usedPercent"`
	// Buffer/Cached
	Buffers uint64 `json:"buffers"`
	Cached  uint64 `json:"cached"`
	// This is the kernel's notion of free memory; RAM chips whose bits nobody
	// cares about the value of right now. For a human consumable number,
	// Available is what you really want.
	free         uint64
	sreClaimAble uint64
}
