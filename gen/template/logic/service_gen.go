package logic

var ServiceGenTemplate = `
package logic

import (
	"{{.ProjectName}}/dao"
	"{{.ProjectName}}/model"
)

// {{.StructName}}Logic 服务
type {{.StructName}}Logic struct {
}

// Create 创建
func ( *{{.StructName}}Logic) Create(p *model.{{.StructName}}) (int64, error) {
	id, err := dao.{{.StructName}}.Create(p)
	if err != nil {
		return 0, err
	}
	  
	return id, nil
}

 
// Delete  ...
func ( *{{.StructName}}Logic) Delete(ids []int64) error {
	return dao.{{.StructName}}.Delete(ids)
}

// Select ...
func (*{{.StructName}}Logic) SelectByID(id int64) (*model.{{.StructName}}, error) {
	data, err := dao.{{.StructName}}.SelectByID(id)
	if err != nil {
		return nil, err
	}

	return data, nil
}


// Update ...
func (*{{.StructName}}Logic) Update(p *model.{{.StructName}}) (*model.{{.StructName}}, error) {
	data, err := dao.{{.StructName}}.Update(p)
	if err != nil {
		return nil, err
	}

	return data, nil
}
`
