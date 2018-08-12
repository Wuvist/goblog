	var userpictemp;
function icon_toggle(me) {
	if (me.value=="启用头像")
	{
		userpicform.action.value="show";
		userpicform.submit();
	}else
	{
		userpicform.action.value="close";
		document.getElementById("userpic").innerHTML='';
		userpicform.submit();
	}

}
function icon_copy(blogger) {
	userpicform.action.value="copy:" + blogger;
	userpic.innerHTML='';
	userpicform.submit();
}
function icon_change() {
	if(document.getElementById("IconBtn").value=="取消修改")
	{
		document.getElementById("userpic").innerHTML=userpictemp;
		document.getElementById("IconBtn").value="修改头像";
	}
	else
	{
		userpictemp=document.getElementById("userpic").innerHTML;
		document.getElementById("IconBtn").value="取消修改";
		document.getElementById("userpic").innerHTML='<IFRAME width=150 height=103 SRC="/userpic_change.aspx" MARGINWIDTH=0 MARGINHEIGHT=0 HSPACE=0 VSPACE=0 FRAMEBORDER=0 allowtransparency></IFRAME>';
	}
}