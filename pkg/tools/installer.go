package tools

type Installer interface {
	Install() error
}

type installer struct {
	WireInstaller Installer
	SqlcInstaller Installer
	DbmlInstaller Installer
}

// Install implements Installer.
func (t *installer) Install() error {

	var err error

	// Install wire
	if e1 := t.WireInstaller.Install(); e1 != nil {
		err = e1
	}

	// Install sqlc
	if e2 := t.SqlcInstaller.Install(); e2 != nil {
		err = e2
	}

	// Install dbml-cli
	if e3 := t.DbmlInstaller.Install(); e3 != nil {
		err = e3
	}

	return err
}

func New(wireInstaller Installer, sqlcInstaller Installer, dbmlInstaller Installer) Installer {
	return &installer{
		WireInstaller: wireInstaller,
		SqlcInstaller: sqlcInstaller,
		DbmlInstaller: dbmlInstaller,
	}
}
