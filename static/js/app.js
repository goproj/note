angular.module('starter', ['ionic', 'starter.controllers', 'starter.services', 'LocalForageModule'])

.run(function($ionicPlatform, $state, $rootScope, $location, Auth) {
	$ionicPlatform.ready(function() {
		if (window.cordova && window.cordova.plugins.Keyboard) {
			cordova.plugins.Keyboard.hideKeyboardAccessoryBar(true);
		}
		if (window.StatusBar) {
			StatusBar.styleDefault();
		}
	});

	$rootScope.$on('$stateChangeStart',
		function(event, toState, toParams, fromState, fromParams) {
			if (toState.name.substring(0,3)==='app' && !Auth.isAuth()) {
				$location.path('/login')
			}
		});


})

.config(function($stateProvider, $urlRouterProvider) {
	$stateProvider

	.state('login', {
		url: '/login',
		templateUrl: '/static/templates/login.html',
		controller: 'LoginCtrl'
	})

	.state('register', {
		url: '/register',
		templateUrl: '/static/templates/register.html',
		controller: 'RegisterCtrl'
	})

	.state('app', {
		url: "/app",
		abstract: true,
		templateUrl: '/static/templates/menu.html',
		controller: 'MenuCtrl'
	})

	.state('app.note', {
		url: "/note/:state",
		views: {
			'menuContent': {
				templateUrl: '/static/templates/note.html',
				controller: 'NoteCtrl'
			}
		}
	});



	$urlRouterProvider.otherwise('/app/note/all');
});