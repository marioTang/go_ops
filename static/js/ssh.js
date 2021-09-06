$(document).ready(function () {
    //注册
    $("#addssh").validate({
        rules: {
            hostname: {
                required: true,
                rangelength: [1, 100]
            },
            username: {
                required: true,
                rangelength: [1, 30]
            },
            port: {
                required: true,
                rangelength: [1, 20],
            },
        },
        messages: {
            hostname: {
                required: "请输入主机名名",
                rangelength: "输入主机名称"
            },
            username: {
                required: "请输入用户名",
                rangelength: "请输入用户名"
            },
            port: {
                required: "请确认端口号",
                rangelength: "端口号必须是1-20位",
            },
        },



        submitHandler: function (form) {
            var editid = $("#editid").val()
            var urlStr = "/ssh";
            if (editid != "") {
                var  urlStr = "/sshinfo"
            }
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {

                    if (data.code == 1) {
                        $("#alert").html(data.message)
                    }
                    else if (data.code == 2) {
                        $("#alert").text(data.message)
                    }
                },
            })
        }
    });

    //添加监控项目
    $("#addmonitor").validate({
        rules: {
            groupname: {
                required: true,
                rangelength: [1, 100]
            },
            projectname: {
                required: true,
                rangelength: [1, 100]
            },
            hostname: {
                required: true,
                rangelength: [1, 30]
            },
            crontime: {
                required: true,
                rangelength: [1, 30]
            },
            scriptpath: {
                required: true,
                rangelength: [1, 20],
            },
        },
        messages: {
            groupname: {
                required: "请必须输入项目名称",
                rangelength: "请必须输入项目名称"
            },
            projectname: {
                required: "请必须输入定时任务名称",
                rangelength: "请必须输入定时任务名称"
            },
            hostname: {
                required: "请必须输入主机名称",
                rangelength: "请必须输入主机名称"
            },
            scriptpath: {
                required: "请必须输入脚本路径",
                rangelength: "请必须输入脚本路径",
            },
            crontime: {
                required: "请必须定时任务时间",
                rangelength: "请必须定时任务时间"
            }
        },



        submitHandler: function (form) {
             var editid = $("#editid").val()
            var urlStr = "/monitor";
            if (editid != "") {
                var  urlStr = "/updatamonitor"
            }
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {

                    if (data.code == 1) {
                        $("#alert").html(data.message)
                    }
                    else if (data.code == 2) {
                        $("#alert").text(data.message)
                    }
                },
            })
        }
    });


    $("#adduser").validate({
        rules: {
            username: {
                required: true,
                rangelength: [1, 100]
            },
            password: {
                required: true,
                rangelength: [1, 30]
            },
            groupname: {
                required: true,
                rangelength: [1, 30]
            },
            remarks: {
                required: true,
                rangelength: [1, 20],
            },
            accesskeyid: {
                required: true,
                rangelength: [1, 20],
            },
            accesskeysecret: {
                required: true,
                rangelength: [1, 20],
            },
        },
        messages: {
            username: {
                required: "请输入用户名",
                rangelength: "输入用户名"
            },
            password: {
                required: "请输入密码",
                rangelength: "请输入密码"
            },
            groupname: {
                required: "请输入商户名称",
                rangelength: "请输入商户名称"
            },
            remarks: {
                required: "请输入备注",
                rangelength: "请输入备注",
            },
            accesskeyid: {
                required: "请输入accesskeyid",
                rangelength: "请输入accesskeyid",
            },
            accesskeysecret: {
                required: "请输入accesskeysecret",
                rangelength: "请输入accesskeysecret",
            },
        },



        submitHandler: function (form) {
            $(form).ajaxSubmit({
                url: "/userinfo",
                type: "post",
                dataType: "json",
                success: function (data, status) {

                    if (data.code == 1) {
                        $("#alert").html(data.message)
                    }
                    else if (data.code == 2) {
                        $("#alert").text(data.message)
                    }
                },
            })
        }
    });

    $('#submit').click(function(){
        var options=$("#selectdata option:selected").val();  //获取选中的项
        var cmd = $("#cmd").val();
        $.ajax({
            url: "/sshrun",
            method: 'POST',
            data: {selectdata: options, cmd: cmd},
            success: function (data, status) {
                if (data.code == 1) {
                    $("#alert").text(data.message)
                }
                else if (data.code == 2) {
                    $("#alert").text(data.message)
                }
            }
        })

    })

    $('#plmxs_submit').click(function(){
        var options=$("#selectdata option:selected").val();  //获取选中的项
        var ue = UE.getEditor('content').getPlainTxt()
        var filename = $("#filename").val();

        $.ajax({
            url: "/plmxs",
            method: 'POST',
            data: {options: options,ue:ue, filename:filename},
            success: function (data, status) {
                alert("1111111111111111111")
                if (data.code == 1) {
                    alert(text(data.message))
                }
                else if (data.code == 2) {
                    alert("fffffffffffff")
                }
            }
        })

    })


    $("#addssh").validate({
        rules: {
            hostname: {
                required: true,
                rangelength: [1, 100]
            },
            username: {
                required: true,
                rangelength: [1, 30]
            },
            port: {
                required: true,
                rangelength: [1, 20],
            },
        },
        messages: {
            hostname: {
                required: "请输入主机名名",
                rangelength: "输入主机名称"
            },
            username: {
                required: "请输入用户名",
                rangelength: "请输入用户名"
            },
            port: {
                required: "请确认端口号",
                rangelength: "端口号必须是1-20位",
            },
        },



        submitHandler: function (form) {
            var editid = $("#editid").val()
            var urlStr = "/ssh";
            if (editid != "") {
                var  urlStr = "/sshinfo"
            }
            $(form).ajaxSubmit({
                url: urlStr,
                type: "post",
                dataType: "json",
                success: function (data, status) {

                    if (data.code == 1) {
                        $("#alert").html(data.message)
                    }
                    else if (data.code == 2) {
                        $("#alert").text(data.message)
                    }
                },
            })
        }
    });


    $("#search").validate({
        rules: {
            searchname: {
                required: true,
                rangelength: [1, 100]
            }
        },
        messages: {
            searchname: {
                required: "",
                rangelength: "请输入搜索的名称"
            },
        },



        submitHandler: function (form) {
            var urlStr = "/seachname";
            $(form).ajaxSubmit({
                url: urlStr,
                type: "get",
                dataType: "json",
                success: function (data, status) {
                },
            })
        }
    });


});

