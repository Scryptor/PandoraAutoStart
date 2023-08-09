package main

import (
	"PandoraConnectGo/PandoraConnect"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Изменить на свои
const login = "1079312388"
const pass = "n8V8xqbF"

func main() {
	// Почему я не храню объект пандоры? Я не знаю, сколько у них может длиться сессия, проще каждый раз переконнект делать, тогда точно будет работать.
	mux := http.NewServeMux()
	mux.HandleFunc("/start", HandleVehicleStarter)
	mux.HandleFunc("/stop", HandleVehicleStopper)
	// так же можно сделать команду стоп двигатель, проверку баланса и т.п., но моя задача была просто запустить через Алису, поэтому если кому надо, то сам или ко мне за прайс
	fmt.Println("Сервер запущен по адресу http://127.0.0.1:4443, для запуска машины https://127.0.0.1:4443/start")
	log.Fatal("Сервер упал: ", http.ListenAndServe(":4443", mux))
}

func HandleVehicleStarter(w http.ResponseWriter, r *http.Request) {
	go PandoraStarter()
	// Запускает движок, советую добавить какой-нибудь токен, чтобы любой посторонний человек не смог запустить авто, зная просто url
	w.Header().Set("Content-Type", "application/json")
	warn, _ := json.Marshal("Отправлена команда на запуск")
	_, _ = w.Write(warn)
}

func HandleVehicleStopper(w http.ResponseWriter, r *http.Request) {
	go PandoraStopper()
	w.Header().Set("Content-Type", "application/json")
	warn, _ := json.Marshal("Отправлена команда на остновку")
	_, _ = w.Write(warn)
}

func PandoraStarter() {
	Pandora := PandoraConnect.NewPConnect(login, pass, "ru")
	err := Pandora.Authorize()
	if err != nil {
		log.Println("Не могу подключиться к серверу", err)
	}
	// Тут тоже, при желании можно отрефакторить, и сделать хоть с записью пусков в базу данных и выводом чего угодно, но я, как говорил этой цели не преследую, мне нужен был фуникционал, я его получил ), другие проекты ждут, и так все утро потратил на это
	_ = Pandora.StartVehicle()

	time.Sleep(2 * time.Minute)
	Pandora.Cansel()
}
func PandoraStopper() {
	Pandora := PandoraConnect.NewPConnect(login, pass, "ru")
	err := Pandora.Authorize()
	if err != nil {
		log.Println("Не могу подключиться к серверу", err)
	}
	_ = Pandora.StopVehicle()

	time.Sleep(2 * time.Minute)
	Pandora.Cansel()
}
