package service

import (
	"app-hris-server/data/dto"
	"app-hris-server/data/entity"
	"app-hris-server/helper"
	"app-hris-server/repository"
	"app-hris-server/validation"
	"errors"
	"log"
	"math/rand"
	"os"
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

// Attendance implements AttendanceService
func (service *AttendanceServiceImpl) Attendance(request *dto.AttendanceCreateDTO) *dto.AttendanceResponseDTO {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	errorValidation := service.Validate.Struct(request)
	if errorValidation != nil {
		msgError := validation.TextValidation(errorValidation.Error())
		return &dto.AttendanceResponseDTO{
			Error:   true,
			Message: msgError,
		}
	} else {
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

				checkPath := "assets/images/attendance"
				if _, err := os.Stat(checkPath); errors.Is(err, os.ErrNotExist) {
					err := os.Mkdir(checkPath, os.ModePerm)
					if err != nil {
						log.Println(err)
					}
				}

				const charsets = "abcdefghijklmnopqrstuvwxyz" +
					"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

				var seededRand *rand.Rand = rand.New(
					rand.NewSource(time.Now().UnixNano()))

				b := make([]byte, 32)
				for i := range b {
					b[i] = charsets[seededRand.Intn(len(charsets))]
				}

				path := checkPath + "/" + string(b) + formatFile
				request.Path = "/" + path
				absent := entity.Attendance{}
				absent.UserId = request.UserId
				absent.Date = request.Date
				if strings.ToLower(request.Kode) == "clockin" {
					absent.TimeClockin = request.Time
					absent.ImageClockin = path
					absent.NoteClockin = request.Note
					absent.LatitudeClockin = request.Latitude
					absent.LongitudeClockin = request.Longitude

					result := service.AttendanceRepository.CheckAttendance(tx, absent.UserId, absent.Date)

					if result.ID == 0 {
						res := service.AttendanceRepository.Attendance(tx, &absent)

						return dto.ToAttendanceResponseDTO(&entity.Attendance{
							ID:                res.ID,
							UserId:            res.UserId,
							ImageClockin:      res.ImageClockin,
							TimeClockin:       res.TimeClockin,
							LatitudeClockin:   res.LatitudeClockin,
							LongitudeClockin:  res.LongitudeClockin,
							NoteClockin:       res.NoteClockin,
							ImageClockout:     res.ImageClockout,
							TimeClockout:      res.TimeClockout,
							LatitudeClockout:  res.LatitudeClockout,
							LongitudeClockout: res.LongitudeClockout,
							NoteClockout:      res.NoteClockout,
						})
					} else {
						absent.TimeClockout = result.TimeClockout
						absent.ImageClockout = result.ImageClockout
						absent.NoteClockout = result.NoteClockout
						absent.LatitudeClockout = result.LatitudeClockout
						absent.LongitudeClockout = result.LongitudeClockout

						res := service.AttendanceRepository.UpdateAttendance(tx, &absent)

						return dto.ToAttendanceResponseDTO(&entity.Attendance{
							ID:                res.ID,
							UserId:            res.UserId,
							ImageClockin:      res.ImageClockin,
							TimeClockin:       res.TimeClockin,
							LatitudeClockin:   res.LatitudeClockin,
							LongitudeClockin:  res.LongitudeClockin,
							NoteClockin:       res.NoteClockin,
							ImageClockout:     res.ImageClockout,
							TimeClockout:      res.TimeClockout,
							LatitudeClockout:  res.LatitudeClockout,
							LongitudeClockout: res.LongitudeClockout,
							NoteClockout:      res.NoteClockout,
						})
					}

				} else {
					absent.TimeClockout = request.Time
					absent.ImageClockout = path
					absent.NoteClockout = request.Note
					absent.LatitudeClockout = request.Latitude
					absent.LongitudeClockout = request.Longitude

					result := service.AttendanceRepository.CheckAttendance(tx, absent.UserId, absent.Date)
					if result.ID == 0 {

						res := service.AttendanceRepository.Attendance(tx, &absent)

						return dto.ToAttendanceResponseDTO(&entity.Attendance{
							ID:                res.ID,
							UserId:            res.UserId,
							ImageClockin:      res.ImageClockin,
							TimeClockin:       res.TimeClockin,
							LatitudeClockin:   res.LatitudeClockin,
							LongitudeClockin:  res.LongitudeClockin,
							NoteClockin:       res.NoteClockin,
							ImageClockout:     res.ImageClockout,
							TimeClockout:      res.TimeClockout,
							LatitudeClockout:  res.LatitudeClockout,
							LongitudeClockout: res.LongitudeClockout,
							NoteClockout:      res.NoteClockout,
						})

					} else {

						absent.TimeClockin = result.TimeClockin
						absent.ImageClockin = result.ImageClockin
						absent.NoteClockin = result.NoteClockin
						absent.LatitudeClockin = result.LatitudeClockin
						absent.LongitudeClockin = result.LongitudeClockin

						res := service.AttendanceRepository.UpdateAttendance(tx, &absent)

						return dto.ToAttendanceResponseDTO(&entity.Attendance{
							ID:                res.ID,
							UserId:            res.UserId,
							ImageClockin:      res.ImageClockin,
							TimeClockin:       res.TimeClockin,
							LatitudeClockin:   res.LatitudeClockin,
							LongitudeClockin:  res.LongitudeClockin,
							NoteClockin:       res.NoteClockin,
							ImageClockout:     res.ImageClockout,
							TimeClockout:      res.TimeClockout,
							LatitudeClockout:  res.LatitudeClockout,
							LongitudeClockout: res.LongitudeClockout,
							NoteClockout:      res.NoteClockout,
						})
					}

				}

			}
		}
	}
}

// CheckAttendance implements AttendanceService
func (service *AttendanceServiceImpl) CheckAttendance(request *dto.AttendanceCreateDTO) bool {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)
	if tx.Error != nil {
		return true
	} else {
		result := service.AttendanceRepository.CheckAttendance(tx, request.UserId, request.Date)
		if strings.ToLower(request.Kode) == "clockin" {
			if result.ID == 0 {
				return false
			} else if result.TimeClockin == "" {
				return false
			} else {
				return true
			}
		} else {
			if result.ID == 0 {
				return false
			} else if result.TimeClockout == "" {
				return false
			} else {
				return true
			}
		}
	}
}
