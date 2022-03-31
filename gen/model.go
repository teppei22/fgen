package gen

type Config struct {
	outputPath string
}

type CodeInfo struct {
}

type TempInfo struct {
	Name         string
	TemplatePath string
	OutputDir    string
}

type OutputImplData struct {
	Handler     ImplFileData
	Usecase     ImplFileData
	Persistence ImplFileData
	Repository  ImplFileData
}

type ImplFileData struct {
	PkgName      string
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
