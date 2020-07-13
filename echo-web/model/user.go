package model

import (
	"encoding/json"
	"time"

	"letinvr.com/echo-web/module/auth"
	"letinvr.com/echo-web/module/log"
)

// func (u *User) TraceGetUserById(id uint64) *User {
// 	if s := u.Trace(); s != nil {
// 		defer s.Finish()
// 	}

// 	user := User{}
// 	var count int64
// 	db := DB().Where("id = ?", id)
// 	if err := Cache(db).First(&user).Count(&count).Error; err != nil {
// 		log.Debugf("GetUserById error: %v", err)
// 		return nil
// 	}

// 	return &user
// }

func (u *User) GetUserById(id uint64) *User {
	user := User{}
	var count int64
	db := DB().Where("id = ?", id)
	if err := Cache(db).First(&user).Count(&count).Error; err != nil {
		log.Debugf("GetUserById error: %v", err)
		return nil
	}

	return &user
}

func (u *User) GetUserList(m map[string]string, page uint64, pageSize uint64) ([]*User, uint64, error) {
	var users []*User
	//users := make([]*User, 0)
	//var users map[string]interface{}
	db := DB()
	//查询条件
	if m["gender"] > "0" {
		db = db.Where("gender = ?", m["gender"])
	}
	if m["nickname"] != "" {
		db = db.Where("nickname like ?", m["nickname"])
	}

	var count int
	db = db.Find(&users).Count(&count)
	//分页
	var totalpage uint64
	if page > 0 && pageSize > 0 {
		db = db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	//获取结果集
	db = Cache(db).Find(&users).Order("id asc")
	var err error
	if err := db.Error; err != nil {
		log.Debugf("GetUserList error: %v", err)
	}
	//totalpage = uint64(math.Ceil(float64(count) / float64(pageSize)))
	totalpage = uint64(count)
	return users, totalpage, err
}

//保存数据
func (u User) Save(user User) *User {
	//now := time.Now()
	//CreatedAt, _ := time.Parse("2006-01-02 15:04:05", now.Format("2006-01-02 15:04:05"))
	//CreatedAt, _ := time.Parse("2006-01-02 15:04:05", now.Format("2006-01-02 15:04:05"))
	//user["CreatedAt"] = CreatedAt

	if err := Cache(db).Create(&user).Error; err != nil {
		return nil
	}
	return &user
}

//保存数据
func (u User) Update(id int64, user User) *User {
	//now := time.Now()
	//CreatedAt, _ := time.Parse("2006-01-02 15:04:05", now.Format("2006-01-02 15:04:05"))
	//CreatedAt, _ := time.Parse("2006-01-02 15:04:05", now.Format("2006-01-02 15:04:05"))
	//user["CreatedAt"] = CreatedAt
	if err := Cache(db).Model(u).Where("id=?", id).Update(&user).Error; err != nil {
		return nil
	}
	return &user
}

//删除数据
func (u User) Delete(id int) error {
	if err := Cache(db).Model(u).Delete(&User{}, id).Error; err != nil {
		log.Debugf("Delete error: %v", err)
		return err
	}
	return nil
}

func (u *User) GetUserByNicknamePwd(nickname string, pwd string) *User {
	user := User{}
	if err := DB().Where("nickname = ? AND password = ?", nickname, pwd).First(&user).Error; err != nil {
		log.Debugf("GetUserByNicknamePwd error: %v", err)
		return nil
	}
	return &user
}

func (u *User) AddUserWithNicknamePwd(nickname string, pwd string) *User {
	user := User{Nickname: nickname, Password: pwd, Birthday: time.Now()}

	if err := DB().Create(&user).Error; err != nil {
		return nil
	}
	return &user
}

type Time time.Time

// func (t *Time) UnmarshalJSON(data []byte) (err error) {
// 	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
// 	*t = Time(now)
// 	return
// }

// func (t Time) MarshalJSON() ([]byte, error) {
// 	b := make([]byte, 0, len(timeFormart)+2)
// 	b = append(b, '"')
// 	b = time.Time(t).AppendFormat(b, timeFormart)
// 	b = append(b, '"')
// 	return b, nil
// }

const (
	timeFormart = "2006-01-02 15:04:05"
)

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}

type User struct {
	Id            uint64    `json:"id,omitempty"`
	Nickname      string    `form:"nickname" json:"nickname,omitempty"`
	Email         string    `json:"email,omitempty"`
	Password      string    `form:"password" json:"-"`
	Gender        int64     `json:"gender"`
	Birthday      time.Time `json:"birthday"`
	Face          string    `json:"face"`
	CreatedAt     time.Time `gorm:"column:created_time" json:"created_time,omitempty"`
	UpdatedAt     time.Time `gorm:"column:updated_time" json:"updated_time,omitempty"`
	authenticated bool      `form:"-" db:"-" json:"-"`
}
type User1 struct {
	User
	Age int
}

func (this User1) MarshalJSON() ([]byte, error) {
	// 定义一个该结构体的别名
	type AliasUser User1
	// 定义一个新的结构体
	tmpUser := struct {
		AliasUser
		Birthday  string `json:"birthday"`
		CreatedAt string `json:"created_time"`
		UpdatedAt string `json:"updated_time"`
	}{
		AliasUser: (AliasUser)(this),
		Birthday:  this.Birthday.Format(timeFormart),
		CreatedAt: this.CreatedAt.Format(timeFormart),
		UpdatedAt: this.UpdatedAt.Format(timeFormart),
	}
	return json.Marshal(tmpUser)
}

// GetAnonymousUser should generate an anonymous user model
// for all sessions. This should be an unauthenticated 0 value struct.
func GenerateAnonymousUser() auth.User {
	return &User{}
}

func (u User) TableName() string {
	return "user"
}

// Login will preform any actions that are required to make a user model
// officially authenticated.
func (u *User) Login() {
	// Update last login time
	// Add to logged-in user's list
	// etc ...
	u.authenticated = true
}

// Logout will preform any actions that are required to completely
// logout a user.
func (u *User) Logout() {
	// Remove from logged-in user's list
	// etc ...
	u.authenticated = false
}

func (u *User) IsAuthenticated() bool {
	return u.authenticated
}

func (u *User) UniqueId() interface{} {
	return u.Id
}

// GetById will populate a user object from a database model with
// a matching id.
func (u *User) GetById(id interface{}) error {
	if err := DB().Where("id = ?", id).First(&u).Error; err != nil {
		return err
	}
	return nil
}
