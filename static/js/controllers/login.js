angular.module('starter.controllers.Login', [])

.controller('LoginCtrl', function($rootScope, $scope, $location, $state, Auth, Util) {
	$scope.loginForm = {};

	$scope.login = function() {
		var email = $scope.loginForm.email;
		var password = $scope.loginForm.password;

		if (!email) {
			Util.toast('请输入邮箱');
		} else if (!password) {
			Util.toast('请输入密码');
		} else {
			Auth.login(email, password).then(function(result) {
				if (result.err !== 0) {
					Util.toast(result.msg);
				} else {
					Auth.setAuth();
					$state.go('app.note');
				}
			});
		}
	}
});