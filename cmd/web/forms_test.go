package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_Form_Has(t *testing.T) {
	form := NewForm(nil)

	has := form.Has("whatever")
	if has {
		t.Error("form show has field when it should not")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")

	form = NewForm(postedData)
	has = form.Has("a")

	if !has {
		t.Error("show form does not has field when it should")
	}
}

func Test_Form_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)
	form := NewForm(r.PostForm)

	form.Required("a", "b", "c")

	if form.Valid() {
		t.Error("form show valid when required field are missing")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "b")
	postedData.Add("c", "c")

	r, _ = http.NewRequest("POST", "/whatevet", nil)
	r.PostForm = postedData

	form = NewForm(r.PostForm)

	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("shows post does not have required field when it does")
	}
}

func Test_Form_Check(t *testing.T) {
	form := NewForm(nil)

	form.Check(false, "password", "password is required")
	if form.Valid() {
		t.Error("Valid() return false, and it should be true when calling Valid()")
	}
}

func Test_Form_ErrorGet(t *testing.T) {
	form := NewForm(nil)
	form.Check(false, "password", "password is required")
	s := form.Errors.Get("password")

	if len(s) == 0 {
		t.Error("should have an error returned from Get, but do not")
	}

	s = form.Errors.Get("whatever")
	if len(s) != 0 {
		t.Error("should not have an error")
	}
}
