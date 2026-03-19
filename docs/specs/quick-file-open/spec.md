# Quick File Open Spec

## 1. 개요

감지된 파일 경로를 클릭하여 바로 파일을 열 수 있는 기능입니다.

## 2. 요구사항 (EARS)

### 2.1 기능 요구사항

**WHEN** 사용자가 하이퍼링크를 클릭하면 **THEN** 시스템은 해당 파일을 엽니다 **SHALL**

**WHERE** 라인 번호가 포함된 경로이면 **THE SYSTEM SHALL** 해당 라인으로 이동

**IF** 기본 에디터가 설정되어 있으면 **THE SYSTEM SHALL** 그 에디터를 사용

**WHILE** 파일이 열리면 **THE SYSTEM SHALL** 터미널 포커스를 유지

### 2.2 비기능 요구사항

**WHEN** 클릭이 발생하면 **THEN** 파일은 200ms 이내에 열려야 합니다 **SHALL**

## 3. 사용자 스토리

```gherkin
Feature: Quick File Open

  Scenario: Open file with click
    Given 사용자가 "./file.md" 하이퍼링크를 클릭
    Then 시스템은 file.md를 엽니다

  Scenario: Open file at specific line
    Given 사용자가 "file.md:10" 하이퍼링크를 클릭
    Then 시스템은 file.md를 열고 10번째 줄로 이동

  Scenario: Open with custom editor
    Given 사용자가 vim을 기본 에디터로 설정
    And 하이퍼링크를 클릭
    Then 시스템은 vim으로 파일을 엽니다
```

## 4. 기능 목록

- [ ] 마우스 클릭으로 파일 열기 (계획중)
- [ ] 키보드 단축키로 파일 열기 (계획중)
- [ ] 라인 번호로 이동 (계획중)
- [ ] 기본 에디터 연동 (계획중)
- [ ] 커스텀 에디터 설정 (계획중)
- [ ] 터미널 포커스 유지 (계획중)
- [x] 문서 완성 (spec.md, plan.md, acceptance.md, CHANGELOG.md)

## 5. 에디터 지원

### 기본 지원
- VS Code: `code --goto file.md:10`
- Vim: `vim +10 file.md`
- Nano: `nano +10 file.md`
- Emacs: `emacs +10 file.md`

### 설정 파일 예시
```yaml
# ~/.config/mdterm/config.yaml
editor: code
editor_args:
  - --goto
```

## 6. 인터페이스

```go
type FileOpener interface {
    OpenFile(path string, line int) error
    GetDefaultEditor() string
    SetEditor(editor string) error
}
```

## 7. 변경 이력

- 2026-03-19: 스펙 초기 생성
