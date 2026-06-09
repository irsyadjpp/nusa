# Changelog - NUSA Platform

All notable changes to the NUSA Platform will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned - Sprint 4
- AI Copilot untuk semua workspace
- Analytics dan reporting dashboard
- Student progress tracking
- Principal dashboard
- Multi-school support enhancement
- Curriculum versioning system
- Learning recommendation engine

### Planned - Sprint 3.5
- Resource authorization architecture (5-layer)
- Unified versioning strategy untuk semua entities
- Evidence storage architecture
- Observability dan operational readiness
- Permission caching optimization

## [0.3.0] - 2026-06-09

### Added - Sprint 3B-3C
- Frontend architecture blueprint dengan 10 section lengkap
- TPSet aggregate implementation di semua layer (domain, application, repository, handler)
- Complete domain implementation untuk CP, TP, ATP, Modul Ajar
- Assessment, Evidence, Evaluation, Achievement services
- Narrative Report generation system
- Frontend workspace components untuk semua 8 workspace:
  - TP Workspace (8 components)
  - ATP Workspace (5 components)
  - Modul Ajar Workspace (5 components)
  - Assessment Designer (8 components)
  - Rubric Designer (5 components)
  - Evidence Workspace (6 components)
  - Evaluation Workspace (6 components)
  - Achievement Dashboard (6 components)
  - Narrative Report Builder (9 components)
- TanStack Query dan Zustand untuk state management
- Query dan Command services untuk semua domain
- Comprehensive API clients untuk semua endpoints
- AI Orchestration Architecture dengan 6 workflow
- Sprint 4 Roadmap dengan 42 fitur untuk 8 epik

### Changed
- Refactored architecture governance documentation
- Centralized markdown documentation structure di `docs/centralized/`
- Migration implementation untuk TPSet dan domain entities
- Docker configuration merged dengan Podman support
- Updated README.md untuk solo-developer context
- Updated CONTRIBUTING.md untuk single-maintainer project

### Fixed
- Import errors di frontend components (getEvaluationsByEvidence, axios, lib/utils)
- Dependency conflicts (vite downgrade 8.0.16 → 7.3.5 untuk @cloudflare/vite-plugin compatibility)
- LayoutContextProvider import path
- Duplicate migration files removed
- PostgreSQL extensions SQL moved ke internal/database/

## [0.2.0] - 2026-06-08

### Added - Sprint 3A
- Domain stabilization dengan KKTP embedded di TP
- Assessment domain refactor untuk reference TP
- Evidence dan Evaluation aggregates
- Achievement service (runtime calculation only, no persistence)
- Backend API implementation untuk semua domain
- Database migrations dengan 30 tables
- Frontend implementation untuk 7 workspace phases
- OpenAPI specification generation
- Workflow validation dan defect identification (9 defects)

### Changed
- TP domain: embedded SuccessCriteria (KKTPCriteria) sebagai Value Object
- Assessment domain: replaced ModulAjarID dengan TPID, added TPVersionNo dan SuccessCriteriaSnapshot
- Evaluation model: added revision tracking capability (revision_no, teacher_feedback)
- Database schema: added success_criteria JSONB di TP table

### Fixed
- 9 defects identified dari workflow validation:
  - DEF-001 (CRITICAL): Evaluation updates in-place without creating revisions
  - DEF-002 (CRITICAL): No integration between Report and Achievement services
  - DEF-003 (HIGH): TP entities lack individual version tracking
  - DEF-004 (HIGH): RevisionNo field never incremented
  - DEF-005 (HIGH): NarrativeReport does not reference achievement data
  - DEF-006 (MEDIUM): TP update protection when downstream assessments exist
  - DEF-007 (MEDIUM): Assessment snapshot incomplete
  - DEF-008 (MEDIUM): Evaluation history query not implemented
  - DEF-009 (MEDIUM): Teacher feedback history not preserved

## [0.1.0] - 2026-06-01

### Added - Initial MVP
- Project initialization
- Basic architecture setup (Modular Monolith, DDD Lite, Layered Architecture)
- Domain model untuk Kurikulum Merdeka (CP, TP, ATP, Modul Ajar, Assessment, Evidence)
- Initial database schema dengan core tables
- Basic frontend structure dengan React dan TypeScript
- Authentication system (admin@nusa.local / admin123)
- Multi-school support foundation
- Role-based access control (RBAC)
- JWT-based authentication dengan refresh tokens
- Docker/Podman configuration untuk local development

### Changed
- Initial commit dari project template

## [0.0.1] - 2026-05-15

### Added
- Project repository initialization
- Basic project structure
- MIT License
- Initial documentation

---

## Versioning Scheme

- **Major version (X.0.0)**: Breaking changes, major architectural refactoring, database migration breaking changes
- **Minor version (0.X.0)**: New features, significant additions, sprint releases
- **Patch version (0.0.X)**: Bug fixes, small improvements, documentation updates

## Sprint Progress

### ✅ Completed
- **Sprint 3A**: Domain implementation, backend API, database migration
- **Sprint 3B**: Frontend implementation, workspace components, state management
- **Sprint 3C**: UAT validation, integration testing, frontend architecture blueprint

### 🚧 In Progress
- **Sprint 3.5**: Resource authorization, unified versioning, evidence storage, observability

### 📋 Planned
- **Sprint 4 Q3**: AI Copilot, Analytics, Student Progress Tracking (14 P0 features)
- **Sprint 4 Q4**: Principal Dashboard, Multi-School Support, Curriculum Versioning (28 P1+P2 features)

## Roadmap Highlights

### Sprint 4 Features (42 total)
- **AI Copilot** (6 workflows): TP, ATP, Modul Ajar, Assessment, Rubric, Narrative Report generation
- **Analytics** (8 features): Teacher performance, student progress, class achievement, competency mastery
- **Student Progress Tracking** (5 features): Individual trajectory, intervention alerts, parent portal
- **Principal Dashboard** (6 features): School overview, teacher performance, compliance monitoring
- **Multi-School Support** (5 features): Cross-school analytics, resource sharing, standardized reporting
- **Curriculum Versioning** (6 features): Version history, approval workflows, rollback capabilities
- **Learning Recommendation Engine** (6 features): Personalized learning paths, content recommendation

Target: 50+ school onboarding by Q4 2026

---

## Catatan Penting

- Semua perubahan besar harus direview terhadap `docs/ARCHITECTURE_FREEZE_V2.md`
- Breaking changes akan diumumkan minimal 1 sprint sebelumnya
- Security fixes akan di-release segera sebagai patch version
- Dokumentasi harus selalu update bersama dengan kode
- Database migrations harus memiliki backward compatibility sebaik mungkin

## Solo Developer Notes

Sebagai proyek solo-developer:
- Timeline bisa bervariasi tergantung ketersediaan
- Priority bisa berubah berdasarkan feedback dan kebutuhan
- Kolaborasi dari komunitas sangat diapresiasi
- Focus pada kualitas di atas kecepatan
- Architecture governance strict untuk maintainability

Untuk detail lebih lanjut tentang setiap sprint, lihat dokumen di folder `docs/` dan `docs/centralized/`.
