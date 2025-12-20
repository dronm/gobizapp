package errs

type ErrorCode string

const (
	UnknownError        ErrorCode = "UNKNOWN_ERROR"
	InternalError       ErrorCode = "INTERNAL_ERROR"
	BadRequest          ErrorCode = "BAD_REQUEST"
	Timeout             ErrorCode = "REQUEST_TIMEOUT"
	ValidationFailed    ErrorCode = "VALIDATION_FAILED"
	NoSession           ErrorCode = "NO_SESSION"
	NotAllowed          ErrorCode = "NOT_ALLOWED"
	NotLogger           ErrorCode = "NOT_LOGGER"
	ExtNotSupported     ErrorCode = "EXT_NOT_SUPPORTED"
	ConstUndefined      ErrorCode = "CONST_UNDEFINED"
	ClientNotFoundByINN ErrorCode = "CLIENT_NOT_FOUND_BY_INN"
	DadataKeyUndefined  ErrorCode = "DADATA_KEY_UNDEFINED"
	BaseURLUndefined    ErrorCode = "BASE_URL_UNDEFINED"
	UserCredIncorrect   ErrorCode = "USER_CRED_INCORRECT"
	UserEmailNotFound   ErrorCode = "USER_EMAIL_NOT_FOUND"
	DBKeyExists         ErrorCode = "DB_KEY_EXISTS"
	DBRefExists         ErrorCode = "DB_REF_EXISTS"
)

var errorRegistry = map[ErrorCode]string{
	UnknownError:        "Unknown error",
	InternalError:       "An internal error occurred",
	BadRequest:          "Bad request",
	Timeout:             "Request timeout",
	ValidationFailed:    "Validation failed",
	NoSession:           "Session not found",
	NotAllowed:          "Not allowed",
	NotLogger:           "User is not logged",
	ExtNotSupported:     "Extention is not supported",
	ConstUndefined:      "Constant %s not defined",
	ClientNotFoundByINN: "По запросу ничего не найдено",
	DadataKeyUndefined:  "Не задан ключ dadata.ru",
	UserCredIncorrect:   "Неверное имя пользователя или пароль",
	UserEmailNotFound:   "Электронная почта не неайдена",
	DBKeyExists:         "Нарушение уникального ключа",
	DBRefExists:         "Существуют ссылки",
}

func ErrorDescr(code ErrorCode) string {
	descr, ok := errorRegistry[code]
	if !ok {
		descr = errorRegistry["UNKNOWN_ERROR"]
	}

	return descr
}
