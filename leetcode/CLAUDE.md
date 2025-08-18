# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview
Quickies is a Go workspace for algorithm practice, coding challenges, and quick programming experiments. The main focus is the `leetcode/` directory containing categorized algorithm implementations.

## Essential Commands

### Development
- `go run <file.go>` - Run a single Go file
- `go fmt ./...` - Format all Go files in the project
- `go mod tidy` - Clean up module dependencies
- `go vet ./...` - Run Go's static analysis tool

### Testing
- `go test ./...` - Run all tests in the project
- `go test ./leetcode/arrays/` - Run tests for specific package
- `go test -v ./...` - Run tests with verbose output
- `go test -cover ./...` - Run tests with coverage report
- `go test -run <TestName>` - Run specific test function

### Building
- `go build ./...` - Build all packages (check for compilation errors)
- `go build <file.go>` - Build specific file

## Code Architecture

### Package Structure
All LeetCode solutions use `package leetcode` and are organized into subdirectories by problem category:
- **arrays/** - Array manipulation and optimization algorithms
- **strings/** - String processing and text algorithms  
- **linked-lists/** - Linked list operations (defines shared `ListNode` struct)
- **hash-tables/** - Hash map/set-based lookup solutions

### Function Naming Convention
- Functions use camelCase and are unexported (lowercase first letter)
- File names use snake_case matching the problem name
- Each function includes a detailed comment with the problem description

### Code Patterns
- Solutions prioritize optimal time complexity over readability when appropriate
- Use Go's built-in `map[type]type` for hash table solutions
- Linked list problems define `ListNode` struct with `Val int` and `Next *ListNode`
- Functions typically return the most direct solution type (indices, booleans, modified structures)
- No external dependencies - pure Go standard library only

### Testing Strategy
When adding tests:
- Create `*_test.go` files alongside implementation files
- Use standard Go testing package
- Test edge cases including empty inputs, single elements, and boundary conditions
- Include benchmark tests for performance-critical algorithms
