package logic

var ServiceGenTemplate = `
package logic

import (
	"{{.ProjectName}}/dao"
	"{{.ProjectName}}/model"
)

// {{.StructName}}Service 服务
type {{.StructName}}Service struct {
}

// Create 创建
func ( *{{.StructName}}Service) Create(p *model.{{.StructName}}) (int64, error) {
	id, err := dao.{{.StructName}}.Create(p)
	if err != nil {
		return 0, err
	}
	  
	return id, nil
}

 
// Delete  ...
func ( *{{.StructName}}Service) Delete(ids []int64) error {
	return dao.{{.StructName}}.Delete(ids)
}

// Select ...
func (*{{.StructName}}Service) SelectByID(id int64) (*model.{{.StructName}}, error) {
	data, err := dao.{{.StructName}}.SelectByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}


// Update ...
func (*{{.StructName}}Service) Update(p *model.{{.StructName}}) (*model.{{.StructName}}, error) {
	data, err := dao.{{.StructName}}.Update(p)
	if err != nil {
		return nil, err
	}

	return data, nil
}
`
