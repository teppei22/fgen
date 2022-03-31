package {{ .Model.PkgName }}

import "time"

type {{ .Model.Name }} struct {
	ID          int       `json:"id" param:"id"`
    
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
