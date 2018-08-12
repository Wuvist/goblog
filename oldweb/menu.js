function addlink(b) {
	if (document.getElementById("newlink").innerHTML=='')
	{
		document.getElementById("newlink").innerHTML = ' <div id="addlink"><FORM action="/savelink.aspx" method="post">' + word_url + '：<INPUT class="input" type="text" maxLength="100" size="40" value="http://" name="url">' + word_public + '：<INPUT type="checkbox" CHECKED name="reveal"><BR>' + word_title + '：<INPUT class="input" type="text" maxLength="50" size="40" name="link"> <INPUT class="logsub" type="submit" value="' + word_save + '"></FORM>';
	}
	else
	{document.getElementById("newlink").innerHTML='';}

}
function shows(link) {
document.getElementById("viewfriend").innerHTML = '<IFRAME width=580 height=400 SRC="/friendlist.aspx?blogger=' + link+ '" MARGINWIDTH=0 MARGINHEIGHT=0 HSPACE=0 VSPACE=0 FRAMEBORDER=0></IFRAME>';
}
function hides(link) {
document.getElementById("viewfriend").innerHTML = '';
document.getElementById("newlink").innerHTML = '';
}
function opens(link) {
window.open('/' + link,'_blank');
}