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

  @override
  void didChangeDependencies() {
    final args = ModalRoute.of(context)?.settings.arguments;
    assert(args != null, "Check args");
    print(args);
    List<User> help = args as List<User>;
    _users = help;
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

  // Заглушка для списка пользователей
  //final List<String> _users = ['Alice', 'Bob', 'Charlie', 'Diana','And','Pop','Wik',"Wed",'Marik','Andro','Pipa','Tata'];

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(   
        //automaticallyImplyLeading: false, 
        title: const Column(
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
                    imageUrl: 'https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTH-BPnpCcCSysCnqNBWDDAnYGNgnFyRpCrP4cQ6NGgqA&s', 
                  ), //Text(_users[index])
                  selected: index == _selectedUserIndex,
                  onTap: () {
                    setState(() {
                      _selectedUserIndex = index;
                    });
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
            child: ChatPanel(user_name: _users[_selectedUserIndex].name),
          ),
        ],
      ),
    );
  }
}

