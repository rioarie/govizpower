package models

import (
	"govizpower/datamappers"
	"log"
)

type ApiModel struct {
	BaseModel
}

func (a *ApiModel) GetData() datamappers.ApiMapper {

	mapper := datamappers.ApiMapper{}

	statement := `
		SELECT KD_RBM, GPS_L, GPS_B, AVG(KWH_KVARH), AVG(KWH_LWBP), AVG(KWH_WBP)
		FROM stmt201512 
		WHERE KD_RBM <> "" AND KD_RBM <> " "
		GROUP BY KD_RBM
	`

	rows, err := a.Query(statement)
	if err != nil {
		log.Println("Error Query: ", err)
	}

	rows.Next() {
		err := rows.Scan(
			&mapper.Rbm,
			&mapper.GPS_L,
			&mapper.GPS_B,
			&mapper.Kwh_kvarh,
			&mapper.Kwh_lwbp,
			&mapper.Kwh_wbp,
		)
	}

	return mapper
}
