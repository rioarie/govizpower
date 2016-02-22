package models

import (
	mappers "govizpower/datamappers"
	"log"
)

type ApiModel struct {
	BaseModel
}

// GetDataKota func queri data power comsumtion average per city
// Param string blth date(YYYYMM)
// Return slice []KotaMapper
func (a *ApiModel) GetDataKota(blth string) mappers.KotaContainer {

	statement := `
		SELECT A.ID_UPJ, B.NAMAUPJ, A.GPS_L, A.GPS_B, AVG(A.KWH_KVARH) AS KWH_KVARH, 
        (CASE WHEN (AVG(A.KWH_LWBP) < 0)
              THEN 0        
              ELSE AVG(A.KWH_LWBP)
              END    
        ) AS KWH_LWBP, 
        AVG(A.KWH_WBP) AS KWH_WBP
		FROM stmt` + blth + ` A
		LEFT JOIN m_upj B ON A.ID_UPJ = B.IDUPJ
		WHERE A.ID_UPJ NOT IN (114, 361) AND A.ID_UPJ < 1000
		AND A.GPS_L <> '' AND A.GPS_B <> '' AND A.GPS_L IS NOT NULL AND A.GPS_B IS NOT NULL
		GROUP BY ID_UPJ
		ORDER BY B.NAMAUPJ
	`
	data := mappers.KotaMapper{}
	container := make(mappers.KotaContainer, 0, 100)

	rows, err := a.Query(statement)
	if err != nil {
		log.Println("Error Query: ", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&data.Id,
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

// GetDetailKota func Query detail of city
// Param string idcity, string blth date(YYYYMM)
// Return slice []DetailMapper
func (a *ApiModel) GetDetailKota(id, blth string) mappers.DetailContainer {
	statement := `
		SELECT IDPELANGGAN, GPS_L, GPS_B, AVG(KWH_KVARH) AS KWH_KVARH, 
		(CASE WHEN (AVG(KWH_LWBP) < 0)
              THEN 0        
              ELSE AVG(KWH_LWBP)
              END    
        ) AS KWH_LWBP, 
        AVG(KWH_WBP) AS KWH_WBP
		FROM stmt` + blth + ` 
		WHERE ID_UPJ = ` + id + `
		AND GPS_L <> '' AND GPS_B <> '' AND GPS_L IS NOT NULL AND GPS_B IS NOT NULL
		GROUP BY IDPELANGGAN
		LIMIT 0,20000
	`

	data := mappers.DetailMapper{}
	container := make(mappers.DetailContainer, 0, 10000)

	rows, err := a.Query(statement)
	if err != nil {
		log.Println("Error Query:", err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&data.Id,
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
