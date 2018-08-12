<%@ WebHandler Language="VB" Class="Friends" %>

Imports System
Imports System.Web
Imports Newtonsoft.Json
Imports SqlConnHelper

Public Class Friends : Implements IHttpHandler, IRequiresSessionState
    
    Public Sub ProcessRequest(ByVal context As HttpContext) Implements IHttpHandler.ProcessRequest
        context.Response.ContentType = "text/plain"
        If String.IsNullOrEmpty(context.Request.Form("action")) Then
            Dim sql As String
            Dim dt As DataTable
            sql = "select_friendlist " & context.Session("bloggerID")
            dt = connHelper.retrieve(sql, CommandType.Text, Nothing)
            Dim jw As New JsonWriter(context.Response.Output)
            jw.WriteStartArray()
            For Each dr As DataRow In dt.Rows
                jw.WriteStartObject()
                jw.WritePropertyName("id")
                jw.WriteValue(dr.Item("id"))
            
                jw.WritePropertyName("nick")
                jw.WriteValue(dr.Item("nick"))
            
                jw.WritePropertyName("icon")
                jw.WriteValue(app.GetIconPath(dr.Item("id"), dr.Item("userpic") Mod 10))
                jw.WriteEndObject()
            Next
            jw.WriteEndArray()
            jw.Close()
        Else
            Dim pa As New ParaHelper
            pa.addVarchar("@friendid", context.Request.Form("friendid"), 50)
            pa.addVarchar("@blogger", context.Session("blogger"), 50)
            connHelper.exeNonQuery("del_friend", CommandType.StoredProcedure, pa.parameters)
            context.Response.Write(context.Request.Form("friendid"))
        End If
        
        
    End Sub
 
    Public ReadOnly Property IsReusable() As Boolean Implements IHttpHandler.IsReusable
        Get
            Return False
        End Get
    End Property

End Class