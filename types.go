package nsclient

const (
	MESSAGE_TYPE_ERROR = iota
	MESSAGE_TYPE_TEST
	MESSAGE_TYPE_INSERT
)

const (
	_ = iota
	// MessageTypeError - Тип сообщения: ошибка
	MessageTypeError
	// MessageTypeObjectlist -  Тип сообщения: перечень
	MessageTypeObjectlist
	// MessageTypeFile - Тип сообщения: файл
	MessageTypeFile
)

// MessagePackage -структурка для хранения мессаджа
type MessagePackage struct {
	Type    int    `json:"type"`
	Message []byte `json:"message"`
}

// Message - структура для хранения сообщения
type Message struct {
	Body     string `json:"body"`
	Subject  string `json:"subject"`
	UserFrom string `json:"user_from"`
	UserTo   string `json:"user_to"`
	Type     int    `json:"type"`
}
