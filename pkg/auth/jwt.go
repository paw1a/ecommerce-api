package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/paw1a/ecommerce-api/internal/config"
	"github.com/paw1a/ecommerce-api/internal/domain"
	"time"
)

type TokenProvider interface {
	CreateToken(admin domain.Admin) (string, error)
}

type Provider struct {
	cfg *config.Config
}

func NewTokenProvider(cfg *config.Config) *Provider {
	return &Provider{cfg: cfg}
}

func (p *Provider) CreateToken(admin domain.Admin) (string, error) {
	expTime := time.Minute * time.Duration(p.cfg.JWT.AccessTokenTime)
	claims := jwt.MapClaims{
		"adminID": admin.ID,
		"exp":     time.Now().Add(expTime).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(p.cfg.JWT.Secret))
}
