import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:frontend/model/poop_location.dart';
import 'package:frontend/requests/poop_locations.dart';
import 'package:latlong2/latlong.dart';
import 'package:url_launcher/url_launcher.dart';
import 'package:frontend/map/marker/marker.dart';
import 'package:frontend/map/poop_location_dialog.dart';
import 'package:flutter_map_location_marker/flutter_map_location_marker.dart';

class MapWidget extends StatefulWidget {
  const MapWidget({super.key});

  @override
  State<MapWidget> createState() => _MapWidgetState();
}

class _MapWidgetState extends State<MapWidget> {
  // poop location stuff
  List<PoopLocation> poopLocations = [];
  // user location stuff
  late final Stream<LocationMarkerPosition?> positionStream;
  StreamSubscription<LocationMarkerPosition?>? positionSubscription;
  late LatLng? currentLocation;
  AlignOnUpdate alignPositionOnUpdate = AlignOnUpdate.once;
  late StreamController<double?> followCurrentLocationStreamController;

  List<Marker> createMarkers() {
    return poopLocations
        .map((poopLocation) => makeMarker(poopLocation, 80, 80, 40, context))
        .toList();
  }

  Future<void> _loadPoopLocations() async {
    final locs = await getAllPoopLocations();
    setState(() {
      poopLocations = locs;
    });
  }

  @override
  void initState() {
    _loadPoopLocations();
    const factory = LocationMarkerDataStreamFactory();
    positionStream = factory.fromGeolocatorPositionStream().asBroadcastStream();
    positionSubscription = positionStream.listen((position) {
      if (position != null) {}
    });

    followCurrentLocationStreamController = StreamController<double?>();
    super.initState();
  }

  @override
  void dispose() {
    positionSubscription?.cancel();
    followCurrentLocationStreamController.close();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        FlutterMap(
          options: MapOptions(
            initialCenter: const LatLng(40.0, -73.94),
            initialZoom: 19,
            onTap: (tapPosition, point) => {
              showAddPoopLocationDialog(
                  context,
                  point,
                  (PoopLocation pl) => setState(() {
                        poopLocations.add(pl);
                      }))
            },
          ),
          children: [
            TileLayer(
              urlTemplate: 'https://tile.openstreetmap.org/{z}/{x}/{y}.png',
              userAgentPackageName: 'com.example.app',
              maxNativeZoom: 19,
            ),
            CurrentLocationLayer(
              alignPositionOnUpdate: alignPositionOnUpdate,
              alignPositionStream: followCurrentLocationStreamController.stream,
            ),
            MarkerLayer(
              markers: createMarkers(),
            ),
            RichAttributionWidget(
              attributions: [
                TextSourceAttribution(
                  'OpenStreetMap contributors',
                  onTap: () => launchUrl(
                      Uri.parse('https://openstreetmap.org/copyright')),
                ),
              ],
            ),
          ],
        ),
      ],
    );
  }
}
