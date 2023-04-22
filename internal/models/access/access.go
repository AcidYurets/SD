package access

import "strings"

type Type string

// Возможные права доступа
const (
	ReadAccess   Type = "r" // Право видеть содержание события
	InviteAccess Type = "i" // Право на приглашение других пользователей
	UpdateAccess Type = "u" // Право на изменение
	DeleteAccess Type = "d" // Право на удаление
)

// Возможные комбинации прав доступа
const (
	OnlyReadAccess         Type = "r"    // Право видеть содержание события
	ReadInviteAccess       Type = "ri"   // Право на чтение и приглашение других пользователей
	ReadInviteUpdateAccess Type = "riu"  // Право на чтение, приглашение других пользователей и изменение
	FullAccess             Type = "riud" // Право на чтение, приглашение других пользователей, изменение и удаление
)

func (s Type) String() string {
	return string(s)
}

func (s Type) IsValid() bool {
	all := ReadAccess + InviteAccess + UpdateAccess + DeleteAccess
	for _, token := range s {
		if !strings.ContainsRune(string(all), token) {
			return false
		}
	}
	return true
}
