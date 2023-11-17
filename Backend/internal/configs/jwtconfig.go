package configs

import "os"

func JwtPubKey() ([]byte, error) {
	pubKey, err := os.ReadFile("jwtRS256/jwtRS256.key.pub")
	if err != nil {
		return nil, err
	}
	return pubKey, nil
}

func JwtPrvKey() ([]byte, error) {
	prvKey, err := os.ReadFile("jwtRS256/jwtRS256.key")
	if err != nil {
		return nil, err
	}
	return prvKey, err
}
