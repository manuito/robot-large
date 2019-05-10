// Here command processing with raspberryPi

// A received command = 
// 1 Letter : X
// 3 number : 123
// Then send ACK + X

// See robotoby for protocol : 
// Fixed order for attributes :
// 0]  => MOx  => Monitoring state (always on = 1)
// 1]  => MLxy => Left motor state
// 2]  => MRxy => Right motor state
// 3]  => DLxy => Left distance value
// 4]  => DRxy => Right distance value
// 5]  => RXxy => Rotation X° 0 to 360° / 180° = straight
// 6]  => RYxy => Rotation Y° 0 to 360° / 180° = straight
// 7]  => RZxy => Rotation Z° 0 to 360° / 180° = straight
// 8]  => AXxy => Accelerometer X
// 9]  => AYxy => Accelerometer Y
// 10] => AZxy => Accelerometer Z
// 11] => TPxy => Temperature °C
// 12] => BAxy => RGB Band state
// 13] => FAxy => Face state
// 14] => BExy => Beeper state

char cmdBuffer[4];
int serIn; 

bool shortMode = false;

uint8_t monitoringState;
uint8_t faceActionState;
uint8_t bandActionState;

void startSerial(){
  Serial.begin(115200);
}

void sendReady(){
  Serial.println("READY");
}

bool checkCommand(){
  if (Serial.available()) {
    digitalWrite(LED_STARTED, HIGH);   // turn the LED on (HIGH is the voltage level)
    Serial.readBytes(cmdBuffer, Serial.available());
    Serial.print("ACK");
    Serial.println(cmdBuffer[0]);
    return true;
  }
  return false;
}

void processCommand(){
  
  switch (cmdBuffer[0]) {
    case 'M': 
      monitoringState = 1;
      break;
    case 'm': 
      monitoringState = 0;
      break;
    case 'B':
      setBuzzerAction(cmdBuffer[3]);
      break;
    case 'b':
      setBuzzerAction('0');
      break;
    case 'S': 
      goStraight(toBoolean(cmdBuffer[1]), toInt(cmdBuffer[2], cmdBuffer[3]));
      break;
    case 's': 
      goStraight(true, 0);
      break;
    case 'L': 
      turnLeft(toInt(cmdBuffer[2], cmdBuffer[3]));
      break;
    case 'R': 
      turnRight(toInt(cmdBuffer[2], cmdBuffer[3]));
      break;
    case 'H':
      setFaceAction(cmdBuffer[3]);
      break;
    case 'h':
      setFaceAction('0');
      break;
    case 'F':
      setBandAction(cmdBuffer[3]);
      break;
    case 'f':
      setBandAction('0');
      break;
  }
  
  cmdBuffer[1] = 0;
  cmdBuffer[2] = 0;
  cmdBuffer[3] = 0;
}

bool toBoolean(byte val){
  return val != 48;
}

uint8_t toInt(byte first, byte second){
  return (first - 48) * 10 + (second - 48);
}

void updateActions(){
  
}

void doMonitoring(){

  if(monitoringState){
    Serial.print("MO1|ML");
    Serial.print(getLeftState());
    Serial.print("|MR");
    Serial.print(getRightState());
    Serial.print("|DL");
    Serial.print(getLeftDist());
    Serial.print("|DR");
    Serial.print(getRightDist());
    Serial.print("|RX");
    Serial.print(getAngleX());
    Serial.print("|RY");
    Serial.print(getAngleY());
    Serial.print("|RZ");
    Serial.print(getAngleZ());
    Serial.print("|AX");
    Serial.print(getAccelX());
    Serial.print("|AY");
    Serial.print(getAccelY());
    Serial.print("|AZ");
    Serial.print(getAccelZ());
    Serial.print("|TP");
    Serial.print(getTemperature());
    Serial.print("|BA");
    Serial.print(getBandAction());
    Serial.print("|FA");
    Serial.print(getFaceAction());
    Serial.print("|BE");
    Serial.println(getBuzzerAction());
  }
}

