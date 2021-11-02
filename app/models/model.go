package models

import "jasy/goblog/pkg/types"

type BaseModel struct {
	ID uint64
}

func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}