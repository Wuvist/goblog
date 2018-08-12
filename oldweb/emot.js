function v(s) {
	if (current_emot==s)
	{
		current_emot="";
		e.innerHTML='';
	}
	else
	{
		current_emot=s;
		vv(s);
	}
}
function vv(s) {
	var str,t;
	str="";
	if(s=='QQ')
	{		
		for (i=1;i<81;i++ )
		{
			str+='<a href="#" onclick=w("QQ';
			if (i<10)
			{
				t='0';
			}
			else
				t='';
			str += t + new String(i) + '");><img border="0" src="images/emot/QQ';
			str += t + new String(i) + '.gif"></a> ';
		}
		e.innerHTML =str;
	}
	if(s=='flag')
	{
		for (i=1;i<29;i++ )
		{
			str+='<a href="#" onclick=w("flag';
			if (i<10)
			{
				t='0';
			}
			else
				t='';
			str += t + new String(i) + '");><img border="0" src="images/emot/flag';
			str += t + new String(i) + '.gif"></a> ';
		}
		e.innerHTML =str;
	}
	if(s=='emot')
	{
		for (i=1;i<29;i++ )
		{
			str+='<a href="#" onclick=w("em';
			if (i<10)
			{
				t='0';
			}
			else
				t='';
			str += t + new String(i) + '");><img border="0" src="images/emot/em';
			str += t + new String(i) + '.gif"></a> ';
		}
		e.innerHTML =str;
	}
	if(s=='neko')
	{
		for (i=1;i<129;i++ )
		{
			str+='<a href="#" onclick=w("neko';
			str += new String(i) + '");><img border="0" src="images/emot/neko';
			str += new String(i) + '.gif"></a> ';
		}
		e.innerHTML =str;
	}
	if(s=='face')
	{
		for (i=1;i<56;i++ )
		{
			str+='<a href="#" onclick=w("face/';
			str += new String(i) + '");><img border="0" src="images/emot/face/';
			str += new String(i) + '.gif"></a> ';
		}
		e.innerHTML =str;
	}
}