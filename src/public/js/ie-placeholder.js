$('#n_name').focus(function() {
    $('div.n-name').hide();
});

$('#n_name').focusout(function() {
    if($('#n_name').val() == "") {
        $('div.n-name').show();
    }
});

$('#n_email').focus(function() {
    $('div.n-email').hide();
});

$('#n_email').focusout(function() {
    if($('#n_email').val() == "") {
        $('div.n-email').show();
    }
});            