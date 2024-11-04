package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"zatrasz75/tz_market/internal/models"
	"zatrasz75/tz_market/pkg/postgres"
)

type Store struct {
	*postgres.Postgres
}

func New(pg *postgres.Postgres) *Store {
	return &Store{pg}
}

// SaveBuilding сохраняет объект Building в базе данных
func (s *Store) SaveBuilding(ctx context.Context, b models.Building) error {
	// Начать транзакцию
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("не удалось запустить транзакцию: %w", err)
	}
	defer tx.Rollback(ctx)

	query := "INSERT INTO Building (name, city, year_built, floors) VALUES ($1, $2, $3, $4)"
	_, err = tx.Exec(ctx, query, b.Name, b.City, b.YearBuilt, b.Floors)
	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса INSERT: %w", err)
	}

	// Фиксация транзакции
	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("не удалось зафиксировать транзакцию: %w", err)
	}

	return nil
}

// CheckBuilding проверяет, есть ли такая запись с именем и городом
func (s *Store) CheckBuilding(ctx context.Context, b models.Building) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM Building WHERE name = $1)`

	// Выполняем запрос и проверяем, существует ли запись
	err := s.Pool.QueryRow(ctx, query, b.Name).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("ошибка при проверке существования записи: %w", err)
	}

	return exists, nil
}

// GetBuildings возвращает список строений, с возможностью фильтрации по городу, году и кол-ву этажей (параметры не обязательные)
func (s *Store) GetBuildings(ctx context.Context, city string, year int, floors int) ([]models.Building, error) {
	var query strings.Builder
	query.Reset()
	query.WriteString("SELECT id, name, city, year_built, floors FROM building WHERE")

	var ownerArgs []interface{}
	var ownerArgIndex int

	// Добавляем фильтры, если параметры не пустые
	if city != "" {
		query.WriteString(" city = $" + strconv.Itoa(ownerArgIndex+1))
		ownerArgs = append(ownerArgs, city)
		ownerArgIndex++
	}
	if year != 0 {
		query.WriteString(" AND year_built = $" + strconv.Itoa(ownerArgIndex+1))
		ownerArgs = append(ownerArgs, year)
		ownerArgIndex++
	}
	if floors != 0 {
		query.WriteString(" AND floors = $" + strconv.Itoa(ownerArgIndex+1))
		ownerArgs = append(ownerArgs, floors)
		ownerArgIndex++
	}

	// Выполняем запрос
	rows, err := s.Pool.Query(ctx, query.String(), ownerArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Сканируем результаты в срез Building
	var buildings []models.Building
	for rows.Next() {
		var building models.Building
		err = rows.Scan(&building.ID, &building.Name, &building.City, &building.YearBuilt, &building.Floors)
		if err != nil {
			return nil, err
		}
		buildings = append(buildings, building)
	}

	return buildings, rows.Err()
}
