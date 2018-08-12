<%@ WebHandler Language="VB" Class="BloggerAPI" %>
<%@ Assembly Name="CookComputing.XmlRpcV2" %>
Imports CookComputing.XmlRpc
Imports SqlConnHelper

<XmlRpcService()> _
Public Class BloggerAPI
    Inherits XmlRpcService
    Implements CookComputing.Blogger.IBlogger
    
       
    Public Function deletePost(ByVal appKey As String, ByVal postid As String, ByVal username As String, ByVal password As String, ByVal publish As Boolean) As Boolean Implements CookComputing.Blogger.IBlogger.deletePost

    End Function

    Public Function editPost(ByVal appKey As String, ByVal postid As String, ByVal username As String, ByVal password As String, ByVal content As String, ByVal publish As Boolean) As Object Implements CookComputing.Blogger.IBlogger.editPost

    End Function

    Public Function getCategories(ByVal blogid As String, ByVal username As String, ByVal password As String) As CookComputing.Blogger.Category() Implements CookComputing.Blogger.IBlogger.getCategories
        Dim BloggerId As Integer
        Dim sql As String
        Dim pa As New ParaHelper
        Dim i As Integer
        username = "Wuvist"
        BloggerId = Blogger.Auth(username, password)
        
        Dim dt As New DataTable
        sql = "select UserDefineCategory.* from UserDefineCategory WHERE Blogger=" & BloggerId
        dt = connHelper.retrieve(sql, CommandType.Text, Nothing)
        
        Dim cs(dt.Rows.Count - 1) As CookComputing.Blogger.Category
        i = 0
        For Each dr As DataRow In dt.Rows
            cs(i) = New CookComputing.Blogger.Category
            cs(i).categoryid = dr("index")
            cs(i).title = dr("cate")
            cs(i).htmlUrl = "http://www.blogwind.com/" & username & "/cate" & dr("index") & ".shtml"
            i += 1
        Next

        Return cs
    End Function

    Public Function getPost(ByVal appKey As String, ByVal postid As String, ByVal username As String, ByVal password As String) As CookComputing.Blogger.Post Implements CookComputing.Blogger.IBlogger.getPost

    End Function

    Public Function getRecentPosts(ByVal appKey As String, ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal numberOfPosts As Integer) As CookComputing.Blogger.Post() Implements CookComputing.Blogger.IBlogger.getRecentPosts

    End Function

    Public Function getTemplate(ByVal appKey As String, ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal templateType As String) As String Implements CookComputing.Blogger.IBlogger.getTemplate

    End Function

    Public Function getUserInfo(ByVal appKey As String, ByVal username As String, ByVal password As String) As CookComputing.Blogger.UserInfo Implements CookComputing.Blogger.IBlogger.getUserInfo

    End Function

    Public Function getUsersBlogs(ByVal appKey As String, ByVal username As String, ByVal password As String) As CookComputing.Blogger.BlogInfo() Implements CookComputing.Blogger.IBlogger.getUsersBlogs
        Dim b(0) As CookComputing.Blogger.BlogInfo
        Dim bb As New CookComputing.Blogger.BlogInfo
        bb.blogid = "1"
        bb.blogName = "Wuvist"
        bb.url = "http://www.blog.com/Wuvist"
        b(0) = bb
        Return b

    End Function

    Public Function newPost(ByVal appKey As String, ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal content As String, ByVal publish As Boolean) As String Implements CookComputing.Blogger.IBlogger.newPost
        Dim aid As Integer
        Dim BloggerId As Integer
        Dim pa As New ParaHelper
        BloggerId = Blogger.Auth(username, password)
        pa.addVarchar("@title", "12312", 50)
        pa.addVarchar("@blogger", username, 50)
        pa.addNText("@content", content)
        pa.addInt("@cate", Categories.GetBloggerCategoryIdByName(BloggerId, "乱弹"))

        'If mytags.Visible = True And mytags.Text <> "" Then
        '    Dim t As String
        '    t = mytags.Text
        '    mytags.Text = ""
        '    t = t.Replace("，", ",")
        '    t = t.Replace(";", ",")
        '    t = t.Replace("；", ",")
        '    t = t.Replace("'", "")
        '    t = t.Replace("""", "")
        '    t = t.Replace("|", "")
        '    '加多一个逗号，方便存储过程切分
        '    t = t & ","
        '    pa.addNVarchar("@tags", t, 50)
        'Else
        '    'cn.addPara("@tags", SqlDbType.NVarChar, 50, "")
        '    pa.addNVarchar("@tags", "", 50)
        'End If
        pa.addNVarchar("@tags", "", 50)
        aid = connHelper.exeInteger("insert_doc", CommandType.StoredProcedure, pa.parameters, False)
        Return aid
    End Function

    Public Function setTemplate(ByVal appKey As String, ByVal blogid As String, ByVal username As String, ByVal password As String, ByVal template As String, ByVal templateType As String) As Boolean Implements CookComputing.Blogger.IBlogger.setTemplate

    End Function
End Class