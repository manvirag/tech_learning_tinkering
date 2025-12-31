package models

import "time"

type PackageCategory string

const (
	Premium PackageCategory = "PREMIUM"
	Gold    PackageCategory = "GOLD"
	Silver  PackageCategory = "SILVER"
)

type Package interface {
	GetType() PackageCategory
	GetRemainingClassCount() int
	SubRemainingClassCount(credit int)
}

type PremiumPackage struct {
	MaxClassCount       int
	Type                PackageCategory
	BuyDate             time.Time
	RemainingClassCount int
}

func (pp *PremiumPackage) GetType() PackageCategory {
	return pp.Type
}
func (pp *PremiumPackage) GetRemainingClassCount() int {
	return pp.RemainingClassCount
}

func (pp *PremiumPackage) SubRemainingClassCount(credit int) {
	pp.RemainingClassCount = pp.RemainingClassCount - credit
}

type GoldPackage struct {
	MaxClassCount       int
	Type                PackageCategory
	BuyDate             time.Time
	RemainingClassCount int
}

func (pp *GoldPackage) GetType() PackageCategory {
	return pp.Type
}
func (pp *GoldPackage) GetRemainingClassCount() int {
	return pp.RemainingClassCount
}
func (pp *GoldPackage) SubRemainingClassCount(credit int) {
	pp.RemainingClassCount = pp.RemainingClassCount + credit
}

type SilverPackage struct {
	MaxClassCount       int
	Type                PackageCategory
	BuyDate             time.Time
	RemainingClassCount int
}

func (pp *SilverPackage) GetType() PackageCategory {
	return pp.Type
}

func (pp *SilverPackage) GetRemainingClassCount() int {
	return pp.RemainingClassCount
}
func (pp *SilverPackage) SubRemainingClassCount(credit int) {
	pp.RemainingClassCount = pp.RemainingClassCount + credit
}

func PackageFactory(category PackageCategory) (Package, error) { // current hardcoding class count
	switch category {
	case Premium:
		return &PremiumPackage{
			MaxClassCount:       10,
			Type:                Premium,
			BuyDate:             time.Now(),
			RemainingClassCount: 10,
		}, nil
	case Gold:
		return &GoldPackage{
			MaxClassCount:       5,
			Type:                Gold,
			BuyDate:             time.Now(),
			RemainingClassCount: 5,
		}, nil
	case Silver:
		return &PremiumPackage{
			MaxClassCount:       2,
			Type:                Silver,
			BuyDate:             time.Now(),
			RemainingClassCount: 2,
		}, nil
	}

	return nil, nil
}
