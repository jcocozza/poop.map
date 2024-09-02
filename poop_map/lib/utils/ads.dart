import 'dart:io';
import 'package:flutter/material.dart';
import 'package:poop_map/utils/read_config.dart';
import 'package:google_mobile_ads/google_mobile_ads.dart';

/// Useful for ad info on a platform by platform basis
class AdHelper {
  final Config config;

  AdHelper({required this.config});

  String get bannerAdUnitId {
    if (Platform.isAndroid) {
      return config.androidBanner;
    } else if (Platform.isIOS) {
      return config.iosBanner;
    } else {
      throw UnsupportedError('Unsupported platform');
    }
  }

  String get adMobId {
    if (Platform.isAndroid) {
      return config.androidAdMobAppId;
    } else if (Platform.isIOS) {
      return config.iosAdMobAppId;
    } else {
      throw UnsupportedError('Unsupported platform');
    }
  }
}

class AdBanner extends StatefulWidget {
  final AdHelper adHelper;

  const AdBanner({
    super.key,
    required this.adHelper,
  });


  @override
  State<AdBanner> createState() => _AdBannerState();
}

class _AdBannerState extends State<AdBanner> {
  BannerAd? _ad;

  Future<InitializationStatus> _initGoogleMobileAds() {
    return MobileAds.instance.initialize();
  }

  @override
  void initState() {
    _initGoogleMobileAds();
    BannerAd ad = BannerAd(
      adUnitId: widget.adHelper.bannerAdUnitId,
      request: const AdRequest(),
      size: AdSize.banner,
      listener: BannerAdListener(
        onAdFailedToLoad: (ad, err) {
          print('Failed to load a banner ad: ${err.message}');
        },
      ),
    );
    setState(() {
      _ad = ad;
      _ad?.load();
    });
    super.initState();
  }

  @override
  void dispose() {
    _ad?.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return AdWidget(
      ad: _ad as BannerAd,
    );
  }
}
