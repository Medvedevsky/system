package main

import (
	"fmt"
	"math/rand"
	"os/exec"
	"runtime"
	"time"

	"github.com/TheTitanrain/w32"
	"github.com/micmonay/keybd_event"
)

var websites = []string{
	"https://pixel.one",
	"https://lk.clubpixel.ru/dashboard",
	"https://ln1034.listokcrm.ru/",
	"https://www.google.com/",
	"https://dzen.ru/",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		doRandomActivity()
		time.Sleep(10 * time.Second)
	}
}

func doRandomActivity() {
	// Примерно по 35% шанс для каждого действия
	roll := rand.Intn(100)
	switch {
	case roll < 35:
		moveMouseRandom()
	case roll < 70:
		typeRandomKey()
	default:
		openRandomWebsite()
	}
}

func moveMouseRandom() {
	x := rand.Intn(800) + 100
	y := rand.Intn(600) + 100
	w32.SetCursorPos(x, y)
	fmt.Println("🖱 Мышь перемещена в:", x, y)
}

func typeRandomKey() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Println("❌ Ошибка клавиатуры:", err)
		return
	}

	keys := []int{
		keybd_event.VK_ENTER, keybd_event.VK_SPACE,
		keybd_event.VK_A, keybd_event.VK_S,
		keybd_event.VK_D, keybd_event.VK_F,
		keybd_event.VK_TAB,
	}

	key := keys[rand.Intn(len(keys))]
	kb.SetKeys(key)

	if err := kb.Launching(); err != nil {
		fmt.Println("❌ Ошибка нажатия клавиши:", err)
		return
	}
	fmt.Printf("⌨ Нажата клавиша: %d\n", key)
}

func openRandomWebsite() {
	url := websites[rand.Intn(len(websites))]
	fmt.Println("🌐 Открытие сайта:", url)
	openBrowser(url)
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
