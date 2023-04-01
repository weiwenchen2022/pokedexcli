package pokecache

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestAddGet(t *testing.T) {
	t.Parallel()

	const internal = 5 * time.Second

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("(%q, %q)", c.key, c.val), func(t *testing.T) {
			cache := NewCache(internal)
			cache.Add(c.key, c.val)

			val, ok := cache.Get(c.key)
			if !ok {
				t.Error("expected to find key")
			}

			if diff := cmp.Diff(string(c.val), string(val)); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	t.Parallel()

	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

	cache := NewCache(baseTime)

	const key = "https://example.com"
	cache.Add(key, []byte("testdata"))

	_, ok := cache.Get(key)
	if !ok {
		t.Error("expected to find key")
	}

	<-time.After(waitTime)
	_, ok = cache.Get(key)
	if ok {
		t.Error("expected to not find key")
	}
}
