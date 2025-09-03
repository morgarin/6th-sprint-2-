package main

import (
	"fmt"
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	fmt.Println("Запускаем сервер")
	srv, err := server.New(log.Default())
	if err != nil {
		panic(err)
	}
	if err := srv.HttpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Завершаем работу")
}

//создать логгер slog(пакет открыть)
//инициализировать пакет morse
//инициализировать хендлеры
//создать сервер
