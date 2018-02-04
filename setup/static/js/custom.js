$(function(){
    var textarea = $("textarea");
    
    $("form").submit(function(e){
        e.preventDefault();

        jQuery.post("/raw", textarea.val())
        .done(function() {
            alert( "Done" );

            jQuery.get("/finalize")
            .always(function(){
                location.href = "/";
            })
        })
        .fail(function(e) {
            alert("Error: "+e.responseText );
        })
    })  
})