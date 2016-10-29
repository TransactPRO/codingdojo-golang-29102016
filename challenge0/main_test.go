package main

import "testing"

func TestChallenge(t *testing.T) {

	output, err := solve("11 9 1\n14 90 232\n111 15 111")
	if err != nil {
		t.Errorf("got an error %s", err.Error())
	}

	if output != "1 16 21" {
		t.Error(output)
	}

}

func TestInputParams(t *testing.T) {

	_, err := solve("11 zzz")
	if err == nil {
		t.Error("Got no error")
	}
}

func TestComputeResult(t *testing.T) {
	output, err := computeResult([]string{"2", "4", "19"})

	if err != nil {
		t.Errorf("got an error %s", err.Error())
	}
	if output != "9" {
		t.Error(output)
	}
}
