import 'package:flutter_dotenv/flutter_dotenv.dart';

String baseURL() {
  String base = dotenv.env['backend_url']!;
  return base;
}

Uri createURL(String path) {
  final url = Uri.parse(baseURL() + path);
  return url;
}

String getAPIKey() {
  return dotenv.env['api_key']!;
}
