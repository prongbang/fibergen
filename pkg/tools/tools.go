package tools

import (
	"github.com/prongbang/fibergen/pkg/arch"
	"github.com/prongbang/fibergen/pkg/command"
)

type Tools interface {
	Install() error
}

type tool struct{}

// Install implements Tools.
func (*tool) Install() error {
	cmd := command.New()
	arc := arch.New()

	// Install sqlc
	_, err := cmd.Run("sqlc", "version")
	if err != nil {
		if arc.IsDarwinArm64() {
			_, err = cmd.RunAsync("arch", "-arm64", "brew", "install", "sqlc")
		} else {
			_, err = cmd.Run("brew", "install", "sqlc")
		}
	}

	return err
}

func New() Tools {
	return &tool{}
}
