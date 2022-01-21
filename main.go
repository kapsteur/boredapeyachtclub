package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type URIJson struct {
	Image      string `json:"image"`
	Attributes []struct {
		TraitType string `json:"trait_type"`
		Value     string `json:"value"`
	} `json:"attributes"`
}

func main() {

	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	log.Printf("APP Started")
	log.Printf("PORT:%s", os.Getenv("PORT"))
	http.HandleFunc("/random", RenderRandom)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func RenderRandom(w http.ResponseWriter, req *http.Request) {

	params := req.URL.Query()
	refresh := 30
	if params.Get("refresh") != "" {
		refreshTmp, err := strconv.Atoi(params.Get("refresh"))
		if err == nil {
			refresh = refreshTmp
		}
	}

	img := GetRandom()
	resp, err := http.Get(img)
	if err != nil {
		fmt.Println("Error: File could not be opened")
		log.Printf("err:%s", err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	imgDecoded, _, err := image.Decode(resp.Body)

	if err != nil {
		log.Printf("Error:%s", err)
	}

	r, g, b, a := imgDecoded.At(100, 100).RGBA()
	fmt.Fprintf(w, "<html><head><meta http-equiv=\"refresh\" content=\"%d;URL=/random\"></head><body style=\"background-color: rgba(%d,%d,%d,%d);\"><img src=\"%s\" style=\"width:100%%;position:absolute;bottom:0;\" /></body></html>", refresh, r/257, g/257, b/257, a/257, img)
}

func GetRandom() string {
	rand.Seed(time.Now().Unix())
	r := rand.Intn(10000)

	resp, err := http.Get(fmt.Sprintf("https://gateway.ipfs.io/ipfs/QmeSjSinHpPnmXmspMjwiXyN6zS4E9zccariGR3jxcaWtq/%d", r))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	var uriJson URIJson

	err = json.NewDecoder(resp.Body).Decode(&uriJson)
	if err != nil {
		return ""
	}

	log.Printf("%v", uriJson)

	return strings.Replace(uriJson.Image, "ipfs:/", "https://gateway.ipfs.io/ipfs/", -1)
}
