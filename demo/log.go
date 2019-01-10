package main

import (
	"os"

	"github.com/Sirupsen/logrus"
)

const (
	//	логи
	LOG_FILE_NAME = "ns.log"
	LOG_FILE_DIR  = "./"
)

// 	создаем переменные
// 	переменная логов
var log = logrus.New()

// 	файл логов
var logFile *os.File

// 	инитим
func init() {

	// 	пробуем открыть файл
	logFile, err := os.OpenFile(LOG_FILE_DIR+LOG_FILE_NAME, os.O_RDWR, 0666)
	if err != nil {
		//	Файла нет - создаем
		logFile, err = os.Create(LOG_FILE_DIR + LOG_FILE_NAME)
		if err != nil {
			log.Error(err)
			return
		}
	}
	// привязываем вывод к файлу логов
	log.Out = logFile

	log.Info("Инициализация логов")
}
