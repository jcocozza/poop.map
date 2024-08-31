import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:poop_map/model/poop_location.dart';
import 'package:poop_map/widgets/map/marker/marker_info_dialog.dart';

Marker makeMarker(PoopLocation poopLocation, double width, double height, double iconSize, BuildContext ctx) {
  return Marker(
    point: poopLocation.location(),
    width: width,
    height: height,
    child: GestureDetector(
      onTap: () => showMarkerInfoDialog(ctx, poopLocation),
      child: Icon(Icons.location_pin, color: poopLocation.locationType.color, size: iconSize,),
    ));
}
