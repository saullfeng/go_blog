package service

import (
	. "blog_api/db"
	"blog_api/entity"
	"blog_api/utils"

	"github.com/jinzhu/gorm"
)

type Blog entity.Blog

func (Blog) TableName() string {
	return "blog"
}

//查询博客内容及类型名
func (blog *Blog) FindtOneTypeName() (b *Blog) {
	b = new(Blog)
	Db.Table(" blog b").Select("b.*,bt.name as type_name").
		Joins("left join blog_type bt on b.typeId = bt.id").
		Where("b.id = ?", blog.Id).Order("bt.sort asc").Find(b)
	return
}

//查询博客下一条
func (blog *Blog) FindNextOne() (b *Blog) {
	b = new(Blog)
	result := Db.Where("id > ?", blog.Id).First(b)
	if result.Error != nil {
		return nil
	}
	return
}

//查询博客上一条
func (blog *Blog) FindLastOne() (b *Blog) {
	b = new(Blog)
	result := Db.Where("id < ?", blog.Id).Order("id").First(b)
	if result.Error != nil {
		return nil
	}
	return
}

//查询博客评论
func (blog *Blog) FindCommentByBlog() []Comment {
	comment := make([]Comment, 0)
	result := Db.Table("comment").Where("blog_id = ? and status = 1", blog.Id).Order("add_time asc").Find(&comment)
	if result.Error != nil {
		return nil
	}
	return comment
}

//查询博客类型统计
func (blog *Blog) FindByTypeCount() (count int) {
	Db.Model(blog).Where("typeId = ?", blog.TypeId).Count(&count)
	return
}

//查找博客列表
func (blog *Blog) FindList(page *utils.Page) ([]*Blog, error) {
	bs := make([]*Blog, 0)
	curDb := Db.Table("blog b").Select("b.*,bt.name as type_name").
		Joins("left join blog_type bt on b.typeId = bt.id")
	if blog.TypeId > 0 {
		curDb = curDb.Where("b.typeId = ?", blog.TypeId)
	}
	result := curDb.Limit(page.Size).Offset(page.GetStart()).Order("`add_time` asc").Find(&bs)
	return bs, result.Error
}

//查找单个博客
func (blog *Blog) FindOne() (b *Blog) {
	b = new(Blog)
	Db.Where("id = ?", blog.Id).First(b)
	return
}

//数量统计
func (blog *Blog) Count() (count int) {
	Db.Model(blog).Count(&count)
	return
}

//更新点击次数
func (blog *Blog) UpdateClick() *gorm.DB {
	return Db.Model(blog).Where("id = ?", blog.Id).Update("click_hit", gorm.Expr("click_hit + ?", 1))

}

//更新评论次数
func (blog *Blog) UpdateReplay() *gorm.DB {
	return Db.Model(blog).Where("id = ?", blog.Id).Update("replay_hit", gorm.Expr("replay_hit + ?", 1))
}

//博客插入
func (blog *Blog) Insert() *gorm.DB {
	return Db.Create(blog)
}

//博客更新
func (blog *Blog) Update() *gorm.DB {
	return Db.Save(blog)
}

//博客删除
func (blog *Blog) Delete() *gorm.DB {
	return Db.Delete(blog)
}
