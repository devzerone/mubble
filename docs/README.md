# mubble 문서

마크다운 기반 터미널 애플리케이션 프로젝트 문서입니다.

## 📋 문서 구조

```
docs/
├── README.md           # 이 파일 (문서 인덱스)
├── DESIGN.md           # 프로젝트 전체 설계
├── specs/              # 피처별 스펙 문서
│   ├── markdown-input/      # ✅ 완료 (spec.md, plan.md, acceptance.md, CHANGELOG.md)
│   ├── file-path-detection/ # ✅ 완료 (spec.md, CHANGELOG.md)
│   ├── quick-file-open/     # 📋 계획중 (spec.md만)
│   ├── command-execution/   # 📋 계획중 (spec.md만)
│   └── theming/             # 📋 계획중 (spec.md만)
└── adr/                # 아키텍처 의사결정 기록
    ├── 001-tech-stack.md
    ├── 002-architecture.md
    ├── 003-path-detection.md
    └── 004-markdown-mode.md
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
- [x] **markdown-input** - ✅ 완료 (마크다운 입력 및 렌더링)
- [x] **file-path-detection** - ✅ 완료 (파일 경로 자동 감지)
- [x] **quick-file-open** - 📋 문서 완료 (빠른 파일 열기)
- [ ] command-execution - 📋 계획중 (터미널 명령 실행)
- [ ] theming - 📋 계획중 (테마 및 설정)

### 구현 완료 기능
- ✅ 마크다운 실시간 렌더링
- ✅ 문법 하이라이팅 (헤더, 볼드, 이탤릭, 코드, 링크)
- ✅ 모드 전환 (Ctrl+M: 터미널 ↔ 마크다운)
- ✅ 분할 화면 레이아웃
- ✅ 파일 경로 자동 감지 (정규식)
- ✅ 파일 존재 확인 및 시각적 표시
- ✅ 상대경로, 절대경로, 홈 디렉토리 경로 지원
- ✅ 라인 번호 포함 경로 지원

### 진행중 기능
- 🚀 파일 클릭으로 열기 (다음 구현 예정)

### 최근 변경사항
- 2026-03-19: 프로젝트 초기화
- 2026-03-19: 마크다운 렌더러 구현 완료
- 2026-03-19: 마크다운 모드 전환 기능 구현 완료
- 2026-03-19: 파일 경로 감지 기능 구현 완료
- 2026-03-19: 전체 문서 업데이트
- 2026-03-19: 파일 클릭으로 열기 기능 문서 완료 📋

---

**문서 관리**: 이 문서는 프로젝트 진행 상황에 맞춰 지속적으로 업데이트됩니다.
