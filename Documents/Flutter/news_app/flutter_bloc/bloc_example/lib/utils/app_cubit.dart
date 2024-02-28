import 'package:bloc_example/cubit/movie_cubit/movie_list_cubit.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class AppCubit {
  const AppCubit._();

  static Widget initCubit(Widget widget) {
    return MultiBlocProvider(providers: [
      BlocProvider<MovieCubit>(
        create: (BuildContext context) => MovieCubit(),
      ),
    ], child: widget);
  }
}
