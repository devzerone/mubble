# ADR 002: 아키텍처 패턴 결정

## 상태
승인됨

## 날짜
2026-03-19

## 문맥
TUI 애플리케이션의 구조를 설계하기 위해 적절한 아키텍처 패턴을 선택해야 합니다.

## 결정
**Component-based MVU (Model-View-Update) 패턴**을 사용합니다.

### 아키텍처 구성
```
┌─────────────┐
│    View     │ (Lipgloss로 렌더링)
└─────────────┘
       ↑
       │ (Display)
       │
┌─────────────┐
│    Model    │ (AppState)
└─────────────┘
       ↑
       │ (Update)
       │
┌─────────────┐
│   Events    │ (키보드, 마우스, 타이머)
└─────────────┘
```

## 상세 설명

### Model (모델)
- 애플리케이션의 전체 상태를 보관
```go
type AppState struct {
    CurrentInput     string
    RenderedMarkdown string
    DetectedPaths    []FilePath
    CommandHistory   []string
    Config           Config
}
```

### View (뷰)
- 상태를 기반으로 UI 렌더링
- Lipgloss를 사용한 스타일링
- 순수 함수: 상태만 입력으로 받음

### Update (업데이트)
- 이벤트를 받아 새 상태를 반환
- 순수 함수: 부작용 없음

## 대안 고려

### 옵션 1: MVU (Bubbletea) ✅ 선택
**장점**
- 예측 가능한 상태 변화
- 테스트 용이성
- 명확한 데이터 흐름
- Bubbletea 프레임워크 지원

### 옵션 2: MVC
**장점**
- 익숙한 패턴

**단점**
- TUI에서는 뷰와 컨트롤러의 분리가 모호함
- 상태 관리가 복잡해질 수 있음

### 옵션 3: Component-based (React-like)
**장점**
- 모듈화된 컴포넌트

**단점**
- Go에서는 비순차적
- 과도한 추상화

## 결정 근거

1. **단순성**: Bubbletea의 MVU 패턴은 간단하고 이해하기 쉬움
2. **테스트 가능성**: 순수 함수로 테스트가 매우 용이
3. **확장성**: 새로운 이벤트와 상태를 쉽게 추가 가능
4. **프레임워크 지원**: Bubbletea가 이 패턴을 완벽하게 지원

## 컴포넌트 구조

```
md-terminal/
├── ui/
│   ├── model.go          # 애플리케이션 상태
│   ├── view.go           # 렌더링 로직
│   ├── update.go         # 이벤트 처리
│   └── components/
│       ├── input.go      # 입력 컴포넌트
│       ├── preview.go    # 미리보기 컴포넌트
│       └── status.go     # 상태바 컴포넌트
├── pathfinder/           # 경로 감지 모듈
├── executor/             # 명령 실행 모듈
└── config/               # 설정 모듈
```

## 결과
- Bubbletea의 MVU 패턴을 사용하여 명확하고 테스트 가능한 코드
- 모듈화된 컴포넌트로 유지보수성 향상
- 단방향 데이터 흐름으로 예측 가능한 동작
