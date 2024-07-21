package main

import (
	"fmt"

	"github.com/TwiN/go-away"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
type Message struct{
	Text string `json:"text"`
	Sender string `json:"sender"`
}
func main() {
	fmt.Println("Welcome....")
	app:=fiber.New()
	app.Use(cors.New(cors.Config{
        AllowOrigins: "*", // Allow all origins, adjust as needed
        AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))
	app.Post("api/v1/message",DetectAbusiveHandler)
	app.Listen(":3000")
}
func DetectAbusiveHandler(c *fiber.Ctx)error{
	message :=new(Message)
	if err := c.BodyParser(message); err != nil{
		return c.JSON(fiber.Map{"error":"failed to parse json data"})
	}
	isAbusive,abusiveWord:=DetectAbusive(message.Text)
	if isAbusive{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"abusive:":abusiveWord})
	}
	return nil
}
func DetectAbusive(message string)(bool,string){
	if goaway.IsProfane(message){
		abusive:=goaway.ExtractProfanity(message)
		return true,abusive
	}
	return false,""
}
