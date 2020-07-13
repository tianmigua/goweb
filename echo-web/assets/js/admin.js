/**
 * 后台侧边菜单选中状态
 */
$('.layui-nav-item').find('a').removeClass('layui-this');
//$('.layui-nav-tree').find('a[href="' + GV.current_controller + '"]').parent().addClass('layui-this').parents('.layui-nav-item').addClass('layui-nav-itemed');
function logout(){
    layer.open({
        shade: false,
        content: '确定退出？',
        btn: ['确定', '取消'],
        yes: function (index) {
            $.ajax({
                url: '/logout',
                type: "post",
                success: function (data) {
                    if(data.code==1){
                        window.location.href = '/login'
                    }
                }
            });
            layer.close(index);
        }
    });
}
function clearCache()
{
    layer.open({
        shade: false,
        content: '确定清除？',
        btn: ['确定', '取消'],
        yes: function (index) {
            $.post(
                '/index/clear_cache',
                function(data){
                    layer.msg(data.msg);
                }
            )
            layer.close(index);
        }
    });
}
function sortLaravel(e){

    $.ajax({
        url: $(e).data('action'),
        type: "post",
        data:{sort:$(e).val(),id:$(e).data('id')},
        success: function (data) {
            if(data.code==200){
                setTimeout(function(){
                    location.reload();
                },1000)
            }
            layer.msg(data.msg);
        }
    });
}
function setPrice(id,url,type)
{
    layer.prompt(
        {
            title: '请输入定价',
            formType: 0,
            maxlength: 10
        },
        function(price, index){
            $.ajax({
                url: url,
                type: "post",
                data:{id:id,price:price,type:type},
                success: function (data) {
                    layer.msg(data.msg,{time:1000},function(){
                        location.reload();
                    });
                }
            });
        }
    );
}
function setPrice1(e)
{
    var id = $(e).data('id');
    var url = $(e).data('action');
    var type = $(e).attr('name');
    if(!isNumber($(e).val())){
        layer.msg('请填写正确数字');
        return false;
    }
    $.ajax({
        url: url,
        type: "post",
        data:{id:id,price:$(e).val(),type:type},
        success: function (data) {
            layer.msg(data.msg,{time:1000});
        }
    });
}
function toEpg(id,url)
{
    layer.confirm('确认发布到EPG？', {
        btn: ['确认','取消'] //按钮
    }, function(){
        $.ajax({
            url: url,
            type: "post",
            data:{id:id},
            success: function (data) {
                if(data.code==200){
                    setTimeout(function(){
                        location.reload();
                    },1000)
                }
                layer.msg(data.msg);
            }
        });
    }, function(index){
        layer.close(index);
    });
}
function permission(e)
{
    if($(e).is(':checked')){
        $('.parent_'+$(e).val()).prop('checked',true);
    }else{
        $('.parent_'+$(e).val()).prop('checked',false);
    }
}