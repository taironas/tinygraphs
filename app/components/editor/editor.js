'use strict';

angular.module('tinygraphs.editor', [])
    .controller('EditorCtrl', ['$scope', '$location', '$routeParams', EditorCtrl]);

function EditorCtrl($scope, $location, $routeParams) {

    var init = function () {
	if($location.$$hash == 'tryitout'){
	    $scope.editorName = $location.search().name;
	    $scope.editorShape = $location.search().shape;
	    $scope.editorTheme = $location.search().theme;
	    $scope.editorColorNumber = $location.search().numcolors;
	}else{
	    $scope.editorName = 'tinygraphs';
	    $scope.editorShape = 'squares';
	    $scope.editorTheme = 'frogideas';
	    $scope.editorColorNumber = '4';
	}
	$scope.absUrl = $location.absUrl();
    };
    init();

    $scope.editorChanged = function(){
	$location.hash('tryitout')
	    .search({
		name: $scope.editorName, 
		shape: $scope.editorShape, 
		theme: $scope.editorTheme, 
		numcolors: $scope.editorColorNumber
	    });
	$scope.absUrl = $location.absUrl();
    };    
}
