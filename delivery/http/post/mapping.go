package post

import (
	"sharing-vision-2021/common/helper"
	"sharing-vision-2021/delivery/http/post/model"
	entity "sharing-vision-2021/domain/post"
)

func mapRequestAddPost(req *model.Post, e *entity.Post) {
	e.Title = req.Title
	e.Content = req.Content
	e.Category = req.Category
	e.Status = entity.Status(helper.Capitalize(req.Status))
}
