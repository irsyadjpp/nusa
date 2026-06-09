# Pull Request Template

## Description
<!-- 
Brief description of the changes made in this PR.
What does this PR do? Why is it needed?
-->

## Type of Change
<!-- Mark the relevant option with an "x" -->

- [ ] Bug fix (non-breaking change which fixes an issue)
- [ ] New feature (non-breaking change which adds functionality)
- [ ] Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] Documentation update
- [ ] Refactoring (no functional changes)
- [ ] Performance improvement
- [ ] Test update/addition

## Related Issues
<!-- 
Link to related issues using GitHub syntax:
Fixes #123
Related to #456
-->

## Changes Made
<!-- 
Describe the specific changes made:
- Backend: ...
- Frontend: ...
- Database: ...
- Documentation: ...
-->

### Backend Changes
- [ ] Domain layer changes
- [ ] Service layer changes
- [ ] Repository changes
- [ ] Handler/API changes
- [ ] Database migrations
- [ ] Tests added/updated

### Frontend Changes
- [ ] Component changes
- [ ] Page changes
- [ ] State management changes
- [ ] API integration changes
- [ ] UI/UX improvements
- [ ] Tests added/updated

## Testing
<!-- 
Describe how you tested these changes:
- Manual testing
- Automated tests
- Integration tests
-->

### Backend Tests
```bash
# Commands to run backend tests
cd backend
go test ./... -v
```
- [ ] All tests passing

### Frontend Tests
```bash
# Commands to run frontend tests
cd frontend
npm test
```
- [ ] All tests passing

### Manual Testing
- [ ] Tested locally
- [ ] Tested in development environment
- [ ] Browser(s) tested: Chrome, Firefox, Safari

## Screenshots (if applicable)
<!-- Add screenshots for UI changes -->

## Checklist
<!-- Mark the relevant options with an "x" -->

- [ ] My code follows the project's architecture guidelines (ARCHITECTURE_FREEZE_V2.md)
- [ ] I have performed a self-review of my code
- [ ] I have commented my code, particularly in hard-to-understand areas
- [ ] I have made corresponding changes to the documentation
- [ ] My changes generate no new warnings
- [ ] I have added tests that prove my fix is effective or that my feature works
- [ ] New and existing unit tests pass locally with my changes
- [ ] I have updated the CHANGELOG.md for significant changes
- [ ] I have followed the Conventional Commits specification for commit messages

## Architecture Compliance
<!-- 
This is a solo-developer project with strict architecture governance.
Please confirm your changes comply with the architecture freeze:
-->

- [ ] Follows DDD Lite principles (no CQRS, Event Sourcing, Event Bus)
- [ ] Adheres to layered architecture: Handler → Service → Repository → PostgreSQL
- [ ] No business logic in handlers
- [ ] No direct repository access from handlers
- [ ] Uses existing aggregates and bounded contexts
- [ ] No new domains outside Kurikulum Merdeka scope

## Breaking Changes
<!-- 
If this PR includes breaking changes, please describe them here.
This is critical for a solo-maintained project.
-->

- [ ] No breaking changes
- [ ] Breaking changes documented below:
  - 
  - 

## Additional Context
<!-- 
Any other context about the PR:
- Dependencies added
- Performance considerations
- Security considerations
- Deployment considerations
-->

## Notes for Reviewer
<!-- 
Any specific areas you'd like the reviewer to focus on:
-->

---

## Solo Developer Notes

As the sole maintainer, I appreciate your contribution! Here's what to expect:

- **Review Timeline**: 1-3 business days (depending on availability)
- **Review Focus**: Architecture compliance, code quality, testing, documentation
- **Response Style**: I'll provide detailed feedback and work with you iteratively
- **Merge Decision**: I'll merge once all feedback is addressed and tests pass

Thank you for contributing to NUSA Platform! 🇮🇩
