import 'package:flutter/material.dart';
import 'package:poop_map/model/poop_location.dart';
import 'package:poop_map/widgets/star_rating.dart';

class MarkerInfoDialog extends StatelessWidget {
  final PoopLocation poopLocation;

  const MarkerInfoDialog({
    super.key,
    required this.poopLocation
  });

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Column(
        children: [
          StarRating(rating: poopLocation.rating),
          Text('Poop Location: ${poopLocation.name} - ${poopLocation.locationType.displayName}'),
        ],
      ),
      content: Text( // TODO: clean this up and make it better
          '${poopLocation.name} is rated ${poopLocation.rating} and has been around since ${poopLocation.firstCreated}\n\n${poopLocation.notes}'),
      actions: <Widget>[
        TextButton(
          child: const Text('Close'),
          onPressed: () {
            Navigator.of(context).pop();
          },
        ),
      ],
    );
  }
}

void showMarkerInfoDialog(BuildContext ctx, PoopLocation poopLocation) async {
  return showDialog(
    context: ctx,
    builder: (BuildContext ctx) {
      return MarkerInfoDialog(poopLocation: poopLocation);
    }
  );
}
