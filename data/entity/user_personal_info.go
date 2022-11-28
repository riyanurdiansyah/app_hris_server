package entity

type UserPersonalInfo struct {
	ID                int    `gorm:"column:id;primaryKey;autoIncrement"`
	IdEmployee        string `gorm:"column:id_employee"`
	IdUser            int    `gorm:"column:id_user"`
	NamaDepan         string `gorm:"column:nama_depan"`
	NamaBelakang      string `gorm:"column:nama_belakang"`
	JenisKelamin      string `gorm:"column:jenis_kelamin"`
	TempatLahir       string `gorm:"column:tempat_lahir"`
	TanggalLahir      string `gorm:"column:tanggal_lahir"`
	NoHP              string `gorm:"column:no_hp"`
	Telepon           string `gorm:"column:telepon"`
	StatusPernikahan  string `gorm:"column:status_pernikahan"`
	Agama             string `gorm:"column:agama"`
	NomorId           string `gorm:"column:nomor_id"`
	TipeId            string `gorm:"column:tipe_id"`
	TanggalKadaluarsa string `gorm:"column:tanggal_kadaluarsa"`
	AlamatKTP         string `gorm:"column:alamat_ktp"`
	AlamatDomisili    string `gorm:"column:alamat_domisili"`
	GolonganDarah     string `gorm:"column:golongan_darah"`
	CreatedAt         string `gorm:"column:created_at"`
	UpdatedAt         string `gorm:"column:updated_at"`
}
