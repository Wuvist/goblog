Imports sqlconnhelper

Public Class Categories
    Public Shared Function GetBloggerCategoryIdByName(ByVal BloggerId As Integer, ByVal CategoryName As String) As Integer
        Dim pa As New ParaHelper
        pa.addInt("blogger", BloggerId)
        pa.addVarchar("cate", CategoryName, 50)
        Return connHelper.exeInteger("select [index] from UserDefineCategory where blogger=@blogger and cate=@cate", CommandType.Text, pa.parameters, False)
    End Function

    Public Shared Function GetBloggerDefaultCategoryId(ByVal BloggerId As Integer)
        Dim pa As New ParaHelper
        pa.addInt("blogger", BloggerId)
        Return connHelper.exeInteger("select top 1 [index] from UserDefineCategory where blogger=@blogger order by [index]", CommandType.Text, pa.parameters, False)

    End Function
End Class
