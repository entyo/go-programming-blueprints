$(function() {
  var socket = null;
  var msgBox = $("#chatbox textarea");
  var messages = $("#messages");
  $("#chatbox").submit(function() {
    if (!msgBox.val()) return false;
    if (!socket) {
      alert("Error: Failed to make connection by WebSocket");
      return false;
    }
    socket.send(msgBox.val());
    msgBox.val("");
    return false;
  });
  if (!window["WebSocket"]) {
    alert("Error: Your browser doesn't support WebSocket");
  } else {
    socket = new WebSocket("ws://localhost:8080/room");
    socket.onclose = function() {
      alert("Connection has been closed");
    }
    socket.onmessage = function(e) {
      messages.append("<li>").text(e.data);
    }
  }
});
