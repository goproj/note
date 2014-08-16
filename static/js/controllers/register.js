angular.module('starter.controllers.Register', [])

.controller('RegisterCtrl', function($rootScope, $scope, $state, Auth, Util) {
	$scope.registerForm = {};

	$scope.register = function() {
		var username = $scope.registerForm.username ? $scope.registerForm.username : $scope.registerForm.email;
		var email = $scope.registerForm.email;
		var password = $scope.registerForm.password;
		var confirmPassword = $scope.registerForm.confirmPassword;


		if (!email) {
			Util.toast('请输入正确邮箱');
		} else if (!password) {
			Util.toast('请输入密码');
		} else if (password !== confirmPassword) {
			Util.toast('两次输入密码不一致');
		} else {
			Auth.register(username, email, password).then(function(result) {
				if (result.err !== 0) {
					Util.toast(result.msg);
				}else{
					Auth.setAuth();
					$state.go('app.note');
				}
			});
		}
	}
});