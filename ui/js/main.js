/*jslint browser: true*/
/*global GibberishAES, $, jQuery, alert, console*/

var API = (function () {
    'use strict';

    var enc, dec, shaObj;
    enc = GibberishAES.enc("This sentence is super secret", "ultra-strong-password");
    console.log(enc);
    dec = GibberishAES.dec(enc, "ultra-strong-password");
    console.log(dec);

    shaObj = new jsSHA("SHA-256", "TEXT");
    console.log(shaObj.getHash("HEX"));
    var QuickNote = function () {
        var editorElementId = "quick-note-editor";
        var apiEndpoint = "";
        var self = this;

    Load = function(){
        CKEDITOR.replace( this.editorElementId);
        CKEDITOR.config.height = '400';   // CSS unit (percent).
        CKEDITOR.instances[this.editorElementId].setData('Start typing your note here');

        jQuery('#saveBtn').click(this.SaveNote);
    };

    DisplayNote = function(){
        console.log("DisplayNote!!" + window.location.pathname);

    };

    SaveNote = function(){
        console.log("SaveNote!!!!!1");
        console.log("this: ", this);
        console.log("self: ", self);
        self.EncryptNote();
    };

    EncryptNote = function(){
        console.log("EncryptNote!");
        console.log("QuickNoteAPI: ", API.Endpoint)
    };

    GenerateHashPass = function(){

    };
    API = new QuickNoteAPI();
};

function QuickNoteAPI(){
    this.Endpoint = "localhost:7001";

}


console.log(QuickNote);
QuickNote.Load();

}());
