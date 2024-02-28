import 'package:json_annotation/json_annotation.dart';

part 'movie_list_response.g.dart';

@JsonSerializable(fieldRename: FieldRename.snake)
class MovieListResponse {
  int? id;
  int? page;
  List<Results>? results;
  int? totalPages;
  int? totalResults;

  MovieListResponse({this.id, this.page, this.results, this.totalPages, this.totalResults});

  MovieListResponse.fromJson(Map<String, dynamic> json) {
    id = json['id'];
    page = json['page'];
    if (json['results'] != null) {
      results = <Results>[];
      json['results'].forEach((v) {
        results!.add(new Results.fromJson(v));
      });
    }
    totalPages = json['total_pages'];
    totalResults = json['total_results'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['id'] = this.id;
    data['page'] = this.page;
    if (this.results != null) {
      data['results'] = this.results!.map((v) => v.toJson()).toList();
    }
    data['total_pages'] = this.totalPages;
    data['total_results'] = this.totalResults;
    return data;
  }
}

@JsonSerializable(fieldRename: FieldRename.snake)
class Results {
  String? description;
  int? favoriteCount;
  int? id;
  int? itemCount;
  String? iso6391;
  String? iso31661;
  String? listType;
  String? name;
  String? posterPath;

  Results({this.description, this.favoriteCount, this.id, this.itemCount, this.iso6391, this.iso31661, this.listType, this.name, this.posterPath});

  Results.fromJson(Map<String, dynamic> json) {
    description = json['description'];
    favoriteCount = json['favorite_count'];
    id = json['id'];
    itemCount = json['item_count'];
    iso6391 = json['iso_639_1'];
    iso31661 = json['iso_3166_1'];
    listType = json['list_type'];
    name = json['name'];
    posterPath = json['poster_path'];
  }

  Map<String, dynamic> toJson() {
    final Map<String, dynamic> data = new Map<String, dynamic>();
    data['description'] = this.description;
    data['favorite_count'] = this.favoriteCount;
    data['id'] = this.id;
    data['item_count'] = this.itemCount;
    data['iso_639_1'] = this.iso6391;
    data['iso_3166_1'] = this.iso31661;
    data['list_type'] = this.listType;
    data['name'] = this.name;
    data['poster_path'] = this.posterPath;
    return data;
  }
}
