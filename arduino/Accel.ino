// Accel
MPU6050 mpu6050(Wire);

void startAccel(){
  // Start accel (I2C)
  Wire.begin();
  mpu6050.begin();
}

void zeroAccel(){
  mpu6050.calcGyroOffsets(300);
}

void updateAccel(){
  mpu6050.update();
}

uint8_t getAngleX(){
  return (uint8_t) (mpu6050.getAngleX() + 180);
}

uint8_t getAngleY(){
  return (uint8_t) (mpu6050.getAngleY() + 180);
}

uint8_t getAngleZ(){
  return (uint8_t) (mpu6050.getAngleZ() + 180);
}

uint8_t getAccelX(){
  // For now : disabled
  // return mpu6050.getAccX() * 1000;
  return 0;
}

uint8_t getAccelY(){
  // For now : disabled
  // return mpu6050.getAccY() * 1000;
  return 0;
}

uint8_t getAccelZ(){
  // For now : disabled
  // return mpu6050.getAccZ() * 1000;
  return 0;
}

float getTemperature() {
  return (uint8_t) mpu6050.getTemp();
}

