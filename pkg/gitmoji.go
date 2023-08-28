package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
)

var gitmojiPersistFileName = "gitmojis.json"

func getGitmojisCachePath() (string, error) {
	return utils.GetCacheDirCreateIfNotExists(ProgramName)
}

func cacheGitmojisTo(cachePath string, gitmojis Gitmojis) {
	vip := viper.New()
	vip.SetConfigType("json")
	vip.Set("gitmojis", gitmojis.Gitmojis)
	var cacheFilePath = path.Join(cachePath, gitmojiPersistFileName)
	err := vip.WriteConfigAs(cacheFilePath)
	if err != nil {
		log.Fatalf("writting gitmoji cache did not work %s", err)
	}
	log.Debugf("Written gitmojis to cache %s", cacheFilePath)
}

func CacheGitmojis(gitmojis Gitmojis) {
	gitmojisCachePath, err := getGitmojisCachePath()
	if err != nil {
		log.Fatalf("getting gitmoji cache dir did not work %s", err)
	}
	cacheGitmojisTo(gitmojisCachePath, gitmojis)
}

func UpdateGitmojis(config Config) (gitmojis Gitmojis) {
	log.Info("update gitmojis local cache")
	log.Infof("Reading from %s and write to cache", config.GitmojisUrl)
	gitmojis, err := getGitmojisHttp(config)
	if err != nil {
		log.Fatalf("Some error happened %s", err)
	}
	CacheGitmojis(gitmojis)
	return
}

func GetGitmojis(config Config) (gitmojis Gitmojis) {
	gitmojis, err := GetGitmojisCache()
	if err != nil {
		log.Info("Haven't been able to read gitmojis from cache")
		UpdateGitmojis(config)
		gitmojis, err = GetGitmojisCache()
		if err != nil {
			log.Fatalf("getting gitmojis from cache did not work after update cache %s", err)
		}
	}
	return
}

func GetGitmojisCache() (gitmojis Gitmojis, err error) {
	cachePath, err := getGitmojisCachePath()
	if err != nil {
		log.Fatal("not able to get the cache file path")
	}
	gitmojis, err = getGitmojisCacheFrom(cachePath)
	return
}

func getGitmojisCacheFrom(cachePath string) (gitmojis Gitmojis, err error) {
	vip := viper.New()
	vip.SetConfigType("json")
	vip.AddConfigPath(cachePath)
	vip.SetConfigName(gitmojiPersistFileName)
	if err = vip.ReadInConfig(); err != nil {
		log.Debugf("Did not find cached values in %s", cachePath)
		return
	}

	err = vip.Unmarshal(&gitmojis)
	return
}

func getGitmojisHttp(config Config) (gitmojis Gitmojis, err error) {
	log.Info(os.Getenv("HTTP_PROXY"))
	log.Info(os.Getenv("HTTPS_PROXY"))
	res, err := http.Get(config.GitmojisUrl)
	if err != nil {
		fmt.Println("error", err)
	}

	if res.Body != nil {
		defer func(Body io.ReadCloser) {
			err = Body.Close()
			if err != nil {
				return
			}
		}(res.Body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &gitmojis)
	if err != nil {
		return
	}
	log.Debugf("Finished retreiving gitmojis from %s", config.GitmojisUrl)
	return
}

func FindGitmoji(emojiOrCode string, gitmojis []Gitmoji) *Gitmoji {
	compareCode := func(code string, gitmoji Gitmoji) bool {
		return gitmoji.Code == code
	}

	compareEmoji := func(emoji string, gitmoji Gitmoji) bool {
		return gitmoji.Emoji == emoji
	}

	re := regexp.MustCompile(`:(.*?):`)
	isCode := re.MatchString(emojiOrCode)
	var compareFkt func(string, Gitmoji) bool
	if isCode {
		compareFkt = compareCode
	} else {
		compareFkt = compareEmoji
	}
	for _, gitmoji := range gitmojis {
		if compareFkt(emojiOrCode, gitmoji) {
			return &gitmoji
		}
	}
	return nil
}
