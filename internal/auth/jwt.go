package auth

import (
	"Adeeb_Go/internal/config"
	"Adeeb_Go/internal/database/sqlc"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTUserClaims struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Roles []sqlc.Role `json:"roles"`
}

func CreateJWT(ttl time.Duration, user JWTUserClaims, permissions []string) (string, error) {
	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["user"] = user               // Our custom data.
	claims["exp"] = now.Add(ttl).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()          // The time at which the token was issued.
	claims["permissions"] = permissions
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(config.JWT_PRIVATE)
	if err != nil {
		return "", fmt.Errorf("create token error: %w", err)
	}

	return token, nil
}

func NewPermission(signedFor []sqlc.Role) []string {
	var permission []string
	for _, service := range signedFor {
		switch service {
		case sqlc.RoleManagement:
			permission = append(permission, fmt.Sprintf("%v:read", sqlc.RoleManagement), fmt.Sprintf("%v:write", sqlc.RoleManagement))
		case sqlc.RoleDBA:
			permission = append(permission, fmt.Sprintf("%v:read", sqlc.RoleDBA), fmt.Sprintf("%v:write", sqlc.RoleDBA))
		case sqlc.RoleAnalytics:
			permission = append(permission, fmt.Sprintf("%v:read", sqlc.RoleAnalytics), fmt.Sprintf("%v:write", sqlc.RoleAnalytics))
		default:
			return []string{}
		}
	}
	return permission
}

func ValidateToken(authHeader string, onlyAuthorizedFor []string) (jwt.MapClaims, error) {
	if authHeader == "" {
		return nil, fmt.Errorf("Empty header")
	}

	token := authHeader[7:]

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return config.JWT_PUBLIC, nil
	})

	if err != nil {
		return nil, fmt.Errorf("parse error: %v", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	} else if exp := claims["exp"].(float64); int64(exp) < time.Now().Unix() { // you can check claims type in claims.StandardClaims
		return nil, fmt.Errorf("error: token expired")
	}

	// permissions in not interface{}, it's []interface{}. So we need get it, then loop on it and add it item by item.
	permissionsInterface := claims["permissions"].([]interface{})
	permissions := make([]string, len(permissionsInterface))
	for i, v := range permissionsInterface {
		permissions[i] = v.(string)
	}
	if len(onlyAuthorizedFor) > 0 {
		isAuthorized := isAuthorized(onlyAuthorizedFor, permissions)
		if !isAuthorized {
			return nil, fmt.Errorf("not authorized")
		}
	}

	return claims, nil
}
