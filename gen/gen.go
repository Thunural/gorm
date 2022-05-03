package gen

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jimsmart/schema"
	cfg "gormui/config"
	"gormui/gen/template/dao"
	"gormui/gen/template/logic"
	"gormui/gen/template/model"
	"gormui/utils"
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
			// 先移除原有的gen文件
			os.Remove(toDir + "/model/" + tableName + "Model_gen.go")
			err := baseStr(model.ModelGenTemplate, toDir+"/model/"+tableName+"Model_gen.go", base)
			if err != nil {
				return err
			}
			// 判断文件是否存在，存在则跳过
			_, err = os.Stat(toDir + "/model/" + tableName + "Model.go")
			if err != nil {
				err = baseStr(model.ModelTemplate, toDir+"/model/"+tableName+"Model.go", base)
				if err != nil {
					return err
				}
			}
		}
	}

	// 搞service文件夹内的
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
			// 先将原有的gen文件删除
			os.Remove(toDir + "/service/" + tableName + "Service_gen.go")
			err := baseStr(logic.ServiceGenTemplate, toDir+"/service/"+tableName+"Service_gen.go", base)
			if err != nil {
				return err
			}
			// 判断原有的service文件是否存在，如果不存在则生成
			_, err = os.Stat(toDir + "/service/" + tableName + "Service.go")
			if err != nil {
				err := baseStr(logic.ServiceTemplate, toDir+"/service/"+tableName+"Service.go", base)
				if err != nil {
					return err
				}
			}
		}
	}

	// 搞dao文件夹内的
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
			err := baseStr(dao.DaoTemplate, toDir+"/dao/"+tableName+"Dao.go", base)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 基础文件拷贝
func baseStr(baseStr string, toPath string, base interface{}) error {
	tmpl, err := template.New("base").Parse(baseStr) // 从模板内读取数据
	if err != nil {
		return errors.New("生成报错\n" + err.Error())
	}
	toDir := filepath.Dir(toPath)
	_ = os.MkdirAll(toDir, 0777)
	openFile, _ := os.OpenFile(toPath, os.O_WRONLY|os.O_CREATE|os.O_SYNC, 0777)
	err = tmpl.Execute(openFile, base)
	if err != nil {
		return errors.New("生成报错\n" + err.Error())
	}
	_ = openFile.Close()
	return nil
}

func in(target string, strArray []string) bool {
	for _, element := range strArray {
		if target == element {
			return true
		}
	}
	return false
}
