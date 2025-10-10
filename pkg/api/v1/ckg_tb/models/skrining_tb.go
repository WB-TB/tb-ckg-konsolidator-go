package models

type DataSkriningTBRaw struct {
	// Identitas Pasien
	PasienID           int     `bson:"pasien_id"`
	PasienNIK          string  `bson:"nik"`
	PasienNama         string  `bson:"pasien_name"`
	PasienJenisKelamin string  `bson:"jenis_kelamin"`
	PasienTglLahir     string  `bson:"tgl_lahir"`
	PasienUsia         int     `bson:"usia"`
	PasienPekerjaan    string  `bson:"pekerjaan"` //TODO: saat ini data belum tersedia di dwh mongodb
	PasienProvinsi     *string `bson:"provinsi_pasien"`
	PasienKabkota      *string `bson:"kabkota_pasien"`
	PasienKecamatan    *string `bson:"kecamatan_pasien"`
	PasienKelurahan    *string `bson:"kelurahan_pasien"`
	PasienNoHandphone  string  `bson:"no_handphone"`

	// Data Kunjungan
	KodeFaskes     *string `bson:"kode_faskes"`
	NamaFaskes     *string `bson:"nama_faskes"`
	ProvinsiFaskes *string `bson:"provinsi_faskes"`
	KabkotaFaskes  *string `bson:"kabkota_faskes"`
	TglPemeriksaan string  `bson:"tgl_pemeriksaan"`

	// Data Hasil Pemeriksaan
	BeratBadan            *int    `bson:"berat_badan"`
	TinggiBadan           *int    `bson:"tinggi_badan"`
	StatusImt             *string `bson:"imt"`
	KekuranganGizi        *string `bson:"kekurangan_gizi"`
	Merokok               *string `bson:"merokok"`
	PerokokPasif          *string `bson:"perokok_pasif"`
	LansiaDiatas65        *string `bson:"lansia_lebih_dari_65"`
	IbuHamil              *string `bson:"ibu_hamil"`
	HasilGds              *int    `bson:"hasil_gds"`
	HasilGdp              *int    `bson:"hasil_gdp"`
	HasilGdpp             *int    `bson:"hasil_gdpp"`
	PemeriksaanChestXray  *string `bson:"pemeriksaan_chest_xray"`
	MetodePemeriksaanTb   *string `bson:"metode_pemeriksaan_tb"`
	HasilPemeriksaanTbBta *string `bson:"hasil_pemeriksaan_tb_bta"`
	HasilPemeriksaanTbTcm *string `bson:"hasil_pemeriksaan_tb_tcm"`
	HasilPemeriksaanDm    *string `bson:"hasil_pemeriksaan_dm"`
	HasilPemeriksaanHt    *string `bson:"hasil_pemeriksaan_ht"`
	InfeksiHivAids        *string `bson:"inveksi_hiv_aids"`

	// Data Skrining TB
	GejalaDanTandaBatuk                *string `bson:"gejala_dan_tanda_batuk"`
	GejalaDanTandaBbTurun              *string `bson:"gejala_dan_tanda_bb_turun"`                // dipersiapkan untuk variabel baru/tambahan Andak & Dewasa
	GejalaDanTandaDemamHilangTimbul    *string `bson:"gejala_dan_tanda_demam_hilang_timbul"`     // dipersiapkan untuk variabel baru/tambahan Andak & Dewasa
	GejalaDanTandaLesuMalaise          *string `bson:"gejala_dan_tanda_lesu_malaise"`            // dipersiapkan untuk variabel baru/tambahan Anak
	GejalaDanTandaBerkeringatMalam     *string `bson:"gejala_dan_tanda_berkeringat_malam"`       // dipersiapkan untuk variabel baru/tambahan Dewasa
	GejalaDanTandaPembesaranKelenjarGB *string `bson:"gejala_dan_tanda_pembesaran_getah_bening"` // dipersiapkan untuk variabel baru/tambahan Dewasa
	KontakPasienTbc                    *string `bson:"kontak_pasien_tbc"`
	GejalaDanTandaTbc                  *string `bson:"gejala_dan_tanda_tbc"`

	TindakLanjutPenegakanDiagnosa *string `bson:"tindak_lanjut_penegakan_diagnosa"`
}

type DataSkriningTBResult struct {
	// Identitas Pasien
	PasienID                 int     `json:"pasien_ckg_id"`
	PasienNIK                string  `json:"pasien_nik"`
	PasienNama               string  `json:"pasien_nama"`
	PasienJenisKelamin       string  `json:"pasien_jenis_kelamin"`
	PasienTglLahir           string  `json:"pasien_tgl_lahir"`
	PasienUsia               int     `json:"pasien_usia"`
	PasienPekerjaan          string  `bson:"pekerjaan"` //TODO: saat ini data belum tersedia di dwh mongodb
	PasienProvinsiSatusehat  *string `json:"pasien_provinsi_satusehat"`
	PasienKabkotaSatusehat   *string `json:"pasien_kabkota_satusehat"`
	PasienKecamatanSatusehat *string `json:"pasien_kecamatan_satusehat"`
	PasienKelurahanSatusehat *string `json:"pasien_kelurahan_satusehat"`
	PasienProvinsiSitb       *string `json:"pasien_provinsi_sitb"`
	PasienKabkotaSitb        *string `json:"pasien_kabkota_sitb"`
	PasienKecamatanSitb      *string `json:"pasien_kecamatan_sitb"`
	PasienKelurahanSitb      *string `json:"pasien_kelurahan_sitb"`
	PasienNoHandphone        string  `json:"pasien_no_handphone"`

	// Data Kunjungan
	KodeFaskesSatusehat *string `json:"periksa_faskes_satusehat"`
	KodeFaskesSITB      *string `json:"periksa_faskes_sitb"`
	TglPemeriksaan      string  `json:"periksa_tgl"`

	// Data Hasil Pemeriksaan
	BeratBadan  *int    `json:"hasil_berat_badan"`
	TinggiBadan *int    `json:"hasil_tinggi_badan"`
	StatusImt   *string `json:"hasil_imt"`
	HasilGds    *int    `json:"hasil_gds"`
	HasilGdp    *int    `json:"hasil_gdp"`
	HasilGdpp   *int    `json:"hasil_gdpp"`

	// Data Faktor Risiko
	KekuranganGizi *string `json:"risiko_kekurangan_gizi"`
	Merokok        *string `json:"risiko_merokok"`
	PerokokPasif   *string `json:"risiko_perokok_pasif"`
	LansiaDiatas65 *string `json:"risiko_lansia"`
	IbuHamil       *string `json:"risiko_ibu_hamil"`
	RiwayatDm      *string `json:"risiko_dm"`
	RiwayatHt      *string `json:"risiko_hipertensi"`
	InfeksiHivAids *string `json:"risiko_hiv_aids"`

	// Skrining gejala dan tanda
	GejalaBatuk                *string `json:"gejala_batuk"`
	GejalaBbTurun              *string `json:"gejala_bb_turun"`                // dipersiapkan untuk variabel baru/tambahan Andak & Dewasa
	GejalaDemamHilangTimbul    *string `json:"gejala_demam_hilang_timbul"`     // dipersiapkan untuk variabel baru/tambahan Andak & Dewasa
	GejalaLesuMalaise          *string `json:"gejala_lesu_malaise"`            // dipersiapkan untuk variabel baru/tambahan Anak
	GejalaBerkeringatMalam     *string `json:"gejala_berkeringat_malam"`       // dipersiapkan untuk variabel baru/tambahan Dewasa
	GejalaPembesaranKelenjarGB *string `json:"gejala_pembesaran_getah_bening"` // dipersiapkan untuk variabel baru/tambahan Dewasa
	KontakPasienTbc            *string `json:"kontak_pasien_tbc"`
	HasilSkriningTbc           *string `json:"hasil_skrining_tbc"`
	TerdugaTb                  *string `json:"terduga_tb"`

	// Pemeriksaan Lab TB
	MetodePemeriksaanTb   *string `json:"pemeriksaan_tb_metode"`
	HasilPemeriksaanTbBta *string `json:"pemeriksaan_tb_bta"`
	HasilPemeriksaanTbTcm *string `json:"pemeriksaan_tb_tcm"`
}

type DataSkriningTBOutput struct {
	Count     int64                  `json:"totalRecords"`
	PageTotal int64                  `json:"totalPage"`
	PageSize  int64                  `json:"sizePerPage"`
	Page      int                    `json:"currentPage"`
	Results   []DataSkriningTBResult `json:"results"`
}
