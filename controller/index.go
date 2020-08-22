package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber"
)

//RegisterRoutes register all app routes
func RegisterRoutes(app fiber.Router) {
	app.Get("/sum", sum)
	app.Get("/other", other)
}

func sum(ctx *fiber.Ctx) {
	firstParam, errFirstParam := strconv.ParseFloat(ctx.Query("a"), 32)
	secondParam, errSecondParam := strconv.ParseFloat(ctx.Query("b"), 32)

	if errFirstParam == nil && errSecondParam == nil {
		ctx.Status(200).Send(firstParam + secondParam)
		return
	}
	ctx.Status(400).Send(fmt.Sprintf("Error on trying to sum %s and %s", ctx.Query("a"), ctx.Query("b")))
}

func other(ctx *fiber.Ctx) {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(100)
	b := rand.Intn(100)
	otherSumEndpoint := os.Getenv("OTHER_ENDPOINT")
	resp, err := http.Get(fmt.Sprintf("%s?a=%d&b=%d", otherSumEndpoint, a, b))
	if err != nil {
		ctx.Status(400).Send(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	ctx.Status(200).Send(fmt.Sprintf("sum %v and %v equals %v", a, b, string(body)))

}
