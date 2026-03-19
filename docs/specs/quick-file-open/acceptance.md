# Quick File Open Acceptance Criteria

## 1. 기본 파일 열기

### Scenario 1.1: 클릭으로 파일 열기
```gherkin
Given 사용자가 마크다운 모드에 있음
And "./docs/README.md"가 녹색 하이퍼링크로 표시됨
When 해당 하이퍼링크를 클릭
Then 시스템은 기본 에디터로 README.md를 엽니다
And 터미널 포커스는 유지됨
```

### Scenario 1.2: 존재하지 않는 파일
```gherkin
Given 사용자가 "./missing.md"가 빨간색으로 표시됨
When 해당 경로를 클릭
Then 시스템은 에러 메시지를 표시
And "파일을 찾을 수 없습니다: missing.md" 메시지 확인
```

## 2. 라인 번호로 이동

### Scenario 2.1: VS Code로 라인 이동
```gherkin
Given 사용자가 "src/main.go:25"를 클릭
And VS Code가 기본 에디터로 설정됨
When 시스템이 파일을 엽니다
Then VS Code가 main.go를 열고 25번째 줄로 커서를 이동
```

### Scenario 2.2: Vim으로 라인 이동
```gherkin
Given 사용자가 vim을 기본 에디터로 설정
And "file.md:10"을 클릭
When 시스템이 파일을 엽니다
Then vim이 file.md를 열고 10번째 줄로 커서를 이동
```

### Scenario 2.3: 라인 번호 없는 파일
```gherkin
Given 사용자가 "README.md"를 클릭 (라인 번호 없음)
When 시스템이 파일을 엽니다
Then 에디터가 파일의 첫 번째 줄부터 표시
```

## 3. 다양한 경로 패턴

### Scenario 3.1: 상대 경로
```gherkin
Given 사용자가 "./src/utils/helper.go"를 클릭
When 파일이 존재하면
Then 시스템은 절대 경로로 해석하여 파일을 엽니다
```

### Scenario 3.2: 상위 디렉토리
```gherkin
Given 사용자가 "../config/app.yaml"를 클릭
When 파일이 존재하면
Then 시스템은 상위 디렉토리의 파일을 엽니다
```

### Scenario 3.3: 홈 디렉토리
```gherkin
Given 사용자가 "~/Documents/notes.md"를 클릭
When 파일이 존재하면
Then 시스템은 홈 디렉토리의 파일을 엽니다
```

## 4. 에디터 연동

### Scenario 4.1: 자동 에디터 감지
```gherkin
Given 시스템이 시작될 때
When 사용자가 설정 파일을 생성하지 않음
Then 시스템은 다음 우선순위로 에디터를 선택
  1. code (VS Code)
  2. vim
  3. nano
```

### Scenario 4.2: 커스텀 에디터 설정
```gherkin
Given 사용자가 ~/.config/mubble/config.yaml에 설정
```yaml
editor: custom-editor
editor_args:
  - --line
```
When "file.md:5"를 클릭
Then 시스템은 "custom-editor --line 5 file.md"를 실행
```

## 5. 터미널 포커스

### Scenario 5.1: 백그라운드 실행
```gherkin
Given 사용자가 파일을 클릭하여 엶니다
When 에디터가 실행될 때
Then 터미널은 입력을 계속 받을 수 있어야 함
And mubble은 계속 실행 중이어야 함
```

### Scenario 5.2: 포어그라운드 실행
```gherkin
Given 사용자가 Shift+Click를 누름
When 파일을 엽니다
Then 터미널은 에디터가 종료될 때까지 대기
```

## 6. 키보드 단축키

### Scenario 6.1: Tab 키로 파일 열기
```gherkin
Given 사용자가 커서가 파일 경로 위에 있음
When Tab 키를 누르면
Then 시스템은 해당 파일을 엽니다
```

### Scenario 6.2: Enter 키로 파일 열기
```gherkin
Given 사용자가 커서가 파일 경로 위에 있음
And Enter 키를 누르면
Then 시스템은 해당 파일을 엽니다
```

## 7. 에러 처리

### Scenario 7.1: 에디터를 찾을 수 없음
```gherkin
Given 사용자가 파일을 클릭
When 설정된 에디터를 찾을 수 없음
Then 시스템은 "에디터를 찾을 수 없습니다" 메시지를 표시
And 대안 에디터 목록을 표시
```

### Scenario 7.2: 파일 읽기 권한 없음
```gherkin
Given 사용자가 읽기 권한이 없는 파일을 클릭
When 시스템이 파일을 열려고 시도
Then "권한이 없습니다" 에러를 표시
```

## 8. 성능

### Scenario 8.1: 빠른 파일 열기
```gherkin
Given 사용자가 파일을 클릭
When 200ms 이내에
Then 에디터가 시작되어야 함
```

### Scenario 8.2: 여러 파일 열기
```gherkin
Given 사용자가 3개의 파일을 연속으로 클릭
When 각 파일이 200ms 이내에 열리면
Then 모든 에디터가 정상적으로 실행됨
```

## 9. 사용자 피드백

### Scenario 9.1: 열기 성공 메시지
```gherkin
Given 사용자가 파일을 클릭
When 파일이 성공적으로 열리면
Then 상태바에 "file.md를 열었습니다" 메시지를 표시
```

### Scenario 9.2: 열기 실패 메시지
```gherkin
Given 사용자가 파일을 클릭
When 파일 열기가 실패하면
Then 상태바에 "파일을 열 수 없습니다: [에러]" 메시지를 표시
And 빨간색으로 강조
```

## 10. 통합 테스트

### Scenario 10.1: 전체 워크플로우
```gherkin
Given 사용자가 마크다운 모드에 있음
And 다음 텍스트를 입력:
  """
  # 프로젝트 문서

  문서는 [./docs/README.md](./docs/README.md)를 참고하세요
  코드는 [src/main.go:25](src/main.go:25)에 있습니다
  """
When 각 경로를 클릭
Then README.md가 열리고
And main.go가 25번째 줄로 열림
```

### Scenario 10.2: 에디터 전환
```gherkin
Given 사용자가 vim으로 파일을 엶습니다
When 설정을 code로 변경
And 다른 파일을 클릭
Then VS Code가 파일을 엽니다
```
