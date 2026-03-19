# CHANGELOG - File Path Detection

## [0.1.0] - 2026-03-19

### Added
- 초기 스펙 문서 생성 (spec.md)
- CHANGELOG 생성

### Implemented
- ✅ 정규식 기반 파일 경로 감지
- ✅ 상대경로 지원 (./, ../)
- ✅ 절대경로 지원
- ✅ 홈 디렉토리 경로 (~/)
- ✅ 라인 번호 포함 경로 (file.md:10)
- ✅ 파일 존재 여부 확인
- ✅ 시각적 표시 (녹색/빨간색)
- ✅ 마크다운 렌더러 통합

### Status
- **Current Status**: COMPLETED
- **Integration**: 마크다운 모드에서 실시간 경로 감지

### Features
- **PathDetector 인터페이스**: 경로 감지 및 해석
- **FilePath 구조체**: 경로 정보 (원본, 해석, 존재 여부, 라인 번호)
- **정규식 패턴**: 4가지 경로 패턴 지원
- **Lipgloss 스타일링**: 존재 여부에 따른 컬러 표시

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
