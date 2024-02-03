package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{
		campaignService,
	}
}

func (h *campaignHandler) GetCampaigns(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.Query("user_id"))
	campaigns, err := h.campaignService.FindCampaigns(userID)

	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaigns(campaigns)
	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", formatter)
	ctx.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaignByID(ctx *gin.Context) {

	var input campaign.GetCampaignDetailInput

	err := ctx.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error to get detail of campaign", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.campaignService.FindCampaignByID(input.ID)

	if campaignDetail.ID == 0 {
		response := helper.APIResponse("An error occurred while getting campaign details based on that campaign ID", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if err != nil {
		response := helper.APIResponse("Error to get campaign detail", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaignDetail(campaignDetail)
	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", formatter)
	ctx.JSON(http.StatusOK, response)

}

func (h *campaignHandler) SaveCampaign(ctx *gin.Context) {

	var input campaign.CreateCampaignInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Error to create campaign", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := ctx.MustGet("currentUser").(user.User)
	input.User = currentUser
	campaign, err := h.campaignService.SaveCampaign(input)

	if err != nil {
		response := helper.APIResponse("Error to create campaign ", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to Create Campaign", http.StatusCreated, "success", campaign)
	ctx.JSON(http.StatusCreated, response)

}

func (h *campaignHandler) UpdateCampaign(ctx *gin.Context) {

	var input campaign.GetCampaignDetailInput

	err := ctx.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Error to get detail of campaign", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var campaign campaign.CreateCampaignInput
	err = ctx.ShouldBindJSON(&campaign)

	if err != nil {
		response := helper.APIResponse("Error to create campaign", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	newCampaign, err := h.campaignService.UpdateCampaign(input.ID, campaign)

	if err != nil {
		response := helper.APIResponse("Error to create campaign ", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to Update Campaign", http.StatusOK, "success", newCampaign)
	ctx.JSON(http.StatusOK, response)

}

func (h *campaignHandler) SaveCampaignImage(ctx *gin.Context) {

	var input campaign.CreateCampaignImageInput
	err := ctx.ShouldBind(&input)

	fmt.Println(input)

	if err != nil {
		response := helper.APIResponse("Error to create campaign image", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := ctx.FormFile("image")

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := ctx.MustGet("currentUser").(user.User)
	path := fmt.Sprintf("images/%d-%s", currentUser.ID, file.Filename)

	err = ctx.SaveUploadedFile(file, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
	}

	_, err = h.campaignService.SaveCampaignImage(input, path)

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Image successfuly uploaded", http.StatusOK, "success", data)
	ctx.JSON(http.StatusOK, response)
}
