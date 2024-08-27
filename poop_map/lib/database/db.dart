import 'dart:convert';

import 'package:flutter/widgets.dart';
import 'package:http/http.dart' as http;
import 'package:path/path.dart';
import 'package:poop_map/model/poop_location.dart';
import 'package:sqflite/sqflite.dart';
import 'package:path_provider/path_provider.dart' as path_provider;


/*
Future<Database> createDatabase() async {
  WidgetsFlutterBinding.ensureInitialized();
  final documentsDirectory = await path_provider.getApplicationDocumentsDirectory();
  final path = join(documentsDirectory.path, 'poop_locations_database.db');
  print("db is located at path: $path");
  final database = openDatabase(
    // Set the path to the database. Note: Using the `join` function from the
    // `path` package is best practice to ensure the path is correctly
    // constructed for each platform.
    //join(await getDatabasesPath(), 'poop_locations_database.db'),
    //"/Users/josephcocozza/Repositories/poop.map/test/poop_locations_database.db",
    path,
    onCreate: (db, version) {
      // Run the CREATE TABLE statement on the database.
      return db.execute(
        'CREATE TABLE poop_locations(uuid TEXT PRIMARY KEY NOT NULL, latitude REAL NOT NULL, longitude REAL NOT NULL, rating INTEGER, first_created TEXT NOT NULL, location_type TEXT, name TEXT)',
      );
    },
    // Set the version. This executes the onCreate function and provides a
    // path to perform database upgrades and downgrades.
    version: 1,
  );

  return database;
}
*/

/*
Future<void> insertPoopLocation(Database db, PoopLocation poopLocation) async {
  int _ = await db.insert(
    'poop_locations',
    poopLocation.toMap(),
    conflictAlgorithm: ConflictAlgorithm.replace,
  );
}
*/

Future<void> insertPoopLocation(PoopLocation poopLocation) async {
  const baseUrl = 'http://localhost:8080/api/create';
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
    // Check the response status
    if (response.statusCode == 201) {
      // Request was successful
      print('Response data: ${response.body}');
    } else {
      // Request failed
      print('Failed to send request: ${response.statusCode}');
    }
  } catch (e) {
    print('Error: $e');
  }
}

/*
Future<List<PoopLocation>> getAllPoopLocations(Database db) async {
  final List<Map<String, Object?>> poopLocMaps = await db.query('poop_locations');

  return [
    for (final {
          'uuid': uuid as String,
          'latitude': latitude as double,
          'longitude': longitude as double,
          'rating': rating as int,
          'first_created': firstCreated as String,
          'location_type': locationType as String,
          'name': name as String,
        } in poopLocMaps)
        PoopLocation(uuid: uuid, latitude: latitude, longitude: longitude, rating: rating, firstCreated: firstCreated, locationType: LocationTypeExtension.fromString(locationType), name: name)
  ];
}
*/

Future<List<PoopLocation>> getAllPoopLocations() async {
  final url = Uri.parse("http://localhost:8080/api/list_all");
  try {
    final response = await http.get(url);
    if (response.statusCode == 200) {
      // Request successful, parse the JSON response
      final decodedResponse = json.decode(response.body);
      print("THIS IS THE DECODED RESPONSE");
      print(decodedResponse);
      // Convert each item in the response to a PoopLocation
      return decodedResponse.map<PoopLocation>((item) {
        return PoopLocation(
          uuid: item['uuid'] as String,
          latitude: (item['latitude'] as num).toDouble(),
          longitude: (item['longitude'] as num).toDouble(),
          rating: item['rating'] as int,
          firstCreated: item['first_created'] as String,
          locationType: LocationTypeExtension.fromString(item['location_type'] as String),
          name: item['name'] as String,
        );
      }).toList();
    } else {
      // Request failed
      print('Request failed with status: ${response.statusCode}');
      return [];
    }
  } catch (e) {
    // Handle any errors that occurred during the request
    print('Error: $e');
    return [];
  }
}