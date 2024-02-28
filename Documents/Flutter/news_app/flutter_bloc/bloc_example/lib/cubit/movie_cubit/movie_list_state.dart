part of 'movie_list_cubit.dart';

abstract class MovieState extends Equatable {
  const MovieState();

  @override
  List<Object?> get props => [];
}

class MovieInitial extends MovieState {}

class MovieSuccess extends MovieState {
  final MovieListResponse response;

  const MovieSuccess(this.response);

  @override
  List<Object?> get props => [response];
}

class MovieLoading extends MovieState {}

class MovieFailure extends MovieState {
  final String message;

  const MovieFailure({required this.message});

  @override
  List<Object?> get props => [message];
}
