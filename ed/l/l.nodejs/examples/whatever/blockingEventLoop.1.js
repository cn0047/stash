// Js is single thread and this thread is busy with infinite loop...
while (true) {
  setTimeout(() => console.log(100), 0);
}
