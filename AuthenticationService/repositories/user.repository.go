package repositories

import (
	"AuthenticationService/interfaces"
	"AuthenticationService/models"
	"database/sql"
)

type userRepositoryImpl struct {
	sqlDb *sql.DB
}

func NewUserRepository(db *sql.DB) interfaces.UserRepository {
	return &userRepositoryImpl{
		sqlDb: db,
	}
}

func (userRepo *userRepositoryImpl) Register(email, password string) (*models.User, error) {
	query := `INSERT INTO users(email, password_hash) values (? , ?);`
	result, err := userRepo.sqlDb.Exec(query, email, password)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	return userRepo.GetUserById(id)

}

func (userRepo *userRepositoryImpl) Login(email string) (*models.User, error) {
	query := `SELECT id, email, password_hash FROM users WHERE email = ?`
	row := userRepo.sqlDb.QueryRow(query, email)
	user := &models.User{}

	err := row.Scan(&user.Id, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (userRepo *userRepositoryImpl) GetUserById(id int64) (*models.User, error) {
	query := `SELECT id, email FROM users WHERE id = ?`
	row := userRepo.sqlDb.QueryRow(query, id)
	user := &models.User{}

	err := row.Scan(&user.Id, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
