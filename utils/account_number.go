package utils

import (
	"fmt"
)

// LuhnCheckDigit calculates the Luhn check digit for a numeric string
func LuhnCheckDigit(number string) int {
	sum := 0
	double := false

	// Process digits from right to left
	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}

	checkDigit := (10 - (sum % 10)) % 10
	return checkDigit
}

// GenerateAccountNumber generates a bank account number with branch code, sequential ID, and check digit
func GenerateAccountNumber(branchCode string, seq int) string {
	seqStr := fmt.Sprintf("%07d", seq) // zero-pad to 7 digits
	base := branchCode + seqStr
	checkDigit := LuhnCheckDigit(base)
	return fmt.Sprintf("%s%d", base, checkDigit)
}
