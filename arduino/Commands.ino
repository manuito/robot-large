// Here command processing with raspberryPi

// A received command = 
// 1 Letter : X
// 3 number : 123
// Then send ACK + X

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
      doBeep();
      break;
    case 'T':
      doSingMarioTheme();
      break;
    case 'W':
      doSingUnderworldTheme();
      break;
    case 'S': 
      if(shortMode){goStraight(true, 40);}
      else{goStraight(toBoolean(cmdBuffer[1]), toInt(cmdBuffer[2], cmdBuffer[3]));}
      break;
    case 's': 
      goStraight(true, 0);
      break;
    case 'L': 
      if(shortMode){turnLeft(30);}
      else{turnLeft(toInt(cmdBuffer[2], cmdBuffer[3]));}
      break;
    case 'R': 
      if(shortMode){turnRight(30);}
      else{turnRight(toInt(cmdBuffer[2], cmdBuffer[3]));}
      break;
    case 'F':
      if(shortMode){setFaceAction('1');}
      else{setFaceAction(cmdBuffer[3]);}
      break;
    case 'E':
      setFaceAction('2');
      break;
    case 'f':
      setFaceAction('0');
      break;
    case 'U':
      if(shortMode){setBandAction('1');}
      else{setBandAction(cmdBuffer[3]);}
      break;
    case 'K':
      setBandAction('2');
      break;
    case 'u':
      setBandAction('0');
      break;
    case 'H':
      shortMode = true;
      break;
    case 'h':
      shortMode = false;
      cmdBuffer[1] = 0;
      cmdBuffer[2] = 0;
      cmdBuffer[3] = 0;
      break;
  }
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
    Serial.print("|AX");
    Serial.print(getAccelX());
    Serial.print("|AY");
    Serial.print(getAccelY());
    Serial.print("|AZ");
    Serial.print(getAccelZ());
    Serial.print("|GX");
    Serial.print(getGyroX());
    Serial.print("|GY");
    Serial.print(getGyroY());
    Serial.print("|GZ");
    Serial.print(getGyroZ());
    Serial.print("|TP");
    Serial.println(getTemperature());
  }
}

