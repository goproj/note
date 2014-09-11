angular.module('starter.services.Note', [])

.factory('Note', function($localForage, $http, Util) {
	var Note = {
		addNote: function(note) {
			return $http.post('/note', note).then(function(result) {
				return result.data.data;
			});
		},
		getNotes: function(done) {

			if (done) {
				return $http.get('/note/done').then(function(result) {
					return result.data.data;
				});
			} else {
				return $http.get('/note/todo').then(function(result) {
					return result.data.data;
				});
			}
		},
		deleteNote: function(id) {
			return $http.delete('/note/'+id);
		},
		editNote: function(note) {
			return $http.put('/note/'+id,{
				content:content
			}).then(function(result){
				alert(JSON.stringify(result));
			});
		},
		mark:function(id,done){
			return $http.get('/note/'+id+'/mark/'+done);
		}

	}
	return Note;



})