package arch

import "runtime"

type Arch interface {
	IsDarwinArm64() bool
}

type a struct{}

// IsDarwinArm64 implements Arch.
func (*a) IsDarwinArm64() bool {
	os := runtime.GOOS
	arch := runtime.GOARCH
	return os == "darwin" && arch == "arm64"
}

func New() Arch {
	return &a{}
}
