# Markdown Input Implementation Plan

## 1. 기술적 접근 방식

### 1.1 아키텍처
Bubbletea의 MVU 패턴을 사용하여 마크다운 입력 컴포넌트를 구현합니다.

```
TextInput (Model)
    ↓
Update (Event Handler)
    ↓
Render (View)
    ↓
MarkdownRenderer (Service)
```

### 1.2 핵심 라이브러리
- **Bubbletea**: TUI 프레임워크
- **Lipgloss**: 스타일링
- **Chroma**: 문법 하이라이팅 (코드 블록)
- **Goldmark**: 마크다운 파서

## 2. 구현 태스크

### Phase 1: 기본 입력창 (1일)
- [ ] TextInput 모델 구현
- [ ] Bubbletea 기본 이벤트 처리
- [ ] 렌더링 기본 구조
- [ ] 입력 버퍼 관리

### Phase 2: 마크다운 렌더링 (2일)
- [ ] Goldmark 마크다운 파서 통합
- [ ] 인라인 요소 렌더링 (볼드, 이탤릭, 코드)
- [ ] 블록 요소 렌더링 (헤더, 리스트, 인용구)
- [ ] 링크 처리

### Phase 3: 문법 하이라이팅 (2일)
- [ ] Chroma 통합
- [ ] 코드 블록 하이라이팅
- [ ] 인라인 코드 하이라이팅
- [ ] 테마 적용

### Phase 4: 멀티라인 입력 (1일)
- [ ] Shift+Enter 처리
- [ ] 들여쓰기 유지
- [ ] 커서 위치 관리
- [ ] 멀티라인 버퍼 처리

### Phase 5: 자동 완성 (1일)
- [ ] 마크다운 문법 패턴 정의
- [ ] Tab 키 처리
- [ ] 제안 목록 표시
- [ ] 선택 및 적용

### Phase 6: 미리보기 (1일)
- [ ] 분할 화면 레이아웃
- [ ] 실시간 미리보기
- [ ] 스크롤 동기화
- [ ] 렌더링 최적화

## 3. 파일 구조

```
internal/ui/
├── markdown/
│   ├── input.go          # TextInput 모델
│   ├── renderer.go       # 마크다운 렌더러
│   ├── highlighter.go    # 문법 하이라이팅
│   ├── autocomplete.go   # 자동 완성
│   └── preview.go        # 미리보기
└── components/
    ├── textarea.go       # 텍스트 영역 컴포넌트
    └── viewport.go       # 뷰포트 컴포넌트
```

## 4. 데이터 모델

```go
type TextInput struct {
    Value      string
    Cursor     int
    Mode       InputMode // SingleLine, MultiLine
    History    []string
    AutoComplete *AutoComplete
}

type MarkdownRenderer struct {
    parser     markgold.Markdown
    highlighter chroma.Highlighter
    theme      Theme
}

type AutoComplete struct {
    Suggestions []Suggestion
    Selected    int
    Visible     bool
}
```

## 5. 의존성 관계

- **file-path-detection**: 마크다운 텍스트에서 경로 감지에 의존
- **theming**: 테마 설정에 의존
- **command-execution**: Enter 키 처리에 의존

## 6. API 설계

```go
// TextInput 모델
type TextInputModel struct {
    textInput textinput.Model
    renderer  *MarkdownRenderer
    preview   *PreviewPane
}

// Bubbletea 인터페이스 구현
func (m TextInputModel) Init() tea.Cmd
func (m TextInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd)
func (m TextInputModel) View() string

// 마크다운 렌더러
func (r *MarkdownRenderer) Render(input string) string
func (r *MarkdownRenderer) RenderInline(input string) string
func (r *MarkdownRenderer) RenderBlock(input string) string
```

## 7. 테스트 전략

### 단위 테스트
- [ ] 마크다운 파서 테스트
- [ ] 렌더러 출력 테스트
- [ ] 하이라이팅 테스트
- [ ] 자동 완성 로직 테스트

### 통합 테스트
- [ ] Bubbletea 메시지 흐름 테스트
- [ ] 멀티라인 입력 테스트
- [ ] 미리보기 동기화 테스트

### 성능 테스트
- [ ] 대용량 텍스트 렌더링 (10,000줄)
- [ ] 실시간 입력 응답 시간 (< 100ms)

## 8. 예상 작업 기간

- **총 작업 시간**: 8일
- **개발자**: 1명
- **종속 작업**: theming (테마 설정 필요)

## 9. 위험 요소 및 완화 계획

### 위험 1: 성능 저하
- **원인**: 대용량 텍스트 렌더링
- **완화**: 가상 스크롤, 렌더링 최적화

### 위험 2: 복잡한 마크다운 문법
- **원인**: 모든 마크다운 기능 지원
- **완화**: MVP에서는 기본 기능만 구현

### 위험 3: Bubbletea 학습 곡선
- **원인**: MVU 패턴 익숙하지 않음
- **완화**: 간단한 예제부터 시작

## 10. 롤백 계획

구현이 실패할 경우:
- **Phase 1**: 기본 텍스트 입력만 유지
- **Phase 2**: 렌더링 없이 텍스트만 표시
- **Phase 3**: 하이라이팅 제외

## 11. 성공 기준

- [ ] 모든 Phase 1-6 태스크 완료
- [ ] 단위 테스트 커버리지 80% 이상
- [ ] 성능 기준 충족 (100ms 이내 렌더링)
- [ ] acceptance.md의 모든 시나리오 통과
