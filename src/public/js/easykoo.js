String.prototype.endWith = function (s) {
    if (s == null || s == "" || this.length == 0 || s.length > this.length)
        return false;
    if (this.substring(this.length - s.length) == s)
        return true;
    else
        return false;
    return true;
};

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