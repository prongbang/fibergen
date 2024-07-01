package template

import (
	"fmt"
	"strings"

	"github.com/ettle/strcase"
)

func DataSourceCrudPostgrest(
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
		"github.com/prongbang/sqlxwrapper/pqwrapper"
	)
	
	type DataSource interface {
		Count(params Params) int64
		FindList(params Params) []{model}
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
		sql := "SELECT COUNT(id) FROM {table} WHERE 1=1 %s"
		wheres := ""
		args := []any{}
	
		sql = fmt.Sprintf(sql, wheres)
	
		return pqwrapper.Count(conn, sql, args...)
	}
	
	func (d *dataSource) FindList(params Params) []{model} {
		conn := d.Driver.{driver}()
		sql := "SELECT {columns} FROM {table} {alias} WHERE 1=1 %s ORDER BY {alias}.id"
		wheres := ""
		args := []any{}
	
		if params.Limit > 0 && params.Offset >= 0 {
			sql += " LIMIT ?"
			args = append(args, params.Limit)
			sql += " OFFSET ?"
			args = append(args, params.Offset)
		}
	
		sql = fmt.Sprintf(sql, wheres)
	
		return pqwrapper.SelectList[{model}](conn, sql, args...)
	}
	
	func (d *dataSource) FindById(id {pk}) {model} {
		conn := d.Driver.{driver}()
		sql := "SELECT {columns} FROM {table} {alias} WHERE {alias}.id = $1"
	
		return pqwrapper.SelectOne[{model}](conn, sql, id)
	}
	
	func (d *dataSource) Create(data *Create{model}) error {
		conn := d.Driver.{driver}()
		sql := "INSERT INTO {table} ({fields}) VALUES ({questions}) RETURNING id"
		args := []any{
			{values}
		}
		tx, err := pqwrapper.Create(conn, sql, []any{&data.Id}, args...)
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
	
		tx, err := pqwrapper.Update(conn, sql, set, params)
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
		sql := "DELETE FROM {table} WHERE id = $1"
	
		tx, err := pqwrapper.Delete(conn, sql, id)
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
	tmpl = strings.ReplaceAll(tmpl, "{name}", strings.ToLower(model))

	// Query
	tmpl = strings.ReplaceAll(tmpl, "{table}", strings.ToLower(model))
	tmpl = strings.ReplaceAll(tmpl, "{alias}", strings.ToLower(model)[:1])
	tmpl = strings.ReplaceAll(tmpl, "{columns}", queryColumns)

	// Insert
	tmpl = strings.ReplaceAll(tmpl, "{values}", insertValues)
	tmpl = strings.ReplaceAll(tmpl, "{fields}", insertFields)

	questions := []string{}
	qstr := strings.Split(insertQuestions, ",")
	for i, _ := range qstr {
		questions = append(questions, fmt.Sprintf("$%d", i+1))
	}
	tmpl = strings.ReplaceAll(tmpl, "{questions}", strings.Join(questions, ","))

	// Update
	tmpl = strings.ReplaceAll(tmpl, "{sets}", updateSets)

	return tmpl
}
