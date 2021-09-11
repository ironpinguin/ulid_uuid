package main

import (
	"testing"
)

func TestToUlidFail(t *testing.T) {
	_, err := toUlid("jfkldsfaj")
	if err == nil {
		t.Error("expected an error")
	}
}

func TestToUlidSucess(t *testing.T) {
	result, err := toUlid("cfa45f5d-9c38-4772-b39a-036a0b9f8d30")
	if err != nil {
		t.Errorf("not expected error: %v", err)
	}
	if result != "6FMHFNV71R8XSB76G3D85SZ39G" {
		t.Errorf("Result %s is not the expected ulid '6FMHFNV71R8XSB76G3D85SZ39G'", result)
	}
}

func TestToUuidFail(t *testing.T) {
	_, err := toUuid("ADSJFKEWIFJWFEW")
	if err == nil {
		t.Error("expected an error")
	}
}

func TestToUuidSucess(t *testing.T) {
	result, err := toUuid("08A1YW3WAH8SNTQVYGDB2EP69T")
	if err != nil {
		t.Errorf("not expected error: %v", err)
	}
	if result != "08507dc1-f151-466b-abef-d06ac4eb193a" {
		t.Errorf("Result %s is not the expected uuid '08507dc1-f151-466b-abef-d06ac4eb193a'", result)
	}
}
