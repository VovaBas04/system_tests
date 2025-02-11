package service

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"ginProject1/internal/tasks/config"
	"ginProject1/internal/tasks/neuron_model"
	"ginProject1/internal/tasks/postprocess"
	"ginProject1/internal/tasks/repository"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type IService interface {
	ReadNeuronModelAnswer(response *http.Response) (*postprocess.GigaChatResponse, error)
	ToResponse(testCases []postprocess.TestCase) []map[string]string
	Do(gigaChat *neuron_model.GigaChatRequest) (*http.Response, error)
	SaveAction(input string, output string) error
}

type NeuronApiService struct {
	repository repository.IRepository
}

func NewNeuronApiServiceService(repository repository.IRepository) *NeuronApiService {
	return &NeuronApiService{repository: repository}
}

func (service *NeuronApiService) ReadNeuronModelAnswer(response *http.Response) (*postprocess.GigaChatResponse, error) {
	defer response.Body.Close()
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	gigaChatResponse := &postprocess.GigaChatResponse{}
	err = json.Unmarshal(bytes, gigaChatResponse)
	if err != nil {
		return nil, err
	}

	return gigaChatResponse, err
}

func (service *NeuronApiService) ToResponse(testCases []postprocess.TestCase) []map[string]string {
	result := make([]map[string]string, len(testCases))
	for index, testCase := range testCases {
		result[index] = testCase.ToResponse()
	}

	return result
}

func (service *NeuronApiService) Do(gigaChat *neuron_model.GigaChatRequest) (*http.Response, error) {
	const urlApi = "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"
	body := map[string]interface{}{
		"model": gigaChat.Model,
		"messages": []interface{}{
			map[string]interface{}{
				"role":    "user",
				"content": gigaChat.Prompt,
			},
		},
		"max_tokens":         gigaChat.MaxTokens,
		"n":                  1,
		"stream":             false,
		"repetition_penalty": 1,
		"update_interval":    0,
	}
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, urlApi, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	token, err := service.getToken()
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	return client.Do(request)
}

func (service *NeuronApiService) getToken() (string, error) {
	cacheToken, err := service.repository.GetTokenFromCache()
	if err == nil {
		return cacheToken, nil
	}
	if !errors.Is(err, redis.Nil) {
		return "", err
	}

	const urlToken = "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
	urlValues := url.Values{}
	urlValues.Add("scope", "GIGACHAT_API_PERS")
	body := urlValues.Encode()

	request, err := http.NewRequest(http.MethodPost, urlToken, strings.NewReader(body))
	if err != nil {
		return "", err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Accept", "application/json")

	uuid4, _ := uuid.NewUUID()
	request.Header.Set("RqUID", uuid4.String())

	token, err := config.GetBasicToken()
	if err != nil {
		return "", err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", token))
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if response.StatusCode != http.StatusOK {
		return "", errors.New(string(data))
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(data, &tokenResponse)
	if err != nil {
		return "", err
	}

	service.repository.SetTokenToCache(tokenResponse.AccessToken)

	return tokenResponse.AccessToken, nil
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   int    `json:"expire_at"`
}

func (service *NeuronApiService) SaveAction(input string, output string) error {
	return service.repository.SaveAction(input, output)
}
