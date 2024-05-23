import 'package:flutter/material.dart';

class ChatBubble extends StatefulWidget {
  final String imageUrl;
  final String chatTitle;
  final String last_mes;
  const ChatBubble({
    required this.imageUrl,
    required this.chatTitle,
    super.key, required this.last_mes,
  });

  @override
  State<StatefulWidget> createState() => _ChatBubble();
}

class _ChatBubble extends State<ChatBubble> {
  Widget _buildThumbnailImage() {
    return SizedBox(
      width: 64,
      height: 64,
      child: ClipRRect(
        borderRadius: BorderRadius.circular(32),
        child: Image.network(
          widget.imageUrl,
          fit: BoxFit.fill,
          errorBuilder: (
            BuildContext context,
            Object exception,
            StackTrace? stackTrace,
          ) {
            return CircleAvatar(
              radius: 32,
              backgroundColor: Theme.of(context).colorScheme.primaryContainer,
              child: Text(widget.chatTitle[0]),
            );
          },
        ),
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(left: 12, right: 12, top: 12, bottom: 12),
      padding: const EdgeInsets.only(left: 3, right: 3),
      alignment: Alignment.center,
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(8),
      ),
      child: Row(
        children: [
          _buildThumbnailImage(),
          const Padding(padding: EdgeInsets.only(right: 12)),
          Expanded(
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(
                  widget.chatTitle,
                  style: TextStyle(
                    color: Theme.of(context).colorScheme.onPrimaryContainer,
                    fontSize: 20,
                  ),
                  overflow: TextOverflow.ellipsis,
                ),
                Text(
                  widget.last_mes,
                  style: TextStyle(
                    color: Colors.grey,
                    fontSize: 16,
                  ),
                  overflow: TextOverflow.ellipsis,
                ),
                // Другие элементы, такие как последнее сообщение или время, могут быть добавлены здесь
              ],
            ),
          ),
          const Padding(padding: EdgeInsets.only(right: 12)),
        ],
      ),
    );
  }
}
