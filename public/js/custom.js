$(document).ready(function(){
  console.log("inside validate");
  $("#clientForm").validate({
    rules: {
      firstName: "required",
      lastName: "required",
      address: "required",
      city: "required",
      state:{
        required: true,
        minlength:2
      },
      zip: {
        required: true
      },
      entranceDate: "required",
      transactionDate: "required"

    }
  });
});

$("#phone","#zip").keypress(function(event){
  if (event.which < 48 || event.which > 57)
    {
        event.preventDefault();
    }
});

$("#phone").keypress(function(e){
  var phoneField = $("#phone");
  var phone = $("#phone").val();
  console.log(phone.length)


  if(phone.length == 3 || phone.length == 7){
    phone=phone+"-";
    console.log(phone)
    phoneField.val(phone);
  }

})

$("#state").keypress(function(e){
  e = e || window.event;
    var charCode = (typeof e.which == "undefined") ? e.keyCode : e.which;
    var charStr = String.fromCharCode(charCode);
    if (/\d/.test(charStr)) {
        return false;
    }
});
