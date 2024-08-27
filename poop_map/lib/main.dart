import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:flutter_map_location_marker/flutter_map_location_marker.dart';
import 'package:latlong2/latlong.dart';
import 'package:poop_map/requests/requests.dart';
import 'package:poop_map/model/poop_location.dart';
import 'package:geolocator/geolocator.dart';
import 'unpack_polyline.dart';

void main() {
  runApp(const MyApp());
}
class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'poop.map',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const MyHomePage(title: 'click to add a poop location'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key, required this.title});
  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  late final Stream<LocationMarkerPosition?> _positionStream;
  LatLng? _currentLocation;
  List<PoopLocation> _poopLocations = [];
  List<Polyline<Object>> _polylines = [];
  bool _isLoading = false;

  @override
  void initState() {
    super.initState();
    _loadPoopLocations();
    //_getCurrentLocation();
    const factory = LocationMarkerDataStreamFactory();
    _positionStream =
        factory.fromGeolocatorPositionStream().asBroadcastStream();
  }

  Future<void> _loadPoopLocations() async {
    try {
      final locations = await getAllPoopLocations();
      setState(() {
        _poopLocations = locations;
        print("poop locations set!");
        print(_poopLocations);
      });
    } catch (e) {
      print('Error loading poop locations: $e');
    }
  }

  void _showAddPoopLocationDialog(LatLng location) async {
    final _nameController = TextEditingController();
    final _ratingController = TextEditingController();
    LocationType _selectedLocationType = LocationType.regular;

    return showDialog<void>(
      context: context,
      barrierDismissible: false,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text('Add Poop Location'),
          content: SingleChildScrollView(
            child: ListBody(
              children: <Widget>[
                TextField(
                  controller: _nameController,
                  decoration: const InputDecoration(labelText: 'Name'),
                ),
                TextField(
                  controller: _ratingController,
                  decoration: const InputDecoration(labelText: 'Rating'),
                  keyboardType: TextInputType.number,
                ),
                DropdownButton<LocationType>(
                    value: _selectedLocationType,
                    items: LocationType.values.map((LocationType locationType) {
                      return DropdownMenuItem<LocationType>(
                        value: locationType,
                        child: Text(locationType.displayName),
                      );
                    }).toList(),
                    onChanged: (LocationType? newValue) {
                      setState(() {
                        _selectedLocationType = newValue!;
                      });
                    })
              ],
            ),
          ),
          actions: <Widget>[
            TextButton(
              child: const Text('Cancel'),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
            ElevatedButton(
              child: const Text('Add'),
              onPressed: () async {
                final name = _nameController.text;
                final rating = int.tryParse(_ratingController.text) ?? 0;
                final locationType = _selectedLocationType;

                if (name.isNotEmpty) {
                    PoopLocation pl = createPoopLocation(
                      location.latitude,
                      location.longitude,
                      rating,
                      locationType,
                      name,
                    );
                  await insertPoopLocation(pl);
                  setState(() {
                    _poopLocations.add(pl);
                  });
                  Navigator.of(context).pop();
                }
              },
            ),
          ],
        );
      },
    );
  }

  void _showMarkerInfo(PoopLocation poopLocation) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: Text(
              'Poop Location: ${poopLocation.name} - ${poopLocation.locationType.displayName}'),
          content: Text(
              '${poopLocation.name} is rated ${poopLocation.rating} and has been around since ${poopLocation.firstCreated}'),
          actions: <Widget>[
            TextButton(
              child: const Text('Close'),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
          ],
        );
      },
    );
  }

  Future<void> _getCurrentLocation() async {
    try {
      Position position = await Geolocator.getCurrentPosition(locationSettings: const LocationSettings(accuracy: LocationAccuracy.high,),);
      setState(() {
        _currentLocation = LatLng(position.latitude, position.longitude);
      });
      print("Current Location: ${_currentLocation!.latitude}, ${_currentLocation!.longitude}");
    } catch (e) {
      print("Error getting location: $e");
    }
  }

  void _floatingButtonCallback() async {
    setState(() {
      _isLoading = true;
    });
    //await _getCurrentLocation();
    //_currentLocation = const LatLng(40.730610, -73.935242);
    LocationMarkerPosition? _foo = await _positionStream.first;
    _currentLocation = LatLng(_foo!.latitude, _foo.longitude);

    ClosestPoopLocation? closestPoopLocation = await getClosestPoopLocation(_currentLocation!.latitude, _currentLocation!.longitude);
    setState(() {
      _polylines.add(Polyline(points: decodePolyline(closestPoopLocation!.route).unpackPolyline()));
      _isLoading = false;
    });
  }
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          backgroundColor: Theme.of(context).colorScheme.inversePrimary,
          title: Text(widget.title),
        ),
        body: Stack(
          children: [
            FlutterMap(
              options: MapOptions(
                initialCenter: const LatLng(
                    40.730610, -73.935242), // Center the map over NYC
                initialZoom: 19,
                onTap: (tapPosition, point) => setState(() {
                  _showAddPoopLocationDialog(point);
                  //_markerPositions.add(point);
                }),
              ),
              children: [
                TileLayer(
                  // Display map tiles from any source
                  urlTemplate: 'https://tile.openstreetmap.org/{z}/{x}/{y}.png', // OSMF's Tile Server
                  userAgentPackageName: 'com.example.app',
                  maxNativeZoom: 19, // Scale tiles when the server doesn't support higher zoom levels
                ),
                CurrentLocationLayer(alignPositionOnUpdate: AlignOnUpdate.always),
                MarkerLayer(
                  markers: _poopLocations
                      .map((poopLocation) => Marker(
                            point: poopLocation.location(),
                            width: 80,
                            height: 80,
                            child: GestureDetector(
                              onTap: () => {_showMarkerInfo(poopLocation)},
                              child: const Icon(Icons.location_pin,
                                  color: Colors.red, size: 40),
                            ),
                          ))
                      .toList(),
                ),
                PolylineLayer(polylines: _polylines),
                const RichAttributionWidget(
                  // Include a stylish prebuilt attribution widget that meets all requirments
                  attributions: [
                    TextSourceAttribution(
                      'OpenStreetMap contributors',
                      //onTap: () => launchUrl(Uri.parse('https://openstreetmap.org/copyright')), // (external)
                    ),
                  ],
                ),
              ],
            ),
            if (_isLoading)
              const Center(
                child: Column(
                  children: [
                    CircularProgressIndicator(),
                    Text("determining route to closest poop location. this can take a minute...")
                  ]
                ), // Display loading wheel
              ),
            FloatingActionButton(onPressed: _floatingButtonCallback,
              child: const Icon(Icons.navigation_rounded),
            ),
          ],
        ));
  }
}
