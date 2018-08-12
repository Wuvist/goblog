Imports SqlConnHelper
Namespace blogwind
    Partial Class write
        Inherits BWPage
        Protected blogger As String
        Protected ph As pageHelper
        Private Shared ReadOnly Log As log4net.ILog = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType)


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
            ph = New pageHelper(Me.Session, Me.Request)

            If Session("blogger") = "" Then
                If IsPostBack Then
                    Session("blogger") = b.Text
                End If
            End If

            If Session("blogger") = "" Then
                Response.Write("请重新登陆")
                Response.End()
            End If

            'cn.exeNonQuery("update blogger set writer=0 where id='" & blogger & "'", CommandType.Text, False)
            Button1.Text = ph.g("发表")
            Button2.Text = ph.g("保存")

            If Not Page.IsPostBack Then
                b.Text = Session("blogger")
                Dim ds As New DataSet
                Dim sql As String

                blogger = Session("blogger")

                Dim pa As New ParaHelper
                sql = "select_cate_by_Blogger"
                pa.addVarchar("@blogger", blogger, 50)
                cate.DataSource = connHelper.retrieve(sql, CommandType.StoredProcedure, pa.parameters)
                cate.DataBind()

                Dim dr As DataRow
                If Request.QueryString("edit") <> "" Then
                    mytags.Visible = False
                    ds.Tables.Add("Article")
                    sql = "select_doc"

                    pa.clearPara()
                    pa.addVarchar("@blogger", blogger, 50)
                    pa.addInt("@id", Request.QueryString("edit"))

                    dr = connHelper.exeDataRow(sql, CommandType.StoredProcedure, pa.parameters)

                    title.Text = dr.Item("Title")
                    WTB1.Text = dr.Item("Content")
                    cate.Items.FindByValue(dr.Item("cate_index")).Selected = True
                    Button2.Enabled = False
                    message.Text = "请修改已发表的文章。"
                Else
                    sql = "select TS from blogger where id='" & Session("blogger") & "'"

                    If connHelper.exeScalar(sql, CommandType.Text, Nothing) Then
                        ds.Tables.Add("TS")
                        sql = "select_TS"
                        pa.clearPara()
                        pa.addVarchar("@blogger", blogger, 50)
                        dr = connHelper.exeDataRow(sql, CommandType.StoredProcedure, pa.parameters)

                        Try
                            title.Text = dr.Item("Title")
                        Catch ex As Exception
                            title.Text = ""
                        End Try
                        Try
                            WTB1.Text = dr.Item("Content")
                        Catch ex As Exception
                            WTB1.Text = ""
                        End Try
                        Try
                            cate.Items.FindByValue(dr.Item("cate_index")).Selected = True
                        Catch ex As Exception

                        End Try

                        message.Text = "请继续编辑未完成的网志。"
                    End If
                End If

                If Request.QueryString("con") <> "" Then
                    WTB1.Text = Request.QueryString("con")
                End If

                cate.Items(0).Selected = True
            End If
        End Sub

        Private Sub Button1_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Button1.Click
            Dim sql, msg As String
            Dim aid As Integer
            Dim pa As New ParaHelper

            If Request.QueryString("edit") <> "" Then
                aid = CInt(Request.QueryString("edit"))
                sql = "update_doc"

                pa.addVarchar("@title", title.Text, 50)
                pa.addVarchar("@blogger", Session("blogger"), 50)
                pa.addNText("@content", WTB1.Text)
                pa.addInt("@id", aid)
                pa.addInt("@cate", cate.SelectedItem.Value)

                connHelper.exeNonQuery(sql, CommandType.StoredProcedure, pa.parameters)
                Response.Redirect("postguide.aspx?action=edited&aid=" & aid)
                'Response.Redirect("system/admin_article.aspx")
            Else
                pa.addVarchar("@title", title.Text, 50)
                pa.addVarchar("@blogger", Session("blogger"), 50)
                pa.addNText("@content", WTB1.Text)
                pa.addInt("@cate", cate.SelectedItem.Value)

                If mytags.Visible = True And mytags.Text <> "" Then
                    Dim t As String
                    t = mytags.Text
                    mytags.Text = ""
                    t = t.Replace("，", ",")
                    t = t.Replace(";", ",")
                    t = t.Replace("；", ",")
                    t = t.Replace("'", "")
                    t = t.Replace("""", "")
                    t = t.Replace("|", "")
                    '加多一个逗号，方便存储过程切分
                    t = t & ","
                    pa.addNVarchar("@tags", t, 50)
                Else
                    'cn.addPara("@tags", SqlDbType.NVarChar, 50, "")
                    pa.addNVarchar("@tags", "", 50)
                End If

                aid = connHelper.exeInteger("insert_doc", CommandType.StoredProcedure, pa.parameters, False)

                Dim accs As blogwind.WindMill.AccountCollection
                Dim AccTypes As Integer() = {blogwind.WindMill.Account.AccountTypes.MSNSpace, _
                    blogwind.WindMill.Account.AccountTypes.Baidu, _
                    blogwind.WindMill.Account.AccountTypes.WordPress, _
                    blogwind.WindMill.Account.AccountTypes.Blogger, _
                    blogwind.WindMill.Account.AccountTypes.MetaWebLog}

                Dim query As New SubSonic.Select(blogwind.WindMill.Account.Schema.Provider)
                query.From("accounts")
                query.Where("user_id").IsEqualTo(Me.ph.bloggerID)
                query.And("account_type").In(AccTypes)
                accs = query.ExecuteAsCollection(Of blogwind.WindMill.AccountCollection)()

                For Each acc As blogwind.WindMill.Account In accs
                    If acc.AccountType = blogwind.WindMill.Account.AccountTypes.Baidu Then
                        Try
                            baidu.Post(acc, title.Text, WTB1.Text, cate.SelectedItem.Text)
                        Catch ex As Exception
                            Log.Debug(ex.ToString)
                            acc.Status = 1
                            acc.Save()
                        End Try
                    Else
                        Try
                            Metaweblog.Post(acc, title.Text, WTB1.Text, cate.SelectedItem.Text)
                        Catch ex As Exception
                            Log.Debug(ex.ToString)
                            acc.Status = 1
                            acc.Save()
                        End Try
                    End If
                Next


                Response.Redirect("postguide.aspx?action=added&aid=" & aid.ToString)
            End If
        End Sub

        Private Sub Button2_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Button2.Click
            Dim sql As String
            sql = "insert_ts"
            Dim pa As New ParaHelper
            pa.addVarchar("@title", title.Text, 50)
            pa.addVarchar("@blogger", Session("blogger"), 50)
            pa.addNText("@content", WTB1.Text)
            pa.addInt("@cate", cate.SelectedItem.Value)

            connHelper.exeNonQuery(sql, CommandType.StoredProcedure, pa.parameters)
            message.Text = "保存成功！"
        End Sub
    End Class
End Namespace
