package models

import (
	"govizpower/datamappers"
	"log"
)

type ApiModel struct {
	BaseModel
}

func (a *ApiModel) GetDataKota() datamappers.KotaMapper {

	mapper := datamappers.KotaMapper{}
	//data := datamappers.ApiMapper{}

	statement := `
		SELECT B.NAMAUPJ, A.GPS_L, A.GPS_B, AVG(A.KWH_KVARH), AVG(A.KWH_LWBP), AVG(A.KWH_WBP)
		FROM stmt201512 A
		LEFT JOIN m_upj B ON A.ID_UPJ = B.IDUPJ
        WHERE A.ID_UPJ NOT IN (114, 361)
		GROUP BY ID_UPJ
		ORDER BY B.NAMAUPJ
	`
	//a.Query(&mapper, statement)

	rows, err := a.Query(statement)
	if err != nil {
		log.Println("Error Query: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&mapper.Kota,
			&mapper.Gps_L,
			&mapper.Gps_B,
			&mapper.Kwh_kvarh,
			&mapper.Kwh_lwbp,
			&mapper.Kwh_wbp,
		)

		if err != nil {
			log.Println("Error Fetching data:", err)
		}
		log.Println(mapper)
	}

	return mapper
}
