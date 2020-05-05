package os

type ExecFormat string

// https://en.wikipedia.org/wiki/Comparison_of_operating_system_kernels
const (
	ExecFormat_HUNK    ExecFormat = "HUNK"
	ExecFormat_ELF     ExecFormat = "ELF"
	ExecFormat_MachO   ExecFormat = "Mach-O"
	ExecFormat_NLM     ExecFormat = "NLM"
	ExecFormat_LX      ExecFormat = "LX"
	ExecFormat_PE      ExecFormat = "PE"
	ExecFormat_AOUT    ExecFormat = "a.out"
	ExecFormat_Unknown ExecFormat = "Unknown"
)
