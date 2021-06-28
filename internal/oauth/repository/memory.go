package repository

import (
	"context"
	"encoding/json"
	"github.com/rianekacahya/boilerplate/domain/entity"
	"github.com/rianekacahya/boilerplate/pkg/goerror"
)

func (re *oauthRepository) CheckSessionExist(ctx context.Context, key string) (bool, error) {
	exist, err := re.dependency.Rdb.Exists(ctx, key).Result()
	if err != nil {
		re.dependency.Logger.Error(err)
		return false, goerror.Wrap(err, goerror.ErrCodeDataRead, "error when get data session")
	}

	if exist == 1 {
		return true, nil
	}

	return false, nil
}

func (re *oauthRepository) GetSession(ctx context.Context, key string) (*entity.Session, error) {
	var session = new(entity.Session)

	result, err := re.dependency.Rdb.Get(ctx, key).Bytes()
	if err != nil {
		re.dependency.Logger.Error(err)
		return nil, goerror.Wrap(err, goerror.ErrCodeDataRead, "error when get data session")
	}

	// unmarshall session
	err = json.Unmarshal(result, session)
	if err != nil {
		re.dependency.Logger.Error(err)
		return nil, goerror.Wrap(err, goerror.ErrCodeFormatting, "error when serialize data session")
	}

	return session, nil
}
