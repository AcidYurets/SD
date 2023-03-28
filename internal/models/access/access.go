package access

import "strings"

type Type string

// Системные права
const (
	ReadAccess   Type = "r" // Право видеть содержание события
	InviteAccess Type = "i" // Право на приглашение других пользователей
	UpdateAccess Type = "u" // Право на изменение
	FullAccess   Type = "d" // Право на удаление
)

func (s Type) String() string {
	return string(s)
}

func (s Type) IsValid() bool {
	all := ReadAccess + InviteAccess + UpdateAccess + FullAccess
	for _, token := range s {
		if !strings.ContainsRune(all, token) {
			return false
		}
	}
	return true
}
