angular.module('starter.controllers.Note', [])

.controller('NoteCtrl', function($scope, $state, $ionicModal, Note) {

	var titles = {
		finished: '历史',
		unfinished: '待办'
	};

		

	var state = $state.params.state;
	var loadNotes = function() {
		if (state === 'finished') {
			return Note.getNotes(true).then(function(notes){
				$scope.notes = notes;
			});
		} else if (state === 'unfinished') {
			return Note.getNotes(false).then(function(notes){
				$scope.notes = notes;
			});
		}
	};

	

	$scope.title = titles[state];
	$scope.notes = [];
	$scope.shouldShowDelete=false;

	loadNotes();

	

	$scope.closeEditNoteModal = function() {
		$scope.modal.hide();
	};

	$scope.toggleShowDelete=function(){
		$scope.shouldShowDelete=!$scope.shouldShowDelete;
	};


	$scope.delete=function($event,id){
		$event.stopPropagation();
		Note.deleteNote(id).then(function(){
	 		loadNotes();
	 	});
	};

	$scope.openEditNoteModal = function(note) {
		$scope.currentNote = angular.copy(note);
		$scope.modal.show();
	};

	$scope.edit = function(id,content) {
		Note.editNote(id,content).then(function(data) {
			loadNotes();
			$scope.modal.hide();
		});
	};

	$scope.toggleDone = function(note) {
		var done=note.done?0:1;
		var id=note.id;

		Note.mark(id,done).then(function() {
			loadNotes();
		});
	};

	$scope.noteDetail=function(note){
		Note.currentNote=note;
		$state.go('app.noteDetail',{id:note.id});
	};

	$scope.$on('note:add', function(e, note) {
		if (state!=='finished'){
			$scope.notes.unshift(note);
		}
	});

});