package token

import (
	"crypto/rsa"
	"github.com/google/uuid"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

type (
	Scope string
	Token struct {
		accessTokenExp  time.Duration
		refreshTokenExp time.Duration
		leeway          time.Duration
		issuer          string
		jwt             jwt.NestedBuilder
		key             *rsa.PrivateKey
	}
)

type token struct {
	TokenType            string
	AccessToken          string
	AccessTokenExpireIn  time.Time
	RefreshToken         string
	RefreshTokenExpireIn time.Time
}

type Claims struct {
	jwt.Claims
	Scope Scope `json:"scope,omitempty"`
}

const (
	TokenType = "JWT"

	ScopeBasic          Scope = "BASIC"
	ScopeOnboard        Scope = "ONBOARD"
	ScopeAdministration Scope = "ADMINISTRATION"
)

var level = map[Scope]int{
	ScopeBasic:          1,
	ScopeAdministration: 3,
}

func NewToken(key *rsa.PrivateKey, issuer string, accessTokenExp, refreshTokenExp, leeway time.Duration) (*Token, error) {
	signer, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.RS256, Key: key}, (&jose.SignerOptions{}).WithContentType(TokenType))
	if err != nil {
		return nil, err
	}

	enc, err := jose.NewEncrypter(jose.A128GCM, jose.Recipient{Algorithm: jose.RSA_OAEP, Key: &key.PublicKey}, (&jose.EncrypterOptions{}).WithContentType(TokenType))
	if err != nil {
		return nil, err
	}

	return &Token{
		jwt:             jwt.SignedAndEncrypted(signer, enc),
		key:             key,
		accessTokenExp:  accessTokenExp,
		refreshTokenExp: refreshTokenExp,
		leeway:          leeway,
		issuer:          issuer,
	}, nil
}

func (t *Token) GenerateToken(clientID string, sessionID uuid.UUID, scope Scope) (*token, error) {
	iat := time.Now()
	nbf := iat.Add(t.leeway)
	exp := iat.Add(t.accessTokenExp)
	rexp := iat.Add(t.refreshTokenExp)

	access, err := t.encode(Claims{
		Claims: jwt.Claims{
			Issuer:    t.issuer,
			Subject:   sessionID.String(),
			Expiry:    jwt.NewNumericDate(exp),
			NotBefore: jwt.NewNumericDate(nbf),
			IssuedAt:  jwt.NewNumericDate(iat),
			ID:        clientID,
		},
		Scope: scope,
	})

	if err != nil {
		return nil, err
	}

	refresh, err := t.encode(Claims{
		Claims: jwt.Claims{
			Issuer:  t.issuer,
			Subject: sessionID.String(),
			Expiry:  jwt.NewNumericDate(rexp),
			NotBefore: jwt.NewNumericDate(nbf),
		},
		Scope: scope,
	})

	if err != nil {
		return nil, err
	}

	return &token{
		TokenType:            TokenType,
		AccessToken:          access,
		AccessTokenExpireIn:  exp,
		RefreshToken:         refresh,
		RefreshTokenExpireIn: rexp,
	}, nil
}

func (t *Token) encode(c Claims) (string, error) {
	return t.jwt.Claims(c).CompactSerialize()
}

func (t *Token) Decode(token string) (Claims, error) {
	var result Claims

	parsed, err := jwt.ParseSignedAndEncrypted(token)
	if err != nil {
		return result, err
	}

	decrypted, err := parsed.Decrypt(t.key)
	if err != nil {
		return result, err
	}

	if err = decrypted.Claims(&t.key.PublicKey, &result); err != nil {
		return result, err
	}

	return result, err
}

func (t *Token) Validate(c Claims) error {
	e := jwt.Expected{
		Issuer: t.issuer,
		Time:   time.Now(),
	}

	return c.Validate(e)
}

func IsAllowedScope(rule Scope, accessor Scope) bool {
	if rule == accessor {
		return true
	}
	var tier = level[rule]
	return tier != 0 && level[rule] < level[accessor]
}
