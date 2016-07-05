/*jslint browser: true*/
/*global QuickNote, GibberishAES, $, jQuery, alert, console, CKEDITOR, JsSHA*/
function QuickNoteAPI() {
    'use strict';
    this.Endpoint = "localhost:7001";
}

QuickNote = function () {
    'use strict';

    this.Load = function () {
        try {
            CKEDITOR.replace('quick-note-editor');
            CKEDITOR.config.height = '400';   // CSS unit (percent).
            $('#passwordModal').modal('show');
        } catch(err) {
            this.DisplayErr('Error. Please try again later.' + err);
        }
    };

    this.DisplayNote = function () {
        console.log("DisplayNote!!" + window.location.pathname);

        try {
            $.ajax({    url: "http://localhost:7000/g/" + window.location.hash.slice(1),
                        type: "GET",
                        beforeSend: function(xhr){xhr.setRequestHeader('xauthhash', qn.GenerateHashPass(qn.pass));},
                        success: function(data){
                                    console.log("!!!!!resp: " + data);
                                    var payload = GibberishAES.dec(data, qn.pass);
                                    CKEDITOR.instances["quick-note-editor"].setData(payload);
                                },
                        error: function (xhr, ajaxOptions, thrownError) {
                               alert(xhr.status);
                               alert(thrownError);
                           },
                        error: this.DisplayErr
            });
        } catch (err) {
            console.log("SHIT SON", err);
            this.DisplayErr(err);
        }
    };

    this.SaveNote = function () {
        try {
            var xauthhash = this.GenerateHashPass(this.pass),
                payload = this.EncryptNote(CKEDITOR.instances["quick-note-editor"].getData(), this.pass);
                console.log("Save note: CHECKSUM: " + xauthhash + " payload: " + payload);
                $.ajax({    url: "http://localhost:7000/save/",
                            type: "POST",
                            beforeSend: function(xhr){xhr.setRequestHeader('xauthhash', xauthhash);},
                            success: this.SetSaveSuccess,
                            data: payload,
                            error: this.DisplayErr
                });
        } catch (err) {
            this.DisplayErr(err);
        }
    };

    this.EncryptNote = function (payload, pass) {
        console.log("EncryptNote");
        if (pass === undefined || pass === '') {
            console.log("Param error: pass:", pass);
            throw "System Error!";
        }

        if (payload === undefined || payload === '') {
            console.log("Param error: pass:", pass);
            throw "You can't save an empty note!";
        }

        return GibberishAES.enc(payload, pass);
    };

    this.GenerateHashPass = function (pass) {
        console.log("GenerateHashPass");
        if (pass === undefined || pass === '') {
            console.log("Param error: pass:", pass);
            throw "System Error!";
        }
        var shaObj = new JsSHA("SHA-256", "TEXT");
        shaObj.update(pass);
        console.log(shaObj.getHash("HEX"));
        return shaObj.getHash("HEX");
    };

    this.SavePassword = function () {
        this.pass = $("#notePass").val();
        this.DisplayNote();
    };

    this.API = new QuickNoteAPI();
    this.SetSaveSuccess = function(xhr, reqStatus, reqError){
        console.log(xhr);
        $("#statusContainer").html('<div class="alert alert-success"><a href="#" class="close" data-dismiss="alert" aria-label="close">&times;</a><strong>Note saved successfully!</strong></div>');
    }

    this.DisplayErr = function(msg){
        $("#statusContainer").html('<div class="alert alert-danger"><a href="#" class="close" data-dismiss="alert" aria-label="close">&times;</a><strong>Error try again later!</strong></div>');
    }
};
