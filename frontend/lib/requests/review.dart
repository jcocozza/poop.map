import 'dart:convert';
import 'package:frontend/model/review.dart';
import 'package:http/http.dart' as http;
import 'package:frontend/requests/api.dart';

Future<List<Review>> getAllReviewsByPoopLocation(String poopLocationUUID) async {
  final url = createURL("/poop-location/$poopLocationUUID/review");
  Map<String, String> headers = {
    'Authorization': getAPIKey(),
  };
  final response = await http.get(url, headers: headers);
  if (response.statusCode == 200) {
    print(response.body);
    final decodedRespose = json.decode(response.body);
    List<Review> lst =
        decodedRespose['data'].map((js) => Review.fromJson(js)).toList();
    return lst;
  } else {
    print(response.body);
    return [];
  }
}

Future<void> createReview(String poopLocationUUID, Review review) async {
  final url = createURL('/poop-location/$poopLocationUUID/review');
  Map<String, String> headers = {
    'Authorization': getAPIKey(),
  };

  final json = jsonEncode(review.toJson());

  final response = await http.put(url, headers: headers, body: json);
  if (response.statusCode == 201) {
    return;
  } else {
    print(response.body);
    return;
  }
}

Future<void> upvote(String reviewUUID) async {
  final url = createURL('/review/$reviewUUID/upvote');
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

Future<void> downvote(String reviewUUID) async {
  final url = createURL('/review/$reviewUUID/downvote');
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
