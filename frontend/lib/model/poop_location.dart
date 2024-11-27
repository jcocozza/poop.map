import 'package:latlong2/latlong.dart';
import 'package:flutter/material.dart';

class PoopLocation {
  String uuid;
  String name;
  double latitude;
  double longitude;
  String firstCreated;
  String lastModified;
  String locationType;
  bool seasonal;
  List<String> seasons;
  bool accessible;
  int upvotes;
  int downvotes;

  PoopLocation({
    required this.uuid,
    required this.name,
    required this.latitude,
    required this.longitude,
    required this.firstCreated,
    required this.lastModified,
    required this.locationType,
    required this.seasonal,
    required this.seasons,
    required this.accessible,
    required this.upvotes,
    required this.downvotes,
  });

  // return the lat/lng coord
  LatLng location() {
    return LatLng(latitude, longitude);
  }

  MaterialColor color() {
    switch (locationType) {
      case "regular":
        return Colors.red;
      case "porta potty":
        return Colors.green;
      case "outhouse":
        return Colors.brown;
      case "other":
      default:
        return Colors.blue;
    }
  }

  List<Widget> seasonsIcons() {
    Map<String, Widget> seasonsMap = {
      "summer": const Column(children: [Icon(Icons.sunny), Text("summer")]),
      "spring": const Column(children: [Icon(Icons.eco), Text("spring")]),
      "winter": const Column(children: [Icon(Icons.ac_unit), Text("winter")]),
      "fall": const Column(children: [Icon(Icons.nature), Text("fall")]),
    };
    return seasons
        .map((season) => seasonsMap[season] ?? const Icon(Icons.help_outline))
        .toList();
  }

  // A method to convert the PoopLocation instance to JSON
  Map<String, dynamic> toJson() {
    return {
      'uuid': uuid,
      'name': name,
      'latitude': latitude,
      'longitude': longitude,
      //'first_created': firstCreated,
      //'last_modified': lastModified,
      'location_type': locationType,
      'seasonal': seasonal,
      'seasons': seasons,
      'accessible': accessible,
      'upvotes': upvotes,
      'downvotes': downvotes,
    };
  }

  factory PoopLocation.fromJson(Map<String, dynamic> json) {
    return PoopLocation(
      uuid: json['uuid'],
      name: json['name'],
      latitude: json['latitude'],
      longitude: json['longitude'],
      firstCreated: json['first_created'],
      lastModified: json['last_modified'],
      locationType: json['location_type'],
      seasonal: json['seasonal'],
      seasons: json['seasons'],
      accessible: json['accessible'],
      upvotes: json['upvotes'],
      downvotes: json['downvotes'],
    );
  }
}
