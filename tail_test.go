package tail

import (
	"testing"

	"github.com/go-zoo/tail/memcache"
)

func TestTemplateCreation(t *testing.T) {
	cache := memcache.New()
	tpl := New("test", "help.go", cache)
	if tpl == nil {
		t.Fail()
	}
}

func TestTemplateCaching(t *testing.T) {
	cache := memcache.New()
	tpl := New("test", "tail_test.go", cache)
	if tpl.get() == nil {
		t.Fail()
	}
}

func TestMultipleTemplate(t *testing.T) {
	cache := memcache.New()
	tpl1 := New("1", "tail_test.go", cache)
	tpl2 := New("2", "tail.go", cache)
	if tpl1.get() == nil || tpl2.get() == nil {
		t.Fail()
	}
}

func TestMultipleCache(t *testing.T) {
	c1 := memcache.New()
	c2 := memcache.New()
	tpl1 := New("1", "tail_test.go", c1)
	tpl2 := New("2", "tail.go", c2)
	if tpl1.get() == nil || tpl2.get() == nil {
		t.Fail()
	}
}
