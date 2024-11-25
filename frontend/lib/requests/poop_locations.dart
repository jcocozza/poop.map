import 'dart:convert';
import 'package:frontend/model/poop_location.dart';
import 'package:http/http.dart' as http;
import 'package:frontend/requests/api.dart';

Future<List<PoopLocation>> getAllPoopLocations() async {
  final url = createURL("/poop-location");
  Map<String, String> headers = {
    'Authorization': getAPIKey(),
  };
  final response = await http.get(url, headers: headers);
  if (response.statusCode == 200) {
    final decodedRespose = json.decode(response.body);
    List<PoopLocation> lst =
        decodedRespose['data'].map((js) => PoopLocation.fromJson(js)).toList();
    return lst;
  } else {
    return [];
  }
}

Future<void> createPoopLocation(PoopLocation poopLocation) async {
  final url = createURL("/poop-location");
  Map<String, String> headers = {
    'Authorization': getAPIKey(),
  };

  final json = jsonEncode(poopLocation.toJson());

  final response = await http.put(url, headers: headers, body: json);
  if (response.statusCode == 201) {
    return;
  } else {
    print(response.body);
    return;
  }
}

Future<void> upvote(String poopLocationUUID) async {
  final url = createURL('/poop-location/$poopLocationUUID/upvote');
  Map<String, String> headers = {
    'Authorization': getAPIKey(),
  };
  final response = await http.patch(url, headers: headers);
  if (response.statusCode == 200) {
    return;
  } else {
    return;
  }
}

Future<void> downvote(String poopLocationUUID) async {
  final url = createURL('/poop-location/$poopLocationUUID/downvote');
  Map<String, String> headers = {
    'Authorization': getAPIKey(),
  };
  final response = await http.patch(url, headers: headers);
  if (response.statusCode == 200) {
    return;
  } else {
    return;
  }
}
