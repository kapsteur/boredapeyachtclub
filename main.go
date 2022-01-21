package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
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

	http.HandleFunc("/random", RenderRandom)
	http.ListenAndServe(":80", nil)
}

func RenderRandom(w http.ResponseWriter, req *http.Request) {

	params := req.URL.Query()
	refresh := 5
	if params.Get("refresh") != "" {
		refreshTmp, err := strconv.Atoi(params.Get("refresh"))
		if err != nil {
			refresh = refreshTmp
		}
	}

	img := GetRandom()
	fmt.Fprintf(w, "<html><head><meta http-equiv=\"refresh\" content=\"%d;URL=/random\"></head><body><img src=\"%s\" style=\"width: 100%%;\" /></body></html>", refresh, img)
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
