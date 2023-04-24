# Go CPU Monitor

A real-time CPU usage monitoring web application built with Go, Fiber, WebSockets, and Preact.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Built With](#built-with)

## Installation

1. Clone the repository:
   git clone https://github.com/your_username/cpu-monitor-go.git

2. Install the required Go packages:
   go get -u github.com/gofiber/fiber/v2
   go get -u github.com/gofiber/websocket/v2
   go get -u github.com/shirou/gopsutil/v3

## Usage

1. Run the application:
   go run main.go

2. Open a web browser and navigate to `http://localhost:3000`.

## Built With

- [Go](https://golang.org/) - The programming language used
- [Fiber](https://github.com/gofiber/fiber) - The web framework used
- [WebSockets](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API) - For real-time communication
- [Preact](https://preactjs.com/) - The lightweight JavaScript library used for rendering
