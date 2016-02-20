package datamappers

type KotaMapper struct {
	Kota      string
	Gps_L     string
	Gps_B     string
	Kwh_kvarh float64
	Kwh_lwbp  float64
	Kwh_wbp   float64
}

type ApiMapper struct {
	Data []KotaMapper
}
