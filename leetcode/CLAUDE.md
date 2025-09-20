# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview
Quickies is a Go workspace for algorithm practice, coding challenges, and quick programming experiments. The main focus is the `leetcode/` directory containing categorized algorithm implementations.

## Tutoring Approach
When working with users on LeetCode problems, act as a tutor rather than providing direct solutions. Your role is to:
- Guide users through the problem-solving process step by step
- Ask leading questions to help them discover the solution approach
- Explain underlying concepts and algorithms (time/space complexity, data structures, patterns)
- Help them understand why certain approaches work better than others
- Encourage them to think through edge cases and test scenarios
- Only provide code hints or partial solutions when they're stuck, not complete implementations
- Focus on building their problem-solving intuition and algorithmic thinking skills

## Essential Commands

### Development
**IMPORTANT**: All commands must be run from the `/leetcode/` directory (module root).

- `go run <file.go>` - Run a single Go file
- `go fmt ./...` - Format all Go files in the project  
- `go mod tidy` - Clean up module dependencies
- `go vet ./...` - Run Go's static analysis tool

### Linting
- `golangci-lint run` - Run comprehensive linter with project configuration
- `golangci-lint run ./...` - Run linter on all packages
- `golangci-lint run --fix` - Auto-fix issues where possible

### Testing
- `go test ./...` - Run all tests in the project
- `go test ./arrays/` - Run tests for specific package
- `go test -v ./...` - Run tests with verbose output
- `go test -cover ./...` - Run tests with coverage report
- `go test -run <TestName>` - Run specific test function

### Building
- `go build ./...` - Build all packages (check for compilation errors)
- `go build <file.go>` - Build specific file

### Pre-commit Hooks
The repository has an active pre-commit hook that automatically runs:
1. `go test ./...` - All tests must pass
2. `golangci-lint run` - Linting must pass with no errors
3. `go fmt ./...` - Code must be properly formatted
Commits will be blocked if any of these checks fail.

## Code Architecture

### Package Structure
All LeetCode solutions use `package leetcode` and are organized into subdirectories by problem category:
- **arrays/** - Array manipulation and optimization algorithms
- **strings/** - String processing and text algorithms  
- **linked-lists/** - Linked list operations (defines shared `ListNode` struct)
- **hash-tables/** - Hash map/set-based lookup solutions and memoization
- **trees/** - Binary tree algorithms (defines shared `TreeNode` struct and queue implementations)

### Function Naming Convention
- Functions use camelCase and are unexported (lowercase first letter)
- File names use snake_case matching the problem name
- Each function includes a detailed comment with the problem description

### Code Patterns
- Solutions prioritize optimal time complexity over readability when appropriate
- Use Go's built-in `map[type]type` for hash table solutions
- Linked list problems define `ListNode` struct with `Val int` and `Next *ListNode`
- Tree problems use shared `TreeNode` struct and queue implementations in `trees/types.go`
- Functions typically return the most direct solution type (indices, booleans, modified structures)
- No external dependencies - pure Go standard library only
- Generic queue interface and implementations available for complex data structure problems

### Testing Strategy
When adding tests:
- Create `*_test.go` files alongside implementation files
- Use standard Go testing package with table-driven test pattern
- Test edge cases including empty inputs, single elements, and boundary conditions
- Include benchmark tests for performance-critical algorithms
- Use helper functions for complex data setup (e.g., `createLinkedList`, `linkedListToSlice`)
- Follow naming convention: `TestFunctionName` and `BenchmarkFunctionName`

### Project Configuration
- **golangci-lint**: Comprehensive linter with security checks (`gosec`), complexity analysis (`gocyclo`), and code quality rules
- **Go version**: 1.25.0 with module support
- **Module name**: `leetcode` (module root is in the `leetcode/` directory)
