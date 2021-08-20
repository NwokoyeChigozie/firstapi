

    $(document).ready(function(){
     $("#send_m").click(function(event){
      event.preventDefault();
      var email = $("#email").val();
      var name = $("#name").val();
      var message = $("#message").val();
      var send_m =    "send_m";
	//console.log("email: " + email + " name: " + name+" message: "+ message +" csrf: "+ document.getElementsByName('csrfmiddlewaretoken')[0].value)
          if(name.trim()==''){
            $("#name").css({"background-color": "transparent", "border-color": "red"});
			$(".error_message").html("<p style='' class='alert alert-danger text-center'>Name field is empty</p>");
    }else{
            $("#name").css({"border-color": "green"});
                if(email.trim()==''){
                   $(".error_message").html("<p style='' class='alert alert-danger text-center'>Email field is empty</p>");
                  $("#email").css({"background-color": "transparent", "border-color": "red"});
    }else{
          $("#email").css({"border-color": "green"});
	if(message.trim()==''){
                   $(".error_message").html("<p style='' class='alert alert-danger text-center'>Please, type in your message</p>");
                  $("#message").css({"background-color": "transparent", "border-color": "red"});
	}else{
		//console.log("email: " + email + " name: " + name+" message: "+ message +" csrf: "+ document.getElementsByName('csrfmiddlewaretoken')[0].value)
       $(".error_message").html("<p style='' class='alert alert-primary text-center'>Please wait...</p>");
$.ajax({

    type: "POST",
    url: window.location.protocol + "//" + window.location.host +"/contact",
    contentType: "application/json",
    dataType: "json",
    data: {       
        // csrfmiddlewaretoken: document.getElementsByName('csrfmiddlewaretoken')[0].value,
                  name : name,
				  email : email,
                  message : message
    },
    success: function(response) {
$(".error_message").html(response);
//console.log(response);
        
    },
    error: function(response) {
 $(".error_message").html("<p class='alert alert-danger text-center' style='text-align:center'>An error occured<p>");
     }      
            });


 
          }
	}
      }




     });   
    });