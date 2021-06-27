package repository

import (
	"context"
	"database/sql"
	"github.com/rianekacahya/boilerplate/domain/entity"
	"github.com/rianekacahya/boilerplate/pkg/goerror"
	"strings"
)

func (re *oauthRepository) GetClientByClientID(ctx context.Context, clientID string) (*entity.Clients, error) {
	var (
		query  strings.Builder
		result = new(entity.Clients)
	)

	// select
	query.WriteString(`SELECT id, client_id, client_secret, channel, status, created_at, created_by, updated_at, updated_by `)

	// from
	query.WriteString(`FROM clients `)

	// where
	query.WriteString(`WHERE client_id = $1`)

	// execution query
	row := re.dependency.Dbr.QueryRowContext(ctx, query.String(), clientID)
	err := row.Scan(
		&result.ID, &result.ClientID, &result.ClientSecret,
		&result.Channel, &result.Status,
		&result.CreatedAt, &result.CreatedBy,
		&result.UpdatedAt, &result.UpdatedBy,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, goerror.Wrap(err, goerror.ErrCodeNoResult, "data client not found")
		}

		return nil, goerror.Wrap(err, goerror.ErrCodeDataRead, "error when get data client")
	}

	return result, nil
}
