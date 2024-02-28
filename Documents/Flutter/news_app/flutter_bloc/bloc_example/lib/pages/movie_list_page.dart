import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../cubit/movie_cubit/movie_list_cubit.dart';

class MovieListPage extends StatefulWidget {
  const MovieListPage({super.key});

  @override
  State<MovieListPage> createState() => _MovieListPageState();
}

class _MovieListPageState extends State<MovieListPage> {
  late MovieCubit cubit;

  @override
  void initState() {
    cubit = context.read<MovieCubit>();
    fetchList();
    super.initState();
  }

  void fetchList() {
    cubit.getList();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(),
        body: Padding(
          padding: const EdgeInsets.all(15.0),
          child: Container(
            child: BlocBuilder<MovieCubit, MovieState>(builder: (context, state) {
              if (state is MovieLoading) {
                return Center(child: CircularProgressIndicator());
              }
              if (state is MovieSuccess) {
                return ListView.builder(
                    itemCount: state.response.results?.length ?? 0,
                    itemBuilder: (builder, index) {
                      return Text(state.response.results?[index].name ?? '-');
                    });
              }
              if (state is MovieFailure) {
                return Text('Error Data');
              }
              return Container();
            }),
          ),
        ));
  }
}
