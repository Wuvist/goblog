Imports SqlConnHelper
Namespace blogwind
    Partial Class userpic_change
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
            If Session("blogger") = "" Then
                Response.End()
            End If
        End Sub

        Private Sub Button1_Click(ByVal sender As System.Object, ByVal e As System.EventArgs) Handles Button1.Click
            Dim err As Boolean
            Dim e_name As String
            err = False
            Dim width, height As Integer
            Dim im As System.Drawing.Image
            Dim output As Bitmap
            Dim blogger As String
            blogger = Session("blogger")
            If blogger = "" Then
                Response.End()
            End If

            If err = False Then
                e_name = Right(Up_file.PostedFile.FileName, 3).ToLower
                If Not (e_name = "jpg" Or e_name = "gif" Or e_name = "png" Or e_name = "") Then
                    hint.Text = "头像需为jpg/gif/png格式"
                    err = True
                End If
                If err = False Then
                    im = System.Drawing.Image.FromStream(Up_file.PostedFile.InputStream)
                    width = im.Width
                    height = im.Height

                    If im.Height > 1024 Or im.Width > 768 Then
                        hint.Text = "头像宽高限1024*768象素"
                        err = True
                        im.Dispose()
                    End If
                End If


            End If

            If err = False Then

                If width > 100 Or height > 100 Then
                    Dim callb As System.Drawing.Image.GetThumbnailImageAbort

                    If width > 100 Then
                        height = height * 100 / width
                        width = 100
                    End If

                    If height > 100 Then
                        width = width * 100 / height
                        height = 100
                    End If
                    Dim thumbnailImage As System.Drawing.Image
                    thumbnailImage = im.GetThumbnailImage(width, height, callb, New System.IntPtr)
                    output = New Bitmap(thumbnailImage)
                Else
                    output = New Bitmap(im)
                End If

                output.Save(Server.MapPath("images/userpic/" & blogger & ".jpg"), System.Drawing.Imaging.ImageFormat.Jpeg)
                'output.Save(Server.MapPath(blogger & ".jpg"), System.Drawing.Imaging.ImageFormat.Jpeg)
                output.Dispose()
                connHelper.exeNonQuery("update blogger set userpic=11 where id='" & blogger & "'", CommandType.Text, Nothing)
                im.Dispose()

                Response.Write("<script>window.parent.location.reload();</script>")

            End If

        End Sub
    End Class
End Namespace