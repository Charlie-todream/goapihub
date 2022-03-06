package policies

import (
	"github.com/gin-gonic/gin"
	"goapihub/app/models/topic"
	"goapihub/pkg/auth"
)

func CanModifyTopic(c *gin.Context,_topic topic.Topic) bool {

	return auth.CurrentUID(c) == _topic.UserID

}