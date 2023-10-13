program HappyBirthdayOlesya;

uses crt, sysutils;

procedure PrintWithAnimation(text: string);
var i: integer;
begin
  for i := 1 to Length(text) do
  begin
    write(text[i]);
    Delay(100);
  end;
end;

procedure HappyBirthdayMessage(name: string);
begin
  clrscr;
  PrintWithAnimation('Happy Birthday, ' + name + '!');
  writeln;
  writeln;
  Delay(1000);
  PrintWithAnimation('Wishing you a day filled with joy and happiness!!!');
  Delay(1000);
  writeln;
  writeln;
end;

begin
  clrscr;
  HappyBirthdayMessage('<NAME>');
  readln;
end.
