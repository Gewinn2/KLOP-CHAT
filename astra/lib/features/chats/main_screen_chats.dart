import 'package:astra/features/chats/add_chat.dart';
import 'package:astra/features/chats/chat_panel.dart';
import 'package:astra/features/chats/struct_people.dart';
import 'package:astra/repo/repo_post.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

class MainChatScreen extends StatefulWidget {
  const MainChatScreen({super.key});


  @override
  _MainChatScreen createState() => _MainChatScreen();
}

class _MainChatScreen extends State<MainChatScreen> {

  List<User> _users = [];
  List<User> _users_important = [];
  String jwt = "";
  String user_id = "";
  String user_name = "";

  @override
  void didChangeDependencies() {
    final args = ModalRoute.of(context)?.settings.arguments;
    assert(args != null, "Check args");
    print(args);
    List<Object> help = args as List<Object>;
    _users = help[0] as List<User>;
    jwt = help[1] as String;
    user_id = help[2] as String;
    _users_important = help[3] as List<User>;
    user_name = help[4] as String;
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

  final TextEditingController _msgController = TextEditingController();
  late User us;
  bool flag = false;



  // Заглушка для списка пользователей
  //final List<String> _users = ['Alice', 'Bob', 'Charlie', 'Diana','And','Pop','Wik',"Wed",'Marik','Andro','Pipa','Tata'];

  @override
  Widget build(BuildContext context) {
    String chatName;
    String url='';
    String user_id_str;
    return Scaffold(
      appBar: AppBar(   
        //automaticallyImplyLeading: false, 
         iconTheme: IconThemeData(
          color: Color.fromRGBO(59, 3, 102, 1), // Установите нужный цвет стрелочки здесь
        ),
        toolbarHeight: 80.0,
        title:  Row(
          children: [
            SizedBox(width: 10),
            Text(
              "Klop Chat",
              style: TextStyle( 
                color: Color.fromRGBO(59, 3, 102, 1),
                fontSize: 40,
                fontFamily: 'Inter',
                fontWeight: FontWeight.w700,
              )
              ),
              SizedBox(width: 60,),
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
                  backgroundColor: Color.fromRGBO(164, 128, 222, 1),
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
            child: Column(
              children: [
                SizedBox(height: 10,),
                Container(
                  width: 350, // Ширина контейнера
                  height: 40, // Высота контейнера
                  padding: EdgeInsets.symmetric(horizontal: 8.0, vertical: 4.0), // Уменьшенные отступы
                  decoration: BoxDecoration(
                    color: Color.fromRGBO(230, 206, 241, 1), // Фиолетовый цвет фона
                    borderRadius: BorderRadius.circular(10.0), // Сглаживание краев
                  ),
                  alignment: Alignment.center, // Центрирование текста по вертикали и горизонтали
                  child: Text(
                    user_name,
                    style: TextStyle(
                      color: Color.fromRGBO(31, 2, 53, 1), // Цвет текста
                      fontSize: 24, // Размер шрифта
                    ),
                    overflow: TextOverflow.ellipsis, // Троеточие, если текст не помещается
                  ),
                ),
                SizedBox(height: 15,),
                Expanded(
                child: TextField(
                  controller: _msgController,
                  decoration: InputDecoration(
                    filled: true,
                    fillColor: Theme.of(context).colorScheme.onInverseSurface,
                    hintText: "Введите название чата",
                    prefixIcon: IconButton(
                      icon: Icon(Icons.search),
                      onPressed: () async {
                        try{
                          us = await getUser_by_name(jwt,_msgController.text);
                          setState(() {
                            flag = true;
                          });
                        }catch(e){
                          print(e);
                        }
                      },
                    ), // Иконка лупы теперь в виде кнопки
                    border: const OutlineInputBorder(
                      borderSide: BorderSide.none,
                      borderRadius: BorderRadius.all(
                        Radius.circular(12),
                      ),
                    ),
                  ),
                ),
              ),
              flag?ChatBubble(imageUrl: us.photo, chatTitle: us.name, last_mes: us.message_created_at) : Container(),


                //SizedBox(height: 10,),
                !(_users_important.length == 0)?Text(
                  'Рекомендованные чаты',
                  style: TextStyle( 
                      color: Color.fromRGBO(59, 3, 102, 1),
                      fontSize: 22,
                      fontFamily: 'Inter',
                      fontWeight: FontWeight.w700,
                    ),
                ) : Container(),
                SizedBox(height: 10,),
                (_users_important.length == 0)? Text('Добавьте чаты',
                style: TextStyle( 
                      color: Color.fromRGBO(59, 3, 102, 1),
                      fontSize: 22,
                      fontFamily: 'Inter',
                      fontWeight: FontWeight.w700,
                    ),)
                :Flexible(
                  child: ListView.builder(
                    itemCount: _users_important.length,
                    itemBuilder: (context, index) {
                      return ListTile(
                        title: ChatBubble(
                          chatTitle: _users_important[index].name, 
                          imageUrl: url, 
                          last_mes: _users_important[index].content, 
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
                //SizedBox(height: 10,),
                !(_users_important.length == 0)?const Text(
                  'Все чаты',
                  style: TextStyle( 
                      color: Color.fromRGBO(59, 3, 102, 1),
                      fontSize: 22,
                      fontFamily: 'Inter',
                      fontWeight: FontWeight.w700,
                    ),
                ): Container(),
                SizedBox(height: 10,),
                !(_users.length == 0)?Flexible(
                  child: ListView.builder(
                    itemCount: _users.length,
                    itemBuilder: (context, index) {
                      return ListTile(
                        title: ChatBubble(
                          chatTitle: _users[index].name, 
                          imageUrl: url, 
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
                ) : Container(),
              ],
            ),
          ),
          // Разделительная линия
          VerticalDivider(
            width: 1, // Ширина линии
            color: Colors.grey, // Цвет линии
          ),
          flag_tab?// Панель чат
          Expanded(
            flex: 8, // Занимает 3/4 пространства
            child: ChatPanel(user: _users[_selectedUserIndex],jwt: jwt,message: message, flag: flag_tab, user_id_tot_sami: user_id,),
          ):Container(),
        ],
      ),
    );
  }
}

