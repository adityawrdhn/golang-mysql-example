package controllers

import (
	//"errors"
	//"regexp"
	//"strings"

	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/validation"
	//"github.com/dgrijalva/jwt-go"
)

// Predefined const error strings.
const (
	ErrInputData    = "Error Input Data"
	ErrDatabase     = "Database Error"
	ErrDupUser      = "Duplicate User"
	ErrNoUser       = "No User "
	ErrPass         = "Password Error"
	ErrNoUserPass   = "No User Password"
	ErrNoUserChange = "no user change"
	ErrInvalidUser  = "invalid user"
	ErrOpenFile     = "open file error"
	ErrWriteFile    = "write file error"
	ErrSystem       = "system error"
)

// UserData definition.
type UserSuccessLoginData struct {
	AccessToken string `json:"access_token"`
	UserName    string `json:"username"`
}

// CreateDevice definition.
type CreateObjectData struct {
	Id int `json:"id"`
}

// GetDevices definition.
type GetDeviceData struct {
	TotalCount int64       `json:"total_count"`
	Devices    interface{} `json:"devices"`
}

// GetAirAds definition.
type GetAirAdData struct {
	TotalCount int64       `json:"total_count"`
	AirAds     interface{} `json:"airads"`
}

// Predefined controller error/success values.
var (
	successReturn   = &Response{200, 0, "Success", "Success"}
	err404          = &Response{404, 404, "File Not Found", "File Not Found"}
	errInputData    = &Response{400, 10001, "Error Input Data", "Error input Data"}
	errDatabase     = &Response{500, 10002, "Database Error", "Database Error"}
	errUserToken    = &Response{500, 10002, "User Token error", "User Token Error"}
	errDupUser      = &Response{400, 10003, "Duplicate User", "Duplicate User"}
	errNoUser       = &Response{400, 10004, "No User Error", "No User Error"}
	errPass         = &Response{400, 10005, "Password Error", "Password Error"}
	errNoUserOrPass = &Response{400, 10006, "No user or Password Error", "No user or Password Error"}
	errNoUserChange = &Response{400, 10007, "No user change", "No user change"}
	errInvalidUser  = &Response{400, 10008, "Invalid User", "Invalid User"}
	errOpenFile     = &Response{500, 10009, "Open File Error", "Open File Error"}
	errWriteFile    = &Response{500, 10010, "Write File Error", "Write File Error"}
	errSystem       = &Response{500, 10011, "System Error", "System Error"}
	errExpired      = &Response{400, 10012, "Expired Session", "Expired Session"}
	errPermission   = &Response{400, 10013, "Permission Error", "Permission Error"}
)

// BaseController definition.
//type BaseController struct {
//	beego.Controller
//}

// RetError return error information in JSON.
func (base *BaseController) RetError(e *Response) {
	if mode := beego.AppConfig.String("runmode"); mode == "prod" {
		e.Data = ""
	}

	base.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
	base.Ctx.ResponseWriter.WriteHeader(e.Status)
	base.Data["json"] = e
	base.ServeJSON()
	base.StopRun()
}

var sqlOp = map[string]string{
	"eq": "=",
	"ne": "<>",
	"gt": ">",
	"ge": ">=",
	"lt": "<",
	"le": "<=",
}

// ParseQueryParm parse query parameters.
//   query=col1:op1:val1,col2:op2:val2,...
//   op: one of eq, ne, gt, ge, lt, le
//func (base *BaseController) ParseQueryParameter() (v map[string]string, o map[string]string, err error) {
//	var nameRule = regexp.MustCompile("^[a-zA-Z0-9_]+$")
//	queryVal := make(map[string]string)
//	queryOp := make(map[string]string)
//
//	query := base.GetString("query")
//	if query == "" {
//		return queryVal, queryOp, nil
//	}
//
//	for _, cond := range strings.Split(query, ",") {
//		kov := strings.Split(cond, ":")
//		if len(kov) != 3 {
//			return queryVal, queryOp, errors.New("Query format != k:o:v")
//		}
//
//		var key string
//		var value string
//		var operator string
//		if !nameRule.MatchString(kov[0]) {
//			return queryVal, queryOp, errors.New("Query key format is wrong")
//		}
//		key = kov[0]
//		if op, ok := sqlOp[kov[1]]; ok {
//			operator = op
//		} else {
//			return queryVal, queryOp, errors.New("Query operator is wrong")
//		}
//		value = strings.Replace(kov[2], "'", "\\'", -1)
//
//		queryVal[key] = value
//		queryOp[key] = operator
//	}
//
//	return queryVal, queryOp, nil
//}

// ParseOrderParameter parse order parameters.
//   order=col1:asc|desc,col2:asc|esc,...
//func (base *BaseController) ParseOrderParameter() (o map[string]string, err error) {
//	var nameRule = regexp.MustCompile("^[a-zA-Z0-9_]+$")
//	order := make(map[string]string)
//
//	v := base.GetString("order")
//	if v == "" {
//		return order, nil
//	}
//
//	for _, cond := range strings.Split(v, ",") {
//		kv := strings.Split(cond, ":")
//		if len(kv) != 2 {
//			return order, errors.New("Order format != k:v")
//		}
//		if !nameRule.MatchString(kv[0]) {
//			return order, errors.New("Order key format is wrong")
//		}
//		if kv[1] != "asc" && kv[1] != "desc" {
//			return order, errors.New("Order val isn't asc/desc")
//		}
//
//		order[kv[0]] = kv[1]
//	}
//
//	return order, nil
//}

// ParseLimitParameter parse limit parameter.
//   limit=n
//func (base *BaseController) ParseLimitParameter() (l int64, err error) {
//	if v, err := base.GetInt64("limit"); err != nil {
//		return 10, err
//	} else if v > 0 {
//		return v, nil
//	} else {
//		return 10, nil
//	}
//}

// ParseOffsetParameter parse offset parameter.
//   offset=n
//func (base *BaseController) ParseOffsetParameter() (o int64, err error) {
//	if v, err := base.GetInt64("offset"); err != nil {
//		return 0, err
//	} else if v > 0 {
//		return v, nil
//	} else {
//		return 0, nil
//	}
//}

// VerifyForm use validation to verify input parameters.
//func (base *BaseController) VerifyForm(obj interface{}) (err error) {
//	valid := validation.Validation{}
//	ok, err := valid.Valid(obj)
//	if err != nil {
//		return err
//	}
//	if !ok {
//		str := ""
//		for _, err := range valid.Errors {
//			str += err.Key + ":" + err.Message + ";"
//		}
//		return errors.New(str)
//	}
//
//	return nil
//}

// ParseToken parse JWT token in http header.
//func (base *BaseController) ParseToken() (t *jwt.Token, e *ControllerError) {
//	authString := base.Ctx.Input.Header("Authorization")
//	beego.Debug("AuthString:", authString)
//
//	kv := strings.Split(authString, " ")
//	if len(kv) != 2 || kv[0] != "Bearer" {
//		beego.Error("AuthString invalid:", authString)
//		return nil, errInputData
//	}
//	tokenString := kv[1]
//
//	// Parse token
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		return []byte("secret"), nil
//	})
//	if err != nil {
//		beego.Error("Parse token:", err)
//		if ve, ok := err.(*jwt.ValidationError); ok {
//			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
//				// That's not even a token
//				return nil, errInputData
//			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
//				// Token is either expired or not active yet
//				return nil, errExpired
//			} else {
//				// Couldn't handle this token
//				return nil, errInputData
//			}
//		} else {
//			// Couldn't handle this token
//			return nil, errInputData
//		}
//	}
//	if !token.Valid {
//		beego.Error("Token invalid:", tokenString)
//		return nil, errInputData
//	}
//	beego.Debug("Token:", token)
//
//	return token, nil
//}
