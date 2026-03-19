# md-terminal 문서

마크다운 기반 터미널 애플리케이션 프로젝트 문서입니다.

## 📋 문서 구조

```
docs/
├── README.md           # 이 파일 (문서 인덱스)
├── DESIGN.md           # 프로젝트 전체 설계
├── specs/              # 피처별 스펙 문서
│   ├── markdown-input/
│   ├── file-path-detection/
│   ├── quick-file-open/
│   ├── command-execution/
│   └── theming/
└── adr/                # 아키텍처 의사결정 기록
    ├── 001-tech-stack.md
    ├── 002-architecture.md
    └── 003-path-detection.md
```

## 🚀 빠른 시작

1. **전체 설계 확인**: [DESIGN.md](./DESIGN.md) 읽기
2. **아키텍처 의사결정 확인**: [adr/](./adr/) 디렉토리 확인
3. **피처별 스펙 확인**: [specs/](./specs/)에서 원하는 피처 선택

## 📚 주요 문서

### [DESIGN.md](./DESIGN.md)
프로젝트 전체 설계 문서
- 시스템 아키텍처
- 기술 스택
- 전체 기능 목록
- 데이터 모델
- 인터페이스 설계
- 개발 로드맵

### [specs/](./specs/)
피처별 상세 스펙 (EARS 형식)

### [adr/](./adr/)
아키텍처 의사결정 기록 (ADR)

## 🔄 문서 작성 플로우

```
1. 프로젝트 시작
   /project-init → DESIGN.md 생성

2. 피처 기획
   /spec-create → docs/specs/{feature}/ 생성

3. 의사결정
   /adr-create → docs/adr/ 생성

4. 변경 추적
   /spec-update → CHANGELOG.md 업데이트
```

## 📊 현재 진행 상황

### 피처 목록
- [ ] markdown-input - 마크다운 입력 및 렌더링
- [ ] file-path-detection - 파일 경로 자동 감지
- [ ] quick-file-open - 빠른 파일 열기
- [ ] command-execution - 터미널 명령 실행
- [ ] theming - 테마 및 설정

### 최근 변경사항
- 2026-03-19: 프로젝트 초기화

---

**문서 관리**: 이 문서는 `/project-init` 실행 시 자동으로 생성되며, 새로운 피처가 추가될 때마다 업데이트됩니다.
