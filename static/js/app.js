angular.module('starter', ['ionic', 'starter.controllers','starter.services','LocalForageModule'])

.run(function($ionicPlatform) {
  $ionicPlatform.ready(function() {
    if(window.cordova && window.cordova.plugins.Keyboard) {
      cordova.plugins.Keyboard.hideKeyboardAccessoryBar(true);
    }
    if(window.StatusBar) {
      StatusBar.styleDefault();
    }
  });
})

.config(function($stateProvider, $urlRouterProvider) {
  $stateProvider

    .state('app', {
      url: "/app",
      abstract: true,
      templateUrl: "/static/templates/menu.html",
      controller: 'MenuCtrl'
    })

    .state('app.note',{
      url:"/note/:state",
      views:{
        'menuContent':{
          templateUrl: "/static/templates/note.html",
          controller:'NoteCtrl'
        }
      }
    })

   
  $urlRouterProvider.otherwise('/app/note/all');
});

