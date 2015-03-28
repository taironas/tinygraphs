'use strict';

// Declare app level module which depends on views, and components
angular.module('tinygraphs', [
    'ngRoute',
    'smoothScroll',
    'tinygraphs.view1',
    'tinygraphs.view2',
    'tinygraphs.version',
    'tinygraphs.holder',
    'tinygraphs.prettify',
    'tinygraphs.editor'
]).
    config(['$locationProvider', '$routeProvider', 
	    function($locationProvider, $routeProvider) {
		$routeProvider.otherwise({redirectTo: '/'});
	    }
	   ]);
