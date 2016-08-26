$(document).ready(function(){
  console.log("inside validate");
  $("#clientForm").validate({
    rules: {
      firstName: "required",
      lastName: "required",
      // address: "required",
      // city: "required",
      // state:{
      //   required: true,
      //   minlength:2
      // },
      phone: {
        number: true,
        required: true,
        minlength:10
      },
      // zip: {
      //   required: true,
      //   minlength:5
      // },
      lastTransaction: "required",
      entranceDate: "required",
      transactionDate: "required"

    }
  });
});

$("#phone, #zip").keypress(function(e){
  console.log("Inside phone or zip keypress");
  if (e.which < 48 || e.which > 57){
      return false;
  }
});

$("#state").keypress(function(e){
  e = e || window.event;
    var charCode = (typeof e.which == "undefined") ? e.keyCode : e.which;
    var charStr = String.fromCharCode(charCode);
    if (/\d/.test(charStr)) {
        return false;
    }
});

function deleteConfirmation(id){
  $("#deleteModalContainer").html(`
    <div class="modal fade" id="deleteModal" tabindex="-1" role="dialog">
      <div class="modal-dialog">
        <div class="modal-content">
          <div class="color-line"></div>
          <div class="modal-title">

              <h3 class="center">Are you sure you want to delete this client?</h3>

          </div>
          <div class="modal-body"></div>
          <div class="modal-footer mymodal">
            <a href="/delete/`+id+`" class="btn btn-danger">Delete</a>
            <button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
          </div>
        </div><!-- /.modal-content -->
      </div><!-- /.modal-dialog -->
    </div>
  `);
  $("#deleteModal").modal("show");
}

function getClient(id){

  $.ajax({
     url: "/client/"+id,
     type: "GET",
     async: false,
     success: function(data){
         if(data.includes('<meta')){
              window.location.replace("/login");
         }else{
             $("#tempModal").html(data);
             $("#clientModal").modal("show");
         }
     },
     fail: function(){
         alert("Your fault");
     }
  });
}

$('#clientForm').submit(function(event){
    event.preventDefault();

    var formData = new FormData($(this)[0])
    formData.append("lastTransaction", $(".selectpicker").selectpicker('val'))
    var url = $('#formAction').val();
    console.log(url);
    $.ajax({
    url: url,
    type: 'POST',
    data: formData,
    async: false,
    cache: false,
    contentType: false,
    processData: false,
    success: function (returndata) {
        console.log(returndata)
      window.location.replace("/clients");
    }
  });

  return false;
})

function getTransactions(){
    $.ajax({
        url:"transactions",
        type: "GET",
        success: function(data){
            //todo: Populate the graphical stuff inside this function
            console.log(data);
        },
        error: function(error){
            console.log(error);
        }
    });
}
