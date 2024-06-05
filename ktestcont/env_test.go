package ktestcont

import (
    "errors"
    "os"
    "testing"
)

func TestEnv(t *testing.T) {
    t.Parallel()

    env, err := NewEnv(t)
    if err != nil {
        t.Fatal(err)
    }

    sharedDir := env.SharedDir()
    stat, err := os.Stat(sharedDir)
    if err != nil {
        t.Fatal(err)
    }

    if !stat.IsDir() {
        t.Error("env.SharedDir() is not a directory")
    }

    env.Cleanup()

    if _, err := os.Stat(sharedDir); !errors.Is(err, os.ErrNotExist) {
        t.Errorf("Expected sharedDir %s to not exist, instead saw error %s", sharedDir, err)
    }
}
