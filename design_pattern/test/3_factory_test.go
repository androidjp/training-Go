package test

import (
	"../factory"
	"testing"
)

func TestEventFactory(t *testing.T) {
	f := factory.Factory{}
	e := f.Create(factory.Start)
	if e.EventType() != factory.Start {
		t.Errorf("expect event.Start, but actual %v.", e.EventType())
	}
	e = f.Create(factory.End)
	if e.EventType() != factory.End {
		t.Errorf("expect event.End, but actual %v.", e.EventType())
	}
}

func TestEventFactory2(t *testing.T) {
	e := factory.OfStart()
	if e.EventType() != factory.Start {
		t.Errorf("expect event.Start, but actual %v.", e.EventType())
	}
	e = factory.OfEnd()
	if e.EventType() != factory.End {
		t.Errorf("expect event.End, but actual %v.", e.EventType())
	}
}
