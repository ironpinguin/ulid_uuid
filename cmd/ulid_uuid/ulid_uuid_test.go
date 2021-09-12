package main

import (
	"testing"

	"github.com/google/uuid"
)

func TestToUlidFail(t *testing.T) {
	_, err := toUlid("jfkldsfaj")
	if err == nil {
		t.Error("expected an error")
	}
	if !uuid.IsInvalidLengthError(err) {
		t.Errorf("unexpcted Error type %v", err)
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

func TestToUlidFail2(t *testing.T) {
	_, err := toUlid("cfa45f5k-9c38-4772-!39a-036a0b9f8d30")
	if err == nil {
		t.Errorf("expected an error")
	}
	if err.Error() != "invalid UUID format" {
		t.Errorf("wrong error: %v", err)
	}
}

func TestToUuidFail(t *testing.T) {
	_, err := toUUID("ADSJFKEWIFJWFEW")
	if err == nil {
		t.Error("expected an error")
	}
}

func TestToUuidFailUnMarshalOverflow(t *testing.T) {
	_, err := toUUID("FFMHFNV71R8XSB76G3D85SZ39G")
	if err == nil {
		t.Error("expected an error")
	}
}

func TestToUuidSucess(t *testing.T) {
	result, err := toUUID("08A1YW3WAH8SNTQVYGDB2EP69T")
	if err != nil {
		t.Errorf("not expected error: %v", err)
	}
	if result != "08507dc1-f151-466b-abef-d06ac4eb193a" {
		t.Errorf("Result %s is not the expected uuid '08507dc1-f151-466b-abef-d06ac4eb193a'", result)
	}
}

func TestCheckType(t *testing.T) {
	var result string
	var err error

	result, err = checkType("08A1YW3WAH8SNTQVYGDB2EP69T")
	if err != nil {
		t.Errorf("not expected error: %v", err)
	}
	if result != "08507dc1-f151-466b-abef-d06ac4eb193a" {
		t.Errorf("Result %s is not the expected ulid '08507dc1-f151-466b-abef-d06ac4eb193a'", result)
	}

	_, err = checkType("FFMHFNV71R8XSB76G3D85SZ39G")
	if err == nil {
		t.Errorf("expected error")
	}

	result, err = checkType("cfa45f5d-9c38-4772-b39a-036a0b9f8d30")
	if err != nil {
		t.Errorf("not expected error: %v", err)
	}
	if result != "6FMHFNV71R8XSB76G3D85SZ39G" {
		t.Errorf("Result %s is not the expected uuid '6FMHFNV71R8XSB76G3D85SZ39G'", result)
	}

	_, err = checkType("cfa45f5k-9c38-4772-!39a-036a0b9f8d30")
	if err == nil {
		t.Errorf("expected error")
	}
}
