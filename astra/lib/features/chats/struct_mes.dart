import 'dart:typed_data';
import 'dart:ui';

import 'package:flutter/material.dart';

import 'dart:ui';
import 'package:flutter/material.dart';
import 'package:flutter/painting.dart';

// {
//     "message_id": 3,
//     "content": "Это сообщение с веба",
//     "user_id": 4,
//     "chat_id": 13,
//     "created_at": "2024-05-22T00:00:00Z"
//   },



class StructMessage extends StatelessWidget {
  final String message;
  final String id_who_send;
  final String id_i;

  const StructMessage({
    Key? key,
    required this.message,
    required this.id_who_send,
    required this.id_i,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // Проверяем, совпадают ли идентификаторы
    bool isSentByMe = id_i == id_who_send;

    // Выбираем выравнивание в зависимости от отправителя
    Alignment messageAlignment = isSentByMe ? Alignment.centerRight : Alignment.centerLeft;

    // Выбираем цвет в зависимости от отправителя
    Color messageColor = isSentByMe ? Colors.pink : Colors.purple;

    return Align(
      alignment: messageAlignment,
      child: Container(
        padding: EdgeInsets.symmetric(horizontal: 10.0, vertical: 8.0), // Отступы внутри контейнера
        margin: EdgeInsets.symmetric(horizontal: 15.0, vertical: 5.0), // Отступы вокруг контейнера
        constraints: BoxConstraints(maxWidth: MediaQuery.of(context).size.width * 0.8), // Максимальная ширина сообщения
        decoration: BoxDecoration(
          color: messageColor, // Цвет фона сообщения
          borderRadius: BorderRadius.circular(8.0), // Скругление углов
        ),
        child: Text(
          message,
          style: TextStyle(
            color: Colors.white,
            fontFamily: 'Inter',
            fontSize: 16,
            fontWeight:FontWeight.w500,
            ), // Цвет текста
        ),
      ),
    );
  }
}