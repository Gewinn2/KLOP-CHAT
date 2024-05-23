import 'package:flutter/material.dart';

import 'package:flutter/material.dart';

Future<List<String>> addChat(
  BuildContext context,
) async {
  final TextEditingController _msgController = TextEditingController();
  final TextEditingController _msgController1 = TextEditingController();
  final TextEditingController _msgController2 = TextEditingController();
  final Widget okButton = TextButton(
    style: ButtonStyle(
      overlayColor: MaterialStatePropertyAll(
        Theme.of(context).colorScheme.error.withAlpha(32),
      ),
    ),
    child: Text(
      "Отмена",
      style: TextStyle(
        color: Theme.of(context).colorScheme.error,
      ),
    ),
    onPressed: () {
      Navigator.of(context).pop();
    },
  );
  final Widget publishButton = TextButton(
    style: const ButtonStyle(
      overlayColor: MaterialStatePropertyAll(Color.fromARGB(64, 96, 11, 129)),
    ),
    child: const Text(
      "Создать",
      style: TextStyle(
        color: Color.fromARGB(255, 84, 9, 107),
      ),
    ),
    onPressed: () {
      List<String>help = [_msgController.text,_msgController1.text,_msgController2.text];
      Navigator.of(context).pop(help); // Возвращаем текст
    },
  );

  final AlertDialog alert = AlertDialog(
    title: const Text(
      "Введите название чата",
      style: TextStyle(
        fontSize: 28,
        fontWeight: FontWeight.bold,
      ),
    ),
    content: ConstrainedBox(
        constraints: BoxConstraints(
        minWidth: 300, // Минимальная ширина диалога
        maxWidth: 500, // Максимальная ширина диалога
        minHeight: 100, // Минимальная высота диалога
        maxHeight: 300, // Максимальная высота диалога
      ),
      child: Column(
        children: [
          Expanded(
            child: TextField(
              controller: _msgController,
              decoration: InputDecoration(
                filled: true,
                fillColor: Theme.of(context).colorScheme.onInverseSurface,
                hintText: "Введите название чата",
                border: const OutlineInputBorder(
                  borderSide: BorderSide.none,
                  borderRadius: BorderRadius.all(
                    Radius.circular(12),
                  ),
                ),
              ),
            ),
          ),
          Expanded(
            child: TextField(
              controller: _msgController1,
              decoration: InputDecoration(
                filled: true,
                fillColor: Theme.of(context).colorScheme.onInverseSurface,
                hintText: "Введите url ссылку на аву",
                border: const OutlineInputBorder(
                  borderSide: BorderSide.none,
                  borderRadius: BorderRadius.all(
                    Radius.circular(12),
                  ),
                ),
              ),
            ),
          ),
          Expanded(
            child: TextField(
              controller: _msgController2,
              decoration: InputDecoration(
                filled: true,
                fillColor: Theme.of(context).colorScheme.onInverseSurface,
                hintText: "Введите имена пользователя через пробел",
                border: const OutlineInputBorder(
                  borderSide: BorderSide.none,
                  borderRadius: BorderRadius.all(
                    Radius.circular(12),
                  ),
                ),
              ),
            ),
          ),
        ],
      ),
    ),
    backgroundColor: Theme.of(context).colorScheme.background,
    surfaceTintColor: Colors.white,
    actions: [
      okButton,
      publishButton,
    ],
  );

  List<String>? chatName = await showDialog<List<String>>(
    context: context,
    builder: (BuildContext context) {
      return alert;
    },
  );

  return chatName ?? []; // Возвращаем текст или пустую строку, если диалог был отменен
}
