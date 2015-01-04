'use strict';

angular.module('myApp.holder.holder-directive', [])

.directive('holderFix', ['version', function() {
  return {
    link: function(scope, element, attrs) {
	attrs.$set('data-src', attrs.holderFix);
	Holder.run({images:element[0]});
    }
  };
}]);
