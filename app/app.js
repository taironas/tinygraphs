'use strict';

// Declare app level module which depends on views, and components
angular.module('tinygraphs', [
    'ngRoute',
    'smoothScroll',
    'tinygraphs.view1',
    'tinygraphs.view2',
    'tinygraphs.version',
    'tinygraphs.holder',
    'tinygraphs.prettify'
]).
    config(['$routeProvider', function($routeProvider) {
	$routeProvider.otherwise({redirectTo: '/'});
    }]);
