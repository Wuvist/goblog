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
    ''' <summary>
    ''' Controller class for accounts
    ''' </summary>
		<System.ComponentModel.DataObject()> Public Partial Class AccountController
    
        ' Preload our schema..
        Dim thisSchemaLoad As Account = New Account()
        Private strUserName As String = String.Empty
        Protected ReadOnly Property UserName() As String
            Get
				If strUserName.Length = 0 Then
		        
    				If Not System.Web.HttpContext.Current Is Nothing Then
						strUserName = System.Web.HttpContext.Current.User.Identity.Name
					Else
		        		strUserName = System.Threading.Thread.CurrentPrincipal.Identity.Name
					End If
					Return strUserName
				End If
				Return strUserName
			End Get
        End Property
        <DataObjectMethod(DataObjectMethodType.Select, True)> Public Function FetchAll() As AccountCollection
        
            Dim coll As AccountCollection = New AccountCollection()
            Dim qry As Query = New Query(Account.Schema)
            coll.LoadAndCloseReader(qry.ExecuteReader())
            Return coll
            
        End Function
        <DataObjectMethod(DataObjectMethodType.Select, True)> Public Function FetchByID(ByVal Id As Object) As AccountCollection 
        
            Dim coll As AccountCollection = New AccountCollection().Where("id", Id).Load()
            Return coll
        
        End Function
        
        <DataObjectMethod(DataObjectMethodType.Select, True)> Public Function FetchByQuery(ByVal qry As SubSonic.Query) As AccountCollection 
        
            Dim coll As AccountCollection = New AccountCollection()
            coll.LoadAndCloseReader(qry.ExecuteReader())
            Return coll
        
        End Function
        <DataObjectMethod(DataObjectMethodType.Delete, True)> Public Function Delete(ByVal Id As Object) as Boolean
        
            Return (Account.Delete(Id) = 1)
        
        End Function
        <DataObjectMethod(DataObjectMethodType.Delete, False)> Public Function Destroy(ByVal Id As Object) as Boolean
        
            Return (Account.Destroy(Id) = 1)
        
        End Function
        
    	
	    ''' <summary>
	    ''' Inserts a record, can be used with the Object Data Source
	    ''' </summary>
        <DataObjectMethod(DataObjectMethodType.Insert, True)> Public Sub Insert(ByVal UserId As Integer,ByVal AccountType As Short,ByVal ServerHost As String,ByVal Blogid As String,ByVal UserName As String,ByVal Pwd As String,ByVal Status As Byte,ByVal CreateOn As DateTime)
	   
		    Dim item As Account = New Account()
		    
            item.UserId = UserId
            
            item.AccountType = AccountType
            
            item.ServerHost = ServerHost
            
            item.Blogid = Blogid
            
            item.UserName = UserName
            
            item.Pwd = Pwd
            
            item.Status = Status
            
            item.CreateOn = CreateOn
            
	    
		    item.Save(UserName)
	   
	   End Sub
    	
	    ''' <summary>
	    ''' Updates a record, can be used with the Object Data Source
	    ''' </summary>
        <DataObjectMethod(DataObjectMethodType.Update, True)> Public Sub Update(ByVal Id As Integer,ByVal UserId As Integer,ByVal AccountType As Short,ByVal ServerHost As String,ByVal Blogid As String,ByVal UserName As String,ByVal Pwd As String,ByVal Status As Byte,ByVal CreateOn As DateTime)
	    
		    Dim item As Account = New Account()
		    
				item.Id = Id
				
				item.UserId = UserId
				
				item.AccountType = AccountType
				
				item.ServerHost = ServerHost
				
				item.Blogid = Blogid
				
				item.UserName = UserName
				
				item.Pwd = Pwd
				
				item.Status = Status
				
				item.CreateOn = CreateOn
				
		    item.MarkOld()
		    item.Save(UserName)
	    
	    End Sub
    End Class
End Namespace
