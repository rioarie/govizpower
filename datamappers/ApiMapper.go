package datamappers

type KotaMapper struct {
	Id        int     `db:"ID_UPJ" json:"id"`
	Kota      string  `db:"NAMAUPJ" json:"kota"`
	Gps_L     float64 `db:"GPS_L" json:"gps_l"`
	Gps_B     float64 `db:"GPS_B" json:"gps_b"`
	Kwh_kvarh float64 `json:"kwh_kvarh"`
	Kwh_lwbp  float64 `json:"kwh_lwbp"`
	Kwh_wbp   float64 `json:"kwh_wbp"`
}

type DetailMapper struct {
	Id        int     `db:"IDPELANGGAN" json:"id"`
	Gps_L     float64 `db:"GPS_L" json:"gps_l"`
	Gps_B     float64 `db:"GPS_B" json:"gps_b"`
	Kwh_kvarh float64 `json:"kwh_kvarh"`
	Kwh_lwbp  float64 `json:"kwh_lwbp"`
	Kwh_wbp   float64 `json:"kwh_wbp"`
}

type KotaContainer []KotaMapper
type DetailContainer []DetailMapper
