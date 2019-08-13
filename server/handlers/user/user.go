package user

import (
	repoUser "tobi-api/lib/database/repositories/user"
	"tobi-api/lib/log"
	"tobi-api/server/handlers"

	"github.com/gin-gonic/gin"
)

// POSTUser 新增使用者
func POSTUser(c *gin.Context) {
	jsonUser := new(repoUser.User)
	err := c.BindJSON(jsonUser)
	if err != nil {
		log.Error("handler.POSTUser.BindJSON", err)
		c.AbortWithStatus(400)
		return
	}

	repo := repoUser.NewRepo()
	err = repo.CreateUser(jsonUser)
	if err != nil {
		log.Error("handler.POSTUser.CreateUser", err)
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, handlers.OK(true))

}

// DELETEUser 刪除使用者
func DELETEUser(c *gin.Context) {
	id := c.Param("id")
	repo := repoUser.NewRepo()
	err := repo.DeleteUser(&id)
	if err != nil {
		log.Error("handler.DELETEUser.DeleteUser", err)
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, handlers.OK(true))

}

// PUTUser 更新使用者
func PUTUser(c *gin.Context) {
	repo := repoUser.NewRepo()
	id := c.Param("id")
	jsonUser := new(repoUser.User)

	err := c.BindJSON(jsonUser)
	if err != nil {
		log.Error("handler.PUTUser.BindJSON", err)
		c.AbortWithStatus(400)
		return
	}

	_, err = repo.ReadUser(&id)
	if err != nil {
		log.Error("handler.PUTUser.ReadUser", err)
		c.AbortWithStatus(404)
		return
	}

	err = repo.UpdateUser(&id, jsonUser)
	if err != nil {
		log.Error("handler.PUTUser.UpdateUser", err)
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, handlers.OK(true))

}

// GETUsers 取得所有使用者
func GETUsers(c *gin.Context) {
	repo := repoUser.NewRepo()
	res, err := repo.ReadUsers()
	if err != nil {
		log.Error("handler.GETUsers.ReadUsers", err)
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, handlers.OK(res))

}

// GETUser 取得使用者
func GETUser(c *gin.Context) {
	id := c.Param("id")
	repo := repoUser.NewRepo()
	res, err := repo.ReadUser(&id)
	if err != nil {
		log.Error("handler.GETUser", err)
		c.AbortWithStatus(404)
		return
	}
	c.JSON(200, handlers.OK(res))

}

// GETDoubleUsers 取得兩份的所有使用者
func GETDoubleUsers(c *gin.Context) {
	repo := repoUser.NewRepo()
	res, err := repo.DoubleReadUsers()
	if err != nil {
		log.Error("handler.GETDoubleUsers", err)
		c.AbortWithStatus(400)
		return
	}
	c.JSON(200, handlers.OK(res))
}
