package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"letinvr.com/echo-web/model"
)

func UserHandler(c *Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		panic(err)
	}

	var User model.User
	u := User.GetUserById(id)

	c.Set("tmpl", "web/user/add")
	c.Set("data", map[string]interface{}{
		"title": "User",
		"user":  u,
	})

	return nil
}

func UDeleteHandler(c *Context) error {
	//接收参数
	idStr := c.FormValue("ids")
	idArr := strings.Split(strings.Trim(idStr, "[]"), ",")
	//开启事物
	tx := model.DB().Begin()
	var user model.User
	for i := 0; i < len(idArr); i++ {
		id, _ := strconv.Atoi(idArr[i])
		err := user.Delete(id)
		if err != nil {
			tx.Callback()
			return c.outJson(400, "删除失败", nil)
		}
	}
	tx.Commit()
	return c.outJson(200, "删除成功", nil)
}

func UEditHandler(c *Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		panic(err)
	}
	var User model.User
	userInfo := User.GetUserById(id)
	//Birthday := userInfo.CreatedAt.Format("2006-01-02 15:04:05")
	c.Set("tmpl", "web/edit")
	c.Set("data", map[string]interface{}{
		"title":    "编辑用户",
		"userInfo": userInfo,
	})
	// var c1 *gin.Context
	// c1.HTML(http.StatusOK, "web/edit.html", pongo2.Context{"title": "编辑用户", "userInfo": userInfo})
	return nil
}

type Susess struct {
	Code int
	Msg  string
}

func UUpdateHandler(c *Context) error {
	//获取参数
	idStr := c.FormValue("id")          //用户ID
	nickname := c.FormValue("nickname") //用户名
	password := c.FormValue("password") //密码
	sex := c.FormValue("sex")           //性别
	birthday := c.FormValue("birthday") //出生日期
	email := c.FormValue("email")       //邮箱
	face := c.FormValue("cover")        //邮箱

	//Id处理，转换64位
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		panic(err)
	}

	//性别处理，转换64位
	Gender, err := strconv.ParseInt(sex, 10, 64)
	if err != nil {
		panic(err)
	}
	//User结构体
	now := time.Now()
	//CreatedAt, _ := time.Parse("2006-01-02 15:04:05", now.Format("2006-01-02 15:04:05"))
	Birthday, _ := time.Parse("2006-01-02", birthday)
	//CreatedAt, _ = t.Parse("2006-01-02 15:04:05", now.Format("2006-01-02 15:04:05"))
	userData := model.User{Nickname: nickname, Email: email, Face: face, Gender: Gender, Password: password, Birthday: Birthday, CreatedAt: now, UpdatedAt: now}
	//保存数据
	var User model.User
	if User.Update(id, userData) != nil {
		// su := &Susess{
		// 	Code: 0,
		// 	Msg:  "成功",
		// }
		// c.JSON(http.StatusOK, su)
		return c.outJson(200, "修改成功", nil)
	}
	return c.outJson(400, "修改失败，请稍后再试！", nil)
}

func UAddHandler(c *Context) error {
	c.Set("tmpl", "web/add")
	c.Set("data", map[string]interface{}{
		"title": "添加用户",
	})
	return nil
}

func USaveHandler(c *Context) error {
	//params := c.FormParams() //用户名
	//获取参数
	nickname := c.FormValue("nickname") //用户名
	password := c.FormValue("password") //密码
	sex := c.FormValue("sex")           //性别
	birthday := c.FormValue("birthday") //出生日期
	email := c.FormValue("email")       //邮箱
	face := c.FormValue("cover")        //邮箱

	//性别处理，转换64位
	Gender, err := strconv.ParseInt(sex, 10, 64)
	if err != nil {
		panic(err)
	}
	//User结构体
	now := time.Now()
	//CreatedAt, _ := time.Parse("2006-01-02 15:04:05", now.Format("2006-01-02 15:04:05"))
	Birthday, _ := time.Parse("2006-01-02", birthday)
	//CreatedAt, _ = t.Parse("2006-01-02 15:04:05", now.Format("2006-01-02 15:04:05"))
	userData := model.User{Nickname: nickname, Email: email, Face: face, Gender: Gender, Password: password, Birthday: Birthday, CreatedAt: now, UpdatedAt: now}
	//保存数据
	var User model.User
	if User.Save(userData) != nil {
		return c.outJson(200, "添加成功", nil)
	}
	return c.outJson(400, "添加失败，请稍后再试！", nil)
}

func UListHandler(c *Context) error {
	//接受参数
	nickname := c.FormValue("keywords") //昵称
	gender := c.FormValue("sex")        //性别
	pageStr := c.FormValue("page")      //分页
	pageSizeStr := c.FormValue("limit") //分页参数
	//分页数处理，转换64位
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		panic(err)
	}
	pageSize, err := strconv.ParseUint(pageSizeStr, 10, 64)
	if err != nil {
		panic(err)
	}
	//查询条件map
	condition := map[string]string{"nickname": nickname, "gender": gender}
	var User model.User
	//获取用户列表
	list, count, err := User.GetUserList(condition, page, pageSize)
	listA, _ := json.Marshal(list)
	fmt.Println(string(listA))
	now := time.Now()
	userLists := map[int]interface{}{}
	for i, v := range list {
		//value := reflect.ValueOf(&v.CreatedAt) // 参数必须为指针地址
		//CreatedAt := v.CreatedAt.Format("2006-01-02 15:04:05")
		//value.Elem().Set(CreatedAt)
		//fmt.Println(value)
		year, _ := strconv.Atoi(v.Birthday.Format("2006"))
		age := now.Year() - year
		userLists[i] = model.User1{
			User: *v,
			Age:  age,
		}
	}
	//fmt.Printf("%+v\n", userLists)
	//fmt.Println("")
	jsonData := map[string]interface{}{"code": 0, "count": count, "data": userLists}
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(c.Response()).Encode(jsonData)
}

func (c *Context) outJson(code int, msg string, data map[string]interface{}) error {
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(code)

	jsonData := map[string]interface{}{"code": code, "msg": msg}
	if data != nil {
		jsonData["data"] = data
	}
	return json.NewEncoder(c.Response()).Encode(jsonData)
}
