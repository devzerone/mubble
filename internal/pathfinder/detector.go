package pathfinder

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// FilePath는 감지된 파일 경로 정보를 나타냅니다.
type FilePath struct {
	Original    string // 원본 경로
	Resolved    string // 해석된 절대 경로
	Exists      bool   // 파일 존재 여부
	LineNumber  int    // 라인 번호 (있는 경우)
	StartPos    int    // 텍스트 시작 위치
	EndPos      int    // 텍스트 종료 위치
}

// Detector는 파일 경로를 감지합니다.
type Detector struct {
	patterns []*regexp.Regexp
}

// NewDetector는 새로운 경로 감지기를 생성합니다.
func NewDetector() *Detector {
	return &Detector{
		patterns: []*regexp.Regexp{
			// 상대경로: ./file.md, ../file.md
			regexp.MustCompile(`[\./]+[^\s\]]+`),

			// 홈 디렉토리: ~/file.md
			regexp.MustCompile(`~[^\s\]]+`),

			// 파일 확장자: file.md, file.txt
			regexp.MustCompile(`[^\s\]]+\.[a-z]{2,4}(:\d+)?`),

			// 라인 번호 포함: file.md:10
			regexp.MustCompile(`[^\s\]]+\.md:\d+`),
		},
	}
}

// DetectPaths는 텍스트에서 파일 경로를 감지합니다.
func (d *Detector) DetectPaths(input string) []FilePath {
	paths := []FilePath{}

	for _, pattern := range d.patterns {
		matches := pattern.FindAllStringIndex(input, -1)

		for _, match := range matches {
			startPos := match[0]
			endPos := match[1]
			original := input[startPos:endPos]

			// 라인 번호 분리
			lineNumber := 0
			pathOnly := original

			if strings.Contains(original, ":") {
				parts := strings.Split(original, ":")
				if len(parts) == 2 {
					pathOnly = parts[0]
					if num, err := strconv.Atoi(parts[1]); err == nil {
						lineNumber = num
					}
				}
			}

			// 경로 해석
			resolved, err := d.ResolvePath(pathOnly)
			exists := false
			if err == nil {
				exists = d.ValidatePath(resolved)
			}

			paths = append(paths, FilePath{
				Original:   original,
				Resolved:   resolved,
				Exists:     exists,
				LineNumber: lineNumber,
				StartPos:   startPos,
				EndPos:     endPos,
			})
		}
	}

	return paths
}

// ResolvePath는 상대경로를 절대경로로 해석합니다.
func (d *Detector) ResolvePath(relPath string) (string, error) {
	// 홈 디렉토리 확장
	if strings.HasPrefix(relPath, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			return relPath, err
		}
		relPath = filepath.Join(home, relPath[2:])
	}

	// 현재 작업 디렉토리 기준으로 해석
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		return relPath, err
	}

	return absPath, nil
}

// ValidatePath는 파일 존재 여부를 확인합니다.
func (d *Detector) ValidatePath(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}

	// 디렉토리가 아닌 파일만 유효
	return !info.IsDir()
}

// RenderPath는 경로를 적절한 스타일로 렌더링합니다.
func (d *Detector) RenderPath(fp FilePath) string {
	// 나중에 lipgloss 스타일을 적용할 수 있도록 플레이스홀더 반환
	if fp.Exists {
		return "[GREEN]" + fp.Original + "[RESET]"
	}
	return "[RED]" + fp.Original + "[RESET]"
}
