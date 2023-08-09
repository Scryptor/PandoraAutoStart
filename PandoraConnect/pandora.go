package PandoraConnect

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"
)

type PConnect struct {
	Client   http.Client
	Login    string
	Password string
	Sid      string
	Lang     string
	Active   bool
	NameF    string
	Cansel   context.CancelFunc
	ctx      context.Context
}

func NewPConnect(login, pass, lang string) *PConnect {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())

	return &PConnect{
		Client:   http.Client{Timeout: 10 * time.Second, Jar: jar},
		Login:    login,
		Password: pass,
		Lang:     lang,
		Cansel:   cancel,
		ctx:      ctx,
	}
}

// Authorize Авторизация на сервере пандоры
func (Pcon *PConnect) Authorize() error {
	data := url.Values{}
	data.Set("login", Pcon.Login)
	data.Set("password", Pcon.Password)
	data.Set("lang", Pcon.Lang)

	req, err := http.NewRequest("POST", "https://p-on.ru/api/users/login", strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 YaBrowser/23.7.1.1140 Yowser/2.5 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("Connection", "keep-alive")
	resp, err := Pcon.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if !json.Valid(body) {
		return errors.New("Incorrect Json " + string(body))
	}
	var PandoraAuthJson PandoraAuthorizeJson
	err = json.Unmarshal(body, &PandoraAuthJson)
	if err != nil {
		return err
	}
	if PandoraAuthJson.Status == "success" {
		fmt.Println("Успешная авторизация на сервере Pandora")
		fmt.Printf("Sid текущей сессии: %s\n", PandoraAuthJson.SessionID)
	} else {
		return errors.New("авторизация на сервере сигнализации не удалась")
	}

	time.Sleep(200 * time.Millisecond)
	err = Pcon.GetSettings()
	if err != nil {
		return err
	}

	// Апдейты состояния авто
	ctxU, _ := context.WithCancel(Pcon.ctx)
	go Pcon.PandoraUpdates(ctxU)
	// Запускаем пинг каждые 30 сек, дабы с эмулировать работу сайта
	ctxL, _ := context.WithCancel(Pcon.ctx)
	go Pcon.IamAlive(ctxL)

	time.Sleep(200 * time.Millisecond)
	err = Pcon.GetUserProfile()
	if err != nil {
		Pcon.Cansel()
		return err
	}
	time.Sleep(200 * time.Millisecond)
	pdList, err := Pcon.GetDevices()
	if err != nil {
		Pcon.Cansel()
		return errors.New("не удалось получить список устройств")
	}
	if len(pdList) == 0 {
		Pcon.Cansel()
		return errors.New("нет подключенных устройств")
	}
	fmt.Printf("Версия прошивки: %s\n"+
		"Имя машины: %s\n"+
		"Телефон сигнализации: %s\n",
		pdList[0].Firmware,
		pdList[0].Name,
		pdList[0].Phone)
	return nil
}

// GetDevices Получение списка устройств
func (Pcon *PConnect) GetDevices() (PandoraDevicesList, error) {
	data := url.Values{}

	req, err := http.NewRequest("GET", "https://p-on.ru/api/devices", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 YaBrowser/23.7.1.1140 Yowser/2.5 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("Connection", "keep-alive")

	resp, err := Pcon.Client.Do(req)
	if err != nil {
		log.Println(err)
		return PandoraDevicesList{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return PandoraDevicesList{}, err
	}
	if !json.Valid(body) {
		log.Println("IncorrectJson", string(body))
		return PandoraDevicesList{}, err
	}
	var PandoraDList PandoraDevicesList
	err = json.Unmarshal(body, &PandoraDList)
	if err != nil {
		log.Println(err)
		log.Println(string(body))
		return PandoraDevicesList{}, err
	}
	return PandoraDList, nil
}

// GetSettings Получаем настройки, тут нет ничего интересно, чисто для эмуляции поведения сайта
func (Pcon *PConnect) GetSettings() error {
	data := url.Values{}

	req, err := http.NewRequest("GET", "https://p-on.ru/api/settings", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 YaBrowser/23.7.1.1140 Yowser/2.5 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("Connection", "keep-alive")

	resp, err := Pcon.Client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	if !json.Valid(body) {
		log.Println("IncorrectJson", string(body))
		return err
	}
	return nil
}

// GetUserProfile Получает профиль пользователя, из него нам нужно только name_f (id сигнализации, чтобы потом ему команды отправлять)
func (Pcon *PConnect) GetUserProfile() error {
	data := url.Values{}

	req, err := http.NewRequest("GET", "https://p-on.ru/api/users/profile", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 YaBrowser/23.7.1.1140 Yowser/2.5 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("Connection", "keep-alive")

	resp, err := Pcon.Client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	if !json.Valid(body) {
		log.Println("IncorrectJson", string(body))
		return err
	}
	if !json.Valid(body) {
		log.Println("IncorrectJson", string(body))
		return err
	}
	var PandoraUProf PandoraUserProfileJson
	err = json.Unmarshal(body, &PandoraUProf)
	if err != nil {
		log.Println(err)
		log.Println(string(body))
		return err
	}
	Pcon.NameF = PandoraUProf.Response.NameF
	fmt.Println("Id сигнализации", Pcon.NameF)
	return nil
}

// StartVehicle Запускаем машину
func (Pcon *PConnect) StartVehicle() error {
	log.Println("Запускаю двигатель")
	data := url.Values{}
	data.Set("id", Pcon.NameF)
	data.Set("command", "4")
	req, err := http.NewRequest("POST", "https://p-on.ru/api/devices/command", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 YaBrowser/23.7.1.1140 Yowser/2.5 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("Connection", "keep-alive")

	resp, err := Pcon.Client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	if !json.Valid(body) {
		log.Println("IncorrectJson", string(body))
		return err
	}
	log.Println(string(body))
	log.Println("Команда на запуск двигателя успешно отправлена")
	return nil
}

// StopVehicle Останавливаем машину
func (Pcon *PConnect) StopVehicle() error {
	log.Println("Останавливаю двигатель")
	data := url.Values{}
	data.Set("id", Pcon.NameF)
	data.Set("command", "8")
	req, err := http.NewRequest("POST", "https://p-on.ru/api/devices/command", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 YaBrowser/23.7.1.1140 Yowser/2.5 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\"Not.A/Brand\";v=\"8\", \"Chromium\";v=\"114\", \"YaBrowser\";v=\"23\"")
	req.Header.Add("Connection", "keep-alive")

	resp, err := Pcon.Client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	if !json.Valid(body) {
		log.Println("IncorrectJson", string(body))
		return err
	}
	log.Println(string(body))
	log.Println("Команда на запуск двигателя успешно отправлена")
	return nil
}
