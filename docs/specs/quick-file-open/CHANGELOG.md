# CHANGELOG - Quick File Open

## [0.1.0] - 2026-03-19

### Added
- 초기 스펙 문서 생성 (spec.md)
- 구현 계획 추가 (plan.md)
- 인수 기준 추가 (acceptance.md)
- CHANGELOG 생성

### Features
- 마우스 클릭으로 파일 열기 (계획중)
- 키보드 단축키로 파일 열기 (계획중)
- 라인 번호로 이동 (계획중)
- 기본 에디터 연동 (계획중)
- 커스텀 에디터 설정 (계획중)
- 터미널 포커스 유지 (계획중)

### Status
- **Current Status**: PLANNED
- **Next Steps**: 구현 시작

### Implementation Plan

#### Phase 1: 기본 파일 열기 (1일)
- FileOpener 인터페이스 구현
- 기본 에디터 감지
- 단순 파일 열기

#### Phase 2: 라인 번호 지원 (1일)
- 라인 번호 파싱
- 에디터별 라인 번호 옵션

#### Phase 3: 마우스 클릭 처리 (1일)
- Bubbletea 마우스 이벤트
- 클릭 위치 감지

#### Phase 4: 커스텀 에디터 (1일)
- 설정 파일 파싱
- 커스텀 에디터 실행

#### Phase 5: 터미널 포커스 유지 (1일)
- 백그라운드 실행
- 프로세스 관리

### Dependencies
- file-path-detection: ✅ 완료
- markdown-input: ✅ 완료

### Estimated Duration
- **Total**: 5 days
- **Developer**: 1 person

---

## 버전 규칙

이 프로젝트는 [Semantic Versioning](https://semver.org/)을 따릅니다:

- **MAJOR**: 기존 기능의 호환되지 않는 변경
- **MINOR**: 새로운 기능 추가 (기존 기능과 호환)
- **PATCH**: 버그 수정

## 변경 유형

- **Added**: 새로운 기능 추가
- **Changed**: 기존 기능의 변경
- **Deprecated**: 향후 제거될 기능
- **Removed**: 제거된 기능
- **Fixed**: 버그 수정
- **Security**: 보안 관련 수정
