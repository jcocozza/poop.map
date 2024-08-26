import 'package:flutter/widgets.dart';
import 'package:path/path.dart';
import 'package:poop_map/model/poop_location.dart';
import 'package:sqflite/sqflite.dart';
import 'package:path_provider/path_provider.dart' as path_provider;


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

Future<void> insertPoopLocation(Database db, PoopLocation poopLocation) async {
  int _ = await db.insert(
    'poop_locations',
    poopLocation.toMap(),
    conflictAlgorithm: ConflictAlgorithm.replace,
  );
}

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