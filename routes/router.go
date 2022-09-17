package routes

import (
	"github.com/gin-gonic/gin"
	v1 "qiublog/api/v1"
	"qiublog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	router := r.Group("api/v1")

	{
		//文章
		router.POST("article/add", v1.ReleaseArticle)  //发布文章
		router.GET("article/list", v1.GetsArticle)     //获取文章列表
		router.GET("article/:id", v1.GetArticle)       //获取单个文章
		router.PUT("article/:id", v1.ModifyArticle)    //更新文章
		router.DELETE("article/:id", v1.DeleteArticle) //删除文章
		//用户
		//菜单子项
		router.POST("menuchild/add", v1.AddMenuchild) //添加菜单子项
		router.GET("menuchild", v1.GetMenuchild)      //获取菜单子项
		router.PUT("menuchild/set", v1.SetMenuchild)  //设置菜单子项
		//分类
		router.POST("category/add", v1.AddCategory) //添加分类
		//上传
		router.POST("upload/image", v1.Upload) //上传文件
	}
	err := r.Run(utils.HttpPort)
	if err != nil {
		return
	}
}
