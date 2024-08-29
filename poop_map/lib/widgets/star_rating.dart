import 'package:flutter/material.dart';

class StarRating extends StatelessWidget {
  final int rating;
  final Function(int)? onRatingChanged;
  final Color color;
  final bool readOnly;

  const StarRating({
    super.key,
    required this.rating,
    this.onRatingChanged,
    this.color = Colors.amber,
    this.readOnly = true
  });

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisSize: MainAxisSize.min,
      children: List.generate(5, (index) {
        Widget starIcon = Icon(
          index < rating ? Icons.star : Icons.star_border,
          color: color,
          size: 30,
        );
        if (readOnly) {
          return starIcon;
        } else {
        return IconButton(
          icon: starIcon,
          onPressed: () => onRatingChanged?.call(index + 1),
          iconSize: 30,
        );
        }
      }),
    );
  }
}