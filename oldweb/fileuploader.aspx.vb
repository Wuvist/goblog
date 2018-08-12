Imports System.IO
Imports System.Drawing
Imports System.Drawing.Imaging
Imports SqlConnHelper
Namespace blogwind
    Partial Class fileuploader
        Inherits System.Web.UI.Page
        Private _uploadID As String
        Public hint As String
        Public ReadOnly Property uploadID() As String
            Get
                Return _uploadID
            End Get
        End Property

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
            If Session("blogger") = "" Then
                'Response.End()
            End If

        End Sub

        Private Sub btn_upload_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles btn_upload.Click
            If Not file1.PostedFile.FileName.ToLower.EndsWith("jpg") Then
                hint = "请上传jpg格式文件"
            Else
                Try
                    hint = blogwind.Photos.SaveImg(file1.PostedFile.InputStream)
                Catch ex As blogwind.Photos.InvalidPhotoException
                    hint = "不好意思，博客风无法识别你上传的图片"
                End Try
            End If
        End Sub
    End Class
End Namespace