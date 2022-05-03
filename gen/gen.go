package gen

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimsmart/schema"
	cfg "gormui/config"
	"gormui/utils"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func Project(c cfg.Param) error {
	outFiles := c.OutFiles
	ProjectName := c.ProjectName
	toDir := outFiles + "/" + ProjectName
	_ = os.Mkdir(toDir, 0777)

	// 搞model文件夹内的

	{
		if in("model", c.CheckType) {
			os.Mkdir(toDir+"/model", 0777)
			// 获取所有数据库表
			dbStr := c.UserName + ":" + c.Password + "@tcp(" + c.Address + ":" + c.Port + ")/" + c.Database + "?charset=utf8&parseTime=true&loc=Local"
			var db, err = sql.Open("mysql", dbStr)
			if err != nil {
				return errors.New("打开数据库失败:\n" + err.Error())
			}
			defer db.Close()
			var tables = make([]string, 0)
			// parse or read tables
			tables, err = schema.TableNames(db)
			if err != nil {
				return errors.New("类型获取失败:\n" + err.Error())
			}
			// generate go files for each table
			for _, tableName := range tables {
				structName := utils.FmtFieldName(tableName)
				if structName[len(structName)-1] == 's' {
					structName = structName[0 : len(structName)-1]
				}

				modelInfo := GenerateStruct(db, tableName, structName, "model", true, true, true)

				var base = struct {
					StructName string
					TableName  string
					Fields     []string
				}{
					StructName: structName,
					TableName:  tableName,
					Fields:     modelInfo.Fields,
				}
				// 同时生成gen
				baseStr("./template/model_gen.go.tpl", toDir+"/model/"+tableName+"Model_gen.go", base)
				baseStr("./template/model.go.tpl", toDir+"/model/"+tableName+"Model.go", base)
			}
			{
				baseStr("./template/common.go.tpl", toDir+"/model/common.go", "")
			}
		}
	}

	// 搞service文件夹内的
	{
		if in("service", c.CheckType) {
			os.Mkdir(toDir+"/service", 0777)
			dbStr := c.UserName + ":" + c.Password + "@tcp(" + c.Address + ":" + c.Port + ")/" + c.Database + "?charset=utf8&parseTime=true&loc=Local"
			var db, err = sql.Open("mysql", dbStr)
			if err != nil {
				return errors.New("打开数据库失败：\n" + err.Error())
			}
			defer db.Close()
			var tables = make([]string, 0)
			// parse or read tables
			tables, err = schema.TableNames(db)
			// generate go files for each table
			for _, tableName := range tables {
				structName := utils.FmtFieldName(tableName)
				if structName[len(structName)-1] == 's' {
					structName = structName[0 : len(structName)-1]
				}

				modelInfo := GenerateStruct(db, tableName, structName, "model", true, false, true)

				var base = struct {
					ProjectName  string
					StructName   string
					FieldsCreate []string
					Fields       []string
				}{
					ProjectName:  ProjectName,
					StructName:   structName,
					FieldsCreate: modelInfo.Fields, //这里应该去掉主键，但是懒得弄了
					Fields:       modelInfo.Fields,
				}
				baseStr("./template/service.go.tpl", toDir+"/service/"+tableName+"Service.go", base)
			}
		}
	}

	// 搞dao文件夹内的
	{
		log.Println(in("dao", c.CheckType))
		if in("dao", c.CheckType) {
			_ = os.Mkdir(toDir+"/dao", 0777)
			dbStr := c.UserName + ":" + c.Password + "@tcp(" + c.Address + ":" + c.Port + ")/" + c.Database + "?charset=utf8&parseTime=true&loc=Local"
			var db, err = sql.Open("mysql", dbStr)
			if err != nil {
				return errors.New("打开数据库失败：\n" + err.Error())
			}
			defer db.Close()
			// parse or read tables
			var tables = make([]string, 0)
			tables, err = schema.TableNames(db)
			// generate go files for each table
			for _, tableName := range tables {
				structName := utils.FmtFieldName(tableName)
				if structName[len(structName)-1] == 's' {
					structName = structName[0 : len(structName)-1]
				}

				var base = struct {
					ProjectName string
					StructName  string
				}{
					ProjectName: ProjectName,
					StructName:  structName,
				}
				baseStr("./template/dao.go.tpl", toDir+"/dao/"+tableName+"Dao.go", base)
			}
		}
	}
	return nil
}

// 基础文件拷贝
func baseStr(fileName string, toPath string, base interface{}) {
	tmpl, err := template.ParseFiles(fileName) // 从模板内读取数据
	if err != nil {
		fmt.Println(err)
		return
	}
	type Main struct {
		ProjectName string
	}
	toDir := filepath.Dir(toPath)
	_ = os.MkdirAll(toDir, 0777)
	openFile, _ := os.OpenFile(toPath, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0755)
	err = tmpl.Execute(openFile, base)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = openFile.Close()
}

func in(target string, strArray []string) bool {
	for _, element := range strArray {
		if target == element {
			return true
		}
	}
	return false
}
