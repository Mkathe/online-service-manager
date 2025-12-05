package store

import (
	"database/sql"
	"effMobile/internal/model"
	"fmt"

	"github.com/google/uuid"
	"github.com/hashicorp/go-hclog"
)

type StoreService struct {
	db     *sql.DB
	logger hclog.Logger
}

func NewStoreService(db *sql.DB, logger hclog.Logger) *StoreService {
	return &StoreService{db: db, logger: logger}
}

func (s *StoreService) GetServices() ([]model.Service, error) {
	rows, err := s.db.Query("SELECT * FROM services")
	if err != nil {
		s.logger.Error("Failed to get services", "error", err)
		return nil, err
	}
	var services []model.Service
	for rows.Next() {
		var service model.Service
		err = rows.Scan(&service.Id, &service.Name, &service.Price, &service.UserId, &service.StartDate, &service.EndDate)
		if err != nil {
			s.logger.Error("Failed to get services", "error", err)
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (s *StoreService) CreateService(service model.Service) error {
	query := `INSERT INTO services (Id, Name, Price, UserId, StartDate, EndDate)
				VALUES ($1, $2, $3, $4, $5, $6)`

	service.Id = uuid.New()
	_, err := s.db.Exec(query, service.Id, service.Name, service.Price, service.UserId, service.StartDate, service.EndDate)
	if err != nil {
		s.logger.Error("Failed to create service", "error", err)
		return err
	}
	return nil
}

func (s *StoreService) UpdateService(id uuid.UUID, newService model.Service) error {
	queryUpdate := `UPDATE services SET
                    Name = $1,
                    Price = $2,
                    UserId = $3,
                    StartDate = $4,
                    EndDate = $5
                WHERE Id = $6`

	result, err := s.db.Exec(
		queryUpdate,
		newService.Name,
		newService.Price,
		newService.UserId,
		newService.StartDate,
		newService.EndDate,
		id)
	if err != nil {
		s.logger.Error("Failed to update newService", "error", err)
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		s.logger.Error("Failed to update newService", "error", err)
		return err
	}
	s.logger.Info("Updated rows", "info", rows)

	return nil
}

func (s *StoreService) DeleteService(id uuid.UUID) error {
	_, err := s.db.Exec("DELETE FROM services WHERE Id = $1", id)
	if err != nil {
		s.logger.Error("Failed to delete service", "error", err)
		return err
	}

	return nil
}

func (s *StoreService) GetTotalSum(filter model.TotalCostFilter) (int, error) {
	if filter.From.After(filter.To) {
		filter.From, filter.To = filter.To, filter.From
	}

	query := `SELECT SUM(price) 
				FROM services 
                  WHERE StartDate >= $1 AND StartDate <= $2`

	args := []any{filter.From, filter.To}
	idx := 3

	if filter.UserID != uuid.Nil {
		query += fmt.Sprintf(" AND UserId = $%d", idx)
		args = append(args, filter.UserID)
		idx++
	}

	if filter.ServiceName != "" {
		query += fmt.Sprintf(" AND Name = $%d", idx)
		args = append(args, filter.ServiceName)
		idx++
	}

	var total int
	err := s.db.QueryRow(query, args...).Scan(&total)
	if err != nil {
		s.logger.Error("Failed to get total sum", "error", err)
		return -1, err
	}
	return total, nil
}
