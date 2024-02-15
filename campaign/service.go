package campaign

import (
	"errors"
	"fmt"
)

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
	FindCampaignByID(ID int) (Campaign, error)
	SaveCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(id int, input CreateCampaignInput) (Campaign, error)
	SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error)
	DeleteCampaignByID(input DeleteCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil

	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) FindCampaignByID(ID int) (Campaign, error) {
	campaign, err := s.repository.FindByID(ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) SaveCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.Description = input.Description
	campaign.ShortDescription = input.ShortDescription
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID

	newCampaign, err := s.repository.Save(campaign)

	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) UpdateCampaign(id int, input CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(id)

	if err != nil {
		return campaign, err
	}

	if campaign.UserID != input.User.ID {
		return Campaign{}, errors.New("not an owner of the campaign")
	}

	campaign.Name = input.Name
	campaign.Description = input.Description
	campaign.ShortDescription = input.ShortDescription
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID
	newCampaign, err := s.repository.Update(campaign)

	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) SaveCampaignImage(input CreateCampaignImageInput, fileLocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindByID(input.CampaignID)

	fmt.Println(campaign)

	if err != nil {
		return CampaignImage{}, err
	}

	if campaign.UserID != input.User.ID {
		return CampaignImage{}, errors.New("not an owner of the campaign")
	}

	isPrimary := 0

	if input.IsPrimary {
		isPrimary = 1

		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return CampaignImage{}, err
		}

	}

	campaignImage := CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.FileName = fileLocation
	campaignImage.IsPrimary = isPrimary
	newCampaignImage, err := s.repository.CreateImage(campaignImage)

	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil
}

func (s *service) DeleteCampaignByID(input DeleteCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)

	if err != nil {
		return campaign, errors.New("Campaign not found")
	}

	if campaign.UserID != input.User.ID {
		return Campaign{}, errors.New("not an owner of the campaign")
	}

	_, err = s.repository.Delete(campaign)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
