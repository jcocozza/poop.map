import 'package:flutter/material.dart';

class SeasonsSelector extends StatefulWidget {
  final Map<String, bool> seasons;
  const SeasonsSelector({super.key, required this.seasons});

  @override
  State<SeasonsSelector> createState() => _SeasonsSelectorState();
}

class _SeasonsSelectorState extends State<SeasonsSelector> {
  final Map<String, bool> seasons;

  @override
  Widget build(BuildContext context) {
    return Row(
      children: [
        Column(
          children: [
            const Text("Summer"),
            Checkbox(
              value: seasons['summer'],
              onChanged: (bool? value) {
                setState(() {
                  seasons['summer'] = value!;
                });
              },
            ),
          ],
        ),
        Column(
          children: [
            const Text("winter"),
            Checkbox(
              value: seasons['winter'],
              onChanged: (bool? value) {
                setState(() {
                  seasons['winter'] = value!;
                });
              },
            ),
          ],
        ),
        Column(
          children: [
            const Text("fall"),
            Checkbox(
              value: seasons['fall'],
              onChanged: (bool? value) {
                setState(() {
                  seasons['fall'] = value!;
                });
              },
            ),
          ],
        ),
        Column(
          children: [
            const Text("spring"),
            Checkbox(
              value: seasons['spring'],
              onChanged: (bool? value) {
                setState(() {
                  seasons['spring'] = value!;
                });
              },
            ),
          ],
        ),
      ],
    );
  }
}
