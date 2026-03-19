# File Path Detection Spec

## 1. 개요

마크다운 텍스트에서 파일 경로를 자동으로 감지하고 하이퍼링크로 변환하는 기능입니다.

## 2. 요구사항 (EARS)

### 2.1 기능 요구사항

**WHEN** 텍스트에 파일 경로 패턴이 나타나면 **THEN** 시스템은 자동으로 감지 **SHALL**

**WHERE** 상대경로가 입력될 때 **THE SYSTEM SHALL** 절대 경로로 해석

**IF** 파일이 존재하면 **THE SYSTEM SHALL** 녹색 하이퍼링크로 표시

**WHILE** 파일이 존재하지 않으면 **THE SYSTEM SHALL** 빨간색으로 표시

### 2.2 비기능 요구사항

**WHEN** 텍스트가 입력될 때 **THEN** 경로 감지는 50ms 이내에 완료 **SHALL**

**WHERE** 여러 경로가 있을 때 **THE SYSTEM SHALL** 모두 감지

## 3. 사용자 스토리

```gherkin
Feature: File Path Detection

  Scenario: Detect relative path
    Given 사용자가 "./file.md"를 입력
    And file.md가 존재
    Then 시스템은 녹색 하이퍼링크로 표시

  Scenario: Detect non-existent file
    Given 사용자가 "./missing.md"를 입력
    And missing.md가 존재하지 않음
    Then 시스템은 빨간색으로 표시

  Scenario: Detect path with line number
    Given 사용자가 "file.md:10"을 입력
    Then 시스템은 라인 번호를 포함하여 감지
```

## 4. 기능 목록

- [x] 상대경로 감지 (./, ../)
- [x] 절대경로 감지
- [x] 홈 디렉토리 경로 (~/)
- [x] 라인 번호 포함 경로 (file.md:10)
- [x] 파일 존재 여부 확인
- [x] 시각적 표시 (존재: 녹색, 미존재: 빨간색)
- [x] 정규식 기반 패턴 매칭
- [x] 마크다운 렌더러와 통합

## 5. 지원 경로 패턴

```
./file.md
../file.md
~/Documents/file.md
file.md
file.md:10
./src/main.go:25
```

## 6. 인터페이스

```go
type PathDetector interface {
    DetectPaths(input string) []FilePath
    ResolvePath(relPath string) (string, error)
    ValidatePath(path string) bool
}

type FilePath struct {
    Original   string
    Resolved   string
    Exists     bool
    LineNumber int
    StartPos   int
    EndPos     int
}
```

## 7. 변경 이력

- 2026-03-19: 스펙 초기 생성
- 2026-03-19: 구현 완료 및 마크다운 렌더러 통합
- 2026-03-19: 정규식 기반 경로 감지 구현
