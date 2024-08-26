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

  const PoopLocation({
    required this.uuid,
    required this.latitude,
    required this.longitude,
    required this.rating,
    required this.firstCreated,
    required this.locationType,
    required this.name,
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
    };
  }
}

PoopLocation createPoopLocation(double latitude, double longitude, int rating, LocationType locationType, String name) {
  const uuid = Uuid();

  DateTime now = DateTime.now();
  String currentDate = DateFormat.yMMMMd('en_US').format(now);
  return PoopLocation(uuid: uuid.v4(), latitude: latitude, longitude: longitude, rating: rating, firstCreated: currentDate, locationType: locationType, name: name);
}