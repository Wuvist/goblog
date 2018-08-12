Public Class util
    Public Shared Function GetCTDes(ByVal ct As ini.CommentorType) As String
        Select Case ct
            Case ini.CommentorType.Blogger
                Return "作者"
            Case ini.CommentorType.BloggerFans
                Return "关心作者的用户"
            Case ini.CommentorType.RegisterUser
                Return "博客风用户"
            Case ini.CommentorType.BloggerFriend
                Return "作者的朋友"
            Case ini.CommentorType.BloggerGoodFriend
                Return "作者的知己"
            Case ini.CommentorType.UnrevealUser
                Return "低调的博客风用户"
            Case Else
                Return "游客"
        End Select
    End Function

    Public Shared Function GetPagingSQL(ByVal sql As String, ByVal OrderSQL As String, ByVal start As Integer, ByVal count As Integer) As String
        Dim i As Integer
        Dim r As String
        i = sql.IndexOf("from")

        r = "select * from (" & sql.Substring(0, i) & ",ROW_NUMBER() over (order by " & OrderSQL & ") as rowNum "
        r &= sql.Substring(i) & ") as mytable where rowNum between " & Str(start) & " and " & Str(start + count)

        Return r
    End Function

    Public Shared Function GetPagingSQL(ByVal sql As String, ByVal Start As Integer, ByVal Count As Integer) As String
        Dim i, j As Integer
        Dim r As String
        i = sql.IndexOf("from ")
        If i < 0 Then
            i = sql.IndexOf("FROM ")
        End If

        Dim OrderSQL As String
        j = sql.IndexOf("order by", System.StringComparison.OrdinalIgnoreCase)
        OrderSQL = sql.Substring(j)
        sql = sql.Substring(0, j)

        r = sql.Substring(0, i) & " from (" & sql.Substring(0, i) & ",ROW_NUMBER() over (" & OrderSQL & ") as rowNum "
        r &= sql.Substring(i) & ") as mytable where rowNum between " & Str(Start + 1) & " and " & Str(Start + Count)

        Return r
    End Function

    Public Shared Function IsIP(ByVal ip As String) As Boolean
        Return Regex.IsMatch(ip, "^((2[0-4]\d|25[0-5]|[01]?\d\d?)\.){3}(2[0-4]\d|25[0-5]|[01]?\d\d?)$")
    End Function

    Public Shared Function GetIP() As String
        Dim result As String = [String].Empty

        result = HttpContext.Current.Request.ServerVariables("HTTP_X_FORWARDED_FOR")
        If result Is Nothing OrElse result = [String].Empty Then
            result = HttpContext.Current.Request.ServerVariables("REMOTE_ADDR")
        End If

        If result Is Nothing OrElse result = [String].Empty Then
            result = HttpContext.Current.Request.UserHostAddress
        End If

        If result Is Nothing OrElse result = [String].Empty OrElse Not util.IsIP(result) Then
            Return "0.0.0.0"
        End If

        Return result
    End Function
    Public Shared Function IsInternalIP(ByVal IP As String) As Boolean
        If IP.StartsWith("10.") Then Return True
        If IP.StartsWith("192.168.") Then Return True
        If IP.StartsWith("172.16.") Then Return True
        If IP.StartsWith("172.17.") Then Return True
        If IP.StartsWith("172.18.") Then Return True
        If IP.StartsWith("172.19.") Then Return True
        If IP.StartsWith("172.20.") Then Return True
        If IP.StartsWith("172.21.") Then Return True
        If IP.StartsWith("172.22.") Then Return True
        If IP.StartsWith("172.23.") Then Return True
        If IP.StartsWith("172.24.") Then Return True
        If IP.StartsWith("172.25.") Then Return True
        If IP.StartsWith("172.26.") Then Return True
        If IP.StartsWith("172.27.") Then Return True
        If IP.StartsWith("172.28.") Then Return True
        If IP.StartsWith("172.29.") Then Return True
        If IP.StartsWith("172.30.") Then Return True
        If IP.StartsWith("172.31.") Then Return True

        Return False
    End Function

    Public Shared Function CountStr(ByVal value As String, ByVal keyword As String) As Integer
        Dim i, j As Integer
        Do
            i = value.IndexOf(keyword, i)

            If i > -1 Then
                j += 1
                i += 1
            End If
        Loop Until i < 0
        Return j
    End Function

    Public Shared Function AddHyperLink(ByVal value As String) As String
        Dim r As New System.Text.RegularExpressions.Regex("(http://\S+)")
        Return r.Replace(value, "<a href='$1'>$1</a>")
    End Function

    Public Shared Function GetPublicIP() As String
        Dim result As String = [String].Empty

        result = HttpContext.Current.Request.ServerVariables("HTTP_X_FORWARDED_FOR")
        If result Is Nothing OrElse result = [String].Empty OrElse IsInternalIP(result) Then
            result = HttpContext.Current.Request.ServerVariables("REMOTE_ADDR")
        End If

        If result Is Nothing OrElse result = [String].Empty OrElse IsInternalIP(result) Then
            result = HttpContext.Current.Request.UserHostAddress
        End If

        If result Is Nothing OrElse result = [String].Empty OrElse Not util.IsIP(result) Then
            Return "0.0.0.0"
        End If

        Return result
    End Function

    Public Shared Function MD5(ByVal plain As String) As Byte()
        Dim md5cp As New System.Security.Cryptography.MD5CryptoServiceProvider
        Dim data() As Byte
        data = System.Text.Encoding.UTF8.GetBytes(plain)
        Return md5cp.ComputeHash(data)
    End Function

    Public Shared Function Hex2Binary(ByVal msg As String) As Byte()
        Dim i As Integer
        Dim bin(msg.Length / 2 - 1) As Byte
        For i = 0 To bin.Length - 1
            bin(i) = Int32.Parse(msg.Substring(i * 2, 2), System.Globalization.NumberStyles.HexNumber)
        Next
        Return bin
    End Function

    Public Shared Function UTF2GB2312(ByVal msg As String) As String
        Dim gb As System.Text.Encoding
        gb = System.Text.Encoding.GetEncoding("gb2312")
        Return System.Text.Encoding.UTF8.GetString(System.Text.Encoding.Convert(System.Text.Encoding.UTF8, gb, System.Text.Encoding.UTF8.GetBytes(msg)))
    End Function

    Public Shared Function ToHexString(ByVal data As Byte()) As String
        Dim s As New StringBuilder(data.Length * 2)
        Dim t As String
        For Each b As Byte In data
            t = Hex(b)
            If t.Length = 1 Then
                s.Append("0")
            End If
            s.Append(t)
        Next
        Return s.ToString
    End Function

    Public Shared Function JoinStr(ByVal ar As ArrayList, ByVal joiner As String, Optional ByVal startIndex As Integer = 0) As String
        Dim i As Integer
        Dim sb As New StringBuilder
        For i = startIndex To ar.Count - 2
            sb.Append(ar.Item(i))
            sb.Append(joiner)
        Next
        sb.Append(ar.Item(ar.Count - 1))
        Return sb.ToString
    End Function

    Public Shared Function JoinStr(ByVal ar As String(), ByVal joiner As String, Optional ByVal startIndex As Integer = 0) As String
        Dim i As Integer
        Dim sb As New StringBuilder
        For i = startIndex To ar.Length - 2
            sb.Append(ar(i))
            sb.Append(joiner)
        Next
        sb.Append(ar(ar.Length - 1))
        Return sb.ToString
    End Function
End Class
