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
	''' Strongly-typed collection for the Account class.
	''' </summary>
	<Serializable> _
	Public Partial Class AccountCollection 
	Inherits ActiveList(Of Account, AccountCollection)
	    Public Sub New()
		End Sub
	End Class
	''' <summary>
	''' This is an ActiveRecord class which wraps the accounts table.
	''' </summary>
	<Serializable> _
	Public Partial Class Account 
	Inherits ActiveRecord(Of Account)
		#Region ".ctors and Default Settings"
		
		Public Sub New()
			SetSQLProps()
			InitSetDefaults()
			MarkNew()
		End Sub
		
		Public Sub New(ByVal useDatabaseDefaults As Boolean)
			SetSQLProps()
			If useDatabaseDefaults = True Then
				ForceDefaults()
			End If
			MarkNew()
		End Sub
		Public Sub New(ByVal keyID As Object)
			SetSQLProps()
			InitSetDefaults()
			LoadByKey(keyID)
		End Sub
		Public Sub New(ByVal columnName As String, ByVal columnValue As Object)
			SetSQLProps()
			InitSetDefaults()
			LoadByParam(columnName,columnValue)
		End Sub
		Private Sub InitSetDefaults()
			SetDefaults()
		End Sub
		
		Protected Shared Sub SetSQLProps()
			GetTableSchema()
		End Sub
		#End Region
		
		#Region "Schema and Query Accessor"
		
		Public Shared ReadOnly Property Schema() As TableSchema.Table
			Get
				If BaseSchema Is Nothing Then
					SetSQLProps()
				End If
				Return BaseSchema
			End Get
		End Property
		Private Shared Sub GetTableSchema()
			If (Not IsSchemaInitialized) Then
				'Schema declaration
				Dim schema As TableSchema.Table = New TableSchema.Table("accounts", TableType.Table, DataService.GetInstance("WindMill"))
				schema.Columns = New TableSchema.TableColumnCollection()
				schema.SchemaName = "dbo"
				'columns
				
                
                Dim colvarId As TableSchema.TableColumn = New TableSchema.TableColumn(schema)
                colvarId.ColumnName = "id"
                colvarId.DataType = DbType.Int32
                colvarId.MaxLength = 0
                colvarId.AutoIncrement = true
                colvarId.IsNullable = false
                colvarId.IsPrimaryKey = true
                colvarId.IsForeignKey = false
                colvarId.IsReadOnly = false
                
                
                schema.Columns.Add(colvarId)
                
                Dim colvarUserId As TableSchema.TableColumn = New TableSchema.TableColumn(schema)
                colvarUserId.ColumnName = "user_id"
                colvarUserId.DataType = DbType.Int32
                colvarUserId.MaxLength = 0
                colvarUserId.AutoIncrement = false
                colvarUserId.IsNullable = false
                colvarUserId.IsPrimaryKey = false
                colvarUserId.IsForeignKey = false
                colvarUserId.IsReadOnly = false
                
                
                schema.Columns.Add(colvarUserId)
                
                Dim colvarAccountType As TableSchema.TableColumn = New TableSchema.TableColumn(schema)
                colvarAccountType.ColumnName = "account_type"
                colvarAccountType.DataType = DbType.Int16
                colvarAccountType.MaxLength = 0
                colvarAccountType.AutoIncrement = false
                colvarAccountType.IsNullable = false
                colvarAccountType.IsPrimaryKey = false
                colvarAccountType.IsForeignKey = false
                colvarAccountType.IsReadOnly = false
                
                
                schema.Columns.Add(colvarAccountType)
                
                Dim colvarServerHost As TableSchema.TableColumn = New TableSchema.TableColumn(schema)
                colvarServerHost.ColumnName = "server_host"
                colvarServerHost.DataType = DbType.String
                colvarServerHost.MaxLength = 255
                colvarServerHost.AutoIncrement = false
                colvarServerHost.IsNullable = true
                colvarServerHost.IsPrimaryKey = false
                colvarServerHost.IsForeignKey = false
                colvarServerHost.IsReadOnly = false
                
                
                schema.Columns.Add(colvarServerHost)
                
                Dim colvarBlogid As TableSchema.TableColumn = New TableSchema.TableColumn(schema)
                colvarBlogid.ColumnName = "blogid"
                colvarBlogid.DataType = DbType.String
                colvarBlogid.MaxLength = 50
                colvarBlogid.AutoIncrement = false
                colvarBlogid.IsNullable = true
                colvarBlogid.IsPrimaryKey = false
                colvarBlogid.IsForeignKey = false
                colvarBlogid.IsReadOnly = false
                
                
                schema.Columns.Add(colvarBlogid)
                
                Dim colvarUserName As TableSchema.TableColumn = New TableSchema.TableColumn(schema)
                colvarUserName.ColumnName = "user_name"
                colvarUserName.DataType = DbType.String
                colvarUserName.MaxLength = 50
                colvarUserName.AutoIncrement = false
                colvarUserName.IsNullable = false
                colvarUserName.IsPrimaryKey = false
                colvarUserName.IsForeignKey = false
                colvarUserName.IsReadOnly = false
                
                
                schema.Columns.Add(colvarUserName)
                
                Dim colvarPwd As TableSchema.TableColumn = New TableSchema.TableColumn(schema)
                colvarPwd.ColumnName = "pwd"
                colvarPwd.DataType = DbType.String
                colvarPwd.MaxLength = 500
                colvarPwd.AutoIncrement = false
                colvarPwd.IsNullable = false
                colvarPwd.IsPrimaryKey = false
                colvarPwd.IsForeignKey = false
                colvarPwd.IsReadOnly = false
                
                
                schema.Columns.Add(colvarPwd)
                
                Dim colvarStatus As TableSchema.TableColumn = New TableSchema.TableColumn(schema)
                colvarStatus.ColumnName = "status"
                colvarStatus.DataType = DbType.Byte
                colvarStatus.MaxLength = 0
                colvarStatus.AutoIncrement = false
                colvarStatus.IsNullable = false
                colvarStatus.IsPrimaryKey = false
                colvarStatus.IsForeignKey = false
                colvarStatus.IsReadOnly = false
                
						colvarStatus.DefaultSetting = "((0))"
				
                
                schema.Columns.Add(colvarStatus)
                
                Dim colvarCreateOn As TableSchema.TableColumn = New TableSchema.TableColumn(schema)
                colvarCreateOn.ColumnName = "CreateOn"
                colvarCreateOn.DataType = DbType.DateTime
                colvarCreateOn.MaxLength = 0
                colvarCreateOn.AutoIncrement = false
                colvarCreateOn.IsNullable = false
                colvarCreateOn.IsPrimaryKey = false
                colvarCreateOn.IsForeignKey = false
                colvarCreateOn.IsReadOnly = false
                
						colvarCreateOn.DefaultSetting = "(getdate())"
				
                
                schema.Columns.Add(colvarCreateOn)
                
				BaseSchema = schema
				
				'add this schema to the provider
                'so we can query it later
                DataService.Providers("WindMill").AddSchema("accounts",schema)
			End If
		End Sub
		Public Shared Function CreateQuery() As Query
			Return New Query(Schema)
		End Function
		
		#End Region
		
		#Region "Props"
	
          
        <XmlAttribute("Id")> _
        Public Property Id As Integer 
			Get
				Return GetColumnValue(Of Integer)(Columns.Id)
			End Get
		    
			Set
				SetColumnValue(Columns.Id, Value)
			End Set
		End Property
		
		
		  
        <XmlAttribute("UserId")> _
        Public Property UserId As Integer 
			Get
				Return GetColumnValue(Of Integer)(Columns.UserId)
			End Get
		    
			Set
				SetColumnValue(Columns.UserId, Value)
			End Set
		End Property
		
		
		  
        <XmlAttribute("AccountType")> _
        Public Property AccountType As Short 
			Get
				Return GetColumnValue(Of Short)(Columns.AccountType)
			End Get
		    
			Set
				SetColumnValue(Columns.AccountType, Value)
			End Set
		End Property
		
		
		  
        <XmlAttribute("ServerHost")> _
        Public Property ServerHost As String 
			Get
				Return GetColumnValue(Of String)(Columns.ServerHost)
			End Get
		    
			Set
				SetColumnValue(Columns.ServerHost, Value)
			End Set
		End Property
		
		
		  
        <XmlAttribute("Blogid")> _
        Public Property Blogid As String 
			Get
				Return GetColumnValue(Of String)(Columns.Blogid)
			End Get
		    
			Set
				SetColumnValue(Columns.Blogid, Value)
			End Set
		End Property
		
		
		  
        <XmlAttribute("UserName")> _
        Public Property UserName As String 
			Get
				Return GetColumnValue(Of String)(Columns.UserName)
			End Get
		    
			Set
				SetColumnValue(Columns.UserName, Value)
			End Set
		End Property
		
		
		  
        <XmlAttribute("Pwd")> _
        Public Property Pwd As String 
			Get
				Return GetColumnValue(Of String)(Columns.Pwd)
			End Get
		    
			Set
				SetColumnValue(Columns.Pwd, Value)
			End Set
		End Property
		
		
		  
        <XmlAttribute("Status")> _
        Public Property Status As Byte 
			Get
				Return GetColumnValue(Of Byte)(Columns.Status)
			End Get
		    
			Set
				SetColumnValue(Columns.Status, Value)
			End Set
		End Property
		
		
		  
        <XmlAttribute("CreateOn")> _
        Public Property CreateOn As DateTime 
			Get
				Return GetColumnValue(Of DateTime)(Columns.CreateOn)
			End Get
		    
			Set
				SetColumnValue(Columns.CreateOn, Value)
			End Set
		End Property
		
		
		
		
		#End Region
		
		
			    
	    
		
		
		
		
		
	    
	    'no foreign key tables defined (0)
	    
		
	    
	    'no ManyToMany tables defined (0)
	    
		
		#Region "ObjectDataSource support"
		
		''' <summary>
		''' Inserts a record, can be used with the Object Data Source
		''' </summary>
		Public Shared Sub Insert(ByVal varUserId As Integer,ByVal varAccountType As Short,ByVal varServerHost As String,ByVal varBlogid As String,ByVal varUserName As String,ByVal varPwd As String,ByVal varStatus As Byte,ByVal varCreateOn As DateTime)
			Dim item As Account = New Account()
			
            item.UserId = varUserId
            
            item.AccountType = varAccountType
            
            item.ServerHost = varServerHost
            
            item.Blogid = varBlogid
            
            item.UserName = varUserName
            
            item.Pwd = varPwd
            
            item.Status = varStatus
            
            item.CreateOn = varCreateOn
            
			If Not System.Web.HttpContext.Current Is Nothing Then
				item.Save(System.Web.HttpContext.Current.User.Identity.Name)
			Else
				item.Save(System.Threading.Thread.CurrentPrincipal.Identity.Name)
			End If
		End Sub
		''' <summary>
		''' Updates a record, can be used with the Object Data Source
		''' </summary>
		Public Shared Sub Update(ByVal varId As Integer,ByVal varUserId As Integer,ByVal varAccountType As Short,ByVal varServerHost As String,ByVal varBlogid As String,ByVal varUserName As String,ByVal varPwd As String,ByVal varStatus As Byte,ByVal varCreateOn As DateTime)
			Dim item As Account = New Account()
		    
                item.Id = varId
				
                item.UserId = varUserId
				
                item.AccountType = varAccountType
				
                item.ServerHost = varServerHost
				
                item.Blogid = varBlogid
				
                item.UserName = varUserName
				
                item.Pwd = varPwd
				
                item.Status = varStatus
				
                item.CreateOn = varCreateOn
				
			item.IsNew = False
			If Not System.Web.HttpContext.Current Is Nothing Then
				item.Save(System.Web.HttpContext.Current.User.Identity.Name)
			Else
				item.Save(System.Threading.Thread.CurrentPrincipal.Identity.Name)
			End If
		End Sub
		#End Region
		
		#Region "Typed Columns"
        
        
        Public Shared ReadOnly Property IdColumn() As TableSchema.TableColumn
            Get
                Return Schema.Columns(0)
            End Get
        End Property
        
        
        Public Shared ReadOnly Property UserIdColumn() As TableSchema.TableColumn
            Get
                Return Schema.Columns(1)
            End Get
        End Property
        
        
        Public Shared ReadOnly Property AccountTypeColumn() As TableSchema.TableColumn
            Get
                Return Schema.Columns(2)
            End Get
        End Property
        
        
        Public Shared ReadOnly Property ServerHostColumn() As TableSchema.TableColumn
            Get
                Return Schema.Columns(3)
            End Get
        End Property
        
        
        Public Shared ReadOnly Property BlogidColumn() As TableSchema.TableColumn
            Get
                Return Schema.Columns(4)
            End Get
        End Property
        
        
        Public Shared ReadOnly Property UserNameColumn() As TableSchema.TableColumn
            Get
                Return Schema.Columns(5)
            End Get
        End Property
        
        
        Public Shared ReadOnly Property PwdColumn() As TableSchema.TableColumn
            Get
                Return Schema.Columns(6)
            End Get
        End Property
        
        
        Public Shared ReadOnly Property StatusColumn() As TableSchema.TableColumn
            Get
                Return Schema.Columns(7)
            End Get
        End Property
        
        
        Public Shared ReadOnly Property CreateOnColumn() As TableSchema.TableColumn
            Get
                Return Schema.Columns(8)
            End Get
        End Property
        
        
        #End Region
		
		#Region "Columns Struct"
		Public Structure Columns
			Dim x as Integer
			
            Public Shared Id As String = "id"
            
            Public Shared UserId As String = "user_id"
            
            Public Shared AccountType As String = "account_type"
            
            Public Shared ServerHost As String = "server_host"
            
            Public Shared Blogid As String = "blogid"
            
            Public Shared UserName As String = "user_name"
            
            Public Shared Pwd As String = "pwd"
            
            Public Shared Status As String = "status"
            
            Public Shared CreateOn As String = "CreateOn"
            
		End Structure
		#End Region
	End Class
End Namespace
