package testman

import (
	"context"
	"strings"
	"testing"
	"time"

	"testman/plugin"
)

type (
	T struct {
		*testing.T

		overrides plugin.Overrides
	}

	concreteT = T
)

func (*T) New(t *T, _ ...plugin.Option) *T { return t }

func (t *T) Name() string {
	t.Helper()

	return t.overrides.Name.Call(t.name)()
}

func (t *T) name() string {
	name := t.T.Name()

	idx := strings.Index(name, wrapperTestName)

	if idx >= 0 {
		// +1 for slash
		return name[idx+len(wrapperTestName)+1:]
	}

	return name
}

// Parallel signals that this test is to be run in parallel with (and only with)
// other parallel tests. When a test is run multiple times due to use of
// -test.count or -test.cpu, multiple instances of a single test never run in
// parallel with each other.
func (t *T) Parallel() {
	t.Helper()

	t.overrides.Parallel.Call(t.T.Parallel)()
}

// Chdir calls os.Chdir(dir) and uses Cleanup to restore the current
// working directory to its original value after the test. On Unix, it
// also sets PWD environment variable for the duration of the test.
//
// Because Chdir affects the whole process, it cannot be used
// in parallel tests or tests with parallel ancestors.
func (t *T) Chdir(dir string) {
	t.Helper()

	t.overrides.Chdir.Call(t.T.Chdir)(dir)
}

// Setenv calls os.Setenv(key, value) and uses Cleanup to
// restore the environment variable to its original value
// after the test.
//
// Because Setenv affects the whole process, it cannot be used
// in parallel tests or tests with parallel ancestors.
func (t *T) Setenv(key, value string) {
	t.Helper()

	t.overrides.Setenv.Call(t.T.Setenv)(key, value)
}

// TempDir returns a temporary directory for the test to use.
// The directory is automatically removed when the test and
// all its subtests complete.
// Each subsequent call to t.TempDir returns a unique directory;
// if the directory creation fails, TempDir terminates the test by calling Fatal.
func (t *T) TempDir() string {
	t.Helper()

	return t.overrides.TempDir.Call(t.T.TempDir)()
}

// Log formats its arguments using default formatting, analogous to Println,
// and records the text in the error log. For tests, the text will be printed only if
// the test fails or the -test.v flag is set. For benchmarks, the text is always
// printed to avoid having performance depend on the value of the -test.v flag.
func (t *T) Log(args ...any) {
	t.Helper()

	t.overrides.Log.Call(t.T.Log)(args...)
}

// Logf formats its arguments according to the format, analogous to Printf, and
// records the text in the error log. A final newline is added if not provided. For
// tests, the text will be printed only if the test fails or the -test.v flag is
// set. For benchmarks, the text is always printed to avoid having performance
// depend on the value of the -test.v flag.
func (t *T) Logf(format string, args ...any) {
	t.Helper()

	t.overrides.Logf.Call(t.T.Logf)(format, args...)
}

// Context returns a context that is canceled just before
// Cleanup-registered functions are called.
//
// Cleanup functions can wait for any resources
// that shut down on Context.Done before the test or benchmark completes.
func (t *T) Context() context.Context {
	t.Helper()

	return t.overrides.Context.Call(t.T.Context)()
}

// Deadline reports the time at which the test binary will have
// exceeded the timeout specified by the -timeout flag.
//
// The ok result is false if the -timeout flag indicates “no timeout” (0).
func (t *T) Deadline() (deadline time.Time, ok bool) {
	t.Helper()

	return t.overrides.Deadline.Call(t.T.Deadline)()
}

// Errorf is equivalent to Logf followed by Fail.
func (t *T) Errorf(format string, args ...any) {
	t.Helper()

	t.overrides.Errorf.Call(t.T.Errorf)(format, args...)
}

// Error is equivalent to Log followed by Fail.
func (t *T) Error(args ...any) {
	t.Helper()

	t.overrides.Error.Call(t.T.Error)(args...)
}

// Skip is equivalent to Log followed by SkipNow.
func (t *T) Skip(args ...any) {
	t.Helper()

	t.overrides.Skip.Call(t.T.Skip)(args...)
}

// SkipNow marks the test as having been skipped and stops its execution
// by calling [runtime.Goexit].
// If a test fails (see Error, Errorf, Fail) and is then skipped,
// it is still considered to have failed.
// Execution will continue at the next test or benchmark. See also FailNow.
// SkipNow must be called from the goroutine running the test, not from
// other goroutines created during the test. Calling SkipNow does not stop
// those other goroutines.
func (t *T) SkipNow() {
	t.Helper()

	t.overrides.SkipNow.Call(t.T.SkipNow)()
}

// Skipf is equivalent to Logf followed by SkipNow.
func (t *T) Skipf(format string, args ...any) {
	t.Helper()

	t.overrides.Skipf.Call(t.T.Skipf)(format, args...)
}

// Skipped reports whether the test was skipped.
func (t *T) Skipped() bool {
	t.Helper()

	return t.overrides.Skipped.Call(t.T.Skipped)()
}

// Fail marks the function as having failed but continues execution.
func (t *T) Fail() {
	t.Helper()

	t.overrides.Fail.Call(t.T.Fail)()
}

// FailNow marks the function as having failed and stops its execution
// by calling runtime.Goexit (which then runs all deferred calls in the
// current goroutine).
// Execution will continue at the next test or benchmark.
// FailNow must be called from the goroutine running the
// test or benchmark function, not from other goroutines
// created during the test. Calling FailNow does not stop
// those other goroutines.
func (t *T) FailNow() {
	t.Helper()

	t.overrides.FailNow.Call(t.T.FailNow)()
}

// Failed reports whether the function has failed.
func (t *T) Failed() bool {
	t.Helper()

	return t.overrides.Failed.Call(t.T.Failed)()
}

// Fatal is equivalent to Log followed by FailNow.
func (t *T) Fatal(args ...any) {
	t.Helper()

	t.overrides.Fatal.Call(t.T.Fatal)(args...)
}

// Fatalf is equivalent to Logf followed by FailNow.
func (t *T) Fatalf(format string, args ...any) {
	t.Helper()

	t.overrides.Fatalf.Call(t.T.Fatalf)(format, args...)
}

func (t *T) BaseName() string {
	segments := strings.Split(t.Name(), "/")

	if len(segments) == 0 {
		return ""
	}

	return segments[len(segments)-1]
}

func (t *T) unwrap() *T {
	return t
}
