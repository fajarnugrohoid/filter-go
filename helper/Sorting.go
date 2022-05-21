package helper

import (
	"filterisasi/models/domain"
	"sort"
)

type ByScore []domain.PpdbRegistration

func (m ByScore) Len() int           { return len(m) }
func (m ByScore) Less(i, j int) bool { return m[i].Score > m[j].Score }
func (m ByScore) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

type ByDistance []domain.PpdbRegistration

func (m ByDistance) Len() int           { return len(m) }
func (m ByDistance) Less(i, j int) bool { return m[i].Distance1 < m[j].Distance1 }
func (m ByDistance) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func SortByDistanceAndAge(members []domain.PpdbRegistration) {
	sort.SliceStable(members, func(i, j int) bool {
		mi, mj := members[i], members[j]
		switch {
		case mi.Distance != mj.Distance:
			return mi.Distance < mj.Distance
		default:
			return mi.BirthDate < mj.BirthDate
		}
	})
}

func SortByScoreAndAge(members []domain.PpdbRegistration) {
	sort.SliceStable(members, func(i, j int) bool {
		mi, mj := members[i], members[j]
		switch {
		case mi.Score != mj.Score:
			return mi.Score < mj.Score
		default:
			return mi.BirthDate < mj.BirthDate
		}
	})
}

func SortByAnakGuruAndAge(members []domain.PpdbRegistration) {
	sort.SliceStable(members, func(i, j int) bool {
		mi, mj := members[i], members[j]
		switch {
		case mi.Priority != mj.Priority:
			return mi.Priority < mj.Priority
		case mi.Distance != mj.Distance:
			return mi.Distance < mj.Distance
		default:
			return mi.BirthDate < mj.BirthDate
		}
	})
}
