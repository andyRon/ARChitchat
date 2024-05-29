package models

import "time"

// Post 主题模型类
type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

// CreatedAtDate 格式化创建时间
func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("2006-01-02 15:04:05") // TODO
}

// User 根据主题获取用户信息
func (post *Post) User() (user User) {
	user = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", post.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}
