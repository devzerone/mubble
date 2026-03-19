package markdown

import (
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Renderer는 마크다운을 렌더링합니다.
type Renderer struct{}

// NewRenderer는 새로운 렌더러를 생성합니다.
func NewRenderer() *Renderer {
	return &Renderer{}
}

// Render는 마크다운 텍스트를 렌더링합니다.
func (r *Renderer) Render(input string) string {
	if input == "" {
		return ""
	}

	lines := strings.Split(input, "\n")
	var result strings.Builder

	for _, line := range lines {
		result.WriteString(r.renderLine(line))
		result.WriteString("\n")
	}

	return result.String()
}

// renderLine은 한 줄을 렌더링합니다.
func (r *Renderer) renderLine(line string) string {
	// 헤더 레벨 감지
	if strings.HasPrefix(line, "#") {
		return r.renderHeading(line)
	}

	// 코드 블록
	if strings.HasPrefix(line, "```") {
		return r.renderCodeBlock(line)
	}

	// 인용구
	if strings.HasPrefix(line, ">") {
		return r.renderQuote(line)
	}

	// 리스트
	if strings.HasPrefix(line, "-") || strings.HasPrefix(line, "*") {
		return r.renderList(line)
	}

	// 순서 있는 리스트
	matched, _ := regexp.MatchString(`^\d+\.`, line)
	if matched {
		return r.renderOrderedList(line)
	}

	// 구분선
	if line == "---" || line == "***" {
		return r.renderThematicBreak()
	}

	// 일반 텍스트 (인라인 요소 처리)
	return r.renderInline(line)
}

// renderHeading은 헤딩을 렌더링합니다.
func (r *Renderer) renderHeading(line string) string {
	level := 0
	for _, char := range line {
		if char == '#' {
			level++
		} else {
			break
		}
	}

	if level > 6 {
		level = 6
	}

	text := strings.TrimSpace(line[level:])

	// 헤딩 스타일
	colors := []string{
		"#FF6B6B", // h1 - 빨강
		"#4EC9B0", // h2 - 청록
		"#569CD6", // h3 - 파랑
		"#DCDCAA", // h4 - 노랑
		"#C586C0", // h5 - 보라
		"#9CDCFE", // h6 - 하늘
	}

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(colors[level-1])).
		Bold(true)

	return style.Render(text)
}

// renderCodeBlock은 코드 블록을 렌더링합니다.
func (r *Renderer) renderCodeBlock(line string) string {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DCDCAA")).
		Background(lipgloss.Color("#1E1E1E")).
		Padding(1)

	// 백틱 3개 제거
	backticks := strings.Repeat("`", 3)
	cleanLine := strings.TrimPrefix(line, backticks)
	cleanLine = strings.TrimSuffix(cleanLine, backticks)

	return style.Render(cleanLine)
}

// renderQuote는 인용구를 렌더링합니다.
func (r *Renderer) renderQuote(line string) string {
	text := strings.TrimPrefix(line, ">")
	text = strings.TrimSpace(text)

	quoteStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888"))

	return quoteStyle.Render("| ") + text
}

// renderList는 리스트를 렌더링합니다.
func (r *Renderer) renderList(line string) string {
	text := strings.TrimPrefix(line, "-")
	text = strings.TrimPrefix(text, "*")
	text = strings.TrimSpace(text)

	bulletStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#4EC9B0"))
	textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA"))

	return bulletStyle.Render("* ") + textStyle.Render(r.renderInline(text))
}

// renderOrderedList는 순서 있는 리스트를 렌더링합니다.
func (r *Renderer) renderOrderedList(line string) string {
	re := regexp.MustCompile(`^(\d+)\.\s*(.*)`)
	matches := re.FindStringSubmatch(line)

	if len(matches) == 3 {
		num := matches[1]
		text := matches[2]

		numStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#4EC9B0"))
		textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FAFAFA"))

		return numStyle.Render(num+". ") + textStyle.Render(r.renderInline(text))
	}

	return line
}

// renderThematicBreak는 구분선을 렌더링합니다.
func (r *Renderer) renderThematicBreak() string {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color("#444"))
	return style.Render("---")
}

// renderInline은 인라인 요소를 렌더링합니다.
func (r *Renderer) renderInline(text string) string {
	// 볼드: **text** or __text__
	text = r.renderBold(text)

	// 이탤릭: *text* or _text_
	text = r.renderItalic(text)

	// 인라인 코드: `text`
	text = r.renderInlineCode(text)

	// 링크: [text](url)
	text = r.renderLink(text)

	return text
}

// renderBold는 볼드 텍스트를 렌더링합니다.
func (r *Renderer) renderBold(text string) string {
	re := regexp.MustCompile(`\*\*(.+?)\*\*|__(.+?)__`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		// 내용 추출
		content := strings.TrimPrefix(strings.TrimSuffix(match, "**"), "**")
		content = strings.TrimPrefix(strings.TrimSuffix(content, "__"), "_")

		style := lipgloss.NewStyle().Bold(true)
		return style.Render(content)
	})
}

// renderItalic은 이탤릭 텍스트를 렌더링합니다.
func (r *Renderer) renderItalic(text string) string {
	// 이미 볼드로 처리된 부분은 건너뛰기
	re := regexp.MustCompile(`(?<!\*)\*(?!\*)(.+?)(?<!\*)\*(?!\*)|(?<!_)_(?!_)(.+?)(?<!_)_(?!_)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		content := strings.Trim(match, "*_")
		style := lipgloss.NewStyle().Italic(true)
		return style.Render(content)
	})
}

// renderInlineCode는 인라인 코드를 렌더링합니다.
func (r *Renderer) renderInlineCode(text string) string {
	re := regexp.MustCompile("`.+?`")
	return re.ReplaceAllStringFunc(text, func(match string) string {
		content := strings.Trim(match, "`")
		style := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF6B6B")).
			Background(lipgloss.Color("#2C2C2C")).
			Padding(0, 1)
		return style.Render(content)
	})
}

// renderLink는 링크를 렌더링합니다.
func (r *Renderer) renderLink(text string) string {
	re := regexp.MustCompile(`\[([^\]]+)\]\(([^)]+)\)`)
	return re.ReplaceAllStringFunc(text, func(match string) string {
		submatches := re.FindStringSubmatch(match)
		if len(submatches) == 3 {
			linkText := submatches[1]
			url := submatches[2]

			style := lipgloss.NewStyle().
				Foreground(lipgloss.Color("#569CD6")).
				Underline(true)

			return style.Render(linkText) + lipgloss.NewStyle().Faint(true).Render(" ("+url+")")
		}
		return match
	})
}
