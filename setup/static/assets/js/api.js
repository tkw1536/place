define("api", {
    /** validate a given configuration with the server */
    validate: function (config, callback) {
        // callback(success, fields, message)
        $.post("/setup/validate", JSON.stringify(config), function(data){
            callback(data["success"], data["fields"], data["message"]);
        });
    }, 

    /** Save the configuration into the server */
    save: function(config, callback){
        // callback(success, fields, message)
        $.post("/setup/save", JSON.stringify(config), function(data){
            callback(data["success"], data["fields"], data["message"]);
        });
    }, 

    /** finish the configuration and reload the server */
    finish: function (callback){
        // callback() after request
        $.post("/setup/finish", function() {
            callback();
        });
    }
});