# Contributing to Ellie

Thank you for your interest in contributing to Ellie! This document provides guidelines and instructions for contributing to the project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Testing](#testing)
- [Code Style](#code-style)
- [Pull Request Process](#pull-request-process)
- [Documentation](#documentation)

## Code of Conduct

By participating in this project, you agree to abide by our Code of Conduct. Please be respectful and considerate of others.

### Our Standards

- Be respectful and inclusive
- Be open to different viewpoints and experiences
- Give and gracefully accept constructive feedback
- Focus on what is best for the community
- Show empathy towards other community members

### Unacceptable Behavior

- Harassment, discrimination, or offensive comments
- Trolling, insulting, or derogatory remarks
- Public or private harassment
- Publishing others' private information without permission
- Other conduct which could reasonably be considered inappropriate

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/your-username/ellie.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Push to your fork: `git push origin feature/your-feature-name`
6. Create a Pull Request

## Development Workflow

1. Always work on a feature branch
2. Keep your branch up to date with main
3. Write tests for new features
4. Ensure all tests pass
5. Update documentation as needed
6. Submit a Pull Request

## Testing

We use Go's built-in testing framework. To run tests:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test ./action/...
```

### Test Coverage

We aim to maintain high test coverage. To generate a coverage report:

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Code Style

- Follow Go's standard formatting: `go fmt`
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions small and focused
- Handle errors appropriately
- Use interfaces for better testability

## Pull Request Process

1. Update the README.md with details of changes if needed
2. Update the documentation if needed
3. Ensure tests pass
4. Request review from maintainers
5. Address any feedback
6. Once approved, squash and merge

## Documentation

- Keep documentation up to date
- Use clear and concise language
- Include examples where helpful
- Document any breaking changes
- Update README.md for significant changes

## Questions?

Feel free to open an issue if you have any questions about contributing to Ellie. 