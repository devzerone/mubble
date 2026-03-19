package ui

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/devzerone/mubble/internal/executor"
	"github.com/devzerone/mubble/internal/ui/markdown"
)

// Mode는 애플리케이션 모드를 나타냅니다.
type Mode int

const (
	TerminalMode Mode = iota // 일반 터미널 모드
	MarkdownMode              // 마크다운 모드
)

// Model은 애플리케이션의 전체 상태를 관리합니다.
type Model struct {
	width      int
	height     int
	quitting   bool
	mode       Mode
	textInput  string
	markdown   string
	cursor     int
	renderer   *markdown.Renderer
	fileOpener *executor.FileOpener
	statusMsg  string
	statusType int // 0: info, 1: success, 2: error
}

// NewInitialModel은 초기 모델을 생성합니다.
func NewInitialModel() Model {
	return Model{
		mode:       TerminalMode,
		textInput:  "",
		markdown:   "",
		cursor:     0,
		renderer:   markdown.NewRenderer(),
		fileOpener: executor.NewFileOpener(),
		statusMsg:  "",
		statusType: 0,
	}
}

// Init는 Bubbletea의 초기화 함수입니다.
func (m Model) Init() tea.Cmd {
	return nil
}

// Update는 이벤트를 처리하고 상태를 업데이트합니다.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		{
			// Ctrl+C: 항상 종료
			if msg.Type == tea.KeyCtrlC {
				m.quitting = true
				return m, tea.Quit
			}

			// Ctrl+M: 모드 전환
			if msg.Type == tea.KeyCtrlM {
				if m.mode == TerminalMode {
					m.mode = MarkdownMode
				} else {
					m.mode = TerminalMode
				}
				return m, nil
			}

			// Esc: 모드별 다른 동작
			if msg.Type == tea.KeyEsc {
				if m.mode == MarkdownMode {
					m.mode = TerminalMode
					return m, nil
				}
				m.quitting = true
				return m, tea.Quit
			}

			// Enter: 모드별 다른 동작
			if msg.Type == tea.KeyEnter {
				if m.mode == MarkdownMode {
					m.textInput += "\n"
					m.cursor++
					m.markdown = m.renderer.Render(m.textInput)
				} else {
					m.textInput = ""
					m.cursor = 0
					m.markdown = ""
				}
				return m, nil
			}

			// Backspace: 모든 모드 동일
			if msg.Type == tea.KeyBackspace {
				if len(m.textInput) > 0 {
					m.textInput = m.textInput[:len(m.textInput)-1]
					if m.cursor > 0 {
						m.cursor--
					}
					m.markdown = m.renderer.Render(m.textInput)
				}
				return m, nil
			}

			// 문자 입력
			if msg.Type == tea.KeyRunes {
				m.textInput += string(msg.Runes)
				m.cursor += len(msg.Runes)
				m.markdown = m.renderer.Render(m.textInput)
				return m, nil
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.MouseMsg:
		// 마우스 클릭 처리
		if msg.Type == tea.MouseLeft {
			// 파일 열기 시도
			if m.mode == MarkdownMode {
				return m.handleMouseClick(msg)
			}
		}
		return m, nil

	case clearStatusMsg:
		m.statusMsg = ""
		m.statusType = 0
		return m, nil
	}

	return m, nil
}

// handleMouseClick은 마우스 클릭을 처리합니다.
func (m Model) handleMouseClick(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	// 현재는 간단하게 상태 메시지만 표시
	// 나중에 실제 클릭 위치 계산 구현
	m.statusMsg = "🖱️ 마우스 클릭 감지됨"
	m.statusType = 1

	// 파일 열기 시�레 실행 (데모)
	// 실제로는 클릭된 위치의 경로를 감지해야 함
	// 지금은 상태 메시지만 표시
	return m, tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
		return clearStatusMsg{}
	})
}

// clearStatusMsg는 상태 메시지를 지우는 메시지입니다.
type clearStatusMsg struct{}

// View는 UI를 렌더링합니다.
func (m Model) View() string {
	if m.quitting {
		return "종료하는 중...\n"
	}

	if m.mode == MarkdownMode {
		return m.renderMarkdownMode()
	}

	return m.renderTerminalMode()
}

// renderTerminalMode는 일반 터미널 모드를 렌더링합니다.
func (m Model) renderTerminalMode() string {
	// 스타일 정의
	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#86E1EC")).
		Background(lipgloss.Color("#3C3C3C")).
		Padding(0, 2)

	inputStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#1E1E1E")).
		Padding(1)

	// 레이아웃 구성
	title := titleStyle.Render("mubble - 터미널 모드 (Ctrl+M: 마크다운 모드)")
	input := inputStyle.Render(m.getTerminalInput())

	// 전체 화면 구성
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"\n",
		input,
		"\n",
		"Ctrl+C: 종료 | Ctrl+M: 마크다운 모드",
	)
}

// renderMarkdownMode는 마크다운 모드를 렌더링합니다.
func (m Model) renderMarkdownMode() string {
	// 스타일 정의
	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#4EC9B0")).
		Background(lipgloss.Color("#3C3C3C")).
		Padding(0, 2)

	inputStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#1E1E1E")).
		Padding(1).
		Width(m.width / 2)

	previewStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#252525")).
		Padding(1).
		Width(m.width / 2).
		Height(m.height - 5) // 제목과 도움말 공간 제외

	// 레이아웃 구성
	title := titleStyle.Render("mubble - 마크다운 모드 (Ctrl+M: 터미널 모드)")
	input := inputStyle.Render(m.getMarkdownInput())
	preview := previewStyle.Render(m.getMarkdownPreview())

	// 분할 화면 구성
	topSection := lipgloss.JoinHorizontal(lipgloss.Top, input, preview)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"\n",
		topSection,
		"\n",
		m.getStatusBar(),
	)
}

// getTerminalInput은 터미널 모드 입력 라인을 반환합니다.
func (m Model) getTerminalInput() string {
	if m.textInput == "" {
		return "$ 명령을 입력하세요..."
	}
	return "$ " + m.textInput
}

// getMarkdownInput은 마크다운 모드 입력 라인을 반환합니다.
func (m Model) getMarkdownInput() string {
	if m.textInput == "" {
		return "> 마크다운을 입력하세요...\n"
	}
	lines := ""
	for i, line := range splitLines(m.textInput) {
		if i == 0 {
			lines += "> " + line + "\n"
		} else {
			lines += "  " + line + "\n"
		}
	}
	return lines
}

// getMarkdownPreview는 마크다운 미리보기를 반환합니다.
func (m Model) getMarkdownPreview() string {
	if m.markdown == "" {
		return "📄 미리보기가 여기에 표시됩니다..."
	}
	return m.markdown
}

// splitLines는 텍스트를 라인으로 분리합니다.
func splitLines(text string) []string {
	return strings.Split(text, "\n")
}

// getStatusBar는 상태바를 반환합니다.
func (m Model) getStatusBar() string {
	if m.statusMsg == "" {
		return "Ctrl+C: 종료 | Ctrl+M/Esc: 터미널 모드 | Enter: 새 줄"
	}

	style := lipgloss.NewStyle()
	if m.statusType == 1 { // success
		style = style.Foreground(lipgloss.Color("#4EC9B0")) // 녹색
	} else if m.statusType == 2 { // error
		style = style.Foreground(lipgloss.Color("#FF6B6B")) // 빨간색
	}

	return style.Render("│ " + m.statusMsg + " │ Ctrl+C: 종료 | Ctrl+M: 터미널 모드 | Enter: 새 줄")
}

// clearStatus는 상태 메시지를 지웁니다.
func (m Model) clearStatus() tea.Msg {
	return clearStatusMsg{}
}
