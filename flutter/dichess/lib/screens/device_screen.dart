import 'dart:async';
import 'dart:typed_data';

import 'package:flutter/material.dart';
import 'package:flutter_bluetooth_serial/flutter_bluetooth_serial.dart';
import 'package:flutter_chess_board/flutter_chess_board.dart';

import 'package:dichess/bluetoothpb/bluetoothpb.pb.dart';
import 'package:dichess/util/length_decoder.dart';

class DeviceScreen extends StatefulWidget {
  const DeviceScreen({Key key, this.device}) : super(key: key);

  final BluetoothDevice device;

  @override
  _DeviceState createState() => _DeviceState();
}

class _DeviceState extends State<DeviceScreen> {
  StreamSubscription<Uint8List> _streamSubscrition;
  ChessBoardController _chessController;
  ChessBoard _chessBoard;

  @override
  void initState() {
    super.initState();

    _chessController = ChessBoardController();
    _chessBoard = ChessBoard(
      enableUserMoves: false,
      chessBoardController: _chessController,
    );

    BluetoothConnection.toAddress(widget.device.address).then((connection) {
      connection.output.add(Uint8List.fromList("something somethhing".runes.toList()));


      _streamSubscrition = connection.input.transform(StreamTransformer.fromBind(lengthDecoder)).listen((r) {
        var response = Response.fromBuffer(r);
        print(response);

        if (response.hasChessBoard()) {
          setState(() {
            print(_chessController.game.load(response.chessBoard.fen));
            print(_chessController.refreshBoard());
          });
        }
      });

      // Disconnected
      _streamSubscrition.onDone(() {
        Navigator.pop(context);
      });
    });
  }

  @override
  void dispose() {
    _streamSubscrition?.cancel();

    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(widget.device.name),
      ),
      body: _chessBoard,
    );
  }
}