package controllers

import (
	"apigetspot/models"
	"apigetspot/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"time"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

//URL Mapping
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *UserController) Post() {
	token := c.Ctx.Input.Header("Authorization")
	//id := c.Ctx.Input.Header("id")
	et := utils.EasyToken{}
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	validation, err := et.ValidateToken(token)
	if !validation {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = Response{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}
	var v models.User

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if message := utils.CheckNewUserPost(v.Name, v.Username, v.Password,
			v.Email, v.Location, v.Phone, v.Gender); message != "ok" {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, message, ""}
			c.ServeJSON()
			return
		}
		if models.CheckUsername(v.Username) {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, "The username " + v.Username + " is not available", ""}
			c.ServeJSON()
			return
		}
		if models.CheckEmail(v.Email) {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, "email is already exist, Please try another email", ""}
			c.ServeJSON()
			return
		}

		if user, err := models.AddUser(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			var returnData = &UserSuccessLoginData{user.Token, user.Username}
			c.Data["json"] = &Response{0, 0, "success", returnData}
		} else {
			c.Data["json"] = &Response{1, 1, "error", err.Error()}
		}
	} else {
		c.Data["json"] = &Response{1, 1, "error", err.Error()}
	}
	c.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (c *UserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int = 10
	var offset int

	token := c.Ctx.Input.Header("Authorization")
	//id := c.Ctx.Input.Header("id")
	et := utils.EasyToken{}
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	validation, err := et.ValidateToken(token)
	if !validation {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = Response{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	} else {
		fields = strings.Split("Name,Username,Email,Location,Phone,Gender", ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get User by id
// @Param	id		path 	string	true "The key for static block"
// @Success 200 {object} models.AirAd
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	token := c.Ctx.Input.Header("Authorization")
	//idStr := c.Ctx.Input.Param("id")
	idStr := c.Ctx.Input.Param(":id")
	//token := c.Ctx.Input.Param(":token")
	et := utils.EasyToken{}
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	valido, err := et.ValidateToken(token)
	if !valido {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = Response{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}

	id, _ := strconv.Atoi(idStr)
	v, err := models.GetUserById(id)
	if v == nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()

}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (c *UserController) Put() {
	token := c.Ctx.Input.Header("Authorization")
	//id := c.Ctx.Input.Header("id")
	et := utils.EasyToken{}
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	validation, err := et.ValidateToken(token)
	if !validation {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = Response{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.User{Id: id}
	fmt.Println("put tes", v)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUserById(&v); err == nil {
			c.Data["json"] = successReturn
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (c *UserController) Delete() {
	token := c.Ctx.Input.Header("Authorization")
	//id := c.Ctx.Input.Header("id")
	et := utils.EasyToken{}
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	validation, err := et.ValidateToken(token)
	if !validation {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = Response{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteUser(id); err == nil {
		c.Data["json"] = successReturn
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @router /register [post]
func (c *UserController) Register() {
	var v models.User

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if message := utils.CheckNewUserPost(v.Name, v.Username, v.Password,
			v.Email, v.Location, v.Phone, v.Gender); message != "ok" {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, message, ""}
			c.ServeJSON()
			return
		}
		if models.CheckUsername(v.Username) {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, "The username " + v.Username + " is not available", ""}
			c.ServeJSON()
			return
		}
		if models.CheckEmail(v.Email) {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, "Email is already exist, Please try another email", ""}
			c.ServeJSON()
			return
		}

		if user, err := models.AddUser(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			var returnData = &UserSuccessLoginData{user.Token, user.Username}
			c.Data["json"] = &Response{0, 0, "success", returnData}
		} else {
			c.Data["json"] = &Response{1, 1, "error", err.Error()}
		}
	} else {
		c.Data["json"] = &Response{1, 1, "error", err.Error()}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get User by id
// @Param	id		path 	string	true "The key for static block"
// @Success 200 {object} models.AirAd
// @Failure 403 :id is empty
// @router /activateAccount/Activation_Code=:token&Email_From=:email [get]
// func (c *UserController) ActivateAccount() {
// 	var reqData struct {
// 		Activation_Code string `valid:"required"`
// 		Email_From      string `valid:"required"`
// 	}
// 	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqData); err == nil {
// 		if message := utils.CheckActivationAccount(reqData.Activation_Code, reqData.Email_From); message != "ok" {
// 			c.Ctx.ResponseWriter.WriteHeader(403)
// 			c.Data["json"] = Response{403, 403, message, ""}
// 			c.ServeJSON()
// 			return
// 		}
// 		if ok, user := models.Activate(reqData.Activation_Code, reqData.Email_From); ok {
// 			et := utils.EasyToken{}
// 			validation, err := et.ValidateToken(user.Token)
// 			if !validation {
// 				et = utils.EasyToken{
// 					Username: user.Username,
// 					Uid:      int64(user.Id),
// 					Expires:  time.Now().Unix() + 2*3600,
// 				}
// 				token, err = et.GetToken()
// 				if token == "" || err != nil {
// 					c.Data["json"] = errUserToken
// 					c.ServeJSON()
// 					return
// 				} else {
// 					models.UpdateUserToken(user, token)
// 				}
// 			} else {
// 				token = user.Token
// 			}
// 			models.UpdateUserLastLogin(user)

// 			var returnData = &UserSuccessLoginData{token, user.Username}
// 			c.Data["json"] = &Response{0, 0, "ok", returnData}
// 		} else {
// 			c.Data["json"] = &errNoUserOrPass
// 		}

// 	}
// }

// @router /login [post]
func (c *UserController) Login() {
	var reqData struct {
		Username string `valid:"Required"`
		Password string `valid:"Required"`
	}
	var token string

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqData); err == nil {
		if message := utils.CheckUsernamePassword(reqData.Username, reqData.Password); message != "ok" {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, message, ""}
			c.ServeJSON()
			return
		}
		if ok, user := models.Login(reqData.Username, reqData.Password); ok {
			et := utils.EasyToken{}
			validation, err := et.ValidateToken(user.Token)
			if !validation {
				et = utils.EasyToken{
					Username: user.Username,
					Uid:      int64(user.Id),
					Expires:  time.Now().Unix() + 2*3600,
				}
				token, err = et.GetToken()
				if token == "" || err != nil {
					c.Data["json"] = errUserToken
					c.ServeJSON()
					return
				} else {
					models.UpdateUserToken(user, token)
				}
			} else {
				token = user.Token
			}
			models.UpdateUserLastLogin(user)

			var returnData = &UserSuccessLoginData{token, user.Username}
			c.Data["json"] = &Response{0, 0, "ok", returnData}
		} else {
			c.Data["json"] = &errNoUserOrPass
		}
	} else {
		c.Data["json"] = &errNoUserOrPass
	}
	c.ServeJSON()
}

// @router /updateAccount [put]
func (c *UserController) UpdateAccount() {
	token := c.Ctx.Input.Header("Authorization")
	//id := c.Ctx.Input.Header("id")
	et := utils.EasyToken{}
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	validation, err := et.ValidateToken(token)
	if !validation {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = Response{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}
	tokens := c.Ctx.Input.Header("Authorization")
	v := models.User{Token: tokens}
	fmt.Println("tes", &v)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if message := utils.CheckUpdateUserPost(v.Name, v.Username,
			v.Email, v.Location, v.Phone, v.Gender); message != "ok" {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, message, ""}
			c.ServeJSON()
			return
		}
		if models.CheckUsernameExceptMe(v.Username, tokens) {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, "The username " + v.Username + " is not available", ""}
			c.ServeJSON()
			return
		}
		if models.CheckEmailExceptMe(v.Email, tokens) {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, "Email is already exist, Please try another email", ""}
			c.ServeJSON()
			return
		}

		//
		if err := models.UpdateUserByToken(&v, tokens); err == nil {
			c.Data["json"] = successReturn
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @router /updatePassword [put]
func (c *UserController) UpdatePassword() {
	token := c.Ctx.Input.Header("Authorization")
	//id := c.Ctx.Input.Header("id")
	et := utils.EasyToken{}
	//token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	validation, err := et.ValidateToken(token)
	if !validation {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = Response{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}
	var reqData struct {
		OldPassword    string `valid:"Required"`
		NewPassword    string `valid:"Required"`
		RetypePassword string `valid:"Required"`
	}
	// var token string
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqData); err == nil {
		if message := utils.CheckPassword(reqData.OldPassword, reqData.NewPassword, reqData.RetypePassword); message != "ok" {
			c.Ctx.ResponseWriter.WriteHeader(403)
			c.Data["json"] = Response{403, 403, message, ""}
			c.ServeJSON()
			return
		}
		if models.CheckPassword(reqData.OldPassword) {
			if reqData.NewPassword == reqData.RetypePassword || err == nil {
				if user {

				}
				if err := models.UpdatePassword(reqData.OldPassword, reqData.NewPassword, token); err == nil {
					c.Data["json"] = successReturn
				} else {
					c.Data["json"] = err.Error()
				}
			} else {
				c.Data["json"] = &errNewPassDoNotMatch
			}
		} else {
			c.Data["json"] = &errOldPassDoNotMatch
		}

	} else {
		c.Data["json"] = &errNoUserOrPass
	}
	c.ServeJSON()
}

// @Title 认证测试
// @Description 测试错误码
// @Success 200 {object}
// @Failure 401 unauthorized
// @router /auth [get]
func (c *UserController) Auth() {
	et := utils.EasyToken{}
	token := strings.TrimSpace(c.Ctx.Request.Header.Get("Authorization"))
	validation, err := et.ValidateToken(token)
	if !validation {
		c.Ctx.ResponseWriter.WriteHeader(401)
		c.Data["json"] = Response{401, 401, fmt.Sprintf("%s", err), ""}
		c.ServeJSON()
		return
	}

	c.Data["json"] = Response{0, 0, "is login", ""}
	c.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = successReturn
	u.ServeJSON()
}
