在线论坛
----

参考：https://laravelacademy.org/post/21519

ARChitchat，基于Go语言构建在线论坛。

## 1 整体设计与数据模型

![](images/image-15850625575141.jpg)

这个在线论坛需要具备用户认证功能（注册、登录、退出等），认证后的用户才能创建新的群组、以及在群组中发表主题，访客用户访问论坛首页可以查看群组列表，进入指定群组页面可以查看对应的主题信息。



### 技术方案

![](images/image-20240519003616881.png)

Client 代表客户端发起请求的用户，虚框内是部署在服务器已启动的在线论坛应用，Multiplexer 代表路由器（比如 [gorilla/mux](https://github.com/gorilla/mux) ），Handler 代码处理器/处理器方法，数据库操作位于处理器方法中，Templates 代表最终展示给用户的经过模板引擎编译过的视图模板。

Go Web项目因为静态语言和实现机制的缘故不需要借助类似 php-fpm、nginx 这种额外的HTTP 服务器、反向代理服务器，Go 应用**以单文件形式部**署，静态资源和视图模板的部署与传统动态语言不一样等。

### 数据模型

- 用户（User）
- 群组（Thread）
- 主题（Post）



在本项目开发时，会把用户会话（Session）也存储到数据库，所以需要一个额外的会话模型，此外，为了简化应用，我们不会真的像Google网上论坛那样对用户做权限管理，整个应用只包含一种用户类型，并且具备所有操作权限：

![](images/image-DraggedImage-1 00.36.28.png)



## 2 通过模型类与MySQL数据库交互

### 项目初始化

目录/文件的作用介：

- `main.go`：应用入口文件
- `config.json`：全局配置文件
- `handlers`：用于存放处理器代码（可类比为 MVC 模式中的控制器目录）
- `logs`：用于存放日志文件
- `models`：用于存放与数据库交互的模型类
- `public`：用于存放前端资源文件，比如图片、CSS、JavaScript 等
- `routes`：用于存放路由文件和路由器实现代码
- `views`：用于存放视图模板文件





```sh
go mod init github.com/andyron/architchat
```



### 创建数据表

```sql
create table users (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  name       varchar(255),
  email      varchar(255) not null unique,
  password   varchar(255) not null,
  created_at timestamp not null
);
    
create table sessions (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      varchar(255),
  user_id    integer references users(id),
  created_at timestamp not null
);
    
create table threads (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  topic      text,
  user_id    integer references users(id),
  created_at timestamp not null
);
    
create table posts (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  body       text,
  user_id    integer references users(id),
  thread_id  integer references threads(id),
  created_at timestamp not null
);
```



### 与数据库交互

#### 数据库驱动

第三方包 [go-mysql-driver](https://github.com/go-sql-driver/mysql) 

```sh
go get github.com/go-sql-driver/mysql
```

#### 数据库连接



#### 用户相关模型类



#### 主题相关模型类





## 3 访问论坛首页

### 3.1 定义路由器

基于 [gorilla/mux](https://github.com/gorilla/mux) 来实现路由器：

```sh
go get github.com/gorilla/mux
```

仿照Laravel框架对Go路由处理器代码进行拆分



### 3.2 启动HTTP服务器



> ```go
> 	. "github.com/andyron/architchat/routes"
> 
> 	_ "github.com/go-sql-driver/mysql"
> ```
>
>  `.` 别名，通过这种方式引入的包可以直接调用包中对外可见的变量、方法和结构体，而不需要加上包名前缀。
>
>  `_` 别名，只会调用该包里定义的 `init` 方法。🔖



### 3.3 处理静态资源



### 3.4 编写处理器实现



#### 首页处理器方法

#### 创建视图模板

使用 Go 自带的 `html/template` 作为模板引擎

主布局文件 `layout.html`

顶部导航模板 `navbar.html`

首页视图模板 `index.html`



引入多个视图模板是为了提高模板代码的复用性，因为对于同一个应用的不同页面来说，可能基本布局、页面顶部导航和页面底部组件都是一样的。

#### 渲染视图模板



#### 注册首页路由

### 访问论坛首页



## 4 通过 Cookie + Session 实现用户认证

### 编写全局辅助函数



## 5 创建群组和主题功能实现



## 6 日志和错误处理



## 7 通过单例模式获取全局配置



## 8 消息、视图和日期时间本地化



```sh
go get -u github.com/nicksnyder/go-i18n/v2/i18n
go get -u github.com/nicksnyder/go-i18n/v2/goi18n
```

### 消息本地化



### 视图本地化



### 日期时间本地化





## 9 部署 Go Web 应用



## 10 通过 Viper 读取配置文件并实现热加载





