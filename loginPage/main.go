package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func LoginLogoutMiddleware() gin.HandlerFunc{
    return func(c *gin.Context) {
        method := c.Request.Method
        path := c.Request.URL.Path
        ip := c.ClientIP()
        println("LOG:",method,path,"requested Ip:",ip)
        c.Next()
    }
}

func main() {
    router := gin.Default()
    
    router.LoadHTMLGlob("templates/*")
    
    store := cookie.NewStore([]byte("secret123"))
    router.Use(sessions.Sessions("mysession",store))
    
    
    router.GET("/login",func(c *gin.Context) {
        c.HTML(http.StatusOK, "login.html",nil)
    })
    
    router.POST("/login",LoginLogoutMiddleware(),func(c *gin.Context) {
        username := c.PostForm("username")
        password := c.PostForm("password")
        
        if username == "admin" && password == "1234"{
            session := sessions.Default(c)
            session.Set("user",username)
            session.Save()
            
            c.Redirect(http.StatusSeeOther,"/home")
            return
        }
        c.String(http.StatusUnauthorized,"invalid username or password")
    })
    
    router.GET("/home", func(c *gin.Context) {
        session := sessions.Default(c)
        user := session.Get("user")
        
        if user == nil{
            c.Redirect(http.StatusSeeOther,"/login")
        }
        
        c.HTML(http.StatusOK, "home.html", gin.H{
            "user":user,
        })
    })
    
    router.GET("/logout", func(c *gin.Context) {
        session := sessions.Default(c)
        session.Clear()
        session.Save()
        
        c.Redirect(http.StatusSeeOther,"/login")
    })
    
    router.Run(":8080")
}