package repositories

import (
	"context"
	"filterisasi/collection"
	"filterisasi/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitData(ctx context.Context, database *mongo.Database, optionTypes map[string][]*models.PpdbOption, schoolOption []models.PpdbOption) map[string][]*models.PpdbOption {
	for _, opt := range schoolOption {

		fmt.Printf(opt.Id.String())

		var studentRegistrations []models.PpdbRegistration
		studentRegistrations = collection.GetRegistrations(ctx, database, "sma", opt.Id)
		studentHistories := make([]models.PpdbRegistration, len(studentRegistrations), cap(studentRegistrations))
		copy(studentHistories, studentRegistrations)

		tmpOpt := &models.PpdbOption{
			Id:                  opt.Id,
			Name:                opt.Name,
			Quota:               opt.Quota,
			Type:                opt.Type,
			AddQuota:            0,
			Filtered:            0,
			UpdateQuota:         true,
			NeedQuotaFirstOpt:   0,
			PpdbRegistration:    studentRegistrations,
			RegistrationHistory: studentHistories,
			HistoryShifting:     make([]models.PpdbRegistration, 0),
		}
		/*
			tmpOptAbk := &models.PpdbOption{
				Id:                  opt.Id,
				Name:                opt.Name,
				Quota:               opt.Quota,
				Type:                opt.Type,
				AddQuota:            0,
				Filtered:            0,
				UpdateQuota:         true,
				NeedQuotaFirstOpt:   0,
				PpdbRegistration:    studentRegistrations,
				RegistrationHistory: studentHistories,
				HistoryShifting:     make([]models.PpdbRegistration, 0),
			}
			tmpOptKetm := &models.PpdbOption{
				Id:                  opt.Id,
				Name:                opt.Name,
				Quota:               opt.Quota,
				Type:                opt.Type,
				AddQuota:            0,
				Filtered:            0,
				UpdateQuota:         true,
				NeedQuotaFirstOpt:   0,
				PpdbRegistration:    studentRegistrations,
				RegistrationHistory: studentHistories,
				HistoryShifting:     make([]models.PpdbRegistration, 0),
			}
			tmpOptKondisiTertentu := &models.PpdbOption{
				Id:                  opt.Id,
				Name:                opt.Name,
				Quota:               opt.Quota,
				Type:                opt.Type,
				AddQuota:            0,
				Filtered:            0,
				UpdateQuota:         true,
				NeedQuotaFirstOpt:   0,
				PpdbRegistration:    studentRegistrations,
				RegistrationHistory: studentHistories,
				HistoryShifting:     make([]models.PpdbRegistration, 0),
			}
			tmpOptPerpindahan := &models.PpdbOption{
				Id:                  opt.Id,
				Name:                opt.Name,
				Quota:               opt.Quota,
				Type:                opt.Type,
				AddQuota:            0,
				Filtered:            0,
				UpdateQuota:         true,
				NeedQuotaFirstOpt:   0,
				PpdbRegistration:    studentRegistrations,
				RegistrationHistory: studentHistories,
				HistoryShifting:     make([]models.PpdbRegistration, 0),
			}
			tmpOptAnakGuru := &models.PpdbOption{
				Id:                  opt.Id,
				Name:                opt.Name,
				Quota:               opt.Quota,
				Type:                opt.Type,
				AddQuota:            0,
				Filtered:            0,
				UpdateQuota:         true,
				NeedQuotaFirstOpt:   0,
				PpdbRegistration:    studentRegistrations,
				RegistrationHistory: studentHistories,
				HistoryShifting:     make([]models.PpdbRegistration, 0),
			}
			tmpOptRaporUmum := &models.PpdbOption{
				Id:                  opt.Id,
				Name:                opt.Name,
				Quota:               opt.Quota,
				Type:                opt.Type,
				AddQuota:            0,
				Filtered:            0,
				UpdateQuota:         true,
				NeedQuotaFirstOpt:   0,
				PpdbRegistration:    studentRegistrations,
				RegistrationHistory: studentHistories,
				HistoryShifting:     make([]models.PpdbRegistration, 0),
			}
			tmpOptPerlombaan := &models.PpdbOption{
				Id:                  opt.Id,
				Name:                opt.Name,
				Quota:               opt.Quota,
				Type:                opt.Type,
				AddQuota:            0,
				Filtered:            0,
				UpdateQuota:         true,
				NeedQuotaFirstOpt:   0,
				PpdbRegistration:    studentRegistrations,
				RegistrationHistory: studentHistories,
				HistoryShifting:     make([]models.PpdbRegistration, 0),
			}
			tmpOptZonasi := &models.PpdbOption{
				Id:                  opt.Id,
				Name:                opt.Name,
				Quota:               opt.Quota,
				Type:                opt.Type,
				AddQuota:            0,
				Filtered:            0,
				UpdateQuota:         true,
				NeedQuotaFirstOpt:   0,
				PpdbRegistration:    studentRegistrations,
				RegistrationHistory: studentHistories,
				HistoryShifting:     make([]models.PpdbRegistration, 0),
			}*/

		switch opt.Type {
		case "abk":
			optionTypes["abk"] = append(optionTypes["abk"], tmpOpt)
			break
		case "ketm":
			optionTypes["ketm"] = append(optionTypes["ketm"], tmpOpt)
			break
		case "kondisi-tertentu":
			optionTypes["kondisi-tertentu"] = append(optionTypes["kondisi-tertentu"], tmpOpt)
			break
		case "perpindahan":
			optionTypes["perpindahan"] = append(optionTypes["perpindahan"], tmpOpt)
			break
		case "anak-guru":
			optionTypes["anak-guru"] = append(optionTypes["anak-guru"], tmpOpt)
			break
		case "rapor-umum":
			optionTypes["rapor-umum"] = append(optionTypes["rapor-umum"], tmpOpt)
			break
		case "perlombaan":
			optionTypes["perlombaan"] = append(optionTypes["perlombaan"], tmpOpt)
			break
		case "zonasi":
			optionTypes["zonasi"] = append(optionTypes["zonasi"], tmpOpt)
			break
		}

	}

	TmpIdSMAAbk, _ := primitive.ObjectIDFromHex("000000000000000000000011")
	TmpIdSMAKetm, _ := primitive.ObjectIDFromHex("000000000000000000000012")
	TmpIdSMAKondisiTertentu, _ := primitive.ObjectIDFromHex("000000000000000000000013")
	TmpIdSMAPerpindahan, _ := primitive.ObjectIDFromHex("000000000000000000000014")
	TmpIdSMAAnakGuru, _ := primitive.ObjectIDFromHex("000000000000000000000015")
	TmpIdSMAPrestasiNilaiRapor, _ := primitive.ObjectIDFromHex("000000000000000000000016")
	TmpIdSMAPrestasiPerlombaan, _ := primitive.ObjectIDFromHex("000000000000000000000017")
	TmpIdSMAZonasi, _ := primitive.ObjectIDFromHex("000000000000000000000018")

	/*
		TmpIdSMKAbk, _ := primitive.ObjectIDFromHex("000000000000000000000021")
		TmpIdSMKKetm, _ := primitive.ObjectIDFromHex("000000000000000000000022")
		TmpIdSMKKondisiTertentu, _ := primitive.ObjectIDFromHex("000000000000000000000023")
		TmpIdSMKPerpindahan, _ := primitive.ObjectIDFromHex("000000000000000000000024")
		TmpIdSMKAnakGuru, _ := primitive.ObjectIDFromHex("000000000000000000000025")
		TmpIdSMKPrioritasTerdekat, _ := primitive.ObjectIDFromHex("000000000000000000000026")
		TmpIdSMKPrestasiKejuaraan, _ := primitive.ObjectIDFromHex("000000000000000000000027")
		TmpIdSMKPrestasiIndustri, _ := primitive.ObjectIDFromHex("000000000000000000000028")
		TmpIdSMKPrestasiNilaiRapor, _ := primitive.ObjectIDFromHex("000000000000000000000029")
	*/
	tmpSMAAbk := &models.PpdbOption{
		Id:                  TmpIdSMAAbk,
		Name:                "TemporarySMAAbk",
		Type:                "abk",
		Quota:               0,
		Filtered:            1,
		UpdateQuota:         false,
		PpdbRegistration:    nil,
		RegistrationHistory: nil,
		HistoryShifting:     nil,
	}
	tmpSMAKetm := &models.PpdbOption{
		Id:                  TmpIdSMAKetm,
		Name:                "TemporarySMAKetm",
		Type:                "ketm",
		Quota:               0,
		Filtered:            1,
		UpdateQuota:         false,
		PpdbRegistration:    nil,
		RegistrationHistory: nil,
		HistoryShifting:     nil,
	}
	tmpSMAKondisiTertentu := &models.PpdbOption{
		Id:                  TmpIdSMAKondisiTertentu,
		Name:                "TemporarySMAKondisiTertentu",
		Type:                "kondisi-tertentu",
		Quota:               0,
		Filtered:            1,
		UpdateQuota:         false,
		PpdbRegistration:    nil,
		RegistrationHistory: nil,
		HistoryShifting:     nil,
	}
	tmpSMAPerpindahan := &models.PpdbOption{
		Id:                  TmpIdSMAPerpindahan,
		Name:                "TemporarySMAPerpindahan",
		Type:                "perpindahan",
		Quota:               0,
		Filtered:            1,
		UpdateQuota:         false,
		PpdbRegistration:    nil,
		RegistrationHistory: nil,
		HistoryShifting:     nil,
	}
	tmpSMAAnakGuru := &models.PpdbOption{
		Id:                  TmpIdSMAAnakGuru,
		Name:                "TemporarySMAAnakGuru",
		Type:                "anak-guru",
		Quota:               0,
		Filtered:            1,
		UpdateQuota:         false,
		PpdbRegistration:    nil,
		RegistrationHistory: nil,
		HistoryShifting:     nil,
	}
	tmpSMAPrestasiNilaiRapor := &models.PpdbOption{
		Id:                  TmpIdSMAPrestasiNilaiRapor,
		Name:                "TemporarySMAPrestasiNilaiRapor",
		Type:                "rapor-umum",
		Quota:               0,
		Filtered:            1,
		UpdateQuota:         false,
		PpdbRegistration:    nil,
		RegistrationHistory: nil,
		HistoryShifting:     nil,
	}
	tmpSMAPrestasiPerlombaan := &models.PpdbOption{
		Id:                  TmpIdSMAPrestasiPerlombaan,
		Name:                "TemporarySMAPrestasiPerlombaan",
		Type:                "perlombaan",
		Quota:               0,
		Filtered:            1,
		UpdateQuota:         false,
		PpdbRegistration:    nil,
		RegistrationHistory: nil,
		HistoryShifting:     nil,
	}
	tmpSMAZonasi := &models.PpdbOption{
		Id:                  TmpIdSMAZonasi,
		Name:                "TemporarySMAZonasi",
		Type:                "zonasi",
		Quota:               0,
		Filtered:            1,
		UpdateQuota:         false,
		PpdbRegistration:    nil,
		RegistrationHistory: nil,
		HistoryShifting:     nil,
	}
	optionTypes["abk"] = append(optionTypes["abk"], tmpSMAAbk)
	optionTypes["ketm"] = append(optionTypes["ketm"], tmpSMAKetm)
	optionTypes["kondisi-tertentu"] = append(optionTypes["kondisi-tertentu"], tmpSMAKondisiTertentu)
	optionTypes["perpindahan"] = append(optionTypes["perpindahan"], tmpSMAPerpindahan)
	optionTypes["anak-guru"] = append(optionTypes["anak-guru"], tmpSMAAnakGuru)
	optionTypes["rapor-umum"] = append(optionTypes["rapor-umum"], tmpSMAPrestasiNilaiRapor)
	optionTypes["perlombaan"] = append(optionTypes["perlombaan"], tmpSMAPrestasiPerlombaan)
	optionTypes["zonasi"] = append(optionTypes["zonasi"], tmpSMAZonasi)

	return optionTypes
}
