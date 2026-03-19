package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/devzerone/mubble/internal/ui"
)

func main() {
	// 초기 모델 생성
	model := ui.NewInitialModel()

	// Bubbletea 프로그램 시작
	p := tea.NewProgram(
		model,
		tea.WithAltScreen(),       // 대체 화면 모드
		tea.WithMouseCellMotion(), // 마우스 모션 지원
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("알 수 없는 오류: %v", err)
		os.Exit(1)
	}
}
