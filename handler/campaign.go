package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/salmaqnsGH/crowdfunding-app/user"

	"github.com/gin-gonic/gin"
	"github.com/salmaqnsGH/crowdfunding-app/campaign"
	"github.com/salmaqnsGH/crowdfunding-app/helper"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Failed get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Lists of campaigns", http.StatusOK, "success", campaign.FormatCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("Failed get detail of campaign 1", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignByID(input)
	if err != nil {
		response := helper.APIResponse("Failed get detail of campaign 2", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Lists of campaigns", http.StatusOK, "success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create campaign 1", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("current_user").(user.User)

	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("Failed to create campaign 2", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success create campaign", http.StatusOK, "success", campaign.FormatCampaign(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var inputID campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign 1", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput
	currentUser := c.MustGet("current_user").(user.User)
	inputData.User = currentUser

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to update campaign 2", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedCampaign, err := h.service.UpdateCampaign(inputID, inputData)
	if err != nil {
		response := helper.APIResponse("Failed to update campaign 3", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success updated campaign", http.StatusOK, "success", campaign.FormatCampaign(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CreateCampaignImageInput

	err := c.ShouldBind(&input)

	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to upload campaign image 1", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("current_user").(user.User)
	input.User = currentUser

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image 2", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := currentUser.ID
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image 3", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveCampaignImage(input, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload campaign image 4", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Successfully upload campaign image", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)

}
