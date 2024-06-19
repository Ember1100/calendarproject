package service

import (
	"fmt"
	"time"

	"github.com/calendarproject/common"
	"github.com/calendarproject/model"
	"github.com/calendarproject/model/req"
	"github.com/calendarproject/model/resp"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"log"

	"strconv"
)

type IPostController interface {
	Create(ctx *gin.Context)

	Delete(ctx *gin.Context)

	Put(ctx *gin.Context)

	SelectByID(ctx *gin.Context)

	PageList(ctx *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func NewPostController() IPostController {
	db := common.GetDb()
	_ = db.AutoMigrate(&model.SysCalendar{})
	return PostController{DB: db}
}

func (p PostController) PageList(ctx *gin.Context) {

	var pageQuery req.PageQuery
	if err := ctx.ShouldBind(&pageQuery); err != nil {
		log.Println(err)
		resp.Fail(ctx, gin.H{"数据": pageQuery}, "数据验证错误")
		return
	}

	if pageQuery.CreateID == "" || len(pageQuery.CreateID) == 0 {
		resp.Fail(ctx, nil, "create_id不能为空")
		return
	}

	// 转换分页参数
	pageNum, _ := strconv.Atoi(pageQuery.PageNum)
	if pageNum <= 0 {
		pageNum = 1
	}

	pageSize, _ := strconv.Atoi(pageQuery.PageSize)
	if pageSize <= 0 {
		pageSize = 10
	}

	// 分页
	var posts []model.SysCalendar
	p.DB.Where("create_id = ?", pageQuery.CreateID).Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 记录的总条数
	var total int64
	p.DB.Where("create_id = ?", pageQuery.CreateID).Model(model.SysCalendar{}).Count(&total)

	resp.Success(ctx, gin.H{
		"data":     posts,
		"pageNum":  pageNum,
		"pageSize": pageSize,
		"total":    total,
	}, "查询成功")

}

func (p PostController) Create(ctx *gin.Context) {
	var requestPost req.PostRequest

	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Println(err)
		resp.Fail(ctx, gin.H{"数据": requestPost}, "数据验证错误")
		return

	}

	// 添加数据
	post := model.SysCalendar{}
	post.CreateID = requestPost.CreateID
	post.WarnContext = requestPost.WarnContext
	post.SendType = requestPost.SendType

	time, err := time.Parse("2006-01-02 15:04:05", requestPost.WarnDate)
	if err == nil {
		post.WarnDate = time
	}
	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
	}

	resp.Success(ctx, gin.H{
		"提醒信息": post,
	}, "创建成功")

}

func (p PostController) Delete(ctx *gin.Context) {
	// 获取path的ID
	postID := ctx.Params.ByName("id")

	var post model.SysCalendar
	if err := p.DB.First(&post, postID).Error; err != nil {
		resp.Fail(ctx, nil, "数据不存在")
		return
	}

	p.DB.Delete(&post)

	resp.Success(ctx, gin.H{"post": post}, "删除成功")

}

func (p PostController) Put(ctx *gin.Context) {
	var requestPost req.PostRequest

	if err := ctx.ShouldBind(&requestPost); err != nil {
		resp.Fail(ctx, nil, "数据验证错误")
		return

	}

	// 获取path的ID
	postID := ctx.Params.ByName("id")
	var post model.SysCalendar
	if err := p.DB.First(&post, postID).Error; err != nil {
		resp.Fail(ctx, nil, "数据不存在")
		return

	}

	// 更新数据
	if err := p.DB.Model(&post).Updates(requestPost).Error; err != nil {
		resp.Fail(ctx, nil, "更新失败")
		return

	}

	resp.Success(ctx, gin.H{"post": post}, "更新成功")

}

func (p PostController) SelectByID(ctx *gin.Context) {

	// 获取path的ID
	postID := ctx.Params.ByName("id")

	var post model.SysCalendar

	if err := p.DB.First(&post, postID).Error; err != nil {
		fmt.Println(err)
		resp.Fail(ctx, nil, "数据不存在")
		return
	}

	resp.Success(ctx, gin.H{"post": post}, "查询成功")

}
