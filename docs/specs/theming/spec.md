# Theming Spec

## 1. 개요

테마 및 설정을 관리하는 기능입니다.

## 2. 요구사항 (EARS)

### 2.1 기능 요구사항

**WHEN** 애플리케이션이 시작되면 **THEN** 시스템은 설정 파일을 로드 **SHALL**

**WHERE** 사용자가 테마를 변경하면 **THE SYSTEM SHALL** 즉시 적용

**IF** 설정 파일이 없으면 **THE SYSTEM SHALL** 기본 설정을 사용

**WHILE** 설정이 변경되면 **THE SYSTEM SHALL** 설정 파일에 저장

### 2.2 비기능 요구사항

**WHEN** 테마가 변경되면 **THEN** 적용은 50ms 이내에 완료 **SHALL**

## 3. 사용자 스토리

```gherkin
Feature: Theming

  Scenario: Load default theme
    Given 애플리케이션이 시작
    And 설정 파일이 없음
    Then 시스템은 기본 테마를 적용

  Scenario: Change theme
    Given 사용자가 테마를 "dark"로 변경
    Then 시스템은 어두운 테마를 적용

  Scenario: Custom editor setting
    Given 사용자가 기본 에디터를 "vim"으로 설정
    Then 파일 열기 시 vim이 사용됨
```

## 4. 기능 목록

- [ ] 밝은/어두운 테마
- [ ] 컬러 스킴 설정
- [ ] 설정 파일 로드/저장
- [ ] 런타임 테마 전환
- [ ] 커스텀 에디터 설정
- [ ] 기본 설정 제공

## 5. 설정 파일 구조

```yaml
# ~/.config/mdterm/config.yaml
theme: dark
editor: code
auto_open: true
max_history: 100

themes:
  dark:
    background: "#1e1e1e"
    foreground: "#d4d4d4"
    primary: "#4fc1ff"
    secondary: "#ce9178"
    success: "#4ec9b0"
    error: "#f48771"
  light:
    background: "#ffffff"
    foreground: "#000000"
    primary: "#0066cc"
    secondary: "#a31515"
    success: "#008000"
    error: "#cd3131"
```

## 6. 인터페이스

```go
type Config struct {
    Theme     string
    Editor    string
    AutoOpen  bool
    MaxHistory int
    Themes    map[string]Theme
}

type Theme struct {
    Background string
    Foreground string
    Primary    string
    Secondary  string
    Success    string
    Error      string
}

type ConfigManager interface {
    Load() (Config, error)
    Save(config Config) error
    GetTheme(name string) (Theme, error)
}
```

## 7. 변경 이력

- 2026-03-19: 스펙 초기 생성
