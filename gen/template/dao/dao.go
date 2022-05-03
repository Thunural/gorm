package dao

var DaoTemplate = `
package dao
import (
	"{{.ProjectName}}/model"
	"{{.ProjectName}}/global"
	"time"
)

var (
    {{.StructName}} = {{.StructName}}Dao{}
)

// TODO 你的逻辑写在这里
`
