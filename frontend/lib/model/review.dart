class Review {
  String uuid;
  String poopLocationUUID;
  int rating;
  String comment;
  String time;
  int upvotes;
  int downvotes;

  Review(
      {required this.uuid,
      required this.poopLocationUUID,
      required this.rating,
      required this.comment,
      required this.time,
      required this.upvotes,
      required this.downvotes});

  Map<String, dynamic> toJson() {
    return {
      'uuid': uuid,
      'poop_location_uuid': poopLocationUUID,
      'rating': rating,
      'comment': comment,
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
      comment: json['comment'],
      time: json['time'],
      upvotes: json['upvotes'],
      downvotes: json['downvotes'],
    );
  }
}
