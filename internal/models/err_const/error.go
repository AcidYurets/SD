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

	MsgBadRequest = "ошибка параметров запроса"
	ErrBadRequest = Const(MsgBadRequest)

	MsgJsonUnMarshal = "не удалось декодировать JSON"
	ErrJsonUnMarshal = Const(MsgJsonUnMarshal)
	MsgJsonMarshal   = "не удалось упаковать данные в JSON"
	ErrJsonMarshal   = Const(MsgJsonMarshal)

	// DB
	ErrDatabaseRecordNotFound = Const("запись не найдена")
	ErrUniqueViolation        = Const("нарушение уникальности ключа")

	/* Ошибки транзакций БД */
	ErrTransactionNotActive = Const("Транзакция не активна")
	ErrRollbackTransaction  = Const("ошибка отката изменений транзакции")

	ErrValidateModel = Const("Ошибка валидации модели")

	ErrDocumentRequired = Const("не указан документ")

	// Права доступа
	ErrAccessDenied = Const("недостаточно прав")
)
