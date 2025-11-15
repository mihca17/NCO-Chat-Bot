package repository

import (
	"NCO-Chat-Bot/config"
	"NCO-Chat-Bot/models"
	"database/sql"
	"fmt"
	"sync"
)

// SQLiteRepository реализует интерфейс Repository для SQLite
type SQLiteRepository struct {
	config *config.Config
	db     *sql.DB
	mu     sync.RWMutex
}

// NewSQLiteRepository создает новый репозиторий для SQLite
func NewSQLiteRepository(db *sql.DB, config *config.Config) *SQLiteRepository {
	return &SQLiteRepository{
		config: config,
		db:     db,
	}
}

// База данных функции
func (r *SQLiteRepository) SaveNCO(nco *models.NCO) error {
	_, err := r.db.Exec(`
        INSERT OR REPLACE INTO nco (x, y, name, owner)
        VALUES (?, ?, ?, ?)
    `, nco.X, nco.Y, nco.Name, nco.Owner)
	return err
}

func (r *SQLiteRepository) GetByID(ID int64) (*models.NCO, error) {
	var nco models.NCO

	err := r.db.QueryRow(`
        SELECT nco.X, nco.Y, nco.Name, nco.Owner
        FROM nco WHERE user_id = ?
    `, ID).Scan(&nco.X, &nco.Y, &nco.Name, &nco.Owner)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("ошибка получения НКО: %v", err)
	}

	return &nco, nil
}

func (r *SQLiteRepository) DeleteByID(ID int64) error {
	result, err := r.db.Exec("DELETE FROM nco WHERE id = ?", ID)
	if err != nil {
		return fmt.Errorf("ошибка удаления НКО: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("ошибка получения количества удаленных строк: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("НКО с ID %d не найдено", ID)
	}

	return nil
}
