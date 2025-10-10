package impl

import (
	"context"
	"fhir-sirs/app/config"
	"fhir-sirs/pkg/api/v1/ckg_tb/models"
	"fhir-sirs/pkg/api/v1/ckg_tb/utils"
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CKGTBRepo struct {
	useCache     bool
	cacheWilayah map[string]models.MasterWilayah
	cacheFaskes  map[string]models.MasterFaskes
}

func NewCKGTBRepo() *CKGTBRepo {
	return &CKGTBRepo{
		useCache:     config.GetConfig().CkgTbUseCache,
		cacheWilayah: map[string]models.MasterWilayah{},
		cacheFaskes:  map[string]models.MasterFaskes{},
	}
}

func (r *CKGTBRepo) GeTbtDataFiltered(ctx context.Context, conn *mongo.Client, tanggal string, halaman int) (*models.DataSkriningTBOutput, error) {
	if err := r._ValidateTanggal(tanggal); err != nil {
		return nil, err
	}

	conf := config.GetConfig()
	filter := bson.M{
		"$and": []bson.M{
			{"tgl_pemeriksaan": tanggal},
			{"$or": []bson.M{
				{"gejala_dan_tanda_batuk": "Ya"},

				// TODO: Variabel baru diaktifkan setelah form CKG diupdate
				// {"gejala_dan_tanda_bb_turun": "Ya"},
				// {"gejala_dan_tanda_demam_hilang_timbul": "Ya"},
				// {"gejala_dan_tanda_berkeringat_malam": "Ya"},
				// {"gejala_dan_tanda_pembesaran_getah_bening": "Ya"},

				// TODO: Diskusi terkait pemanfaatan data di SITB
				// jika melihat logic skrining TBC di SITB kontak dengan pasien TBC tidak langsung serta merta positif terduga
				// sementara di CKG saat kontak dengan pasien TB maka akan diarahkan ke periksa BTA atau TCM
				// {"kontak_pasien_tbc": "Ya"},
			}},
		},
	}

	// collection := conn.Database(config.GetConfig().MongoDatabaseName).Collection(config.GetConfig().CkgTbMongoCollectionTransaction)
	// ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	// defer cancel()
	timeoutTransaction := time.Duration(conf.CkgTbGetDataTimeout) * time.Second
	ctxTransaction, collectionTransaction := utils.GetCollection(ctx, conn, conf.CkgTbMongoCollectionTransaction, timeoutTransaction)

	timeoutMasterWilayah := time.Duration(60) * time.Second
	ctxMasterWilayah, collectionMasterWilayah := utils.GetCollection(ctx, conn, "master_wilayah", timeoutMasterWilayah)

	timeoutMasterFaskes := time.Duration(60) * time.Second
	ctxMasterFaskes, collectionMasterFaskes := utils.GetCollection(ctx, conn, "master_faskes", timeoutMasterFaskes)

	// Hitung total count data sesuai filter
	totalCount, err := collectionTransaction.CountDocuments(ctxTransaction, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to count documents: %w", err)
	}

	fmt.Printf("%d data found\n", totalCount)

	// Inisialisasi slice dengan kapasitas berdasarkan total count
	results := make([]models.DataSkriningTBResult, 0, totalCount)

	// Pagination dengan batch 1000 record
	batchSize := int64(conf.CkgTbGetDataPageSize)
	skip := int64(0)
	minTotal := int64(math.Min(float64(totalCount), float64(batchSize)))
	pageTotal := int64(math.Ceil(float64(minTotal) / float64(batchSize)))

	// paksa halaman tidak melebihi pageTotal
	if halaman > int(pageTotal) && pageTotal > 0 {
		halaman = int(pageTotal)
	}

	skip = int64(halaman-1) * batchSize

	// for skip < totalCount {
	// Hitung berapa banyak record yang akan diambil dalam batch ini
	limit := batchSize
	if skip+batchSize > totalCount {
		limit = totalCount - skip
	}

	// Query dengan pagination
	opts := options.Find().
		SetLimit(limit).
		SetSkip(skip)

	cursor, err := collectionTransaction.Find(ctxTransaction, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find documents at skip %d: %w", skip, err)
	}

	var batchResults []models.DataSkriningTBRaw
	if err := cursor.All(ctxTransaction, &batchResults); err != nil {
		cursor.Close(ctxTransaction)
		return nil, fmt.Errorf("failed to decode batch results at skip %d: %w", skip, err)
	}
	cursor.Close(ctxTransaction)

	// Proses setiap record dalam batch
	for _, raw := range batchResults {
		res := models.DataSkriningTBResult{
			// Identitas Pasien
			PasienID:                 raw.PasienID,
			PasienNIK:                raw.PasienNIK,
			PasienNama:               raw.PasienNama,
			PasienJenisKelamin:       raw.PasienJenisKelamin,
			PasienTglLahir:           raw.PasienTglLahir,
			PasienUsia:               raw.PasienUsia,
			PasienPekerjaan:          raw.PasienPekerjaan,
			PasienProvinsiSatusehat:  raw.PasienProvinsi,
			PasienKabkotaSatusehat:   raw.PasienKabkota,
			PasienKecamatanSatusehat: raw.PasienKecamatan,
			PasienKelurahanSatusehat: raw.PasienKelurahan,
			PasienNoHandphone:        raw.PasienNoHandphone,

			// Data Kunjungan
			KodeFaskesSatusehat: raw.KodeFaskes,
			TglPemeriksaan:      raw.TglPemeriksaan,

			// Data Hasil Pemeriksaan
			BeratBadan:  raw.BeratBadan,
			TinggiBadan: raw.TinggiBadan,
			StatusImt:   raw.StatusImt,
			HasilGds:    raw.HasilGds,
			HasilGdp:    raw.HasilGdp,
			HasilGdpp:   raw.HasilGdpp,

			// Data Faktor Risiko
			KekuranganGizi: raw.KekuranganGizi,
			Merokok:        raw.Merokok,
			PerokokPasif:   raw.PerokokPasif,
			LansiaDiatas65: raw.LansiaDiatas65,
			IbuHamil:       raw.IbuHamil,
			RiwayatDm:      raw.HasilPemeriksaanDm,
			RiwayatHt:      raw.HasilPemeriksaanHt,
			InfeksiHivAids: raw.InfeksiHivAids,

			// Skrining gejala dan tanda
			GejalaBatuk:                raw.GejalaDanTandaBatuk,
			GejalaBbTurun:              raw.GejalaDanTandaBbTurun,
			GejalaDemamHilangTimbul:    raw.GejalaDanTandaDemamHilangTimbul,
			GejalaLesuMalaise:          raw.GejalaDanTandaLesuMalaise,
			GejalaBerkeringatMalam:     raw.GejalaDanTandaBerkeringatMalam,
			GejalaPembesaranKelenjarGB: raw.GejalaDanTandaPembesaranKelenjarGB,
			KontakPasienTbc:            raw.KontakPasienTbc,
			HasilSkriningTbc:           raw.GejalaDanTandaTbc,
		}

		r._HitungHasilSkrining(raw, &res)
		r._MappingMasterData(ctxMasterWilayah, ctxMasterFaskes, collectionMasterWilayah, collectionMasterFaskes, raw, &res)

		// jprint, _ := json.MarshalIndent(res, "", "  ") // "" for no prefix, " " for 2-space indent
		// fmt.Printf(" --> Result %d: ", i+1)
		// fmt.Println(string(jprint))

		results = append(results, res)
	}

	// Update skip untuk batch berikutnya
	// skip += batchSize
	// }

	return &models.DataSkriningTBOutput{
		Count:     totalCount,
		PageTotal: pageTotal,
		PageSize:  batchSize,
		Page:      halaman,
		Results:   results,
	}, nil
}

func (r *CKGTBRepo) PostTbPatientStatus(ctx context.Context, conn *mongo.Client, input []models.StatusPasienTBInput) ([]models.StatusPasienTBResult, error) {
	conf := config.GetConfig()
	results := make([]models.StatusPasienTBResult, 0, len(input))

	timeoutPasienTb := time.Duration(60) * time.Second
	ctxPasienTb, collectionPasienTb := utils.GetCollection(ctx, conn, "pasien_tb", timeoutPasienTb)

	timeoutTransaction := time.Duration(conf.CkgTbGetDataTimeout) * time.Second
	ctxTransaction, collectionTransaction := utils.GetCollection(ctx, conn, conf.CkgTbMongoCollectionTransaction, timeoutTransaction)

	for i, item := range input {
		res := models.StatusPasienTBResult{
			PasienCkgID: item.PasienCkgID,
			TerdugaID:   item.TerdugaID,
			PasienTbID:  item.PasienTbID,
			PasienNIK:   item.PasienNIK,
			IsError:     false,
			Respons:     "",
		}

		// Validas data input
		err := r._ValidateSkriningData(item, i)
		if err != nil {
			res.IsError = true
			res.Respons = err.Error()
			results = append(results, res)
			continue
		}

		// Simpan atau update database.
		resExist, err := utils.FindPasienTb(ctxPasienTb, collectionPasienTb, item)
		if resExist != nil { // sudah ada status
			if utils.IsNotEmptyInt(resExist.PasienCkgID) {
				res.PasienCkgID = resExist.PasienCkgID
			}
			// Ditemukan tapi id CKG belum di set (kemungkinan SITB mendahului lapor)
			if !utils.IsNotEmptyInt(resExist.PasienCkgID) && utils.IsNotEmptyString(item.PasienNIK) {
				var transaction models.DataSkriningTBRaw
				filterTx := bson.D{
					{Key: "nik", Value: item.PasienNIK},
				}

				errTx := collectionTransaction.FindOne(ctxTransaction, filterTx).Decode(&transaction)
				if errTx == nil { // Transaksi layanan CKG ditemukan
					item.PasienCkgID = &transaction.PasienID
					res.PasienCkgID = &transaction.PasienID
				}
			}

			// jaga-jaga kalau pasienTbID tidak dikirim oleh sitb di pengiriman berikutnya
			if utils.IsNotEmptyString(resExist.PasienTbID) && item.PasienTbID == nil {
				item.PasienTbID = resExist.PasienTbID
				res.PasienTbID = resExist.PasienTbID
			}

			if utils.IsNotEmptyString(item.PasienTbID) {
				if utils.IsNotEmptyString(item.StatusDiagnosis) {
					if !utils.IsNotEmptyString(item.DiagnosisLabMetode) {
						item.DiagnosisLabMetode = nil
					}
					if !utils.IsNotEmptyString(item.DiagnosisLabHasil) {
						item.DiagnosisLabHasil = nil
					}
				} else {
					item.StatusDiagnosis = nil
					item.DiagnosisLabMetode = nil
					item.DiagnosisLabHasil = nil
					item.TanggalMulaiPengobatan = nil
					item.TanggalSelesaiPengobatan = nil
					item.HasilAkhir = nil
				}
			} else {
				item.StatusDiagnosis = nil
				item.DiagnosisLabMetode = nil
				item.DiagnosisLabHasil = nil
				item.TanggalMulaiPengobatan = nil
				item.TanggalSelesaiPengobatan = nil
				item.HasilAkhir = nil
			}

			msg, err1 := utils.UpdatePasienTb(ctxPasienTb, collectionPasienTb, item)
			res.Respons = msg
			if err1 != nil {
				res.IsError = true
			}
		} else if err == mongo.ErrNoDocuments { // status baru
			// Coba cari di transaksi
			if utils.IsNotEmptyString(item.PasienNIK) {
				var transaction models.DataSkriningTBRaw
				filterTx := bson.D{
					{Key: "nik", Value: item.PasienNIK},
				}

				errTx := collectionTransaction.FindOne(ctxTransaction, filterTx).Decode(&transaction)
				if errTx == nil { // Transaksi layanan CKG ditemukan
					item.PasienCkgID = &transaction.PasienID
					res.PasienCkgID = &transaction.PasienID
				} else { // SITB duluan dilaporkan oleh CKG
					item.PasienCkgID = nil
					res.PasienCkgID = nil
				}
			}

			if utils.IsNotEmptyString(item.PasienTbID) {
				if utils.IsNotEmptyString(item.StatusDiagnosis) {
					if !utils.IsNotEmptyString(item.DiagnosisLabMetode) {
						item.DiagnosisLabMetode = nil
					}
					if !utils.IsNotEmptyString(item.DiagnosisLabHasil) {
						item.DiagnosisLabHasil = nil
					}
				} else {
					item.StatusDiagnosis = nil
					item.DiagnosisLabMetode = nil
					item.DiagnosisLabHasil = nil
					item.TanggalMulaiPengobatan = nil
					item.TanggalSelesaiPengobatan = nil
					item.HasilAkhir = nil
				}
			} else {
				item.StatusDiagnosis = nil
				item.DiagnosisLabMetode = nil
				item.DiagnosisLabHasil = nil
				item.TanggalMulaiPengobatan = nil
				item.TanggalSelesaiPengobatan = nil
				item.HasilAkhir = nil
			}

			msg, err1 := utils.InsertPasienTb(ctxPasienTb, collectionPasienTb, item)
			res.Respons = msg
			if err1 != nil {
				res.IsError = true
			}
		} else {
			continue
		}

		results = append(results, res)
	}

	return results, nil
}

func (r *CKGTBRepo) _HitungHasilSkrining(raw models.DataSkriningTBRaw, res *models.DataSkriningTBResult) {
	hasilSkrining := "Tidak"

	if raw.PasienUsia < 15 {
		// Gejala batuk dan sudah lebih dari 14 hari
		if raw.GejalaDanTandaBatuk != nil && *raw.GejalaDanTandaBatuk == "Ya" {
			hasilSkrining = "Ya"
		}
		if raw.GejalaDanTandaBbTurun != nil && *raw.GejalaDanTandaBbTurun == "Ya" {
			hasilSkrining = "Ya"
		}
		if raw.GejalaDanTandaDemamHilangTimbul != nil && *raw.GejalaDanTandaDemamHilangTimbul == "Ya" {
			hasilSkrining = "Ya"
		}
		if raw.GejalaDanTandaLesuMalaise != nil && *raw.GejalaDanTandaLesuMalaise == "Ya" {
			hasilSkrining = "Ya"
		}

		// bersihkan gejala untuk dewasa
		res.GejalaBerkeringatMalam = nil
		res.GejalaPembesaranKelenjarGB = nil
	} else { // 15 tahun ke atas
		if raw.InfeksiHivAids != nil && *raw.InfeksiHivAids == "Ya" {
			// Cukup gejala batuk tanpa harus melihat sudah 14 hari atau tidak
			if raw.GejalaDanTandaBatuk != nil && *raw.GejalaDanTandaBatuk == "Ya" {
				hasilSkrining = "Ya"
			}
			if raw.GejalaDanTandaBbTurun != nil && *raw.GejalaDanTandaBbTurun == "Ya" {
				hasilSkrining = "Ya"
			}
			if raw.GejalaDanTandaDemamHilangTimbul != nil && *raw.GejalaDanTandaDemamHilangTimbul == "Ya" {
				hasilSkrining = "Ya"
			}
			if raw.GejalaDanTandaBerkeringatMalam != nil && *raw.GejalaDanTandaBerkeringatMalam == "Ya" {
				hasilSkrining = "Ya"
			}
			if raw.GejalaDanTandaPembesaranKelenjarGB != nil && *raw.GejalaDanTandaPembesaranKelenjarGB == "Ya" {
				hasilSkrining = "Ya"
			}
		} else {
			// Gejala batuk dan sudah lebih dari 14 hari
			if raw.GejalaDanTandaBatuk != nil && *raw.GejalaDanTandaBatuk == "Ya" {
				hasilSkrining = "Ya"
			}
			if raw.GejalaDanTandaBbTurun != nil && *raw.GejalaDanTandaBbTurun == "Ya" {
				hasilSkrining = "Ya"
			}
			if raw.GejalaDanTandaDemamHilangTimbul != nil && *raw.GejalaDanTandaDemamHilangTimbul == "Ya" {
				hasilSkrining = "Ya"
			}
			if raw.GejalaDanTandaBerkeringatMalam != nil && *raw.GejalaDanTandaBerkeringatMalam == "Ya" {
				hasilSkrining = "Ya"
			}
			if raw.GejalaDanTandaPembesaranKelenjarGB != nil && *raw.GejalaDanTandaPembesaranKelenjarGB == "Ya" {
				hasilSkrining = "Ya"
			}
		}

		// bersihkan gejala untuk anak
		res.GejalaLesuMalaise = nil
	}

	res.HasilSkriningTbc = &hasilSkrining
	if hasilSkrining == "Ya" {
		if raw.MetodePemeriksaanTb != nil {
			metodePemeriksaanTB := strings.ToUpper(*raw.MetodePemeriksaanTb)
			switch metodePemeriksaanTB {
			case "TCM":
				if raw.HasilPemeriksaanTbTcm != nil {
					res.MetodePemeriksaanTb = &metodePemeriksaanTB

					//TODO: koordinasikan mapping nilai TCM dengan DE
					// convert hasil TCM ke ["not_detected", "rif_sen", "rif_res", "rif_indet", "invalid", "error", "no_result", "tdl"]
					mapTcm := map[string]string{
						"neg":       "not_detected",
						"rif sen":   "rif_sen",
						"rif res":   "rif_res",
						"rif indet": "rif_indet",
						"invalid":   "invalid",
						"error":     "error",
						"no result": "no_result",
					}
					if utils.IsNotEmptyString(raw.HasilPemeriksaanTbTcm) {
						tcm := strings.ToLower(*raw.HasilPemeriksaanTbTcm)
						if val, ok := mapTcm[tcm]; ok {
							res.HasilPemeriksaanTbTcm = &val
						}
					}
				}
			case "BTA":
				if raw.HasilPemeriksaanTbBta != nil {
					res.MetodePemeriksaanTb = &metodePemeriksaanTB

					//TODO: koordinasikan mapping nilai BTA dengan DE
					// convert hasil BTA ke ["negatif", "positif"]
					var hasilTbBta *string
					if utils.IsNotEmptyString(raw.HasilPemeriksaanTbBta) {
						bta := strings.ToLower(*raw.HasilPemeriksaanTbBta)
						hasilTbBta = &bta
					}
					res.HasilPemeriksaanTbBta = hasilTbBta
				}
			}
		}

		res.TerdugaTb = &hasilSkrining
	}
}

func (r *CKGTBRepo) _MappingMasterData(ctxMasterWilayah context.Context, ctxMasterFaskes context.Context, collectionMasterWilayah *mongo.Collection, collectionMasterFaskes *mongo.Collection, raw models.DataSkriningTBRaw, res *models.DataSkriningTBResult) {
	if utils.IsNotEmptyString(raw.PasienKelurahan) {
		kelurahan, _ := utils.FindMasterWilayah(*raw.PasienKelurahan, ctxMasterWilayah, collectionMasterWilayah, r.useCache, &r.cacheWilayah)
		if kelurahan != nil {
			res.PasienKelurahanSatusehat = raw.PasienKelurahan
			res.PasienKelurahanSitb = kelurahan.KelurahanID
		}
	}

	if utils.IsNotEmptyString(raw.PasienKecamatan) {
		kecamatan, _ := utils.FindMasterWilayah(*raw.PasienKecamatan, ctxMasterWilayah, collectionMasterWilayah, r.useCache, &r.cacheWilayah)
		if kecamatan != nil {
			res.PasienKecamatanSatusehat = raw.PasienKecamatan
			res.PasienKecamatanSitb = kecamatan.KecamatanID
		}
	}

	if utils.IsNotEmptyString(raw.PasienKabkota) {
		kabupaten, _ := utils.FindMasterWilayah(*raw.PasienKabkota, ctxMasterWilayah, collectionMasterWilayah, r.useCache, &r.cacheWilayah)
		if kabupaten != nil {
			res.PasienKabkotaSatusehat = raw.PasienKabkota
			res.PasienKabkotaSitb = kabupaten.KabupatenID
		}
	}

	if utils.IsNotEmptyString(raw.PasienProvinsi) {
		provinsi, _ := utils.FindMasterWilayah(*raw.PasienProvinsi, ctxMasterWilayah, collectionMasterWilayah, r.useCache, &r.cacheWilayah)
		if provinsi != nil {
			res.PasienProvinsiSatusehat = raw.PasienProvinsi
			res.PasienProvinsiSitb = provinsi.ProvinsiID
		}
	}

	if utils.IsNotEmptyString(raw.KodeFaskes) {
		faskes, _ := utils.FindMasterFaskes(*raw.KodeFaskes, ctxMasterFaskes, collectionMasterFaskes, r.useCache, &r.cacheFaskes)
		if faskes != nil {
			res.KodeFaskesSatusehat = raw.KodeFaskes
			res.KodeFaskesSITB = faskes.ID
		}
	}
}

func (r *CKGTBRepo) _ValidateTanggal(tanggal string) error {
	// Parse tanggal dengan format YYYY-MM-DD
	date, err := time.Parse("2006-01-02", tanggal)
	if err != nil {
		return fmt.Errorf("format tanggal harus YYYY-MM-DD")
	}

	// Cek jika tanggal di bawah 2025-01-01
	minDate := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)
	if date.Before(minDate) {
		return fmt.Errorf("tanggal tidak boleh di bawah 2025-01-01")
	}

	// Cek jika tanggal lebih dari hari ini
	today := time.Now().UTC().Truncate(24 * time.Hour)
	if date.After(today) {
		return fmt.Errorf("tanggal tidak boleh lebih dari hari ini")
	}

	return nil
}

func (r *CKGTBRepo) _ValidateSkriningData(item models.StatusPasienTBInput, i int) error {
	// TerdugaID dan PasienNIK tidak boleh kosong
	if item.TerdugaID == nil || *item.TerdugaID == "" {
		return fmt.Errorf("validation error at index %d: terduga_id is required", i)
	}
	if item.PasienNIK == nil || *item.PasienNIK == "" {
		return fmt.Errorf("validation error at index %d: pasien_nik is required", i)
	}

	// 1=TBC SO,
	// 2=TBC RO,
	// 3= Bukan TBC
	statusDiagnosis := []string{"TBC SO", "TBC RO", "Bukan TBC"}

	// Paling tidak StatusTerduga, atau DiagnosisLabHasil harus ada
	if item.PasienTbID != nil && (item.StatusDiagnosis == nil || !slices.Contains(statusDiagnosis, *item.StatusDiagnosis)) {
		return fmt.Errorf("validation error at index %d: at least one of status_terduga, or status_diagnosa must be provided", i)
	} else {
		item.StatusDiagnosis = nil // abaikan
	}

	if utils.IsNotEmptyString(item.StatusDiagnosis) {
		if !utils.IsNotEmptyString(item.DiagnosisLabMetode) {
			return fmt.Errorf("validation error at index %d: diagnosis_lab_metode is required when status_diagnosa is provided", i)
		}
		if !utils.IsNotEmptyString(item.DiagnosisLabHasil) {
			return fmt.Errorf("validation error at index %d: diagnosis_lab_hasil is required when status_diagnosa is provided", i)
		}
	}

	// Hasil Akhir
	// 1= Sembuh,
	// 2= Pengobatan Lengkap,
	// 3= Pengobatan Gagal ,
	// 4= Meninggal,
	// 5= Putus berobat (lost to follow up),
	// 6= Tidak dievaluasi/pindah,
	// 7= Gagal karena Perubahan Diagnosis, "
	statusAkhir := []string{"Sembuh", "Pengobatan Lengkap", "Pengobatan Gagal", "Meninggal", "Putus berobat (lost to follow up)", "Tidak dievaluasi/pindah", "Gagal karena Perubahan Diagnosis"}
	if item.HasilAkhir != nil && !slices.Contains(statusAkhir, *item.HasilAkhir) {
		return fmt.Errorf("validation error at index %d: hasil_akhir must be one of %v", i, statusAkhir)
	}

	return nil
}
