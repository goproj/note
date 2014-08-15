angular.module('starter.services.Note', [])

.factory('Note', function($localForage, Util) {
	var Note = {
		addNote: function(note) {
			return $localForage.getItem('notes').then(function(notes) {
				return notes ? notes : [];
			}).then(function(notes) {
				notes.unshift(note);
				return $localForage.setItem('notes', notes);
			}).then(function() {
				Util.toast('添加任务成功');
			});
		},
		getNotes: function() {
			return $localForage.getItem('notes').then(function(notes) {
				return notes ? notes : [];
			});
		},
		deleteNote: function(id) {
			return Note.getNotes().then(function(notes) {
				_.remove(notes, function(item) {
					return id === item.id;
				});
				return notes;
			}).then(function(notes) {
				return $localForage.setItem('notes', notes);
			});
		},
		editNote: function(note) {
			return Note.getNotes().then(function(notes) {
				var index=_.findIndex(notes, function(item) {
					return note.id === item.id;
				});
				notes[index]=note;
				return notes;
			}).then(function(notes) {
				return $localForage.setItem('notes', notes);
			});
		}

	}
	return Note;



})