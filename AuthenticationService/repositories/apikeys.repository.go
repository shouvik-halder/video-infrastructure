package repositories

import (
	"AuthenticationService/interfaces"
	"AuthenticationService/models"
	"database/sql"
)

type apiKeysRepositoryImpl struct {
	sqldb *sql.DB
}

func NewApiKeysRepositor(db *sql.DB) interfaces.ApiKeysRepository {
	return &apiKeysRepositoryImpl{
		sqldb: db,
	}
}

func (apiRepo *apiKeysRepositoryImpl) Create(userId int64, keyId, keyHash string) (*models.ApiKey, error) {
	query := `INSERT INTO api_keys(user_id, key_id, key_hash) VALUES (?, ?, ?);`
	result, err := apiRepo.sqldb.Exec(query, userId, keyId, keyHash)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	return apiRepo.GetApiKeyById(id)
}

func (apiRepo *apiKeysRepositoryImpl) Get(keyId string, userId int64) (*models.ApiKey, error) {
	query := `SELECT user_id, key_id, key_hash FROM api_keys WHERE key_id= ? AND user_id = ? AND revoked_at IS NULL`
	row := apiRepo.sqldb.QueryRow(query, keyId, userId)
	apiKey := &models.ApiKey{}
	err := row.Scan(&apiKey.UserId, &apiKey.KeyId, &apiKey.KeyHash)
	if err != nil {
		return nil, err
	}

	return apiKey, nil
}

func (apiRepo *apiKeysRepositoryImpl) Revoke(keyId string) (int64, error) {
	query := `UPDATE api_keys SET revoked_at = CURRENT_TIMESTAMP where key_id = ? AND revoked_at IS NULL`
	row, err := apiRepo.sqldb.Exec(query, keyId)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := row.RowsAffected()
	return rowsAffected, nil
}

func (apiRepo *apiKeysRepositoryImpl) GetApiKeyById(id int64) (*models.ApiKey, error) {
	query := `SELECT key_id, key_hash from api_keys where id = ?`
	row := apiRepo.sqldb.QueryRow(query, id)

	apiKey := &models.ApiKey{}

	err := row.Scan(&apiKey.KeyId, &apiKey.KeyHash)
	if err != nil {
		return nil, err
	}

	return apiKey, nil

}
