import 'package:firebase_core/firebase_core.dart';
import 'package:flutter/material.dart';
import 'package:frontend/logger.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:firebase_auth/firebase_auth.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Firebase.initializeApp();
  await dotenv.load(fileName: '.env');
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Google Auth Example',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: MyHomePage(title: 'Flutter Demo Home Page'),
    );
  }
}

class MyHomePage extends StatefulWidget {
  MyHomePage({super.key, required this.title});
  final String title;

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  // アクセストークン
  String widgetCurrentToken = '';
  // CustomToken
  String widgetCustomToken = '';
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).colorScheme.inversePrimary,
        title: Text(widget.title),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            ElevatedButton(
              onPressed: () async {
                try {
                  final FirebaseAuth auth = FirebaseAuth.instance;
                  await auth.signInWithEmailAndPassword(
                    email: dotenv.get('EMAIL'),
                    password: dotenv.get('PASSWORD'),
                  );
                  final User? currentUser = auth.currentUser;
                  if (currentUser != null) {
                    final idTokenResult = await currentUser.getIdTokenResult();
                    final token = idTokenResult.token;
                    setState(() {
                      widgetCurrentToken = token!;
                    });
                    logger.d('IDトークン: $token');
                    // IDトークンをバックエンドに送信
                    final response = await http.post(
                      Uri.parse(
                          'http://10.0.2.2:8080/getCustomToken'), //android emulator 想定
                      headers: {
                        'Content-Type': 'application/json',
                        'AuthToken': token!,
                      },
                    );
                    if (response.statusCode == 200) {
                      final customToken = response.body;
                      final customTokenJson = jsonDecode(customToken);

                      setState(() {
                        widgetCustomToken = customTokenJson['customToken'];
                      });
                      logger.d('CustomToken: $customToken');
                    }
                  } else {
                    print('サインインに失敗しました');
                  }
                } catch (e) {
                  print("サインイン失敗: $e");
                }
              },
              child: const Text('サインインした後にカスタムトークンを取得'),
            ),
            ElevatedButton(
              /// CustomTokenを使用してrefreshTokenを取得
              onPressed: () async {
                final Uri signInUrl = Uri.parse(
                    'https://identitytoolkit.googleapis.com/v1/accounts:signInWithCustomToken?key=${dotenv.get('API_KEY')}');

                final refleshTokenData = await http.post(
                  signInUrl,
                  headers: {'Content-Type': 'application/json'},
                  body: jsonEncode({
                    'token': widgetCustomToken,
                    'returnSecureToken': true,
                  }),
                );

                if (refleshTokenData.statusCode == 200) {
                  final result = jsonDecode(refleshTokenData.body);
                  logger.d("refreshToken: ${result['refreshToken']}");

                  /// refreshTokenを使用してIDトークンを再取得
                  final Uri getNewTokenUrl = Uri.parse(
                      'https://securetoken.googleapis.com/v1/token?key=${dotenv.get('API_KEY')}');

                  final reult = await http.post(
                    getNewTokenUrl,
                    headers: {'Content-Type': 'application/json'},
                    body: jsonEncode({
                      'grant_type': 'refresh_token',
                      'refresh_token': result['refreshToken'],
                    }),
                  );
                  if (reult.statusCode == 200) {
                    final responseData = jsonDecode(reult.body);
                    logger.d("IDトークン: ${responseData['id_token']}");
                    setState(() {
                      widgetCurrentToken = responseData["id_token"];
                    });
                  }
                } else {
                  print('Failed to sign in with Custom Token.');
                }
              },
              child: const Text('Reflesh Token'),
            ),
            ElevatedButton(
              onPressed: () async {
                final response = await http.get(
                  Uri.parse("http://10.0.2.2:8080/example"),
                  headers: {
                    'Content-Type': 'application/json',
                    'AuthToken': widgetCurrentToken,
                  },
                );
                logger.d(response.body);
              },
              child: const Text('[正常系] APIリクエスト'),
            ),
            ElevatedButton(
              onPressed: () async {
                final response = await http.get(
                  Uri.parse("http://10.0.2.2:8080/example"),
                  headers: {
                    'Content-Type': 'application/json',
                    'AuthToken': "invalidToken",
                  },
                );
                logger.d(response.body);
              },
              child: const Text('[401] Unauthorized APIリクエスト'),
            ),
          ],
        ),
      ),
    );
  }
}
