package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	data []string
	options Options
	expected []string
}{
	{
		[]string {
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
		},
		NewOptions(
			false,
			true,
			true,
			0,
			0,
			false,
		),
		[]string {
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		},
	},
	{
		[]string {
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
		},
		NewOptions(
			true,
			true,
			true,
			0,
			0,
			false,
		),
		[]string {
			"3 I love music.",
			"1 ",
			"2 I love music of Kartik.",
			"1 Thanks.",
		},
	},
	{
		[]string {
			"I love music.",
			"I love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
		},
		NewOptions(
			false,
			true,
			false,
			0,
			0,
			false,
		),
		[]string {
			"I love music.",
			"I love music of Kartik.",
		},
	},
	{
		[]string {
			"I love music.",
			"i love music.",
			"I love music.",
			"",
			"I love music of Kartik.",
			"I love music of Kartik.",
			"Thanks.",
		},
		NewOptions(
			false,
			true,
			true,
			0,
			0,
			true,
		),
		[]string {
			"I love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		},
	},
	{
		[]string {
			"We love music.",
			"I love music.",
			"They love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		NewOptions(
			false,
			true,
			true,
			1,
			0,
			false,
		),
		[]string {
			"We love music.",
			"",
			"I love music of Kartik.",
			"Thanks.",
		},
	},
	{
		[]string {
			"I love music.",
			"A love music.",
			"C love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
		NewOptions(
			false,
			true,
			true,
			0,
			1,
			false,
		),
		[]string {
			"I love music.",
			"",
			"I love music of Kartik.",
			"We love music of Kartik.",
			"Thanks.",
		},
	},
}

func TestUniq(t *testing.T) {
	for _, e := range tests {
		result := Uniq(e.options, e.data)

		if !reflect.DeepEqual(e.expected, result) {
			t.Fatalf("Test has been failed")
		}
	}
}
