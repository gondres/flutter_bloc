import 'package:dio/dio.dart';
import 'package:bloc_example/base/base_param.dart' as baseParam;

import 'app_preferences.dart';

class DioInterceptor extends Interceptor {
  @override
  Future<void> onRequest(RequestOptions options, RequestInterceptorHandler handler) async {
    SharedPref.getToken().then((value) => options.headers[baseParam.authorization] = 'Bearer ${value}');
    options.headers[baseParam.content_type] = baseParam.app_json;
    super.onRequest(options, handler);
  }
}
