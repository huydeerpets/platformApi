/**
 * oss 文件上传
 */

var accessid = ''
var accesskey = ''
var host = ''
var policyBase64 = ''
var signature = ''
var callbackbody = ''
var key = ''
var expire = 0
var g_object_name = ''
var g_object_name_type = ''
var now = timestamp = Date.parse(new Date()) / 1000;
var g_objType = '0';//0-图片；1-语音；2-视频
var bucketDir = '';//默认OSS的路径

function send_request() {
  var xmlhttp = null;
  if (window.XMLHttpRequest) {
    xmlhttp = new XMLHttpRequest();
  } else if (window.ActiveXObject) {
    xmlhttp = new ActiveXObject("Microsoft.XMLHTTP");
  }

  if (xmlhttp != null) {
    serverUrl = '/OSS/GetPolicyToken'
    xmlhttp.open("GET", serverUrl, false);
    xmlhttp.send(null);
    return xmlhttp.responseText
  } else {
    alert("您的浏览器不支持 XMLHTTP.");
  }
}

function get_signature() {
  //可以判断当前expire是否超过了当前时间,如果超过了当前时间,就重新取一下.3s 做为缓冲
  now = timestamp = Date.parse(new Date()) / 1000;
  if (expire < now + 3) {
    body = send_request()
    var obj = eval("(" + body + ")");
    host = obj['host']
    policyBase64 = obj['policy']
    accessid = obj['accessid']
    signature = obj['signature']
    expire = parseInt(obj['expire'])
    callbackbody = obj['callback']
    key = obj['dir'];
    bucketDir = key;
    return true;
  }
  return false;
};

function set_upload_param(up, filename, ret) {
  if (ret == false) {
    ret = get_signature()
  }
  if (bucketDir != '') {
    key = bucketDir;
  }
  if (g_objType == '1') {
    key += "voice/";
  } else if (g_objType == '2') {
    key += "vedio/";
  }
  g_object_name = key;
  if (filename != '') {
    calculate_object_name(filename);
  }
  new_multipart_params = {
    'key': g_object_name,
    'policy': policyBase64,
    'OSSAccessKeyId': accessid,
    'success_action_status': '200', //让服务端返回200,不然，默认会返回204
    'callback': callbackbody,
    'signature': signature,
  };
  up.setOption({
    'url': host,
    'multipart_params': new_multipart_params
  });
  //up.start();
}
function calculate_object_name(filename) {
  if (g_object_name_type == 'local_name') {
    g_object_name += "${filename}"
  }
  else if (g_object_name_type == 'random_name') {
    suffix = get_suffix(filename)
    g_object_name = key + random_string2() + suffix;
  } else {
    g_object_name = key + filename;
  }
  return ''
}
function add0(m) {
  return m < 10 ? '0' + m : m;
}

function get_suffix(filename) {
  pos = filename.lastIndexOf('.')
  suffix = ''
  if (pos != -1) {
    suffix = filename.substring(pos)
  }
  return suffix;
}

function random_string2() {
  var time = new Date();
  var y = time.getFullYear();
  var m = time.getMonth() + 1;
  var d = time.getDate();
  var h = time.getHours();
  var mm = time.getMinutes();
  var s = time.getSeconds();
  var ms = time.getMilliseconds();

  return y + add0(m) + add0(d) + add0(h) + add0(mm) + add0(s) + ms;
}

function progressHtml(randStrID) {
  var html = [];
  html.push('<img src="" id="img_' + randStrID + '" style="display:none"/>');
  html.push('<div style="height:10px; border:2px solid #09F;margin:5px" id="parent_' + randStrID + '">');
  html.push('<div style="width:0; height:100%; background-color:#09F; text-align:center; line-height:10px; font-size:20px; font-weight:bold;" id="progess_' + randStrID + '"></div>');
  html.push('</div>');
  return html.join("");
}

function initCKEditorUpload(ckeditorObj) {
  var uploader = new plupload.Uploader({
    disable_statistics_report: false,
    runtimes: 'html5,flash,silverlight,html4',
    browse_button: 'cke_32',//如果工具拦中添加其他的话，可能需要调整该值
    container: 'container',//指的是容器ID
    drop_element: 'container',
    max_file_size: '4mb',
    flash_swf_url: '/static/js/plupload/js/Moxie.swf',
    silverlight_xap_url: '/static/js/plupload/Moxie.xap',
    dragdrop: true,
    chunk_size: '4mb',
    multi_selection: true,
    get_new_uptoken: false,
    unique_names: false,
    save_key: false,
    auto_start: true,
    log_level: 5,
    filters: {
      mime_types: [ //只允许上传图片和zip文件
        { title: "Image files", extensions: "jpg,jpeg,gif,png,bmp" }
      ],
      max_file_size: '4mb', //最大只能上传400kb的文件
      prevent_duplicates: true //不允许选取重复文件
    },
    init: {
      BeforeChunkUpload: function (up, file) {
        console.log("before chunk upload:", file.name);
      },
      FilesAdded: function (up, files) {
        var htmlAry = [];
        for (var i = 0; i < files.length; i++) {
          var randStrID = files[i].id;
          htmlAry.push(progressHtml(randStrID));

          if (i == files.length - 1) {
            ckeditorObj.insertHtml(htmlAry.join(""));
          }
        }
        get_signature();//获取OSS访问的签名
        up.start();
      },
      BeforeUpload: function (up, file) {
        set_upload_param(up, file.name, true);//可以自定义文件名称
      },
      UploadProgress: function (up, file) {
        var randStrID = file.id;
        $("#progess_" + randStrID, $("body", $(ckeditorObj.document.$))).css("width", file.percent + "%").text(file.percent + "%");
      },
      UploadComplete: function () {
        //所有文件上传完成
      },
      FileUploaded: function (up, file, info) {
        if (info.status == 200) {
          var randStrID = file.id;
          var url = host + "/" + g_object_name;
          ckeditorObj.document.getById("img_" + randStrID)
            .setAttributes({ "src": url, "data-cke-saved-src": url })
            .removeAttributes(["style", "id"]);
          ckeditorObj.document.getById("parent_" + randStrID).remove();
        } else {
          top.layer.alert("status:=" + info.status + " info:=" + info.response);
        }
      },
      Error: function (up, err, errTip) {
        if (err.code == -600) {
          top.layer.alert("选择的图片文件太大了！");
        }
        else if (err.code == -601) {
          top.layer.alert("图片上传失败！");
        }
        else if (err.code == -602) {
          top.layer.alert("图片已经上传过一遍了！");
        }
        else {
          top.layer.alert("图片上传失败！");
        }
      }
    }
  });
  uploader.bind('BeforeUpload', function () {
    console.log("hello man, i am going to upload a file");
  });

  uploader.bind('FileUploaded', function () {
    console.log('hello man,a file is uploaded');
  });
  uploader.init();
}