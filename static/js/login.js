$(document).ready(function () {
    //登录
    $("#login-form").validate({
        rules: {
            username: {
                required: true,
                rangelength: [5, 20]
            },
            password: {
                required: true,
                rangelength: [5, 20]
            },
            captcha: {
                required: true,
                rangelength: [1, 10]
            }
        },
        messages: {
            username: {
                required: "请输入用户名",
                rangelength: "用户名必须是5-20位"
            },
            password: {
                required: "请输入密码",
                rangelength: "密码必须是5-20位"
            }
            ,captcha: {
                required: "验证码",
                rangelength: "验证码必须是1-10位"
            }
        },
        submitHandler: function (form) {
            var urlStr = "/login"
            alert("urlStr:" + urlStr)
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert("data:" + data.message)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/"
                        }, 1000)
                    }
                },
                error: function (data, status) {
                    alert("err:" + data.message)

                }
            });
        }
    });
});