package app

import (
	"effMobile/internal/dto"
	"effMobile/internal/model"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const (
	invalidBodyResponse       = "Invalid body"
	invalidParametersResponse = "Invalid parameters"
)

type ServiceRepository interface {
	GetServices() ([]model.Service, error)
	CreateService(service model.Service) error
	UpdateService(id uuid.UUID, service model.Service) error
	DeleteService(id uuid.UUID) error
	GetTotalSum(filter model.TotalCostFilter) (int, error)
}

func (s *server) HealthCheck(ctx *fiber.Ctx) error {
	err := s.db.Ping()
	if err != nil {
		s.logger.Error("Failed to connect to database", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON("Database error")
	}

	return ctx.Status(fiber.StatusOK).JSON("OK")
}

func (s *server) GetServices(ctx *fiber.Ctx) error {
	services, err := s.storeService.GetServices()
	if err != nil {
		s.logger.Error("Error getting services from repo", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON("Cannot get services")
	}

	return ctx.Status(fiber.StatusOK).JSON(services)
}

func (s *server) CreateService(ctx *fiber.Ctx) error {
	reqBody := ctx.Body()
	var serviceDTO dto.ServiceDTO
	err := json.Unmarshal(reqBody, &serviceDTO)
	if err != nil {
		s.logger.Error("Error unmarshalling body", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(invalidBodyResponse)
	}

	start, err := time.Parse("01-2006", serviceDTO.StartDate)
	if err != nil {
		s.logger.Error("Error parsing end date", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(invalidBodyResponse)
	}

	var end time.Time
	if serviceDTO.EndDate != "" {
		end, err = time.Parse("01-2006", serviceDTO.EndDate)
		if err != nil {
			s.logger.Error("Error parsing end date", "error", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(invalidBodyResponse)
		}
	}

	service := model.Service{
		Name:      serviceDTO.Name,
		Price:     serviceDTO.Price,
		UserId:    serviceDTO.UserId,
		StartDate: start,
		EndDate:   end,
	}

	err = s.storeService.CreateService(service)
	if err != nil {
		s.logger.Error("Error creating service", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON("Cannot create service")
	}

	return ctx.Status(fiber.StatusOK).JSON("Created service")
}

func (s *server) UpdateService(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		s.logger.Error("Error parsing service id", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(invalidParametersResponse)
	}

	reqBody := ctx.Body()
	var serviceDTO dto.ServiceDTO
	err = json.Unmarshal(reqBody, &serviceDTO)
	if err != nil {
		s.logger.Error("Error unmarshalling body", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(invalidBodyResponse)
	}

	start, err := time.Parse("01-2006", serviceDTO.StartDate)
	if err != nil {
		s.logger.Error("Error parsing end date", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(invalidBodyResponse)
	}

	var end time.Time
	if serviceDTO.EndDate != "" {
		end, err = time.Parse("01-2006", serviceDTO.EndDate)
		if err != nil {
			s.logger.Error("Error parsing end date", "error", err)
			return ctx.Status(fiber.StatusInternalServerError).JSON(invalidBodyResponse)
		}
	}

	service := model.Service{
		Name:      serviceDTO.Name,
		Price:     serviceDTO.Price,
		UserId:    serviceDTO.UserId,
		StartDate: start,
		EndDate:   end,
	}

	err = s.storeService.UpdateService(uid, service)
	if err != nil {
		s.logger.Error("Error updating service", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON("Cannot update service")
	}

	return ctx.Status(fiber.StatusOK).JSON("Updated service")
}

func (s *server) DeleteService(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		s.logger.Error("Error parsing uuid", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(invalidParametersResponse)
	}
	err = s.storeService.DeleteService(uid)
	if err != nil {
		s.logger.Error("Error deleting service", "error", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON("Cannot delete service")
	}

	return ctx.Status(fiber.StatusOK).JSON("Deleted service")
}

func (s *server) GetTotalCosts(ctx *fiber.Ctx) error {
	fromStr := ctx.Query("from")
	toStr := ctx.Query("to")

	if fromStr == "" || toStr == "" {
		s.logger.Warn("Error getting total costs", "warning", "from", "to")
		return ctx.Status(fiber.StatusBadRequest).JSON(invalidParametersResponse)
	}

	from, err := time.Parse("01-2006", fromStr)
	if err != nil {
		s.logger.Error("Error parsing 'from', expected MM-YYYY", "error", err)
		return ctx.Status(fiber.StatusBadRequest).JSON("invalid 'from', expected MM-YYYY")
	}

	to, err := time.Parse("01-2006", toStr)
	if err != nil {
		s.logger.Error("Error parsing 'to', expected MM-YYYY", "error", err)
		return ctx.Status(fiber.StatusBadRequest).JSON("invalid 'to', expected MM-YYYY")
	}

	var userID uuid.UUID
	userIdStr := ctx.Query("user_id")
	if userIdStr != "" {
		userID, err = uuid.Parse(userIdStr)
		if err != nil {
			s.logger.Error("Error parsing 'user_id' query param", "error", err)
			return ctx.Status(fiber.StatusBadRequest).JSON(invalidParametersResponse)
		}
	}

	serviceName := ctx.Query("service_name")

	filter := model.TotalCostFilter{
		From:        from,
		To:          to,
		UserID:      userID,
		ServiceName: serviceName,
	}

	total, err := s.storeService.GetTotalSum(filter)
	if err != nil {
		s.logger.Error("Error getting total sum", "error", err)
		return ctx.Status(fiber.StatusBadRequest).JSON("Not found")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"total_cost": total})
}
