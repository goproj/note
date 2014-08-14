angular.module('starter.services.Util', [])

.factory('Util', function($ionicLoading) {
	var Util={
		toast:function(content){
			$ionicLoading.show({
				template:content,
				duration:1000
			});
		}

	}
	return Util;
	


})