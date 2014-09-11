angular.module('starter.controllers.Note', [])

.controller('NoteCtrl', function($scope, $state, $ionicModal, Note) {

	var titles = {
		all: '全部',
		finished: '已完成',
		unfinished: '待完成'
	};

		

	var state = $state.params.state;
	var filter = function() {
		if (state === 'finished') {
			Note.getNotes(true).then(function(notes){
				$scope.notes = notes;
			});
		} else if (state === 'unfinished') {
			Note.getNotes(false).then(function(notes){
				$scope.notes = notes;
			});
		}
	};

	

	$scope.title = titles[state];
	$scope.notes = [];

	filter();

	$ionicModal.fromTemplateUrl('modal/edit-note.html', {
		scope: $scope
	}).then(function(modal) {
		$scope.modal = modal;
	});

	$scope.closeEditNoteModal = function() {
		$scope.modal.hide();
	};

	$scope.delete = function(note) {
		Note.deleteNote(note.id).then(function(){
			filter();
		});
	};

	$scope.openEditNoteModal = function(note) {
		$scope.currentNote = angular.copy(note);
		$scope.modal.show();
	};

	$scope.edit = function(id,content) {
		Note.editNote(id,note).then(function() {
			$scope.modal.hide();
		});
	};

	$scope.toggleDone = function(note) {
		var done=note.done?0:1;
		var id=note.id;

		Note.mark(id,done).then(function() {
			filter();
		});
	};

	$scope.$on('note:add', function(e, note) {
		if (state!=='finished'){
			$scope.notes.unshift(note);
		}
	});

});