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
				
			};
			registerForm.username=username;
			registerForm.password=CryptoJS.MD5(password).toString();
			registerForm.email=email;

			return $http.get('/register',{
				params:registerForm
			}).then(function(result){
				return result.data;
			});
		},
		login:function(email,password){
			var loginForm={
				email:email,
				password:CryptoJS.MD5(password).toString()
			};
			return $http.get('/login',{
				params:loginForm
			}).then(function(result){
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