// See http://stackoverflow.com/a/12191469 for communicating using bash.

int ledPin = 13, relayPin = 10, relayStatus = 0, OPEN = 0, CLOSED = 1;

void setup()
{
  Serial.begin(9600);
  pinMode(ledPin, OUTPUT);
  pinMode(relayPin, OUTPUT);
}

void loop()
{
  // wait for PC to send something
  while(Serial.available() == 0);

  // read what was sent
  int val = Serial.read() - '0';

  // switch over val and decide what to do
  switch (val) {
  case 0:
    openRelay();
    break;
  case 1:
    closeRelay();
    break;
  default:
    Serial.println("Invalid!");
    break;
  }

  // flush the serial as we only care about one char
  while(Serial.available() > 0) {
    Serial.read();
  }
}

void closeRelay()
{
  if (relayStatus == CLOSED)
  {
    Serial.println("The relay is already closed");
    return;
  }
  relayStatus = CLOSED;
  digitalWrite(ledPin, HIGH);
  digitalWrite(relayPin, HIGH);
  Serial.println("The relay is now closed");
}

void openRelay()
{
  if (relayStatus == OPEN)
  {
    Serial.println("The relay is already open");
    return;
  }
  relayStatus = OPEN;
  digitalWrite(ledPin, LOW);
  digitalWrite(relayPin, LOW);
  Serial.println("The relay is now open");
}

