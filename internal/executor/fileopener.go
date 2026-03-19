package executor

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/devzerone/mubble/internal/pathfinder"
)

// FileOpener는 파일을 여는 기능을 제공합니다.
type FileOpener struct {
	defaultEditor string
	customEditor  string
}

// EditorConfig는 에디터 설정을 나타냅니다.
type EditorConfig struct {
	Name    string
	Command  string
	Args    []string
	LineArg string // 라인 번호 인자 형식
}

// NewFileOpener는 새로운 FileOpener를 생성합니다.
func NewFileOpener() *FileOpener {
	return &FileOpener{
		defaultEditor: "",
		customEditor:  "",
	}
}

// OpenFile은 파일을 엽니다.
func (fo *FileOpener) OpenFile(path string, line int) error {
	editor := fo.getEditor()
	cmd := fo.buildCommand(editor, path, line)

	// 터미널에서 실행하므로 백그라운드로 실행
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("📂 열기 중: %s\n", path)
	if line > 0 {
		fmt.Printf("📍 줄 번호: %d\n", line)
	}
	fmt.Printf("✏️  에디터: %s\n\n", editor)

	return cmd.Start()
}

// OpenFileInBackground은 백그라운드에서 파일을 엽니다.
func (fo *FileOpener) OpenFileInBackground(path string, line int) error {
	editor := fo.getEditor()
	cmd := fo.buildCommand(editor, path, line)

	return cmd.Start()
}

// GetDefaultEditor는 기본 에디터를 반환합니다.
func (fo *FileOpener) GetDefaultEditor() string {
	return fo.getEditor()
}

// SetEditor는 커스텀 에디터를 설정합니다.
func (fo *FileOpener) SetEditor(editor string) error {
	fo.customEditor = editor
	return nil
}

// getEditor는 사용할 에디터를 반환합니다.
func (fo *FileOpener) getEditor() string {
	if fo.customEditor != "" {
		return fo.customEditor
	}

	if fo.defaultEditor != "" {
		return fo.defaultEditor
	}

	// 자동 감지
	editor := fo.detectEditor()
	if editor != "" {
		fo.defaultEditor = editor
		return editor
	}

	// 기본값
	return "vi"
}

// detectEditor는 사용 가능한 에디터를 감지합니다.
func (fo *FileOpener) detectEditor() string {
	editors := []EditorConfig{
		{Name: "code", Command: "code", Args: []string{"--goto"}, LineArg: "--goto"},
		{Name: "vim", Command: "vim", Args: []string{}, LineArg: "+"},
		{Name: "vi", Command: "vi", Args: []string{}, LineArg: "+"},
		{Name: "nano", Command: "nano", Args: []string{}, LineArg: "+"},
		{Name: "emacs", Command: "emacs", Args: []string{}, LineArg: "+"},
	}

	for _, editor := range editors {
		if _, err := exec.LookPath(editor.Command); err == nil {
			fmt.Printf("✅ 에디터 감지: %s\n", editor.Name)
			return editor.Command
		}
	}

	fmt.Println("⚠️  에디터를 찾을 수 없습니다. 기본 vi 사용")
	return "vi"
}

// buildCommand는 에디터 실행 명령을 생성합니다.
func (fo *FileOpener) buildCommand(editor string, path string, line int) *exec.Cmd {
	args := []string{}

	// 라인 번호 처리
	if line > 0 {
		// VS Code
		if strings.Contains(editor, "code") {
			args = append(args, "--goto", fmt.Sprintf("%s:%d", path, line))
		} else {
			// Vim, Nano, Emacs 등
			args = append(args, fmt.Sprintf("+%d", line), path)
		}
	} else {
		args = append(args, path)
	}

	return exec.Command(editor, args...)
}

// GetEditorConfig는 에디터 설정을 반환합니다.
func (fo *FileOpener) GetEditorCommand(fp pathfinder.FilePath) (*exec.Cmd, error) {
	editor := fo.getEditor()
	path := fp.Resolved

	if fp.LineNumber > 0 {
		line := fp.LineNumber
		if strings.Contains(editor, "code") {
			return exec.Command(editor, "--goto", fmt.Sprintf("%s:%d", path, line)), nil
		} else {
			return exec.Command(editor, fmt.Sprintf("+%d", line), path), nil
		}
	}

	return exec.Command(editor, path), nil
}
