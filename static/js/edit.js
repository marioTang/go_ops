$(document).ready(function () {

    $("#edituser-form").validate({
        rules: {
            username: {
                required: true,
                rangelength: [1, 100]
            },

            edit: {
                required: true,
                rangelength: [1, 100]
            },
            deldet: {
                required: true,
                rangelength: [1, 20]
            },
            accesskeyid: {
                required: true,
                rangelength: [1, 100]
            },
            accesskeysecret: {
                Level: true,
                rangelength: [1, 20]
            }
        },
        messages: {
            username: {
                rangelength: "不能为空"
            },
            edit: {
                rangelength: "不能为空"
            },
            deldet: {
                rangelength: "不能为空"
            },
            accesskeyid: {
                rangelength: "不能为空"
            }
            ,
            accesskeysecret: {
                rangelength: "不能为空"
            }
        },
        submitHandler: function (form) {

            $(form).ajaxSubmit({
                url: "/userinfo",
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert(data.message)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/userinfo"
                        }, 1000)
                    }
                },
                error: function (data) {
                    alert("err:" + data.message )

                }
            });
        }
    });

});