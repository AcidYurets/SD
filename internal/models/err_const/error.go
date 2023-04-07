package err_const

// Константные ошибки

// Const тип используемый для константных ошибок, позволяет избегать возможных мутаций значений ошибок.
// Не рекомендуется использовать их для создания ошибок в рамках бизнес-логики.
type Const string

func (e Const) Error() string {
	return string(e)
}

const (
	ErrIdValidate   = Const("неверно указан Id объекта")
	ErrUuidValidate = Const("неверно указан Uuid объекта")

	MsgUnauthorized = "не авторизованный доступ"
	ErrUnauthorized = Const(MsgUnauthorized)
	ErrInvalidToken = Const("некорректный токен")

	MsgBadRequest = "ошибка параметров запроса"
	ErrBadRequest = Const(MsgBadRequest)

	MsgJsonUnMarshal = "не удалось декодировать JSON"
	ErrJsonUnMarshal = Const(MsgJsonUnMarshal)
	MsgJsonMarshal   = "не удалось упаковать данные в JSON"
	ErrJsonMarshal   = Const(MsgJsonMarshal)

	ErrDatabaseRecordNotFound = Const("запись не найдена")
	ErrUniqueViolation        = Const("нарушение уникальности ключа")

	ErrTransactionNotActive = Const("Транзакция не активна")
	ErrRollbackTransaction  = Const("ошибка отката изменений транзакции")

	ErrValidateModel = Const("Ошибка валидации модели")

	ErrAccessDenied          = Const("недостаточно прав")
	ErrMissingRequiredFields = Const("не хватает необходимых полей")
)
