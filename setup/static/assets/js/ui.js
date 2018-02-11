define("ui", function(require){
    var
        $ = require("jquery"),
        api = require("api");

    /** set of all field names */
    var _fieldNames =
        ['BindAddress', 'BindAddress', 'WebhookPath', 'GitURL', 'GitBranch', 'GitSSHKeyPath', 'GitUsername', 'GitPassword', 'GitCloneTimeout', 'GitHubSecret', 'GitLabSecret', 'Debug', 'StaticPath', 'ProxyURL', 'BuildScript', ];

    /** getters for all the fields */
    var _fieldGetters = {
        '_': function(elem){ return elem.val(); }, // default getter
        'GitCloneTimeout': function(elem){ return parseInt(elem.val())* Math.pow(10, 9); },
        'Debug': function(elem) { return elem.is(":checked"); }
    }

    /** gets the form field belonging to the given name */
    var getField = function(name) {
        return $(document.getElementById(name)).eq(0);
    }

    /** gets the value of a field */
    var getValue = function(name) {
        var getter = _fieldGetters.hasOwnProperty(name) ? _fieldGetters[name] : _fieldGetters['_'];
        return getter(getField(name));
    }

    /** reads the configuration as inputted in the form */
    var getConfig = function() {
        var cfg = {};

        _fieldNames.forEach(function(e){
            cfg[e] = getValue(e);
        });

        return cfg;
    }


    /** validates the form and makes appropriate color changes */
    var validateForm = function(){
        api.validate(getConfig(), renderValidation);
    }


    /** saves the form and show the next step of the UI */
    var applyForm = function(){
        api.save(getConfig(), function(success, fields, message){
            if(!success){
                renderValidation(success, fields, message);
            } else {
                $("#step1").addClass("d-none");
                $("#step2").removeClass("d-none");
            }
        })
    }

    var restart = function(){
        api.finish(function(){

            // hide everything
            $(document.body).addClass("d-none");

            // wait for a second and then reload
            window.setTimeout(function(){
                location.href = "/";
            }, 1000);
        })
    }

    /** renders a given validation */
    var renderValidation = function(success, fields, message) {
        var $message = $(document.getElementById("message")).addClass("d-none"); // hide the message
        var $fields = $(_fieldNames.map(getField).map(function(e){return e[0]})).removeClass('is-invalid').removeClass('is-valid'); // remove all previous validation

        // write the message (if it failed)
        if(!success){
            var message = $('<div>').text(message ? message : "Unknown error, please try again").text();
            $message.removeClass("d-none").html(message.replace(/\n/g, '<br/>'))
        }

        // and
        $fields.addClass('is-valid');
        (fields || []).forEach(function(n){
            getField(n).removeClass('is-valid').addClass('is-invalid');
        });
    }

    return {
        'validateForm': validateForm,
        'applyForm': applyForm,
        'restart': restart
    }
});
