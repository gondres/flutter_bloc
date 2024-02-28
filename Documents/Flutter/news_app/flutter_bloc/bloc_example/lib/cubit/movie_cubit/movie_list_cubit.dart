import 'package:bloc_example/data/response/movie_list_response.dart';
import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../data/repository/movie_repository.dart';

part 'movie_list_state.dart';

class MovieCubit extends Cubit<MovieState> {
  MovieCubit() : super(MovieInitial());

  final MovieRepositories repositories = MovieRepositories();

  void getList() async {
    emit(MovieLoading());
    await repositories.getListMovie().then((value) {
      if (value.isSuccess && value.dataResponse is MovieListResponse) {
        final res = value.dataResponse as MovieListResponse;
        emit(MovieSuccess(res));
      } else {
        emit(MovieFailure(message: value.dataResponse));
      }
    });
  }
}
