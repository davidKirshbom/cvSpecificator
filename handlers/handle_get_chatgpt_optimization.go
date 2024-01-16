package handlers

import (
	"os"
	"strings"

	"github.com/davidKirshbom/cvSpecificator/models"
	httpClient "github.com/davidKirshbom/cvSpecificator/utils/http_client"
)

func HandleGetChatgptOptimization(experiences []models.Experience,jobDescription string) []models.Experience {
	client:=httpClient.New()
	defer client.Close()
	optimizeExperiences := []models.Experience{}
	for _,experience := range experiences{
	threadResult:= openAndRunThread(client,jobDescription,experience.Role.Explain)
	waitForThreadToComplete(client,threadResult.ThreadId,threadResult.RunId)
	messagesList:=getMessages(client,threadResult.ThreadId)
	optimizeExperience := experience.Clone()
	optimizeExperience.Role.Explain =messagesList.Data[0].Content[0].Text.Value
	optimizeExperiences = append(optimizeExperiences,optimizeExperience )
	}
	return optimizeExperiences 
}

func  openAndRunThread(client httpClient.HttpClient,jobDescription,experience string) ThreadResult{
	var createThreadResult = OpenAiThreadOpenResponse{}
	// open assistant thread
	client.Fetch(httpClient.HttpFetchOptions{
		Method: "POST",
		Url:    "https://api.openai.com/v1/threads/runs",
		Headers: map[string]string{
			"Authorization": "Bearer " + os.Getenv("OPENAI_API_KEY"),
			"Content-Type":  "application/json",
			"OpenAI-Beta":   "assistants=v1",
		},
		Body: strings.ReplaceAll(`{"assistant_id":"`+os.Getenv("OPENAI_ASSISTANT_ID")+`","thread":{"messages": [{"role":"user","content":"job description: `+jobDescription+`.and my experience at one work place `+experience+`"  }]}}`, "\\", ""),
	}, &createThreadResult)
	return ThreadResult{ThreadId: createThreadResult.ThreadId,RunId:createThreadResult.ID}
}
func waitForThreadToComplete(client httpClient.HttpClient,threadId string,runId string) OpenAiThreadRunResponse{
	runResponse := OpenAiThreadRunResponse{}
	for runResponse.Status != "completed" {
		client.Fetch(httpClient.HttpFetchOptions{
			Method: "GET",
			Url:    `https://api.openai.com/v1/threads/` + threadId + `/runs/` + runId,
			Headers: map[string]string{
				"Authorization": "Bearer " + os.Getenv("OPENAI_API_KEY"),
				"Content-Type":  "application/json",
				"OpenAI-Beta":   "assistants=v1",
			},
			
		}, &runResponse)
	}
	return runResponse
}
func getMessages(client httpClient.HttpClient,threadId string) OpenAiThreadMessage{
	messagesList := OpenAiThreadMessage{}
	client.Fetch(httpClient.HttpFetchOptions{
		Method: "GET",
		Url:    `https://api.openai.com/v1/threads/` + threadId + `/messages`,
		Headers: map[string]string{
			"Authorization": "Bearer " + os.Getenv("OPENAI_API_KEY"),
			"Content-Type":  "application/json",
			"OpenAI-Beta":   "assistants=v1",
		},
	}, &messagesList)
	return messagesList
}
type ThreadResult struct{
	ThreadId string `json:"thread_id"`
	RunId string `json:"run_id"`
}
type OpenAiThreadOpenResponse struct {
	ID string `json:"id"`
	ThreadId string `json:"thread_id"`
	AssistantId string `json:"assistant_id"`
	Object string `json:"object"`
	CreatedAt int `json:"created_at"`
	Metadata any `json:"metadata"`
	Status string `json:"status"`
	Model string `json:"model"`
	Instructions string `json:"instructions"`
  
}
type OpenAiThreadRunResponse struct {
	ID string `json:"id"`
	Object string `json:"object"`
	CreatedAt int `json:"created_at"`
	AssistantId string `json:"assistant_id"`
	ThreadId string `json:"thread_id"`
	Status string `json:"status"`
	StartedAt int `json:"started_at"`
	ExpiresAt int `json:"expires_at"`
	CancelledAt int `json:"cancelled_at"`
	FailedAt int `json:"failed_at"`
	CompletedAt int `json:"completed_at"`
	LastError int `json:"last_error"`
	Model string `json:"model"`
	Instructions string `json:"instructions"`
	Tools []struct{
		Type string `json:"type"`
	}
	FileIds []string `json:"file_ids"`
	Metadata any `json:"metadata"`
	
}

type OpenAiThreadMessage struct {
	Object string `json:"object"`
	Data []struct{
		Id string `json:"id"`
		Object string `json:"object"`
		CreatedAt int `json:"created_at"`
		ThreadId string `json:"thread_id"`
		Role string `json:"role"`
		Content []struct{
			Type string `json:"type"`
			Text struct{
				Value string `json:"value"`
				Annotations []string `json:"annotations"`
			} `json:"text"`
		} `json:"content"`
		FileIds []string `json:"file_ids"`
		AssistantId string `json:"assistant_id"`
		RunId string `json:"run_id"`
		Metadata any `json:"metadata"`
	} `json:"data"`
	FirstId string `json:"first_id"`
	LastId string `json:"last_id"`
	HasMore bool `json:"has_more"`
}
	