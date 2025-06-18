package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://www.fiverr.com/ut_works/develop-a-game-prototype-in-unreal-engine-using-blueprints-and-cpp?context_referrer=tailored_homepage_perseus&source=recently_viewed_gigs&ref_ctx_id=81fdc44fce1549e3ba55dfbbcee88d3f&context=recommendation&pckg_id=1&pos=1&context_alg=recently_viewed&imp_id=94934311-06cc-4d35-b9f6-974360f9773e"

func main() {
	fmt.Println("Welcome to urls in Golang")
	//fmt.Println(myurl)
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.User)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	fmt.Println(result.Query())
}
