// LED Matrix
LedControl lc = LedControl(LED_MRX_DATA_IN,LED_MRX_CLK,LED_MRX_CS,1);

bool faceSwap = false;
uint8_t faceCpt = 0;
uint8_t faceCurrentAction;

void startFace(){
  // Wakeup LED matrix
  lc.shutdown(0,false);
  lc.setIntensity(0,1);
  lc.clearDisplay(0);
}

void setFaceAction(byte action){
  faceOff();
  faceCurrentAction = action - 48;
  faceCpt = 0;
}

uint8_t getFaceAction(){
  return faceCurrentAction;
}

void doFaceAction(){
  
   switch (faceCurrentAction) {
    case 1: 
      lookRightOrLeft();
      break;
    case 2: 
      displayEarthBeat();
      break;
    case 4:
      displayLightLoad();
      break;
    default:
      break;
  }
}
void doFaceShortAction(){
  
   switch (faceCurrentAction) {
    case 3: 
      displayRadar();
      break;
  }
}

void faceOff(){
  lc.clearDisplay(0);
}

void displayEarthBeat(){
  if(faceSwap){
    lc.setRow(0,0,B00000000);
    lc.setRow(0,1,B00011100);
    lc.setRow(0,2,B00111100);
    lc.setRow(0,3,B01111000);
    lc.setRow(0,4,B01111000);
    lc.setRow(0,5,B00111100);
    lc.setRow(0,6,B00011100);
    lc.setRow(0,7,B00000000);
  } else {
    lc.setRow(0,0,B00011110);
    lc.setRow(0,1,B00111111);
    lc.setRow(0,2,B01111110);
    lc.setRow(0,3,B11111000);
    lc.setRow(0,4,B11111000);
    lc.setRow(0,5,B01111110);
    lc.setRow(0,6,B00111111);
    lc.setRow(0,7,B00011110);
  }
  faceSwap = !faceSwap;
}

void lookRightOrLeft(){
  if(faceSwap){
    lc.setRow(0,0,B00000000);
    lc.setRow(0,1,B00101010);
    lc.setRow(0,2,B01001110);
    lc.setRow(0,3,B01000000);
    lc.setRow(0,4,B01000000);
    lc.setRow(0,5,B01001010);
    lc.setRow(0,6,B00101110);
    lc.setRow(0,7,B00000000);
  } else {
    lc.setRow(0,0,B00000000);
    lc.setRow(0,1,B00101110);
    lc.setRow(0,2,B01001010);
    lc.setRow(0,3,B01000000);
    lc.setRow(0,4,B01000000);
    lc.setRow(0,5,B01001110);
    lc.setRow(0,6,B00101010);
    lc.setRow(0,7,B00000000);
  }
  faceSwap = !faceSwap;
}

// Max 5
void displayLoadingFace(uint8_t stepLoad){
   switch (stepLoad) {
    case 1: 
      lc.setRow(0,0,B00000000);
      lc.setRow(0,1,B11000011);
      lc.setRow(0,2,B10100111);
      lc.setRow(0,3,B10011111);
      lc.setRow(0,4,B10011111);
      lc.setRow(0,5,B10100111);
      lc.setRow(0,6,B11000011);
      lc.setRow(0,7,B00000000);
      lc.setRow(0,0,B00000000);
      break;
    case 2: 
      lc.setRow(0,0,B00000000);
      lc.setRow(0,1,B11000011);
      lc.setRow(0,2,B10100111);
      lc.setRow(0,3,B10111011);
      lc.setRow(0,4,B10111011);
      lc.setRow(0,5,B10100111);
      lc.setRow(0,6,B11000011);
      lc.setRow(0,7,B00000000);
      break;
    case 3: 
      lc.setRow(0,1,B11000011);
      lc.setRow(0,2,B11100101);
      lc.setRow(0,3,B11111001);
      lc.setRow(0,4,B11111001);
      lc.setRow(0,5,B11100101);
      lc.setRow(0,6,B11000011);
      lc.setRow(0,7,B00000000);
      break;
   }
}

void displayLightLoad() {

  // SetLed KO ?
  switch (faceCpt) {
    case 1: 
      lc.setRow(0,0,B00000000);
      lc.setRow(0,1,B00000000);
      lc.setRow(0,2,B00000000);
      lc.setRow(0,3,B00010000);
      lc.setRow(0,4,B00000000);
      lc.setRow(0,5,B00000000);
      lc.setRow(0,6,B00000000);
      lc.setRow(0,7,B00000000);
      break;
    case 2: 
      lc.setRow(0,0,B00000000);
      lc.setRow(0,1,B00000000);
      lc.setRow(0,2,B00000000);
      lc.setRow(0,3,B00001000);
      lc.setRow(0,4,B00000000);
      lc.setRow(0,5,B00000000);
      lc.setRow(0,6,B00000000);
      lc.setRow(0,7,B00000000);
      break;
    case 3: 
      lc.setRow(0,0,B00000000);
      lc.setRow(0,1,B00000000);
      lc.setRow(0,2,B00000000);
      lc.setRow(0,3,B00000000);
      lc.setRow(0,4,B00001000);
      lc.setRow(0,5,B00000000);
      lc.setRow(0,6,B00000000);
      lc.setRow(0,7,B00000000);
      break;
    case 4: 
      lc.setRow(0,0,B00000000);
      lc.setRow(0,1,B00000000);
      lc.setRow(0,2,B00000000);
      lc.setRow(0,3,B00000000);
      lc.setRow(0,4,B00010000);
      lc.setRow(0,5,B00000000);
      lc.setRow(0,6,B00000000);
      lc.setRow(0,7,B00000000);
      break;
   }
   
  faceCpt++;
  if(faceCpt > 4){
    faceCpt = 1;
  }
}

void displayRadar(){

  if(getLeftDist() < 30){
    if(getLeftDist() < 20){
      if(getLeftDist() < 10){
        lc.setRow(0,5,B11110000);
        lc.setRow(0,6,B11111100);
        lc.setRow(0,7,B11111110);
      } else {
        lc.setRow(0,5,B11000000);
        lc.setRow(0,6,B11110000);
        lc.setRow(0,7,B11111000);
      }
    } else {
      lc.setRow(0,5,B00000000);
      lc.setRow(0,6,B11000000);
      lc.setRow(0,7,B11100000);
    }
  } else {
      lc.setRow(0,5,B00000000);
      lc.setRow(0,6,B00000000);
      lc.setRow(0,7,B00000000);
  }
  
  if(getRightDist() < 30){
    if(getRightDist() < 20){
      if(getRightDist() < 10){
        lc.setRow(0,2,B11110000);
        lc.setRow(0,1,B11111100);
        lc.setRow(0,0,B11111110);
      } else {
        lc.setRow(0,2,B11000000);
        lc.setRow(0,1,B11110000);
        lc.setRow(0,0,B11111000);
      }
    } else {
      lc.setRow(0,2,B00000000);
      lc.setRow(0,1,B11000000);
      lc.setRow(0,0,B11100000);
    }
  } else {
      lc.setRow(0,2,B00000000);
      lc.setRow(0,1,B00000000);
      lc.setRow(0,0,B00000000);
  }
}
