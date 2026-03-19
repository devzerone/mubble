# Markdown Input Acceptance Criteria

## 1. 헤더 렌더링

### Scenario 1.1: H1 헤더
```gherkin
Given 사용자가 입력창에 "# 제목"을 입력
When 텍스트가 렌더링될 때
Then 시스템은 큰 폰트와 굵은 스타일로 "제목"을 표시
And 폰트 크기는 일반 텍스트의 2배여야 한다
```

### Scenario 1.2: H2-H6 헤더
```gherkin
Given 사용자가 "## 소제목"을 입력
When 텍스트가 렌더링될 때
Then 시스템은 중간 크기의 굵은 스타일로 "소제목"을 표시
And 헤더 레벨에 따라 폰트 크기가 감소해야 한다
```

## 2. 텍스트 스타일링

### Scenario 2.1: 볼드 텍스트
```gherkin
Given 사용자가 "**중요한 텍스트**"를 입력
When 텍스트가 렌더링될 때
Then 시스템은 "중요한 텍스트"를 굵은 폰트로 표시
And 별표 문자는 표시하지 않는다
```

### Scenario 2.2: 이탤릭 텍스트
```gherkin
Given 사용자가 "*이탤릭*"를 입력
When 텍스트가 렌더링될 때
Then 시스템은 "이탤릭"을 기울어진 폰트로 표시
```

### Scenario 2.3: 인라인 코드
```gherkin
Given 사용자가 "`code`"를 입력
When 텍스트가 렌더링될 때
Then 시스템은 "code"를 다른 배경색과 고정폭 폰트로 표시
```

## 3. 코드 블록

### Scenario 3.1: 기본 코드 블록
```gherkin
Given 사용자가 다음을 입력:
  """
  ```go
  func main() {
      fmt.Println("Hello")
  }
  ```
  """
When 텍스트가 렌더링될 때
Then 시스템은 코드를 회색 배경의 박스로 표시
And 고정폭 폰트를 사용한다
```

### Scenario 3.2: 문법 하이라이팅
```gherkin
Given Go 언어 코드 블록이 입력됨
When 코드가 렌더링될 때
Then 키워드(func, var, const)는 파란색으로 표시
And 문자열은 녹색으로 표시
And 주석은 회색으로 표시
```

### Scenario 3.3: 다양한 언어 지원
```gherkin
Given JavaScript, Python, Rust 코드가 입력됨
When 각 코드가 렌더링될 때
Then 각 언어의 문법에 맞게 하이라이팅 되어야 한다
```

## 4. 리스트

### Scenario 4.1: 순서 없는 리스트
```gherkin
Given 사용자가 다음을 입력:
  """
  - 항목 1
  - 항목 2
  - 항목 3
  """
When 텍스트가 렌더링될 때
Then 시스템은 각 항목 앞에 불렛 포인트를 표시
And 항목들은 세로로 정렬되어야 한다
```

### Scenario 4.2: 순서 있는 리스트
```gherkin
Given 사용자가 다음을 입력:
  """
  1. 첫 번째
  2. 두 번째
  3. 세 번째
  """
When 텍스트가 렌더링될 때
Then 시스템은 자동 번호 매기기를 표시
```

### Scenario 4.3: 중첩 리스트
```gherkin
Given 사용자가 중첩된 리스트를 입력
When 텍스트가 렌더링될 때
Then 하위 항목들은 들여쓰기 되어야 한다
And 불렛 포인트 스타일이 달라야 한다
```

## 5. 링크

### Scenario 5.1: 일반 링크
```gherkin
Given 사용자가 "[텍스트](url)"을 입력
When 텍스트가 렌더링될 때
Then 시스템은 "텍스트"를 파란색 밑줄로 표시
And 링크 문법은 표시하지 않는다
```

### Scenario 5.2: 이미지 링크
```gherkin
Given 사용자가 "![alt](image.png)"을 입력
When 텍스트가 렌더링될 때
Then 시스템은 이미지를 렌더링하거나 대체 텍스트를 표시
```

## 6. 멀티라인 입력

### Scenario 6.1: Shift+Enter로 새 줄
```gherkin
Given 사용자가 첫 번째 줄에 텍스트를 입력
And Shift+Enter를 누름
When 새 줄이 추가될 때
Then 커서는 새 줄의 시작 부분에 위치
And 명령은 실행되지 않는다
```

### Scenario 6.2: 들여쓰기 유지
```gherkin
Given 사용자가 들여쓰기된 텍스트를 입력
And Shift+Enter를 누름
When 새 줄이 추가될 때
Then 이전 줄의 들여쓰기가 유지되어야 한다
```

### Scenario 6.3: Enter로 명령 실행
```gherkin
Given 사용자가 텍스트를 입력
And Enter를 누름 (Shift 없이)
When 명령이 실행될 때
Then 입력창은 비워진다
And 명령 결과가 표시된다
```

## 7. 자동 완성

### Scenario 7.1: 헤더 자동 완성
```gherkin
Given 사용자가 "#"을 입력하고 Tab을 누름
When 자동 완성 메뉴가 표시될 때
Then "## ", "### " 등의 제안이 표시되어야 한다
```

### Scenario 7.2: 코드 블록 자동 완성
```gherkin
Given 사용자가 "`"를 입력하고 Tab을 누름
When 자동 완성 메뉴가 표시될 때
Then "```go", "```js" 등의 제안이 표시되어야 한다
```

### Scenario 7.3: 제안 선택
```gherkin
Given 자동 완성 메뉴가 표시됨
And 사용자가 위/아래 화살표로 탐색
When Enter를 눌러 제안을 선택하면
Then 선택된 제안이 입력창에 삽입된다
```

## 8. 미리보기

### Scenario 8.1: 실시간 미리보기
```gherkin
Given 사용자가 마크다운을 입력
When 텍스트가 변경될 때마다
Then 미리보기 창이 즉시 업데이트되어야 한다
```

### Scenario 8.2: 스크롤 동기화
```gherkin
Given 사용자가 입력창을 스크롤
When 스크롤 위치가 변경되면
Then 미리보기 창도 동일한 위치로 스크롤되어야 한다
```

## 9. 성능

### Scenario 9.1: 빠른 렌더링
```gherkin
Given 사용자가 텍스트를 입력
When 입력이 발생한 후
Then 렌더링은 100ms 이내에 완료되어야 한다
```

### Scenario 9.2: 대용량 텍스트 처리
```gherkin
Given 10,000줄의 마크다운 텍스트가 입력됨
When 텍스트가 렌더링될 때
Then 애플리케이션은 멈추지 않고 응답해야 한다
And 스크롤은 부드러워야 한다
```

## 10. 에러 처리

### Scenario 10.1: 잘못된 마크다운 문법
```gherkin
Given 사용자가 잘못된 마크다운 문법을 입력
When 텍스트가 렌더링될 때
Then 시스템은 원본 텍스트를 그대로 표시
And 에러 메시지는 표시하지 않는다
```

### Scenario 10.2: 지원하지 않는 기능
```gherkin
Given 사용자가 지원하지 않는 마크다운 기능을 입력
When 텍스트가 렌더링될 때
Then 시스템은 해당 부분을 일반 텍스트로 표시
```

## 11. 접근성

### Scenario 11.1: 키보드 내비게이션
```gherkin
Given 키보드만 사용하는 사용자
When 모든 기능에 접근할 때
Then 마우스 없이 모든 작업을 완료할 수 있어야 한다
```

### Scenario 11.2: 색상 대비
```gherkin
Given 시각 장애가 있는 사용자
When 텍스트가 렌더링될 때
Then 텍스트와 배경색은 충분한 대비를 가져야 한다
```

## 12. 통합 테스트

### Scenario 12.1: 복합 마크다운 문서
```gherkin
Given 사용자가 다음과 같은 복합 문서를 입력:
  """
  # 프로젝트 문서

  ## 개요
  이 프로젝트는 **마크다운 터미널**입니다.

  ## 기능
  - 실시간 렌더링
  - 문법 하이라이팅

  ```go
  func main() {
      fmt.Println("Hello")
  }
  ```

  자세한 내용은 [문서](./docs.md)를 참고하세요.
  """
When 문서가 렌더링될 때
Then 모든 요소가 올바르게 렌더링되어야 한다
And 헤더, 스타일, 리스트, 코드, 링크가 모두 올바르게 표시되어야 한다
```
