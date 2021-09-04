package main

import "testing"

func TestRun(t *testing.T) {
	err := runApplication()

	if err != nil {
		t.Errorf("failed runApplication()")
	}

}
