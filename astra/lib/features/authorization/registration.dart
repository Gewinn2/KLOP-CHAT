import 'package:flutter/material.dart';

class Registration extends StatefulWidget {
  const Registration({super.key});


  @override
  _Registration createState() => _Registration();
}

class _Registration extends State<Registration> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(  
        automaticallyImplyLeading: false, 
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
                  width: 550,
                  padding: EdgeInsets.symmetric(horizontal: 16, vertical: 8), // Добавьте отступы внутри контейнера
                  decoration: BoxDecoration(
                    color: Color.fromRGBO(236, 223, 245, 1), // Фон контейнера
                    borderRadius: BorderRadius.all(Radius.circular(16)), // Закругленные углы
                  ),
                  child: Text(
                    "Регистрация в Klop Chat",
                    textAlign: TextAlign.center,
                    style: TextStyle(
                      fontFamily: 'Inter',
                      fontSize: 32,
                      color: Color.fromRGBO(59, 3, 102, 1),
                      fontWeight: FontWeight.w700,
                    ),
                  ),
                ),
                SizedBox(height: 70,),
                Text( 
                        "Ваше имя",
                        textAlign: TextAlign.start,
                        style: TextStyle( 
                          color: Color.fromRGBO(59, 3, 102, 1),
                          fontFamily: 'Inter',
                          fontSize: 16,
                          fontWeight: FontWeight.w700,
                        ),
                      ),   
                SizedBox(height: 20,),
                Container(
                  width: 400,
                  height: 45,
                  child: TextField(
                          decoration: InputDecoration(
                            floatingLabelBehavior: FloatingLabelBehavior.always,
                            labelText: 'Ваше имя',
                            border: OutlineInputBorder(
                              borderRadius: BorderRadius.circular(10),
                            ),
                          ),
                          autofocus: true,
                        ),
                ),
                SizedBox(height: 20,),
                Text( 
                        "Ваша почта",
                        textAlign: TextAlign.start,
                        style: TextStyle( 
                          color: Color.fromRGBO(59, 3, 102, 1),
                          fontFamily: 'Inter',
                          fontSize: 16,
                          fontWeight: FontWeight.w700,
                        ),
                      ),   
                SizedBox(height: 20,),
                Container(
                  width: 400,
                  height: 45,
                  child: TextField(
                          decoration: InputDecoration(
                            floatingLabelBehavior: FloatingLabelBehavior.always,
                            labelText: 'Почта',
                            border: OutlineInputBorder(
                              borderRadius: BorderRadius.circular(10),
                            ),
                          ),
                          autofocus: true,
                        ),
                ),
                SizedBox(height: 20,),
                Text( 
                        "Придумайте пароль для аккаунта",
                        textAlign: TextAlign.start,
                        style: TextStyle( 
                          color: Color.fromRGBO(59, 3, 102, 1),
                          fontFamily: 'Inter',
                          fontSize: 16,
                          fontWeight: FontWeight.w700,
                        ),
                      ),   
                SizedBox(height: 20,),
                Container(
                  width: 400,
                  height: 45,
                  child: TextField(
                          decoration: InputDecoration(
                            floatingLabelBehavior: FloatingLabelBehavior.always,
                            labelText: 'Пароль',
                            border: OutlineInputBorder(
                              borderRadius: BorderRadius.circular(10),
                            ),
                          ),
                          autofocus: true,
                        ),
                ),
                SizedBox(height: 20,),
                Text( 
                        "Повторите пароль",
                        textAlign: TextAlign.start,
                        style: TextStyle( 
                          color: Color.fromRGBO(59, 3, 102, 1),
                          fontFamily: 'Inter',
                          fontSize: 16,
                          fontWeight: FontWeight.w700,
                        ),
                      ),   
                SizedBox(height: 20,),
                Container(
                  width: 400,
                  height: 45,
                  child: TextField(
                          decoration: InputDecoration(
                            floatingLabelBehavior: FloatingLabelBehavior.always,
                            labelText: 'Пароль',
                            border: OutlineInputBorder(
                              borderRadius: BorderRadius.circular(10),
                            ),
                          ),
                          autofocus: true,
                        ),
                ),
                SizedBox(height: 70,),
                ElevatedButton(
                      onPressed: () {
                        Navigator.popAndPushNamed(
                            context,
                            '/',
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
                        "Зарегистрироваться",
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
        ),
      ),
    );
  }
}



