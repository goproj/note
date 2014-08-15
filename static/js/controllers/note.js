angular.module('starter.controllers.Note', [])

.controller('NoteCtrl', function($scope, $state, $ionicModal, Note) {

	var titles = {
		all: '全部',
		finished: '已完成',
		unfinished: '待完成'
	};

	var state = $state.params.state;
	var notes=[];
	var filter=function(notes){
		if (state === 'finished') {
			notes = _.filter(notes, function(note) {
				return note.done === true;
			});
		} else if (state === 'unfinished') {
			notes = _.filter(notes, function(note) {
				return note.done === false;
			});
		}

		$scope.notes = notes;
	}

	$scope.title = titles[state];
	$scope.notes = [];

	Note.getNotes().then(function(items) {
		notes=items;
		filter(notes);
	});

	$ionicModal.fromTemplateUrl('modal/edit-note.html', {
		scope: $scope
	}).then(function(modal) {
		$scope.modal = modal;
	});

	$scope.closeEditNoteModal = function() {
		$scope.modal.hide();
	};

	$scope.delete = function(note) {
		_.remove($scope.notes, function(item) {
			return note.id===item.id;
		});
		Note.deleteNote(note.id);
	};

	$scope.openEditNoteModal = function(note) {
		$scope.currentNote = note;
		$scope.modal.show();
	};

	$scope.edit=function(note){
		Note.editNote(note).then(function(){
			$scope.modal.hide();
		});
	};

	$scope.toggleDone = function(note) {
		note.done=!note.done;
		Note.editNote(note).then(function(){
			filter(notes);
		});
	};

	$scope.$on('note:add', function(e, note) {
		$scope.notes.unshift(note);
	});



});