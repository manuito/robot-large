int leftSpeed, rightSpeed ;

/*
 * Provided commands for motor
 */

void setupMotors(){
  // Motors
  pinMode(LEFT_DIR_PIN, OUTPUT);
  pinMode(LEFT_DRV_PIN, OUTPUT);
  pinMode(RIGHT_DIR_PIN, OUTPUT);
  pinMode(RIGHT_DRV_PIN, OUTPUT);
}

void goStraight(bool direction, uint8_t speed){
  driveLeft(direction, speed);
  driveRight(direction, speed);
}

void turnLeft(uint8_t speed){
  driveLeft(true, speed);
  driveRight(false, speed);
}

void turnRight(uint8_t speed){
  driveLeft(false, speed);
  driveRight(true, speed);
}

/**
 * Internal commands
 */

void driveLeft(bool direction, uint8_t speed){ 

  // Apply command
  digitalWrite(LEFT_DIR_PIN, direction);
  analogWrite(LEFT_DRV_PIN, speed);

  // Update state
  if(direction){
    leftSpeed = speed;
  } else {
    leftSpeed = 0-speed;    
  }
}

void driveRight(bool direction, uint8_t speed){ 

  // Apply command
  digitalWrite(RIGHT_DIR_PIN, direction);
  analogWrite(RIGHT_DRV_PIN, speed);

  // Update state
  if(direction){
    rightSpeed = speed;
  } else {
    rightSpeed = -speed;    
  }  
}

// For monitoring
int getLeftState(){
  return leftSpeed;
}

// For monitoring
int getRightState(){
  return rightSpeed;
}

