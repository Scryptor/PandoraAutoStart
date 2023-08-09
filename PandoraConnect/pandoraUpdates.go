package PandoraConnect

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// IamAlive Своеобразный пинг от пандоры, проверяет есть ли коннект
func (Pcon *PConnect) IamAlive(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case _ = <-ticker.C:
			log.Println("Пинг")
			data := url.Values{}
			data.Set("num_click", "0")
			req, err := http.NewRequest("POST", "https://p-on.ru/api/iamalive", strings.NewReader(data.Encode()))
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
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}
			if !json.Valid(body) {
				log.Println("IncorrectJson", string(body))
			}
			var PandoraUProf PandoraPingJson
			err = json.Unmarshal(body, &PandoraUProf)
			if err != nil {
				log.Println(err)
				log.Println(string(body))
			}
			log.Println("Статус пинга: ", PandoraUProf.Status)
		case <-ctx.Done():
			return
		}
	}
}
func (Pcon *PConnect) PandoraUpdates(ctx context.Context) {
	ticker := time.NewTicker(6 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case _ = <-ticker.C:
			log.Println("Пинг")
			data := url.Values{}
			req, err := http.NewRequest("GET", fmt.Sprintf("https://p-on.ru/api/updates?ts=%d", time.Now().Unix()), strings.NewReader(data.Encode()))
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
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
			}
			if !json.Valid(body) {
				log.Println("IncorrectJson", string(body))
			}
			var PandoraUpdates PandoraUpdatesJson
			err = json.Unmarshal(body, &PandoraUpdates)
			if err != nil {
				log.Println(err)
				log.Println(string(body))
			}
			log.Println("Updates: ", PandoraUpdates)
		case <-ctx.Done():
			return
		}
	}
}
