/*global
    parseInt, w2ui, getDepMeth, getDepository, $, app, getBUDfromBID, getCurrentBusiness, console,
    form_dirty_alert, getFormSubmitData, formRecDiffer, formRefreshCallBack, addDateNavToToolbar,
    getGridReversalSymbolHTML, dateControlString, w2utils, saveDepositForm, w2confirm,
    delete_confirm_options,
*/

"use strict";

function getDepositInitRecord(BID, BUD){
    var y = new Date();
    return {
        recid:          0,
        check:          0,
        DID:            0,
        BID:            BID,
        BUD:            BUD,
        DEPID:          0,
        DEPName:        "",
        DPMID:          0,
        DPMName:        "",
        Dt:             dateControlString(y),
        FLAGS:          0,
        Amount:         0.0,
        ClearedAmount:  0.0,
        LastModTime:    y.toISOString(),
        LastModBy:      0,
        CreateTS:       y.toISOString(),
        CreateBy:       0,
    };
}


//---------------------------------------------------------------------------------
// buildDepositElements - changes the main view of the program to a grid with
//                variable name svc + 'Grid'
//
//---------------------------------------------------------------------------------
function buildDepositElements() {
    //------------------------------------------------------------------------
    //          deposit Grid
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'depositGrid',
        url: '/v1/deposit',
        multiSelect: false,
        show: {
            toolbar        : true,
            footer         : true,
            toolbarAdd     : true,   // indicates if toolbar add new button is visible
            toolbarDelete  : false,   // indicates if toolbar delete button is visible
            toolbarSave    : false,   // indicates if toolbar save button is visible
            selectColumn   : false,
            expandColumn   : false,
            toolbarEdit    : false,
            toolbarSearch  : false,
            toolbarInput   : false,
            searchAll      : false,
            toolbarReload  : true,
            toolbarColumns : true,
        },
        columns: [
            {field: 'recid',        hidden: true,  caption: 'recid',        size: '40px',  sortable: true},
            {field: 'DID',          hidden: false, caption: 'DID',          size: '50px',  sortable: true, style: 'text-align: center'},
            {field: 'BID',          hidden: false, caption: 'BID',          size: '50px',  sortable: true, style: 'text-align: center'},
            {field: 'BUD',          hidden: false, caption: 'BUD',          size: '50px',  sortable: true, style: 'text-align: center'},
            {field: 'DEPID',        hidden: true,  caption: 'DEPID',        size: '50px',  sortable: true, style: 'text-align: center'},
            {field: 'DEPName',      hidden: false, caption: 'Depository',   size: '80px',  sortable: true, style: 'text-align: center'},
            {field: 'DPMID',        hidden: true,  caption: 'DPMID',        size: '50px',  sortable: true, style: 'text-align: center'},
            {field: 'DPMName',      hidden: false, caption: 'Method',       size: '150px', sortable: true, style: 'text-align: center'},
            {field: 'Dt',           hidden: false, caption: 'Date',         size: '100px', sortable: true, style: 'text-align: center'},
            {field: 'Amount',       hidden: false, caption: 'Amount',       size: '100px', sortable: true, style: 'text-align: right', render: 'money'},
            {field: 'ClearedAmount',hidden: false, caption: 'ClearedAmount',size: '100px', sortable: true, style: 'text-align: right', render: 'money'},
        ],
        onRefresh: function(event) {
            event.onComplete = function() {
                var sel_recid = parseInt(this.last.sel_recid);
                if (app.active_grid == this.name && sel_recid > -1) {
                    if (app.new_form_rec) {
                        this.selectNone();
                    }
                    else{
                        this.select(app.last.grid_sel_recid);
                    }
                }
            };
        },
        onClick: function(event) {
            event.onComplete = function () {
                var yes_args = [this, event.recid];
                var no_args = [this];
                var no_callBack = function(grid) {
                        grid.select(app.last.grid_sel_recid);
                        return false;
                    };
                var yes_callBack = function(grid, recid) {
                        var x = getCurrentBusiness();
                        var Bid = x.value;
                        var Bud = getBUDfromBID(Bid);

                        var getUIInfo = function(bid,x) {
                            return $.get('/v1/uival/' + bid + x );
                        };

                        $.when( getUIInfo(Bid,"/app.depmeth"),
                                getUIInfo(Bid,"/app.Depositories"))
                        .done( function(dpmArgs,depArgs) {
                            if (typeof dpmArgs[0] == 'string') {
                                app.depmeth = JSON.parse(dpmArgs[0]);
                                w2ui.depositForm.get('DPMName').options.items = app.depmeth[Bud];
                            } else if (dpmArgs[0].status != 'success') {
                                w2ui.depositForm.message(dpmArgs[0].message);
                            }

                            if (typeof depArgs[0] == 'string') {
                                app.Depositories = JSON.parse(depArgs[0]);
                                w2ui.depositForm.get('DEPName').options.items = app.Depositories[Bud];
                            } else if (depArgs[0].status != 'success') {
                                w2ui.depositForm.message(depArgs[0].message);
                            }

                            w2ui.depositForm.refresh();
                            app.last.grid_sel_recid = parseInt(recid);
                            grid.select(app.last.grid_sel_recid); // keep highlighting current row in any case
                            var rec = grid.get(recid);
                            var myurl = '/v1/deposit/' + rec.BID + '/' + rec.DID;
                            var urlgrid = '/v1/depositlist/' + rec.BID + '/' + rec.DID;
                            setToDepositForm("depositLayout","depositForm",myurl,urlgrid,700,true);
                        })
                        .fail( function() { console.log('Error getting /v1/uival/' + x.value + '/{app.depmeth | app.Depositories}'); });
                    };

                form_dirty_alert(yes_callBack, no_callBack, yes_args, no_args); // warn user if form content has been changed
            };
        },
        onAdd: function (/*event*/) {
            var yes_args = [this];
            var no_callBack = function() { return false; };
            var yes_callBack = function(grid) {
                // reset it
                app.last.grid_sel_recid = -1;
                grid.selectNone();

                // Insert an empty record...
                var x = getCurrentBusiness(),
                BID=parseInt(x.value),
                BUD = getBUDfromBID(BID),
                f = w2ui.depositForm;
                var record = getDepositInitRecord(BID, BUD);
                f.record = record;
                var getUIInfo = function(bid,x) {
                    return $.get('/v1/uival/' + bid + x );
                };

                $.when( getUIInfo(BID,"/app.depmeth"),
                        getUIInfo(BID,"/app.Depositories"))
                .done( function(dpmArgs,depArgs) {
                    if (typeof dpmArgs[0] == 'string') {
                        app.depmeth = JSON.parse(dpmArgs[0]);
                        w2ui.depositForm.get('DPMName').options.items = app.depmeth[BUD];
                    } else if (dpmArgs[0].status != 'success') {
                        w2ui.depositForm.message(dpmArgs[0].message);
                    }

                    if (typeof depArgs[0] == 'string') {
                        app.Depositories = JSON.parse(depArgs[0]);
                        w2ui.depositForm.get('DEPName').options.items = app.Depositories[BUD];
                    } else if (depArgs[0].status != 'success') {
                        w2ui.depositForm.message(depArgs[0].message);
                    }
                })
                .fail( function() { console.log('Error getting /v1/uival/' + x.value + '/{app.depmeth | app.Depositories}'); });

                f.refresh();
                setToDepositForm('depositLayout', 'depositForm', '/v1/deposit/' + BID + '/0','/v1/depositlist/'+BID+'/0', 700);
            };

            // warn user if form content has been changed
            form_dirty_alert(yes_callBack, no_callBack, yes_args);
        },
   });

    addDateNavToToolbar('deposit');


    //------------------------------------------------------------------------
    //          deposit Form
    //------------------------------------------------------------------------
    $().w2form({
        name: 'depositForm',
        style: 'border: 0px; background-color: transparent;',
        header: 'Deposit Detail',
        url: '/v1/deposit',
        formURL: '/webclient/html/formdeposit.html',

        fields: [
            { field: 'recid',         type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'DID',           type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'BID',           type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'DEPID',         type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'DPMID',         type: 'int',   required: true, html: { page: 0, column: 0 } },
            { field: 'FLAGS',         type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'BUD',           type: 'list',  required: true, options: {items: app.businesses}, html: { page: 0, column: 0 } },
            { field: 'Dt',            type: 'date',  required: true },
            { field: 'DEPName',       type: 'list',  required: true, options:  {items: [], selected: {}} },
            { field: 'DPMName',       type: 'list',  required: true, options:  {items: [], selected: {}} },
            { field: 'Amount',        type: 'money', required: true,  html: { page: 0, column: 0 } },
            { field: 'ClearedAmount', type: 'money', required: false, html: { page: 0, column: 0 } },
            { field: 'LastModTime',   type: 'time',  required: false, html: { page: 0, column: 0 } },
            { field: 'LastModBy',     type: 'int',   required: false, html: { page: 0, column: 0 } },
            { field: 'CreateTS',      type: 'time',  required: false, html: { page: 0, column: 0 } },
            { field: 'CreateBy',      type: 'int',   required: false, html: { page: 0, column: 0 } },
        ],
        toolbar: {
            items: [
                { id: 'btnNotes', type: 'button', icon: 'fa fa-sticky-note-o' },
                //{ id: 'formSave', type: 'button', caption: 'Save', icon: 'fa fa-check'},
                { id: 'bt3', type: 'spacer' },
                { id: 'btnClose', type: 'button', icon: 'fa fa-times' },
            ],
            onClick: function (event) {
                switch(event.target) {
                case 'btnClose':
                    var no_callBack = function() { return false; },
                        yes_callBack = function() {
                            w2ui.toplayout.hide('right',true);
                            w2ui.depositGrid.render();
                        };
                    form_dirty_alert(yes_callBack, no_callBack);
                    break;
                case 'formSave':
                    saveDepositForm();
                }
            },
        },
        actions: {
            save: saveDepositForm,
         },
        onRefresh: function(event) {
            event.onComplete = function() {
                var f = this;
                var r = f.record;
                var header = "Edit Deposit ({0})";
                var bud = r.BUD.text;
                var dpmid = r.DPMID;
                var depid = r.DEPID;

                f.get("DPMName").options.selected = getDepMeth(bud, dpmid);
                f.get("DEPName").options.selected = getDepository(bud, depid);
                formRefreshCallBack(f, "DID", header);
            };
        },
        onChange: function(event) {
            event.onComplete = function() {
                // formRecDiffer: 1=current record, 2=original record, 3=diff object
                var diff = formRecDiffer(this.record, app.active_form_original, {});
                // if diff == {} then make dirty flag as false, else true
                if ($.isPlainObject(diff) && $.isEmptyObject(diff)) {
                    app.form_is_dirty = false;
                } else {
                    app.form_is_dirty = true;
                }
            };
        },
        onSubmit: function(target, data) {
            delete data.postData.record.LastModTime;
            delete data.postData.record.LastModBy;
            delete data.postData.record.CreateTS;
            delete data.postData.record.CreateBy;
            // modify form data for server request
            getFormSubmitData(data.postData.record);
        },
    });

    //------------------------------------------------------------------------
    //  depositListGrid - For new deposits, it lists all that are not
    //                    currently part of a deposit. Any combination can be
    //                    selected to be part of the new deposit.
    //
    //                    For existing deposits, it lists the receipts that
    //                    belong to the deposit.
    //------------------------------------------------------------------------
    $().w2grid({
        name: 'depositListGrid',
        url: '/v1/depositlist',
        multiSelect: false,
        show: {
            toolbar        : true,
            footer         : true,
            toolbarAdd     : false,   // indicates if toolbar add new button is visible
            toolbarDelete  : false,   // indicates if toolbar delete button is visible
            toolbarSave    : false,   // indicates if toolbar save button is visible
            selectColumn   : false,
            expandColumn   : false,
            toolbarEdit    : false,
            toolbarSearch  : false,
            toolbarInput   : false,
            searchAll      : false,
            toolbarReload  : false,
            toolbarColumns : false,
        },
        columns: [
            {field: 'recid',    caption: 'recid',        hidden: true,  size: '40px',  sortable: true  },
            {field: 'Check',    caption: '',             hidden: false, size: '40px',  editable: { type: 'checkbox' } },
            {field: 'reversed', size: '10px', style: 'text-align: center', sortable: true,
                    render: function (record /*, index, col_index*/) {
                        if (typeof record === "undefined") {
                            return;
                        }
                        if ( (record.FLAGS & app.rcptFLAGS.RCPTREVERSED) !== 0 ) { // if reversed then
                            return getGridReversalSymbolHTML();
                        }
                        return '';
                    },
            },
            {field: 'RCPTID',   caption: 'Receipt ID',   hidden: false, size: '80px',  sortable: true, style: 'text-align: right'},
            {field: 'Dt',       caption: 'Date',         hidden: false, size: '80px',  sortable: true, style: 'text-align: right'},
            {field: 'ARID',     caption: 'ARID',         hidden: true,  size: '150px', sortable: false },
            {field: 'AcctRule', caption: 'Account Rule', hidden: true,  size: '150px', sortable: true  },
            {field: 'Amount',   caption: 'Amount',       hidden: false, size: '100px', sortable: true, style: 'text-align: right', render: 'money'},
            {field: 'BID',      caption: 'BUD',          hidden: true,  size: '40px',  sortable: false },
            {field: 'TCID',     caption: 'TCID',         hidden: true,  size: '40px',  sortable: false },
            {field: 'PMTID',    caption: 'PMTID',        hidden: true,                 sortable: false },
            {field: 'PMTName',  caption: 'Payment Type', hidden: false, size: '100px', sortable: true, style: 'text-align: center' },
            {field: 'DocNo',    caption: 'Document No.', hidden: false, size: '100px', sortable: true, style: 'text-align: right'},
            {field: 'Payors',   caption: 'Payors',       hidden: false, size: '200px', sortable: true  },
            {field: 'FLAGS',    caption: 'FLAGS',        hidden: true,  size: '20px',  sortable: false  },
        ],
        onLoad: function(event) {
            event.done(function () {
                if (w2ui.depositListGrid.summary.length === 0) {
                    var x = getCurrentBusiness(),
                    BID=parseInt(x.value),
                    BUD = getBUDfromBID(BID);
                    var rec = {recid: 's-1', DID: 0, BID: BID, BUD: BUD, DEPID: 0, DEPName: "", DPMID: 0, DPMName: "", Dt: null, FLAGS: 0, Amount: 0.0, ClearedAmount: 0.0, LastModTime: null, LastModBy: 0, CreateTS: null, CreateBy: 0, w2ui:{summary: true}, };
                    w2ui.depositListGrid.add(rec);
                }
                calcTotalCheckedReceipts();
            });
        },
        onClick: function(event) {
            event.done(function () {
                if (event.column == 1) {
                    calcTotalCheckedReceipts();
                }
            });
        },
    });

    addDateNavToToolbar('depositList');


    //------------------------------------------------------------------------
    //  depositFormButtons - Save, Save And Add, Delete buttons
    //------------------------------------------------------------------------
    $().w2form({
        name: 'depositFormBtns',
        style: 'border: 0px; background-color: transparent;',
        url: '/v1/deposit',
        formURL: '/webclient/html/formdepositbtns.html',
        fields: [],
        actions: {
            save: saveDepositForm,
            delete: function() {
                var form = this;
                w2confirm(delete_confirm_options)
                .yes(function() {
                    var tgrid = w2ui.depositForm;
                    var params = {cmd: 'delete', formname: form.name, DID: w2ui.depositForm.record.DID };
                    var dat = JSON.stringify(params);
                    form.url = '/v1/deposit/' + w2ui.depositForm.record.BID + '/' + w2ui.depositForm.record.DID;
                    
                    $.post(form.url, dat, null, "json")
                    .done(function(data) {
                        if (data.status === "error") {
                            form.error(w2utils.lang(data.message));
                            return;
                        }
                        w2ui.toplayout.hide('right',true);
                        w2ui.depositGrid.remove(app.last.grid_sel_recid);
                        w2ui.depositGrid.render();
                    })
                    .fail(function(/*data*/){
                        form.error("Delete Account failed.");
                        return;
                    });
                })
                .no(function() {
                    return;
                });
            },
        },
    });

    //-------------------------------------------------------------------------------
    //  depositLayout - The layout to contain the depositForm and depositDetailGrid
    //-------------------------------------------------------------------------------
    $().w2layout({
        name: 'depositLayout',
        padding: 0,
        panels: [
            { type: 'top',     size: 290,   hidden: false, content: 'top',   resizable: true,  style: app.pstyle },
            { type: 'main',    size: '70%', hidden: false, content: 'main',  resizable: true,  style: app.pstyle },
            { type: 'bottom',  size: 50,    hidden: false, content: 'bottom',resizable: false, style: app.pstyle },
            { type: 'left',    size: '30%', hidden: true },
            { type: 'preview', size: 0,     hidden: true,  content: 'PREVIEW'  },
            { type: 'right',   size: 0,     hidden: true }
        ]
    });
}


//-----------------------------------------------------------------------------
// saveDepositForm - pull the checked Receipts, extend the return values
//      and save the form.
// @params
//-----------------------------------------------------------------------------
function saveDepositForm() {
    var rcpts = getCheckedReceipts();
    var f = w2ui.depositForm;
    f.record.DPMID = f.record.DPMName.id;
    f.record.DEPID = f.record.DEPName.id;
    if (typeof f.record.DID == "string" || typeof f.record.DID == "undefined") {
        f.record.DID = 0;
    }
    if (typeof f.record.FLAGS == "string" || typeof f.record.FLAGS == "undefined") {
        f.record.FLAGS = 0;
    }
    if (typeof f.record.ClearedAmount == "string" || typeof f.record.ClearedAmount == "undefined") {
        f.record.ClearedAmount = 0.0;
    }
    f.save({Receipts: rcpts},function (data) {
        if (data.status == 'error') {
            console.log('ERROR: '+ data.message);
            return;
        }
        w2ui.toplayout.hide('right',true);
        app.form_is_dirty = false;// clean dirty flag of form
        app.last.grid_sel_recid  =-1;// clear the grid select recid
        w2ui.depositGrid.render();
    });
}

//-----------------------------------------------------------------------------
// calcTotalCheckedReceipts - go through all the depositListGrid items and
//      total all the checked receipts. Update the Amount column of the
//      summary row with the total.
// @params
//-----------------------------------------------------------------------------
function calcTotalCheckedReceipts() {
    var t = 0.0;
    var grid = w2ui.depositListGrid;
    var checkcol=0;
    var amtcol=0;
    var flgcol=0;
    for (i = 0; i < grid.columns.length; i++) {
        if (grid.columns[i].field === "Check") {checkcol = i;}
        if (grid.columns[i].field === "Amount") {amtcol = i;}
        if (grid.columns[i].field === "FLAGS") {flgcol = i;}
    }

    for (var i = 0; i < grid.records.length; i++) {
        var x = grid.getCellValue(i,checkcol); // this is what is in the checkbox column
        if (typeof x == "boolean" ) {
            var y = grid.getCellValue(i,flgcol) & 0x4;
            if (x && y === 0) {
                t += grid.getCellValue(i,amtcol);
            }
        }
    }
    grid.set('s-1', { Amount: t });
}

//-----------------------------------------------------------------------------
// getCheckedReceipts - go through depositListGrid items and build a list
//      of the RCPTIDs of the select receipts
// @params
//
// @returns
//      a list of selected receipts
//-----------------------------------------------------------------------------
function getCheckedReceipts() {
    var t = [];
    var grid = w2ui.depositListGrid;
    var i=0;
    var checkcol=0;
    var rcptidcol=0;
    var flagscol=0;
    for (i = 0; i < grid.columns.length; i++) {
        if (grid.columns[i].field === "Check") {checkcol = i;}
        if (grid.columns[i].field === "RCPTID") {rcptidcol = i;}
        //if (grid.columns[i].field === "FLAGS") {flagscol = i;}
    }

    for (i = 0; i < grid.records.length; i++) {
        var x = grid.getCellValue(i,checkcol); // this is what is in the checkbox column
        if (typeof x == "boolean" ) {
            if (x) {
                t.push( grid.getCellValue(i,rcptidcol));
            }
        }
    }
    return t;
}

//-----------------------------------------------------------------------------
// createDepositForm - add the grid and form to the statement layout.  I'm not
//      sure why this is necessary. But if I put this grid and form directly
//      into the layout when it gets created, they do not work correctly.
// @params
//-----------------------------------------------------------------------------
function createDepositForm() {
    w2ui.depositLayout.content('top',   w2ui.depositForm);
    w2ui.depositLayout.content('main',  w2ui.depositListGrid);
    w2ui.depositLayout.content('bottom',w2ui.depositFormBtns);
}

//-----------------------------------------------------------------------------
// setToDepositForm - set to the Deposit Form - puts the depositLayout in
//                    toplayout's right content area. Didn't use the general
//                    call in rutil.js because this form requires the layout
//                    and has multiple parts.
// @params
//   sform   = name of the form
//   url     = request URL for the form
//   [width] = optional, if specified it is the width of the form
//   doRequest =
//-----------------------------------------------------------------------------
function setToDepositForm(slayout, sform, url, urlgrid, width, doRequest) {
    // if not url defined then return
    var url_len=url.length > 0;
    if (!url_len) {
        return false;
    }

    // if form not found then return
    var f = w2ui[sform];
    if (!f) {
        return false;
    }

    // if current grid not found then return
    var g = w2ui[app.active_grid];
    if (!g) {
        return false;
    }

    // if doRequest is defined then take false as default one
    if (!doRequest) {
        doRequest = false;
    }
    f.url = url;
    if (typeof f.tabs.name == "string") {
        f.tabs.click('tab1');
    }
    app.new_form_rec = !doRequest;
    app.form_is_dirty = false;

    var right_panel_content = w2ui.toplayout.get("right").content;
    var fc = w2ui[slayout]; // in this case, we're putting the layout into the content area
    w2ui.depositListGrid.url = urlgrid;
    var showForm = function() {
        // if the same content is there, then no need to render toplayout again
        if ( fc !== right_panel_content) {
            w2ui.toplayout.content('right', fc);
            w2ui.toplayout.sizeTo('right', width);
            w2ui.toplayout.render();
        } else {
            fc.refresh();
        }
        $().w2tag();
        w2ui.toplayout.show('right', true);
    };

    if (doRequest) {
        f.request(function(event) {
            if (event.status === "success") {
                showForm();
                return true;
            }
            else {
                showForm();
                f.message("Could not get form data from server...!!");
                return false;
            }
        });
    } else {
        var sel_recid = parseInt(g.last.sel_recid);
        if (sel_recid > -1) {
            g.unselect(g.last.sel_recid); // if new record is being added then unselect {{the selected record}} from the grid
        }
        showForm();
        return true;
    }
}