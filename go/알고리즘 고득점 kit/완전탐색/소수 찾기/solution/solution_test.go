package solution

import (
	"testing"
)

func Test1(t *testing.T) {
	numbers := "17"
	result := solution(numbers)
	expect := 3

	if result != expect {
		t.Errorf("Expect %d, but %d", expect, result)
	}
}

func Test2(t *testing.T) {
	numbers := "011"
	result := solution(numbers)
	expect := 2

	if result != expect {
		t.Errorf("Expect %d, but %d", expect, result)
	}
}

func TestPrime(t *testing.T) {
	if isPrime(1) {
		t.Errorf("1 is not prime number")
	}

	if !isPrime(2) {
		t.Errorf("2 is prime number")
	}

	if !isPrime(3) {
		t.Errorf("3 is prime number")
	}

	if isPrime(4) {
		t.Errorf("4 is not prime number")
	}

	if !isPrime(5) {
		t.Errorf("5 is prime number")
	}

	if isPrime(6) {
		t.Errorf("6 is not prime number")
	}

	if !isPrime(7) {
		t.Errorf("7 is prime number")
	}

	if isPrime(8) {
		t.Errorf("8 is not prime number")
	}

	if isPrime(9) {
		t.Errorf("9 is not prime number")
	}

	if isPrime(10) {
		t.Errorf("10 is not prime number")
	}
}
