import 'package:logging/logging.dart';
import 'package:flutter/foundation.dart'; // for kReleaseMode

final Logger logger = Logger('poop_map_logger');

/// Instantial the logger listening function
void setUpLogging() {
  Logger.root.level = Level.ALL;
  // in development mode, we can just print logs to the screen
  if (kDebugMode) {
    Logger.root.onRecord.listen((LogRecord record) {
      // ignore: avoid_print
      print('[${record.level.name}] ${record.time}: ${record.message}');
    });
  }
  // for production do nothing
}
