package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	//"github.com/gofiber/fiber/v2"  //not using fiber right now
	"github.com/joho/godotenv" //for loading from the .env file
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)


func main() {
	ctx := context.Background()

	//getting the env variables from the .env file
	err := godotenv.Load()

	if err != nil{
		log.Fatal("Error loading env variables from the .env file: ", err)
	}


	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	
	if err != nil {
    log.Fatal(err)
	}

	defer client.Close()
	
	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text("tell me a story of few lines lol"))
	if err != nil {
		fmt.Print("that was pretty bad")
        log.Fatal(err)
	}

	for i:= 0; i < len(resp.Candidates); i++{
		fmt.Println("element something ", resp.Candidates[i].Content.Parts[0])
		fmt.Println(reflect.TypeOf(resp.Candidates[i].Content.Parts[0]))

	}
	
	log.Println("this is the length of the slice ", len(resp.Candidates))

	
}