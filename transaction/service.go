package transaction

import (
	"bwastartup/campaign"
	"errors"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionByCampaignID(campaignID ParamTransaction) ([]Transaction, error)
}

func NewService(repository Repository, campignRepository campaign.Repository) *service {
	return &service{repository, campignRepository}
}

func (s *service) GetTransactionByCampaignID(input ParamTransaction) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindByID(input.ID)

	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transactions, err := s.repository.FindTransactionByCampaignID(input.ID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil

}
