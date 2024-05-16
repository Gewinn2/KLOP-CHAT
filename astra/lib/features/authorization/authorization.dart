import 'package:astra/features/authorization/password.dart';
import 'package:flutter/material.dart';

class Authorization extends StatefulWidget {
  const Authorization({super.key});


  @override
  _AuthorizationState createState() => _AuthorizationState();
}

class _AuthorizationState extends State<Authorization> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(   
        title: const Text(
          "Klop Chat",
          style: TextStyle( 
            color: Color.fromRGBO(59, 3, 102, 1),
            fontSize: 40,
            fontFamily: 'Inter',
            fontWeight: FontWeight.w700,
          )
          ),
        ),
      body: Center(
        child: Container(
          width: 600, // Ширина квадрата
          height: 700, // Высота квадрата
          child: Container(
            margin: const EdgeInsets.all(20), // Отступы внутреннего контейнера
            decoration: BoxDecoration(
              //color: Theme.of(context).colorScheme.primary, // Цвет фона внутреннего контейнера - белый
              borderRadius: BorderRadius.circular(10), // Закругленные края внутреннего контейнера
            ),
            child: Column(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Container(
                  padding: EdgeInsets.symmetric(horizontal: 16, vertical: 8), // Добавьте отступы внутри контейнера
                  decoration: BoxDecoration(
                    color: Color.fromRGBO(236, 223, 245, 1), // Фон контейнера
                    borderRadius: BorderRadius.all(Radius.circular(16)), // Закругленные углы
                  ),
                  child: Text(
                    "Добро пожаловать в Klop Chat!",
                    style: TextStyle(
                      fontFamily: 'Inter',
                      fontSize: 32,
                      color: Color.fromRGBO(59, 3, 102, 1),
                      fontWeight: FontWeight.w700,
                    ),
                  ),
                ),
                SizedBox(height: 70,),
                Container(
                  width: 400,
                  child: Column( 
                    children: [
                      TextField(
                        decoration: InputDecoration(
                          floatingLabelBehavior: FloatingLabelBehavior.always,
                          labelText: 'Логин',
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(10),
                          ),
                        ),
                        autofocus: true,
                      ),
                      SizedBox(height: 20),
                      PasswordTextField(),
                      SizedBox(height: 70,),
                      ElevatedButton(
                      onPressed: () {
                        Navigator.pushNamed(
                            context,
                            'main_screen_chat',
                            //arguments: data, // Здесь вы передаете данные, которые хотите передать на `ChatsGroupList()`
                          );
                      },
                      style: ButtonStyle(
                        backgroundColor: MaterialStateProperty.all<Color>(Color.fromRGBO(59, 3, 102, 1)), // Фоновый цвет кнопки
                        shape: MaterialStateProperty.all<RoundedRectangleBorder>(
                          RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(10.0), // Радиус закругления краев
                          ),
                        ),
                        minimumSize: MaterialStateProperty.all(Size(400, 50)),
                      ),
                      child: Text(
                        "Вход",
                        style: TextStyle(
                          color: Colors.white, // Цвет текста
                          fontSize: 20,
                          fontWeight: FontWeight.w400,
                        ),
                      ),
                    ),
                        SizedBox(height: 40,),
                      Text( 
                        "У Вас еще нет аккаунта в Klop Chat?",
                        style: TextStyle( 
                          color: Color.fromRGBO(59, 3, 102, 1),
                          fontFamily: 'Inter',
                          fontSize: 16,
                          fontWeight: FontWeight.w700,
                        ),
                      ),
                      SizedBox(height: 20,),
                      ElevatedButton(
                      onPressed: () {
                        // Действия при нажатии на кнопку
                      },
                      style: ButtonStyle(
                        backgroundColor: MaterialStateProperty.all<Color>(Color.fromRGBO(59, 3, 102, 1)), // Фоновый цвет кнопки
                        shape: MaterialStateProperty.all<RoundedRectangleBorder>(
                          RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(10.0), // Радиус закругления краев
                          ),
                        ),
                        minimumSize: MaterialStateProperty.all(Size(400, 50)),
                      ),
                      child: Text(
                        "Регистрация",
                        style: TextStyle(
                          color: Colors.white, // Цвет текста
                          fontSize: 20,
                          fontWeight: FontWeight.w400,
                        ),
                      ),
                    ),
                    ],
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}


