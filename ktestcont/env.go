package ktestcont

import (
    "os"
    "testing"
)

type Env struct {
    t *testing.T
    sharedDir string
}

// Always use this method to create new instances of `Env`.
func NewEnv(t *testing.T) (*Env, error) {
    sharedDir, err := os.MkdirTemp("/share", "ktestcont-*")
    if err != nil {
        return nil, err
    }
    env := &Env{
        t: t,
        sharedDir: sharedDir,
    }
    return env, nil
}

// Accessor for the shared temporary directory.
func (e *Env) SharedDir() string {
    return e.sharedDir
}

// Clean up any resources that were created instead of Env.  Any errors will be
// logged to the Testing.T instance that was passed to NewEnv().  This is
// intended to be called with a `defer` statement as part of test cleanup.
func (e *Env) Cleanup() {
    if err := os.RemoveAll(e.sharedDir); err != nil {
        e.t.Errorf("ktestcont: Error removing shared temporary directory %s %s", e.sharedDir, err)
    }
}
