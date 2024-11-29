class Review {
  String uuid;
  String poopLocationUUID;
  int rating;
  String time;
  int upvotes;
  int downvotes;

  Review(
      {required this.uuid,
      required this.poopLocationUUID,
      required this.rating,
      required this.time,
      required this.upvotes,
      required this.downvotes});

  Map<String, dynamic> toJson() {
    return {
      'uuid': uuid,
      'poop_location_uuid': poopLocationUUID,
      'rating': rating,
      'time': time,
      'upvotes': upvotes,
      'downvotes': downvotes,
    };
  }

  factory Review.fromJson(Map<String, dynamic> json) {
    return Review(
      uuid: json['uuid'],
      poopLocationUUID: json['poop_location_uuid'],
      rating: json['rating'],
      time: json['time'],
      upvotes: json['upvotes'],
      downvotes: json['downvotes'],
    );
  }
}
