package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/AndreasAugustin/go-gitmoji-cli/pkg/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"path"
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

func GetGitmojis() (gitmojis Gitmojis) {
	gitmojis, err := GetGitmojisCache()
	if err != nil {
		log.Info("Haven't been able to read gitmojis from cache")
		log.Infof("Reading from %s and write to cache", ConfigInstance.GitmojisUrl)
		gitmojis, err = getGitmojisHttp()
		if err != nil {
			log.Fatalf("Some error happened %s", err)
		}
		CacheGitmojis(gitmojis)
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

func getGitmojisHttp() (gitmojis Gitmojis, err error) {
	res, err := http.Get(ConfigInstance.GitmojisUrl)
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
	log.Debugf("Finished retreiving gitmojis from %s", ConfigInstance.GitmojisUrl)
	return
}