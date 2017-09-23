
	var smarker;
	var percentagetext = "%";
	function havop(url,target){
		vpath = url.split("?");
		$.ajax({url:vpath[0], data: getJsonFromUrl(vpath[1]),type:"POST",success:function(html){
			$(target).html(html);
		},error:function(e){
			$(target).html(e.responseText);
		}
	});
	}
	//psub --
	//var smarker;
	(function( $ ) {
 
    $.fn.havok = function() {
	
	$(".aput button",this).click(function(e){
		var asub = $(this).parents(".aput");
		var splitz = asub.attr("path").split("?");
		var datav = getJsonFromUrl(splitz[1]);
		datav[$("input",asub).attr("name")] = $("input",asub).val();
				$(".hidden-aspect", asub).css('display', 'block');
		if($("input", asub).val() == ""){
			$("input",asub).css("border-color","red");
		} else {
		$("input", asub).removeAttr("style");
		datav["put"] = $("input", asub).attr("data-key")
		$.ajax({url: splitz[0],type:"POST", data: datav, error:function(e){	1
			$(".hidden-aspect", asub).css('display', 'none');
			$(".side-bay").html("");
		}});
		}
	});
	
	//used in app loading
	
	$(".prefixsub",this).each(function(e,i){
		var prefix = $(this).attr("prefix");
		var prefixsib = $(this).attr("target");
		$(".selfsub", this).click(function(e){
			$(".loader-prog").css('display', 'block');
			$.ajax({url: prefix + $(this).attr("path"), success:function(html){	
				
				$(".loader-prog").css('display', 'none');
				$(prefixsib).html(html);
				
				
			}});
		});
	});
	
	$(".psub",this).each(function(e,i){
		
		var path = $(this).attr("path");
		$('.xsub',this).change(function(e){
			$(".loader-prog").css('display', 'block');
			$.ajax({url: path, type:"POST", data: {target:$(this).attr("name"),data: $(this).val()},success:function(html){
				
				
				
				$(".loader-prog").css('display', 'none');
				$('.alert-stored').html(html);
				
				
			}});
			
		});
		
		$('.isub',this).change(function(e){
			$(".loader-prog").css('display', 'block');
			$.ajax({url: path, type:"POST", data: {target:$(this).attr("name"),data: $(this).is(":checked")},success:function(html){
				
				
				
				$(".loader-prog").css('display', 'none');
				$('.alert-stored').html(html);
				
				
			}});
			
		});
		
		
		
		
	});
	
	$(".xtrigger",this).click(function(e){
		var urlz = ( "" + $(this).attr("path") ).split("?");
		$(".loader-prog").css('display', 'block');
		$.ajax({url:urlz[0],type:"POST", data: getJsonFromUrl(urlz[1]), error:function(e){
			
			$(".loader-prog").css('display', 'none');
			$(".footer-bay").html(e.responseText);
			
		}});
	});
	
	$(".strigger",this).click(function(e){
		
		$(".loader-prog").css('display', 'block');
		$.ajax({url: $(this).attr("path"), success:function(html){
			
		
			$(".loader-prog").css('display', 'none');
			$('.content-window').html(html);
			
			
		}});
	});

	
	$(".marker",this).each(function(e,i){
			
		//get formdata
		
		var marker = $(this);
		var path = $(this).attr("path");
		//smarker = marker
		
		if(marker.attr("data-link") != "") {
			$(".loader-result",marker).load(marker.attr("data-link"));
					
		}
		$('.btn-danger',this).click(function(e){
			 smarker = $(this).closest(".marker");
		    var vpath = $(this).attr("path").split("?")
		   	var req = getJsonFromUrl(vpath[1]);
		   	
		  $.ajax({url: vpath[0], data: req, type:"POST", error:function(e) { 
				//for some reason it goes here
				//console.log(smarker)
				smarker.parent().remove();
			},success:function(html){
					smarker.parent().remove();
			}});
		  	return false;
		});
		$('.dosub', this).click(function(e){
			var cancel = false;
			 smarker = $(this).closest(".marker");
		    var vpath = smarker.attr("path").split("?");
			var req = getJsonFromUrl(vpath[1]);
			$('.xsub', smarker).each(function(e,i){
				if($(this).closest(".marker").attr("path") == smarker.attr("path")){
				req[$(this).attr("name")] = $(this).val();
				if($(this).hasClass("required") && $(this).val() == ""){
					$(this).css('border-color','red');
					cancel = true;
					$(this).tooltip({title:"Required field"});
					$(this).tooltip('show');
				}
			}

			});
			
			//execution
			if(!cancel){
			$(this).button('toggle');
			$('.loader-form').css('display', 'none');
			$('.hidden-aspect', smarker).css('display', 'block');
		$(".loader-result",smarker).html("") 
			$.ajax({url: vpath[0], data: req, type:"POST", error:function(e) { 
				//for some reason it goes here
				
				if ( $(smarker).attr("data-link") != ""){
					$('.hidden-aspect').removeAttr('style');
				//	$(".xsub", smarker).val("");
					$(".loader-result",smarker).load(smarker.attr("data-link"));
					
				}  else  {
				$('.loader-form').removeAttr('style');
				$('.hidden-aspect').removeAttr('style');
			//	$(".xsub", smarker).val("");
				$('.loader-result', smarker).html("<label class=\"label label-primary\">" + e.status +"</label>" + e.responseText);
			
				}
				
				if ($(".loader-result",smarker).html() == "")
					$(".loader-result",smarker).html(e.responseText)

				if($(".side-bay").html() != ""){
					$(".side-bay").html("");
					$("#jstree").jstree("refresh")
				}
			},success:function(html){
				
				//console.log(html);
						//console.log(smarker)
			if (smarker.attr("data-link") != ""){
					$('.hidden-aspect').removeAttr('style');
					$(".xsub", smarker).val("");
					$(".loader-result",smarker).load(smarker.attr("data-link"));
					
				} else {
				$('.loader-form').removeAttr('style');
				$('.hidden-aspect').removeAttr('style');
				//$(".xsub", smarker).val("");
				console.log(html);
				$('.loader-result', smarker).html("<label class=\"label label-primary\">" + e.status +"</label>" + e.responseText);
				
				}

				if($(".side-bay").html() != ""){
					$(".side-bay").html("");
					$("#jstree").jstree("refresh")
				}
			}});
			}
			
			return false;
		});
		
		
		$('.xsub',this).each(function(e,i){
			$(this).keypress(function(e) {
				  if (e.which == '13') {
				     e.preventDefault();
				     $(".dosub", marker).click();
				       return false;

				   }


				 
				});
		});
		
		
	});

   return this;
 
    };
 
}( jQuery ));

function getJsonFromUrl(query) {
	if(query){
  var result = {};
  query.split("&").forEach(function(part) {
    var item = part.split("=");
    result[item[0]] = decodeURIComponent(item[1]);
  });
  return result;
	}
	return {};
}

$(".marker[init='yes']").parent().havok();

$(".auto-marker").havok();


