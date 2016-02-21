package datamappers

type KotaMapper struct {
	Kota      string  `db:"NAMAUPJ"`
	Gps_L     float64 `db:"GPS_L"`
	Gps_B     float64 `db:"GPS_B"`
	Kwh_kvarh float64
	Kwh_lwbp  float64
	Kwh_wbp   float64
}

type ApiMapper struct {
	Data []KotaMapper
}

type KotaContainer []KotaMapper
