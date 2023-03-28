package dto

type Tag struct {
	Uuid        string // Uuid пользователя
	Name        string // Наименование
	Description string // Описание
}

type Tags []*Tag

type CreateTag struct {
	Name        string // Наименование
	Description string // Описание
}

type UpdateTag struct {
	Name        *string // Наименование
	Description *string // Описание
}
