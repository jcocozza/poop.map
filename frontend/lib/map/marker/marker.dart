import 'package:flutter/material.dart';
import 'package:flutter_map/flutter_map.dart';
import 'package:frontend/model/poop_location.dart';
import 'package:frontend/requests/poop_locations.dart';
import 'package:frontend/reviews/reviews.dart';

class MarkerInfoDialog extends StatefulWidget {
  final PoopLocation poopLocation;
  const MarkerInfoDialog({
    super.key,
    required this.poopLocation,
  });

  @override
  State<MarkerInfoDialog> createState() => _MarkerInfoDialogState();
}

class _MarkerInfoDialogState extends State<MarkerInfoDialog> {
  void upvoter() async {
    await upvote(widget.poopLocation.uuid);
    setState(() {
      widget.poopLocation.upvotes += 1;
    });
  }

  void downvoter() async {
    await downvote(widget.poopLocation.uuid);
    setState(() {
      widget.poopLocation.downvotes += 1;
    });
  }

  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: Column(
        children: [
          Text(widget.poopLocation.name),
        ],
      ),
      content: Column(
        children: [
          if (widget.poopLocation.seasonal)
            const Text(
                "This location is seasonal, it might only be open during the following seasons:"),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: widget.poopLocation.seasonsIcons(),
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Text(
                'Votes: ${widget.poopLocation.downvotes}',
              ),
              IconButton(
                icon: const Icon(Icons.arrow_downward),
                onPressed: () => downvoter(),
              ),
              Text(
                'Votes: ${widget.poopLocation.upvotes}',
              ),
              IconButton(
                icon: const Icon(Icons.arrow_upward),
                onPressed: () => upvoter(),
              ),
            ],
          ),
          ReviewList(poopLocationUUID: widget.poopLocation.uuid,),
        ],
      ),
      actions: <Widget>[
        TextButton(
          child: const Text('close'),
          onPressed: () {
            Navigator.of(context).pop();
          },
        )
      ],
    );
  }
}

void showMarkerInfoDialog(
    BuildContext context, PoopLocation poopLocation) async {
  return showDialog(
    context: context,
    builder: (BuildContext context) {
      return MarkerInfoDialog(
        poopLocation: poopLocation,
      );
    },
  );
}

Marker makeMarker(PoopLocation poopLocation, double width, double height,
    double iconSize, BuildContext ctx) {
  return Marker(
      point: poopLocation.location(),
      width: width,
      height: height,
      child: GestureDetector(
        onTap: () => showMarkerInfoDialog(ctx, poopLocation),
        child: Icon(
          Icons.location_pin,
          color: poopLocation.color(),
          size: iconSize,
        ),
      ));
}
