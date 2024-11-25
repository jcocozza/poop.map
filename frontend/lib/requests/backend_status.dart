import 'dart:async';
import 'package:http/http.dart' as http;
import 'package:frontend/requests/api.dart';

Future<void> backendStatus() async {
  final url = createURL("/status");
  Map<String, String> headers = {
    'Authorization': getAPIKey(),
  };
  final response = await http.get(url, headers: headers);
  if (response.statusCode == 200) {
    print(response.body);
  } else {
    throw Exception("failed to get backend status");
  }
}
