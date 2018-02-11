requirejs.config({
    paths: {
        "jquery": "../lib/jquery/jquery-3.3.1.min",
        "bootstrap": "../lib/bootstrap/js/bootstrap.min",
        "popper": "../lib/popper/popper.min"
    },
    shim: {
        "bootstrap": {
            deps: ["jquery", "popper.js"]
        }
    }
});

// bootstrap requires 'popper.js' instead of 'popper'
define('popper.js', function(module){
    window.Popper = require('popper');
})

require(["ui", "jquery", "bootstrap"], function(ui, $, bs){

    // trigger validation on moving to any of the next fields
    $("#form1").find("input").blur(function(){ui.validateForm(); });

    // on submit of the first form, apply the changes
    $("#form1").submit(function(e){
        e.preventDefault();
        ui.applyForm();
    });

    // on submit of the second form, reload the page
    $("#form2").submit(function(e){
      e.preventDefault();
      ui.restart();
    })

    // validate now!
    ui.validateForm();
});
