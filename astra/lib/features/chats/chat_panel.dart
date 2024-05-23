import 'dart:async';

import 'package:astra/features/chats/struct_mes.dart';
import 'package:astra/repo/repo_post.dart';
import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

class ChatPanel extends StatefulWidget {
  final User user;
  final String jwt;
  List<Message>message;
  final bool flag;
  final String user_id_tot_sami;

  ChatPanel({super.key, required this.user,required this.jwt,required this.message, required this.flag, required this.user_id_tot_sami});

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
                child: Text(widget.user.name[0]),
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

   Timer? _timer; // Добавляем переменную для таймера

  @override
  void initState() {
    super.initState();
    _startTimer(); // Запускаем таймер при создании виджета
    // ... ваш код 
  }

  @override
  void dispose() {
    _timer?.cancel(); // Останавливаем таймер при удалении виджета
    super.dispose();
  }

  void _startTimer() {
    _timer = Timer.periodic(Duration(seconds: 100), (timer) async {
      // Вызываем функцию getMessage каждые 10 секунд
      List<Message> help = await getMessage(widget.jwt, widget.user.chat_id);
      update(help);
      // ... ваш код для обработки данных help 
      setState(() {}); // Обновляем состояние виджета, чтобы изменения отобразились
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
        !widget.flag? Text("Нажми на чат") 
        :Expanded( 
          child: ListView.builder(
            itemCount: widget.message.length,
            itemBuilder: (context,index){
              //List<Message> h = widget.message;
              print("--------------------");
              print(widget.message[index].user_id);
              print(widget.user_id_tot_sami);
              return StructMessage(
                  message: widget.message[index].content,
                  id_who_send: widget.message[index].user_id,
                  id_i: widget.user_id_tot_sami,  // сюда впишу номер чела которого мне вернут при входе в ак( айдишник)
                  time1: widget.message[index].created_at,
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
                    _startTimer();
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
