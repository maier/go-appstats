// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package appstats

import (
	"bytes"
	"testing"
)

func TestNewString(t *testing.T) {
	t.Log("NewString")

	t.Log("valid")
	{
		err := NewString("s1")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}

	t.Log("already exists")
	{
		err := NewString("s1")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}
}

func TestSetString(t *testing.T) {
	t.Log("SetString")

	t.Log("valid, exists")
	{
		err := SetString("s1", "bar")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
		expect := `"bar"`
		v := stats.Get("s1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := SetString("s2", "bar")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := `"bar"`
		v := stats.Get("s2")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestNewInt(t *testing.T) {
	t.Log("NewInt")

	t.Log("valid")
	{
		err := NewInt("i1")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}

	t.Log("already exists")
	{
		err := NewInt("i1")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}
}

func TestSetInt(t *testing.T) {
	t.Log("SetInt")

	t.Log("valid, exists")
	{
		err := SetInt("i1", 10)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "10"
		v := stats.Get("i1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := SetInt("i2", 10)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := `10`
		v := stats.Get("i2")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestIncrementInt(t *testing.T) {
	t.Log("IncrementInt")

	t.Log("valid, exists")
	{
		err := IncrementInt("i1")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "11"
		v := stats.Get("i1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := IncrementInt("i3")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "1"
		v := stats.Get("i3")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestDecrementInt(t *testing.T) {
	t.Log("DecrementInt")

	t.Log("valid, exists")
	{
		err := DecrementInt("i1")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "10"
		v := stats.Get("i1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := DecrementInt("i4")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "-1"
		v := stats.Get("i4")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestAddInt(t *testing.T) {
	t.Log("AddInt")

	t.Log("valid")
	{
		err := AddInt("i1", 10)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "20"
		v := stats.Get("i1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := AddInt("i5", -1)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "-1"
		v := stats.Get("i5")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestNewFloat(t *testing.T) {
	t.Log("NewFloat")

	t.Log("valid")
	{
		err := NewFloat("f1")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}

	t.Log("already exists")
	{
		err := NewFloat("f1")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}
}

func TestSetFloat(t *testing.T) {
	t.Log("SetFloat")

	t.Log("valid, exists")
	{
		err := SetFloat("f1", 1.23)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "1.23"
		v := stats.Get("f1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := SetFloat("f2", 1.23)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := `1.23`
		v := stats.Get("f2")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestAddFloat(t *testing.T) {
	t.Log("AddFloat")

	t.Log("valid, exists")
	{
		err := AddFloat("f1", 3.0)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "4.23"
		v := stats.Get("f1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := AddFloat("f3", 1.23)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := `1.23`
		v := stats.Get("f3")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestNewMap(t *testing.T) {
	t.Log("NewMap")

	t.Log("valid")
	{
		err := NewMap("m1")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}

	t.Log("already exists")
	{
		err := NewMap("m1")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}
}

func TestMapAddInt(t *testing.T) {
	t.Log("MapAddInt")

	t.Log("valid")
	{
		err := MapAddInt("m1", "i1", 1)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
		expect := "1"
		v := maps["m1"].Get("i1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, exists")
	{
		err := MapAddInt("m1", "i1", 2)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "3"
		v := maps["m1"].Get("i1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := MapAddInt("m2", "i1", 2)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "2"
		v := maps["m2"].Get("i1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestMapIncrementInt(t *testing.T) {
	t.Log("MapIncrementInt")

	t.Log("valid")
	{
		err := MapIncrementInt("m1", "i1")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
		expect := "4"
		v := maps["m1"].Get("i1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestMapDecrementInt(t *testing.T) {
	t.Log("MapDecrementInt")

	t.Log("valid")
	{
		err := MapDecrementInt("m1", "i1")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
		expect := "3"
		v := maps["m1"].Get("i1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestMapAddFloat(t *testing.T) {
	t.Log("MapAddFloat")

	t.Log("valid")
	{
		err := MapAddFloat("m1", "f1", 1.23)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
		expect := "1.23"
		v := maps["m1"].Get("f1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, exists")
	{
		err := MapAddFloat("m1", "f1", 1.32)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "2.55"
		v := maps["m1"].Get("f1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestMapSet(t *testing.T) {
	t.Log("MapSet")

	t.Log("valid")
	{
		x := bytes.NewBufferString("baz")
		err := MapSet("m1", "s1", x)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
		expect := "baz"
		v := maps["m1"].Get("s1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, exists")
	{
		x := bytes.NewBufferString("baz2")
		err := MapSet("m1", "s1", x)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "baz2"
		v := maps["m1"].Get("s1")
		if v == nil {
			t.Fatal("expected not nil")
		}
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}
