'use strict' 

angular.module('tinygraphs.prettify.prettify-directive', [])

.directive('prettyprint', function(){
    return{
	restrict: 'C',
	link: function postLink(scope, element, attrs){
	    element.html(prettyPrintOne(replaceText(element.html())));
	}
    };
});

function replaceText(str){
    var str1 = String(str);
    return str1.replace(/\n/g,"<br/>");
};
