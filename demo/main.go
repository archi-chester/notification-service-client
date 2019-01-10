package main

import (
	nsclient "gitlab.havana/BDIO/notification-service-client"
)

// 	Клиент для сервиса уведомлений ver. 1.0

//	попытка написать клиент для сервиса сообщений с передачей данных по tcp
func main() {
	log.Info("Клиент для сервиса уведомлений")

	//	var message string
	//	var messageType int

	c := nsclient.NewClient("10.46.2.54", 10445)

	for {
		//		fmt.Println("Введите тип сообщения:")
		//		fmt.Scanln(&message)
		//		messageType, _ = strconv.Atoi(message)
		//		fmt.Println("Введите сообщение:")
		//		fmt.Scanln(&message)

		newMessage := nsclient.Message{
			Body:     "Test message",
			Subject:  "Test Subject",
			UserFrom: "therox",
			UserTo:   "admin",
			Type:     nsclient.MessageTypeObjectlist,
		}
		err := c.Send(newMessage)
		// err := ns_client.Exec(messageType, "10.46.2.54", int(10445), []byte(message))
		if err != nil {
			// handle error
			log.Error("Ошибка отработки Sender", err)
			panic(err.Error())
		}
		break
	}
	// дергаем senderz

	defer logFile.Close()
}
