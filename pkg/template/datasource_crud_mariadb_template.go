package template

import (
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
		"github.com/prongbang/sqlxwrapper/mrwrapper"
		"github.com/prongbang/goerror"
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
	
		return mrwrapper.Count(conn, sql, args...)
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
	
		return mrwrapper.SelectList[{model}](conn, sql, args...)
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
				return goerror.NewError(goerror.Body{Code: "DTB001", Message: "Cannot add a child row"})
			}
			return nil
		}
		return err
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
				return goerror.NewError(goerror.Body{Code: "DTB001", Message: "Cannot add a child row"})
			}
			return nil
		}
		return err
	}
	
	func (d *dataSource) Delete(id {pk}) error {
		conn := d.Driver.{driver}()
		sql := "DELETE FROM {table} WHERE id=?"
	
		tx, err := mrwrapper.Delete(conn, sql, id)
		if err == nil {
			if e := tx.Commit(); e != nil {
				return goerror.NewError(goerror.Body{Code: "DTB001", Message: "Cannot add a child row"})
			}
			return nil
		}
		return err
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
	tmpl = strings.ReplaceAll(tmpl, "{questions}", insertQuestions)

	// Update
	tmpl = strings.ReplaceAll(tmpl, "{sets}", updateSets)

	return tmpl
}
