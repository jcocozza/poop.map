import 'dart:convert';
import 'dart:io';

class Config {
  final String backendUrl;
  final String backendPort;
  final String frontendUrl;
  final String frontendPort;

  Config({
    required this.backendUrl,
    required this.backendPort,
    required this.frontendUrl,
    required this.frontendPort,
  });

  // Factory constructor to create an instance from JSON
  factory Config.fromJson(Map<String, dynamic> json) {
    return Config(
      backendUrl: json['backend_url'],
      backendPort: json['backend_port'],
      frontendUrl: json['frontend_url'],
      frontendPort: json['frontend_port'],
    );
  }
}

/// read the config file
/// does catch the error because I want things to crash if the config doesn't read in
Future<Config> readConfig(String path) async {
  final file = File(path);
  final fileContent = await file.readAsString();
  final jsonData = jsonDecode(fileContent);
  final config = Config.fromJson(jsonData);
  return config;
}

Future<Config> parseConfig(String cfgStr) async {
  final jsonData = jsonDecode(cfgStr);
  final config = Config.fromJson(jsonData);
  return config;
}

