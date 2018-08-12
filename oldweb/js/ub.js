moods=new Array();
moods[0]='饿';moods[1]='离';moods[2]='哭';moods[3]='食';moods[4]='爽';moods[5]='读';moods[6]='拼';moods[7]='泣';moods[8]='闷';moods[9]='羞';moods[10]='暑';moods[11]='烦';moods[12]='哀';moods[13]='拽';moods[14]='囧';moods[15]='穷';moods[16]='睡';moods[17]='空';moods[18]='乐';moods[19]='困';moods[20]='闲';moods[21]='恼';moods[22]='懒';moods[23]='愤';moods[24]='冷';moods[25]='湿';moods[26]='贺';moods[27]='思';moods[28]='怒';moods[29]='聊';moods[30]='厌';moods[31]='忙';moods[32]='吵';moods[33]='泪';moods[34]='热';moods[35]='病';moods[36]='累';moods[37]='乱';moods[38]='吓';moods[39]='赞';moods[40]='美';moods[41]='喜';moods[42]='怨';moods[43]='憋'; moods[44]='渴';moods[45]='伤';moods[46]='痛';moods[47]='幽';moods[48]='静'; moods[49]='恨';moods[50]='凶';moods[51]='爱';

noun_type_mood = new CmdUtils.NounType( "心情",
  ['饿','离','哭','食','爽','读','拼','泣','闷','羞','暑','烦','哀','拽','囧','穷','睡','空','乐','困','闲','恼','懒','愤','冷','湿','贺','思','怒','聊','厌','忙','吵','泪','热','病','累','乱','吓','赞','美','喜','怨','憋','渴','伤','痛','幽','静','恨','凶','爱']
  );


CmdUtils.CreateCommand({
  name: "windmood",
  takes: {"心情": noun_type_mood},
  modifiers: {msg: noun_arb_text},

  homepage: "http://www.blogwind.com/",
  author: {name: "Wuvist", homepage: "http://www.blogwind.com/Wuvist/"},


  preview: function( pblock, mood, mods) {
    pblock.innerHTML = "新心情: 【" + mood.text + "】" + mods.msg.text;
  },
  execute: function(mood, mods) {
	var m, msg;
	for (i=0;i<moods.length ;i++ )
	{
		if(moods[i]==mood.text)break;
	}
	m = i;
	msg = mods.msg.text;
	jQuery.post("http://www.blogwind.com/rpc/moods.aspx",{
	mood:m,
	msg:msg
	  },function(xml){

	});
  }
})
