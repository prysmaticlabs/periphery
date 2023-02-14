package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func Test_purgeOldFiles(t *testing.T) {
	tmp := filepath.Join(os.TempDir(), "purge-test-run")
	if err := os.Mkdir(tmp, 0777); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.RemoveAll(tmp) })
	for i := 0; i < 10; i++ {
		f, err := os.Create(filepath.Join(tmp, fmt.Sprintf("%d.txt", i)))
		if err != nil {
			t.Fatal(err)
		}
		if err := f.Close(); err != nil {
			t.Fatal(err)
		}
	}
	items, err := os.ReadDir(tmp)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 10 {
		t.Fatal("Not 10 items")
	}

	time.Sleep(time.Second * 2)

	for i := 10; i < 20; i++ {
		f, err := os.Create(filepath.Join(tmp, fmt.Sprintf("%d.txt", i)))
		if err != nil {
			t.Fatal(err)
		}
		if err := f.Close(); err != nil {
			t.Fatal(err)
		}
	}
	items, err = os.ReadDir(tmp)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 20 {
		t.Fatal("Not 20 items")
	}

	if err = purgeOldFiles(tmp, time.Second); err != nil {
		t.Fatal(err)
	}

	items, err = os.ReadDir(tmp)
	if err != nil {
		t.Fatal(err)
	}
	if len(items) != 10 {
		t.Fatal("Not 10 items")
	}
}
