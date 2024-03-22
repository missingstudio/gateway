package huggingface

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/missingstudio/ai/gateway/internal/provider/base"
	"github.com/missingstudio/common/errors"
)

type HuggingFaceProvider struct {
	APIKey string
	BaseURL string
}

func (hfp *HuggingFaceProvider) Info() base.ProviderInfo {
	return base.ProviderInfo{
		Name: "HuggingFace",
		Description: "Provider for interacting with HuggingFace's transformer models",
	}
}

func (hfp *HuggingFaceProvider) Models(ctx context.Context) ([]string, error) {
	url := fmt.Sprintf("%s/models", hfp.BaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+hfp.APIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewBadRequest("failed to fetch models from HuggingFace")
	}

	var models []string
	err = json.NewDecoder(resp.Body).Decode(&models)
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (hfp *HuggingFaceProvider) InitiateFineTuning(ctx context.Context, model string, parameters map[string]interface{}) (string, error) {
	url := fmt.Sprintf("%s/fine-tune", hfp.BaseURL)
	payload, err := json.Marshal(map[string]interface{}{
		"model": model,
		"parameters": parameters,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, ioutil.NopCloser(bytes.NewReader(payload)))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+hfp.APIKey)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.NewBadRequest("failed to initiate fine-tuning on HuggingFace")
	}

	var result map[string]string
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result["job_id"], nil
}

func (hfp *HuggingFaceProvider) RetrieveFineTuningResults(ctx context.Context, jobID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/fine-tune/%s", hfp.BaseURL, jobID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+hfp.APIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.NewBadRequest("failed to retrieve fine-tuning results from HuggingFace")
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
