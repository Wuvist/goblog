Imports SqlConnHelper

Namespace blogwind
    Partial Class icon
        Inherits System.Web.UI.UserControl
        Public userpic, cmdlist, formlist As String
        Public isblogger As Integer
        Public deco, enddeco, FriendId As String
        Protected widget As String

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
            Dim ph As New pageHelper(Me.Session, Me.Request)
            Dim blogger As String
            blogger = Session("blogger")
            If Not String.IsNullOrEmpty(blogger) Then
                blogger = blogger.ToLower
            End If

            FriendId = Request.QueryString("blogger")
            If String.IsNullOrEmpty(FriendId) AndAlso (Not String.IsNullOrEmpty(blogger)) Then
                FriendId = blogger
            End If

            If String.IsNullOrEmpty(FriendId) Then
                Exit Sub
            End If
            FriendId = FriendId.ToLower

            Dim i, piccon As Integer
            Dim dr As DataRow
            dr = connHelper.exeDataRow("select userpic, widget from blogger where id='" & FriendId & "'", CommandType.Text, Nothing)
            piccon = dr.Item(0)
            If Not String.IsNullOrEmpty(Request.Url.Query) And Not dr.IsNull(1) Then
                widget = dr.Item(1)
            End If

            If blogger = FriendId Then
                isblogger = 1
                'blogger is viewing own blog
                cmdlist = "<input type=button id='IconBtn'  value='修改头像' class='logbutton' onclick='icon_change()'>"
            Else
                'guest
                isblogger = 0
                If blogger <> "" Then
                    'blogger viewing other's blog
                    isblogger = 2

                    'hoho, I'm admin
                    If blogger.ToLower = "wuvist" Then
                        isblogger = 1
                    End If

                    cmdlist = "<input type=button value='复制头像' class='logbutton' onclick=""icon_copy('" & FriendId & "')"">"
                End If
            End If

            i = piccon Mod 10
            userpic = app.GetIconPath(FriendId, i)

            If String.IsNullOrEmpty(deco) Then
                userpic = "<div class='shadow1'><div class='shadow2'><div class='shadow3'><div class='alpha'><img src='" & userpic & "'></div></div></div></div>"
            Else
                userpic = deco & userpic & enddeco
            End If

            userpic = "<div id=""userpic"" style=""userpic"">" & userpic & "</div><br style=""clear:both"" />"

            If Not String.IsNullOrEmpty(cmdlist) Then
                cmdlist = cmdlist & "<br style=""clear:both"" />"
            End If

            If Not String.IsNullOrEmpty(blogger) Then
                formlist = "<form action=""/userpic.aspx"" method=post name=""userpicform"" style=""display: inline;""><input type=hidden name=""action"" value=""""></form>"
            End If
        End Sub

    End Class
End Namespace