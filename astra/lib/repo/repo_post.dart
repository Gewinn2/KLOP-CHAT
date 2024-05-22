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


class User {
  // {
  //   "chat_id": 0,
  //   "name": "string",
  //   "photo": "string",
  //   "content": "string",
  //   "message_created_at": "2024-05-21T13:11:44.028Z"
  // }
  final String chat_id;
  final String name;
  final String photo;
  final String content;
  final String message_created_at;

  User({
    required this.chat_id,
    required this.name,
    required this.photo,
    required this.content,
    required this.message_created_at,
  });

}

///auth/message
//
///

Future<List<User>> getUser(String jwtToken) async {
  final response = await http.get(
    Uri(
      scheme: 'http',
      host: 'localhost',
      port: 5050,
      path: '/auth/chat',
    ),
    headers: {
      'Authorization': 'Bearer $jwtToken',
    },
  );
  if (response.statusCode == 200) {
    final dynamic decodedData = json.decode(response.body);
    print(decodedData);
    List<User> newsList = [];
    final List<dynamic> jsonList = decodedData as List;
    // final dynamic hi = jsonList[0];
    // print(hi['image_link'][0]);
    newsList = jsonList.map((json) {
      List<dynamic> dynamicList = json['tags']==null? [] : json['tags'] as List<dynamic>;
      List<String> stringList = dynamicList.map((item) => item.toString()).toList();
      return User(
        chat_id: json['chat_id'].toString(), 
        name: json['name'], 
        photo: json['photo'], 
        content: json['content'], 
        message_created_at: json['message_created_at'],
        
      );
    }).toList();
    return newsList;
  } else {
    throw Exception(response.statusCode);
  }
}
