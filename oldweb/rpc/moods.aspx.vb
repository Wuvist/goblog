Imports SqlConnHelper
Imports newtonsoft

Namespace blogwind
    Partial Class rpc_moods
        Inherits System.Web.UI.Page
        Private Shared ReadOnly Log As log4net.ILog = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType)
        Protected ph As pageHelper

        Private Sub Page_Load(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles MyBase.Load
            'Put user code to initialize the page here
            ph = New pageHelper(Me.Session, Me.Request)

            If Request.HttpMethod = "GET" Then
                GetMood()
            Else
                If Not ph.bloggerID > 0 Then
                    Exit Sub
                End If
                SetMood()
            End If
        End Sub

        Private Sub GetMood()
            Dim bloggerid As Integer

            If String.IsNullOrEmpty(Request.QueryString("bloggerid")) Then
                bloggerid = ph.bloggerID
            Else
                bloggerid = app.GetUserId(Request.QueryString("bloggerid"))
            End If

            Dim bm As BloggerMood
            bm = BloggerMood.GetBloggerCureentMood(bloggerid)
            If Not bm Is Nothing Then
                Dim str As String
                str = Json.JavaScriptConvert.SerializeObject(bm)
                Response.Write(str)
            End If
        End Sub

        Private Sub SetMood()
            Dim m As BloggerMood.moods
            m = [Enum].Parse(GetType(BloggerMood.moods), Request.Form("mood"))
            Dim msg As String
            msg = Request.Form("msg")

            Dim bm As New BloggerMood
            bm.Mood = m
            bm.MoodMsg = msg
            bm.BloggerId = ph.bloggerID
            bm.Save()

            Dim accs As blogwind.WindMill.AccountCollection
            Dim AccTypes As Integer() = {5, 6}
            Dim query As New SubSonic.Select(blogwind.WindMill.Account.Schema.Provider)
            query.From("accounts")
            query.Where("user_id").IsEqualTo(Me.ph.bloggerID)
            query.And("account_type").In(AccTypes)
            accs = query.ExecuteAsCollection(Of blogwind.WindMill.AccountCollection)()
            msg = "【" + m.ToString + "】" + msg
            For Each acc As blogwind.WindMill.Account In accs
                Try
                    If acc.AccountType = blogwind.WindMill.Account.AccountTypes.Twitter Then
                        Twitter.Post(acc.UserName, Cryptor.Decrypt(acc.Pwd), Twitter.Urls.Twitter, msg)
                        Log.Info("PostToTwitter: " & Me.ph.blogger)
                    Else
                        Twitter.Post(acc.UserName, Cryptor.Decrypt(acc.Pwd), Twitter.Urls.Fanfou, msg)
                        Log.Info("PostToFanfou: " & Me.ph.blogger)
                    End If
                Catch ex As Exception
                    Log.Debug(ex.ToString)
                    acc.Status = 1
                    acc.Save()
                End Try
            Next
        End Sub
    End Class



End Namespace
