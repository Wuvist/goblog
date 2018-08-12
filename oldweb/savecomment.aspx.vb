Imports SqlConnHelper
Namespace blogwind
    Partial Class savecomment
        Inherits System.Web.UI.Page

#Region " Web Form Designer Generated Code "

        'This call is required by the Web Form Designer.
        <System.Diagnostics.DebuggerStepThrough()> Private Sub InitializeComponent()

        End Sub


        Private Sub Page_Init(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Init
            'CODEGEN: This method call is required by the Web Form Designer
            'Do not modify it using the code editor.
            InitializeComponent()
        End Sub

#End Region

        Private Sub Page_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
            'Put user code to initialize the page here
            Dim sql As String
            Dim CommentCount, ArticleOwner As Integer
            Dim ph As New pageHelper(Me.Session, Me.Request)

            Dim nick, hp, msg, Homepage As String

            Dim ArticleId, utype As Integer
            Dim uid, ip As String
            ip = util.GetPublicIP
            If ip <> "" Then
                ip = Left(ip, 15)
            End If
            uid = Session("blogger")
            If uid Is Nothing Then
                uid = ""
            End If
            nick = Request.Form("AuthorText")
            hp = Request.Form("HomepageText")
            ArticleId = CInt(Request.Form("article_id"))

            If hp.StartsWith("www.") Then
                hp = "http://" & hp
            End If
            If uid = "" Then
                utype = 1
            ElseIf hp.ToLower = "http://www.blogwind.com/" & uid.ToLower OrElse hp.ToLower = "http://www.blogwind.com/" & uid.ToLower & "/" Then
                utype = 3
            Else
                utype = 2
            End If
            Homepage = hp


            Dim pa As New ParaHelper
            pa.addInt("@id", ArticleId)
            'varchar
            pa.addVarchar("@title", Server.HtmlEncode(Request.Form("TitleText")), 50)
            pa.addVarchar("@author", nick, 50)
            sql = hp
            If sql = "http://" Then sql = ""
            pa.addVarchar("@homepage", sql, 255)
            msg = Request.Form("ContentText")

            If utype = 1 Then
                If util.CountStr(msg, "http") > 2 Then
                    Response.Write("博客风认为你是机器人，不如你修改一下内容再重新发评论？")
                    Response.End()
                ElseIf util.CountStr(msg, "www") > 2 Then
                    Response.Write("博客风认为你是机器人，不如你修改一下内容再重新发评论？")
                    Response.End()
                End If
            Else
                msg = util.AddHyperLink(msg)
            End If
            msg = Replace(msg, "  ", "　")
            msg = Replace(msg, vbNewLine, "<br>")

            pa.addNText("@content", msg)
            pa.addTinyInt("@utype", utype)
            pa.addNVarchar("@uid", uid, 50)
            pa.addVarchar("@ip", ip, 15)

            CommentCount = connHelper.exeInteger("insert_comment", CommandType.StoredProcedure, pa.parameters, False)
            If CommentCount = 0 Then
                Response.StatusCode = 404
                Exit Sub
            End If

            Dim dr As DataRow
            sql = "select id, a.blogger from blogger, Articles a where a.[index]=" & ArticleId & " and a.blogger=blogger.[index]"
            dr = connHelper.exeDataRow(sql, CommandType.Text, Nothing)
            ArticleOwner = dr(1)
            hp = "/" & dr(0) & "/" & ArticleId & ".shtml"

            Dim payload As String = ""
            Dim EventId As Integer
            sql = "select id, payload from events where event_type=1 and related_id=" & ArticleId
            pa.clearPara()
            dr = connHelper.exeDataRow(sql, CommandType.Text, Nothing)
            If dr IsNot Nothing Then
                EventId = dr.Item(0)
                If Not dr.IsNull(1) Then
                    payload = dr.Item(1)
                End If
            End If

            If String.IsNullOrEmpty(nick) Then
                nick = "无名氏"
            End If

            If EventId < 1 Then
                sql = "insert into events (owner_id, event_type, payload, related_id)values(@owner_id, 1,@payload, " & ArticleId & ")"
                pa.addInt("owner_id", ArticleOwner)
                payload = nick
            Else
                sql = "update events set update_dt=getdate(), payload=@payload where id=" & EventId
                Dim strs() As String = payload.Split(vbCr)
                Dim HasReplied As Boolean
                Dim i, j As Integer
                For i = 0 To strs.Length - 1
                    If strs(i) = nick Then
                        HasReplied = True
                        Exit For
                    End If
                Next

                If HasReplied Then
                    If i > 0 Then
                        For j = i To 1 Step -1
                            strs(j) = strs(j - 1)
                        Next
                        strs(0) = nick
                    End If
                    payload = util.JoinStr(strs, vbCr)
                Else
                    If strs.Length > 3 Then
                        For j = strs.Length - 1 To 1 Step -1
                            strs(j) = strs(j - 1)
                        Next

                        strs(0) = nick
                        payload = util.JoinStr(strs, vbCr)
                    Else
                        payload = nick & vbCr & util.JoinStr(strs, vbCr)
                    End If
                End If
            End If
            pa.addNVarchar("payload", payload, 500)
            connHelper.exeNonQuery(sql, CommandType.Text, pa.parameters)

            Dim ck As HttpCookie
            ck = New HttpCookie("User")
            ck.Name = "viewer"
            ck.Item("nick") = Convert.ToBase64String(System.Text.Encoding.UTF8.GetBytes(nick))
            ck.Item("hp") = Homepage

            ck.Expires = DateTime.Now.AddDays(365)
            Response.Cookies.Set(ck)
            Response.Cookies("User").Expires = DateTime.Now.AddDays(365)
            Response.Redirect(hp)
        End Sub
    End Class
End Namespace