package qwen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Text2ImageRequest struct {
	Model      string `json:"model"`
	Input      Input  `json:"input"`
	Parameters Params `json:"parameters"`
}

type Input struct {
	Prompt         string `json:"prompt"`
	NegativePrompt string `json:"negative_prompt,omitempty"`
	RefImg         string `json:"ref_img,omitempty"`
}

type Params struct {
	Style    string  `json:"style,omitempty"`
	Size     string  `json:"size,omitempty"`
	N        int     `json:"n,omitempty"`
	Seed     int     `json:"seed,omitempty"`
	Strength float64 `json:"strength,omitempty"`
	RefMode  string  `json:"ref_mode,omitempty"`
}

type Text2ImageResponse struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Output    Output `json:"output"`
	RequestId string `json:"request_id"`
}

type Output struct {
	TaskID     string `json:"task_id"`
	TaskStatus string `json:"task_status"`
}

type ImageRequest struct {
	ApiKey string
	Input  Input
	Params Params
}

func SubmitText2ImageTask(imageRequest *ImageRequest) (*Text2ImageResponse, error) {
	url := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text2image/image-synthesis"
	imageReq := &Text2ImageRequest{
		Model:      "wanx-v1",
		Input:      imageRequest.Input,
		Parameters: imageRequest.Params,
	}
	requestBody, err := json.Marshal(imageReq)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-DashScope-Async", "enable")
	req.Header.Set("Authorization", "Bearer "+imageRequest.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}
	var response *Text2ImageResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

type TaskStatus string

const (
	Running   TaskStatus = "RUNNING"
	Succeeded TaskStatus = "SUCCEEDED"
	Failed    TaskStatus = "FAILED"
)

type TaskMetrics struct {
	Total     int `json:"TOTAL"`
	Succeeded int `json:"SUCCEEDED"`
	Failed    int `json:"FAILED"`
}

type Result struct {
	URL     string `json:"url"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type TaskOutput struct {
	TaskID      string      `json:"task_id"`
	TaskStatus  TaskStatus  `json:"task_status"`
	Results     []Result    `json:"results,omitempty"`
	TaskMetrics TaskMetrics `json:"task_metrics"`
	Code        string      `json:"code,omitempty"`
	Message     string      `json:"message,omitempty"`
}

type Usage struct {
	ImageCount int `json:"image_count"`
}

type CheckText2ImageTaskResponse struct {
	RequestID string     `json:"request_id"`
	Output    TaskOutput `json:"output"`
	Usage     *Usage     `json:"usage,omitempty"`
}

func CheckText2ImageTask(apiKey, taskID string) (CheckText2ImageTaskResponse, error) {
	url := fmt.Sprintf("https://dashscope.aliyuncs.com/api/v1/tasks/%s", taskID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return CheckText2ImageTaskResponse{}, err
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return CheckText2ImageTaskResponse{}, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CheckText2ImageTaskResponse{}, err
	}
	var response CheckText2ImageTaskResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return CheckText2ImageTaskResponse{}, err
	}
	return response, nil
}

func GenerateImage(req *ImageRequest) (string, error) {
	// Step 1: Submit text to image task
	response, err := SubmitText2ImageTask(req)
	if err != nil {
		return "", err
	}
	// Check if the response code is not empty
	if response.Code != "" {
		return "", fmt.Errorf("text to image task submission failed with code: %s, message: %s", response.Code, response.Message)
	}
	taskID := response.Output.TaskID
	// Step 2: Check text to image task status
	time.Sleep(20 * time.Second)
	for {
		checkResponse, err := CheckText2ImageTask(req.ApiKey, taskID)
		if err != nil {
			return "", err
		}
		// Check if task status is SUCCEEDED or FAILED
		if checkResponse.Output.TaskStatus == Succeeded {
			if len(checkResponse.Output.Results) > 0 {
				if checkResponse.Output.Results[0].Code != "" {
					return "", fmt.Errorf("text to image task failed with code:%s, message:%s",
						checkResponse.Output.Results[0].Code, checkResponse.Output.Results[0].Message)
				}
				return checkResponse.Output.Results[0].URL, nil
			} else {
				return "", fmt.Errorf("text to image task failed with no results")
			}
		} else if checkResponse.Output.TaskStatus == Failed {
			return "", fmt.Errorf("text to image task failed with code:%s ,message: %s",
				checkResponse.Output.Code, checkResponse.Output.Message)
		}
		time.Sleep(3 * time.Second)
	}
}
