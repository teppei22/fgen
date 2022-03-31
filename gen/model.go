package gen

type Config struct {
	outputPath string
}

type CodeInfo struct {
}

type InitInfo struct {
	Name         string
	TemplatePath string
	OutputDir    string
}

type LayerInfo struct {
	PkgName      string
	DirPath      string
	Name         string
	TypeName     string
	ReceiverChar string
}

type ModelInfo struct {
	Name   string
	Fields []FieldInfo
}

type FieldInfo struct {
	FieldName string
	TypeName  string
	Tag       struct {
		Json  string
		Param string
	}
}
