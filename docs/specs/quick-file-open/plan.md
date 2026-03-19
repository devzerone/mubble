# Quick File Open Implementation Plan

## 1. 기술적 접근 방식

### 1.1 아키텍처
Bubbletea의 마우스 이벤트 처리를 활용하여 파일 열기 기능을 구현합니다.

```
Mouse Click Event
    ↓
PathDetector detects path
    ↓
FileOpener opens file
    ↓
Editor launches
```

### 1.2 핵심 라이브러리
- **os/exec**: 외부 프로세스 실행 (에디터)
- **Bubbletea**: 마우스 이벤트 처리
- **Lipgloss**: 클릭 가능한 스타일 적용

## 2. 구현 태스크

### Phase 1: 기본 파일 열기 (1일)
- [ ] FileOpener 인터페이스 구현
- [ ] 기본 에디터 감지 (code, vim, nano)
- [ ] 단순 파일 열기 구현
- [ ] 에디터 프로세스 실행

### Phase 2: 라인 번호 지원 (1일)
- [ ] 라인 번호 파싱
- [ ] 에디터별 라인 번호 옵션
- [ ] VS Code --goto 지원
- [ ] Vim +line 지원

### Phase 3: 마우스 클릭 처리 (1일)
- [ ] Bubbletea 마우스 이벤트 리스너
- [ ] 클릭 위치 감지
- [ ] 경로 영역 매핑
- [ ] 클릭 시 파일 열기

### Phase 4: 커스텀 에디터 (1일)
- [ ] 설정 파일 파싱
- [ ] 커스텀 에디터 실행
- [ ] 에디터 인자 지원
- [ ] 기본 에디터 설정

### Phase 5: 터미널 포커스 유지 (1일)
- [ ] 백그라운드 실행
- [ ] 터미널 포커스 유지
- [ ] 프로세스 관리
- [ ] 에러 처리

## 3. 파일 구조

```
internal/
├── executor/
│   ├── fileopener.go        # FileOpener 구현
│   ├── editor.go            # 에디터별 명령 생성
│   └── process.go           # 프로세스 관리
└── config/
    ├── config.go            # 설정 로드
    └── default.yaml         # 기본 설정
```

## 4. 데이터 모델

```go
type FileOpener struct {
    defaultEditor string
    customEditor  string
    editorArgs    []string
}

type EditorConfig struct {
    Name    string
    Command  string
    Args    []string
    LineArg string // 라인 번호 인자 형식
}

type ClickablePath struct {
    Path       pathfinder.FilePath
    StartPos   int
    EndPos     int
}
```

## 5. API 설계

```go
// FileOpener 인터페이스
type FileOpener interface {
    OpenFile(path string, line int) error
    OpenFileInBackground(path string, line int) error
    GetDefaultEditor() string
    SetEditor(editor string) error
}

// Editor 명령 생성
func (fo *FileOpener) GetEditorCommand(path string, line int) (*exec.Cmd, error)

// 클릭 가능한 경로 생성
func MakeClickable(path string, startPos, endPos int) ClickablePath
```

## 6. 에디터 지원

### 자동 감지 에디터
```go
var defaultEditors = []EditorConfig{
    {Name: "VS Code", Command: "code", Args: []string{"--goto"}, LineArg: "--goto"},
    {Name: "Vim", Command: "vim", Args: []string{}, LineArg: "+"},
    {Name: "Nano", Command: "nano", Args: []string{}, LineArg: "+"},
    {Name: "Emacs", Command: "emacs", Args: []string{}, LineArg: "+"},
}
```

### 명령 생성 예시
```bash
# VS Code with line number
code --goto file.md:25

# Vim with line number
vim +25 file.md

# Nano with line number
nano +25 file.md
```

## 7. 테스트 전략

### 단위 테스트
- [ ] 에디터 감지 로직 테스트
- [ ] 명령 생성 로직 테스트
- [ ] 라인 번호 파싱 테스트

### 통합 테스트
- [ ] 파일 열기 테스트
- [ ] 백그라운드 실행 테스트
- [ ] 에러 처리 테스트

### 수동 테스트
- [ ] 마우스 클릭으로 파일 열기
- [ ] 라인 번호로 이동 확인
- [ ] 다양한 에디터 테스트

## 8. 예상 작업 기간

- **총 작업 시간**: 5일
- **개발자**: 1명
- **종속 작업**: file-path-detection (완료됨)

## 9. 위험 요소 및 완화 계획

### 위험 1: 마우스 이벤트 처리 복잡성
- **원인**: Bubbletea 마우스 이벤트 API 이해 필요
- **완화**: 간단한 키보드 단축키 먼저 구현

### 위험 2: 에디터 호환성
- **원인**: 다양한 에디터의 라인 번호 옵션
- **완화**: 주요 에디터(VS Code, Vim)만 먼저 지원

### 위험 3: 백그라운드 실행
- **원인**: 터미널 포커스 유지 복잡
- **완화**: 포어그라운드 실행으로 시작

## 10. 롤백 계획

구현이 실패할 경우:
- **Phase 1**: 키보드 단축키만 지원
- **Phase 2**: 라인 번호 없이 파일만 열기
- **Phase 3**: VS Code만 지원

## 11. 성공 기준

- [ ] 모든 Phase 1-5 태스크 완료
- [ ] 단위 테스트 커버리지 80% 이상
- [ ] 최소 3개 에디터 지원 (VS Code, Vim, Nano)
- [ ] 마우스 클릭으로 파일 열기 동작
- [ ] acceptance.md의 모든 시나리오 통과
