String.prototype.endWith = function (s) {
    if (s == null || s == "" || this.length == 0 || s.length > this.length)
        return false;
    if (this.substring(this.length - s.length) == s)
        return true;
    else
        return false;
    return true;
};

var formatTime = function (timeString) {
    var date = timeString.substr(0, 10)
    var time = timeString.substr(11,8)
    return date + " " + time;
};

var changeLanguage = function (lang) {
    $.ajax('/language/change/' + lang, {
        dataType: 'json',
        type: "GET",
        success: function (data) {
            if (data.success) {
                location.reload();
            }
        }
    });
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