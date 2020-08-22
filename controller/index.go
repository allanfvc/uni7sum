package controller

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/sony/gobreaker"

	"github.com/gofiber/fiber"
)

//RegisterRoutes register all app routes
func RegisterRoutes(app fiber.Router) {
	initBreaker()
	app.Get("/sum", sum)
	app.Get("/other", other)
	app.Get("/other-safe", otherWithBreaker)
	app.Get("/fallback", func(ctx *fiber.Ctx) {
		ctx.Status(400).Send("Sorry, a connection error occurred. Try again in a few minutes.")
	})
	app.Get("/many", func(ctx *fiber.Ctx) {
		ctx.Status(400).Send("Sorry, the service is currently busy. Try again in a few minutes.")
	})
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

func otherWithBreaker(ctx *fiber.Ctx) {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(100)
	b := rand.Intn(100)
	otherSumEndpoint := os.Getenv("OTHER_ENDPOINT")
	url := fmt.Sprintf("%s?a=%d&b=%d", otherSumEndpoint, a, b)
	body, err := getWithBreaker(url)
	if err != nil {
		log.Println(err)
		if errors.Is(err, gobreaker.ErrOpenState) {
			ctx.Redirect("fallback")
		} else if errors.Is(err, gobreaker.ErrOpenState) {
			ctx.Redirect("many")
		} else {
			ctx.Status(400).Send(err)
		}
		return
	}
	ctx.Status(200).Send(fmt.Sprintf("sum %v and %v equals %v", a, b, string(body)))

}
