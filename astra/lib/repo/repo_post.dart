import 'dart:convert';
import 'package:http/http.dart' as http;


Future<http.Response> post_sign_in(
  String email,
  String pass,
) async {
  return http.post(
    Uri(
      scheme: 'http',
      host: 'localhost',
      port: 5050,
      path: '/sign-in',
    ),
    headers: {
      'Content-Type': 'application/json; charset=UTF-8',
    },
    body: jsonEncode({
      'email': email,
      'password': pass,
    }),
  );
}

// {
//   "username": "vanek",
//   "email": "ivanovivan@yandex.ru",
//   "password": "qwerty1234",
//   "photo": "https://gas-kvas.com/grafic/uploads/posts/2023-09/1695826313_gas-kvas-com-p-kartinki-s-kotikami-1.jpg"
// }

Future<http.Response> post_sign_up(
  String username,
  String email,
  String password,
  //String photo,
) async {
  return http.post(
    Uri(
      scheme: 'http',
      host: 'localhost',
      port: 5050,
      path: '/sign-up',
    ),
    headers: {
      'Content-Type': 'application/json; charset=UTF-8',
    },
    body: jsonEncode({
      "username": username,
      'email': email,
      'password': password,
      "photo": "https://gas-kvas.com/grafic/uploads/posts/2023-09/1695826313_gas-kvas-com-p-kartinki-s-kotikami-1.jpg"
    }),
  );
}