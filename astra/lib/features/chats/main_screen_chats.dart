import 'package:astra/features/chats/add_chat.dart';
import 'package:astra/features/chats/chat_panel.dart';
import 'package:astra/features/chats/struct_people.dart';
import 'package:astra/repo/repo_post.dart';
import 'package:flutter/material.dart';

class MainChatScreen extends StatefulWidget {
  const MainChatScreen({super.key});


  @override
  _MainChatScreen createState() => _MainChatScreen();
}

class _MainChatScreen extends State<MainChatScreen> {

  List<User> _users = [];
  String jwt = "";
  String user_id = "";

  @override
  void didChangeDependencies() {
    final args = ModalRoute.of(context)?.settings.arguments;
    assert(args != null, "Check args");
    print(args);
    List<Object> help = args as List<Object>;
    _users = help[0] as List<User>;
    jwt = help[1] as String;
    user_id = help[2] as String;
    // user_name =
    //     help["name"] as String?; // Присваивание значения переменной user_name
    // image_url = help["url"] as String?;
    // vol = help["volume"] as bool;
    // pin = help["pin"] as bool;
    // uid = help["uid"] as String;
    // print("GroupID: $uid");
    // userPhone = help["phone"] as String;

    setState(() {});
    super.didChangeDependencies();
  }


  int _selectedUserIndex = 0;
  bool flag_tab = false;
  List<Message>message = [];
  void update(List<Message> a) {
    setState(() {
      message = a;
    });
  }

  void update_user(List<User>a){
    setState(() {
      _users = a;
    });
  }


  void update_id(int a) {
    setState(() {
      _selectedUserIndex = a;
    });
  }

  // Заглушка для списка пользователей
  //final List<String> _users = ['Alice', 'Bob', 'Charlie', 'Diana','And','Pop','Wik',"Wed",'Marik','Andro','Pipa','Tata'];

  @override
  Widget build(BuildContext context) {
    String chatName;
    String url;
    String user_id_str;
    return Scaffold(
      appBar: AppBar(   
        //automaticallyImplyLeading: false, 
        title:  Row(
          children: [
            SizedBox(height: 20,width: 20,),
            Text(
              "Klop Chat",
              style: TextStyle( 
                color: Color.fromRGBO(59, 3, 102, 1),
                fontSize: 40,
                fontFamily: 'Inter',
                fontWeight: FontWeight.w700,
              )
              ),
              ElevatedButton(
                onPressed: () async {
                  List<String> help = await addChat(context);
                  chatName = help[0];
                  url = help[1];
                  user_id_str = help[2];
                  List<String> user_id = user_id_str.split(' ');
                  print("Chat name: $chatName");
                  try{
                    final response = await create_chat(chatName,url,user_id,jwt);
                          if (response.statusCode == 200) {
                            print("создана группа");
                          } else {
                            // Обработка ошибки
                            print('Ошибка: ${response.statusCode}');
                          }
                          List<User> users = await getUser(jwt);
                          update_user(users);
                  }catch(e){
                    print("оширбка в создании чаата $e" );
                  }
                },
                style: ElevatedButton.styleFrom(
                  backgroundColor: Colors.purple,
                  shape: const CircleBorder(),
                  padding: const EdgeInsets.all(20),
                  elevation: 0, // Убираем тень
                ),
                child: const Icon(Icons.add, color: Colors.white), // Плюсик
              ),
          ],
        ),
        ),
      body: Row(
        children: [
          // Список пользователей
          Expanded(
            flex: 3, // Занимает 1/4 пространства
            child: ListView.builder(
              itemCount: _users.length,
              itemBuilder: (context, index) {
                return ListTile(
                  title: ChatBubble(
                    chatTitle: _users[index].name, 
                    imageUrl: 'https://img.freepik.com/free-photo/abstract-textured-backgound_1258-30627.jpg?size=338&ext=jpg&ga=GA1.1.44546679.1716163200&semt=ais_user', 
                    last_mes: _users[index].content, 
                  ), //Text(_users[index])
                  selected: index == _selectedUserIndex,
                  onTap: () async{
                    update_id(index);
                    setState(() {
                      flag_tab = true;
                    });
                    try {
                          List<Message> help = await getMessage(jwt, _users[index].chat_id);
                          update(help);
                          print(message);
                          
                    } catch (e) {
                          // Обработка исключения при отправке запроса
                        print('Исключение при отправке запроса: $e');
                    }
                  },
                );
              },
            ),
          ),
          // Разделительная линия
          VerticalDivider(
            width: 1, // Ширина линии
            color: Colors.grey, // Цвет линии
          ),
          // Панель чата
          Expanded(
            flex: 8, // Занимает 3/4 пространства
            child: ChatPanel(user: _users[_selectedUserIndex],jwt: jwt,message: message, flag: flag_tab, user_id_tot_sami: user_id,),
          ),
        ],
      ),
    );
  }
}

