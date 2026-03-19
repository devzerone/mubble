# mubble 설계 문서

## 1. 개요

### 1.1 프로젝트 이름
**mubble** (MarkDown Bubble)

### 1.2 프로젝트 목적
- 마크다운 형식의 입력을 지원하는 터미널 애플리케이션
- 파일 경로(상대경로)를 자동으로 감지하고 바로 열 수 있는 기능 제공
- 개발자 친화적인 터미널 경험 제공

### 1.3 타겟 사용자
- CLI 도구를 자주 사용하는 개발자
- 마크다운 문서를 자주 작성하는 개발자
- 터미널에서 파일을 빠르게 열고 싶은 사용자

### 1.4 비기능 요구사항
- **성능**: 입력 지연 없는 실시간 렌더링 (100ms 이내)
- **호환성**: Linux, macOS, Windows 지원
- **확장성**: 플러그인 시스템으로 커스텀 명령 추가 가능

---

## 2. 시스템 아키텍처

### 2.1 아키텍처 패턴
**Component-based TUI Architecture**
- 모델-뷰-업데이트 (MVU) 패턴 사용
- 이벤트 기반 아키텍처
- 모듈화된 컴포넌트 구조

### 2.2 기술 스택

#### 핵심 기술
- **언어**: Go 1.21+
- **TUI 프레임워크**: Bubbletea
- **스타일링**: Lipgloss
- **경로 파싱**: 표준 라이브러리 (path/filepath, regexp)

#### 선택 이유
1. **Go**: 단일 바이너리 배포, 빠른 실행 속도, 강력한 동시성
2. **Bubbletea**: 우아한 TUI 프레임워크, 강력한 업데이트 루프
3. **Lipgloss: 터미널 스타일링을 쉽게 처리

---

## 3. 전체 기능 목록

### 3.1 마크다운 입력 ✅ 구현 완료
- ✅ 마크다운 실시간 렌더링
- ✅ 문법 하이라이팅 (헤더, 볼드, 이탤릭, 코드, 링크)
- ✅ 멀티라인 입력 지원
- [ ] 자동 완성 (계획중)
- [ ] 들여쓰기 유지 (계획중)

### 3.2 파일 경로 감지 ✅ 구현 완료
- ✅ 상대경로 패턴 매칭 (./, ../)
- ✅ 자동 하이퍼링크 변환
- ✅ 파일 존재 여부 확인
- ✅ 다양한 경로 패턴 지원 (./, ../, ~/)
- ✅ 라인 번호 포함 경로 (file.md:10)
- ✅ 시각적 표시 (존재: 녹색, 미존재: 빨간색)

### 3.3 빠른 파일 열기 (계획중)
- [ ] 클릭으로 파일 열기
- [ ] 기본 에디터와 연동
- [ ] 커스텀 에디터 설정
- [ ] 라인 번호로 이동 지원

### 3.4 모드 전환 ✅ 구현 완료
- ✅ Ctrl+M으로 터미널 모드 ↔ 마크다운 모드 전환
- ✅ 분할 화면 레이아웃
- ✅ 모드별 키 동작 차별화
- ✅ 실시간 마크다운 미리보기

### 3.5 터미널 명령 실행 (계획중)
- [ ] 쉘 명령 실행
- [ ] 출력 캡처
- [ ] 명령 히스토리
- [ ] 백그라운드 실행 지원

### 3.6 테마/설정 (계획중)
- [ ] 밝은/어두운 테마
- [ ] 컬러 스킴 설정
- [ ] 설정 파일 지원
- [ ] 런타임 테마 전환
- 마크다운 실시간 렌더링
- 문법 하이라이팅
- 자동 완성
- 멀티라인 입력 지원

### 3.2 파일 경로 감지 (file-path-detection)
- 상대경로 패턴 매칭
- 자동 하이퍼링크 변환
- 파일 존재 여부 확인
- 다양한 경로 패턴 지원 (./, ../, ~/)

### 3.3 빠른 파일 열기 (quick-file-open)
- 클릭으로 파일 열기
- 기본 에디터와 연동
- 커스텀 에디터 설정
- 라인 번호로 이동 지원

### 3.4 터미널 명령 실행 (command-execution)
- 쉘 명령 실행
- 출력 캡처
- 명령 히스토리
- 백그라운드 실행 지원

### 3.5 테마/설정 (theming)
- 밝은/어두운 테마
- 컬러 스킴 설정
- 설정 파일 지원
- 런타임 테마 전환

---

## 4. 데이터 모델

### 4.1 핵심 데이터 구조

```go
// 애플리케이션 상태
type AppState struct {
    CurrentInput     string
    RenderedMarkdown string
    DetectedPaths    []FilePath
    CommandHistory   []string
    HistoryIndex     int
    Config           Config
    Mode             Mode // TerminalMode, MarkdownMode
}

// 애플리케이션 모드
type Mode int

const (
    TerminalMode Mode = iota // 일반 터미널 모드
    MarkdownMode              // 마크다운 모드
)

// 파일 경로 정보
type FilePath struct {
    Original    string
    Resolved    string
    Exists      bool
    LineNumber  int
    StartPos    int
    EndPos      int
}

// 설정
type Config struct {
    Editor         string
    Theme          string
    AutoOpen       bool
    PathPatterns   []string
    MaxHistory     int
}
```

---

## 5. 인터페이스 설계

### 5.1 핵심 인터페이스

```go
// 마크다운 렌더러
type MarkdownRenderer interface {
    Render(input string) string
    HighlightSyntax(input string) string
}

// 경로 감지기
type PathDetector interface {
    DetectPaths(input string) []FilePath
    ResolvePath(relPath string) (string, error)
    ValidatePath(path string) bool
}

// 파일 열기
type FileOpener interface {
    OpenFile(path string, line int) error
    GetDefaultEditor() string
}

// 명령 실행기
type CommandExecutor interface {
    Execute(cmd string) (string, error)
    ExecuteInBackground(cmd string) error
}
```

---

## 6. 프로젝트 구조

```
mubble/
├── cmd/
│   └── mubble/
│       └── main.go              # ✅ 진입점
├── internal/
│   ├── ui/
│   │   ├── model.go             # ✅ Bubbletea Model (모드 전환)
│   │   └── markdown/
│   │       └── renderer.go       # ✅ 마크다운 렌더러
│   ├── pathfinder/
│   │   └── detector.go          # ✅ 경로 감지
│   ├── executor/                # 📁 디렉토리만 생성 (계획중)
│   └── config/                  # 📁 디렉토리만 생성 (계획중)
├── docs/
│   ├── README.md                # ✅ 문서 인덱스
│   ├── DESIGN.md                # ✅ 전체 설계
│   ├── specs/                   # ✅ 피처별 스펙
│   │   ├── markdown-input/      # ✅ spec.md, plan.md, acceptance.md, CHANGELOG.md
│   │   ├── file-path-detection/ # ✅ spec.md, CHANGELOG.md
│   │   ├── quick-file-open/     # 📋 spec.md만 있음
│   │   ├── command-execution/   # 📋 spec.md만 있음
│   │   └── theming/             # 📋 spec.md만 있음
│   └── adr/                     # ✅ 아키텍처 의사결정
│       ├── 001-tech-stack.md
│       ├── 002-architecture.md
│       ├── 003-path-detection.md
│       └── 004-markdown-mode.md
├── .trust.yaml                  # ✅ 품질 설정
├── go.mod                       # ✅ 모듈 정의
├── go.sum
├── mubble                       # ✅ 실행 파일
└── README.md                    # ✅ 프로젝트 개요
```

---

## 7. 개발 로드맵

### Phase 1: MVP ✅ 완료
- [x] 기본 TUI 프레임워크 (Bubbletea)
- [x] 마크다운 입력창 구현
- [x] 정교한 경로 감지 (정규식)
- [x] 모드 전환 (Ctrl+M)
- [x] 분할 화면 레이아웃
- [x] 파일 경로 감지 및 렌더링

### Phase 2: 핵심 기능 (진행중)
- [x] 마크다운 실시간 렌더링 ✅
- [x] 정교한 경로 감지 ✅
- [ ] 파일 클릭으로 열기 (다음)
- [ ] 명령 실행 기능 (계획중)
- [ ] 명령 히스토리 (계획중)

### Phase 3: 고급 기능 (계획중)
- [ ] 테마 시스템
- [ ] 설정 파일 지원
- [ ] 커스텀 에디터 연동
- [ ] 자동 완성
- [ ] 자동 완성

### Phase 4: 릴리스 (1주)
- [ ] 테스트 작성
- [ ] 문서 작성
- [ ] 바이너리 배포
- [ ] 플러그인 시스템 (선택)

---

## 8. 품질 기준

```yaml
# .trust.yaml
quality_gates:
  test_coverage:
    min: 80

  code_review:
    required: true

  go_vet:
    enabled: true

  golangci_lint:
    enabled: true
    max_issues: 10
```

---

## 9. 다음 단계

1. **ADR 작성**: 주요 아키텍처 의사결정 문서화
2. **첫 번째 피처 기획**: markdown-input 기능 상세 설계
3. **개발 시작**: Go 프로젝트 초기화 및 기본 구현
