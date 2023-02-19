package service

import (
	"app-hris-server/data/dto"
	"app-hris-server/data/entity"
	"app-hris-server/helper"
	"app-hris-server/repository"
	"app-hris-server/validation"
	"errors"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type AttendanceServiceImpl struct {
	AttendanceRepository repository.AttendanceRepository

	DB       *gorm.DB
	Validate *validator.Validate
}

func NewAttendanceService(repository repository.AttendanceRepository, db *gorm.DB, validate *validator.Validate) AttendanceService {
	return &AttendanceServiceImpl{
		AttendanceRepository: repository,
		DB:                   db,
		Validate:             validate,
	}
}

// CheckIfDoneClockin implements AttendanceService
func (*AttendanceServiceImpl) CheckIfDoneClockin(request *dto.ClockinCreateDTO) bool {
	panic("unimplemented")
}

// CheckIfDoneClockout implements AttendanceService
func (*AttendanceServiceImpl) CheckIfDoneClockout(request *dto.ClockoutCreateDTO) bool {
	panic("unimplemented")
}

// Clockin implements AttendanceService
func (service *AttendanceServiceImpl) Clockin(request *dto.ClockinCreateDTO) *dto.AttendanceResponseDTO {

	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.AttendanceResponseDTO{
			Error:   true,
			Message: msgError,
		}
	}
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		msgError := validation.TextValidation(tx.Error.Error())
		return &dto.AttendanceResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
		var formatFile string
		if strings.Contains(request.Image.Filename, "jpg") {
			formatFile = ".jpg"
		} else if strings.Contains(request.Image.Filename, "jpeg") {
			formatFile = ".jpeg"
		} else if strings.Contains(request.Image.Filename, "png") {
			formatFile = ".png"
		} else {
			formatFile = ""
		}

		if formatFile == "" {
			return &dto.AttendanceResponseDTO{
				Error:   true,
				Message: "format file must .jpg/.jpeg/.png",
			}
		} else {

			checkPath := "assets/images/menu"
			if _, err := os.Stat(checkPath); errors.Is(err, os.ErrNotExist) {
				err := os.Mkdir(checkPath, os.ModePerm)
				if err != nil {
					log.Println(err)
				}
			}

			path := checkPath + "/" + strings.ToLower(strings.ReplaceAll(strconv.Itoa(request.UserId), " ", "_")+"_"+request.TimeClockin) + formatFile
			request.Path = "/" + path
			absent := entity.ClockIn{
				UserId:           request.UserId,
				ImageClockin:     request.Path,
				TimeClockin:      time.Now().Format(time.RFC3339),
				LatitudeClockin:  request.LatitudeClockin,
				LongitudeClockin: request.LongitudeClockin,
				NoteClockin:      request.NoteClockin,
			}

			res := service.AttendanceRepository.Clockin(tx, &absent)

			return dto.ToAttendanceResponseDTO(&entity.Attendance{
				ID:               res.ID,
				UserId:           res.UserId,
				ImageClockin:     res.ImageClockin,
				TimeClockin:      res.TimeClockin,
				LatitudeClockin:  res.LatitudeClockin,
				LongitudeClockin: res.LongitudeClockin,
				NoteClockin:      res.NoteClockin,
			})
		}
	}
}

// Clockout implements AttendanceService
func (*AttendanceServiceImpl) Clockout(request *dto.ClockoutCreateDTO) *dto.AttendanceResponseDTO {
	panic("unimplemented")
}
