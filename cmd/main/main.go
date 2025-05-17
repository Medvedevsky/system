package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"runtime"
	"time"

	"github.com/TheTitanrain/w32"
)

var websites = []string{
	"https://lk.clubpixel.ru/dashboard",
	"https://ln1034.listokcrm.ru/",
	"https://c1.saas.infomaximum.com/145ac4ace73c43ad80ce71fc583a4d4a/persons",
	"https://lk.clubpixel.ru/marketplace/manage",
	"https://lk.clubpixel.ru/reports",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		openWebsiteAndEmulateActivity()
		time.Sleep(30 * time.Second)
	}
}

func openWebsiteAndEmulateActivity() {
	url := websites[rand.Intn(len(websites))]
	fmt.Println("🌐 Переход по ссылке:", url)
	openBrowser(url)

	// Подождём немного, чтобы страница успела открыться
	time.Sleep(5 * time.Second)

	fmt.Println("🖱 Эмуляция активности на странице в течение 15 секунд...")
	start := time.Now()
	for time.Since(start) < 15*time.Second {
		moveMouseRandom()
		time.Sleep(500 * time.Millisecond)
	}
}

func moveMouseRandom() {
	x := rand.Intn(800) + 100
	y := rand.Intn(600) + 100
	w32.SetCursorPos(x, y)
	fmt.Println("↔ Мышь перемещена в:", x, y)
}

func openBrowser(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default: // Linux
		cmd = exec.Command("xdg-open", url)
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("❌ Ошибка открытия сайта:", err)
	}
}
