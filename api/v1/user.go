package v1

import (
	"Blog/model"
	"Blog/utils/errmsg"
	"Blog/utils/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var code int


//添加用户
func AddUser(c * gin.Context) {
	var data model.User
	_ = c.ShouldBindJSON(&data)
	msg, code := validator.Validate(&data)
	if  code != errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status"  : code,
			"message" : msg,
		})
		return
	}

	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCSE{
		model.CreateUser(&data)
	}else if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}
	c.JSON(http.StatusOK, gin.H{
		"status" :code,
		"data" : data,
		"message" : errmsg.GetErrmsg(code),
	})
}


//删除用户
func DeleteUser(c * gin.Context) {
	id, _ :=strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status" : code,
		"message": errmsg.GetErrmsg(code),
	})

}


//编辑用户
func EditUser(c * gin.Context) {
	var data model.User
	id, _ :=strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED{
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status" : code,
		"message": errmsg.GetErrmsg(code),
	})
}



//查询单个用户
func GetUser(c *gin.Context) {
	id, _ :=strconv.Atoi(c.Param("id"))
	data, code := model.GetUser(id)
	c.JSON(http.StatusOK, gin.H{
		"date"   : data,
		"status" : code,
		"message": errmsg.GetErrmsg(code),
	})

}


//查询用户列表
func GetUsers(c * gin.Context) {
	username := c.Query("username")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, total := model.GetUsers(username, pageSize, pageNum)
	code = errmsg.SUCCSE

	c.JSON(http.StatusOK, gin.H{
		"status" : code,
		"data"   : data,
		"total"  : total,
		"message" : errmsg.GetErrmsg(code),
	})

}



