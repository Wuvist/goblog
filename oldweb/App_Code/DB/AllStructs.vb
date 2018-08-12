Imports System 
Imports System.Text 
Imports System.Data
Imports System.Data.SqlClient
Imports System.Data.Common
Imports System.Collections
Imports System.Collections.Generic
Imports System.ComponentModel
Imports System.Configuration 
Imports System.Xml 
Imports System.Xml.Serialization
Imports SubSonic 
Imports SubSonic.Utilities
Namespace Blogwind.WindMill
	#Region "Tables Struct"
	Public Partial Structure Tables
		Dim x As Integer
		
		Public Shared  Account As String = "accounts"
        
	End Structure
	#End Region
    #Region "View Struct"
    Public Partial Structure Views
		Dim x As Integer
		
    End Structure
    #End Region
	#Region "Query Factories"
	Public Partial Class DB
		Private Sub New()
		End Sub
        Public Shared Function SelectAllColumnsFrom(Of T As {RecordBase(Of T), New})() As [Select]
			Return SubSonic.Select.AllColumnsFrom(Of T)()
		End Function
		Public Shared Function [Select]() As [Select]
			Return New [Select](DataService.Providers("WindMill"))
		End Function
		Public Shared Function [Select](ParamArray ByVal columns As String()) As [Select]
			Return New [Select](DataService.Providers("WindMill"), columns)
		End Function
		Public Shared Function [Select](ParamArray ByVal aggregates As Aggregate()) As [Select]
			Return New [Select](DataService.Providers("WindMill"), aggregates)
		End Function
		Public Shared Function Update(Of T As {ActiveRecord(Of T), New})() As Update
			Return New SubSonic.Update(New T().GetSchema())
		End Function
		
		Public Shared Function Insert() As Insert
			Dim i As Insert = New Insert()
			i.provider=DataService.Providers("WindMill")
			Return i
		End Function
		Public Shared Function Delete() As Delete
			Return New Delete("WindMill")
		End Function
		Public Shared Function Query() As InlineQuery
			   Return New InlineQuery("WindMill")
		End Function
        
	End Class
	#End Region
End Namespace
#Region "Databases"
Public Partial Structure Databases
	Dim x As Integer
	
	Public Shared WindMill As String = "WindMill"
    
End Structure
#End Region
