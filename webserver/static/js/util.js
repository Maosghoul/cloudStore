var util = {};

util.httpGet = function (url, error) {
    $.ajax({
        url: url,
        type: "GET",
        async: true,
        error: error
    });
};

util.httpPost = function (url, data, error) {
    var result = false
    $.ajax({
        url: url,
        type: "POST",
        async: false,
        dataType: "JSON",
        contentType: "application/json",
        data: data,
        success: function () {
            result = true
        },
        error: error
    });
    return result
};

util.httpPostWithResp = function (url, data, error) {
    var resp = new Object()
    $.ajax({
        url: url,
        type: "POST",
        async: false,
        dataType: "JSON",
        contentType: "application/json",
        data: data,
        success: function (res) {
            result = res
        },
        error: error
    });
    return result
}

util.httpFormData = function (url, data, error) {
    var result = false
    $.ajax({
        url: url,
        type: "post",
        async: false,
        data: data,
        contentType: false, //"multipart/form-data",
        processData: false,
        success: function (res) {
            console.log(res.message)
            result = true
        },
        error: error
    });
    return result
};

// 字符串格式化
util.format = function (src) {
    if (arguments.length == 0) return null;
    let args = Array.prototype.slice.call(arguments, 1);
    return src.replace(/\{(\d+)\}/g, function (m, i) {
        return args[i];
    });
};


util.percent = function (v) {
    let n = parseFloat(v);
    return util.format("{0}%", n.toFixed(2))
};

util.str2Int = function (v) {
    return parseInt(v)
};

util.unit = new Array("B", "KB", "MB", "GB");

util.b2string = function (v, rate) {
    let n = v;
    let i = 0;
    for (; n > rate;) {
        n /= rate;
        i++;
        if (i === util.unit.length) {
            break
        }
    }
    return util.format("{0}{1}", Math.round(n), util.unit[i])
};


util.getItem = function (sKey) {
    return decodeURIComponent(document.cookie.replace(new RegExp("(?:(?:^|.*;)\\s*" + encodeURIComponent(sKey).replace(/[-.+*]/g, "\\$&") + "\\s*\\=\\s*([^;]*).*$)|^.*$"), "$1")) || null;
};

util.setItem = function (sKey, sValue, vEnd, sPath, sDomain, bSecure) {
    if (!sKey || /^(?:expires|max\-age|path|domain|secure)$/i.test(sKey)) {
        return false;
    }
    var sExpires = "";
    if (vEnd) {
        switch (vEnd.constructor) {
            case Number:
                sExpires = vEnd === Infinity ? "; expires=Fri, 31 Dec 9999 23:59:59 GMT" : "; max-age=" + vEnd;
                break;
            case String:
                sExpires = "; expires=" + vEnd;
                break;
            case Date:
                sExpires = "; expires=" + vEnd.toUTCString();
                break;
        }
    }
    document.cookie = encodeURIComponent(sKey) + "=" + encodeURIComponent(sValue) + sExpires + (sDomain ? "; domain=" + sDomain : "") + (sPath ? "; path=" + sPath : "") + (bSecure ? "; secure" : "");
    return true;
};

util.hasItem = function (sKey) {
    return (new RegExp("(?:^|;\\s*)" + encodeURIComponent(sKey).replace(/[-.+*]/g, "\\$&") + "\\s*\\=")).test(document.cookie);
};

util.removeItem = function (sKey, sPath, sDomain) {
    if (!sKey || !this.hasItem(sKey)) {
        return false;
    }
    document.cookie = encodeURIComponent(sKey) + "=; expires=Thu, 01 Jan 1970 00:00:00 GMT" + (sDomain ? "; domain=" + sDomain : "") + (sPath ? "; path=" + sPath : "");
    return true;
};

