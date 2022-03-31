package gen

type Config struct {
	OutputPath string
	Model      string
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
	UseCase     ImplFileData
	Persistence ImplFileData
	Repository  ImplFileData
	Model       ModelInfo
}

type ImplFileData struct {
	PkgName       string
	StructName    string
	InterfaceName string
	ReceiverChar  string
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
