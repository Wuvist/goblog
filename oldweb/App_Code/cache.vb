Imports SqlConnHelper

Namespace Blogwind
    Public Class ch
        Public Class item
            Public lastupdate As Date
            Public inteval As Integer = 30
            Public data As Object
        End Class

        Public Shared ht As Hashtable = New Hashtable

        Public Shared Function getbloggerbyHits() As DataTable
            Dim it As item
            If ht.Contains("getbloggerbyHits") Then
                it = ht.Item("getbloggerbyHits")
                If Now.Subtract(it.lastupdate).Minutes > it.inteval Then
                    it.data = connHelper.retrieve("select_blogger_by_hits", CommandType.StoredProcedure, Nothing)
                    it.lastupdate = Now
                End If
            Else
                it = New item
                it.lastupdate = Now
                it.data = connHelper.retrieve("select_blogger_by_hits", CommandType.StoredProcedure, Nothing)
                ht.Add("getbloggerbyHits", it)
            End If
            Return it.data
        End Function

        Public Shared Function getNewblogger() As DataTable
            Dim it As item
            If ht.Contains("getNewblogger") Then
                it = ht.Item("getNewblogger")
                If Now.Subtract(it.lastupdate).Minutes > it.inteval Then
                    it.data = connHelper.retrieve("select_new_Blog", CommandType.StoredProcedure, Nothing)
                    it.lastupdate = Now
                End If
            Else
                it = New item
                it.lastupdate = Now
                it.data = connHelper.retrieve("select_new_Blog", CommandType.StoredProcedure, Nothing)
                ht.Add("getNewblogger", it)
            End If
            Return it.data
        End Function
    End Class

End Namespace
