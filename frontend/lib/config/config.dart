import 'package:flutter_dotenv/flutter_dotenv.dart';

Future<void> loadConfig() async {
  await dotenv.load(fileName: '.env');
}
