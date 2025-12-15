package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func showLoginnPage(c *gin.Context){
    c.HTML(http.StatusOK,"login.html", nil)
}

func handleLogin(c *gin.Context){
    username := c.PostForm("username")
    password := c.PostForm("password")
    
    if username == "admin" && password == "1234"{
        session := sessions.Default(c)
        session.Set("user", username)
        session.Save()
        
        c.Redirect(http.StatusSeeOther,"/home")
        return
    }
    c.HTML(http.StatusUnauthorized, "login.html",gin.H{
        "error":"invalid username or password",
    })
}

func authRequired() gin.HandlerFunc{
    return func(c *gin.Context) {
        session := sessions.Default(c)
        user := session.Get("user")
        if user == nil{
            c.Redirect(http.StatusSeeOther, "/login")
            c.Abort()
            return
        }
        c.Next()
    }
}

func homePage (c *gin.Context){
    session := sessions.Default(c)
    user := session.Get("user")
    c.HTML(http.StatusOK, "login.html",gin.H{
        "user":user,
    })
}

func logout(c *gin.Context){
    session := sessions.Default(c)
    session.Clear()
    session.Save()
    c.Redirect(http.StatusSeeOther, "/login")
}

func main() {
	r := gin.Default()
    
    r.LoadHTMLGlob("practiceHtml/*")
    
    store := cookie.NewStore([]byte("secret123"))
    r.Use(sessions.Sessions("mysession", store))
    
    
}