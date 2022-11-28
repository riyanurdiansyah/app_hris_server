package dto

import "app-hris-server/data/entity"

type UserPersonalInfoResponseDTO struct {
	IdEmployee        string `json:"id_employee"`
	IdUser            int    `json:"id_user"`
	NamaDepan         string `json:"nama_depan"`
	NamaBelakang      string `json:"nama_belakang"`
	JenisKelamin      string `json:"jenis_kelamin"`
	TempatLahir       string `json:"tempat_lahir"`
	TanggalLahir      string `json:"tanggal_lahir"`
	NoHP              string `json:"no_hp"`
	Telepon           string `json:"telepon"`
	StatusPernikahan  string `json:"status_pernikahan"`
	Agama             string `json:"agama"`
	NomorId           string `json:"nomor_id"`
	TipeId            string `json:"tipe_id"`
	TanggalKadaluarsa string `json:"tanggal_kadaluarsa"`
	AlamatKTP         string `json:"alamat_ktp"`
	AlamatDomisili    string `json:"alamat_domisili"`
	GolonganDarah     string `json:"golongan_darah"`
	CreatedAt         string `json:"created_at"`
	UpdatedAt         string `json:"updated_at"`
	Error             bool   `json:"-"`
	Message           string `json:"-"`
}

type UserInfoCreateDTO struct {
	IdEmployee        string `validate:"required" json:"id_employee"`
	IdUser            int    `validate:"required" json:"id_user"`
	NamaDepan         string `json:"nama_depan"`
	NamaBelakang      string `json:"nama_belakang"`
	JenisKelamin      string `json:"jenis_kelamin"`
	TempatLahir       string `json:"tempat_lahir"`
	TanggalLahir      string `json:"tanggal_lahir"`
	NoHP              string `json:"no_hp"`
	Telepon           string `json:"telepon"`
	StatusPernikahan  string `json:"status_pernikahan"`
	Agama             string `json:"agama"`
	NomorId           string `json:"nomor_id"`
	TipeId            string `json:"tipe_id"`
	TanggalKadaluarsa string `json:"tanggal_kadaluarsa"`
	AlamatKTP         string `json:"alamat_ktp"`
	AlamatDomisili    string `json:"alamat_domisili"`
	GolonganDarah     string `json:"golongan_darah"`
}

func ToUserPersonalInfoResponseDTO(ent *entity.UserPersonalInfo) *UserPersonalInfoResponseDTO {

	return &UserPersonalInfoResponseDTO{
		IdEmployee:        ent.IdEmployee,
		NamaDepan:         ent.NamaDepan,
		NamaBelakang:      ent.NamaBelakang,
		JenisKelamin:      ent.JenisKelamin,
		TempatLahir:       ent.TempatLahir,
		TanggalLahir:      ent.TanggalLahir,
		NoHP:              ent.NoHP,
		IdUser:            ent.IdUser,
		Telepon:           ent.Telepon,
		StatusPernikahan:  ent.StatusPernikahan,
		Agama:             ent.Agama,
		NomorId:           ent.NomorId,
		TipeId:            ent.TipeId,
		TanggalKadaluarsa: ent.TanggalKadaluarsa,
		AlamatKTP:         ent.AlamatKTP,
		AlamatDomisili:    ent.AlamatDomisili,
		GolonganDarah:     ent.GolonganDarah,
		CreatedAt:         ent.CreatedAt,
		UpdatedAt:         ent.UpdatedAt,
	}
}
