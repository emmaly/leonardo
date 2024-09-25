package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/emmaly/leonardo"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Initialize the client with your API key
	apiKey := os.Getenv("LEONARDO_API_KEY")
	if apiKey == "" {
		log.Fatal("LEONARDO_API_KEY environment variable is not set")
	}
	client := leonardo.NewClient(apiKey)

	// Get user info
	userInfo, err := client.User.GetUserInfo(context.Background())
	if err != nil {
		log.Fatalf("GetUserInfo failed: %v", err)
	}

	if len(userInfo.UserDetails) == 0 {
		log.Fatal("UserDetails is empty")
	}
	userDetails := userInfo.UserDetails[0]

	// Print user info
	if userDetails.User.ID != nil {
		fmt.Printf("User ID: %s\n", *userDetails.User.ID)
	}
	if userDetails.User.Username != nil {
		fmt.Printf("Username: %s\n", *userDetails.User.Username)
	}
	if userDetails.APIPaidTokens != nil {
		fmt.Printf("API Paid Tokens: %d\n", *userDetails.APIPaidTokens)
	}
	if userDetails.APIPlanTokenRenewalDate != nil {
		fmt.Printf("API Plan Token Renewal Date: %s\n", *userDetails.APIPlanTokenRenewalDate)
	}
	if userDetails.APIConcurrencySlots != nil {
		fmt.Printf("API Concurrency Slots: %d\n", *userDetails.APIConcurrencySlots)
	}
	if userDetails.APISubscriptionTokens != nil {
		fmt.Printf("API Subscription Tokens: %d\n", *userDetails.APISubscriptionTokens)
	}
	if userDetails.PaidTokens != nil {
		fmt.Printf("Paid Tokens: %d\n", *userDetails.PaidTokens)
	}
	if userDetails.SubscriptionGPTTokens != nil {
		fmt.Printf("Subscription GPT Tokens: %d\n", *userDetails.SubscriptionGPTTokens)
	}
	if userDetails.SubscriptionModelTokens != nil {
		fmt.Printf("Subscription Model Tokens: %d\n", *userDetails.SubscriptionModelTokens)
	}
	if userDetails.SubscriptionTokens != nil {
		fmt.Printf("Subscription Tokens: %d\n", *userDetails.SubscriptionTokens)
	}
	if userDetails.TokenRenewalDate != nil {
		fmt.Printf("Token Renewal Date: %s\n", *userDetails.TokenRenewalDate)
	}

	// Generate a prompt
	prompt, err := client.Prompt.GenerateRandomPrompt(context.Background())
	if err != nil {
		log.Fatalf("GenerateRandomPrompt failed: %v", err)
	}

	if prompt.PromptGeneration == nil {
		log.Fatal("Prompt is nil")
	}

	// Print prompt
	if prompt.PromptGeneration.Prompt != nil {
		fmt.Printf("Prompt: %s\n", *prompt.PromptGeneration.Prompt)
	}
	if prompt.PromptGeneration.APICreditCost != nil {
		fmt.Printf("API Credit Cost: %d\n", *prompt.PromptGeneration.APICreditCost)
	}

	// Generate an image from the prompt
	job, err := client.Images.CreateImageGeneration(context.Background(), leonardo.CreateGenerationRequest{
		Prompt:      *prompt.PromptGeneration.Prompt,
		NumImages:   leonardo.Ptr(4),
		PresetStyle: leonardo.Ptr(leonardo.PresetStyleAnime),
	})
	if err != nil {
		log.Fatalf("CreateImageGeneration failed: %v", err)
	}

	// Test for nil job or nil GenerationID
	if job == nil || job.SDGenerationJob.GenerationID == nil {
		log.Fatal("Job is empty or has no GenerationID")
	}

	// Poll for the image generation status
	generation := &leonardo.GetGenerationResponse{}
	for generation.GenerationsByPK.Status == nil || *generation.GenerationsByPK.Status == leonardo.GenerationStatusPending {
		var err error
		generation, err = client.Images.GetImageGeneration(context.Background(), *job.SDGenerationJob.GenerationID)
		if err != nil {
			log.Fatalf("GetImageGeneration failed: %v", err)
		}

		// Test for nil image or nil ID/Status
		if generation == nil || generation.GenerationsByPK.ID == nil || generation.GenerationsByPK.Status == nil {
			log.Fatal("Generation Status is empty or has no status or ID field")
		}

		// Print generation status
		fmt.Printf("Generation ID: %s\n", *generation.GenerationsByPK.ID)
		fmt.Printf("Generation Status: %s\n", *generation.GenerationsByPK.Status)
		fmt.Printf("Generation Status: %s\n", *generation.GenerationsByPK.CreatedAt)
		time.Sleep(1 * time.Second)
	}

	// Print image
	for i, img := range generation.GenerationsByPK.GeneratedImages {
		fmt.Printf("Image %d: %s\n", i, *img.URL)
	}
}
