<html>
<script language="JavaScript">

// Get the window object where the context menu was opened.
var oWindow = external.menuArguments;

var url = oWindow.top.location;
	//alert(url);
	//m=url.match("/http:/\/\/");
	
	//if(url.match("/http:/\/\/"))
	//{
	//	alert("该网页无法添加 ^@^ ");
	//}
var title = oWindow.top.document.title;

// Get the document object exposed through oWindow.
var oDocument = oWindow.document;

// Get the selection from oDocument.
// in oDocument.
var oSelect = oDocument.selection;

// Create a TextRange from oSelect.
var oSelectRange = oSelect.createRange();
//
//// Get the text of the selection.
var sNewText = oSelectRange.text;

var tobj=oWindow.top.document;
var ce=tobj.createElement("DIV");
tobj.body.insertAdjacentElement("AfterBegin",ce);
ce.name="blogwind_bookmark";
ce.id="blogwind_bookmark";
ce.innerHTML="<div style='position:absolute;left:180;top:120; z-Index:8; width:475px; height:300px;;visibility:show;'><iframe width=475 height=250  MARGINWIDTH=0 MARGINHEIGHT=0 HSPACE=0 VSPACE=0 FRAMEBORDER=0 src='http://localhost/bookmark.aspx?url=" + escape(url) + "&title=" + escape(title) + "&des=" + escape(sNewText) + "' ></iframe><div style='text-align=:center;background-color:white; width:475px'><input type='button' value='取消' onclick='javascript:blogwind_bookmark.parentElement.removeChild(blogwind_bookmark)'><br><br></div></div>";

//void(window.open('http://localhost/write.aspx?con='+sNewText,'_blank','scrollbars=no,width=475,height=575,left=75,top=20,status=no,resizable=yes'));
//showModalDialog('http://localhost/write.aspx',window,'dialogWidth:560px; dialogHeight:500px;help:0;status:0;resizeable:1;');

//
//// If nothing was selected, insert some text.
//if (sNewText.length == 0){
//   oSelectRange.text = "INSERT TEXT";
//}
//
//// Otherwise, convert the selection to uppercase.
//else{
//   oSelectRange.text = sNewText.toUpperCase();
//}
</script>
</html>