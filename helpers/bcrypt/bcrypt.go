package bcrypt

import "golang.org/x/crypto/bcrypt"

func HashPassword(pw string) string {
	pwBytes := []byte(pw)
	hashedPW, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hashedPW)
}

func ComparePW(unhashedPW []byte, hashedPW []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPW, unhashedPW)
}
