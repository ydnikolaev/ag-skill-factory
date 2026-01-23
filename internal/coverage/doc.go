// Package coverage provides meta-testing for package coverage enforcement.
//
// This package ensures that all packages in the project have corresponding
// test files. It enforces the TDD Hard Stop Protocol by failing the build
// if any package is missing tests.
//
// The test discovers all Go packages and verifies each has at least one
// *_test.go file, enforcing test coverage at the structural level.
package coverage
