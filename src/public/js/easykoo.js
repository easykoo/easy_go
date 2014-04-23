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
    var time = timeString.substr(11, 8)
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

var cutoff = function (content) {
    var cutoffLine = "----------"
    var index = content.indexOf();
    if (index > 0) {
        content = content.replace(cutoffLine, "");
        var pre = content.substr(0, index);
        var preIndex = pre.lastIndexOf('</p>');
        if (preIndex > 0) {
            preIndex += 4;
        } else {
            preIndex = pre.lastIndexOf('</div>');
            if (preIndex > 0) {
                preIndex += 4;
            } else {
                return content;
            }
        }
        var pre = content.substr(0, preIndex);
        var nex = content.substr(preIndex, content.length);
        return pre + cutoffLine + nex;
    }
    return content;
}
function direct() {
    var winHeight = $(window).height()

    var $top = $('#goTop');
    var $bottom = $('#goBottom');
    var side = $('#side').offset().left;
    var width = $('#side').width();
    var pos = side + width + 25;
    $top.css({
        "left": pos + "px",
        "top": winHeight / 2 - 23 + "px",
        "width": "45px",
        "height": "45px",
        "position": "fixed",
        "opacity": .4
    })
    $bottom.css({
        "left": pos + "px",
        "top": winHeight / 2 + 23 + "px",
        "width": "45px",
        "height": "45px",
        "position": "fixed",
        "opacity": .4
    })
    $(window).scroll(function () {
        var scroll = $(this).scrollTop()
        if (scroll > 0) {
            $top.removeClass("hidden");
        } else {
            $top.addClass('hidden');
        }

        if (scroll + winHeight == $(document).height()) {
            $bottom.addClass('hidden');
        } else {
            $bottom.removeClass("hidden");
        }
    });
    $top.on("click", function () {
        $('body,html').animate({scrollTop: 0}, 500);
        return false;
    })
    $bottom.click(function () {
        $('html, body').animate({scrollTop: $(document).height()}, 300);
        return false;
    });
}