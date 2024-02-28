import 'package:bloc_example/data/service/movie_service.dart';
import 'package:dio/dio.dart';
import 'package:get_it/get_it.dart';

import 'app_dio.dart';

class AppServices {
  final Dio dio;

  AppServices(this.dio);

  final get = GetIt.I;

  registerAppServices(String url) async {
    dio.interceptors.clear();
    dio.options.receiveTimeout = Duration(milliseconds: 35000);
    dio.options.connectTimeout = Duration(milliseconds: 30000);
    dio.interceptors.add(LogInterceptor(responseBody: true, requestBody: true, requestHeader: true, error: true));
    dio.interceptors.add(DioInterceptor());

    if (!get.isRegistered<MovieService>()) {
      get.registerFactory(() => MovieService(dio, baseUrl: url));
    } else {
      get.unregister<MovieService>();
      get.registerFactory(() => MovieService(dio, baseUrl: url));
    }
  }
}
