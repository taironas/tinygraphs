'use strict';

angular.module('tinygraphs.editor', [])
    .controller('EditorCtrl', ['$scope', '$location', EditorCtrl]);

function EditorCtrl($scope, $location) {
    $scope.editorChanged = function(){
	$location.search({
	    name: $scope.editorName, 
	     shape: $scope.editorShape, 
	     theme: $scope.editorTheme, 
	     numcolors: $scope.editorColorNumber
	});
    };    
}
