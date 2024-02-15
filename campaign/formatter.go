package campaign

import (
	"fmt"
	"strings"
)

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

type CampaignDetailFormatter struct {
	ID               int               `json:"id"`
	UserID           int               `json:"user_id"`
	Name             string            `json:"name"`
	ShortDescription string            `json:"short_description"`
	ImageURL         string            `json:"image_url"`
	GoalAmount       int               `json:"goal_amount"`
	CurrentAmount    int               `json:"current_amount"`
	Slug             string            `json:"slug"`
	Description      string            `json:"description"`
	User             UserFormatter     `json:"user"`
	Images           []ImagesFormatter `json:"images"`
	Perks            []string          `json:"perks"`
}

type UserFormatter struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar_url"`
}

type ImagesFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	formatter := CampaignFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		UserID:           campaign.UserID,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		ImageURL:         "",
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		formatter.ImageURL = fmt.Sprintf("http://localhost:8080/%s", campaign.CampaignImages[0].FileName)
	}

	return formatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)

	}
	return campaignsFormatter
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	detail := CampaignDetailFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		UserID:           campaign.UserID,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		ImageURL:         "",
		User: UserFormatter{
			Name:   campaign.User.Name,
			Avatar: campaign.User.AvatarFileName,
		},
		Images: FormatImages(campaign.CampaignImages),
		Slug:   campaign.Slug,
		Perks:  []string{},
	}

	if len(campaign.CampaignImages) > 0 {
		detail.ImageURL = fmt.Sprintf("http://localhost:8080/%s", campaign.CampaignImages[0].FileName)
	}

	// var perks []string
	// for _, perk := range strings.Split(campaign.Perks, ",") {
	// 	perks = append(perks, strings.TrimSpace(perk))
	// }
	detail.Perks = append(detail.Perks, strings.Split(campaign.Perks, ",")...)

	return detail
}

func FormatImages(images []CampaignImage) []ImagesFormatter {

	campaingnImages := []ImagesFormatter{}

	for _, image := range images {
		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		} else {
			isPrimary = false
		}

		formatter := ImagesFormatter{
			ImageURL:  image.FileName,
			IsPrimary: isPrimary,
		}
		campaingnImages = append(campaingnImages, formatter)

	}

	return campaingnImages
}
