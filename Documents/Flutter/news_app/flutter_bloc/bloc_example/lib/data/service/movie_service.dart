import 'package:bloc_example/data/response/movie_list_response.dart';
import 'package:dio/dio.dart';
import 'package:retrofit/http.dart';
import 'package:bloc_example/base/base_path.dart';

part 'movie_service.g.dart';

@RestApi()
abstract class MovieService {
  factory MovieService(Dio dio, {String baseUrl}) = _MovieService;

  @GET(movie_list)
  Future<MovieListResponse> getListMovie();
}
