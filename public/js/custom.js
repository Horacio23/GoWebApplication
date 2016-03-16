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
      phone: {
        number: true,
        minlength:10
      },
      zip: {
        required: true,
        minlength:5
      },
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
  $.ajax("/client/"+id).done(function(data){
      console.log("yay, we got the data");
      console.log(data);
      $("#tempModal").html(data);
      $("#clientModal").modal("show");
    }).fail(function(){
      alert("Your fault");
    })
}
