var layer = layui.layer,
    element=layui.element,
    laydate = layui.laydate,
    upload = layui.upload,
    upload2 = layui.upload,
    layedit = layui.layedit,
    form = layui.form;
/**
 * AJAX全局设置
 */
$.ajaxSetup({
    type:'post',
    dataType: "json",
    headers: {
        'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')
    }
});
layedit.build('content');
/**
 * 通用日期时间选择
 */
$('.datetime').on('click', function () {
    laydate.render({
        elem: this,
        istime: true
    });
});

laydate.render({
    elem: '#datetime1' //指定元素
    ,range: '|'
    ,type: 'datetime'
    ,value: $('#datetime1').val()?$('#datetime1').val():$('#datetime1').data('inure')+' | '+$('#datetime1').data('expire')
});
laydate.render({
    elem: '#datetime2' //指定元素
    ,range: '|'
    ,type: 'datetime'
});

/**
 * 通用批量处理（审核、取消审核、删除）
 */
$('.ajax-action').on('click', function () {
    var _action = $(this).data('action');
    layer.confirm('确定执行此操作？', {
        shade: false,
        btn: ['确定','取消'] //按钮
    }, function(){
        var load = layer.load();
        $.ajax({
            url: _action,
            data: $('.ajax-form').serialize(),
            success: function (info) {
                layer.msg(info.msg,{time:1000},function(){
                    location.reload();
                })
            },
            complete:function(){
                layer.close(load);
            }
        });
        layer.closeAll('dialog');
    }, function(){
        layer.closeAll();
    });
    return false;
});
//执行实例
upload.render({
    elem: '.upload_btn' //绑定元素
    ,url: '/upload' //上传接口
    ,headers:{'X-CSRF-TOKEN': $('meta[name="csrf-token"]').attr('content')}
    ,before: function(obj){ //obj参数包含的信息，跟 choose回调完全一致，可参见上文。
        layer.load(); //上传loading
    }
    ,done: function(res, index, upload){
        layer.closeAll('loading'); //关闭loading
        var item = this.item;
        $('img.'+$(item[0]).attr('target')).attr('src',res.path1);
        $('input.'+$(item[0]).attr('target')).val(res.path2);
    }
    ,error: function(index, upload){
        layer.closeAll('loading'); //关闭loading
    }
});
//多图片上传
upload2.render({
    elem: '.uploads_btn'
    ,url: '/upload'
    ,multiple: true
    ,before: function(obj){ //obj参数包含的信息，跟 choose回调完全一致，可参见上文。
        layer.load(); //上传loading
    }
    ,allDone: function(obj){ //当文件全部被提交后，才触发
        console.log(obj.total); //得到总文件数
        layer.closeAll('loading'); //关闭loading
    }
    ,done: function(res){
        //上传完毕
        console.log(res)
        var html = '<div class="layui-upload-img-box">';
        html += '<img src="'+ res.path1 +'" class="layui-upload-img">';
        html += '<input type="hidden" name="thumb[]" value="'+res.path2+'">';
        html += '<input type="number" maxlength="4" name="imgsort[]" value="0" class="layui-input">';
        html += '<button type="button" class="layui-btn layui-btn-fluid layui-btn-danger layui-btn-xs" onclick="deleteImg(this)">删除</button>';
        html += '</div>';
        $('#imglist').append(html)
    }
});

function deleteImg(e)
{
    $(e).parent('div').remove();
}
/**
 * 通用全选
 */
$('.check-all').on('click', function () {
    $(this).parents('table').find('input[type="checkbox"]').prop('checked', $(this).prop('checked'));
});
/*
* 产品关联内容
* */
$(".ajax-relation").on('click',function (){
    var _href = $(this).attr('data-url');
    layer.open({
        shade: false,
        content: '确定关联吗？',
        btn: ['确定', '取消'],
        yes: function (index) {
            $.ajax({
                url: _href,
                type: "post",
                success: function (info) {
                    if(info.code==200){
                        layer.msg(info.msg,{time:1000},function(){
                            location.reload();
                        })
                    }else{
                        layer.msg(info.msg);
                    }
                }
            });
            layer.close(index);
        }
    });

    return false;
})
/**
 * 通用删除
 */

$('.ajax-delete').on('click', function () {
    var _href = $(this).attr('href');
    layer.open({
        shade: false,
        content: '确定删除？',
        btn: ['确定', '取消'],
        yes: function (index) {
            $.ajax({
                url: _href,
                type: "post",
                success: function (info) {
                    if(info.code==200){
                        layer.msg(info.msg,{time:1000},function(){
                            location.reload();
                        })
                    }else{
                        layer.msg(info.msg);
                    }
                }
            });
            layer.close(index);
        }
    });

    return false;
});

/**
 * 清除缓存
 */
$('#clear-cache').on('click', function () {
    var _url = $(this).data('url');
    if (_url !== 'undefined') {
        $.ajax({
            url: _url,
            success: function (data) {
                if (data.code === 1) {
                    setTimeout(function () {
                        location.href = location.pathname;
                    }, 1000);
                }
                layer.msg(data.msg);
            }
        });
    }
    return false;
});
$('.ajax-panel').on('click',function(){
    return ajax_panel(this);
})

function ajax_panel(o){
    var _url = $(o).data('url');
    var _title = $(o).data('title');
    var _width = $(o).data('width');
    var _height = $(o).data('height');
    if(!_width){
        _width = '700px';
    }
    if(!_height){
        _height = '90%';
    }
    if(_url  != 'undefined'){
        layer.open({
            type: 2,
            title: _title,
            shadeClose: false,
            shade: 0.8,
            scrollbar: true,
            area: [_width, _height],
            content: _url //iframe的url
        });
    }
    return false;
}

// 关闭刷新
$('.ajax-panel-close-flush').on('click',function(){

    var _url = $(this).data('url');
    var _title = $(this).data('title');
    var _width = $(this).data('width');
    var _height = $(this).data('height');
    if(!_width){
        _width = '700px';
    }
    if(!_height){
        _height = '90%';
    }
    if(_url  != 'undefined'){
        layer.open({
            type: 2,
            title: _title,
            shadeClose: false,
            shade: 0.8,
            scrollbar: true,
            area: [_width, _height],
            content: _url, //iframe的url
            btn:['确定'],
            btnAlign: 'r',
            end:function (){
                window.location.reload();
            },
            btn1:function (index){
                layer.close(index)
            }
        });
    }
})

/**
 * 通用表单提交(AJAX方式)
 */
form.on('submit(*)', function (data) {
    var load = layer.load();
    $(data.elem).attr('disabled',true)
    $.ajax({
        url: data.form.action,
        type: data.form.method,
        data: $(data.form).serialize(),
        success: function (data) {
            if(data.code==200){
                layer.msg(data.msg,{time:1500},function(){
                    if(data.address){
                        window.location.href = data.address;
                    }else{
                        parent.location.reload();
                    }
                })
            }else{
                layer.msg(data.msg);
            }
            return false;
        },
        error:function(XMLHttpRequest){
            if( XMLHttpRequest.responseJSON){
                var err = XMLHttpRequest.responseJSON.errors;
                var errs = [];
                Object.keys(err).forEach(function(key){
                    if(key){
                        errs.push(err[key][0]);
                    }
                });
                if(errs.length>0){
                    layer.msg(errs[0]);
                }
                layer.close(load)
                $(data.elem).attr('disabled',false)
            }
        },
        complete:function () {
            setTimeout(()=>{
                layer.close(load)
                $(data.elem).attr('disabled',false)
            },6000)
        }
    });

    return false;
});
form.on('switch(switch)', function(data){
    var index = 0;
    if(data.elem.checked){
        index = 1;
    }else{
        index = -1;
    }
    $.ajax({
        url: '/admin/articles/recommend',
        data:{recommend:index,id:data.value},
        type: "post",
        success: function (data) {
            layer.msg('操作成功',{time:1000},function(){
               location.reload();
            });
        }
    });
});
form.on('checkbox(label)', function(data){
    var _url = $('.layui-form').attr('action');
    var id = $(data.elem).data('id');
    $.ajax({
        url: _url,
        data:{label:data.value,id:id,status:data.elem.checked},
        type: "post",
        success: function (data) {
            layer.msg(data.msg,{time:1000});
        }
    });
});
form.verify({
    username: function(value, item){ //value：表单的值、item：表单的DOM对象
        if(!new RegExp("^[a-zA-Z0-9_\u4e00-\u9fa5\\s·]+$").test(value)){
            return '用户名不能有特殊字符';
        }
        if(/(^\_)|(\__)|(\_+$)/.test(value)){
            return '用户名首尾不能出现下划线\'_\'';
        }
        if(/^\d+\d+\d$/.test(value)){
            return '用户名不能全为数字';
        }
    }

    //我们既支持上述函数式的方式，也支持下述数组的形式
    //数组的两个值分别代表：[正则匹配、匹配不符时的提示文字]
    ,pass: [
        /^[\S]{6,12}$/
        ,'密码必须6到12位，且不能出现空格'
    ]
    ,des_key:function(value,item){
        if(value.length!=16){
            return '请输入16位随机字符串';
        }
        if(isNumber(value)){
            return '请输入字符串、数字和符号的组合';
        }
    }
    ,des_iv:function(value,item){
        if(value.length!=8){
            return '请输入8位随机字符串';
        }
        if(isNumber(value)){
            return '请输入字符串、数字和符号的组合';
        }
    }
});
function isNumber(val){

    var regPos = /^\d+(\.\d+)?$/; //非负浮点数
    var regNeg = /^(-(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9][0-9]*)))$/; //负浮点数
    if(regPos.test(val) || regNeg.test(val)){
        return true;
    }else{
        return false;
    }
}

// post 提交
function $ajax_post($url,$param,fn)
{
    if(undefined == $url) return false;
    if(undefined == $param) return false;
    $.ajax({
        url:$url
        ,async:false
        ,type:"post"
        ,data: $param
        ,dataType: "json"
        ,success: function(msg){
            fn(msg)
        }
    });
}

