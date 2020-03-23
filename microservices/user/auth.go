package user

import (
	"context"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

const (
	tokenHeaderName = "x-testbed-token"
)

func VerifyToken(ctx context.Context, userUUID string, publicKey string) error {
	header, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("Could not get metadata header")
	}

	var tokenStr string
	if tokens, ok := header[tokenHeaderName]; !ok {
		return fmt.Errorf("token header is not found: %#v", header)
	} else if len(tokens) != 1 {
		return fmt.Errorf("token field is not valid: %v", tokens)
	} else {
		tokenStr = tokens[0]
	}
	log.Printf("validate target token is %s (userUUID=%s)", tokenStr, userUUID)

	verifyKey, err := jwt.ParseECPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		return err
	}

	//parts := strings.Split(tokenStr, ".")
	//if err := jwt.SigningMethodES512.Verify(strings.Join(parts[0:2], "."), parts[2], verifyKey); err != nil {
	//	return fmt.Errorf("invalid token: %w", err)
	//}

	type UserClaims struct {
		UserUUID string `json:"user_uuid"`
		jwt.StandardClaims
	}

	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return verifyKey, nil
	})
	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		if claims.UserUUID != userUUID {
			return fmt.Errorf("token is valid, but user uuid is not match (got=%s, exp=%s)", claims.UserUUID, userUUID)
		}
		return nil
	}

	return fmt.Errorf("token is not valid")
}
