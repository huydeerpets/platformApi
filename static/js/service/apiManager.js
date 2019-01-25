var titles = new Array("参数名称", "参数类型", "数据类型", "是否必填", "描述");
var formTitles = new Array("数据名称", "数据类型", "是否必填", "描述");
var typeData = null;
var deleteIds = ",";
var deleteReqSampleIds = ",";
var deleteReqCodeIds = ",";
var deleteFormDataIds = ",";
var deleteRespIds = ",";

$(document).ready(function () {
    if ($("#nodeLevel").val() != "4") {
        $("#name").val('');
        //默认添加一行
        addRow();
    }

    //加载基础数据
    loadBaseData();

    loadRequestParams();

    setBlurEvent();
});

function String2Hex(tmp) {
    var str = '';
    for (var i = 0; i < tmp.length; i++) {
        str += tmp[i].charCodeAt(0).toString(16);
    }
    return str;
}

function setBlurEvent() {
    $("#paramTable>tbody").on("blur", "input:text", function () {
        $(this).css("background", "white");
    });
    $("#formDataTable>tbody").on("blur", "input:text", function () {
        $(this).css("background", "white");
    });
    $("#respTable>tbody").on("blur", "input:text", function () {
        $(this).css("background", "white");
    });
    $("#paramTable>tbody").on("change", ".arg-type", function () {
        $(this).css("background", "white");
    });
    $("#paramTable>tbody").on("change", ".data-type", function () {
        $(this).css("background", "white");
    });
    $("#formDataTable>tbody").on("change", ".data-type", function () {
        $(this).css("background", "white");
    });

    $("#req_field_content").on("blur", "input:text", function () {
        $(this).css("background", "white");
    });
    $("#req_field_content").on("change", ".data-type", function () {
        $(this).css("background", "white");
    });

    $("#respTable>tbody").on("change", ".data-type", function () {
        $(this).css("background", "white");
        if ($(this).val() == "5" || $(this).val() == "6") {//array,object
            var trObj = $(this).parent().parent()[0];
            var btnObj = $(trObj).find("td:last").find(".btn-info");
            if (!btnObj || btnObj.length == 0) {
                var tableId = $(trObj).parent().parent().attr("id");//tbody-table
                var rowIndex = $(trObj).index();
                var btnHtml = '<button type="button" class="btn btn-xs btn-info" onclick="addRespChildTable(\'' + tableId + '\',' + rowIndex + ',this)">添加子项</button>';
                var tds = $(trObj).children();
                $(tds[tds.length - 1]).append(btnHtml);
            }
        } else {
            var trObj = $(this).parent().parent()[0];
            var btnObj = $(trObj).find("td:last").find(".btn-info");
            if (btnObj && btnObj.length>0) {
                btnObj.remove();
            }
        }
    });

    $("#sample_ul").on('click', '.tab-close', function (ev) {
        var ev = window.event || ev;
        ev.stopPropagation();
        //先判断当前要关闭的tab选项卡有没有active类，再判断当前选项卡是否最后一个，如果是则给前一个选项卡以及内容添加active，否则给下一个添加active类
        var len = $("#sample_ul").find("li").length;
        var liObj = $(this).parent().parent();
        var aObj = $(this).parent();
        if (liObj.hasClass('active')) {
            if (liObj.index() == len - 2) {
                liObj.prev().addClass('active');
                $(aObj.attr('href')).prev().addClass('active');
            } else {
                liObj.next().addClass('active');
                $(aObj.attr('href')).next().addClass('active');
            }
        }

        var id = $(liObj).find("#reqSampleId" + liObj.index()).val();
        if (deleteReqSampleIds.indexOf("," + id + ",") == -1) {
            deleteReqSampleIds += id + ",";
        }

        liObj.remove();
        $(aObj.attr('href')).remove();

        //重新设置
        var liArr = $("#sample_ul").find("li");
        var contentArr = $("#sample_content").find(".tab-pane");
        len = liArr.length;
        for (var i = 0; i < len - 1; i++) {
            var id = "reqSample" + i;
            $(liArr[i]).find("a").attr("href", "#" + id);
            $(liArr[i]).find("a").attr("aria-controls", id);
            $(contentArr[i]).attr("id", id);
            $(contentArr[i]).find(":input").attr("id", "req_sample" + i);
            $(contentArr[i]).find(":input").attr("name", "req_sample" + i);
        }
    });

    $("#code_ul").on('click', '.tab-close', function (ev) {
        var ev = window.event || ev;
        ev.stopPropagation();
        //先判断当前要关闭的tab选项卡有没有active类，再判断当前选项卡是否最后一个，如果是则给前一个选项卡以及内容添加active，否则给下一个添加active类
        var len = $("#code_ul").find("li").length;
        var liObj = $(this).parent().parent();
        var aObj = $(this).parent();
        if (liObj.hasClass('active')) {
            if (liObj.index() == len - 2) {
                liObj.prev().addClass('active');
                $(aObj.attr('href')).prev().addClass('active');
            } else {
                liObj.next().addClass('active');
                $(aObj.attr('href')).next().addClass('active');
            }
        }

        var id = $(liObj).find("#reqCodeId" + liObj.index()).val();
        if (deleteReqCodeIds.indexOf("," + id + ",") == -1) {
            deleteReqCodeIds += id + ",";
        }

        liObj.remove();
        $(aObj.attr('href')).remove();

        //重新设置
        var liArr = $("#code_ul").find("li");
        var contentArr = $("#code_content").find(".tab-pane");
        len = liArr.length;
        for (var i = 0; i < len - 1; i++) {
            var id = "reqCode" + i;
            $(liArr[i]).find("a").attr("href", "#" + id);
            $(liArr[i]).find("a").attr("aria-controls", id);
            $(contentArr[i]).attr("id", id);
            $(contentArr[i]).find(":input").attr("id", "req_code" + i);
            $(contentArr[i]).find(":input").attr("name", "req_code" + i);
        }
    });

    //修改标题名称
    $("#sample_ul").on('dblclick', 'a', function (ev) {
        var ev = window.event || ev;
        ev.stopPropagation();
        var that = this;
        var name = $(that).text();
        layer.prompt({
            title: '您可以对【' + name + '】重新命名',
            formType: 0,
            value: name,
            btn: ["确定", "取消"]
        }, function (text, index) {
            if (text && text != "") {
                $(that).contents()[0].textContent = text;
                //$(that).text(text);
                layer.close(index);
            };
        });

    });

    setSubmitBtnEvent();
}



function loadBaseData() {
    //$("#method").empty()
    //$("#method").append("<option value=''>-=请选择=-</option>");
    var langType = $("#langType").val();
    $.ajax({
        url: '/BaseData/GetBaseDataByTypeExt/request_method,data_type,req_arg_type,comp_lang/' + langType,
        async: false,//同步，会阻塞操作
        type: 'GET',//PUT DELETE POST
        success: function (result) {
            if (result.Status == 0) {
                typeData = result.Data;
                setReqMethodSelectValue();
                setCompLangSelectValue();
                setSelectValue(0, 'paramTable', 'req_arg_type');
                setSelectValue(0, 'paramTable', 'data_type');
            }
        }, error: function () {
        }
    });
}

function setCompLangSelectValue() {
    for (var key in typeData) {
        if (key == "comp_lang") {
            var arr = typeData[key];
            for (var idx in arr) {
                $("#lang_type").append("<option value='" + arr[idx].data_code + "'>" + arr[idx].data_name + "</option>");
            }
            break;
        }
    }
}

function setSelectValue(r, containerId, type) {
    var fieldName = "";
    if (type == 'req_arg_type') {
        fieldName = 'arg_type';
    } else if (type == 'data_type') {
        fieldName = 'data_type';
    }
    var containerIds = containerId.split(",");
    for (var key in typeData) {
        if (key == type) {
            var arr = typeData[key];
            for (var idx in arr) {
                for (var j = 0; j < containerIds.length; j++) {
                    $("#" + containerIds[j] + ">tbody").find("#" + fieldName + r).append("<option value='" + arr[idx].data_code + "'>" + arr[idx].data_name + "</option>");
                }

            }
        }
    }
}

function setSelectValueExt(r, containerId, type, val) {
    var fieldName = "";
    if (type == 'req_arg_type') {
        fieldName = 'arg_type';
    } else if (type == 'data_type') {
        fieldName = 'data_type';
    }
    var containerIds = containerId.split(",");
    for (var key in typeData) {
        if (key == type) {
            var arr = typeData[key];
            for (var idx in arr) {
                for (var j = 0; j < containerIds.length; j++) {
                    var trs = $("#" + containerIds[j] + ">tbody").children();
                    var trObj = trs[r];
                    var selected = "";
                    if (arr[idx].data_code == val) {
                        selected = "selected";
                    }
                    $(trObj).find("#" + fieldName).append("<option value='" + arr[idx].data_code + "' " + selected + ">" + arr[idx].data_name + "</option>");
                }

            }
        }
    }
}

function deleteRow(obj) {
    var tr = $(obj).parent().parent();
    var id = $(tr).find("#id" + $(tr).index()).val();
    if(id) {
        if (deleteIds.indexOf("," + id + ",") == -1) {
            deleteIds += id + ",";
        }
    }
    $(tr).remove();
}

function deleteFormRow(obj) {
    var tr = $(obj).parent().parent();
    var id = $(tr).find("#id" + $(tr).index()).val();
    if (deleteFormDataIds.indexOf("," + id + ",") == -1) {
        deleteFormDataIds += id + ",";
    }
    $(tr).remove();
}

function deleteRespRow(obj) {
    var tr = $(obj).parent().parent();//td-tr
    var tableObj = $(tr).parent().parent();//tbody-table
    var tableId = $(tableObj).attr("id");
    var rowIndex = $(tr).index();
    var tablePrefix = tableId + rowIndex;
    var trs = $("#" + tableId + ">tbody").children();
    var ids = [];
    var deleteTrs = [];//需要删除的行
    for (var i = 0; i < trs.length; i++) {
        var tds = $(trs[i]).children();
        if (tds.length == 1) {
            var subTable = $(tds[0]).children();
            var id = $(subTable).attr("id");
            if (id.indexOf(tablePrefix) > -1) {
                deleteTrs.push(trs[i]);
                parseRespTableForDelete(id, ids);
            }
        }
    }
    console.log("ids=" + ids);
    for (var i = 0; i < ids.length; i++) {
        if (deleteRespIds.indexOf("," + ids[i] + ",") == -1) {
            deleteRespIds += ids[i] + ",";
        }
    }
    //删除当前行
    var id = $(tr).find("#id").val();
    if (deleteRespIds.indexOf("," + id + ",") == -1) {
        deleteRespIds += id + ",";
    }
    for (var i = 0; i < deleteTrs.length; i++) {
        $(deleteTrs[i]).remove();
    }
    $(tr).remove();
    console.log("deleteRespIds=" + deleteRespIds);
}

function parseRespTableForDelete(tableId, ids) {
    console.log("******tableId**:=" + tableId)
    var trs = $("#" + tableId + ">tbody").children();
    for (var i = 0; i < trs.length; i++) {
        var tds = $(trs[i]).children();
        var tdLen = tds.length;
        if (tdLen == 1) {
            var subTable = $(tds[0]).children();
            if (subTable && subTable.length > 0) {
                var subId = $(subTable).attr("id");
                console.log("****subid***=" + subId)
                parseRespTableForDelete(subId, ids);
            }
        } else {
            var id = $(tds[tds.length - 1]).find("#id").val();
            if (id) {
                ids.push(id);
            }
        }
    }
}

function parseRespTable(tableId, items) {
    console.log("******tableId**:=" + tableId)
    var trs = $("#" + tableId + ">tbody").children();
    for (var i = 0; i < trs.length; i++) {
        var tds = $(trs[i]).children();
        var tdLen = tds.length;
        if (tdLen == 1) {
            var subTable = $(tds[0]).children();
            if (subTable && subTable.length > 0) {
                var subId = $(subTable).attr("id");
                console.log("****subid***=" + subId)
                parseRespTable(subId, items);
            }
        } else {
            var id = $(tds[tds.length - 1]).find("#id").val();
            if (id) {
                var data_name = $(trs[i]).find("#data_name").val();
                var data_type = $(trs[i]).find("#data_type").val();
                var description = $(trs[i]).find("#description").val();

                var item = {};
                var msgStr = "";
                if (data_name == "") {
                    var tempObj = $(trs[i]).find("#data_name");
                    $(tempObj).css("background", "pink");
                    var tips = "第" + (i + 1) + "行中的【名称】不能为空!"
                    msgStr = tips;
                }
                if (data_type == "") {
                    var tempObj = $(trs[i]).find("#data_type");
                    $(tempObj).css("background", "pink");
                    var tips = "第" + (i + 1) + "行中的【类型】不能为空!"
                    if (msgStr != "")
                        msgStr += "<br>";
                    msgStr += tips;
                }
                if (description == "") {
                    var tempObj = $(trs[i]).find("#description");
                    $(tempObj).css("background", "pink");
                    var tips = "第" + (i + 1) + "行中的【描述】不能为空!"
                    if (msgStr != "")
                        msgStr += "<br>";
                    msgStr += tips;
                }

                item["msg"] = msgStr;

                item["id"] = parseInt(id);
                item["parentId"] = tableId;
                item["order_num"] = i;
                item["data_name"] = data_name;
                item["data_type"] = data_type;
                item["description"] = description;
                items.push(item);
            }
        }
    }
}

function doWithRespParams(respItems, tableId, id) {
    for (var i = 0; i < respItems.length; i++) {
        var item = respItems[i];
        if (!item["isConfirm"]) {//是否未处理过
            if (item.parentId == id) {
                var rowIndex = addRespRow(tableId, item);
                item["isConfirm"] = true;
                if (item.data_type == '5' || item.data_type == '6') {
                    var newTableId = addRespChildTable(tableId, rowIndex);
                    doWithRespParams(respItems, newTableId, item.id);
                }
            }
        }
    }
}

function addRespChildTable(parentTableId, rowIndex, currObj) {
    if (currObj) {
        rowIndex = $(currObj).parent().parent().index();
    }
    var trs = $("#" + parentTableId + ">tbody").children();
    var idPrefix = parentTableId + rowIndex;
    var count = 0;
    for (var i = 0; i < trs.length; i++) {
        var tds = $(trs[i]).children();
        var tdLen = tds.length;
        if (tdLen == 1) {
            var subTable = $(tds[0]).children();
            var id = $(subTable).attr("id");
            var reg = new RegExp(idPrefix + "_[0-9]+$");
            if (reg.test(id)) {
                count++;
            }
        }
    }
    var newTableId = idPrefix + "_" + count;
    var tableHtml = '<tr><td colspan="4"><table id="' + newTableId + '" class="table table-bordered text-center">' +
        '<caption class="text-right" ><button type="button" class="btn btn-xs btn-primary" onclick="addRespRow(\'' + newTableId + '\')">添加</button>&nbsp;&nbsp;<button type="button" class="btn btn-xs btn-default" onclick="impText(this)">导入</button></caption>' +
        '<thead> ' +
        '<tr> ' +
        '<th style="width:25%;text-align: center;">名称</th> ' +
        '<th style="width:20%;text-align: center;">类型</th> ' +
        '<th style="width:30%;text-align: center;">描述</th> ' +
        '<th style="width:25%;text-align: center;">删除</th> ' +
        '</tr> ' +
        '</thead> ' +
        '<tbody> ' +
        '</tbody> ' +
        '</table></td></tr>';
    if (count == 0) {
        var childs = $("#" + parentTableId + ">tbody").children();
        $(childs[rowIndex]).after(tableHtml);
    } else {
        var prevTableId = parentTableId + rowIndex + "_" + (count - 1);
        $("#" + prevTableId).parent().parent().after(tableHtml);
    }
    return newTableId;
}

function addRespRow(tableId, obj) {
    var len = $("#" + tableId + ">tbody").children().length;
    var trHtml = '<tr>' +
        '<td><input type="text" class="form-control" id="data_name" name="data_name"></td>' +
        '<td><select type="text" class="form-control data-type" id="data_type" name="data_type"><option value="">-=请选择=-</option></select></td>' +
        '<td><input type="text" class="form-control" id="description" name="description"></td>' +
        '<td><button type="button" class="btn btn-xs btn-danger" onclick="deleteRespRow(this)">删除</button><input type="hidden" id="id" name="id" value=0><input type="hidden" id="parentId" name="parentId" value=0></td>' +
        '</tr>';
    if (obj) {
        if (tableId.indexOf("_")>-1) {
            var idx = tableId.substring(tableId.lastIndexOf("_") + 1, tableId.length);
            if (idx != obj.idx) {//添加table
                var trObj = $("#respTable>tbody").find("input[value='" + obj.parentId + "']").parent().parent();
                var rowIndex = $(trObj).index();
                var idPrefix = $(trObj).parent().parent().attr("id");
                console.log("tableId=" + idPrefix + " rowIndex=" + rowIndex)
                tableId = addRespChildTable(idPrefix, rowIndex);
                len = $("#" + tableId + ">tbody").children().length;
            }
        }
        
        if (obj.data_type == "5" || obj.data_type == "6") {
            trHtml = '<tr>' +
                '<td><input type="text" class="form-control" id="data_name" name="data_name" value="' + obj.data_name + '"></td>' +
                '<td><select type="text" class="form-control data-type" id="data_type" name="data_type"><option value="">-=请选择=-</option></select></td>' +
                '<td><input type="text" class="form-control" id="description" name="description" value="' + obj.description + '"></td>' +
                '<td><button type="button" class="btn btn-xs btn-danger" onclick="deleteRespRow(this)">删除</button><input type="hidden" id="id" name="id" value="' + obj.id + '"><input type="hidden" id="parentId" name="parentId" value="' + obj.parentId + '"><button type="button" class="btn btn-xs btn-info" onclick="addRespChildTable(\'' + tableId + '\',' + len + ',this)">添加子项</button></td>' +
                '</tr>';
        } else {
            trHtml = '<tr>' +
                '<td><input type="text" class="form-control" id="data_name" name="data_name" value="' + obj.data_name + '"></td>' +
                '<td><select type="text" class="form-control data-type" id="data_type" name="data_type"><option value="">-=请选择=-</option></select></td>' +
                '<td><input type="text" class="form-control" id="description" name="description" value="' + obj.description + '"></td>' +
                '<td><button type="button" class="btn btn-xs btn-danger" onclick="deleteRespRow(this)">删除</button><input type="hidden" id="id" name="id" value="' + obj.id + '"><input type="hidden" id="parentId" name="parentId" value="' + obj.parentId + '"></td>' +
                '</tr>';
        }
    }

    $("#" + tableId + ">tbody").append(trHtml);
    if (obj) {
        setSelectValueExt(len, tableId, 'data_type', obj.data_type);
    } else {
        setSelectValueExt(len, tableId, 'data_type', '');
    }
    return len;
}

function addRow() {
    var len = $("#paramTable>tbody").find("tr").length;
    var trHtml = '<tr>' +
        '<td><input type="text" class="form-control" id="arg_name' + len + '" name="arg_name' + len + '"></td>' +
        '<td><select class="form-control arg-type" id="arg_type' + len + '" name="arg_type' + len + '"><option value="">-=请选择=-</option></select></td>' +
        '<td><select type="text" class="form-control data-type" id="data_type' + len + '" name="data_type' + len + '"><option value="">-=请选择=-</option></select></td>' +
        '<td><select type="text" class="form-control is-require" id="is_require' + len + '" name="is_require' + len + '"><option value="1">是</option><option value="2">否</option></select></td>' +
        '<td><input type="text" class="form-control" id="description' + len + '" name="description' + len + '"></td>' +
        '<td><button type="button" class="btn btn-xs btn-danger" onclick="deleteRow(this)">删除</button><input type="hidden" id="id' + len + '" name="id' + len + '" value=0></td>' +
        '</tr>';
    $("#paramTable>tbody").append(trHtml);
    setSelectValue(len, 'paramTable', 'req_arg_type');
    setSelectValue(len, 'paramTable', 'data_type');
}

function addFormRow() {
    var len = $("#formDataTable>tbody").find("tr").length;
    var trHtml = '<tr>' +
        '<td><input type="text" class="form-control" id="data_name' + len + '" name="data_name' + len + '"></td>' +
        '<td><select type="text" class="form-control data-type" id="data_type' + len + '" name="data_type' + len + '"><option value="">-=请选择=-</option></select></td>' +
        '<td><select type="text" class="form-control is-require" id="is_require' + len + '" name="is_require' + len + '"><option value="1">是</option><option value="2">否</option></select></td>' +
        '<td><input type="text" class="form-control" id="description' + len + '" name="description' + len + '"></td>' +
        '<td><button type="button" class="btn btn-xs btn-danger" onclick="deleteFormRow(this)">删除</button><input type="hidden" id="id' + len + '" name="id' + len + '" value=0></td>' +
        '</tr>';
    $("#formDataTable>tbody").append(trHtml);
    setSelectValue(len, 'formDataTable', 'data_type');
}

function addReqSample(obj) {
    var len = $("#sample_ul").find("li").length - 1;
    var name = "实例" + (len + 1);
    layer.prompt({
        title: '您可以对【' + name + '】重新命名',
        formType: 0,
        value: name,
        btn: ["确定", "取消"]
    }, function (text, index) {
        if (text && text != "") {
            insertReqSampleRow(len, text);
            layer.close(index);
        };
    });
}

function insertReqSampleRow(i, name) {
    var id = "reqSample" + i;
    var navItemHtml = '<li role="presentation" class="tab-list">' +
        '<input type="hidden" id="reqSampleId' + i + '" name="reqSampleId' + i + '" value="0">' +
        '<a href="#' + id + '" aria-controls="' + id + '" role="tab" data-toggle="tab">' + name +
        '<i class="fa fa-remove tab-close"></i></a>' +
        '</li>';
    $("#sample_ul").find("li:last").before(navItemHtml);
    $("#sample_ul").find("li").removeClass("active");
    $("#sample_ul").find("li:last").prev().addClass("active");

    var contentHtml = '<div role="tabpanel" class="tab-pane" id="' + id + '">' +
        '<textarea style="resize:none;height:200px;" class="form-control" id="req_sample' + i + '" name="req_sample' + i + '"></textarea>' +
        '</div >';
    $("#sample_content").append(contentHtml);
    $("#sample_content").find(".tab-pane").removeClass("active");
    $("#sample_content").find(".tab-pane:last").addClass("active");
}

function addReqCode(obj) {
    layer.open({
        type: 1,
        title: '选择语言',
        area: ['360px', '200px'],
        content: $('#computer_lang'),
        btn: ['确定', '取消'],
        btn1: function (_index, layero) {
            var val = $("#lang_type").val();
            var arrTypes = $("#code_ul").find(".langTypeClass");
            for (var j = 0; j < arrTypes.length; j++) {
                if ($(arrTypes[j]).val() == val) {
                    top.layer.alert('该语言代码已存在，不能重复添加!')
                    return;
                }
            }

            var name = $("#lang_type").find("option:selected").text();
            var len = $("#code_ul").find("li").length - 1;
            insertReqCodeRow(len, val, name);
            layer.close(_index);
        }
    });
}
function insertReqCodeRow(i, val, name) {
    var id = "reqCode" + i;
    var navItemHtml = '<li role="presentation" class="tab-list">' +
        '<input type="hidden" id="reqCodeId' + i + '" name="reqCodeId' + i + '" value="0">' +
        '<input type="hidden" class="langTypeClass" id="langType' + i + '" name="langType' + i + '" value="' + val + '">' +
        '<a href="#' + id + '" aria-controls="' + id + '" role="tab" data-toggle="tab">' + name +
        '<i class="fa fa-remove tab-close"></i></a>' +
        '</li>';
    $("#code_ul").find("li:last").before(navItemHtml);
    $("#code_ul").find("li").removeClass("active");
    $("#code_ul").find("li:last").prev().addClass("active");

    var contentHtml = '<div role="tabpanel" class="tab-pane" id="' + id + '">' +
        '<textarea style="resize:none;height:200px;" class="form-control" id="req_code' + i + '" name="req_code' + i + '"></textarea>' +
        '</div >';
    $("#code_content").append(contentHtml);
    $("#code_content").find(".tab-pane").removeClass("active");
    $("#code_content").find(".tab-pane:last").addClass("active");
}