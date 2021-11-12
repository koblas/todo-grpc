package user

import "golang.org/x/crypto/bcrypt"

func validatePassword(password string) string {
	if len(password) < 8 {
		return "length_too_short"
	}

	return ""
}

func passwordEncrypt(password string) ([]byte, string) {
	msg := validatePassword(password)
	if msg != "" {
		return nil, msg
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "encryption_failed"
	}

	return pass, ""
}

func tokenEncrypt(password string) ([]byte, string, error) {
	enc, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return enc, password, err
}

func tokenCompare(user *User, token string) bool {
	return bcrypt.CompareHashAndPassword(*user.VerificationToken, []byte(token)) == nil
}

func passwordCompare(user *User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return err == nil
}
