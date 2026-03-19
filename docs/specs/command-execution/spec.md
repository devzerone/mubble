# Command Execution Spec

## 1. 개요

터미널 명령을 실행하고 결과를 표시하는 기능입니다.

## 2. 요구사항 (EARS)

### 2.1 기능 요구사항

**WHEN** 사용자가 명령을 입력하고 Enter를 누르면 **THEN** 시스템은 명령을 실행 **SHALL**

**WHERE** 명령 실행 중 오류가 발생하면 **THE SYSTEM SHALL** 오류 메시지를 표시

**IF** 명령이 백그라운드 실행 요청이면 **THE SYSTEM SHALL** 백그라운드에서 실행

**WHILE** 명령이 실행 중이면 **THE SYSTEM SHALL** 진행 상태를 표시

### 2.2 비기능 요구사항

**WHEN** 명령이 완료되면 **THEN** 출력은 500ms 이내에 표시 **SHALL**

## 3. 사용자 스토리

```gherkin
Feature: Command Execution

  Scenario: Execute simple command
    Given 사용자가 "ls -la"를 입력하고 Enter를 누름
    Then 시스템은 ls 명령을 실행
    And 결과를 표시

  Scenario: Command with error
    Given 사용자가 "invalid-command"를 입력
    Then 시스템은 오류 메시지를 표시

  Scenario: Background execution
    Given 사용자가 "npm start &"를 입력
    Then 시스템은 백그라운드에서 실행
    And 즉시 새 프롬프트를 표시
```

## 4. 기능 목록

- [ ] 쉘 명령 실행
- [ ] 출력 캡처 및 표시
- [ ] 명령 히스토리 (위/아래 화살표)
- [ ] 백그라운드 실행 (&)
- [ ] 진행 상태 표시
- [ ] 오류 처리

## 5. 명령 히스토리

```go
type CommandHistory struct {
    commands []string
    index    int
    max      int
}

func (h *CommandHistory) Add(cmd string)
func (h *CommandHistory) Previous() string
func (h *CommandHistory) Next() string
```

## 6. 인터페이스

```go
type CommandExecutor interface {
    Execute(cmd string) (string, error)
    ExecuteInBackground(cmd string) error
    IsRunning() bool
    Stop() error
}
```

## 7. 변경 이력

- 2026-03-19: 스펙 초기 생성
