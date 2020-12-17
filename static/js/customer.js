window.onload = function() {

    $('[data-toggle="birthdate"]').datepicker({format: 'yyyy-mm-dd'});
    function get_age(born, now) {
        var birthday = new Date(now.getFullYear(), born.getMonth(), born.getDate());
        if (now >= birthday)
            return now.getFullYear() - born.getFullYear();
        else
            return now.getFullYear() - born.getFullYear() - 1;
    }

    $.validator.addMethod("birth", function (value, element) {
        let dob = $("#birthdate").val();
        let now = new Date();
        let birthdate = dob.split("-");
        let born = new Date(birthdate[0], birthdate[1]-1, birthdate[2]);
        let age=get_age(born,now);
        console.log(dob)
        console.log(now)
        console.log(birthdate)
        console.log(age)
        if ( age >= 18 && age <= 60) {
            console.log(age)
            return true;
        } else {
            return false;
        }
    });

    $("form[name='reqform']").validate({
        rules: {
            FirstName: {
                required: true,
                maxlength: 100
            },
            LastName: {
                required: true,
                maxlength: 100,
            },
            Gender: "required",
            Email: {
                required: true,
                email: true
            },
            BirthDate: {
                required: true,
                birth: true
            },
            Address: {
                maxlength: 200
            }
        },
        messages: {
            FirstName: {
                required: "Please enter your firstname",
                maxlength: "Maximum length 100 char"
            },
            LastName: {
                required: "Please enter your lastname",
                maxlength: "Maximum length 100 char"
            },
            Gender: "Please choose your gender",
            Address: {
                maxlength: "Maximum length 200 char"
            },
            BirthDate: {
                required: "Please provide a birthdate",
                birth: "Your age must from 18 till 60 years"
            },
            Email: "Please enter a valid email address"
        }
    });

    $('#reqform').submit(function (e) {
        if ($("#reqform").valid() == true) {
            $.ajax({
                url: apiurl,
                type: method,
                data: new FormData(this),
                processData: false,
                contentType: false,
                success: function (data) {
                    alert(data)
            }
        });
            e.preventDefault();
        }
    });
}

function Back() {
    if(document.referrer == location.origin + "/search") {
        window.location.replace("/");

    } else {
        history.back();
    }
}