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
        INSERT OR REPLACE INTO nco (x, y, city, region, name, category, description, contacts) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`, nco.X, nco.Y, nco.City, nco.Region, nco.Name, nco.Category, nco.Description, nco.Contacts)
	return err
}

func (r *SQLiteRepository) GetByID(ID int64) (*models.NCO, error) {
	var nco models.NCO

	err := r.db.QueryRow(`
        SELECT x, y, city, region, name, category, description, contacts
        FROM nco WHERE id = ?
    `, ID).Scan(&nco.X, &nco.Y, &nco.City, &nco.Region, &nco.Name, &nco.Category, &nco.Description, &nco.Contacts)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("ошибка получения НКО: %v", err)
	}

	return &nco, nil
}

func (r *SQLiteRepository) GetAll() ([]*models.NCOSimple, error) {
	var ncos []*models.NCOSimple

	rows, err := r.db.Query("SELECT id, x, y, name FROM nco")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var nco models.NCOSimple
		err := rows.Scan(&nco.ID, &nco.X, &nco.Y, &nco.Name)
		if err != nil {
			return nil, err
		}
		ncos = append(ncos, &nco)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ncos, nil
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
