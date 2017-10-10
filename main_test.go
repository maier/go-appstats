// Copyright © 2017 Circonus, Inc. <support@circonus.com>
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package appstats

import "testing"

func TestNewString(t *testing.T) {
	t.Log("NewString")

	t.Log("valid")
	{
		err := NewString("foo")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}

	t.Log("already exists")
	{
		err := NewString("foo")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}
}

func TestSetString(t *testing.T) {
	t.Log("SetString")

	t.Log("valid, exists")
	{
		err := SetString("foo", "bar")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
		expect := `"bar"`
		v := stats.Get("foo")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := SetString("bar", "foo")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := `"foo"`
		v := stats.Get("bar")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestNewInt(t *testing.T) {
	t.Log("NewInt")

	t.Log("valid")
	{
		err := NewInt("foo_int")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}

	t.Log("already exists")
	{
		err := NewInt("foo_int")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}
}

func TestSetInt(t *testing.T) {
	t.Log("SetInt")

	t.Log("valid, exists")
	{
		err := SetInt("foo_int", 10)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "10"
		v := stats.Get("foo_int")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := SetInt("bar_int", 10)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := `10`
		v := stats.Get("bar_int")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestIncrementInt(t *testing.T) {
	t.Log("IncrementInt")

	t.Log("valid, exists")
	{
		err := IncrementInt("foo_int")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "11"
		v := stats.Get("foo_int")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := IncrementInt("bar_int2")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "1"
		v := stats.Get("bar_int2")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestDecrementInt(t *testing.T) {
	t.Log("DecrementInt")

	t.Log("valid, exists")
	{
		err := DecrementInt("foo_int")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "10"
		v := stats.Get("foo_int")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := DecrementInt("bar_int3")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "-1"
		v := stats.Get("bar_int3")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestAddInt(t *testing.T) {
	t.Log("AddInt")

	t.Log("valid")
	{
		err := AddInt("foo_intadd", 10)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		AddInt("foo_intadd", 5)

		expect := "15"
		v := stats.Get("foo_intadd")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := AddInt("bar_intadd", -1)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "-1"
		v := stats.Get("bar_intadd")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestNewFloat(t *testing.T) {
	t.Log("NewFloat")

	t.Log("valid")
	{
		err := NewFloat("foo_float")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}

	t.Log("already exists")
	{
		err := NewFloat("foo_float")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}
}

func TestSetFloat(t *testing.T) {
	t.Log("SetFloat")

	t.Log("valid, exists")
	{
		err := SetFloat("foo_float", 1.23)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "1.23"
		v := stats.Get("foo_float")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := SetFloat("bar_float", 1.23)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := `1.23`
		v := stats.Get("bar_float")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestAddFloat(t *testing.T) {
	t.Log("AddFloat")

	t.Log("valid, exists")
	{
		err := AddFloat("foo_float", 3.0)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "4.23"
		v := stats.Get("foo_float")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := AddFloat("bar_float2", 1.23)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := `1.23`
		v := stats.Get("bar_float2")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestNewMap(t *testing.T) {
	t.Log("NewMap")

	t.Log("valid")
	{
		err := NewMap("foo")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}

	t.Log("already exists")
	{
		err := NewMap("foo")
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
	}
}

func TestMapAddInt(t *testing.T) {
	t.Log("MapAddInt")

	t.Log("valid, exists")
	{
		err := MapAddInt("foo", "bar_i", 1)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
		expect := "1"
		v := maps["foo"].Get("bar_i")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := MapAddInt("bar", "foo_i", 2)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "2"
		v := maps["bar"].Get("foo_i")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}

func TestMapAddFloat(t *testing.T) {
	t.Log("MapAddFloat")

	t.Log("valid, exists")
	{
		err := MapAddFloat("foo", "bar_f", 1.23)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}
		expect := "1.23"
		v := maps["foo"].Get("bar_f")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}

	t.Log("valid, does not exist")
	{
		err := MapAddFloat("bar", "foo_f", 1.32)
		if err != nil {
			t.Fatalf("expected no error, got (%s)", err)
		}

		expect := "1.32"
		v := maps["bar"].Get("foo_f")
		if v.String() != expect {
			t.Fatalf("expected (%s), got (%s)", expect, v.String())
		}
	}
}
