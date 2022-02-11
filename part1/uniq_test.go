package main

import (
	"reflect"
	"testing"
	"./uniq"
)

func TestDefaultSuccess(t *testing.T) {
	data := []string {
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	options := uniq.NewOptions(
		false,
		true,
		true,
		0,
		0,
		false,
	)
	expected := []string {
		"I love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := uniq.Uniq(options, data)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Test with default behavior failed")
	}
}

func TestFlagCSuccess(t *testing.T) {
	data := []string {
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	options := uniq.NewOptions(
		true,
		true,
		true,
		0,
		0,
		false,
	)
	expected := []string {
		"3 I love music.",
		"1 ",
		"2 I love music of Kartik.",
		"1 Thanks.",
	}
	result := uniq.Uniq(options, data)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Test with -c failed")
	}
}

func TestFlagDSuccess(t *testing.T) {
	data := []string {
		"I love music.",
		"I love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	options := uniq.NewOptions(
		false,
		true,
		false,
		0,
		0,
		false,
	)
	expected := []string {
		"I love music.",
		"I love music of Kartik.",
	}
	result := uniq.Uniq(options, data)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Test with -d failed")
	}
}

func TestFlagISuccess(t *testing.T) {
	data := []string {
		"I love music.",
		"i love music.",
		"I love music.",
		"",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}
	options := uniq.NewOptions(
		false,
		true,
		true,
		0,
		0,
		true,
	)
	expected := []string {
		"I love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := uniq.Uniq(options, data)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Test with -d failed")
	}
}

func TestFlagFSuccess(t *testing.T) {
	data := []string {
		"We love music.",
		"I love music.",
		"They love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	options := uniq.NewOptions(
		false,
		true,
		true,
		1,
		0,
		false,
	)
	expected := []string {
		"We love music.",
		"",
		"I love music of Kartik.",
		"Thanks.",
	}
	result := uniq.Uniq(options, data)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Test with -f failed")
	}
}

func TestFlagSSuccess(t *testing.T) {
	data := []string {
		"I love music.",
		"A love music.",
		"C love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	options := uniq.NewOptions(
		false,
		true,
		true,
		0,
		1,
		false,
	)
	expected := []string {
		"I love music.",
		"",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}
	result := uniq.Uniq(options, data)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Test with -s failed")
	}
}
