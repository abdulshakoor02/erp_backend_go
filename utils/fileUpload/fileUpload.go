package fileUpload

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/abdul/erp_backend/config"
)

func UploadToNextcloud(folder, fileName string, fileData []byte) (string, error) {
	fmt.Println("inside fileupload")
	// Define the upload URL
	uploadURL := fmt.Sprintf(
		"%s/%s/%s/%s",
		config.NEXT_CLOUD_URL,
		config.NEXT_CLOUD_USERNAME,
		folder,
		fileName,
	)

	// Create a new HTTP request to upload the file
	req, err := http.NewRequest("PUT", uploadURL, bytes.NewReader(fileData))
	if err != nil {
		return "failed to upload file", fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/octet-stream")
	req.SetBasicAuth(config.NEXT_CLOUD_USERNAME, config.NEXT_CLOUD_PASSWORD)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "failed to upload file", fmt.Errorf("failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status
	if resp.StatusCode != http.StatusCreated {
		fmt.Println(resp)
		return "failed to upload file", fmt.Errorf("failed to upload file: %s", resp.Status)
	}

	return uploadURL, nil
}

func DownloadImageFromNextcloud(nextCloudUrl string) (io.Reader, error) {
	// Create the HTTP request to the WebDAV URL
	req, err := http.NewRequest("GET", nextCloudUrl, nil)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set basic authentication
	req.SetBasicAuth(config.NEXT_CLOUD_USERNAME, config.NEXT_CLOUD_PASSWORD)

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file: status code %d", resp.StatusCode)
	}

	return resp.Body, nil
}
