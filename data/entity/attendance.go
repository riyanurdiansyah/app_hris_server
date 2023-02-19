package entity

type ClockIn struct {
	ID               int     `gorm:"column:id;primaryKey;autoIncrement"`
	UserId           int     `gorm:"column:id_user"`
	ImageClockin     string  `gorm:"column:image_clockin"`
	TimeClockin      string  `gorm:"column:time_clockin"`
	LatitudeClockin  float64 `gorm:"column:latitude_clockin"`
	LongitudeClockin float64 `gorm:"column:longitude_clockin"`
	NoteClockin      string  `gorm:"column:note_clockin"`
}

type ClockOut struct {
	ID                int     `gorm:"column:id;primaryKey;autoIncrement"`
	UserId            int     `gorm:"column:id_user"`
	ImageClockout     string  `gorm:"column:image_clockout"`
	TimeClockout      string  `gorm:"column:time_clockout"`
	LatitudeClockout  float64 `gorm:"column:latitude_clockout"`
	LongitudeClockout float64 `gorm:"column:Longitude_clockout"`
	NoteClockout      string  `gorm:"column:note_clockout"`
}

type Attendance struct {
	ID                int     `gorm:"column:id;primaryKey;autoIncrement"`
	UserId            int     `gorm:"column:id_user"`
	ImageClockin      string  `gorm:"column:image_clockin"`
	TimeClockin       string  `gorm:"column:time_clockin"`
	LatitudeClockin   float64 `gorm:"column:latitude_clockin"`
	LongitudeClockin  float64 `gorm:"column:longitude_clockin"`
	NoteClockin       string  `gorm:"column:note_clockin"`
	ImageClockout     string  `gorm:"column:image_clockout"`
	TimeClockout      string  `gorm:"column:time_clockout"`
	LatitudeClockout  float64 `gorm:"column:latitude_clockout"`
	LongitudeClockout float64 `gorm:"column:Longitude_clockout"`
	NoteClockout      string  `gorm:"column:note_clockout"`
}
