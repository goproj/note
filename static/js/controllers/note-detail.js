angular.module('starter.controllers.NoteDetail', [])

.controller('NoteDetailCtrl', function($scope, $state, $ionicModal, Note) {

	$ionicModal.fromTemplateUrl('modal/edit-note.html', {
		scope: $scope
	}).then(function(modal) {
		$scope.modal = modal;
	});

	$scope.note = Note.currentNote;

	$scope.closeEditNoteModal = function() {
		$scope.modal.hide();
	};

	$scope.openEditNoteModal = function() {
		$scope.editNote = angular.copy($scope.note);
		$scope.modal.show();
	};

	$scope.edit = function(id, content) {
		Note.editNote(id, content).then(function(data) {
			$scope.note=$scope.editNote;
			$scope.modal.hide();
		});
	};

});