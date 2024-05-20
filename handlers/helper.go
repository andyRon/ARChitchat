package handlers

import (
	"errors"
	"fmt"
	. "github.com/andyron/architchat/config"
	"github.com/andyron/architchat/models"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

// 全局辅助函数（主要用在处理器中）

var logger *log.Logger
var config *Configuration
var localizer *i18n.Localizer

// 初始化日志处理器
func init() {
	// 获取全局配置实例
	config = LoadConfig()
	// 获取本地化实例
	localizer = i18n.NewLocalizer(config.LocaleBundle, config.App.Language)

	file, err := os.OpenFile("logs/architchat.log", os.O_CREATE|os.O_APPEND, 0666) // TODO
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

/* 日志函数 */
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

// 命名为danger是为了避免与error重复
func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}
func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

/***************/

// 异常处理统一重定向到错误页面
func errorMessage(w http.ResponseWriter, r *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(w, r, strings.Join(url, ""), 302)
}

// 通过Cookie判断用户是否已登录
func session(w http.ResponseWriter, r *http.Request) (sess models.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// 解析HTML模板（应对需要传入多个模板文件的情况，避免重复编写模板代码）
//func parseTemplateFiles(filenames ...string) (t *template.Template) {
//	var files []string
//	t = template.New("layout")
//	for _, file := range filenames {
//		files = append(files, fmt.Sprintf("views/%s.html", file))
//	}
//	t = template.Must(t.ParseFiles(files...))
//	return
//}

// 生成HTML
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s/%s.html", config.App.Language, file))
	}

	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("layout").Funcs(funcMap)

	templates := template.Must(t.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data) // TODO
}

func Version() string {
	return "0.1"
}

// 日期格式化辅助函数
func formatDate(t time.Time) string {
	datetime := "2006-01-02 15:04:05"
	return t.Format(datetime)
}
