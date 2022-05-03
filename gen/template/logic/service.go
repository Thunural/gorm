package logic

var ServiceTemplate = `
package logic

var (
	{{.StructName}} = {{.StructName}}Service{}
)

// TODO 你的新逻辑可写在这里
`
