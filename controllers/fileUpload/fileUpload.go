package fileUpload

import (
	"fmt"
	"io"

	"github.com/abdul/erp_backend/database/dbAdapter"
	"github.com/abdul/erp_backend/logger"
	tenant "github.com/abdul/erp_backend/models/organization/tenants"
	"github.com/abdul/erp_backend/utils/fileUpload"
	"github.com/gofiber/fiber/v2"
)

var log = logger.Logger

type DownloadUrl struct {
	Url string `json:"url"`
}

func UploadHandler(c *fiber.Ctx) error {
	// Get the filename from the custom header
	fileName := c.Get("filename")
	folder := c.Get("folder")
	tenant_id := c.Get("tenant_id")
	log.Info().Msgf(" uploading logo for tenant id: %v", tenant_id)
	if fileName == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing file name")
	}

	// Parse the file from the FormData
	file, err := c.FormFile("file")
	if err != nil {
		log.Err(err).Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Failed to read file from form data")
	}

	// Open the file
	fileHeader, err := file.Open()
	if err != nil {
		log.Err(err).Msgf("error  %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to open file")
	}
	defer fileHeader.Close()

	// Read the file content into a byte buffer
	fileData, err := io.ReadAll(fileHeader)
	if err != nil {
		log.Err(err).Msgf("error  %v", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to read file content")
	}

	// Validate file data
	if len(fileData) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("Empty file data")
	}

	// Nextcloud credentials and URL

	// Upload file to Nextcloud
	url, err2 := fileUpload.UploadToNextcloud(folder, fileName, fileData)
	if err2 != nil {
		log.Err(err2).Msgf("error  %v", err2)
		return c.Status(fiber.StatusInternalServerError).SendString(err2.Error())
	}
	var tenant, where tenant.Tenant
	tenant.Logo = url
	where.ID = tenant_id
	fmt.Println(tenant_id)
	dbAdapter.DB.Debug()
	if err := dbAdapter.DB.Model(&tenant).Where(&where).Updates(&tenant).Error; err != nil {
		log.Err(err).Msgf("error  %v", err)
		return c.Status(fiber.StatusInternalServerError).
			SendString("file uplaoded not saved try again")
	}
	fmt.Println(tenant)
	// Respond with success
	return c.Status(fiber.StatusCreated).SendString("File uploaded successfully!")
}

func ImagePostHandler(c *fiber.Ctx) error {
	// Parse the JSON body into RequestPayload struct
	var payload DownloadUrl
	if err := c.BodyParser(&payload); err != nil {
		log.Err(err).Msgf("error  %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request payload")
	}

	// Ensure that all required fields are present
	if payload.Url == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Missing URL")
	}

	// Download the image from Nextcloud
	imageStream, err := fileUpload.DownloadImageFromNextcloud(payload.Url)
	if err != nil {
		log.Err(err).Msgf("error  %v", err)
		return c.Status(fiber.StatusInternalServerError).
			SendString(fmt.Sprintf("Error downloading image: %v", err))
	}

	// Set the content type based on the image type (e.g., image/jpeg, image/png)
	c.Set("Content-Type", "image/jpeg")

	// Stream the image back as the response
	_, err = io.Copy(c, imageStream)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusInternalServerError).
			SendString(fmt.Sprintf("Error sending image: %v", err))
	}

	return nil
}
