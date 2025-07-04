package adapters

import (
	"context"
	"encoding/gob"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {
	gob.Register(&GothAccount{})
	gob.Register(&GothUser{})
	gob.Register(&GothSession{})
	gob.Register(&GothVerificationToken{})
	gob.Register(&GothCsrfToken{})
}

// AccountType represents the type of an account.
type AccountType string

// ErrUnimplemented is returned when a method is not implemented.
var ErrUnimplemented = errors.New("not implemented")

const (
	// AccountTypeOAuth2 represents an OAuth2 account type.
	AccountTypeOAuth2 AccountType = "oauth2"
	// AccountTypeOIDC represents an OIDC account type.
	AccountTypeOIDC AccountType = "oidc"
	// AccountTypeSAML represents a SAML account type.
	AccountTypeSAML AccountType = "saml"
	// AccountTypeEmail represents an email account type.
	AccountTypeEmail AccountType = "email"
	// AccountTypeWebAuthn represents a WebAuthn account type.
	AccountTypeWebAuthn AccountType = "webauthn"
)

// GothAccount represents an account in a third-party identity provider.
type GothAccount struct {
	// ID is the unique identifier of the account.
	ID uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;column:id;default:gen_random_uuid();"`
	// Type is the type of the account.
	Type AccountType `json:"type" validate:"required"`
	// Provider is the provider of the account.
	Provider string `json:"provider" validate:"required"`
	// ProviderAccountID is the account ID in the provider.
	ProviderAccountID *string `json:"provider_account_id"`
	// RefreshToken is the refresh token of the account.
	RefreshToken *string `json:"refresh_token"`
	// AccessToken is the access token of the account.
	AccessToken *string `json:"access_token"`
	// ExpiresAt is the expiry time of the account.
	ExpiresAt *time.Time `json:"expires_at"`
	// TokenType is the token type of the account.
	TokenType *string `json:"token_type"`
	// Scope is the scope of the account.
	Scope *string `json:"scope"`
	// IDToken is the ID token of the account.
	IDToken *string `json:"id_token"`
	// SessionState is the session state of the account.
	SessionState string `json:"session_state"`
	// UserID is the user ID of the account.
	UserID *uuid.UUID `json:"user_id"`
	//  User is the user of the account.
	User GothUser `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	// CreatedAt is the creation time of the account.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the update time of the account.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the deletion time of the account.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// GothUser is a user of the application.
type GothUser struct {
	// ID is the unique identifier of the user.
	ID uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;column:id;default:gen_random_uuid()"`
	// Name is the name of the user.
	Name string `json:"name" validate:"required,max=255"`
	// Email is the email of the user.
	Email string `json:"email" gorm:"unique" validate:"required,email"`
	// EmailVerified is true if the email is verified.
	EmailVerified *bool `json:"email_verified"`
	// Image is the image URL of the user.
	Image *string `json:"image" validate:"url"`
	// Password is the password of the user.
	Accounts []GothAccount `json:"accounts" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	// Sessions are the sessions of the user.
	Sessions []GothSession `json:"sessions" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	// CreatedAt is the creation time of the user.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the update time of the user.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the deletion time of the user.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// GothSession is a session for a user.
type GothSession struct {
	// ID is the unique identifier of the session.
	ID uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;column:id;default:gen_random_uuid()"`
	// SessionToken is the token of the session.
	SessionToken string `json:"session_token"`
	// CsrfToken is the CSRF token of the session.
	CsrfToken GothCsrfToken `json:"csrf_token"`
	// CsrfTokenID is the CSRF token ID of the session.
	CsrfTokenID uuid.UUID `json:"csrf_token_id"`
	// UserID is the user ID of the session.
	UserID uuid.UUID `json:"user_id"`
	// User is the user of the session.
	User GothUser `json:"user"`
	// ExpiresAt is the expiry time of the session.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt is the creation time of the session.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the update time of the session.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the deletion time of the session.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// GetUser returns the user of the session.
func (s *GothSession) GetUser() GothUser {
	return s.User
}

// GothCsrfToken is a CSRF token for a user.
type GothCsrfToken struct {
	// ID is the unique identifier of the CSRF token.
	ID uuid.UUID `json:"id" gorm:"primaryKey;unique;type:uuid;column:id;default:gen_random_uuid()"`
	// Token is the unique identifier of the token.
	Token string `json:"token"`
	// ExpiresAt is the expiry time of the token.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt is the creation time of the token.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the update time of the token.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the deletion time of the token.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// IsValid returns true if the session is valid.
func (s *GothSession) IsValid() bool {
	return s.ExpiresAt.After(time.Now())
}

// GetCsrfToken returns the CSRF token.
func (s *GothSession) GetCsrfToken() GothCsrfToken {
	return s.CsrfToken
}

// HasExpired returns true if the session has expired.
func (c GothCsrfToken) HasExpired() bool {
	return c.ExpiresAt.Before(time.Now())
}

// IsValid returns true if the token is valid.
func (c GothCsrfToken) IsValid(token string) bool {
	return c.Token == token
}

// GothVerificationToken is a verification token for a user.
type GothVerificationToken struct {
	// Token is the unique identifier of the token.
	Token string `json:"token" gorm:"primaryKey"`
	// Identifier is the identifier of the token.
	Identifier string `json:"identifier"`
	// ExpiresAt is the expiry time of the token.
	ExpiresAt time.Time `json:"expires_at"`
	// CreatedAt is the creation time of the token.
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt is the update time of the token.
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt is the deletion time of the token.
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Adapter is an interface that defines the methods for interacting with the underlying data storage.
type Adapter interface {
	// CreateUser creates a new user.
	CreateUser(ctx context.Context, user GothUser) (GothUser, error)
	// GetUser retrieves a user by ID.
	GetUser(ctx context.Context, id uuid.UUID) (GothUser, error)
	// GetUserByEmail retrieves a user by email.
	GetUserByEmail(ctx context.Context, email string) (GothUser, error)
	// UpdateUser updates a user.
	UpdateUser(ctx context.Context, user GothUser) (GothUser, error)
	// DeleteUser deletes a user by ID.
	DeleteUser(ctx context.Context, id uuid.UUID) error
	// LinkAccount links an account to a user.
	LinkAccount(ctx context.Context, accountID, userID uuid.UUID) error
	// UnlinkAccount unlinks an account from a user.
	UnlinkAccount(ctx context.Context, accountID, userID uuid.UUID) error
	// CreateSession creates a new session.
	CreateSession(ctx context.Context, userID uuid.UUID, expires time.Time) (GothSession, error)
	// GetSession retrieves a session by session token.
	GetSession(ctx context.Context, sessionToken string) (GothSession, error)
	// UpdateSession updates a session.
	UpdateSession(ctx context.Context, session GothSession) (GothSession, error)
	// RefreshSession refreshes a session.
	RefreshSession(ctx context.Context, session GothSession) (GothSession, error)
	// DeleteSession deletes a session by session token.
	DeleteSession(ctx context.Context, sessionToken string) error
	// CreateVerificationToken creates a new verification token.
	CreateVerificationToken(ctx context.Context, verficationToken GothVerificationToken) (GothVerificationToken, error)
	// UseVerficationToken uses a verification token.
	UseVerficationToken(ctx context.Context, identifier, token string) (GothVerificationToken, error)
}

var _ Adapter = (*UnimplementedAdapter)(nil)

// UnimplementedAdapter is an adapter that does not implement any of the methods.
type UnimplementedAdapter struct{}

// CreateUser creates a new user.
func (a *UnimplementedAdapter) CreateUser(_ context.Context, _ GothUser) (GothUser, error) {
	return GothUser{}, ErrUnimplemented
}

// GetUser retrieves a user by ID.
func (a *UnimplementedAdapter) GetUser(_ context.Context, _ uuid.UUID) (GothUser, error) {
	return GothUser{}, ErrUnimplemented
}

// GetUserByEmail retrieves a user by email.
func (a *UnimplementedAdapter) GetUserByEmail(_ context.Context, _ string) (GothUser, error) {
	return GothUser{}, ErrUnimplemented
}

// GetUserByAccount retrieves a user by account.
func (a *UnimplementedAdapter) GetUserByAccount(_ context.Context, _, _ string) (GothUser, error) {
	return GothUser{}, ErrUnimplemented
}

// UpdateUser updates a user.
func (a *UnimplementedAdapter) UpdateUser(_ context.Context, _ GothUser) (GothUser, error) {
	return GothUser{}, ErrUnimplemented
}

// DeleteUser deletes a user by ID.
func (a *UnimplementedAdapter) DeleteUser(_ context.Context, _ uuid.UUID) error {
	return ErrUnimplemented
}

// LinkAccount links an account to a user.
func (a *UnimplementedAdapter) LinkAccount(_ context.Context, _, _ uuid.UUID) error {
	return ErrUnimplemented
}

// UnlinkAccount unlinks an account from a user.
func (a *UnimplementedAdapter) UnlinkAccount(_ context.Context, _, _ uuid.UUID) error {
	return ErrUnimplemented
}

// CreateSession creates a new session.
func (a *UnimplementedAdapter) CreateSession(_ context.Context, _ uuid.UUID, _ time.Time) (GothSession, error) {
	return GothSession{}, ErrUnimplemented
}

// GetSession retrieves a session by session token.
func (a *UnimplementedAdapter) GetSession(_ context.Context, _ string) (GothSession, error) {
	return GothSession{}, ErrUnimplemented
}

// UpdateSession updates a session.
func (a *UnimplementedAdapter) UpdateSession(_ context.Context, _ GothSession) (GothSession, error) {
	return GothSession{}, ErrUnimplemented
}

// RefreshSession refreshes a session.
func (a *UnimplementedAdapter) RefreshSession(_ context.Context, _ GothSession) (GothSession, error) {
	return GothSession{}, ErrUnimplemented
}

// DeleteSession deletes a session by session token.
func (a *UnimplementedAdapter) DeleteSession(_ context.Context, _ string) error {
	return ErrUnimplemented
}

// CreateVerificationToken creates a new verification token.
func (a *UnimplementedAdapter) CreateVerificationToken(_ context.Context, _ GothVerificationToken) (GothVerificationToken, error) {
	return GothVerificationToken{}, ErrUnimplemented
}

// UseVerficationToken uses a verification token.
func (a *UnimplementedAdapter) UseVerficationToken(_ context.Context, _, _ string) (GothVerificationToken, error) {
	return GothVerificationToken{}, ErrUnimplemented
}
