var app = angular.module("myApp",[]);

var config = {
    headers: {
        'Content-Type': 'application/x-www-form-urlencoded;charset=utf-8;'
    }
};

app.controller("PersonController", ['$scope', '$http', '$location', function ($scope, $http, $location) {
    Users = function(){
        $http({
            method: 'GET',
            url: '/People'
        }).then(function (response) {
            console.log(response);
            $scope.ListOfUsers = response.data;
        }, function (error) {
            $location.path('/');
        });
    }
    Users();
    $scope.Insert = function () {
        var UserData = {
            name: $scope.Name,
            lastName: $scope.LastName
        };
        UserData = JSON.stringify(UserData);
        $http({
            method: 'Post',
            url: '/InsertPeople',
            data: UserData,
            config
        }).then(function (response) {
            console.log(response);
            alert("Agregado");
            location.reload();
        }, function (error) {
            console.log(error.data);
            $scope.error = "Something wrong" + data.ExceptionMessage;
        });
    };
    $scope.Delete = function () {
        id = $scope.Id;
        $http({
            method: 'DELETE',
            url: '/DeletePeople/'+id,
            config
        }).then(function (response) {
            alert("Eliminado");
            location.reload();
        }, function (error) {
            console.log(error.data);
            $scope.error = "Something wrong" + data.ExceptionMessage;
        });
    };
    $scope.Edit = function () {
        id = $scope.Id;
        var UserData = {
            name: $scope.Name,
            lastName: $scope.LastName
        };
        UserData = JSON.stringify(UserData);
        $http({
            method: 'Post',
            url: '/UpdatePeople/'+id,
            data: UserData,
            config
        }).then(function (response) {
           alert("Actualizado");
           location.reload();
        }, function (error) {
            console.log(error.data);
            $scope.error = "Something wrong" + data.ExceptionMessage;
        });
    };
}]);