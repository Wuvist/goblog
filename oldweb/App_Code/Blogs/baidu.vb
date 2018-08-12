Imports Microsoft.VisualBasic

Public Class baidu
    Private Shared ReadOnly Log As log4net.ILog = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType)

    Private Shared Function GetBSUSS(ByVal Username As String, ByVal Password As String) As String
        Dim c As New System.Net.Sockets.TcpClient
        c.Connect("passport.baidu.com", 80)
        c.ReceiveTimeout = 1000
        Dim ns As System.Net.Sockets.NetworkStream
        ns = c.GetStream
        Dim post As String
        post = "tpl_ok=&next_target=&tpl=&skip_ok=&aid=&need_pay=&need_coin=&pay_method=&u=.%2F&return_method=get&more_param=&return_type=&psp_tt=0&username=" & Username & "&password=" & Password & "&mem_pass=on"
        post = "Content-Length: " & post.Length.ToString & vbLf & vbLf & post & vbLf & vbLf

        Dim data As Byte()
        post = System.IO.File.ReadAllText(System.Web.HttpContext.Current.Server.MapPath("/App_Data/baidu_login.txt")) & post

        data = System.Text.Encoding.ASCII.GetBytes(post)
        ns.Write(data, 0, data.Length)

        Dim sr As New System.IO.StreamReader(ns)
        Dim sb As New StringBuilder
        Dim bduss As String
        Dim l As String
        l = sr.ReadLine
        While l.Length > 2
            If l.StartsWith("Set-Cookie: BDUSS=") Then
                Dim i As Integer
                i = l.IndexOf(";")
                bduss = l.Substring(18, i - 18)
            End If
            sb.Append(l)
            l = sr.ReadLine & vbLf
        End While
        ns.Close()

        Return bduss
    End Function

    Public Shared Function Validate(ByVal Username As String, ByVal Password As String) As Boolean
        Dim bduss As String = GetBSUSS(Username, Password)
        Return Not String.IsNullOrEmpty(bduss)
    End Function

    Public Shared Sub Post(ByVal Acc As Blogwind.WindMill.Account, ByVal Title As String, ByVal Content As String, ByVal Category As String)
        Dim c As New Metaweblog
        Dim blog As New CookComputing.MetaWeblog.Post
        blog.title = Title
        blog.description = Content
        blog.categories = New String() {Category}
        blog.dateCreated = Now

        Post(Acc.UserName, Cryptor.Decrypt(Acc.Pwd), blog)
    End Sub

    Public Shared Sub Post(ByVal username As String, ByVal password As String, ByVal newpost As CookComputing.MetaWeblog.Post)
        Dim bduss As String = GetBSUSS(username, password)
        Dim s, post As String
        Dim gb As System.Text.Encoding
        gb = System.Text.Encoding.GetEncoding("gb2312")

        post = "ct=1&cm=1&spRefURL=http%3A%2F%2Fhi.baidu.com%2" & username & _
        "%2Fcreat%2Fblog%2F&spBlogTitle=" & System.Web.HttpUtility.UrlEncode(newpost.title, gb) & _
        "&spBlogText=" & System.Web.HttpUtility.UrlEncode(newpost.description, gb) & _
        "&spBlogCatName=" & System.Web.HttpUtility.UrlEncode(newpost.categories(0), gb) & _
        "&spIsCmtAllow=1&spBlogPower=0&spVcode=&spVerifyKey=&tj=+%B7%A2%B1%ED%CE%C4%D5%C2+"


        Dim c As New System.Net.Sockets.TcpClient
        c.Connect("hi.baidu.com", 80)
        c.ReceiveTimeout = 1000
        Dim ns As System.Net.Sockets.NetworkStream
        ns = c.GetStream

        Dim data As Byte()
        s = System.IO.File.ReadAllText(System.Web.HttpContext.Current.Server.MapPath("/App_Data/baidu_post.txt"))
        s = s.Replace("{username}", username)
        s = s.Replace("{BDUSS}", bduss)
        s = s.Replace("{len}", post.Length)
        s = s.Replace("{post}", post)

        data = System.Text.Encoding.ASCII.GetBytes(s)
        Log.Debug(s.Replace(vbLf, vbNewLine))
        ns.Write(data, 0, data.Length)

        Dim sr As New System.IO.StreamReader(ns)
        Dim sb As New StringBuilder

        Dim l As String
        l = sr.ReadLine
        While l.Length > 2
            sb.Append(l)
            l = sr.ReadLine & vbLf
        End While
        ns.Close()

    End Sub

End Class
