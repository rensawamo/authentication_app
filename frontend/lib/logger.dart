import 'package:flutter/foundation.dart';
import 'package:logger/logger.dart';

// アプリの logger設定を定義
final logger = Logger(
  printer: PrettyPrinter(
    methodCount: kDebugMode ? 2 : 0,
    errorMethodCount: 8,
    lineLength: 80,
    colors: true,
    printEmojis: true,
    printTime: true,
  ),
  level: kDebugMode ? Level.debug : Level.warning,
);

void setupLogging() {
  Logger.level = kDebugMode ? Level.debug : Level.warning;
}
