package main

import (
    "testing"
)

func TestMain_stringInArray_true(t *testing.T) {
        str := "foo"
        array := []string{"foo", "bar"}
        result := stringInArray(str, array)
        if result != true{
                t.Fatalf("Expected true, got %s", result)
        }
}

func TestMain_stringInArray_false(t *testing.T) {
    str := "foo"
    array := []string{"bar", "bar"}
    result := stringInArray(str, array)
    if result != false {
            t.Fatalf("Expected false, got %s", result)
    }
}
