package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/shirou/gopsutil/v3/cpu"
)

type AppState struct {
	cpu  sync.Mutex
	data []float64
}

func main() {
	appState := &AppState{}
	app := fiber.New()

	go func() {
		for {
			percentages, _ := cpu.Percent(time.Second, true)
			appState.cpu.Lock()
			appState.data = percentages
			appState.cpu.Unlock()
			time.Sleep(time.Second)
		}
	}()

	app.Get("/", root)
	app.Get("/api/cpu", appState.getCPU)
	app.Get("/realtime/cpu", websocket.New(appState.realTimeCPU))
	app.Static("/static", "static")

	fmt.Println("Listening on: 3000")
	log.Fatal(app.Listen(":3000"))
}

func root(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendFile("templates/index.html")
}

func (s *AppState) getCPU(c *fiber.Ctx) error {
	s.cpu.Lock()
	data := s.data
	s.cpu.Unlock()

	c.Set("Content-Type", "application/json;charset=utf-8")

	return c.JSON(data)
}

func (s *AppState) realTimeCPU(c *websocket.Conn) {
	defer c.Close()

	for {
		s.cpu.Lock()
		payload, _ := json.Marshal(s.data)
		s.cpu.Unlock()

		err := c.WriteMessage(websocket.TextMessage, payload)
		if err != nil {
			log.Printf("Error writing message: %v", err)
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}
