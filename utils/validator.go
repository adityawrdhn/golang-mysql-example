package utils

import (
	"fmt"
	"github.com/astaxie/beego/validation"
)

func CheckUsernamePassword(username string, password string) (errorMessage string) {
	valid := validation.Validation{}
	valid.Required(username, "Username").Message("Username Required")
	valid.Required(password, "Password").Message("Password Required")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			//c.Ctx.ResponseWriter.WriteHeader(403)
			//c.Data["json"] = Response{403001, 403001,err.Message, ""}
			//c.ServeJSON()
			return fmt.Sprintf("%s", err.Message)
		}
	}
	return fmt.Sprintf("%s", "ok")
}

func CheckNewUserPost(Name string, Username string, Password string, Email string,
	Location string, Phone string, Gender string) (errorMessage string) {
	valid := validation.Validation{}

	valid.Required(Name, "Name").Message("Name is required")
	// valid.Alpha(Name, "name").Message("Name : Invalid Characters")
	valid.Required(Username, "Username").Message("Username is required")
	valid.AlphaNumeric(Username, "Username").Message("Username must be alphanumeric")
	valid.Required(Password, "Password").Message("Password is required")
	valid.MinSize(Password, 8, "Password").Message("Password at least 8-12 characters")
	valid.MaxSize(Password, 12, "Password").Message("Password at least 8-12 characters")
	valid.Required(Email, "Email").Message("Email is required")
	valid.Email(Email, "Email").Message("Email is not valid")
	valid.Required(Location, "Location").Message("Location is required")
	valid.Numeric(Phone, "Phone").Message("Phone number is not valid")
	valid.Required(Phone, "Phone").Message("Phone number is required")
	valid.MinSize(Phone, 6, "Phone").Message("Phone number is not valid -min")
	valid.MaxSize(Phone, 12, "Phone").Message("Phone number is not valid -max")
	valid.Required(Gender, "Gender").Message("Gender is required")
	// valid.Range(Gender, 0, 1, "Gender").Message("Gender type is not valid")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			//c.Ctx.ResponseWriter.WriteHeader(403)
			//c.Data["json"] = Response{403001, 403001,err.Message, ""}
			//c.ServeJSON()
			return fmt.Sprintf("%s", err.Message)
		}
	}
	return fmt.Sprintf("%s", "ok")
}
func CheckUpdateUserPost(Name string, Username string, Email string,
	Location string, Phone string, Gender string) (errorMessage string) {
	valid := validation.Validation{}

	valid.Required(Name, "Name").Message("Name is required")
	// valid.Alpha(Name, "name").Message("Name : Invalid Characters")
	valid.Required(Username, "Username").Message("Username is required")
	valid.AlphaNumeric(Username, "Username").Message("Username must be alphanumeric")
	valid.Required(Email, "Email").Message("Email is required")
	valid.Email(Email, "Email").Message("Email is not valid")
	valid.Required(Location, "Location").Message("Location is required")
	valid.Numeric(Phone, "Phone").Message("Phone number is not valid")
	valid.Required(Phone, "Phone").Message("Phone number is required")
	valid.MinSize(Phone, 6, "Phone").Message("Phone number is not valid -min")
	valid.MaxSize(Phone, 12, "Phone").Message("Phone number is not valid -max")
	valid.Required(Gender, "Gender").Message("Gender is required")
	// valid.Range(Gender, 0, 1, "Gender").Message("Gender type is not valid")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			//c.Ctx.ResponseWriter.WriteHeader(403)
			//c.Data["json"] = Response{403001, 403001,err.Message, ""}
			//c.ServeJSON()
			return fmt.Sprintf("%s", err.Message)
		}
	}
	return fmt.Sprintf("%s", "ok")
}
func CheckPassword(OldPassword string, NewPassword string, RetypePassword string) (errorMessage string) {
	valid := validation.Validation{}

	valid.Required(OldPassword, "Password").Message("Old Password is required")

	valid.Required(NewPassword, "Password").Message("New Password is required")
	valid.MinSize(NewPassword, 8, "Password").Message("New Password at least 8-12 characters")
	valid.MaxSize(NewPassword, 12, "Password").Message("New Password at least 8-12 characters")
	// valid.Range(Gender, 0, 1, "Gender").Message("Gender type is not valid")
	valid.Required(RetypePassword, "Password").Message("Password is required")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			//c.Ctx.ResponseWriter.WriteHeader(403)
			//c.Data["json"] = Response{403001, 403001,err.Message, ""}
			//c.ServeJSON()
			return fmt.Sprintf("%s", err.Message)
		}
	}
	return fmt.Sprintf("%s", "ok")
}

func CheckUserDevicePost(userId int, limit int, offset int) (errorMessage string) {
	valid := validation.Validation{}
	valid.Required(userId, "UserId").Message("User ID is required")
	valid.Min(userId, 1, "UserId").Message("User ID number is not valid")
	//valid.Required(limit, "Limit").Message("Limit必填")
	valid.Range(limit, 0, 20, "Limit").Message("Limit is 20")
	//valid.Required(offset, "Offset").Message("Offset必填")
	valid.Range(offset, 0, 20, "Offset").Message("Offset is 20")
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			//c.Ctx.ResponseWriter.WriteHeader(403)
			//c.Data["json"] = Response{403001, 403001,err.Message, ""}
			//c.ServeJSON()
			return fmt.Sprintf("%s", err.Message)
		}
	}
	return fmt.Sprintf("%s", "ok")

}

func CheckNewAirAdPost(deviceId int, co string, humidity string, temperature string,
	pm25 string, pm10 string, nh3 string, o3 string, suggest string, aqiQuality string) (errorMessage string) {
	valid := validation.Validation{}
	//表单验证
	valid.Required(deviceId, "DeviceId").Message("")
	valid.Required(co, "Co").Message("设备名必填")
	valid.Required(humidity, "Humidity").Message("地址必填")
	valid.Required(temperature, "DeviceId").Message("用户ID必填")
	valid.Required(pm25, "Co").Message("设备名必填")
	valid.Required(pm10, "Humidity").Message("地址必填")
	valid.Required(o3, "DeviceId").Message("用户ID必填")
	valid.Required(suggest, "Co").Message("设备名必填")
	valid.Required(aqiQuality, "Humidity").Message("地址必填")
	valid.Required(nh3, "DeviceId").Message("用户ID必填")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		for _, err := range valid.Errors {
			//c.Ctx.ResponseWriter.WriteHeader(403)
			//c.Data["json"] = Response{403001, 403001,err.Message, ""}
			//c.ServeJSON()
			return fmt.Sprintf("%s", err.Message)
		}
	}
	return fmt.Sprintf("%s", "ok")
}
