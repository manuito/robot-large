// HS Sensors

#define MAX_PING_TIMEOUT 10000 // 10 ms ~ 150 cm

uint8_t leftDist, rightDist;
bool distSwap = false;

const float travelTimeRatio = 0.034/2 ;

void startHsSensors() {
  pinMode(HC_LEFT_TRIG_PIN, OUTPUT);
  pinMode(HC_LEFT_ECHO_PIN, INPUT);
  pinMode(HC_RIGHT_TRIG_PIN, OUTPUT);
  pinMode(HC_RIGHT_ECHO_PIN, INPUT);
  digitalWrite(HC_LEFT_TRIG_PIN, LOW);
  digitalWrite(HC_RIGHT_TRIG_PIN, LOW);
}


void updateDistanceLeft(){
  echoAndPing(HC_LEFT_TRIG_PIN, HC_LEFT_ECHO_PIN, &leftDist);
}

void updateDistanceRight(){
  echoAndPing(HC_RIGHT_TRIG_PIN, HC_RIGHT_ECHO_PIN, &rightDist);
}


void echoAndPing(uint8_t trigPin, uint8_t echoPin, uint8_t* dist){
  
  // Echo
  digitalWrite(trigPin, HIGH);
  delayMicroseconds(10);
  digitalWrite(trigPin, LOW);
  
  // Ping
  long duration = pulseIn(echoPin, HIGH, MAX_PING_TIMEOUT);

  if(duration > 0 ){
    *dist = (uint8_t) (duration * travelTimeRatio);
  }
}

void updateDistances(){
  if(distSwap){
    updateDistanceLeft();
  } else {
    updateDistanceRight();
  }
  distSwap = !distSwap;
}

/**
 * For monitoring
 */
 
uint8_t getLeftDist(){
  return leftDist;
}

uint8_t getRightDist(){
  return rightDist;
}

