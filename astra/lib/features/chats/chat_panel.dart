import 'package:flutter/material.dart';

class ChatPanel extends StatefulWidget {
  final String user_name;

  const ChatPanel({super.key, required this.user_name});

  @override
  _ChatPanelState createState() => _ChatPanelState();
}

class _ChatPanelState extends State<ChatPanel> {
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
                                  widget.user_name,
                                  style: TextStyle(
                                    color: Theme.of(context)
                                        .colorScheme
                                        .onPrimaryContainer,
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
          Expanded(
            child: Center(
              child: Text("Сообщения будут здесь"), // Здесь будет список сообщений
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
                    onPressed: () {
                      // Действие при нажатии на кнопку отправки сообщения
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
