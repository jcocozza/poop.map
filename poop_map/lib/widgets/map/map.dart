import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:flutter_map_location_marker/flutter_map_location_marker.dart';
import 'package:latlong2/latlong.dart';
import 'package:poop_map/model/poop_location.dart';
import 'package:poop_map/requests/requests.dart';
import 'package:poop_map/utils/unpack_polyline.dart';
import 'package:poop_map/utils/read_config.dart';
import 'package:poop_map/widgets/create_poop_location_dialog.dart';
import 'package:poop_map/widgets/loading.dart';
import 'package:poop_map/widgets/map/marker/marker.dart';
import 'package:poop_map/widgets/map/legend.dart';
import 'package:url_launcher/url_launcher.dart';

class PoopMap extends StatefulWidget {
  final Config config;

  const PoopMap({
    super.key,
    required this.config,
  });

  @override
  State<PoopMap> createState() => _MapState();
}

class _MapState extends State<PoopMap> {
  late final Stream<LocationMarkerPosition?> _positionStream;
  StreamSubscription<LocationMarkerPosition?>? _positionSubscription;
  List<PoopLocation> _poopLocations = [];
  List<Polyline<Object>> _polylines = [];
  late LatLng? _currentLocation;
  bool _findingRoute = false;

  Future<void> _loadPoopLocations() async {
    try {
      final locations = await getAllPoopLocations(widget.config);
      setState(() {
        _poopLocations = locations;
        print("poop locations set!");
        print(_poopLocations);
      });
    } catch (e) {
      print('Error loading poop locations: $e');
    }
  }

  List<Marker> _createMarkersFromPoopLocations() {
    return _poopLocations
        .map((poopLocation) => makeMarker(poopLocation, 80, 80, 40, context))
        .toList();
  }

  Future<void> _navigateFromCurrentLocation() async {
    // don't do anything if we don't have a user location
    if (_currentLocation == null) {
      return;
    }
    setState(() {
      _findingRoute = true;
    });
    ClosestPoopLocation? closestPoopLocation = await getClosestPoopLocation(
        widget.config, _currentLocation!.latitude, _currentLocation!.longitude);
    if (closestPoopLocation != null) {
      setState(() {
        _polylines.add(Polyline(
            points:
                decodePolyline(closestPoopLocation.route).unpackPolyline()));
      });
    }
    // stop loading once everything returns, no matter what
    setState(() {
      _findingRoute = false;
    });
  }

  @override
  void initState() {
    _loadPoopLocations();
    const factory = LocationMarkerDataStreamFactory();
    _positionStream =
        factory.fromGeolocatorPositionStream().asBroadcastStream();
    _positionSubscription = _positionStream.listen((position) {
      if (position != null) {
        setState(() {
          _currentLocation = LatLng(position.latitude, position.longitude);
        });
      }
    });

    super.initState();
  }

  @override
  void dispose() {
    _positionSubscription?.cancel();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Stack(children: [
      FlutterMap(
        options: MapOptions(
          initialCenter: const LatLng(40.730610, -73.935242),
          initialZoom: 19,
          onTap: (tapPosition, point) => {
            showAddPoopLocationDialog(
                context,
                widget.config,
                point,
                (PoopLocation pl) => {
                      setState(() {
                        _poopLocations.add(pl);
                      })
                    })
          },
        ),
        children: [
          // in the future these can become their own widgets if they grow
          TileLayer(
            // Display map tiles from any source
            urlTemplate:
                'https://tile.openstreetmap.org/{z}/{x}/{y}.png', // OSMF's Tile Server
            userAgentPackageName: 'com.example.app',
            maxNativeZoom:
                19, // Scale tiles when the server doesn't support higher zoom levels
          ),
          CurrentLocationLayer(alignPositionOnUpdate: AlignOnUpdate.once),
          MarkerLayer(markers: _createMarkersFromPoopLocations()),
          PolylineLayer(polylines: _polylines),
          RichAttributionWidget(attributions: [
            TextSourceAttribution(
              'OpenStreetMap contributors',
              onTap: () => launchUrl(Uri.parse(
                  'https://openstreetmap.org/copyright')), // (external)
            )
          ]),
        ],
      ),
      if (_findingRoute)
        const Loading(
            loadingMessage:
                "determining route to closest poop location. this can take a minute..."),
      Align(
        alignment: Alignment.bottomLeft,
        child: FloatingActionButton(
          onPressed: _navigateFromCurrentLocation,
          child: const Icon(Icons.navigation_rounded),
        ),
      ),
      const Align(
        alignment: Alignment.topRight,
        child: Padding(
          padding: EdgeInsets.all(16.0),
          child: MapLegend(),
        ),
      ),
    ]);
  }
}
