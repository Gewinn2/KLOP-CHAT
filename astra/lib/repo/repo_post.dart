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
      port: 8080,
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
      port: 8080,
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
      port: 8080,
      path: '/auth/chat',
    ),
    headers: {
      'Authorization': 'Bearer $jwtToken',
    },
  );
  if (response.statusCode == 200) {
    final dynamic decodedData = json.decode(response.body);
    print(decodedData);
    if (decodedData == [0]) {return [];}
    List<User> newsList = [];
    final List<dynamic> jsonList = decodedData as List;
    // final dynamic hi = jsonList[0];
    // print(hi['image_link'][0]);
    newsList = jsonList.map((json) {
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

Future<List<User>> getImportantUser(String jwtToken) async {
  final response = await http.get(
    Uri(
      scheme: 'http',
      host: 'localhost',
      port: 8080,
      path: '/auth/chat/priority',
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


// {
//   "content": "Привет",
//   "chat_id": 1
// }


Future<http.Response> post_message(
  String content,
  String chat_id,
  String jwtToken,
  //String photo,
) async {
  int chatIdAsInt = -1;

  try {
    chatIdAsInt = int.parse(chat_id);
    print('chat_id как int: $chatIdAsInt');
  } catch (e) {
    // Обработка ошибки, если строка не содержит число, которое можно преобразовать
    print('Ошибка при преобразовании chat_id в int: $e');
  }
  return http.post(
    Uri(
      scheme: 'http',
      host: 'localhost',
      port: 8080,
      path: '/auth/message',
    ),
    headers: {
      'Authorization': 'Bearer $jwtToken',
    },
    body: jsonEncode({
      'content' : content,
      'chat_id' : chatIdAsInt,
    }),
  );
}

class Message {
    final String message_id;
    final String content;
    final String user_id;
    final String chat_id;
    final String created_at;
    Message({
    required this.message_id,
    required this.content,
    required this.user_id,
    required this.chat_id,
    required this.created_at,
  });
}


Future<List<Message>> getMessage(String jwtToken,String id) async {
  final response = await http.get(
    Uri(
      scheme: 'http',
      host: 'localhost',
      port: 8080,
      path: '/auth/message',
      queryParameters: {'id': id},
    ),
    headers: {
      'Authorization': 'Bearer $jwtToken',
    },
  );
  if (response.statusCode == 200) {
    final dynamic decodedData = json.decode(response.body);
    print(decodedData);
    List<Message> newsList = [];
    final List<dynamic> jsonList = decodedData as List;
    // final dynamic hi = jsonList[0];
    // print(hi['image_link'][0]);
    newsList = jsonList.map((json) {
      return Message(
          
        message_id: json['message_id'].toString(),
        content: json['content'],
       user_id: json['chat_id'].toString(),
       chat_id: json['user_id'].toString(),
       created_at: json['created_at']
      );
    }).toList();
    return newsList;
  } else {
    throw Exception(response.statusCode);
  }
}

Future<String> getUser_by_id(String jwtToken,String id) async {
  final response = await http.get(
    Uri(
      scheme: 'http',
      host: 'localhost',
      port: 8080,
      path: '/user',
      queryParameters: {'id': id},
    ),
    headers: {
      'Authorization': 'Bearer $jwtToken',
    },
  );
  if (response.statusCode == 200) {
    final dynamic decodedData = json.decode(response.body);
    print(decodedData);
    String name = decodedData['username'].toString();
    return name;
  } else {
    throw Exception(response.statusCode);
  }
}


// {
//   "name": "GR2",
//   "photo": "https://www.zastavki.com/pictures/originals/2018Animals___Cats_Large_gray_cat_with_a_surprised_look_123712_.jpg",
//   "user_id_arr": [
//     3
//   ]
// }

Future<http.Response> create_chat(
  String chat_name,
  String photo,
  List<String> user_id_arr,
  String jwtToken,
  //String photo,
) async {
  // Используя int.parse (будет исключение, если строка не является числом)



  return http.post(
    Uri(
      scheme: 'http',
      host: 'localhost',
      port: 8080,
      path: '/auth/chat',
    ),
    headers: {
      'Authorization': 'Bearer $jwtToken',
    },
    body: jsonEncode({
      'name' : chat_name,
      'photo' : photo,
      'username_arr': user_id_arr,
    }),
  );
}



Future<User> getUser_by_name(String jwtToken,String name) async {
  final response = await http.get(
    Uri(
      scheme: 'http',
      host: 'localhost',
      port: 8080,
      path: '/auth/chat/find',
      queryParameters: {'id': name},
    ),
    headers: {
      'Authorization': 'Bearer $jwtToken',
    },
  );
  if (response.statusCode == 200) {
    final dynamic decodedData = json.decode(response.body);
    print(decodedData);
    User newsList = User(
        chat_id: decodedData['chat_id'].toString(), 
        name: decodedData['name'], 
        photo: decodedData['photo'], 
        content: decodedData['content'], 
        message_created_at: decodedData['message_created_at'],
        
      );
    return newsList;
  } else {
    throw Exception(response.statusCode);
  }
}