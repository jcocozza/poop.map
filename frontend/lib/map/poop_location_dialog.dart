import 'package:flutter/material.dart';
import 'package:latlong2/latlong.dart';
import 'package:frontend/model/poop_location.dart';
import 'package:frontend/requests/poop_locations.dart';

class PoopLocationDialog extends StatefulWidget {
  final LatLng location;
  final void Function(PoopLocation) onPoopLocationAdd;

  const PoopLocationDialog({
    super.key,
    required this.location,
    required this.onPoopLocationAdd,
  });

  @override
  State<StatefulWidget> createState() => _PoopLocationDialogState();
}

class _PoopLocationDialogState extends State<PoopLocationDialog> {
  // unique form identifier
  final _formKey = GlobalKey<FormState>();

  // poop location info
  final nameController = TextEditingController();
  String locationType = "regular";
  bool seasonal = false;
  //List<String> seasons = [];
  Map<String, bool> seasons = {
    "summer": false,
    "winter": false,
    "fall": false,
    "spring": false,
  };
  bool accessible = false;
  @override
  Widget build(BuildContext context) {
    return AlertDialog(
      title: const Text("Add Poop Location"),
      content: SingleChildScrollView(
        child: Form(
          key: _formKey,
          child: Column(
            children: <Widget>[
              TextField(
                controller: nameController,
                decoration: const InputDecoration(labelText: 'Name'),
              ),
              Row(
                children: [
                  DropdownMenu<String>(
                    initialSelection: "regular",
                    onSelected: (String? newValue) {
                      setState(() {
                        locationType = newValue!;
                      });
                    },
                    dropdownMenuEntries: [
                      "regular",
                      "porta potty",
                      "outhouse",
                      "other"
                    ].map((loc) {
                      return DropdownMenuEntry(
                        value: loc,
                        label: loc,
                      );
                    }).toList(),
                  ),
                  Row(
                    children: [
                      Checkbox(
                        value: accessible,
                        onChanged: (bool? value) {
                          setState(() {
                            accessible = value!;
                          });
                        },
                      ),
                      const Icon(Icons.accessible),
                    ],
                  ),
                ],
              ),
              Column(
                children: [
                  Center(
                    child: Row(children: [
                      const Text('Seasonal'),
                      Checkbox(
                        value: seasonal,
                        onChanged: (bool? value) {
                          setState(() {
                            seasonal = value!;
                          });
                        },
                      ),
                    ]),
                  ),
                  if (seasonal)
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                      children: [
                        Column(
                          children: [
                            const Text("summer"),
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
                    )
                ],
              ),
            ],
          ),
        ),
      ),
      actions: <Widget>[
        TextButton(
          onPressed: () {
            Navigator.of(context).pop();
          },
          child: const Text('cancel'),
        ),
        ElevatedButton(
          child: const Text('Add'),
          onPressed: () async {
            final name = nameController.text;
            List<String> selectedSeasons = seasons.entries
                .where((entry) => entry.value)
                .map((entry) => entry.key)
                .toList();
            PoopLocation pl = PoopLocation(
              uuid: '',
              name: name,
              latitude: widget.location.latitude,
              longitude: widget.location.longitude,
              firstCreated: '',
              lastModified: '',
              locationType: locationType,
              seasonal: seasonal,
              seasons: selectedSeasons,
              accessible: accessible,
              upvotes: 0,
              downvotes: 0,
            );
            await createPoopLocation(pl);
            if (context.mounted) {
              Navigator.of(context).pop();
              widget.onPoopLocationAdd(pl);
            }
          },
        ),
      ],
    );
  }
}

void showAddPoopLocationDialog(BuildContext context, LatLng location,
    void Function(PoopLocation) onAddPoopLocation) async {
  return showDialog(
      context: context,
      barrierDismissible: false,
      builder: (BuildContext context) {
        return PoopLocationDialog(
          location: location,
          onPoopLocationAdd: onAddPoopLocation,
        );
      });
}
