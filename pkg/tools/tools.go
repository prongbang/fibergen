package tools

import (
	"github.com/prongbang/fibergen/pkg/arch"
	"github.com/prongbang/fibergen/pkg/command"
	"github.com/pterm/pterm"
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
	spinnerSqlc, _ := pterm.DefaultSpinner.Start("Install sqlc")
	_, err := cmd.Run("sqlc", "version")
	if err != nil {
		if arc.IsDarwinArm64() {
			_, err = cmd.RunAsync("arch", "-arm64", "brew", "install", "sqlc")
		} else {
			_, err = cmd.Run("brew", "install", "sqlc")
		}
		if err != nil {
			spinnerSqlc.Fail()
		} else {
			spinnerSqlc.Success()
		}
	} else {
		spinnerSqlc.Success()
	}

	return err
}

func New() Tools {
	return &tool{}
}
