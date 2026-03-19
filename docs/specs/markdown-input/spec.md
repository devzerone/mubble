# Markdown Input Spec

## 1. 개요

마크다운 형식의 입력을 실시간으로 렌더링하고 문법 하이라이팅을 제공하는 기능입니다.

## 2. 요구사항 (EARS)

### 2.1 기능 요구사항

**WHEN** 사용자가 입력창에 텍스트를 입력 **THEN** 시스템은 실시간으로 마크다운을 렌더링 **SHALL**

**WHERE** 마크다운 문법이 입력될 때 **THE SYSTEM SHALL** 해당 부분을 문법 하이라이팅

**WHILE** 사용자가 멀티라인 입력을 작성할 때 **THE SYSTEM SHALL** 자동으로 들여쓰기와 포맷을 유지

**IF** 사용자가 자동 완성을 요청하면 **THE SYSTEM SHALL** 마크다운 문법을 제안

### 2.2 비기능 요구사항

**WHEN** 텍스트가 입력될 때 **THEN** 렌더링은 100ms 이내에 완료 **SHALL**

**WHERE** 대용량 텍스트가 입력될 때 **THE SYSTEM SHALL** 성능 저하 없이 처리

## 3. 사용자 스토리

```gherkin
Feature: Markdown Input

  Scenario: 마크다운 모드 진입
    Given 사용자가 일반 터미널 모드에 있음
    And Ctrl+M을 누름
    Then 시스템은 마크다운 모드로 전환
    And 분할 화면이 표시됨 (입력창 + 미리보기)

  Scenario: Basic markdown rendering
    Given 사용자가 입력창에 "# Header"를 입력
    Then 시스템은 큰 헤더 스타일로 렌더링
    And 미리보기 창에 실시간으로 표시

  Scenario: Code block highlighting
    Given 사용자가 코드 블록을 입력
    Then 시스템은 구문 강조를 적용

  Scenario: Multi-line input
    Given 사용자가 마크다운 모드에서 Enter를 누름
    Then 시스템은 새 줄을 추가
    And 명령을 실행하지 않음

  Scenario: 마크다운 모드 종료
    Given 사용자가 마크다운 모드에 있음
    And Ctrl+M을 다시 누름
    Then 시스템은 일반 터미널 모드로 복귀
```

## 4. 기능 목록

- [x] 마크다운 실시간 렌더링
- [ ] **마크다운 모드 전환 (Ctrl+M)**
- [ ] 분할 화면 레이아웃 (입력창 + 미리보기)
- [ ] 문법 하이라이팅 (헤더, 볼드, 이탤릭, 코드, 링크)
- [ ] 멀티라인 입력 지원
- [ ] 자동 완성
- [ ] 들여쓰기 유지
- [x] 마크다운 미리보기

## 5. 인터페이스

```go
type MarkdownRenderer interface {
    Render(input string) string
    HighlightSyntax(input string) string
    GetPreview(input string) string
}
```

## 6. 변경 이력

- 2026-03-19: 스펙 초기 생성
- 2026-03-19: 마크다운 모드 전환 기능 추가 (Ctrl+M)
