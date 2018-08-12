Public Class langs
    Enum lang
        zh_cn = 0
        en = 1
        de = 2
        jp = 3
    End Enum

    Public Shared defaultLang As lang = lang.zh_cn

    Public Shared Function getLangString(ByVal l As lang) As String
        Select Case l
            Case lang.zh_cn
                Return "zh-CN"
            Case lang.en
                Return "en"
            Case lang.jp
                Return "ja"
            Case lang.de
                Return "de"
        End Select
    End Function

    Public Shared Function getLang(ByVal cookie As String, ByVal userlangs() As String) As lang
        Dim r As lang = defaultLang
        Dim flag As Boolean = False

        'If (Not cookie Is Nothing) AndAlso cookie <> "" Then
        '  Try
        '    i = CInt(cookie)
        '    r = i
        '    Return r
        '  Catch ex As Exception

        '  End Try
        'End If

        Dim s As String
        If userlangs Is Nothing Then Return defaultLang
        For Each s In userlangs
            If Not s Is Nothing Then
                If s = "zh-cn" Then Return lang.zh_cn
                If s.StartsWith("en") Then Return lang.en
                If s.StartsWith("ja") Then Return lang.jp
                If s.StartsWith("de") Then Return lang.de
            End If
        Next
        Return defaultLang
    End Function
End Class

