package dto

type GenerDto struct {
	Kind   uint32 `json:"kind" validate:"required,gt=0,lt=8"`
	Gender uint32 `json:"gender" validate:"required,gte=0,lt=3"`
	Length uint32 `json:"length" validate:"required,gte=1,lt=10"`
	Number uint32 `json:"number" validate:"required,gte=1,lte=150"`
	Repeat bool   `json:"repeat"`
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}
