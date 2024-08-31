import 'package:flutter/material.dart';
import 'package:latlong2/latlong.dart';
import 'package:uuid/uuid.dart';
import 'package:intl/intl.dart';

enum LocationType {
  regular,
  portaPotty,
  outhouse,
  other,
}

extension LocationTypeExtension on LocationType {
  String get displayName {
    switch (this) {
      case LocationType.regular:
        return 'regular';
      case LocationType.portaPotty:
        return 'porta potty';
      case LocationType.outhouse:
        return "outhouse";
      case LocationType.other:
      default:
        return 'other';
    }
  }
  MaterialColor get color {
    switch (this) {
      case LocationType.regular:
        return Colors.red;
      case LocationType.portaPotty:
        return Colors.green;
      case LocationType.outhouse:
        return Colors.brown;
      case LocationType.other:
      default:
        return Colors.blue;
    }
  }
  static LocationType fromString(String value) {
    return LocationType.values.firstWhere(
      (type) => type.displayName == value.toLowerCase(),
      orElse: () => LocationType.other,
    );
  }
}

class PoopLocation {
  final String uuid;
  final double latitude;
  final double longitude;
  final int rating; // rating of 1-5
  final String firstCreated; // date of creation
  final LocationType locationType;
  final String name;
  final String notes;

  const PoopLocation({
    required this.uuid,
    required this.latitude,
    required this.longitude,
    required this.rating,
    required this.firstCreated,
    required this.locationType,
    required this.name,
    required this.notes,
  });

  /// return the lat/lng coord
  LatLng location() {
    return LatLng(latitude, longitude);
  }

  Map<String, Object?> toMap() {
    return {
      'uuid': uuid,
      'latitude': latitude,
      'longitude': longitude,
      'rating': rating,
      'first_created': firstCreated,
      'location_type': locationType.displayName,
      'name': name,
      'notes': notes,
    };
  }

  factory PoopLocation.fromJson(Map<String, dynamic> json) {
    return PoopLocation(
      uuid: json["uuid"] as String,
      latitude: json['latitude'] as double,
      longitude: json['longitude'] as double,
      rating: json['rating'] as int,
      locationType: LocationTypeExtension.fromString(json['location_type'] as String),
      name: json['name'] as String,
      notes: json['notes'] as String,
      firstCreated: json['first_created'] as String,
    );
  }
}

PoopLocation createPoopLocation(double latitude, double longitude, int rating, LocationType locationType, String name, String notes) {
  const uuid = Uuid();

  DateTime now = DateTime.now();
  String currentDate = DateFormat.yMMMMd('en_US').format(now);
  return PoopLocation(uuid: uuid.v4(), latitude: latitude, longitude: longitude, rating: rating, firstCreated: currentDate, locationType: locationType, name: name, notes: notes);
}
