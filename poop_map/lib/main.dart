import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:poop_map/utils/read_config.dart';
import 'package:poop_map/widgets/map/map.dart';

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
  runApp(App(config: config,));
}
class App extends StatelessWidget {
  final Config config;
  const App({super.key, required this.config});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'poop.map',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: Home (title: 'click to add a poop location', config: config,),
    );
  }
}

class Home extends StatefulWidget {
  final Config config;
  const Home({super.key, required this.title, required this.config});
  final String title;

  @override
  State<Home> createState() => _HomeState();
}

class _HomeState extends State<Home> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          backgroundColor: Theme.of(context).colorScheme.inversePrimary,
          title: Text(widget.title),
        ),
        body: Stack(
          children: [
            PoopMap(config: widget.config),
          ],
        ));
  }
}
