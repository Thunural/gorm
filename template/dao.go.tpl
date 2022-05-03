package dao
import (
	"{{.ProjectName}}/model"
	"{{.ProjectName}}/global"
	"time"
)

// {{.StructName}}Dao ...
type {{.StructName}}Dao struct {
}

// Create 增
func (*{{.StructName}}Dao) Create(m *model.{{.StructName}}) (int, error) {
	m.CreateTime = time.Now()
	m.UpdateTime = time.Now()
	err := global.DB.Create(&m).Error
	if err != nil {
		return 0, err
	}
	return m.ID, nil
}

// Delete 删
func (*{{.StructName}}Dao) Delete(ids []int) error {
	err := global.DB.Where("id in(?)", ids).Delete(&model.{{.StructName}}{}).Error
	return err
}

// SelectByID 查
func (*{{.StructName}}Dao) SelectByID(id int64) (*model.{{.StructName}}, error) {
	var m model.{{.StructName}}
	err := global.DB.Where("id = ?", id).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return &m, nil
}


// Update 改  map[string]interface{}{"name": "hello", "age": 18, "actived": false}
func (*{{.StructName}}Dao) Update(m *model.{{.StructName}}) (*model.{{.StructName}}, error) {
	err := global.DB.Model(&m).Updates(m).Error
	if err != nil {
		return nil, err
	}

	err = global.DB.Where("id = ?", m.ID).Last(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// List 列表查询，不支持分页
func (*{{.StructName}}Dao) List() ([]model.{{.StructName}}, error) {
	var m []model.{{.StructName}}
	err := global.DB.Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Page 分页查询
func (*{{.StructName}}Dao) Page(params model.SelectPageReq) ([]model.{{.StructName}}, int64, error) {
	var m []model.{{.StructName}}
	var total int64
	DB := global.DB.Model(&m)
	if params.Page > 0 && params.Size > 0 {
		DB = DB.Limit(params.Size).Offset((params.Page - 1) * params.Size)
	}
	if len(params.Keyword) > 0 {
		DB = DB.Where("keyWord like ?", "%"+params.Keyword+"%")
	}
	if len(params.Sort) > 0 {
		DB = DB.Order(params.Order + params.Sort)
	}
	DB = DB.Count(&total)
	if err := DB.Find(&m).Error; err != nil {
		global.Log.Error(err.Error())
		return nil, 0, err
	}
	return m, total, nil
}
