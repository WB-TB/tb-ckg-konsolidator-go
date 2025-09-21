package models

type DataTunjanganKhusus struct {
	NamaDrSp                   string `json:"nama_drSp" bson:"nama_drSp"`
	NikDrSp                    string `json:"nik_drSp" bson:"nik_drSp"`
	PractitionerID             string `json:"practitioner_id" bson:"practitioner_id"`
	NamaFasyankes              string `json:"nama_fasyankes" bson:"nama_fasyankes"`
	OrganizationID             string `json:"organization_id" bson:"organization_id"`
	Tanggal                    string `json:"tanggal" bson:"tanggal"`
	JumlahPasien               int    `json:"jumlah_pasien" bson:"jumlah_pasien"`
	JumlahBacaanHasilPenunjang int    `json:"jumlah_bacaan_hasil_penunjang" bson:"jumlah_bacaan_hasil_penunjang"`
	JenisPelayanan             string `json:"jenis_pelayanan" bson:"jenis_pelayanan"`
}
