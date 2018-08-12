<%@ WebHandler Language="VB" Class="BSPValidate" %>

Imports System
Imports System.Web

Public Class BSPValidate
    Inherits Blogwind.Blog.BWAuthHandler
    Private Shared ReadOnly Log As log4net.ILog = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType)
    
    Private Sub SaveUser(ByVal account_type As Blogwind.WindMill.Account.AccountTypes, ByVal ServerHost As String, ByVal username As String, ByVal password As String, Optional ByVal blogid As String = Nothing)
        Dim u As New Blogwind.WindMill.Account
        u.AccountType = account_type
        u.UserName = username
        u.Pwd = Cryptor.Encrypt(password)
        u.Blogid = blogid
        u.UserId = Me.ph.bloggerID
        u.CreateOn = Now
        u.ServerHost = ServerHost
        u.Save()
    End Sub
    
    Private Sub Validate(ByVal context As System.Web.HttpContext)
        Dim account_type As Blogwind.WindMill.Account.AccountTypes
        Dim username As String = context.Request("username")
        Dim password As String = context.Request("password")
        Dim blogs As CookComputing.Blogger.BlogInfo()
        account_type = context.Request("account_type")
        
        Select Case account_type
            Case Blogwind.WindMill.Account.AccountTypes.Baidu
                If baidu.Validate(username, password) Then
                    Log.Debug("baidu ok")
                    SaveUser(account_type, "hi.baidu.com", username, password)
                    context.Response.Write(0)
                Else
                    Log.Debug("baidu fail")
                    context.Response.Write(1)
                End If
            Case Blogwind.WindMill.Account.AccountTypes.Blogger, Blogwind.WindMill.Account.AccountTypes.MetaWebLog, _
                 Blogwind.WindMill.Account.AccountTypes.WordPress
                Dim ServerHost As String
                ServerHost = context.Request("server_host")
                blogs = Metaweblog.Validate(ServerHost, username, password)
                If blogs IsNot Nothing Then
                    Log.Debug("metaweblog ok")
                    
                    SaveUser(account_type, ServerHost, username, password, blogs(0).blogid)
                    context.Response.Write(0)
                Else
                    Log.Debug("metaweblog fail")
                    context.Response.Write(1)
                End If
            Case Blogwind.WindMill.Account.AccountTypes.MSNSpace
                blogs = Metaweblog.Validate(Metaweblog.Urls.MSNSpace, username, password)
                If blogs IsNot Nothing Then
                    Log.Debug("msn ok")
                    SaveUser(account_type, Metaweblog.Urls.MSNSpace, username, password, blogs(0).blogid)
                    context.Response.Write(0)
                Else
                    Log.Debug("msn fail")
                    context.Response.Write(1)
                End If
            Case Blogwind.WindMill.Account.AccountTypes.Twitter
                If Twitter.Validate(username, password, Twitter.Urls.TwitterAuth) Then
                    SaveUser(account_type, Twitter.Urls.Twitter, username, password)
                    context.Response.Write(0)
                Else
                    context.Response.Write(1)
                End If
            Case Blogwind.WindMill.Account.AccountTypes.Fanfou
                If Twitter.Validate(username, password, Twitter.Urls.FanfouAuth) Then
                    SaveUser(account_type, Twitter.Urls.Fanfou, username, password)
                    context.Response.Write(0)
                Else
                    context.Response.Write(1)
                End If
        End Select
    End Sub
    
    Private Sub Delete(ByVal context As System.Web.HttpContext)
        Dim id As Integer = context.Request("id")
        Dim u As Blogwind.WindMill.Account
        u = Blogwind.WindMill.Account.FetchByID(id)
        If (u Is Nothing) Then
            context.Response.Write(1)
        End If
        
        If u.UserId <> Me.ph.bloggerID Then
            context.Response.Write(2)
        End If
        
        Blogwind.WindMill.Account.Destroy(id)
        context.Response.Write(0)
    End Sub
    
    Public Overrides Sub ProcessAuthRequest(ByVal context As System.Web.HttpContext)
        Dim action As String
        action = context.Request("action")
        If action = "validate" Then
            Validate(context)
        ElseIf action = "delete" Then
            Delete(context)
        End If

    End Sub
End Class