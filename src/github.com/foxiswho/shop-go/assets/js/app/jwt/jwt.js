define(["jquery", "jquery-form", "bootstrap"], function ($, form, bootstrap) {

	$("#foo").click(function(event) {
		// jwtTester();
        token = $('#token').val();
        console.log(token)
        console.log(url)
        url = $('#url').val();
        $.ajax({
            // async: false,
            type: "GET",//方法类型
            url: url,
            dataType: 'JSON',
            // jsonpCallback:'aajsonp',
            beforeSend: function(request) {
                request.setRequestHeader("Authorization","Bearer "+ token);
                console.log(request)
                // return request;
            },
            success: function (data) {
                // alert("成功");
            },
            error: function () {
                // alert("失败");
            },
            complete: function (xhr) {
                console.log(xhr)
                $('#api_data').val(JSON.stringify(xhr.responseJSON));
            },
        });
	});

	function jwtTester() {
	    token = $('#token').val();
	    console.log(token)
	    url = $('#url').val();
	    $.ajax({
	        // async: false,
            type: "GET",//方法类型
	        url: url,
            dataType: 'JSON',
			// jsonpCallback:'aajsonp',
	        beforeSend: function(request) {
                request.setRequestHeader("Authorization","Bearer "+ token);
                console.log(request)
                // return request;
            },
	        success: function (data) {
	            // alert("成功");
	        },
	        error: function () {
	            // alert("失败");
	        },
	        complete: function (xhr) {
	            console.log(xhr)
	            $('#api_data').val(JSON.stringify(xhr.responseJSON));
	        },
	    });
	}
});