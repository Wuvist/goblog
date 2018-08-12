<%@ WebHandler Language="VB" Class="upload" %>

Imports System
Imports System.Web
Imports System.IO
Imports System.Drawing
Imports System.Drawing.Imaging
Imports SqlConnHelper

Public Class upload
    Inherits Blogwind.Blog.SwfHandler
    
    Public Overrides Sub ProcessAuthRequest(ByVal context As System.Web.HttpContext)
        If Not context.Request.Files(0).FileName.ToLower.EndsWith("jpg") Then
            'not jpg
            context.Response.Write("{error:1}")
        Else
            Try
                context.Response.Write("{file:'" & Blogwind.Photos.SaveImg(context.Request.Files(0).InputStream) & "'}")
            Catch ex As Blogwind.Photos.InvalidPhotoException
                'decode error
                context.Response.Write("{error:2}")
            End Try
        End If
    End Sub
End Class