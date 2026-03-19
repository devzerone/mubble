# ADR 003: 파일 경로 감지 전략

## 상태
구현 완료

## 날짜
2026-03-19
- 생성: 2026-03-19
- 구현 완료: 2026-03-19

## 문맥
마크다운 텍스트에서 파일 경로를 자동으로 감지하고, 사용자가 클릭하여 바로 열 수 있게 해야 합니다.

## 결정
**정규식 기반의 실시간 경로 감지**를 사용합니다.

## 상세 설계

### 1. 경로 패턴 정의

```go
// 지원할 경로 패턴
var pathPatterns = []string{
    // 상대 경로
    `\./[^\s]+`,           // ./file.md
    `\.\./[^\s]+`,         // ../file.md
    `[^\s]+\.[a-z]{2,4}`,  // file.md, file.txt

    // 홈 디렉토리
    `~/[^\s]+`,            // ~/Documents/file.md

    // 라인 번호 포함
    `[^\s]+\.md:\d+`,      // file.md:10

    // URL (마크다운 링크)
    `\[[^\]]+\]\([^\)]+\)`, // [text](url)
}
```

### 2. 감지 프로세스

```
입력 텍스트
    ↓
정규식 매칭
    ↓
경로 추출
    ↓
파일 존재 확인
    ↓
하이퍼링크로 변환
    ↓
클릭 가능한 상태로 표시
```

### 3. 파일 존재 확인

```go
type FilePath struct {
    Original   string  // 원본 경로
    Resolved   string  // 해석된 절대 경로
    Exists     bool    // 파일 존재 여부
    LineNumber int     // 라인 번호 (있는 경우)
    StartPos   int     // 텍스트 시작 위치
    EndPos     int     // 텍스트 종료 위치
}

func ResolvePath(relPath string) (FilePath, error) {
    // 1. 홈 디렉토리 확장
    if strings.HasPrefix(relPath, "~/") {
        home, _ := os.UserHomeDir()
        relPath = filepath.Join(home, relPath[2:])
    }

    // 2. 현재 작업 디렉토리 기준으로 해석
    absPath, err := filepath.Abs(relPath)
    if err != nil {
        return FilePath{}, err
    }

    // 3. 파일 존재 확인
    exists := fileExists(absPath)

    return FilePath{
        Original: relPath,
        Resolved: absPath,
        Exists:   exists,
    }, nil
}
```

### 4. 렌더링 전략

```go
// 존재하는 파일: 녹색 + 밑줄
// 존재하지 않는 파일: 빨간색
// 라인 번호 포함: 파일명 + 라인 번호 강조

func RenderPath(fp FilePath) string {
    style := lipgloss.NewStyle()

    if fp.Exists {
        style.Foreground(lipgloss.Color("86"))  // 녹색
        style.Underline(true)
    } else {
        style.Foreground(lipgloss.Color("203")) // 빨간색
    }

    return style.Render(fp.Original)
}
```

## 대안 고려

### 옵션 1: 정규식 기반 감지 ✅ 선택
**장점**
- 빠른 성능
- 유연한 패턴 매칭
- 구현 단순

**단점**
- 복잡한 경로는 미감지 가능

### 옵션 2: AST 기반 파싱
**장점**
- 정확한 마크다운 파싱

**단점**
- 과도한 엔지니어링
- 느린 성능
- 실시간 감지에 부적합

### 옵션 3: 머신러닝 기반
**장점**
- 높은 정확도

**단점**
- 과도한 복잡성
- 느린 성능
- 추가 의존성

## 결정 근거

1. **성능**: 정규식은 매우 빠르며 실시간 감지에 적합
2. **단순성**: 구현과 유지보수가 쉬움
3. **충분한 정확도**: 대부분의 일반적인 경로 패턴을 감지 가능
4. **확장성**: 필요한 경우 새로운 패턴을 쉽게 추가

## 최적화 전략

1. **디바운싱**: 입력 중에는 감지를 지연
2. **캐싱**: 이미 감지한 경로는 캐시에 저장
3. **비동기 처리**: 파일 존재 확인은 별도 goroutine에서 수행

```go
// 디바운싱 예시
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        if msg.Type == tea.KeyRunes {
            m.input += string(msg.Runes)
            // 디바운스 타이머 설정
            return m, debounce(m.detectPaths, 300*time.Millisecond)
        }
    }
    return m, nil
}
```

## 결과
- 정규식을 사용한 빠르고 정확한 경로 감지
- 실시간으로 경로를 하이퍼링크로 변환
- 파일 존재 여부를 시각적으로 표시
- 클릭하여 바로 파일 열기 가능
