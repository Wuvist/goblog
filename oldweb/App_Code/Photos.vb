Imports Microsoft.VisualBasic
Imports System.Drawing
Imports System.Drawing.Imaging
Imports System.IO
Imports SqlConnHelper
Imports ImageManipulation

Namespace Blogwind
    Public Class Photos
        Private Shared ReadOnly Log As log4net.ILog = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType)

        Enum thumbsize
            i100_100
            i320_240
            i640_480
            full
        End Enum

        Public Shared Function GetSizeFolder(ByVal size As thumbsize) As String
            Select Case size
                Case thumbsize.full
                    Return "full"
                Case thumbsize.i100_100
                    Return "100"
                Case thumbsize.i320_240
                    Return "320"
                Case thumbsize.i640_480
                    Return "640"
            End Select

            Return "unknow"
        End Function

        Public Class InvalidPhotoException
            Inherits System.ApplicationException
        End Class

        Public Shared Function GetPath(ByVal dr As DataRow, ByVal size As thumbsize) As String
            Dim path As String
            path = dr.Item("path")

            If path.IndexOf("\") > 0 Then
                Return "/images/img/" & GetSizeFolder(size) & "/" & path.Replace("\", "/")
            Else
                Return "/images/pic/" & GetSizeFolder(size) & "/" & path
            End If
        End Function

        Public Shared Function SaveImg(ByVal ImgStream As System.IO.Stream) As String
            Dim fpath As String = ini.PhotoBasePath
            Dim DatePath As String = GetDatePath()
            Dim fname As String = DatePath & Now.Ticks.ToString & ".jpg"
            Dim height, width, w, h As Integer
            Dim imgpath As String

            imgpath = Path.Combine(fpath & "\full\", fname)

            Dim fio As New System.IO.FileInfo(imgpath)
            Dim im, thumb As Image

            Try
                im = System.Drawing.Image.FromStream(ImgStream)
                Try
                    im.Save(imgpath)
                Catch ex As System.Runtime.InteropServices.ExternalException
                    Directory.CreateDirectory(fpath & "\full\" & DatePath)
                    Directory.CreateDirectory(fpath & "\640\" & DatePath)
                    Directory.CreateDirectory(fpath & "\320\" & DatePath)
                    Directory.CreateDirectory(fpath & "\100\" & DatePath)
                    im.Save(imgpath)
                End Try


                Dim ici As ImageCodecInfo = Blogwind.Photos.JpgCodec

                Dim ep As EncoderParameters = New EncoderParameters
                ep.Param(0) = New EncoderParameter(System.Drawing.Imaging.Encoder.Quality, CLng(80))

                height = im.Height
                width = im.Width
                If Blogwind.Photos.getThumbSize(height, width, h, w, Blogwind.Photos.thumbsize.i100_100) Then
                    thumb = ImagePro.myGetThumbnailImage(im, w, h, "white")
                    thumb.Save(Path.Combine(fpath & "\100", fname), ici, ep)
                End If
                If Blogwind.Photos.getThumbSize(height, width, h, w, Blogwind.Photos.thumbsize.i320_240) Then
                    thumb = ImagePro.myGetThumbnailImage(im, w, h, "white")
                    thumb.Save(Path.Combine(fpath & "\320", fname), ici, ep)
                End If
                If Blogwind.Photos.getThumbSize(height, width, h, w, Blogwind.Photos.thumbsize.i640_480) Then
                    thumb = ImagePro.myGetThumbnailImage(im, w, h, "white")
                    thumb.Save(Path.Combine(fpath & "\640", fname), ici, ep)
                End If
            Catch ex As Exception
                Log.Error(ex.ToString)
                System.IO.File.Delete(imgpath)
                Throw New InvalidPhotoException
            Finally
                If Not im Is Nothing Then
                    im.Dispose()
                End If
                If Not thumb Is Nothing Then
                    thumb.Dispose()
                End If
            End Try

            Dim pa As New ParaHelper
            pa.addNVarchar("@blogger", System.Web.HttpContext.Current.Session("blogger"), 50)
            pa.addNVarchar("@path", fname, 50)
            pa.addInt("@file_size", fio.Length / 1000)
            pa.addInt("@width", w)
            pa.addInt("@height", h)
            connHelper.exeNonQuery("insert_blog_pic", CommandType.StoredProcedure, pa.parameters)

            Return ("f" & fname & "w" & width.ToString & "h" & height.ToString).Replace("\", "/")
        End Function

        Public Shared ReadOnly Property JpgCodec() As ImageCodecInfo
            Get
                Static Dim ici As ImageCodecInfo

                If ici Is Nothing Then
                    Dim codecs As ImageCodecInfo() = ImageCodecInfo.GetImageEncoders()
                    For Each codec As ImageCodecInfo In codecs
                        If codec.MimeType = "image/jpeg" Then
                            ici = codec
                        End If
                    Next
                End If

                Return ici
            End Get
        End Property

        Public Shared Function GetDatePath() As String
            Dim d As Date = Now
            Return d.Year.ToString("0000") & "\" & d.Month.ToString("00") & d.Day.ToString("00") & "\"
        End Function

        Public Shared Function getThumbSize(ByVal height As Integer, ByVal width As Integer, ByRef h As Integer, ByRef w As Integer, ByVal size As thumbsize) As Boolean
            Dim rate As Double
            Dim f As Boolean = False
            h = height
            Select Case size
                Case thumbsize.i100_100
                    If width > 100 Then
                        rate = width / 100
                        w = 100
                        h = height / rate
                        f = True
                    End If
                    If h > 100 Then
                        rate = height / 100
                        h = 100
                        w = width / rate
                        f = True
                    End If
                Case thumbsize.i320_240
                    If width > 320 Then
                        rate = width / 320
                        w = 320
                        h = height / rate
                        f = True
                    End If
                    If h > 240 Then
                        rate = height / 240
                        h = 240
                        w = width / rate
                        f = True
                    End If
                Case thumbsize.i640_480
                    If width > 640 Then
                        rate = width / 640
                        w = 640
                        h = height / rate
                        f = True
                    End If
                    If h > 480 Then
                        rate = height / 480
                        h = 480
                        w = width / rate
                        f = True
                    End If
            End Select
            Return f
        End Function

    End Class
End Namespace
