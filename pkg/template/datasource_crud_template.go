package template

import (
	"fmt"
)

func DataSourceCrud(
	model string,
	module string,
	path string,
	pk string,
	driver string,
	queryColumns string,
	insertValues string,
	insertFields string,
	insertQuestions string,
	updateSets string,
) string {
	if driver == "mariadb" {
		return DataSourceCrudMariaDb(
			model,
			module,
			path,
			pk,
			"GetMariaDB",
			queryColumns,
			insertValues,
			insertFields,
			insertQuestions,
			updateSets,
		)
	} else if driver == "postgres" {
		return DataSourceCrudPostgrest(
			model,
			module,
			path,
			pk,
			"GetPostgresDB",
			queryColumns,
			insertValues,
			insertFields,
			insertQuestions,
			updateSets,
		)
	}
	panic(fmt.Sprintf("unsupported driver: %s", driver))
}
