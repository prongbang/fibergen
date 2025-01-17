package template

import (
	"github.com/prongbang/fibergen/pkg/tocase"
	"strings"

	"github.com/ettle/strcase"
)

func DataSourceCrudMariaDb(
	model string,
	module string,
	path string,
	pk string,
	driverName string,
	queryColumns string,
	insertValues string,
	insertFields string,
	insertQuestions string,
	updateSets string,
) string {
	tmpl := `package {name}
	
import (
	"fmt"
	"{module}/{path}/database"
	"{module}/internal/pkg/response"
	"{module}/pkg/schema"
	"github.com/iancoleman/strcase"
	"github.com/innotechdevops/core/finder"
	"github.com/prongbang/sqlxwrapper/mrwrapper"
)

type DataSource interface {
	Count(params Params) int64
	FindList(params Params) []{model}
	FindLiteList(params LiteParams) []Lite{model}
	FindById(id {pk}) {model}
	Create(data *Create{model}) error
	Update(data *Update{model}) error
	Delete(id {pk}) error
}

type dataSource struct {
	Driver database.Drivers
}

func (d *dataSource) Count(params Params) int64 {
	conn := d.Driver.{driver}()
	sql := "SELECT COUNT({alias}.id) FROM {table} {alias} WHERE 1=1 %s"
	wheres := ""
	args := []any{}

	sql = fmt.Sprintf(sql, wheres)

	return mrwrapper.Count(conn, sql, args...)
}

func (d *dataSource) FindList(params Params) []{model} {
	conn := d.Driver.{driver}()
	sql := "SELECT {columns} FROM {table} {alias} WHERE 1=1 %s %s "
	wheres := ""
	args := []any{}

	if params.Limit > 0 && params.Offset >= 0 {
		sql += " LIMIT ?"
		args = append(args, params.Limit)
		sql += " OFFSET ?"
		args = append(args, params.Offset)
	}

	order := "ORDER BY {alias}.id "
	params.Sort = strcase.ToSnake(params.Sort)
	if finder.Match(columns, params.Sort) {
		order = fmt.Sprintf("ORDER BY {alias}.%s ", params.Sort)
	}
	if finder.Match(schema.OrderBy, params.Order) {
		order += fmt.Sprintf(" %s ", params.Order)
	}

	sql = fmt.Sprintf(sql, wheres, order)

	return mrwrapper.SelectList[{model}](conn, sql, args...)
}

func (d *dataSource) FindLiteList(params LiteParams) []Lite{model} {
	conn := d.Driver.{driver}()
	sql := "SELECT {columns} FROM {table} {alias} WHERE 1=1 %s %s "
	wheres := ""
	args := []any{}

	order := "ORDER BY {alias}.id "
	params.Sort = strcase.ToSnake(params.Sort)
	if finder.Match(columns, params.Sort) {
		order = fmt.Sprintf("ORDER BY {alias}.%s ", params.Sort)
	}
	if finder.Match(schema.OrderBy, params.Order) {
		order += fmt.Sprintf(" %s ", params.Order)
	}

	sql = fmt.Sprintf(sql, wheres, order)

	return mrwrapper.SelectList[Lite{model}](conn, sql, args...)
}

func (d *dataSource) FindById(id {pk}) {model} {
	conn := d.Driver.{driver}()
	sql := "SELECT {columns} FROM {table} {alias} WHERE {alias}.id = ?"

	return mrwrapper.SelectOne[{model}](conn, sql, id)
}

func (d *dataSource) Create(data *Create{model}) error {
	conn := d.Driver.{driver}()
	sql := "INSERT INTO {table} ({fields}) VALUES ({questions})"
	args := []any{
		{values}
	}
	tx, err := mrwrapper.Create(conn, sql, []any{&data.Id}, args...)
	if err == nil {
		if e := tx.Commit(); e != nil {
			return response.NewCommitError()
		}
		return nil
	}
	return response.NewInsertError()
}

func (d *dataSource) Update(data *Update{model}) error {
	conn := d.Driver.{driver}()
	params := map[string]interface{}{
		"id": data.Id,
	}
	set := ""
	sql := "UPDATE {table} SET %s WHERE id=:id"

	{sets}

	tx, err := mrwrapper.Update(conn, sql, set, params)
	if err == nil {
		if e := tx.Commit(); e != nil {
			return response.NewCommitError()
		}
		return nil
	}
	return response.NewUpdateError()
}

func (d *dataSource) Delete(id {pk}) error {
	conn := d.Driver.{driver}()
	sql := "DELETE FROM {table} WHERE id=?"

	tx, err := mrwrapper.Delete(conn, sql, id)
	if err == nil {
		if e := tx.Commit(); e != nil {
			return response.NewCommitError()
		}
		return nil
	}
	return response.NewDeleteError()
}

func NewDataSource(driver database.Drivers) DataSource {
	return &dataSource{
		Driver: driver,
	}
}`

	tmpl = strings.ReplaceAll(tmpl, "{module}", module)
	tmpl = strings.ReplaceAll(tmpl, "{path}", path)
	tmpl = strings.ReplaceAll(tmpl, "{pk}", pk)
	tmpl = strings.ReplaceAll(tmpl, "{model}", strcase.ToPascal(model))
	tmpl = strings.ReplaceAll(tmpl, "{driver}", driverName)
	tmpl = strings.ReplaceAll(tmpl, "{name}", tocase.ToLower(model))

	// Query
	tmpl = strings.ReplaceAll(tmpl, "{table}", strings.ToLower(model))
	tmpl = strings.ReplaceAll(tmpl, "{alias}", strings.ToLower(model)[:1])
	tmpl = strings.ReplaceAll(tmpl, "{columns}", queryColumns)

	// Insert
	tmpl = strings.ReplaceAll(tmpl, "{values}", insertValues)
	tmpl = strings.ReplaceAll(tmpl, "{fields}", insertFields)
	tmpl = strings.ReplaceAll(tmpl, "{questions}", insertQuestions)

	// Update
	tmpl = strings.ReplaceAll(tmpl, "{sets}", updateSets)

	return tmpl
}
