$(document).ready(function () {

    $("#system-form").validate({
        rules: {
            cpu: {
                required: true,
                rangelength: [1, 100]
            },
            cpunum: {
                required: true,
                rangelength: [1, 20]
            },
            disk: {
                required: true,
                rangelength: [1, 100]
            },
            mem: {
                required: true,
                rangelength: [1, 20]
            }
            ,
            os: {
                required: true,
                rangelength: [1, 100]
            }
            ,
            hostname: {
                required: true,
                rangelength: [1, 100]
            }
            ,
            intranetip: {
                required: true,
                rangelength: [1, 100]
            }
        },
        messages: {
            cpu: {
                rangelength: "不能为空"
            },
            cpunum: {
                rangelength: "不能为空"
            },
            disk: {
                rangelength: "不能为空"
            }
            ,
            mem: {
                rangelength: "不能为空"
            }
            ,
            os: {
                rangelength: "不能为空"
            }
            ,
            hostname: {
                rangelength: "不能为空"
            }
            ,
            intranetip: {
                rangelength: "不能为空"
            }
        },
        submitHandler: function (form) {

            $(form).ajaxSubmit({
                url: "/assetinfo",
                type: "post",
                dataType: "json",
                success: function (data, status) {
                    alert(data.message)
                    if (data.code == 1) {
                        setTimeout(function () {
                            window.location.href = "/assetinfo"
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