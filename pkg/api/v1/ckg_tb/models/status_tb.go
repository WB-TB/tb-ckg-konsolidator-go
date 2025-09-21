package models

type StatusPasienTBInput struct {
	PasienCkgID *int    `json:"pasien_ckg_id" bson:"pasien_ckg_id"`
	TerdugaID   *string `json:"terduga_id" bson:"terduga_id"`     // ID Terduga TB
	PasienTbID  *string `json:"pasien_tb_id" bson:"pasien_tb_id"` // ID Pasien TB SO/RO jika sudah terkonfirmasi positif atau dirawat
	PasienNIK   *string `json:"pasien_nik" bson:"pasien_nik"`

	// Parameter dikendalikan dari SITB
	StatusDiagnosis          *string `json:"status_diagnosa" bson:"status_diagnosa"`         // ["TBC SO", "TBC RO", "Bukan TBC"]
	DiagnosisLabMetode       *string `json:"diagnosa_lab_metode" bson:"diagnosa_lab_metode"` // ["TCM", "BTA"]
	DiagnosisLabHasil        *string `json:"diagnosa_lab_hasil" bson:"diagnosa_lab_hasil"`   // TCM: ["not_detected", "rif_sen", "rif_res", "rif_indet", "invalid", "error", "no_result", "tdl"], BTA: ["negatif", "positif"]
	TanggalMulaiPengobatan   *string `json:"tanggal_mulai_pengobatan" bson:"tanggal_mulai_pengobatan"`
	TanggalSelesaiPengobatan *string `json:"tanggal_selesai_pengobatan" bson:"tanggal_selesai_pengobatan"`
	HasilAkhir               *string `json:"hasil_akhir" bson:"hasil_akhir"` // ["Sembuh", "Pengobatan Lengkap", "Pengobatan Gagal", "Meninggal", "Putus berobat (lost to follow up)", "Tidak dievaluasi/pindah", "Gagal karena Perubahan Diagnosis"]
}

type StatusPasienTBResult struct {
	PasienCkgID *int    `json:"pasien_ckg_id"` // ID CKG
	TerdugaID   *string `json:"terduga_id"`    // ID Terduga TB
	PasienTbID  *string `json:"pasien_tb_id"`  // ID Pasien TB SO/RO jika sudah dirawat
	PasienNIK   *string `json:"pasien_nik"`    // NIK
	IsError     bool    `json:"error"`         // menandakan apakah error atau tidak
	Respons     string  `json:"message"`       // pesan respon pemrosesan
}
