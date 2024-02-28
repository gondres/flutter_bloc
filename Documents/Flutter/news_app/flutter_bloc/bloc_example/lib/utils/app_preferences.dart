import 'package:shared_preferences/shared_preferences.dart';
import 'package:bloc_example/base/base_prefs.dart';

class SharedPref {
  const SharedPref._();

  static const String prefKey = 'bloc_pref';

  static Future<void> setToken(String token) async {
    final prefs = await SharedPreferences.getInstance();

    await prefs.setString(key_token, token);
  }

  static Future<String?> getToken() async {
    final prefs = await SharedPreferences.getInstance();

    return prefs.getString(key_token);
  }
}
