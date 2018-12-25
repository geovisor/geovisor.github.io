
$(function(){
    $("#testNoBtn").bsSuggest({
        url: './data.json',
        emptyTip: '未搜索到您输入的小区，请联系工作人员查询',
        effectiveFields: ["x"],
        searchFields: ["x"],
        effectiveFieldsAlias:{qy: "区域",lm:"路名",xq:"小区"},
        // ignorecase: true,
        showHeader: false,
        getDataMethod: 'firstByUrl',
        showBtn: true,     //不显示下拉按钮
        delayUntilKeyup: true, //获取数据的方式为 firstByUrl 时，延迟到有输入/获取到焦点时才请求数据
        keyField: ["x"],
        clearable: false//更改了清除方式，所以此处失效
    }).on('onDataRequestSuccess', function (e, result) {
        console.log('onDataRequestSuccess: ', result);
    }).on('onSetSelectValue', function (e, keyword, data) {
        console.log('onSetSelectValue: ', keyword, data);
        $("#yb").html(data.y);
        $("#yf").html(data.f);

    }).on('onUnsetSelectValue', function () {
        console.log("onUnsetSelectValue");
    });


    //禁用表单提交
    $("form").submit(function() {
        return false;
    });


});
