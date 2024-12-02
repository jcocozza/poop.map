import 'package:flutter/material.dart';
import 'package:frontend/model/review.dart';
import 'package:frontend/requests/review.dart';

class ReviewViewer extends StatefulWidget {
  final Review review;

  const ReviewViewer({
    super.key,
    required this.review,
  });

  @override
  State<ReviewViewer> createState() => _ReviewViewerState();
}

class _ReviewViewerState extends State<ReviewViewer> {
  void upvoter() async {
    await upvote(widget.review.uuid);
    setState(() {
      widget.review.upvotes += 1;
    });
  }

  void downvoter() async {
    await downvote(widget.review.uuid);
    setState(() {
      widget.review.downvotes += 1;
    });
  }

  @override
  Widget build(BuildContext content) {
    return Column(
      children: [
        Text('${widget.review.rating} - ${widget.review.time}'),
        Text(widget.review.comment),
        Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            Text(
              'Votes: ${widget.review.downvotes}',
            ),
            IconButton(
              icon: const Icon(Icons.arrow_downward),
              onPressed: () => downvoter(),
            ),
            Text(
              'Votes: ${widget.review.upvotes}',
            ),
            IconButton(
              icon: const Icon(Icons.arrow_upward),
              onPressed: () => upvoter(),
            ),
          ],
        ),
      ],
    );
  }
}

class ReviewList extends StatefulWidget {
  final String poopLocationUUID;

  const ReviewList({
    super.key,
    required this.poopLocationUUID,
  });

  @override
  State<ReviewList> createState() => _ReviewListState();
}

class _ReviewListState extends State<ReviewList> {
  List<Review> reviewList = [];

  @override
  Future<void> initState() async {
    final reviews = await getAllReviewsByPoopLocation(widget.poopLocationUUID);
    setState(() {
      reviewList = reviews;
    });
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Column(
        children: reviewList.map((review) {
          return ReviewViewer(
            review: review,
          );
        }).toList(),
      ),
    );
  }
}
