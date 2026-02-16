package main

import (
	"bytes"
	"io"
	"testing"
)

func TestPathTransformFunc(t *testing.T) {
	key := "momsbestpicture"
	pathKey := CASPathTransformFunc(key)
	expectedOriginalKey := "6804429f74181a63c50c3d81d733a12f14a353ff"
	expectedPathname := "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff"
	if pathKey.Pathname != expectedPathname {
		t.Errorf("expected %s, got %s", expectedPathname, pathKey.Pathname)
	}
	if pathKey.Filename != expectedOriginalKey {
		t.Errorf("expected %s, got %s", expectedOriginalKey, pathKey.Filename)
	}
}
func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	key := "momspecials"
	data := []byte("some jpg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Errorf("error writing stream: %s", err)
	}
	if err := s.Delete(key); err != nil {
		t.Errorf("error deleting key: %s", err)
	}
}

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)

	key := "momspecial"
	data := []byte("some jpg bytes")
	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Errorf("error writing stream: %s", err)
	}
	if ok := s.Has(key); !ok {
		t.Errorf("expected key %s to exist", key)
	}
	r, err := s.Read(key)
	if err != nil {
		t.Errorf("error reading stream: %s", err)
	}
	b, _ := io.ReadAll(r)
	if string(b) != "some jpg bytes" {
		t.Errorf("expected %s, got %s", "some jpg bytes", string(b))
	}
	if err := s.Delete(key); err != nil {
		t.Errorf("error deleting key: %s", err)
	}
}
