package main

import (
	"io/ioutil"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestAdditionHandler(t *testing.T) {
	// Test case 1
	req1 := httptest.NewRequest("GET", "/addition?a=3&b=5", nil)
	w1 := httptest.NewRecorder()
	additionHandler(w1, req1)
	resp1 := w1.Result()
	defer resp1.Body.Close()
	body1, _ := ioutil.ReadAll(resp1.Body)
	result1, _ := strconv.Atoi(string(body1))
	expected1 := 8
	if result1 != expected1 {
		t.Errorf("Test case 1 failed, expected %d but got %d", expected1, result1)
	}

	// // Test case 2
	// req2 := httptest.NewRequest("GET", "/addition?a=-2&b=7", nil)
	// w2 := httptest.NewRecorder()
	// additionHandler(w2, req2)
	// resp2 := w2.Result()
	// defer resp2.Body.Close()
	// body2, _ := ioutil.ReadAll(resp2.Body)
	// result2, _ := strconv.Atoi(string(body2))
	// expected2 := 5
	// if result2 != expected2 {
	// 	t.Errorf("Test case 2 failed, expected %d but got %d", expected2, result2)
	// }

	// // Test case 3
	// req3 := httptest.NewRequest("GET", "/addition?a=0&b=0", nil)
	// w3 := httptest.NewRecorder()
	// additionHandler(w3, req3)
	// resp3 := w3.Result()
	// defer resp3.Body.Close()
	// body3, _ := ioutil.ReadAll(resp3.Body)
	// result3, _ := strconv.Atoi(string(body3))
	// expected3 := 0
	// if result3 != expected3 {
	// 	t.Errorf("Test case 3 failed, expected %d but got %d", expected3, result3)
	// }
}
