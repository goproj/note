angular.module('starter.services.Auth', [])

.factory('Auth', function($localForage, $window,$http,Util) {
	var Auth = {
		isAuth:function(){
			return $window.localStorage.getItem('auth')==='true';
		},
		setAuth:function(){
			$window.localStorage.setItem('auth','true');
		},
		setUnauth:function(){
			$window.localStorage.setItem('auth','false');
		},
		register:function(username,email,password){
			var registerForm={
				username:username,
				password:CryptoJS.MD5(password).toString(),
				email:email
				
			};

			return $http.post('/register',registerForm).then(function(result){
				return result.data;
			});
		},
		login:function(email,password){
			var loginForm={
				email:email,
				password:CryptoJS.MD5(password).toString()
			};
			return $http.post('/login',loginForm).then(function(result){
				return result.data;
			});
		},
		logout:function(){
			return $http.get('/logout').then(function(result){
				return result.data;
			});
		}

	}
	return Auth;
});