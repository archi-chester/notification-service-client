package nsclient

import (
	"encoding/json"
	"fmt"
)

// Send - функция посылает сообщение на сервис уведомлений
func (c *Client) Send(msg Message) error {

	messagePack, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("Ошибка упаковки сообщения: %s", err)
	}

	_, err = transferToServer(c.host, c.port, MESSAGE_TYPE_INSERT, messagePack)
	if err != nil {
		// handle error
		return fmt.Errorf("Ошибка передачи по TCP: %s", err)
	}
	// 	возврат ошибки
	return nil
}
