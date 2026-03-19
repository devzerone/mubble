package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/devzerone/mubble/internal/ui/markdown"
)

// Model은 애플리케이션의 전체 상태를 관리합니다.
type Model struct {
	width       int
	height      int
	quitting    bool
	textInput   string
	markdown    string
	cursor      int
	renderer    *markdown.Renderer
}

// NewInitialModel은 초기 모델을 생성합니다.
func NewInitialModel() Model {
	return Model{
		textInput: "",
		markdown:  "",
		cursor:    0,
		renderer:  markdown.NewRenderer(),
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
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			m.quitting = true
			return m, tea.Quit

		case tea.KeyEnter:
			// 나중에 명령 실행 로직 추가
			// 지금은 줄바꿈만 처리
			if msg.Alt { // Alt+Enter는 멀티라인
				m.textInput += "\n"
				m.cursor++
			}
			return m, nil

		case tea.KeyBackspace:
			if len(m.textInput) > 0 {
				m.textInput = m.textInput[:len(m.textInput)-1]
				if m.cursor > 0 {
					m.cursor--
				}
				m.markdown = m.renderer.Render(m.textInput) // 마크다운 렌더링
			}
			return m, nil

		case tea.KeyRunes:
			// 문자 입력
			m.textInput += string(msg.Runes)
			m.cursor += len(msg.Runes)
			m.markdown = m.renderer.Render(m.textInput) // 마크다운 렌더링
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	}

	return m, nil
}

// View는 UI를 렌더링합니다.
func (m Model) View() string {
	if m.quitting {
		return "종료하는 중...\n"
	}

	// 스타일 정의
	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#86E1EC")).
		Background(lipgloss.Color("#3C3C3C")).
		Padding(0, 2)

	inputStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#1E1E1E")).
		Padding(1)

	markdownStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Padding(1)

	// 레이아웃 구성
	title := titleStyle.Render("mubble - 마크다운 터미널")
	input := inputStyle.Render(m.getInputLine())
	markdownView := markdownStyle.Render(m.getMarkdownPreview())

	// 전체 화면 구성
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"\n",
		input,
		"\n",
		markdownView,
		"\n",
		"Ctrl+C 또는 Esc로 종료",
	)
}

// getInputLine은 입력 라인을 반환합니다.
func (m Model) getInputLine() string {
	if m.textInput == "" {
		return "> 마크다운을 입력하세요..."
	}
	return "> " + m.textInput
}

// getMarkdownPreview는 마크다운 미리보기를 반환합니다.
// 나중에 Goldmark를 사용하여 실제 렌더링을 구현할 것입니다.
func (m Model) getMarkdownPreview() string {
	if m.markdown == "" {
		return "미리보기가 여기에 표시됩니다..."
	}
	return "📄 미리보기:\n" + m.markdown
}
