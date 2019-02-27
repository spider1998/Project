package entity

const TableGenderStatistic = "gender_statistic"

type GenderStatistic struct {
	Times	int 		`json:"times"`
	ArtSum 	int 		`json:"art_sum"`
	Male	int 		`json:"male"`
	Female	int 		`json:"female"`
}


func (GenderStatistic) TableName() string {
	return TableGenderStatistic
}
