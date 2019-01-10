package nsclient

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

// Client - тип для работы с клиентом
type Client struct {
	host string
	port int
}

// NewClient - функция создает новый клиент для работы с сервисом уведомлений
func NewClient(host string, port int) *Client {
	return &Client{
		host: host,
		port: port,
	}
}

// transferToServer - посылаем сообщение
func transferToServer(host string, port int, messageType int, messageBody []byte) ([]byte, error) {
	// создаем соединение
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, fmt.Errorf("Ошибка создания сендера: %s", err)
	}
	// 	формируем структуру для передачи
	messageStruct, err := CreateMessage(messageType, messageBody)
	if err != nil {
		return nil, fmt.Errorf("Ошибка формирования структуры сообщения: %s", err)
	}

	// 	переводим структуру в байтовый массив
	plainText, err := PackMessage(&messageStruct)
	if err != nil {
		return nil, fmt.Errorf("Ошибка конвертирования структуры: %s", err)
	}

	// шифруем
	cipherText, err := encrypt(plainText)
	if err != nil {
		return nil, fmt.Errorf("Ошибка шифрования: %s", err)
	}

	// 	формируем префикс
	prefix := make([]byte, 4)
	binary.BigEndian.PutUint32(prefix, uint32(len(cipherText)))

	// 	передаем стрим
	_, err = conn.Write(prefix)
	if err != nil {
		// handle error
		return nil, fmt.Errorf("Ошибка передачи по TCP. Ошибка при передаче префикса: %s", err)
	}
	_, err = conn.Write(cipherText)
	if err != nil {
		// handle error
		return nil, fmt.Errorf("Ошибка передачи по TCP. Ошибка при передачи сообщения: %s", err)
	}

	// 	получаем ответ
	// 	читаем префикс
	prefix = make([]byte, 4)
	_, err = io.ReadFull(conn, prefix)
	if err != nil {
		// handle error
		return nil, fmt.Errorf("Ошибка приема по TCP. Ошибка при приема префикса: %s", err)
	}
	fmt.Println(prefix)

	length := binary.BigEndian.Uint32(prefix)
	// verify length if there are restrictions

	message := make([]byte, int(length))
	_, err = io.ReadFull(conn, message)
	if err != nil {
		// handle error
		return nil, fmt.Errorf("Ошибка приема по TCP. Ошибка при приема тела: %s", err)
	}

	// дешифруем
	plainText, err = decrypt(message)
	if err != nil {
		return nil, fmt.Errorf("Ошибка шифрования: %s", err)
	}

	return plainText, nil
}

// CreateMessage - создание сообщения
func CreateMessage(messageType int, messageBody []byte) (MessagePackage, error) {
	// 	времянка для возврата
	var messageStruct MessagePackage
	// 	заполняем
	// 	тип сообщения
	if messageType != MESSAGE_TYPE_ERROR {
		messageStruct.Type = messageType
	} else {
		return messageStruct, errors.New("Неверный тип сообщения")
	}
	// 	тип сообщения
	if len(messageBody) != 0 && len(messageBody) < 950 {
		messageStruct.Message = messageBody
	} else {
		return messageStruct, errors.New("Неверный размер сообщения")
	}

	// 	возврат
	return messageStruct, nil
}

// PackMessage - запаковываем структуру в байтовый массив
func PackMessage(message *MessagePackage) ([]byte, error) {
	//	Маршалим прочитанное в json
	messagePack, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("Кривой Маршал: %s", err)
	}
	return messagePack, nil
}

// UnpackMessage - распаковываем байтовый массив в структуру
func UnpackMessage(messagePack []byte) (MessagePackage, error) {
	// 	времянка для возврата
	var messageStruct MessagePackage

	//	Анмаршалим прочитанное в структуру
	err := json.Unmarshal(messagePack, &messageStruct)
	if err != nil {
		return messageStruct, fmt.Errorf("Кривой анмаршал: %s", err)
	}
	return messageStruct, nil
}
