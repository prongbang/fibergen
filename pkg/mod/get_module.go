package mod

import (
	"fmt"
	"strings"

	"github.com/prongbang/fibergen/pkg/filex"
)

func GetModule(fx filex.FileX) Mod {
	// Get current path
	pwd, _ := fx.Getwd()

	// Get root project path
	changeToRootProject := "../../../"
	_ = fx.Chdir(changeToRootProject)
	root, _ := fx.Getwd()
	if bt := fx.ReadFile(root + "/go.mod"); bt != "" {
		// Find module
		text := bt
		m := "module "
		s := strings.Index(text, m)
		e := strings.Index(text, "\n")
		if s < 0 && e < 0 {
			return Mod{}
		}
		mod := text[s+len(m) : e]

		// Find app path
		i := strings.LastIndex(mod, "/")
		if i < 0 {
			return Mod{}
		}
		pj := mod[i:]
		ign := "/api"
		c := strings.Index(pwd, fmt.Sprintf("%s/internal/", pj))

		// Find internal/project-name
		ap := pwd[c+len(pj)+1 : len(pwd)-len(ign)]

		_ = fx.Chdir("./" + ap + ign)

		return Mod{
			Module:  mod,
			AppPath: ap,
		}
	}
	return Mod{}
}
