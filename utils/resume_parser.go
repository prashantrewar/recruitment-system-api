package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"recruitment-system/models"
)

func ParseResume(file multipart.File, apiKey string) (models.Profile, error) {
	var profile models.Profile

	// Prepare the file for upload
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("resume", "resume.pdf")
	if err != nil {
		return profile, fmt.Errorf("error creating form file: %w", err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return profile, fmt.Errorf("error copying file: %w", err)
	}

	writer.Close()

	// Make the request to the resume parser API
	req, err := http.NewRequest("POST", "https://api.apilayer.com/resume_parser/upload", body)
	if err != nil {
		return profile, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return profile, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return profile, fmt.Errorf("failed to parse resume, status code: %d", resp.StatusCode)
	}

	var apiResponse struct {
		Name      string `json:"name"`
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		Education []struct {
			Name string `json:"name"`
		} `json:"education"`
		Experience []struct {
			Name  string `json:"name"`
			Dates []string `json:"dates"`
		} `json:"experience"`
		Skills []string `json:"skills"`
	}

	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return profile, fmt.Errorf("error decoding response: %w", err)
	}

	// Map API response to profile model
	profile.Name = apiResponse.Name
	profile.Email = apiResponse.Email
	profile.Phone = apiResponse.Phone
	for _, edu := range apiResponse.Education {
		profile.Education += edu.Name + "; "
	}
	for _, exp := range apiResponse.Experience {
		profile.Experience += exp.Name + " (" + fmt.Sprint(exp.Dates) + "); "
	}
	profile.Skills = fmt.Sprint(apiResponse.Skills)

	return profile, nil
}
