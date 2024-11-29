import 'package:flutter/material.dart';
import 'package:frontend/requests/backend_status.dart';
import 'package:frontend/config/config.dart';
import 'package:frontend/map/map.dart';

void main() async {
  await loadConfig();
  runApp(const App());
}

class App extends StatelessWidget {
  const App({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: "poop.map",
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const Home(title: 'Poop Map'),
    );
  }
}

class Home extends StatefulWidget {
  const Home({super.key, required this.title});

  final String title;

  @override
  State<Home> createState() => _HomeState();
}

class _HomeState extends State<Home> {
  //int _counter = 0;

  //void _incrementCounter() {
  //  print("INCREMENTING THE COUNTER");
  //  backendStatus();
  //  setState(() {
  //    _counter++;
  //  });
  //}

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
      ),
      body: const Stack(
        children: [
          MapWidget(),
        ],
      ),
      //floatingActionButton: FloatingActionButton(
      //  onPressed: _incrementCounter,
      //  tooltip: 'Increment',
      //  child: const Icon(Icons.add),
      //), // This trailing comma makes auto-formatting nicer for build methods.
    );
  }
}
