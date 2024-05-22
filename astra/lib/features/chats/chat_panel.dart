import 'package:astra/features/chats/struct_mes.dart';
import 'package:astra/repo/repo_post.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

class ChatPanel extends StatefulWidget {
  final User user;
  final String jwt;
  List<Message>message;
  final bool flag;

  ChatPanel({super.key, required this.user,required this.jwt,required this.message, required this.flag});

  @override
  _ChatPanelState createState() => _ChatPanelState();
}

class _ChatPanelState extends State<ChatPanel> {
  void clearMessageController() {
    setState(() {
      _msgController.text = '';
    });
     // Устанавливаем текст в пустую строку
    _msgController.selection = TextSelection.fromPosition(TextPosition(offset: 0)); // Сбрасываем позицию курсора в начало текста
}
  final TextEditingController _msgController = TextEditingController();
  Widget _buildThumbnailImage(String url) {
    try {
      return SizedBox(
        width: 48,
        height: 48,
        child: ClipRRect(
          borderRadius: BorderRadius.circular(32),
          child: Image.network(
            url,
            fit: BoxFit.fill,
            height: 256,
            errorBuilder: (
              BuildContext context,
              Object exception,
              StackTrace? stackTrace,
            ) {
              return CircleAvatar(
                radius: 6,
                backgroundColor: Theme.of(context).colorScheme.primaryContainer,
                child: const Text('A'),
              );
            },
          ),
        ),
      );
    } catch (e) {
      return Container();
    }
  }
  void update(List<Message> a) {
    setState(() {
      widget.message = a;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        automaticallyImplyLeading: false,
        backgroundColor: Theme.of(context).colorScheme.background,
        title: Column(
          children: [
            Container(
              width: MediaQuery.of(context).size.width,
              child: ElevatedButton(
                style: ButtonStyle(
                  padding: const MaterialStatePropertyAll(EdgeInsets.zero),
                  surfaceTintColor:
                      const MaterialStatePropertyAll(Colors.transparent),
                  backgroundColor: const MaterialStatePropertyAll(
                    Colors.transparent,
                  ),
                  shadowColor: const MaterialStatePropertyAll(Colors.transparent),
                  shape: MaterialStatePropertyAll(
                    RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(12),
                    ),
                  ),
                ),
                onPressed: () {
                },
                child: Row(
                  children: [
                    _buildThumbnailImage("https://i.pinimg.com/736x/83/8f/84/838f840d51db226a8a0cb79b962f24d5.jpg"),
                    SizedBox(width: 15,),
                    //const Padding(padding: EdgeInsets.only(right: 12)),
                    Container(
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Row(
                            children: [
                              SizedBox(
                                width: MediaQuery.of(context).size.width * 0.35,
                                child: Text(
                                  widget.user.name,
                                  style: TextStyle(
                                    color: Theme.of(context)
                                        .colorScheme
                                        .onPrimaryContainer,
                                        fontFamily: 'Inter',
                                        fontWeight: FontWeight.w500,
                                    fontSize: 20,
                                  ),
                                  softWrap: true,
                                  overflow: TextOverflow.ellipsis,
                                ),
                              ),
            
                            ],
                          ),
                        
                        ],
                      ),
                    ),
                  ],
                ),
              ),
            ),
            SizedBox(height: 5,),
            Divider(
              height: 1, // Высота разделителя, включая линию и пустое пространство над и под ней
              thickness: 1, // Толщина линии разделителя
              color: Colors.grey, // Цвет линии разделителя
            ),
          ],
        ),
      ),
      body: Column(
        mainAxisAlignment: MainAxisAlignment.spaceBetween, // Разделяет пространство между списком сообщений и панелью ввода
        children: [
        //   Column(
        //   children: [
        //     Expanded(
        //       child: ListView.builder(
        //         itemCount: widget.message.length,
        //         itemBuilder: (context, index) {
        //           //final message = widget.message[index].content;
        //           print(widget.user.chat_id);
        //           print(widget.message[index].user_id);
        //           // return StructMessage(
        //           //   message: widget.message[index].content,
        //           //   id_who_send: widget.message[index].user_id,
        //           //   id_i: widget.user.chat_id,
        //           // );
        //           return Text(widget.message[index].content);
        //         },
        //       ),
        //     ),
        //   ],
        // ),
        !widget.flag? Text("Нажми на чат") 
        // :   Column(
        //   children: [
        //     ListView.builder(
        //       itemCount: widget.message.length,
        //       itemBuilder: (context, index) {
        //         //final message = widget.message[index].content;
        //         print(widget.user.chat_id);
        //         print(widget.message[index].user_id);
        //         // return StructMessage(
        //         //   message: widget.message[index].content,
        //         //   id_who_send: widget.message[index].user_id,
        //         //   id_i: widget.user.chat_id,
        //         // );
        //         return Text(widget.message[index].content);
        //         //return Text("gaga");
        //       },
        //     ),
        //   ],
        // ),
        //:Text(widget.message.length.toString()),
        // :ListView.builder(
        //   itemCount: widget.message.length, 
        //   itemBuilder: (context, index) {
        //     var message = widget.message[index];
        //     return ListTile(
        //       title: Text(message.content), 
        //     );
        //   },
        // ),
        // Expanded(
        //     flex: 3, // Занимает 1/4 пространства
        //     child: ListView.builder(
        //       itemCount: _users.length,
        //       itemBuilder: (context, index) {
        //         return ListTile(
        //           title: ChatBubble(
        //             chatTitle: _users[index].name, 
        //             imageUrl: 'https://img.freepik.com/free-photo/abstract-textured-backgound_1258-30627.jpg?size=338&ext=jpg&ga=GA1.1.44546679.1716163200&semt=ais_user', 
        //           ), //Text(_users[index])
        :Expanded( 
          child: ListView.builder(
            itemCount: widget.message.length,
            itemBuilder: (context,index){
              //List<Message> h = widget.message;
              print(widget.user.chat_id);
              print(widget.message[index].user_id);
              return StructMessage(
                  message: widget.message[index].content,
                  id_who_send: widget.message[index].user_id,
                  id_i: '3',  // сюда впишу номер чела которого мне вернут при входе в ак( айдишник)
                );
            },
          ),
        ),
          Align(
            alignment: Alignment.bottomCenter,
            child: Container(
              width: double.infinity,
              decoration: BoxDecoration(
                color: Theme.of(context).colorScheme.onInverseSurface,
                borderRadius: const BorderRadius.only(
                  topLeft: Radius.circular(12),
                  topRight: Radius.circular(12),
                ),
              ),
              padding: const EdgeInsets.all(6),
              child: Row(
                children: [
                  // Заменить на стикеры
                  // IconButton(
                  //   onPressed: () {},
                  //   icon: Icon(Icons.insert_emoticon),
                  // ),
                  Expanded(
                    child: TextField(
                      controller: _msgController,
                      decoration: InputDecoration(
                        filled: true,
                        fillColor: Theme.of(context).colorScheme.onInverseSurface,
                        hintText: "Введите сообщение...",
                        border: const OutlineInputBorder(
                          borderSide: BorderSide.none,
                          borderRadius: BorderRadius.all(
                            Radius.circular(12),
                          ),
                        ),
                      ),
                    ),
                  ),
                  IconButton(
                    onPressed: () {
                      // Действие при нажатии на кнопку прикрепления файла
                    },
                    icon: const Icon(Icons.attach_file_outlined),
                  ),
                  IconButton(
                    onPressed: () async {
                        String content = _msgController.text;

                        try {
                          final response = await post_message(content, widget.user.chat_id,widget.jwt );
                          if (response.statusCode == 200) {
                            print("сообщение отправлено");
                          } else {
                            // Обработка ошибки
                            print('Ошибка: ${response.statusCode}');
                          }
                        } catch (e) {
                          // Обработка исключения при отправке запроса
                          print('Исключение при отправке запроса: $e');
                        }
                        clearMessageController();
                        try {
                          List<Message> help = await getMessage(widget.jwt,widget.user.chat_id);
                          update(help);
                          //print(message);
                          
                    } catch (e) {
                          // Обработка исключения при отправке запроса
                        print('Исключение при отправке запроса: $e');
                    }
                      },
                    icon: const Icon(Icons.send),
                  ),
                ],
              ),
            ),
          ),
        ],
      ),
    );
  }
}
