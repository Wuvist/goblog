
$("#account_type").change(function(){
	if(this.value=="1")
	{
		$("#server_host").val("https://storage.msn.com/storageservice/MetaWeblog.rpc");
		$("#server_host").attr("disabled", true);
	}else if(this.value=="2"){
		$("#server_host").val("Google");
		$("#server_host").attr("disabled", true);
	}else if(this.value=="3"){
		$("#server_host").val("http://myserver.com/blog/xmlrpc.php");
		$("#server_host").attr("disabled", false);
	}else if(this.value=="4"){
		$("#server_host").val("http://myserver.com/xmlrpc.php");
		$("#server_host").attr("disabled", false);
	}else if(this.value=="5"){
		$("#server_host").val("http://twitter.com/statuses/update.xml");
		$("#server_host").attr("disabled", true);
	}else if(this.value=="6"){
		$("#server_host").val("http://api.fanfou.com/statuses/update.xml");
		$("#server_host").attr("disabled", true);
	}
});

$("#submitbtn").click(function(){
	$.post("/rpc/BSPs.ashx",{
		action:"validate",
		server_host:$("#server_host").val(),
		account_type:$("#account_type").val(),
		username:$("#username").val(),
		password:$("#password").val()
	},
		function(xml){
			if(xml=="0")
				window.location.reload();
			else
				alert("Acccount adding fail");
	});
});

function delete_account(id)
{
	$.post("/rpc/BSPs.ashx",{
		action:"delete",
		id:id
	},
		function(xml){
			if(xml=="0")
				$("#account" + id).fadeOut("slow");
			else
				alert("Acccount deleteing fail");
	});
}

function bind_delete_events(selectExp)
{
	$(selectExp).hover(
        function(){
            $(this).find(".x").show();
            $(this).css({backgroundColor:"#E6E6E6"});
        },
        function(){
            $(this).find(".x").hide();
            $(this).css({backgroundColor:"white"});
        }
    );
}

bind_delete_events("#accounts tr");