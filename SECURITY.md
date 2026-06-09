# Security Policy

## Supported Versions

Versi NUSA Platform yang saat ini didukung dengan update keamanan:

| Version | Supported          | Security Updates |
|---------|--------------------|------------------|
| 0.3.x   | :white_check_mark: | :white_check_mark: |
| 0.2.x   | :white_check_mark: | :white_check_mark: |
| < 0.2.0 | :x:                | :x:              |

## Reporting a Vulnerability

### How to Report

Jika Anda menemukan kerentanan keamanan di NUSA Platform, harap **JANGAN** buat public issue. Sebagai gantinya:

1. **Kirim Email**: security@nusa.local (jika ada) atau
2. **GitHub Private Security Advisory**: Gunakan fitur [Security Advisory](https://github.com/sdibonerate85/nusa/security/advisories)
3. **Include Detail**:
   - Deskripsi kerentanan
   - Langkah untuk reproduce
   - Potential impact
   - Suggested fix (jika ada)

### Response Timeline

Sebagai solo maintainer, timeline saya:

- **Initial Response**: 3-5 hari kerja
- **Investigation**: 1-2 minggu tergantung kompleksitas
- **Fix Development**: Variasi tergantung severity
- **Public Disclosure**: Setelah fix tersedia

### Response Process

1. **Acknowledgment**: Saya akan konfirmasi penerimaan report
2. **Validation**: Investigasi dan konfirmasi kerentanan
3. **Fix**: Develop dan test patch
4. **Release**: Release patch version
5. **Disclosure**: Public disclosure dengan credit (jika diminta)

## Security Best Practices

### For Deployments

#### Production Setup
- **Change Default Credentials**: Selalu ubah default admin password
- **Environment Variables**: Jangan hardcode secrets di kode
- **HTTPS**: Gunakan HTTPS di production
- **Database Security**: Enable SSL/TLS untuk PostgreSQL
- **Firewall**: Limit access ke database dan internal services
- **Regular Updates**: Keep dependencies updated

#### Secrets Management
- Jangan commit `.env` files ke git
- Gunakan secret management service di production (AWS Secrets Manager, HashiCorp Vault, dll)
- Rotate secrets secara berkala
- Gunakan strong passwords (minimal 16 karakter dengan mix character types)

#### Network Security
- Jangan expose PostgreSQL langsung ke internet
- Gunakan VPN atau private network untuk admin access
- Configure proper firewall rules
- Rate limiting pada API endpoints

### For Development

#### Code Security
- **Input Validation**: Selalu validate dan sanitize input
- **SQL Injection**: Gunakan parameterized queries (sqlx sudah handle ini)
- **XSS Prevention**: React sudah auto-escape, tapi tetap hati-hati dengan dangerouslySetInnerHTML
- **Authentication**: Jangan implement custom crypto, gunakan library terpercaya
- **Authorization**: Selalu check permissions di server-side

#### Dependencies
```bash
# Audit dependencies regularly
cd backend
go mod verify
go list -json -m all | nancy sleuth

cd frontend
npm audit
npm audit fix
```

#### Code Review
- Review semua code changes dengan security lens
- Perhatikan data flow dan permission checks
- Test authentication dan authorization thoroughly

## Known Security Considerations

### Current Implementation

#### Authentication
- **JWT Tokens**: Dengan expiration dan refresh mechanism
- **Password Hashing**: Menggunakan bcrypt dengan cost factor yang appropriate
- **Session Management**: JWT stateless dengan refresh token rotation

#### Authorization
- **RBAC**: Role-based access control (SYSTEM_ADMIN, ADMIN, TEACHER, PRINCIPAL)
- **School Isolation**: Multi-tenancy dengan school-level data isolation
- **Resource Ownership**: User hanya bisa akses resource dari school mereka

#### Data Protection
- **SQL Injection**: Prevented oleh sqlx parameterized queries
- **XSS**: React auto-escaping, tapi hati-hati dengan user-generated content
- **CSRF**: Need implementation di frontend (currently missing)
- **CORS**: Configurable, default strict di production

### Future Improvements

- [ ] CSRF protection implementation
- [ ] Rate limiting per endpoint
- [ ] Security headers (CSP, X-Frame-Options, etc.)
- [ ] Audit logging untuk sensitive operations
- [ ] 2FA untuk admin accounts
- [ ] Session timeout configuration
- [ ] IP whitelist untuk admin access

## Dependency Security

### Backend (Go)
- **Gin Web Framework**: Regularly updated
- **sqlx**: Stable, actively maintained
- **bcrypt**: Standard crypto library
- **JWT-Go**: Need monitor for updates

### Frontend (React/TypeScript)
- **React**: Regular security updates
- **Axios**: Monitor for security patches
- **Material-UI**: Follow security advisories

### Infrastructure
- **PostgreSQL 18.4**: Latest stable version
- **Podman/Docker**: Regular security updates
- **Nginx** (if used as reverse proxy): Monitor for vulnerabilities

## Security Testing

### Automated Testing
```bash
# Backend security scanning (if tools available)
cd backend
gosec ./...

# Frontend dependency audit
cd frontend
npm audit
```

### Manual Testing Checklist
- [ ] Test SQL injection attempts
- [ ] Test XSS attempts
- [ ] Test authentication bypass
- [ ] Test authorization bypass (horizontal/vertical privilege escalation)
- [ ] Test session hijacking
- [ ] Test CSRF (when implemented)
- [ ] Test file upload security (if applicable)

## Security Incidents

### Incident Response Process

Jika terjadi security incident:

1. **Identification**: Detect dan confirm incident
2. **Containment**: Limit impact (take offline jika perlu)
3. **Eradication**: Remove root cause
4. **Recovery**: Restore dari backup jika perlu
5. **Lessons Learned**: Document dan improve process

### Post-Incident
- Root cause analysis
- Implement preventive measures
- Update documentation
- Consider disclosure jika user data teraffected

## Compliance

### Data Privacy
- Data siswa dan guru adalah sensitive
- Ikuti regulasi privacy Indonesia (jika ada)
- Implement data retention policy
- Secure deletion saat data dihapus

### Educational Data
- Sesuaikan dengan Kemendikbudristek requirements (jika ada)
- Protect student records
- Teacher data protection
- Assessment data security

## Contact

### Security Questions
Untuk pertanyaan security yang bukan vulnerability report:
- **GitHub Discussions**: https://github.com/sdibonerate85/nusa/discussions
- **Label**: Gunakan label `security` untuk pertanyaan security

### Vulnerability Reporting
- **GitHub Security Advisory**: https://github.com/sdibonerate85/nusa/security/advisories
- **Email**: (jika email security disiapkan)

## Security Credits

Terima kasih kepada semua researcher yang bertanggung jawab dalam reporting vulnerabilities. Anda akan diberi credit di security advisory dan changelog.

---

*Last Updated: 2026-06-09*
*Maintainer: Solo Developer*
