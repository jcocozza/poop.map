import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:flutter_map_location_marker/flutter_map_location_marker.dart';
import 'package:latlong2/latlong.dart';
import 'package:poop_map/requests/requests.dart';
import 'package:poop_map/model/poop_location.dart';
import 'package:poop_map/utils/read_config.dart';
import 'unpack_polyline.dart';
import 'widgets/star_rating.dart';

const String appConfigAsset = "../config.json";

Future<Config> loadConfig() async {
  try {
    final String configContent = await rootBundle.loadString(appConfigAsset);
    return parseConfig(configContent);
  } catch (e) {
    print('Error loading config: $e');
    rethrow;
  }
}

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  final config = await loadConfig();
  runApp(MyApp(config: config,));
}
class MyApp extends StatelessWidget {
  final Config config;
  const MyApp({super.key, required this.config});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'poop.map',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: MyHomePage(title: 'click to add a poop location', config: config,),
    );
  }
}

class MyHomePage extends StatefulWidget {
  final Config config;
  const MyHomePage({super.key, required this.title, required this.config});
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
  StreamSubscription<LocationMarkerPosition?>? _positionSubscription;


  @override
  void initState() {
    super.initState();
    _loadPoopLocations();
    //_getCurrentLocation();
    const factory = LocationMarkerDataStreamFactory();
    _positionStream = factory.fromGeolocatorPositionStream().asBroadcastStream();
    _positionSubscription = _positionStream.listen((position) {
    if (position != null) {
      setState(() {
        _currentLocation = LatLng(position.latitude, position.longitude);
      });
    }
    });
  }
  @override
  void dispose() {
    _positionSubscription?.cancel();
    super.dispose();
  }

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

  void _showAddPoopLocationDialog(LatLng location) async {
    final _nameController = TextEditingController();
    int _rating = 1;
    LocationType _selectedLocationType = LocationType.regular;

    return showDialog<void>(
      context: context,
      barrierDismissible: false,
      builder: (BuildContext context) {
        return StatefulBuilder(
          builder: (BuildContext context, StateSetter setStateLoc) {
            return AlertDialog(
              title: const Text('Add Poop Location'),
              content: SingleChildScrollView(
                child: ListBody(
                  children: <Widget>[
                    TextField(
                      controller: _nameController,
                      decoration: const InputDecoration(labelText: 'Name'),
                    ),
                    const Text('Rating:'),
                    StarRating(
                    rating: _rating,
                    readOnly: false,
                    onRatingChanged: (int rating) {
                      setStateLoc(() {
                        _rating = rating;
                      });
                    },
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
                          setStateLoc(() {
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
                    final locationType = _selectedLocationType;
                    if (name.isNotEmpty) {
                        PoopLocation pl = createPoopLocation(
                          location.latitude,
                          location.longitude,
                          _rating,
                          locationType,
                          name,
                        );
                      await insertPoopLocation(widget.config, pl);
                      Navigator.of(context).pop();
                      setState(() {
                        _poopLocations.add(pl);
                      });
                    }
                  },
                ),
              ],
            );
          }
        );
      },
    );
  }

  void _showMarkerInfo(PoopLocation poopLocation) {
    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: Column(
            children: [
              StarRating(rating: poopLocation.rating),
              Text('Poop Location: ${poopLocation.name} - ${poopLocation.locationType.displayName}'),
            ],
          ),
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

void _floatingButtonCallback() async {
  if (_currentLocation == null) {
    print("Error: Current location is not available");
    return;
  }

  setState(() {
    _isLoading = true;
  });

  ClosestPoopLocation? closestPoopLocation = await getClosestPoopLocation(widget.config,
      _currentLocation!.latitude, _currentLocation!.longitude);

  if (closestPoopLocation != null) {
    setState(() {
      _polylines.add(
          Polyline(points: decodePolyline(closestPoopLocation.route).unpackPolyline()));
      _isLoading = false;
    });
  }
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
