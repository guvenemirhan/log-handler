package api

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"net/http"
	"real-time-log-analyze/internal/database"
	"real-time-log-analyze/pkg/models"
	"regexp"
	"strconv"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var logRegex = regexp.MustCompile(`(\d+\.\d+\.\d+\.\d+)\s+-\s+-\s+\[(.*?)\]\s+"(GET|POST|PUT|DELETE)\s+(.*?)\s+HTTP\/\d\.\d"\s+(\d+)\s+(\d+)`)

func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	db, err := database.InitializeDB()
	if err != nil {
		fmt.Println("Error initializing the database:", err)
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Received:", string(p))
		analyzeLogLine(string(p), db)
	}
}

func analyzeLogLine(line string, db *gorm.DB) {
	match := logRegex.FindStringSubmatch(line)
	if len(match) == 0 {
		fmt.Println("Invalid Log line:", line)
		return
	}

	ipAddress := match[1]
	rawTime := match[2]
	method := match[3]
	url := match[4]
	status, err := strconv.Atoi(match[5])
	if err != nil {
		fmt.Println("Conversion error:", err)
	}
	latency, err := strconv.ParseFloat(match[6], 64)
	if err != nil {
		fmt.Println("Conversion error:", err)
	}

	parsedTime, err := time.Parse("02/Jan/2006:15:04:05 -0700", rawTime)
	if err != nil {
		fmt.Println("Could not parse time:", err)
		return
	}

	entry := models.LogEntry{
		IPAddress: ipAddress,
		Timestamp: parsedTime,
		URL:       url,
		Method:    method,
		Status:    status,
		Latency:   latency,
	}
	db.Create(&entry)
}
