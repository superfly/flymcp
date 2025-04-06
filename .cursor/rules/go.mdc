---
description: 
globs: **/*.go
alwaysApply: false
---
# Go Development Rules

## Directory Structure
- Use simple Go project layout
  - ./cmd/ for individual subcommands
  - 
- NEVER create a pkg or internal directory

## Code Changes
- Minimize file changes by keeping related functionality together
- Prefer modifying existing files over creating new ones
- Use interfaces to reduce coupling between components

## Library use
- ALWAYS use the go std library when writing new code
- NEVER use a third party package without explicit user request

## Testing
- Tests MUST fail if the code they're testing isn't implemented
- NEVER modify tests to make them pass without implementing the required functionality
- Write tests before implementing features (TDD approach)