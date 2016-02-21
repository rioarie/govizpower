package models

import (
	"govizpower/datamappers"
	"log"
)

type ApiModel struct {
	BaseModel
}

type KotaContainer []datamappers.KotaMapper

func (a *ApiModel) GetDataKota() KotaContainer {

	statement := `
		SELECT B.NAMAUPJ, A.GPS_L, A.GPS_B, AVG(A.KWH_KVARH), AVG(A.KWH_LWBP), AVG(A.KWH_WBP)
		FROM stmt201512 A
		LEFT JOIN m_upj B ON A.ID_UPJ = B.IDUPJ
        WHERE A.ID_UPJ NOT IN (114, 361)
		GROUP BY ID_UPJ
		ORDER BY B.NAMAUPJ
	`
	data := datamappers.KotaMapper{}
	container := make(KotaContainer, 0, 100)

	rows, err := a.Query(statement)
	if err != nil {
		log.Println("Error Query: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&data.Kota,
			&data.Gps_L,
			&data.Gps_B,
			&data.Kwh_kvarh,
			&data.Kwh_lwbp,
			&data.Kwh_wbp,
		)

		if err != nil {
			log.Println("Error Fetching data:", err)
		}
		container = append(container, data)
	}

	return container
}
