String.prototype.endWith = function (s) {
    if (s == null || s == "" || this.length == 0 || s.length > this.length)
        return false;
    if (this.substring(this.length - s.length) == s)
        return true;
    else
        return false;
    return true;
};

var setCheckSession = function () {
    setInterval(checkSession, 60000);
}

var checkSession = function () {
    $.ajax('common/ajax/checkSession.do', {
        dataType: 'json',
        type: "POST",
        success: function (data) {
            if (!data.success) {
                window.location.href = "index.jsp";
            }
        }
    });
}

var timeStamp2String = function (time) {
    var datetime = new Date();
    datetime.setTime(time);
    var year = datetime.getFullYear();
    var month = datetime.getMonth() + 1 < 10 ? "0" + (datetime.getMonth() + 1) : datetime.getMonth() + 1;
    var date = datetime.getDate() < 10 ? "0" + datetime.getDate() : datetime.getDate();
    var hour = datetime.getHours() < 10 ? "0" + datetime.getHours() : datetime.getHours();
    var minute = datetime.getMinutes() < 10 ? "0" + datetime.getMinutes() : datetime.getMinutes();
    var second = datetime.getSeconds() < 10 ? "0" + datetime.getSeconds() : datetime.getSeconds();
    return hour + ":" + minute + ":" + second + " " + month + "/" + date + "/" + year;
};

var changeLanguage = function (lang) {
    var url = window.location.href;
    url = url.replace('&locale=en', '').replace('&locale=zh_CN', '');
    if (url.endWith("?locale=en") || url.endWith("?locale=zh_CN")) {
        url = url.replace('?locale=en', '').replace('?locale=zh_CN', '');
    }
    url += (url.indexOf('?') > 0 ? '&' : '?') + 'locale=' + lang;
    window.location.href = url;
}

var filterSqlStr = function (value) {
    var sqlStr = sql_str().split(',');
    var flag = false;

    for (var i = 0; i < sqlStr.length; i++) {
        if (value.toLowerCase().indexOf(sqlStr[i]) != -1) {
            flag = true;
            break;
        }
    }
    return flag;
}

var sql_str = function () {
    var str = "and,delete,or,exec,insert,select,union,update,count,*,',join,>,<";
    return str;
}