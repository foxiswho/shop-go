define(["jquery", "jquery-form", "bootstrap"], function ($, form, bootstrap) {

	$("#foo").click(function(event) {
		jwtTester();
	});

	function jwtTester() {
	    token = $('#token').val();
	    url = $('#url').val();
	    $.ajax({
	        async: false,
	        url: url,
            dataType: 'jsonp',
			jsonpCallback:'aajsonp',
	        beforeSend: function(request) {
                request.setRequestHeader("Authorization", "Bearer "+token);
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