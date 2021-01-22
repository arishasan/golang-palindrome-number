package main

import "testing"

func TestPalindrome( t *testing.T ){
	
	actual := testingPalindrome("1 20")
	
	if actual != 10 {
		t.Errorf("1 20 failed, expected %v, got %v", "10", actual)
	}else{
		t.Logf("1 20 success, expected %v, got %v", "10", actual)
	}

}

func TestPalindromeSecond( t *testing.T ){
	
	actual := testingPalindrome("99 100")
	
	if actual != 1 {
		t.Errorf("99 100 failed, expected %v, got %v", "1", actual)
	}else{
		t.Logf("99 100 success, expected %v, got %v", "1", actual)
	}

}

func TestPalindromeThird( t *testing.T ){
	
	actual := testingPalindrome("21 31")

	if actual != 1 {
		t.Errorf("21 31 failed, expected %v, got %v", "1", actual)
	}else{
		t.Logf("21 31 success, expected %v, got %v", "1", actual)
	}

}