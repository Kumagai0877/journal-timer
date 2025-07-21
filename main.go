package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func loadThemes(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var themes []string
	if err := json.Unmarshal(data, &themes); err != nil {
		return nil, err
	}
	return themes, nil
}

func playSound() {
	switch runtime.GOOS {
	case "darwin":
		exec.Command("afplay", "/System/Library/Sounds/Glass.aiff").Run()
	case "linux":
		exec.Command("aplay", "/usr/share/sounds/alsa/Front_Center.wav").Run()
	default:
		fmt.Println("🔇 音声通知は未対応のOSです")
	}
}

func main() {
	// ランダムにテーマ選択
	rand.New(rand.NewSource(time.Now().UnixNano()))
	themeFile := flag.String("f", "themes.json", "テーマファイルのパス")

	themes, err := loadThemes(*themeFile)
	if err != nil {
		fmt.Println("❌ テーマ読み込みエラー:", err)
		os.Exit(1)
	}
	if len(themes) == 0 {
		fmt.Println("⚠️ テーマが空です")
		os.Exit(1)
	}

	// タイマー開始
	var minutes int
	flag.IntVar(&minutes, "t", 10, "タイマー時間（分）")
	flag.Parse()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	selected := themes[r.Intn(len(themes))]

	fmt.Println("📝 今日のテーマ:", selected)

	duration := time.Duration(minutes) * time.Minute
	fmt.Printf("⏳ タイマー開始（%d分）...\n", minutes)

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	done := time.After(duration)
	remaining := int(duration.Seconds())

	for {
		select {
		case <-ticker.C:
			remaining--
			fmt.Printf("\r残り: %02d:%02d", remaining/60, remaining%60)
		case <-done:
			fmt.Println("\n🔔 終了！")
			playSound()
			return
		}
	}
}
