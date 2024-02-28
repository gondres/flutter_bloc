// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'movie_list_response.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

MovieListResponse _$MovieListResponseFromJson(Map<String, dynamic> json) =>
    MovieListResponse(
      id: json['id'] as int?,
      page: json['page'] as int?,
      results: (json['results'] as List<dynamic>?)
          ?.map((e) => Results.fromJson(e as Map<String, dynamic>))
          .toList(),
      totalPages: json['total_pages'] as int?,
      totalResults: json['total_results'] as int?,
    );

Map<String, dynamic> _$MovieListResponseToJson(MovieListResponse instance) =>
    <String, dynamic>{
      'id': instance.id,
      'page': instance.page,
      'results': instance.results,
      'total_pages': instance.totalPages,
      'total_results': instance.totalResults,
    };

Results _$ResultsFromJson(Map<String, dynamic> json) => Results(
      description: json['description'] as String?,
      favoriteCount: json['favorite_count'] as int?,
      id: json['id'] as int?,
      itemCount: json['item_count'] as int?,
      iso6391: json['iso6391'] as String?,
      iso31661: json['iso31661'] as String?,
      listType: json['list_type'] as String?,
      name: json['name'] as String?,
      posterPath: json['poster_path'] as String?,
    );

Map<String, dynamic> _$ResultsToJson(Results instance) => <String, dynamic>{
      'description': instance.description,
      'favorite_count': instance.favoriteCount,
      'id': instance.id,
      'item_count': instance.itemCount,
      'iso6391': instance.iso6391,
      'iso31661': instance.iso31661,
      'list_type': instance.listType,
      'name': instance.name,
      'poster_path': instance.posterPath,
    };
