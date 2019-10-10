package definitions

// OrderBy Order By
type OrderBy struct {
	Index int `json:"index"`
	FieldName string `json:"fieldName"`
	Direction string `json:"direction"`
}
