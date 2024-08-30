import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:poop_map/model/poop_location.dart';
import 'package:poop_map/utils/read_config.dart';

class ClosestPoopLocation {
  PoopLocation poopLocation;
  String route;

  ClosestPoopLocation(this.poopLocation, this.route);

  Map<String, Object?> toMap() {
    return {
      'poop_location': poopLocation,
      'route': route
    };
  }
}

/// Call the api to get the closest poop location and the route to it
Future<ClosestPoopLocation?> getClosestPoopLocation(Config cfg, double currLat, double currLong) async {
  final baseUrl = '${cfg.backendUrl}/api/closest';
  String params = "?latitude=${currLat}&longitude=${currLong}";
  final url = Uri.parse("$baseUrl$params");
  try {
    final response = await http.get(url);
    if (response.statusCode == 200) {
      final decodedResponse = json.decode(response.body);
      print(decodedResponse);
      ClosestPoopLocation cp = ClosestPoopLocation(PoopLocation.fromJson(decodedResponse['poop_location']), decodedResponse['route'] as String);
      print(cp);
      return cp;
    } else {
      print('Request failed with status: ${response.statusCode}');
      return null;
    }
  } catch (e) {
    print('Error: $e');
    return null;
  }
}

/// Call the api to create a new poop location
Future<void> insertPoopLocation(Config cfg, PoopLocation poopLocation) async {
  final baseUrl = '${cfg.backendUrl}/api/create';
  final url = Uri.parse(baseUrl);
  final jsonObj = poopLocation.toMap();
  final jsonBody = jsonEncode(jsonObj);
  try {
    final response = await http.post(
      url,
      headers: {
        'Content-Type': 'application/json', // Set the content type to JSON
      },
      body: jsonBody,
    );
    if (response.statusCode == 201) {
      print('Response data: ${response.body}');
    } else {
      print('Failed to send request: ${response.statusCode}');
    }
  } catch (e) {
    print('Error: $e');
  }
}

/// Call the api to list all poop locations
Future<List<PoopLocation>> getAllPoopLocations(Config cfg) async {
  final url = Uri.parse("${cfg.backendUrl}/api/list_all");
  try {
    final response = await http.get(url);
    if (response.statusCode == 200) {
      final decodedResponse = json.decode(response.body);
      return decodedResponse.map<PoopLocation>((item) {
        return PoopLocation(
          uuid: item['uuid'] as String,
          latitude: (item['latitude'] as num).toDouble(),
          longitude: (item['longitude'] as num).toDouble(),
          rating: item['rating'] as int,
          firstCreated: item['first_created'] as String,
          locationType: LocationTypeExtension.fromString(item['location_type'] as String),
          name: item['name'] as String,
          notes: item['notes'] as String,
        );
      }).toList();
    } else {
      print('Request failed with status: ${response.statusCode}');
      return [];
    }
  } catch (e) {
    print('Error: $e');
    return [];
  }
}
