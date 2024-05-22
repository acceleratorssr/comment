package pwd

import (
	"fmt"
	"testing"
)

func TestHashAndSalt(t *testing.T) {
	fmt.Println(HashAndSalt("999"))
}

func TestCheckPasswords(t *testing.T) {
	hash := HashAndSalt("999")
	fmt.Println(CheckPasswords(hash, "999"))
}
