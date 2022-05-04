package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v7"
	"github.com/paw1a/ecommerce-api/internal/config"
	"github.com/twinj/uuid"
	"time"
)

type RefreshSession struct {
	RefreshToken string        `json:"refreshToken"`
	RefreshExp   int64         `json:"refreshExp"`
	Fingerprint  string        `json:"fingerprint"`
	Claims       jwt.MapClaims `json:"claims"`
}

type AuthDetails struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshInput struct {
	RefreshToken string `json:"-"`
	Fingerprint  string `json:"fingerprint"`
}

type CreateSessionInput struct {
	Fingerprint string
	Claims      jwt.MapClaims
}

type TokenProvider interface {
	CreateJWTSession(input CreateSessionInput) (*AuthDetails, error)
	VerifyToken(tokenString string) (jwt.MapClaims, error)
	Refresh(refreshInput RefreshInput) (*AuthDetails, error)
}

type Provider struct {
	cfg         *config.Config
	redisClient *redis.Client
}

func NewTokenProvider(cfg *config.Config, redisClient *redis.Client) *Provider {
	return &Provider{
		cfg:         cfg,
		redisClient: redisClient,
	}
}

func (p *Provider) CreateJWTSession(input CreateSessionInput) (*AuthDetails, error) {
	accessExpTime := time.Minute * time.Duration(p.cfg.JWT.AccessTokenTime)
	accessExp := time.Now().Add(accessExpTime).Unix()
	input.Claims["exp"] = accessExp

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, input.Claims)
	accessToken, err := unsignedToken.SignedString([]byte(p.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}

	refreshToken := uuid.NewV4().String()
	refreshExpTime := time.Minute * time.Duration(p.cfg.JWT.RefreshTokenTime)
	refreshExp := time.Now().Add(refreshExpTime).Unix()

	session := RefreshSession{
		RefreshToken: refreshToken,
		RefreshExp:   refreshExp,
		Fingerprint:  input.Fingerprint,
		Claims:       input.Claims,
	}

	sessionJson, err := json.Marshal(session)
	if err != nil {
		return nil, err
	}

	p.redisClient.Set(refreshToken, sessionJson, refreshExpTime)

	return &AuthDetails{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (p *Provider) Refresh(refreshInput RefreshInput) (*AuthDetails, error) {
	sessionJson, err := p.redisClient.Get(refreshInput.RefreshToken).Bytes()
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	err = p.redisClient.Del(refreshInput.RefreshToken).Err()
	if err != nil {
		return nil, err
	}

	var session RefreshSession
	err = json.Unmarshal(sessionJson, &session)
	if err != nil {
		return nil, err
	}

	if session.Fingerprint != refreshInput.Fingerprint {
		return nil, errors.New("invalid client fingerprint")
	}

	return p.CreateJWTSession(CreateSessionInput{
		Fingerprint: refreshInput.Fingerprint,
		Claims:      session.Claims,
	})
}

func (p *Provider) VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(p.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token or claims are invalid")
}
