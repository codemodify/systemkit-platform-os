package os

// KernelInfo -
type KernelInfo struct {
	Name       Kernel     `json:"name,omitempty"`       // "Linux"
	Version    string     `json:"version,omitempty"`    // `5.0` on Linux
	BuildTime  string     `json:"buildTime,omitempty"`  // #1 SMP PREEMPT Wed, 29 Apr 2020 16:22:56 +0000
	ExecFormat ExecFormat `json:"execFormat,omitempty"` // ELF
}

// OS -
type OS struct {
	Family   Family     `json:"family,omitempty"`   // ex: Linux, Windows, Darwin
	Name     Name       `json:"name,omitempty"`     // Linux, MINGW, MacOS X
	Version  string     `json:"version,omitempty"`  // `10.0.14393.1066` on Windows
	Kernel   KernelInfo `json:"kernel,omitempty"`   //
	Features []string   `json:"features,omitempty"` // ex: `win32k` on Windows
}

// Info -
func Info() OS {
	return info()
}
