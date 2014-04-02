 $(document).ready(function(){
        $('#subscribe_button').click(function(e){

            //Stop form submission & check the validation
            e.preventDefault();

            // Variable declaration
            var error = false;
            var name = $('#n_name').val();
            var email = $('#n_email').val();

            // Form field validation
            if(name.length == 0){
                var error = true;
                $('#n_name_error').fadeIn(500);
            }else{
                $('#n_name_error').fadeOut(500);
            }
            if(email.length == 0 || email.indexOf('@') == '-1' || email.indexOf('.') == '-1'){
                var error = true;
                $('#n_email_error').fadeIn(500);
            }else{
                $('#n_email_error').fadeOut(500);
            }

            // If there is no validation error, next to process the mail function
            if(error == false){
               // Disable submit button just after the form processed 1st time successfully.
                $('#subscribe_button').attr({'disabled' : 'true', 'value' : 'SUBSCRIBING...' });

                $.post("/subscribe/", $("#n_letter").serialize(),function(result){
                    if(result.status == "ok"){
                        //If the email is sent successfully, remove the submit button
                         $('#subscribe_button').attr({'value' : 'SUBSCRIBED' });
                         $('div#n_success').show();
                         $('div#n_fail').hide();
                    } else {
                        // Enable the submit button again
                        $('#subscribe_button').removeAttr('disabled').attr({'value' : 'NOTIFY ME' });
                        $('div#n_fail').html(result.errors);
                        $('div#n_fail').show();
                        $('div#n_success').hide();
                    }

                });
            }
        });
    });