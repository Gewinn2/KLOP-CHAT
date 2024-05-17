import 'package:astra/features/authorization/authorization.dart';
import 'package:astra/features/authorization/registration.dart';
import 'package:astra/features/chats/main_screen_chats.dart';
import 'package:astra/theme/them.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Chats',
      //initialRoute: firstLaunch ? 'welcome' : '/',
      routes: {
        '/': (context) => const Authorization(),
        'registration': (context) => const Registration(),
        'main_screen_chat': (context) => const MainChatScreen(),
      },
      onUnknownRoute: (settings) {},
    );
  }
}
