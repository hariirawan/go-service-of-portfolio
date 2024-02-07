package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"errors"
	"fmt"
)

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
}

type Service interface {
	GetTransactionsByCampaignID(input GetTransactionsByCampaignIDInput) ([]Transaction, error)
	GetTransactionsByUserID(user user.User) ([]Transaction, error)
}

func NewService(repository Repository, campignRepository campaign.Repository) *service {
	return &service{repository, campignRepository}
}

func (s *service) GetTransactionsByCampaignID(input GetTransactionsByCampaignIDInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindByID(input.ID)

	if err != nil {
		return []Transaction{}, err
	}

	fmt.Println(input.ID, campaign.UserID, input.User.ID)
	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transactions, err := s.repository.FindTransactionByCampaignID(input.ID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil

}

func (s *service) GetTransactionsByUserID(user user.User) ([]Transaction, error) {

	transactions, err := s.repository.FindTransactionByUserID(user.ID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil

}
