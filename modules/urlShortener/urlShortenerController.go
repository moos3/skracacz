package urlShortener

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/moos3/skracacz/utils"
)

var CHAR_MAP string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func getUrlEndpoint(c *gin.Context) {
	shortUrl := c.Param("code")
	urlKey := fmt.Sprintf("url:%d", revertShortUrl(shortUrl))
	client := utils.GetClient()

	longUrl, err := client.Get(urlKey).Result()

	if err != nil {
		c.Error(err)
	}

	c.Redirect(302, longUrl)
}

func postUrlEndpoint(c *gin.Context) {

	rawUrl := c.PostForm("url")
	url, err := url.QueryUnescape(rawUrl)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	urlId := saveUrl(url)
	shortUrl := generateShortUrl(urlId)

	c.JSON(http.StatusOK, gin.H{
		"raw_url":   rawUrl,
		"url":       url,
		"short_url": shortUrl,
	})
}

func generateShortUrl(urlId int) (shortUrl string) {
	charCount := len(CHAR_MAP)

	if urlId == 0 {
		return string(CHAR_MAP[0])
	}

	for urlId > 0 {
		shortUrl = string(CHAR_MAP[urlId%charCount]) + shortUrl
		urlId = (urlId / charCount)
	}

	return shortUrl
}

func revertShortUrl(shortUrl string) (urlId int) {
	charCount := len(CHAR_MAP)
	for _, v := range shortUrl {
		urlId = urlId*charCount + strings.Index(CHAR_MAP, string(v))
	}

	return urlId
}

func saveUrl(url string) (urlId int) {
	client := utils.GetClient()
	nextIdValue, err := client.Get("next.url.id").Result()
	if err != nil {
		fmt.Println(err.Error())
	}

	nextId := 0
	if nextIdValue != "" {
		nextId, err = strconv.Atoi(nextIdValue)
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	client.Set(fmt.Sprintf("url:%d", nextId), url, 0)

	client.Set("next.url.id", nextId+1, 0)
	return nextId
}
