package main

import (
	"fmt"
	"strings"

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
	isAbusive,word:=DetectAbusive("hi you")
	if isAbusive{
		fmt.Println(word)
	}
	app.Use(cors.New(cors.Config{
        AllowOrigins: "*", // Allow all origins, adjust as needed
        AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
        AllowHeaders: "Origin, Content-Type, Accept",
    }))
	app.Post("api/v1/message",DetectAbusiveHandler)
	app.Listen(":3000")
}
var customProfanities = []string{
	"anal",
	"ano",
	"culo",
	"trasero",
	"saco de bolas",
	"bolas",
	"bastardo",
	"perra",
	"zorra",
	"mamón",
	"mamón",
	"paja",
	"cojones",
	"cojones",
	"erección",
	"teta",
	"cabrón",
	"culo",
	"polla",
	"clítoris",
	"verga",
	"negro",
	"mierda",
	"correrse",
	"coño",
	"pene",
	"consolador",
	"imbécil",
	"lesbiana",
	"maricón",
	"carajo",
	"felación",
	"felación",
	"felación",
	"joder",
	"cagador",
	"coño",
	"vete a la mierda",
	"puta",
	"caliente",
	"incesto",
	"imbécil",
	"correrse",
	"labios",
	"masturbar",
	"chocho",
	"desnudo",
	"nazi",
	"negro",
	"pezón",
	"pezones",
	"desnudo",
	"pedófilo",
	"pene",
	"orinar",
	"cagar",
	"porno",
	"polla",
	"prostituta",
	"vello púbico",
	"coño",
	"coño",
	"maricón",
	"violación",
	"violador",
	"retrasado",
	"trabajo anal",
	"escroto",
	"sexo",
	"mierda",
	"puta",
	"corrida",
	"cállate",
	"chúpame",
	"tetas",
	"teta",
	"teta",
	"turd",
	"coño",
	"vagina",
	"paja",
	"puta",
}

func DetectAbusiveHandler(c *fiber.Ctx)error{
	message :=new(Message)
	if err := c.BodyParser(message); err != nil{
		return c.JSON(fiber.Map{"error":"failed to parse json data"})
	}
	isAbusive,abusiveWord:=DetectAbusive(message.Text)
	if isAbusive{
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"abusive":abusiveWord})
	}
	return nil
}
func DetectAbusive(message string)(bool,string){
	// Check against custom profanities
	for _, profanity := range customProfanities {
		if strings.Contains(strings.ToLower(message), strings.ToLower(profanity)) {
			return true, profanity
		}
	}
	if goaway.IsProfane(message){
		abusive:=goaway.ExtractProfanity(message)
		return true,abusive
	}
	return false,""
}
