package transaction

import "gorm.io/gorm"

type Repository interface {
	FindTransactionByCampaignID(campaignID int) ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) FindTransactionByCampaignID(campaignID int) (transactions []Transaction, err error) {
	err = r.db.Preload("User").Preload("Campaign").Where("campaign_id=?", campaignID).Order("created_at desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
