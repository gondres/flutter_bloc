import 'dart:io';

import 'package:bloc_example/data/service/movie_service.dart';
import 'package:dio/dio.dart';
import 'package:get_it/get_it.dart';

import '../model/repositories_response.dart';

class MovieRepositories {
  Future<RepositoriesResponse> getListMovie() async {
    final services = GetIt.I.get<MovieService>();

    late RepositoriesResponse response;

    try {
      await services.getListMovie().then((value) {
        response = RepositoriesResponse(isSuccess: true, dataResponse: value);
        print('sukes ${value}');
      });
    } catch (e) {
      if (e is IOException) {
        response = RepositoriesResponse(isSuccess: false, dataResponse: e.toString());
        print('io exception');
      } else {
        print(' exception ${e.toString()}');
        response = RepositoriesResponse(isSuccess: false, dataResponse: e.toString());
      }
      if (e is DioException) {
        print('dio exception');
        response = RepositoriesResponse(isSuccess: false, dataResponse: e.response?.data['message'].toString() ?? 'Dio Exception');
      }
    }
    print(response);
    return response;
  }
}
