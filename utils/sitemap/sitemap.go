package sitemap

import (
	"encoding/json"
	"fmt"
	"os"
	"qiublog/model"
	"qiublog/utils"
	"qiublog/utils/feeds"
	"strconv"
	"time"
)

var Feed *feeds.Feed
var Author *feeds.Author

func Db() {
	articles := model.GetAllArticle()
	// 先清空之前数据
	Feed.Items = make([]*feeds.Item, 0, len(articles)+100)
	for _, v := range articles {
		Feed.Items = append(Feed.Items, &feeds.Item{
			Title:       v.Title,
			Link:        &feeds.Link{Href: fmt.Sprintf("%s/post/%d", utils.Config.SiteInfo.Url, v.ID)},
			Author:      Author,
			Description: v.Desc,
			Id:          strconv.Itoa(int(v.ID)),
			Updated:     v.UpdatedAt,
			Created:     v.CreatedAt,
		})
	}
	file, err := os.OpenFile("./sitemap.cache", os.O_CREATE|os.O_WRONLY, 0200)
	if err != nil {
		fmt.Println("sitemap缓存写入失败!", err)
		return
	}
	data, _ := json.Marshal(&Feed.Items)
	file.Write(data)
	file.Close()
}

func InitSitemap() {
	// 获取博客建站时间
	CreatedTime := time.Unix(utils.Config.SiteInfo.ConstructionTime, 0)
	Author = &feeds.Author{Name: utils.Config.SiteInfo.User, Email: utils.Config.SiteInfo.Email}
	Feed = &feeds.Feed{
		Title:       utils.Config.SiteInfo.Name,
		Link:        &feeds.Link{Href: utils.Config.SiteInfo.Url},
		Description: utils.Config.SiteInfo.Desc,
		Author:      Author,
		Created:     CreatedTime,
	}

	// 先读取本地的缓存，没有在遍历数据库
	sitemapCache, err := os.ReadFile("./sitemap.cache")
	if err != nil || utils.Config.Server.AppMode == "debug" {
		Db()
	} else {
		err = json.Unmarshal(sitemapCache, &Feed.Items)
		if err != nil {
			Db()
		}
	}
}
