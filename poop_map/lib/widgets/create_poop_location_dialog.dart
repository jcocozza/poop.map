import 'package:flutter/material.dart';
import 'package:latlong2/latlong.dart';
import 'package:poop_map/model/poop_location.dart';
import 'package:poop_map/requests/requests.dart';
import 'package:poop_map/utils/read_config.dart';
import 'package:poop_map/widgets/star_rating.dart';

class PoopLocationDialog extends StatefulWidget {
  final Config cfg;
  final LatLng location;
  final void Function(PoopLocation) onPoopLocationAdd;

  const PoopLocationDialog({
    super.key,
    required this.cfg,
    required this.location,
    required this.onPoopLocationAdd
  });

  @override
  State<StatefulWidget> createState() => _PoopLocationDialogState();
}

class _PoopLocationDialogState extends State<PoopLocationDialog> {
  final _nameController = TextEditingController();
  int _rating = 1;
  LocationType _selectedLocationType = LocationType.regular;

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text("Add Poop Location"),
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
                setState(() {
                  _rating = rating;
                });
              },
            ),
            DropdownButton(
              value: _selectedLocationType,
              items: LocationType.values.map((LocationType locationType) {
                return DropdownMenuItem(
                  value: locationType,
                  child: Text(locationType.displayName)
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
          onPressed: () { Navigator.of(context).pop(); },
          child: const Text("cancel")
        ),
        ElevatedButton(
          child: const Text("Add"),
          onPressed: () async {
            final name = _nameController.text;
            final locationType = _selectedLocationType;
            if (name.isNotEmpty) {
              PoopLocation pl = createPoopLocation(
                widget.location.latitude,
                widget.location.longitude,
                _rating,
                locationType,
                name
              );
              await insertPoopLocation(widget.cfg, pl);
              Navigator.of(context).pop();
              widget.onPoopLocationAdd(pl);
            }
          },
        )
      ],
    );
  }
}

/// Show the dialog for adding a poop location
/// Show the dialog for adding a poop location
void showAddPoopLocationDialog(BuildContext ctx, Config cfg, LatLng location, Function(PoopLocation) onPoopLocationAdd) async {
  return showDialog(
    context: ctx,
    barrierDismissible: false,
    builder: (BuildContext ctx) {
      return PoopLocationDialog(
        cfg: cfg,
        location: location,
        onPoopLocationAdd: onPoopLocationAdd,
      );
    },
  );
}