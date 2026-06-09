# Panduan Kontribusi - NUSA Platform

Terima kasih atas ketertarikan Anda untuk berkontribusi pada NUSA Platform! Ini adalah proyek solo-developer, namun saya sangat menghargai kontribusi dari komunitas.

## 🎯 Konteks Proyek

**Status Proyek**: Solo-developer dengan kualitas produksi
- **Maintainer**: Satu orang (saya)
- **Review Timeline**: 1-3 hari kerja (tergantung ketersediaan)
- **Response Time**: Saya akan berusaha merespon secepat mungkin
- **Bahasa**: Indonesia untuk komunikasi umum, Inggris untuk kode dan dokumentasi teknis

## 🚀 Cara Berkontribusi

### 1. Fork dan Clone Repository
```bash
# Fork repository ini di GitHub
git clone https://github.com/USERNAME/nusa.git
cd nusa
```

### 2. Buat Branch Fitur
```bash
git checkout -b fitur/nama-fitur-anda
```

### 3. Lakukan Perubahan

#### Backend (Go)
- Ikuti arsitektur: Handler → Service → Repository → PostgreSQL
- Tambahkan tes unit untuk setiap fungsi baru
- Pastikan kode sesuai dengan `docs/ARCHITECTURE_FREEZE_V2.md`
- Jalankan format dan lint:
  ```bash
  cd backend
  go fmt ./...
  go vet ./...
  go test ./...
  ```

#### Frontend (React/TypeScript)
- Gunakan komponen yang sudah ada sebelum membuat baru
- Ikuti pola state management (TanStack Query + Zustand)
- Pastikan UX ramah untuk guru SD
- Jalankan:
  ```bash
  cd frontend
  npm install
  npm run lint
  npm run type-check
  ```

### 4. Update Dokumentasi
- Update `CHANGELOG.md` untuk perubahan signifikan
- Tambahkan komentar pada kode yang kompleks
- Update dokumentasi API jika ada perubahan endpoint

### 5. Testing

#### Backend Tests
```bash
cd backend
go test ./... -v
```

#### Frontend Tests
```bash
cd frontend
npm test
```

### 6. Commit dengan Pesan yang Jelas
```bash
git add .
git commit -m "feat: tambahkan fitur X untuk domain Y"
```

Format commit message (Conventional Commits):
- `feat:` - fitur baru
- `fix:` - perbaikan bug
- `docs:` - dokumentasi
- `refactor:` - refactoring
- `test:` - tes
- `chore:` - maintenance
- `perf:` - performance improvement

### 7. Push dan Buat Pull Request
```bash
git push origin fitur/nama-fitur-anda
```

Buat Pull Request di GitHub dengan:
- **Judul yang deskriptif**: "Feat: Add X functionality"
- **Deskripsi perubahan**: Jelaskan apa yang diubah dan kenapa
- **Screenshot**: Jika ada perubahan UI
- **Link ke issue**: Jika PR terkait dengan issue tertentu
- **Checklist**: Tandai apa yang sudah dilakukan

## 📋 Standar Kode

### Backend (Go)
- Gunakan `gofmt` untuk formatting
- Tambahkan error handling yang proper
- Jangan hardcode nilai sensitif (gunakan environment variables)
- Ikuti konvensi penamaan Go: PascalCase untuk exported, camelCase untuk internal
- Business logic hanya di Domain Layer dan Application Layer
- Tidak boleh ada repository access langsung dari Handler

### Frontend (TypeScript)
- Gunakan TypeScript strict mode
- Komponen gunakan PascalCase
- Hooks gunakan camelCase dengan prefix `use`
- Variabel gunakan camelCase
- Konstanta gunakan UPPER_SNAKE_CASE
- State management gunakan TanStack Query + Zustand

## 🏗️ Aturan Arsitektur

### DDD Lite Constraints
⚠️ **PENTING**: Jangan tambahkan hal berikut:
- ❌ CQRS, Event Sourcing, atau Event Bus
- ❌ Command/Query Bus
- ❌ Domain baru di luar MVP
- ❌ Read Models atau Projections
- ✅ Ikuti aggregate boundaries yang sudah ditentukan
- ✅ Business logic hanya di Domain Layer dan Application Layer

### Layered Architecture
```
Handler (HTTP) → Service (Business Logic) → Repository (Data Access) → PostgreSQL
```

- **Handler**: Hanya untuk HTTP request/response, tidak ada business logic
- **Service**: Business logic orchestration, mengkoordinasikan repository dan domain
- **Repository**: Database access saja, tidak ada business logic
- **Domain**: Business rules dan invariants

### Database Changes
- Semua perubahan schema harus melalui migration files
- Migration harus memiliki file `.up.sql` dan `.down.sql`
- Test migration dengan database fresh sebelum PR
- Update `docs/DATABASE_SCHEMA_FREEZE_V1.md` jika ada perubahan signifikan

## 🎨 UX Guidelines

### Teacher-Centric Design
- Guru SD harus bisa menggunakan tanpa training teknis
- Hindari jargon teknis di UI
- Gunakan bahasa Indonesia untuk UI dan UX
- Workflow harus mengikuti alur kerja guru nyata
- Prioritaskan efisiensi kerja guru di atas elegansi teknis

### UX Bahasa
- **UI Labels**: Bahasa Indonesia
- **Error Messages**: Bahasa Indonesia
- **Success Messages**: Bahasa Indonesia
- **Code Comments**: Bahasa Inggris
- **Documentation**: Bahasa Inggris untuk teknis, Indonesia untuk UX

## 🧪 Checklist Review

Sebelum submit PR, pastikan:

### Backend
- [ ] Kode follows architecture freeze (`docs/ARCHITECTURE_FREEZE_V2.md`)
- [ ] Tes unit passing (`go test ./...`)
- [ ] Code formatted (`go fmt ./...`)
- [ ] No hardcoded credentials
- [ ] Error handling proper
- [ ] Documentation updated (jika perlu)

### Frontend
- [ ] TypeScript types tidak ada error
- [ ] Linting tidak ada error
- [ ] UX ramah untuk guru
- [ ] Responsive design (mobile-friendly)
- [ ] No console errors
- [ ] Documentation updated (jika perlu)

### General
- [ ] CHANGELOG.md updated (untuk fitur baru/breaking changes)
- [ ] Commit messages follow Conventional Commits
- [ ] PR description jelas dan lengkap
- [ ] Tidak ada merge conflict

## ⏱️ Timeline Review

Sebagai solo maintainer, timeline saya:

- **Response to PR**: 1-3 hari kerja
- **First Review**: 1-2 hari setelah response
- **Request Changes**: 1 minggu untuk address
- **Merge**: Setelah semua review comments di-address

⚠️ **Catatan**: PR yang tidak aktif > 2 minggu mungkin akan di-close. Jangan ragu untuk re-open jika Anda masih ingin melanjutkan.

## 🐛 Melaporkan Bug

Gunakan [GitHub Issue Template](.github/ISSUE_TEMPLATE/bug_report.md):

1. **Deskripsi Bug**: Jelaskan apa yang terjadi
2. **Langkah Reproduce**: Steps untuk trigger bug
3. **Expected Behavior**: Apa yang seharusnya terjadi
4. **Actual Behavior**: Apa yang sebenarnya terjadi
5. **Environment**: OS, browser, version
6. **Screenshots/Logs**: Jika relevan

## 💡 Request Fitur Baru

Gunakan [GitHub Issue Template](.github/ISSUE_TEMPLATE/feature_request.md):

1. **Use Case**: Jelaskan masalah yang ingin diselesaikan
2. **User Impact**: Bagaimana ini membantu user (guru/admin)
3. **Konteks Domain**: Bagaimana ini relate ke Kurikulum Merdeka
4. **Suggest Implementasi**: (opsional) Ide bagaimana implement

## 💬 Diskusi

Untuk pertanyaan umum atau diskusi:
- Gunakan [GitHub Discussions](https://github.com/sdibonerate85/nusa/discussions)
- Atau buat issue dengan label `question`

## 🌐 Bahasa Komunikasi

- **GitHub Issues**: Bahasa Indonesia atau Inggris (bebas)
- **PR Reviews**: Bahasa Inggris untuk kode, Indonesia untuk UX
- **Discussions**: Bahasa Indonesia (prefer) atau Inggris
- **Code Comments**: Bahasa Inggris

## 🤝 Tipe Kontribusi yang Diapresiasi

- 🐛 Bug fixes
- ✨ Fitur baru (sebisa mungkin sesuai roadmap)
- 📝 Dokumentasi improvement
- 🧪 Tests
- 🎨 UI/UX improvement
- 🌐 Translation (jika ada rencana multi-language)
- 📢 Blog post atau tutorial tentang NUSA

## 🚫 Kontribusi yang Mungkin Ditolak

- ❌ Fitur yang bertentangan dengan architecture freeze
- ❌ Breaking changes tanpa diskusi terlebih dahulu
- ❌ CQRS/Event Sourcing implementation
- ❌ Domain baru di luar scope Kurikulum Merdeka
- ❌ Dependensi baru yang tidak perlu
- ❌ Code yang tidak mengikuti standar proyek

## 📜 Code of Conduct

Silakan baca [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) untuk standar etika komunitas.

## ❓ Pertanyaan?

Jika Anda memiliki pertanyaan:
1. Cek dokumentasi di folder `docs/`
2. Cari di [GitHub Issues](https://github.com/sdibonerate85/nusa/issues)
3. Buat [GitHub Discussion](https://github.com/sdibonerate85/nusa/discussions)
4. Tag saya di PR jika butuh clarification

## 🙏 Terima Kasih

Kontribusi Anda sangat berharga untuk proyek ini dan untuk pendidikan Indonesia! Setiap kontribusi, sekecil apapun, membantu meningkatkan kualitas platform ini.

---

*Dibuat dengan ❤️ oleh solo developer, tapi terbuka untuk kolaborasi komunitas* 🇮🇩
