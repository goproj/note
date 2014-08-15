angular.module('starter.controllers.Menu', [])

.controller('MenuCtrl', function($rootScope, $scope, $ionicModal, Note, Util) {

	$scope.note = {};

	$ionicModal.fromTemplateUrl('modal/add-note.html', {
		scope: $scope
	}).then(function(modal) {
		$scope.modal = modal;
	});

	$scope.closeAddNoteModal = function() {
		$scope.modal.hide();
	},

	$scope.openAddNoteModal = function() {
		$scope.modal.show();
	};

	$scope.addNote = function() {
		$scope.note.createdAt = Date.now();
		$scope.note.id = $scope.note.createdAt;
		$scope.note.done = false;

		if (!$scope.note.content||$scope.note.content.trim() === '') {
			Util.toast('请输入内容');
		} else {
			Note.addNote($scope.note).then(function() {
				$scope.modal.hide();
				$rootScope.$broadcast('note:add', $scope.note);
				$scope.note = {};
			});
		}

	}
});