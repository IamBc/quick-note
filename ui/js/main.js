/*jslint browser: true*/
/*global QuickNote, GibberishAES, $, jQuery, alert, console, CKEDITOR, JsSHA*/
function QuickNoteAPI() {
    'use strict';
    this.Endpoint = "localhost:7001";
}

QuickNote = function () {
    /*
    enc = GibberishAES.enc("This sentence is super secret", "ultra-strong-password");
    console.log(enc);
    dec = GibberishAES.dec(enc, "ultra-strong-password");
    console.log(dec);

    shaObj = new JsSHA("SHA-256", "TEXT");
    console.log(shaObj.getHash("HEX"));

    console.log("this: ", this);*/
    'use strict';

    this.Load = function () {
        CKEDITOR.replace('quick-note-editor');
        CKEDITOR.config.height = '400';   // CSS unit (percent).
        CKEDITOR.instances[this.editorElementId].setData('Start typing your note here');

        jQuery('#saveBtn').click(this.SaveNote);
    };

    this.DisplayNote = function () {
        console.log("DisplayNote!!" + window.location.pathname);
    };

    this.SaveNote = function () {
        console.log("SaveNote!!!!!1");
        console.log("this: ", this);
    };

    this.EncryptNote = function () {
        console.log("EncryptNote!");
        console.log("QuickNoteAPI: ", API.Endpoint);
    };

    this.GenerateHashPass = function () {
        console.log("foobar");
    };
    this.API = new QuickNoteAPI();
};
