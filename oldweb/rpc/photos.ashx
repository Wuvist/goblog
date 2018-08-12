<%@ WebHandler Language="VB" Class="photos" %>

Imports System
Imports System.Web
Imports SqlConnHelper

Public Class photos
    Inherits Blogwind.Blog.BWAuthHandler
    
    Private Shared ReadOnly Log As log4net.ILog = log4net.LogManager.GetLogger(System.Reflection.MethodBase.GetCurrentMethod().DeclaringType)
    
    Public Overrides Sub ProcessAuthRequest(ByVal context As System.Web.HttpContext)
        Dim id As Integer
        id = context.Request("id")
        Dim action As String = context.Request("action")
        Select Case action
            Case "delete"
                Dim sql As String
                Dim uid As Integer = app.GetUid(Me.ph.blogger)
                Dim dr As DataRow
                sql = "select * from pic where uid=" & uid & " and id=" & id
                
                dr = connHelper.exeDataRow(sql, CommandType.Text, Nothing)
                If dr Is Nothing Then
                    context.Response.Write("1")
                    Exit Sub
                End If
                
                Dim prefix, path As String
                path = dr.Item("path")
                If path.IndexOf("\") > 0 Then
                    prefix = "images\img\"
                Else
                    prefix = "images\pic\"
                End If
                
                Dim folders As String() = {"100\", "320\", "480\", "full\"}
                For Each folder As String In folders
                    Try
                        System.IO.File.Delete(context.Server.MapPath("..\" & prefix & folder & path))
                    Catch ex As Exception

                    End Try
                Next
                
                sql = "delete from pic where id=" & id
                connHelper.exeNonQuery(sql, CommandType.Text, Nothing)
                context.Response.Write("0")
        End Select

    End Sub
End Class